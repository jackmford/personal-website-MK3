---
title: "The SLO Review"
date: "2025-01-24"
slug: "the-slo-review"
blurb: "Ensuring You Don't Set It and Forget It"
---

A persistent thorn in your side when implementing service level objectives (SLOs) will be the revisit strategy. As companies scale to a level that inhibits them from having embedded site reliability engineers (SREs) working closely with application teams day-to-day on reliability needs, it inevitably falls to the application team to own their reliability strategy.

In an environment like this, how do we ensure SLOs defined by (or defined for—[see inherited vs. developed reliability strategies](https://becomingreliable.substack.com/p/inherited-vs-developed-reliability)) teams stay fresh? Engineering teams juggle a thousand priorities, and SLOs can easily get moved into the "done" column and forgotten. [Since reliability is perhaps the most important feature of any system](https://sre.google/workbook/reaching-beyond/), how do we ensure the goals around an application's reliability get the attention they deserve?

### Strategies for Keeping Your SLOs Fresh 🍅:

1. **Bake Them into Your Planning Ceremonies**

- Regardless of whether your team uses Scrum, Agile, or Waterfall, you likely plan your work somehow. If, for example, you operate on a two-week sprint, make it habit to review your SLOs during planning. Has your service been performing in accordance with your goals? If it hasn’t, now is the perfect time to allocate some capacity to figure out why. You’re already planning, so leverage that context to prioritize reliability.

2. **Set Up a Recurring Meeting on a Cadence**

- New SLOs require more care and feeding, consider scheduling a meeting every couple of weeks. As they mature, you can extend your revisit schedule. The specific frequency matters less than consistently dedicating time to evaluate performance, targets, and trends. Entropy comes for everything, your SLOs will go stale if you don't work on them. Build it into your team’s rhythm so it doesn’t fall through the cracks.

3. **Configure Simple Notifications When You’re Out of Budget**

- When SLOs are new, you probably set the wrong target. It isn't your fault, it's difficult to know how reliable you have been in the past if you've never measured it! Set up an alert to notify your team when you’ve run out of error budget. This notification will, at a minimum, keep your SLO on your mind if it is underperforming your target. I recommend pairing it with one of the two above strategies. Over time, this alert should mature as your team's confidence in the SLO target grows. In the early stages, though, simplicity keeps the team plugged into their SLOs.

### Staying Plugged In

Regardless of how you implement these strategies, the main idea is to configure a simple system that ensures regular attention to your SLOs. Prioritization is hard; we can't always keep all of the balls in the air. These systems at least automate _some_ of that away difficult decision making away. They ensure reliability remains a focal point, rather than an afterthought.
