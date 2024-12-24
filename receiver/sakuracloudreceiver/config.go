// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package sakuracloudreceiver // import "github.com/sacloud/opentelemetry-collector-sakuracloud/receiver/sakuracloudreceiver"

import (
	"github.com/sacloud/opentelemetry-collector-sakuracloud/receiver/sakuracloudreceiver/internal/metadata"
	"go.opentelemetry.io/collector/config/confighttp"
	"go.opentelemetry.io/collector/config/configopaque"
	"go.opentelemetry.io/collector/receiver/scraperhelper"
)

// Config holds all the parameters to start an HTTP server that can be sent logs from CloudFlare
type Config struct {
	confighttp.ClientConfig        `mapstructure:",squash"`
	scraperhelper.ControllerConfig `mapstructure:",squash"`
	metadata.MetricsBuilderConfig  `mapstructure:",squash"`

	AccessToken       configopaque.String `mapstructure:"access_token"`
	AccessTokenSecret configopaque.String `mapstructure:"access_token_secret"`
	Zones             []string            `mapstructure:"zones"`
	Endpoint          string              `mapstructure:"endpoint"`
}

func (c *Config) Validate() error {
	return nil
}
