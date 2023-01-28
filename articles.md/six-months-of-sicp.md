---
title: 'Six Months of SICP'
subtitle: 'Still Slowly, Still Steadily'
author: 'Patrick Bucher'
date: 2023-01-28T22:00:00
lang: en
---

Last summer, I decided to work through SICP, at least through the first two
chapters, and probably also through the third. I wrote about my progress [after
the first month](./2022-08-30-one-month-of-sicp.html). Back then, I established a daily
habit, but was also a bit pessimistic whether or not I'd stick to it. I worked
three days for the company, teached two days a week, and on the weekends, I had
to prepare the lessons for the next week. I also knew that I was moving to a new
place in October. So how could I possibly continue with my daily efforts?

To this day, I didn't miss a single day. Even on the day of our move, I was able
to [do an
exercise](https://github.com/patrickbucher/sicp/blob/master/diary.md#2022-10-04-tu).
I installed MIT Scheme, Emacs, and Geiser on a laptop a couple of days before,
and I pushed my commit through a hotspot provided by my smartphone. I kept using
this setup until a proper internet connection was established in my new flat.

Fortunately, I wasn't sick for a single day, besides the occasional headaches.
But since I worked on SICP as the first thing in the morning, nothing could
really get in its way. I made it a priority, and so it worked out fine. However,
I sometimes had to work on SICP in the train on my way to work or back home, but
I did at least one commit every day. Sometimes it was just an entry into my
diary that I read a couple of pages, but usually I also committed code. On some
exercises, I had to work many days, sometimes about a week or so. But the most
important thing was to start the editor, to look at the code, and at least to
try it.

# Challenges

The second chapter had fewer mathematical proofs and exercises where I needed to
trace long chains of function calls. It was very hard, nonetheless.

- Some examples and exercises were about image processing. They were not too
  difficoult, after all, but required quite some additional work. First, the
  setup of the image library wasn't documented. This is actually quite
  understandable, because the implementation from the 80s and 90s probably
  wouldn't have worked on a modern machine anyway. I figured out a way to do the
  exercises in Racket. Its Scheme dialect is basically a superset of MIT Scheme.
  Having figured out my setup and the way the API worked, I went through the
  exercises rather quickly. DrRacket isn't great if you're used to Vim and
  Emacs, but it worked quite nicely.
- The last part of chapter 2 was quite a bit annoying. In fact, I considered
  putting away SICP on multiple occasions. There's a myth about SICP: that it
  only introduces state in the third chapter; after more than 200 pages, that
  is. Indeed, state isn't introduced until the third chapter, but the exercises
  in the last part of chapter 2 heavily rely on state: A lookup table for
  generic procedures can only be managed using state. Figuring out the workings
  of the `set!` procedure hasn't been an issue, but the whole arithmetics
  package that was developed was quite involved. Different examples and
  exercises provided different ways of doing computations on different types,
  and they didn't always play well together. I had to skip exercise 2.92, which
  was a sacrifice I had to do in order to finish the chapter. I am very glad
  that I choose this tradeoff, which allowed me to keep going.

So chapter 2 offered quite some challanges; the first due to the short lifespan
of technology, the second one, I daresay, due to some didactic imperfections. I
learned a lot about functional programming in the chapter, and writing tail-call
optimized recursive functions really became second nature to me.

I've spent six months on two chapters and worked through 216 pages. The book has
five chapters on 610 pages, i.e. three chapters on almost 400 pages are left. If
I'd continue at the current pace, I could finish SICP roughly at the end of the
year, if not later.

# What's Next?

My initial plan was to do at least the first two chapters, but probably also the
third one. The subjects of the third chapter—state, scope, data structures,
concurrency, and streams—are important concepts that well transfer to other
functional programming languages and environments. It's also one of the biggest
chapters with 140 pages.

The fourth chapter is about evaluation and the inner workings of LISP. The fifth
chapter is about register machines and simulations. Those two chapters might
lead to the often quoted LISP enlightment. However, I might consider a break
after chapter 3. With this much Scheme exposure under my belt, I probably could
pick up SICP again years later on chapter 4.

Chapter 3 might take three months, so I have plenty of time to consider my
future plans. My long-term plan is to stick to functional programming, and I'd
like to pursue my path on three different tracks:

## LISP

Having finished at least three of the five chapters of SICP, I'll probably leave
Scheme behind and try out different LISPs:

1. _Clojure_: I already used it for some toy examples (computing soccer league
   tables, simulating the Game of Life), and it yielded very concise code.  This
   would be a practical choice that could also be used for writing web
   applications.
2. _Racket_: Its Scheme dialect comes with nice additions such as hash tables
   and support for concurrent programming, and might be even useful for very
   practical things.

## ML

I had some exposure to Standard ML and Haskell before, and I consider its type
system to be extremely helpful, especially when I look at the messy arithmetic
package I was developing in chapter 2 of SICP. There are a lot languages in the
ML family, of which I consider the following the most interesting:

1. _Elm_: A domain-specific language for writing web frontends that compile to
   JavaScript. I worked through some basic examples, but the clean way of
   handling state really impressed me. I'd like to rewrite some of my old
   JavaScript toy programs in Elm.
2. _Haskell_: This I consider to be the holy grail of functional programming.  I
   really want to figure out how to write real-world applications with a pure
   functional programming language.
3. _OCaml_: This is supposed to be a more relaxed version of Haskell, which is
   quite close to Elm, and provides a compiler that produces rather fast
   binaries.

## Erlang

This was my initial motivation for functional programming, because it is also
about my second pet subject: concurrency. So I have to learn it!

1. _Erlang_: Even though Elixir is considered to be the primary choice nowadays,
   one needs to know the host language when dealing with the hosted language.
   The Prolog-like syntax doesn't bother me at all; and the boiler-plate code
   doesn't look too scary.
2. _Elixir_: For web applications, I'd probably use Elixir instead of Erlang.
   Elixir comes with some really nice additions, such as nice tooling and the
   pipe operator.

---

So I have plenty of ideas, but I should stick to one subject at a time, to which
I devote some daily practice. It worked quite well with SICP so far. And so it
will for my next endeavour, hopefully!
