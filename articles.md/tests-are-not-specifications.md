---
title: Tests Are Not Specifications
subtitle: Boomers vs. the (All Too) Silent Generation
author: Patrick Bucher
date: 2021-05-02T10:00:00
lang: en
---

In _The Clean Coder_, Chapter 7 (_Acceptance Tests_), Robert C. Martin writes
(p. 109):

> But the real reason these tests aren't redundant is that their primary
> function _is not testing_. The fact that they are tests is incidental. […]
> The fact that they automatically verify the design, structure, and behaviour
> that they specify is wildly useful, **but the specification is their true
> purpose**. [Bold emphasis mine.]

OK, boomer, here's your spec:

    assert f(0) == 1
    assert f(-1) == -2
    assert f(+1) == 4
    assert f(+4) == 13
    assert f(-9) == -26

Now implement `f(x)`. The requirements are clear, aren't they?

You're still here? Where's the implementation!? I told you everything you need.
You got test cases! They _are_ the specification; you just told me this, and of
course, _Uncle Bob is always right_, they told me in school.

Still not done? OK, you're fired. Here's the Python implementation. Easy, isn't
it?

    def f(x):
        return (x + 2) * 3 - 5

Those young developers are no good, after all. Let's hire one of those old
programmers, please bring in that old, more-arrogant-than-Uncle-Bob, grumpy
Dijkstra guy ([EWD
981](https://www.cs.utexas.edu/users/EWD/transcriptions/EWD09xx/EWD981.html)):

> At this stage I can give you some behavioural advice. Contrary to what
> misguided—and misguiding—educationists may have told you, don't waste your
> time looking at specific examples. Trying to come to grips with a large set by
> looking at a few—hopefully representative—elements of that set is one of the
> most ineffective ways of spending one's time. Go immediately for the general
> case, or, in other words, **attack the set not by looking at specific elements
> but by analysing the definition of the set.** [Bold emphasis mine.]

So you're not interested in some randomly picked elements of both the problem
and solution set, but in those set's definition? Is that all you need? Well,
here's that definition:

    y = f(x) = (x + 2) * 3 - 5

But I can also provide you with a set of test cases, or _specs_, as they are
called by some developers. Oh, you don't need them? But what is QA supposed to
do all day long then? You mean we can fire them? Wow, that sounds great. You're
hired, Mr. Dijkstra!

Did I hear this right, Mr. Dijkstra, _tests are not specifications_? Now you're
talking…
