// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package sakuracloudreceiver // import "github.com/sacloud/opentelemetry-collector-sakuracloud/receiver/sakuracloudreceiver"

import (
	"context"
	"fmt"
	"time"

	client "github.com/sacloud/api-client-go"
	"github.com/sacloud/opentelemetry-collector-sakuracloud/receiver/sakuracloudreceiver/internal/metadata"
	"github.com/sacloud/opentelemetry-collector-sakuracloud/receiver/sakuracloudreceiver/internal/sacloud"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/scraper/scrapererror"
	"go.uber.org/zap"
)

// sakuraCloudScraper handles scraping of RabbitMQ metrics
type sakuraCloudScraper struct {
	logger   *zap.Logger
	cfg      *Config
	settings component.TelemetrySettings
	mb       *metadata.MetricsBuilder

	sacloudClient *sacloud.Client
}

// newScraper creates a new scraper
func newScraper(logger *zap.Logger, cfg *Config, settings receiver.Settings) *sakuraCloudScraper {
	return &sakuraCloudScraper{
		logger:   logger,
		cfg:      cfg,
		settings: settings.TelemetrySettings,
		mb:       metadata.NewMetricsBuilder(cfg.MetricsBuilderConfig, settings),
	}
}

// start starts the scraper by creating a new HTTP Client on the scraper
func (s *sakuraCloudScraper) start(ctx context.Context, host component.Host) error {
	httpClient, err := s.cfg.ClientConfig.ToClient(ctx, host, s.settings)
	if err != nil {
		return err
	}

	clientOptions := &client.Options{
		AccessToken:       string(s.cfg.AccessToken),
		AccessTokenSecret: string(s.cfg.AccessTokenSecret),
		Gzip:              true,
		HttpClient:        httpClient,
		UserAgent:         fmt.Sprintf("opentelemetry-collector-receiver-sakuracloud/%s", Version),
	}

	sacloudClient, err := sacloud.NewClient(s.cfg.Endpoint, s.cfg.Zones, clientOptions)
	if err != nil {
		return err
	}

	s.sacloudClient = sacloudClient
	return nil
}

// scrape collects metrics from the RabbitMQ API
func (s *sakuraCloudScraper) scrape(ctx context.Context) (pmetric.Metrics, error) {
	ts := pcommon.NewTimestampFromTime(time.Now())
	errs := &scrapererror.ScrapeErrors{}

	s.scrapeServerMetrics(ctx, ts, errs)

	return s.mb.Emit(), errs.Combine()
}

func (s *sakuraCloudScraper) scrapeServerMetrics(ctx context.Context, ts pcommon.Timestamp, errs *scrapererror.ScrapeErrors) {
	servers, err := s.sacloudClient.Server.Find(ctx)
	if err != nil {
		errs.AddPartial(1, err)
		return
	}

	mcfg := s.cfg.Metrics

	// TODO workerプールを導入する
	for _, server := range servers {
		up := 0
		if server.InstanceStatus.IsUp() {
			up = 1
		}
		s.mb.RecordSakuracloudServerUpDataPoint(ts, int64(up), server.ID.String(), server.Name, server.ZoneName)

		if !server.InstanceStatus.IsUp() {
			continue
		}

		if mcfg.SakuracloudServerCPUTime.Enabled {
			v, err := s.sacloudClient.Server.MonitorCPU(ctx, server.ZoneName, server.ID, ts.AsTime())
			if err != nil {
				errs.AddPartial(1, err)
				return
			}
			cpuTimeMilliSecPerCore := v.CPUTime / float64(server.CPU) * 1000 // コア数で割る、かつ単位をmsに
			s.mb.RecordSakuracloudServerCPUTimeDataPoint(ts, cpuTimeMilliSecPerCore,
				server.ID.String(), server.Name, server.ZoneName)
		}

		if mcfg.SakuracloudServerNetworkInterfaceSend.Enabled || mcfg.SakuracloudServerNetworkInterfaceReceive.Enabled {
			for i, nic := range server.Interfaces {
				v, err := s.sacloudClient.Server.MonitorNIC(ctx, server.ZoneName, nic.ID, ts.AsTime())
				if err != nil {
					errs.AddPartial(1, err)
					return
				}
				if mcfg.SakuracloudServerNetworkInterfaceSend.Enabled {
					send := v.Send * 8
					s.mb.RecordSakuracloudServerNetworkInterfaceSendDataPoint(ts, send,
						server.ID.String(), server.Name, server.ZoneName,
						nic.ID.String(), int64(i))
				}
				if mcfg.SakuracloudServerNetworkInterfaceReceive.Enabled {
					receive := v.Receive * 8
					s.mb.RecordSakuracloudServerNetworkInterfaceReceiveDataPoint(ts, receive,
						server.ID.String(), server.Name, server.ZoneName,
						nic.ID.String(), int64(i))
				}
			}
		}

		if mcfg.SakuracloudServerDiskRead.Enabled || mcfg.SakuracloudServerDiskWrite.Enabled {
			for i, disk := range server.Disks {
				v, err := s.sacloudClient.Server.MonitorDisk(ctx, server.ZoneName, disk.ID, ts.AsTime())
				if err != nil {
					errs.AddPartial(1, err)
					return
				}
				if mcfg.SakuracloudServerDiskRead.Enabled {
					read := v.Read
					s.mb.RecordSakuracloudServerDiskReadDataPoint(ts, read,
						server.ID.String(), server.Name, server.ZoneName,
						disk.ID.String(), int64(i))
				}
				if mcfg.SakuracloudServerDiskWrite.Enabled {
					write := v.Write
					s.mb.RecordSakuracloudServerDiskWriteDataPoint(ts, write,
						server.ID.String(), server.Name, server.ZoneName,
						disk.ID.String(), int64(i))
				}
			}
		}
	}
}
