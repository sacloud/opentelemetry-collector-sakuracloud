// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package sakuracloudreceiver // import "github.com/sacloud/opentelemetry-collector-sakuracloud/receiver/sakuracloudreceiver"

import (
	"context"
	"errors"
	"time"

	"github.com/sacloud/iaas-api-go"
	"github.com/sacloud/opentelemetry-collector-sakuracloud/receiver/sakuracloudreceiver/internal/metadata"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/confighttp"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/receiver/scraperhelper"
	"go.opentelemetry.io/collector/scraper"
)

var errConfigNotSakuraCloud = errors.New("config was not a SAKURA Cloud receiver config")

// NewFactory returns the component factory for the sakuracloudreceiver
func NewFactory() receiver.Factory {
	return receiver.NewFactory(
		metadata.Type,
		createDefaultConfig,
		receiver.WithMetrics(createMetricsReceiver, metadata.MetricsStability),
	)
}

func createMetricsReceiver(
	_ context.Context,
	params receiver.Settings,
	rConf component.Config,
	consumer consumer.Metrics,
) (receiver.Metrics, error) {
	cfg, ok := rConf.(*Config)
	if !ok {
		return nil, errConfigNotSakuraCloud
	}

	sakuraCloudScraper := newScraper(params.Logger, cfg, params)
	s, err := scraper.NewMetrics(sakuraCloudScraper.scrape, scraper.WithStart(sakuraCloudScraper.start))
	if err != nil {
		return nil, err
	}

	return scraperhelper.NewScraperControllerReceiver(&cfg.ControllerConfig, params, consumer, scraperhelper.AddScraper(metadata.Type, s))
}

func createDefaultConfig() component.Config {
	return &Config{
		ControllerConfig: scraperhelper.ControllerConfig{
			CollectionInterval: time.Minute,
		},
		ClientConfig: confighttp.ClientConfig{},
		Zones:        iaas.SakuraCloudZones,
		Endpoint:     iaas.SakuraCloudAPIRoot,
	}
}
