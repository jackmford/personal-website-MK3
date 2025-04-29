---
title: "Error Budget + Burn Rate Cheatsheet"
date: "2025-04-28"
slug: "error-budget-burn-rate-cheatsheet"
blurb: "My go to cheat sheet when I inevitably forget error budget math!"
---

## Formulas

### 1. Error Rate

Measures the proportion of "bad" events relative to total events over a period.

**Formula:**
`Error Rate = Bad Events / Total Events`

**Usage:**
Used to determine the amount of *time* or *budget* burned during a specific window.

---

### 2. Burn Rate

Measures how *fast* you are consuming your error budget relative to the allowed rate based on your SLO period. A burn rate of `1` means you're consuming budget exactly on track; higher numbers mean faster consumption.

**Formula:**
`Burn Rate = Observed Error Rate / (1 - SLO Target)`

**Where:**
- `Observed Error Rate`: Actual error rate measured (e.g., 0.1 for 10%).
- `SLO Target`: Target service level objective (e.g., 0.99 for 99%).
- `(1 - SLO Target)`: Total Error Budget expressed as a percentage (e.g., 0.01 for 1%).

**Example:**
- Observed Error Rate = 10% (0.1)
- SLO Target = 99% (0.99)
- Calculation:  
  `Burn Rate = 0.1 / (1 - 0.99) = 0.1 / 0.01 = 10`
- **Interpretation:** You are burning through your budget **10x faster** than allowed over the SLO period.

---

### 3. Budget Consumed During Alerting Window (%)

Calculates what percentage of your *total* error budget (for the *entire* SLO period) is consumed if the observed burn rate persists for the duration of your *alerting window*.

**Formula:**  
`% Budget Consumed = Burn Rate * (Alerting Window Duration / SLO Period Duration) * 100`

**Where:**
- `Burn Rate`: As calculated above.
- `Alerting Window Duration`: Duration your alert evaluates (e.g., 60 minutes).
- `SLO Period Duration`: Full SLO period (e.g., 43,200 minutes for 30 days).

> **⚠️ Make sure units match between Alerting Window and SLO Period!**

**Example:**
- Burn Rate = 10
- Alerting Window = 60 minutes
- SLO Period = 30 days (43,200 minutes)
- Calculation:  
  `% Budget Consumed = 10 * (60 / 43200) * 100 = 1.388%`
- **Interpretation:** If a burn rate of 10 persists for 60 minutes, **1.39%** of the 30-day error budget is consumed.

**Calculating Time Burned from %:**
- Total Budget Time = 432 minutes (for 99% SLO over 30 days)
- `Time Burned = Total Budget Time * (% Budget Consumed / 100)`
- Example:  
  `Time Burned = 432 * 0.01388 = 6 minutes`

> **💡 Note:**  
> Substituting the burn rate formula gives:  
> `(Observed Error Rate / Total Error Budget %) * (Alerting Window Duration / SLO Period Duration) * 100`  
> This connects the observed error rate to budget consumption over time.

---

### 4. Required Burn Rate for Theoretical Budget Consumption

Calculates the `Burn Rate` you should alert on to catch a certain % of error budget burned within an alerting window. (Rearranges Formula #3.)

_This answers: "How fast am I burning if I consume X% of my budget within Y window?"_

**Formula:**  
`Required Burn Rate = (% Budget To Consume / 100) * (SLO Period Duration / Alerting Window Duration)`

**Where:**
- `% Budget To Consume`: Desired % of budget to trigger an alert (e.g., 2%).
- `SLO Period Duration`: Full SLO measurement period.
- `Alerting Window Duration`: Time window for the alert.

**Key Concepts:**
- The "ideal" percentage to burn per window is `1 / (number of windows)`.
- Anything over that is a *multiplier* relative to ideal consumption.

**Example (7-Day SLO):**
- Goal: Alert when 10% of budget is consumed within 1 hour.
- SLO Period = 7 days × 24 hours = 168 hours
- Alerting Window = 1 hour
- Calculation:  
  `Required Burn Rate = (10 / 100) * (168 / 1) = 16.8`
- **Interpretation:** Set an alert to fire if burn rate exceeds **16.8** over a 1-hour window.

**Example Walkthrough (7-Day SLO, 99% Target):**
- SLO Target = 99% ➔ Error Budget = 1% (0.01)
- SLO Period = 168 hours
- Total Budget Time = `0.01 * 168 = 1.68 hours`
- Goal: Alert if 10% of 1.68 hours = 0.168 hours are consumed in 1 hour.
- Observed Error Rate needed:  
  `0.168 / 1 = 0.168` (16.8%)
- Checking Required Burn Rate:  
  `0.168 / 0.01 = 16.8` — matches!

---

## Calculation Examples

### Scenario Setup

- **SLO Target:** 99% (Error Budget = 1% or 0.01)
- **SLO Period:** 30 days (43,200 minutes)
- **Total Error Budget Time:** `0.01 * 43200 = 432 minutes`
- **Alerting Window:** 60 minutes

---

### Example 1: High Error Rate (3%)

- Observed Error Rate = 3% (0.03)
- **Time Burned in Window:** `0.03 * 60 = 1.8 minutes`
- **% Budget Burned in Window:** `(1.8 / 432) * 100 = 0.416%`
- **Burn Rate:** `0.03 / 0.01 = 3`
- **Verify with Burn Rate:**  
  `3 * (60 / 43200) * 100 = 0.416%` (Matches!)

**Interpretation:**  
An error rate of 3% leads to a burn rate of 3. We consume 1.8 minutes (0.42%) of the total monthly budget in that hour.

---

### Example 2: Steady Low Error Rate (0.01%)

- Observed Error Rate = 0.01% (0.0001)
- **Time Burned in Window:** `0.0001 * 60 = 0.006 minutes`
- **Burn Rate:** `0.0001 / 0.01 = 0.01`
- **Verify with Burn Rate:**  
  `0.01 * (60 / 43200) * 100 = 0.00139%`

**Interpretation:**  
A very low error rate results in a negligible burn rate (<1) and minimal budget consumption.

---

### Example 3: Sustained High Error Rate (10%) — How Long Until Budget Depleted?

- Observed Error Rate = 10% (0.1)
- **Burn Rate:** `0.1 / 0.01 = 10`
- **Time to Deplete Budget:**  
  `SLO Period Duration / Burn Rate`
- Calculations:  
  `30 days / 10 = 3 days`  
  or  
  `43200 minutes / 10 = 4320 minutes = 72 hours = 3 days`

**Interpretation:**  
A 10% sustained error rate would deplete the entire 30-day budget in just 3 days.

---

### Example 4: Using Required Burn Rate Method

- Goal: How fast to burn 10% of budget in 10 hours?
- **Calculation:**  
  `Required Burn Rate = (10 / 100) * (43200 / 600) = 0.1 * 72 = 7.2`

**Interpretation:**  
A burn rate of **7.2** would burn 10% of the monthly budget over a 10-hour window.

---

