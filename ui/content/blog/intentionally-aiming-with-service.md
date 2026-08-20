---
title: "Intentionally Aiming With Service Level Objectives"
date: "2024-12-17"
slug: "intentionally-aiming-with-service"
blurb: "“If you can't measure it, you can't improve it.”"
---

“If you can’t measure it, you can’t improve it.”

— Peter Drucker

It is 3:00 AM and the pager goes off. Everything is on fire. Engineers, managers, and directors get called in. You spend the next several days trying to answer critical questions: What went wrong? How did this happen? How many customers were impacted, and for how long? Postmortems are scheduled and action items are assigned and implemented. Day by day, things return to the routine, and that 3:00 AM wake-up call becomes a distant memory. But key details fade: How long were customers impacted? How many errors occurred? Incidents blur together as time passes.

Then it happens again. The cycle repeats—incident, fix, forget. While it might feel like progress is being made, there’s no way to definitively show that things **are** getting better. Without a measurable target, there’s no way to know if you’re heading in the right direction. Improvement is impossible without a goal, and this applies to more than just system reliability.

# The Power of Setting Goals

By defining and committing to measurable targets tied to key service indicators, you establish a clear aim. Tracking adherence to these targets over time reveals whether you are improving, regressing, or maintaining performance. This is the core benefit of defining Service Level Objectives (SLOs): **the ability to measure progress and direction**. Without SLOs, your understanding of critical performance metrics remains clouded. Memories are short; few can recall outages or performance dips from six months ago—even one month can be challenging.

Without this understanding, prioritization becomes guesswork. Should you focus on new features, or should you set up failover to a secondary region to ensure uptime during an outage? Without a historical understanding of your metrics and goals, it’s impossible to know.

# Take Action: Define Your Targets

Set a target. Aim at **something**. For example, you might define an SLO to keep error rates below 1% or latency from 200 responses < 500ms over a rolling 30-day period. If you’re off target, adjust your approach and begin again. This iterative process ensures continuous improvement and clear priorities.

[Read about SLOs from Google’s SRE book](https://sre.google/sre-book/service-level-objectives/).
