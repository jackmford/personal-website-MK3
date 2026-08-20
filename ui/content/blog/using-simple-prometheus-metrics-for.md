---
title: "Using Simple Prometheus Metrics for Easy SLO Instrumentation"
date: "2024-12-25"
slug: "using-simple-prometheus-metrics-for"
blurb: "Setting up Service Level Objectives (SLOs) might sound complex at first, but as with anything, it pays to start simple."
---

Setting up Service Level Objectives (SLOs) might sound complex at first, but as with anything, it pays to start simple. If you have a Prometheus instance, you can quickly implement simple and effective SLOs using Prometheus [counters](https://prometheus.io/docs/concepts/metric_types/#counter). A couple of counters are enough to track errors, successful requests, and total requests for creating SLOs.

This article walks you through setting up basic SLOs using Prometheus counters, providing examples and practical advice to get started.

## **Instrumentation**

To illustrate, imagine a web page handler tracking user interactions. We use two Prometheus counters: one increments each time a request is made, and the other increments when a request is successfully handled and sent to the browser. Here’s how you can define these counters in Go:

```
var successCounter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "home_success_request_count",
		Help: "Number of requests successfully handled by the home handler",
	},
)

var totalCounter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "home_total_request_count",
		Help: "Total number of requests handled by the home handler",
	},
)
```

These counters can be registered and exported using one of the [Prometheus client libraries](https://prometheus.io/docs/instrumenting/clientlibs/). Once the metrics are available in Prometheus, you can use them to define simple SLO calculations.

## **Defining SLOs**

### **Success Rate SLO**

Using the success request counter and the total request counter in combination with the [Prometheus increase query function](https://prometheus.io/docs/prometheus/latest/querying/functions/#increase), you can calculate the success rate over a given period. For example, to determine the success rate over the past seven days:

```
increase(home_success_request_count[7d]) / increase(home_total_request_count[7d]) * 100
```

This formula provides the percentage of successful requests over the last seven days. You can set a target SLO based on this value, such as:

> **99% of requests must be successful in the last 7 days.**

Graphing this metric in tools like Grafana makes it easier to monitor your target and identify trends.

### **Error Rate and Burn Rate**

To calculate the error rate, subtract the successful request counter from the total request counter and divide the result by the total request counter:

```
(increase(home_total_request_count[7d]) - increase(home_success_request_count[7d])) / increase(home_total_request_count[7d]) * 100
```

This gives the error rate percentage. From here, you can explore burn rate calculations for alerting and monitoring. Burn rates help you measure how quickly you’re consuming your error budget and can provide actionable insights for incident management.

To learn more about burn rate alerting, check out my article [Error Budgets and Alerting](https://substack.com/home/post/p-153549568?source=queue).

## **Starting Somewhere**

Prometheus counters provide a straightforward way to begin tracking and monitoring SLOs. With just two counters, you can measure success rates, error rates, and calculate burn rates for better alerting, which can completely transform team culture and quality of life for on-call engineering teams. The most daunting step of any journey is the first, always make it as simple as you can.
