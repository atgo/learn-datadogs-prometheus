Expose prometheus metrics to Datadogs
====

Golang app propose 2 endpoints:

 - GET /buggy, increase buggy counter.
 - GET /metrics, expose prometheus metrics

Datadog agent v6 configured scrap app's metrics in prometheus format.

Expecting: Metrics are reported into Datadog console.

## Testing locally

- Change `DD_API_KEY` in docker-compose.yaml file
- Build golang app: `CGO_ENABLED=0 GOOS=linux go build -o app cmd/prometheus/main.go`
- docker-compose up
- check Datadog console
- docker-compose down
