# SakuraCloud Receiver

<!-- status autogenerated section -->
| Status        |           |
| ------------- |-----------|
| Stability     | [development]: metrics   |
| Distributions | [] |
| Issues        | [![Open issues](https://img.shields.io/github/issues-search/open-telemetry/opentelemetry-collector-contrib?query=is%3Aissue%20is%3Aopen%20label%3Areceiver%2Fsakuracloud%20&label=open&color=orange&logo=opentelemetry)](https://github.com/open-telemetry/opentelemetry-collector-contrib/issues?q=is%3Aopen+is%3Aissue+label%3Areceiver%2Fsakuracloud) [![Closed issues](https://img.shields.io/github/issues-search/open-telemetry/opentelemetry-collector-contrib?query=is%3Aissue%20is%3Aclosed%20label%3Areceiver%2Fsakuracloud%20&label=closed&color=blue&logo=opentelemetry)](https://github.com/open-telemetry/opentelemetry-collector-contrib/issues?q=is%3Aclosed+is%3Aissue+label%3Areceiver%2Fsakuracloud) |
| [Code Owners](https://github.com/open-telemetry/opentelemetry-collector-contrib/blob/main/CONTRIBUTING.md#becoming-a-code-owner)    | [@yamamoto-febc](https://www.github.com/yamamoto-febc) |

[development]: https://github.com/open-telemetry/opentelemetry-collector/blob/main/docs/component-stability.md#development
<!-- end autogenerated section -->

TBD overview

## Getting Started

TBD getting started

## Configuration

- `access_token`
- `access_token_secret`

### Example:

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
```
