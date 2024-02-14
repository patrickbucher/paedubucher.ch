---
title: Testing Is Not About Programming
subtitle: Programming vs. Software Engineering
author: Patrick Bucher
date: 2018-12-16T12:00:00+0100
lang: en
---

Programming is the process of writing code in order to solve a problem.
Software Engineering is programming, combined with the factors time and other
programmers.

Tests cannot prove the correctness of a program. A test case makes sure that a
couple of hand-picked values from the input set are mapped correctly to a
couple of hand-picked values from the solution set. As Dijkstra wrote, in order
to get a program right, forget about the individual items of the set and work
with the set definition—which is quite the contrary of writing test cases. So
tests are not the right tool when it comes to getting programs right in the
first place.

However, tests can be very helpful in the area of Software Engineering. There, they
aren't used to proof the correctness of a program, but to make sure that a test
covered program doesn't change its test-covered behaviour as time goes, and
other programmers work on it.
