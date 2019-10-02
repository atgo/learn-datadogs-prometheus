Expose prometheus metrics to Datadogs
====

Golang app propose 2 endpoints:

 - GET /buggy, increase buggy counter.
 - GET /metrics, expose prometheus metrics

Datadog agent v6 configured scrap app's metrics in prometheus format.

Expecting: Metrics are reported into Datadog console.
