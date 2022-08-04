# Mock-Exporter

mock-exporter allows you to mock any Prometheus exporters such as node-exporter, kube-state-metrics, etc.

## How to use

1. Configure the metrics which you want to mock in yaml file.

2. Start mock-exporter server.

Then your Prometheus can collect your metrics exported from mock-exporter server. (Make sure your Prometheus has job config to scrape your mock-exporter server.)

## License

MIT License. see [LICENSE](./LICENSE)

## Author

Teppei Sudo (https://github.com/sp-yduck)