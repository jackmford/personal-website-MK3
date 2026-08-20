---
title: "Error Budgets and Alerting"
date: "2024-12-23"
slug: "error-budgets-and-alerting"
blurb: "Emphasizing Simplicity"
---

Error budgets are confusing. Personally, anytime I need to remind myself of one of their intricacies, I end up Googling or writing out a formula for longer than I should need to. Working with development teams that are brand new to error budgets adds a whole extra set of hurdles. My goal here is to simplify the essential formulas and concepts as succinctly as possible, creating a reference for myself and others.

### **What Is an Error Budget?**

An error budget is a bucket of acceptable bad behavior. When you agree to a Service Level Objective (SLO) target, you automatically get your bucket full of acceptable bad behavior. For example, if your SLO target is 99% of responses returning a 200 status code, you’re allowed 1% of responses to **not** return a 200. That’s your error budget. If you use up the entire bucket of bad behavior, your SLO is violated.

#### **Total Error Budget Calculation**

**Error Budget = 100% - SLO Target (in percentage form)**

### **Understanding Burn Rate**

The most confusing aspect of an error budget is the burn rate—the rate at which you are consuming the bucket of bad behavior. Imagine your bucket is full of snacks, and someone is eating them. A high burn rate means they’re real hungry, chowing down on snacks, depleting them quickly. A low burn rate means maybe the bucket is out on the dining room table and they grab a snack here or there throughout the day. Ideally, your error budget burns at a steady rate and lasts precisely as long as your SLO window.

In reality, bad behavior comes in peaks and valleys. Sometimes you’ll have leftover budget, and sometimes you’ll use too much and violate your SLO target.

#### **Current Burn Rate**

**Current Burn Rate = Error Rate / (1 - SLO Target)**

Example: For an error rate of 10% (0.1) with a 99% SLO target:

_Current Burn Rate = 0.1 / (1 - 0.99) = 10_

#### **Maximum Burn Rate**

**Max Burn Rate = 1 / (1 - SLO Target)**

Example: For a 99% SLO target:

_Max Burn Rate = 1 / (1 - 0.99) = 100_

This means that if all behavior within the SLO is bad (an error rate of 100%), the burn rate would hit 100.

**Tip:** When configuring alerts, avoid setting thresholds higher than the max burn rate. Such thresholds will never be reached and render the alert ineffective.

#### **Time to Error Budget Exhaustion**

**Time to Exhaustion = SLO Window / Burn Rate**

Examples:

- If your SLO window is 7 days and your burn rate is 1: _7 / 1 = 7 days_
- If your burn rate is 10: _7 / 10 = 0.7 days (16.8 hours)_

[![](https://substackcdn.com/image/fetch/%24s_!56Mo!,w_1456,c_limit,f_auto,q_auto:good,fl_progressive:steep/https%3A%2F%2Fsubstack-post-media.s3.amazonaws.com%2Fpublic%2Fimages%2F6181b4c1-d69c-4491-902f-102467fdcf05_2400x1200.png)](https://substackcdn.com/image/fetch/%24s_!56Mo!,f_auto,q_auto:good,fl_progressive:steep/https%3A%2F%2Fsubstack-post-media.s3.amazonaws.com%2Fpublic%2Fimages%2F6181b4c1-d69c-4491-902f-102467fdcf05_2400x1200.png)

#### **Error Budget Consumed**

**Error Budget Consumed = (SLO Window * 24 hours * Budget Consumed) / Alerting Window**

Example: For a 7-day SLO window and 5% budget consumption over 1 hour:

_(7 * 24 * 0.05) / 1 = 8.4_

This means burning at a rate of 8.4 over one hour consumes 5% of your error budget.

## Alerting Recommendations

For baseline values, refer to [Google’s recommendations](https://sre.google/workbook/alerting-on-slos/#recommended_parameters_for_an_slo_based_a). These assume a 99.9% alerting target but can be adjusted to match your own error budget consumption calculations.

### Pairing Long and Short Alerting Windows

Using both long and short alerting windows reduces alert reset times after resolving an issue. For instance, if you set an alert for 5% error budget consumption over one hour, adding a shorter alert (e.g., 5 minutes) ensures alerts stop firing quickly once the issue is resolved.

### Keep It Simple

When configuring alerts, simplicity is key. Overly complex alert configurations that confuse your team are counterproductive. Start with straightforward alerts that everyone understands. An alert with clear logic is far more valuable than an intricate system that loses your team in complexity.

---

Error budgets and alerting are powerful tools for managing service reliability. By prioritizing clarity and simplicity, you can maximize their effectiveness and ensure your team is on the same page.
