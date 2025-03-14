# opentelemetry-collector-sakuracloud

OpenTelemetry Collector Receiver for SAKURA Cloud

# Example Configuration for OpenTelemetry Collector

```yaml
receivers:
  sakuracloud:
    collection_interval: 10s
    access_token: ${env:SAKURACLOUD_ACCESS_TOKEN}
    access_token_secret: ${env:SAKURACLOUD_ACCESS_TOKEN_SECRET}
    metrics:
      sakuracloud.server.up:
        enabled: true
      sakuracloud.server.cpu_time:
        enabled: true
      sakuracloud.server.network_interface.send:
        enabled: true
      sakuracloud.server.network_interface.receive:
        enabled: true
      sakuracloud.server.disk.read:
        enabled: true
      sakuracloud.server.disk.write:
        enabled: true

exporters:
  debug:
    verbosity: detailed
#  otlp/mackerel:
#    endpoint: otlp.mackerelio.com:4317
#    compression: gzip
#    headers:
#      Mackerel-Api-Key: ${env:MACKEREL_APIKEY}    

service:
  pipelines:
    metrics:
      receivers: [sakuracloud]
      exporters: [debug]
      #exporters: [otlp/mackerel]
```

## License

`sacloud/opentelemetry-collector-sakuracloud` Copyright (C) 2024-2025 [The sacloud/opentelemetry-collector-sakuracloud Authors](AUTHORS).

This project is published under [Apache 2.0 License](LICENSE).
