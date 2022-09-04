---
title: One Month of SICP
subtitle: Slowly, but Steadily
author: Patrick Bucher
date: 2022-08-30T18:43:00
lang: en
---

When I finally got my bachelor's degree in Computer Science, I set another goal
for my further professional development: I wanted to learn functional
programming. I got a glimpse of Prolog and Racket during my studies, and used
higher-order functions with lambda expressions in Java and JavaScript once in a
while. I also peaked into Common LISP on multiple occasions, but I never made it
a priority.

# Erlang? No, Haskell and MOOCs!

Now I wanted to dig deeper, and Erlang was the language that interested me the
most. The main reason for this choice was the concurrency model based on actors
and messaging, and the Prolog syntax didn't bother me at all.

For some reason I got sidetracked and ended up learning a bit of Haskell. I
worked halfway through [Programming
Haskell](https://www.cambridge.org/ch/academic/subjects/computer-science/programming-languages-and-applied-logic/programming-haskell-2nd-edition),
which taught me a lot about functional programming. I skipped the second part;
probably because I got sidetracked again: I remember playing around with OpenBSD
and FreeBSD at that time. But I also spent quite some time with machine learning
in the fall and winter with Coursera MOOCs, and only got back to functional
programming when I worked through the Scala MOOC in early 2021.

I also started a MOOC called [Programming Languages Part
A](https://www.coursera.org/learn/programming-languages), which uses Standard ML
as a teaching language. This was the only MOOC I did _not_ finish—and the last
one I started to this day. I just had enough of learning with videos, it was
just too tiresome for me. So I decided to pick up a book; _the_ book: [Structure
and Interpretation of Computer
Programs](https://mitpress.mit.edu/9780262510875/structure-and-interpretation-of-computer-programs/).

# Elixir and Clojure

Unfortunately, I got stuck in the middle of the first chapter. Up to then, I
worked through all the exercises with great care. But the math probably was too
difficult for me, and I rather wanted to do something practical with functional
programming. So I left SICP unfinished and wanted to learn Erlang—finally—but
got sidetracked again: this time to its fancy little brother, Elixir.

I picked up a book—[Learn Functional Programming with
Elixir](https://pragprog.com/titles/cdc-elixir/learn-functional-programming-with-elixir/)—that
left out some of the best parts of Elixir, but had a lot of silly examples in
it. I quite liked Elixir, but the book killed my motivation. And learning some
LISP was still on my bucket list. Therefore, I tried another hosted language:
Clojure. I also had to learn some Clojure for my job in order to adjust the
[Riemann](http://riemann.io/) configuration being used to ingest metrics. I
worked through [Getting
Clojure](https://pragprog.com/titles/roclojure/getting-clojure/), which I
consider a great introduction to the language.

# Wasteful Meandering

I didn't have time for functional programming in spring/summer 2021, because I
was preparing lessons for a teaching side-job I got into on short notice.
However, I managed to work through a small eBook called [Functional Programming in
Python](https://leanpub.com/functionalprogramminginpython), whose lessons I can
apply in my day job—at least partially (no monads). Teaching was also what
occupied most of my time in fall and winter. So 2021 passed without much further
effort.

In the first half of 2022, I found some time again to improve my functional
programming skills. I picked up Elixir again; this time with the excellent book
[Elixir in
Action](https://www.manning.com/books/elixir-in-action-second-edition). However,
the book's focus is on concurrency and error handling in the second half, and I
felt that I first needed to sharpen my skills in functional programming
techniques before digging into the applications.

Back to Clojure I was, and this time I bought a lot of literature. [Clojure for
the Brave and True](https://www.braveclojure.com/) is quite witty and informal,
but focused too much on games for my taste. [Programming
Clojure](https://pragprog.com/titles/shcloj3/programming-clojure-third-edition/)
really was what I was looking for—but it was too fast-paced for me, and I
noticed that I wasn't really learning the mechanisms of the language, and I got
stuck once more.

Frustrated as I was, I went on a walk to overthink my meandering learning path
that lead me nowhere. After two years, I was able to write some of my stock
programs (Connect Four, computing league tables from game results) in Clojure
and Elixir under heavy googling, but that pretty much was it. Having turned 35
this summer, I realized that I should not waste more of my time like this. No
matter what programming language I am learning, I just need to stick with it in
order to make some real progress.

# Back to SICP

So I decided to go back to SICP, and I started completely from scratch. I
deleted my old repository from GitHub, and even from my backup server. On the
30th of July 2022, I pledged myself to work on SICP every day from now on; even
if I only read a single page or think about an exercise for five minutes. And
finally, I [managed to do
so](https://github.com/patrickbucher/sicp/blob/master/diary.md).

After a month, I'm quite far into chapter 1, which I want to finish in
September. The upcoming move to another place and my teaching duty certainly
will disrupt my streak in early October, but until then my «daily SICP» habit
will be strong enough to smoothly continue after being settled in my new home.

My plan is to at least work through chapters 1 and 2. Chapter 3, dealing with
subjects such as concurrency, streams and delayed evaluation, looks promising
for my further journey. Chapters 4 and 5 I might consider later; chapters 1-3
probably will be a strong-enough foundation to take Clojure or Elixir/Erlang
into consideration again. But for the rest of 2022, I'll focus on SICP.

## Lessons Learned

What lessons did I learn in the last 30 days?

First, hard exercises require pen and paper. The computer is just too
distracting, and formulas are way easier to understand drawn out on a sheet of
paper rather than mangled into a text file. Seeing something on paper, as
opposed to holding it in your mind, frees up resources for actually thinking
about it.

Second, if my mind goes blank when working on an exercise, then there are just
too many things I haven't understood yet. Probably I didn't fully comprehend the
examples earlier in the chapter. So back to the examples, back to pen and paper;
and only back to the exercise after I carefully worked through the examples.

Third, being unsure about my first step into a mathematical proof, I just cannot
find the willpower to step through it completely. I first need a little hint
whether or not my initial idea is valid. After looking at a single line of a
sample solution, I can find the confidence to work through the complete proof.

Fourth, thinking hard about a problem for a while, then leaving it for a day or
night, and going back to it in the evening or on the next day helps a lot. I
often got stuck in the wrong place and missed the actual problem. Getting out of
the problem and back in after a couple of hours gives you a new perspective. So
far, I was able to solve all the problems within 24 hours, but with multiple
attempts at it.
