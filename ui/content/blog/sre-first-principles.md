---
title: "The SRE First Principles Problem"
date: "2026-02-07"
slug: "sre-first-principles"
blurb: "Nebulous, ever changing, and difficult to solve"
---

Large enterprises weren’t built in a day. Like Rome (you know the saying), they are built brick by brick, feature by feature, over years, decades, in some cases centuries. If only we could be a fly on the wall for all of the architectural design sessions over those years. We all occasionally fall into the trap of looking at a decision made before our time in a place and scratch our heads, frustrated, thinking, “... Okaaay, but why?”.

The reality is, those choices in the moment (in the general case) were probably smart, sound decisions. The engineers that came before you were just as smart as you are, operating with the context and constraints they had. They had their reasons.

But those reasons pile up. Decisions are made, behavioral patterns harden, and before you know it, you have a jumbled mess that lacks any standardization across the software landscape. Business Unit A operates differently than Business Unit B because of a choice that was made before half the workforce was even onboarded.

This becomes a salient problem when platform engineering or central site reliability teams begin to scale up. Everything is easy in a bubble. To accomplish anything at true scale, the challenge becomes one of _standardization_ as much as, or more than, an engineering problem.

Imagine you want to spin up an enterprise-wide solution for displaying service level objective numbers because that would be a cool thing to have and management would love it, great! This sounds like a cool way for everyone to have wide visibility into how their peers (and their dependencies) are performing against their SLO targets. The SaaS landscape is full of tools that can accomplish this ask, if you don’t already have one (which you probably do). There are even LLMs to help you write whatever templating language or code you need to spin up the dashboards and store them in source.

You begin. Unit A’s data is clean and tidy. Easy to integrate with. Using their data as a proving ground for your idea was a sweet way to prove out your idea. You’re good to continue building, as some initial value has been seen in your design.

Moving along to Business Unit B. You talk to some folks, understand where their resources are and how they are organized. Their data is living in a legacy system that for some nebulous (likely valid at the time the decision was made) reason, they must stay integrated with it. Okay, you can build around that.

You discover next that their applications have a different naming scheme than Business Unit A’s. They prepend their app IDs to their app names instead of relying solely on an ID system. You need to build around this as well. You continue this for five more business units and you’ve sunk way more time into this project than you initially planned, discovering that even within individual business units things differ from area to area.

Your SLO product launches, but must exist with caveats: “We can’t support that area in this view because their data comes from X not Y.” Your observability posture at the company suffers because of it. Your central engineering team spends more time managing configuration drift across the enterprise than building solutions to reliability problems.

If you are trying to build out your SLO or SRE program and Business Unit A has a completely different definition and idea of what an “application” even is, your problems go far beyond just implementing the observability tool. You’ve run into a first principles issue. There is no common understanding on what the most basic resources in the enterprise even are. If that is the case, every standard or enterprise solution you try to build becomes infinitely more complex and tangled.

As we hear time and time again, SRE is as much a human problem as it is a technology problem. I’d wager that most SREs are going to have an easier time setting up a log ingestion pipeline that can handle massive scales of records than getting a large enterprise to completely overhaul their enterprise data model. Before you jump into building enterprise observability solutions, take a deep breath, try to relax (these are hard problems), and then go talk to your business units.
