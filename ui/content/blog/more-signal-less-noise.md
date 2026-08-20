---
title: "More Signal, Less Noise"
date: "2025-02-04"
slug: "more-signal-less-noise"
blurb: "Don't blast your devs with context-poor instrumentation"
---

In working at large enterprises, many of us are living in what I refer to as an "inherited" reliability environment. This concept - which I've explored in a [previous post](https://becomingreliable.substack.com/p/inherited-vs-developed-reliability) - attempts to describe organizations in which observability and reliability resources are often managed by teams far removed from those who actually use them to troubleshoot applications. Application engineers may be expected to use telemetry or tools to understand their reliability that they themselves did not configure and do not utilize on a day to day basis. While this is a problem in and of itself; we need approaches and strategies to work within this problem.

In chapter one of "Seeking SRE" by David N. Blank-Edelman, Coburn Watson - an accomplished reliability engineering leader - explains context vs. control in SRE organizations. Watson gives details on how Netflix uses a mainly context driven strategy to help engineering teams with their reliability. If an application is experiencing issues or degraded performance, they aim to ensure the engineers related to that application are getting the proper amount of context _about_ the issue to know how to troubleshoot it.

In a control heavy model, engineers may be informed about their issue or degraded performance with "punitive action", as Watson says. The amount of context and control that an organization requires often changes depending on the type of service they provide. For instance, software related to human safety likely leans more toward a control heavy approach.

Consider the difference in these two alerts:

```
Alert: High Latency
Service: Landing Page
Status: Warning
Timestamp: 2024-02-03 14:22:17
```

```
Alert: Landing Page Performance Degradation
Severity: High Latency Warning
Timestamp: 2024-02-03 14:22:17

Details:
- Current P99 Latency: 1.2s (Threshold: 500ms)
- Affected Endpoint: /user/profile
- Impacted Users: Approximately 35% of active sessions
- Related Traces: [Link to distributed tracing]
- Recent Deployments: Backend update at 14:15

Business Impact:
- Estimated user experience degradation
- Potential revenue loss: $X,XXX per hour

Visualization:
- [Direct link to real-time performance dashboard]
```

Regardless of the organizational model, the tools and telemetry utilized by engineering teams _especially at an organization with an inherited reliability strategy_ should not be light on context. Context is the information that allows engineers to understand what the data that their platform or SRE team spun up for them means and how to act on it.

When working in an inherited model and enriching applications with logs, service level objectives, traces, etc. for other engineering teams, the goal is to ensure that when alerts fire, and they will, the related data provides engineers with enough context to guide them toward the correct resolution.

Meaningful context such as:

- What specifically is the behavior causing this alert to fire?
- How long has this been happening?
- Where is it coming from? i.e. what app or service?
- Where can you go to see it in a visualization?
- How does this impact end-users? (And it _should_ be impacting your end user if you're getting alerted on it.)

The configuration of reliability solutions by platform teams, SRE teams, and operations teams on behalf of engineering organizations can so easily become a quest of trying to ensure everything is "instrumented" while failing to prioritize the collective understanding amongst everyone of _why_ it is being configured and how to use it when it is needed.

To mitigate this, when instrumenting observability in systems you do not "own", imagine anything you instrument is for a brand new engineer going through their first on-call experience. Documentation should be comprehensive and clear to guide the engineers who may be utilize it in critical situations.

Observability and monitoring solutions are only effective if they can be understood by those who rely on them.

I highly recommend reading the full chapter in [Seeking SRE](https://www.oreilly.com/library/view/seeking-sre/9781491978856/).
