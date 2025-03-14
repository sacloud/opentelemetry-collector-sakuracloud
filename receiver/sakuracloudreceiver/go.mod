module github.com/sacloud/opentelemetry-collector-sakuracloud/receiver/sakuracloudreceiver

go 1.22.0

require (
	github.com/google/go-cmp v0.6.0
	github.com/sacloud/api-client-go v0.2.10
	github.com/sacloud/iaas-api-go v1.14.0
	github.com/sacloud/packages-go v0.0.10
	github.com/sacloud/webaccel-api-go v1.2.0
	github.com/stretchr/testify v1.10.0
	go.opentelemetry.io/collector/component v0.116.1-0.20241220212031-7c2639723f67
	go.opentelemetry.io/collector/component/componenttest v0.116.1-0.20241220212031-7c2639723f67
	go.opentelemetry.io/collector/config/confighttp v0.116.0
	go.opentelemetry.io/collector/config/configopaque v1.22.1-0.20241220212031-7c2639723f67
	go.opentelemetry.io/collector/confmap v1.22.0
	go.opentelemetry.io/collector/consumer v1.22.1-0.20241220212031-7c2639723f67
	go.opentelemetry.io/collector/consumer/consumertest v0.116.1-0.20241220212031-7c2639723f67
	go.opentelemetry.io/collector/pdata v1.22.1-0.20241220212031-7c2639723f67
	go.opentelemetry.io/collector/receiver v0.116.1-0.20241220212031-7c2639723f67
	go.opentelemetry.io/collector/receiver/receivertest v0.116.1-0.20241220212031-7c2639723f67
	go.opentelemetry.io/collector/scraper v0.116.0
	go.uber.org/zap v1.27.0
)

require (
	github.com/benbjohnson/clock v1.3.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/fatih/structs v1.1.0 // indirect
	github.com/felixge/httpsnoop v1.0.4 // indirect
	github.com/fsnotify/fsnotify v1.8.0 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-viper/mapstructure/v2 v2.2.1 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/hashicorp/errwrap v1.0.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/hashicorp/go-retryablehttp v0.7.5 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/compress v1.17.11 // indirect
	github.com/knadh/koanf/maps v0.1.1 // indirect
	github.com/knadh/koanf/providers/confmap v0.1.0 // indirect
	github.com/knadh/koanf/v2 v2.1.2 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pierrec/lz4/v4 v4.1.21 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rs/cors v1.11.1 // indirect
	github.com/sacloud/go-http v0.1.8 // indirect
	go.opentelemetry.io/collector/client v1.22.1-0.20241220212031-7c2639723f67 // indirect
	go.opentelemetry.io/collector/config/configauth v0.116.0 // indirect
	go.opentelemetry.io/collector/config/configcompression v1.22.0 // indirect
	go.opentelemetry.io/collector/config/configtelemetry v0.116.1-0.20241220212031-7c2639723f67 // indirect
	go.opentelemetry.io/collector/config/configtls v1.22.1-0.20241220212031-7c2639723f67 // indirect
	go.opentelemetry.io/collector/config/internal v0.116.0 // indirect
	go.opentelemetry.io/collector/consumer/consumererror v0.116.1-0.20241220212031-7c2639723f67 // indirect
	go.opentelemetry.io/collector/consumer/xconsumer v0.116.1-0.20241220212031-7c2639723f67 // indirect
	go.opentelemetry.io/collector/extension v0.116.0 // indirect
	go.opentelemetry.io/collector/extension/auth v0.116.0 // indirect
	go.opentelemetry.io/collector/pdata/pprofile v0.116.1-0.20241220212031-7c2639723f67 // indirect
	go.opentelemetry.io/collector/pipeline v0.116.1-0.20241220212031-7c2639723f67 // indirect
	go.opentelemetry.io/collector/receiver/xreceiver v0.116.1-0.20241220212031-7c2639723f67 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.56.0 // indirect
	go.opentelemetry.io/otel v1.32.0 // indirect
	go.opentelemetry.io/otel/metric v1.32.0 // indirect
	go.opentelemetry.io/otel/sdk v1.32.0 // indirect
	go.opentelemetry.io/otel/sdk/metric v1.32.0 // indirect
	go.opentelemetry.io/otel/trace v1.32.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	go.uber.org/ratelimit v0.3.0 // indirect
	golang.org/x/crypto v0.31.0 // indirect
	golang.org/x/net v0.33.0 // indirect
	golang.org/x/sys v0.28.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20241015192408-796eee8c2d53 // indirect
	google.golang.org/grpc v1.69.0 // indirect
	google.golang.org/protobuf v1.36.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
