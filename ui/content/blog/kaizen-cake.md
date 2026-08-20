---
title: "Kaizen Cake"
date: "2025-12-03"
slug: "kaizen-cake"
blurb: "Integrating reliability iteratively"
---

I’ll occasionally bust out a weekend project with some unfamiliar tech stack to get my feet wet with something new if I begin to feel stagnant with my current batch of tools. If the project doesn’t stick, I move on and leave it in the graveyard of my Github repositories. Every now and again, I’ll revisit one of these “set it and forget it” projects, realizing that in most cases, while I picked a few things up during the process, maybe a surface level familiarity with whatever technology it was I was exploring, I didn’t really, truly, _learn_ anything.

Enterprises, teams, departments, whatever level of organizational structure it may be, often attempt a similar approach when it comes to establishing a reliability practice. They are told the promise that if you can just get this base level of observability established, you have the golden ticket. So, naturally, they will set aside a block of time to do just that, and likely do it well. They will get the right metrics, right traces, and right logs all in place. They will set up some service level objective targets and configure their alerting. And then they’ll move on. Back to “normal” day to day operations, writing features, fixing application bugs, writing tests.

When the time comes and shit hits the fan, the trap that I too encounter with my dead weekend projects will be sprung. By working in a big “lift and shift” manner, there is no familiarity that grows over time with the technologies or implementations that were conquered during that short stint of effort. Regardless of how sound the implementation is, if a group is not intimately familiar with their reliability posture, in the heat of an incident they will not magically gain that needed familiarity.

Srihari Sriraman outlined an idea in his article [The common sense unit of work](https://blog.nilenso.com/blog/2025/09/17/the-common-sense-unit-of-work/) that a shippable unit of work for a product shouldn’t be an entire feature, those are too big. They need to be broken down into smaller pieces and provide immediate value to the customer. He frames his common sense unit of work as a slice of cake that contains multiple different layers of work, for example UI/UX, data and database work, pipelines, etc. Instead of tackling **any one of those** all at once, instead, tackle a subsection of each to ship out and provide value as quickly as you can. All we need to do to avoid our earlier treacherous trap, is add reliability focused work as a layer of the cake.

> _**Kaizen**_ is a Japanese concept in business studies which asserts that significant positive results may be achieved due the cumulative effect of many, often small (and even trivial), improvements to all aspects of a company’s operations.

Kaizen became popular in the business world for being an effective approach to improving business operations in highly complex enterprises. Sriraman’s common sense unit of work brought it to mind immediately. By taking an iterative approach to reliability work, as Kaizen does to business operations, context over time about how the observability posture is built around the corresponding service can be gained and learned at a much more natural and appropriate pace for us as humans.

It’s well known that repetition over time when learning new material is incredibly important. The [forgetting curve](https://en.wikipedia.org/wiki/Forgetting_curve), a hypothesis from the 1880s by Herman Ebbinghaus which was more recently replicated with similar findings, attempts to quantify how quickly we forget new information when we put no effort into attempting to retain it. According to the curve, up to 90% of newly learned information is forgotten within one month of learning it if it isn’t revisited.

Building a reliability habit into your normal working cadence is incredibly important for being well equipped for incidents. Things will go wrong, change inevitably introduces instability into a system, and when they do, having a gradually, iteratively grown understanding of how your observability systems exist and work alongside your business applications will outperform a fuzzy idea gained from a short burst of work every time. You don’t show up on race day and magically rise to the occasion if you haven’t put in the work day by day all of those months leading up to it.

As James Clear said, **“You do not rise to the level of your goals. You fall to the level of your systems.”**
