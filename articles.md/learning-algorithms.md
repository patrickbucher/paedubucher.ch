---
title: Learning Algorithms
subtitle: Slowly, but Steadily (again)
author: Patrick Bucher
date: 2023-07-27T16:22:00
lang: en
---

A year ago, I was stuck in my efforts to properly learn Clojure. So I decided to
take a step back and work through [SICP](https://github.com/patrickbucher/sicp):
slowly, but steadily. I sticked to it rather well, even worked on an exercise on
the day I moved, and on Christmas, too. However, I had to stop in March, because
working on the first project for my own company (Composed GmbH) besides working
two other jobs was too much.

# Satisfied with SICP

I quickly took up SICP again in July to learn about lazy evaluation and streams,
which is also used in Clojure. But now that I got out of it what I needed, I'm
back at Clojure. Those nine months of SICP with Scheme and Racket really paid
off. And since I'm not particularly interested in building my own LISP, I won't
continue with SICP for the moment… but I'll leave the bookmark where it is
(almost at the end of the third chapter).

So what's up next? A lot of teaching from August to February (15 lessons a week)
including a completely new course (on software testing), which takes up three
days a week. Then there's a lot to do for my own company; maybe project work, or
building up some infrastructure and creating a proper website (probably using
Hugo). But I'd also like to learn something new besides, and doing so
systematically (again).

# Programming Languages…?

I looked at a lot of different programming languages in the last couple of
years. I wrote some code in AWK, Bash, C, C#, Clojure, Elixir, Elm, Erlang, Go,
Haskell, Io, Java, JavaScript, Lua, Octave, PHP, Python, Prolog, R, Racket,
Ruby, Rust, Scheme, Standard ML, and SQL. I was exposed to 25 languages in the
last, say, five years. This is very broad, but therefore also rather shallow. So
instead of picking up fresh ones (like OCaml or Perl), I'd rather get deeper
into the ones I already know. For me, the most interesting ones are:

1. **Erlang** and **Elixir**: Actually two (functional) languages both based on
   the same technology; the Erlang VM (Beam) and OTP. I'm very interested in
   Concorrency, and Go's CSP model was a revelation in that respect for me. Now
   I'd like to learn more about the Actor model, and on how to build resilient
   applications.  For practical purposes (e.g. building web applications), I'll
   use Elixir. But since that's a hosted language, I should also spend some time
   with the host language Erlang.
2. **Clojure**: In my opinion, this is the most beautiful language. The basic
   data structures (lists, vectors, maps, and sets) just feel right. There's a
   powerful and intuitive API for them, they can be made up with nested
   literals, and once you get the hang on persistent data structures, you don't
   want to go back. I already know the host language (Java), and this
   interoperability makes it a great choice for practical projects like web
   applications. I haven't used macros yet, so there's even more power lurking
   beneath. Rich Hickey created Clojure, because he was frustrated with all the
   other languages. The result is a huge pleasure to work with, even though I
   probably need to improve my workflow with Emacs.
3. **Rust**: Just in case Erlang/Elixir and Clojure should be too slow or too
   memory intensive for certain tasks, Rust will solve both issues. It's quite
   functional in some respects (iterators, enums, pattern matching) and strongly
   typed. However, its collections are basically mutable, and many operations on
   them only work if they are declared so. This is one step back from the
   persistent data structures of the other languages/platforms, but exactly what
   is needed for efficiency.

With those three (or four) platforms (or languages), I'll cover a lot of ground.
I _did_ neglect languages suited for writing web frontends, but this is not my
main concern. (And Elm would probably be my choice for it.)

But what should I do with these programming languages? Without a proper project,
I won't stick to them and do something else.

# …and Algorithms!

For the lack of a productive project (with deadlines and payment, but also with
the option to pivot to a language I already know rather well), I have to make up
some kind of project.

There's a very big book on the very bottom of my bookshelf: [Introduction to
Algorithms (4th Edition)](https://mitpress.mit.edu/9780262046305/), which I now
own since its release in spring 2022. I wasn't able to read a single page of it
yet, but now the time has come. Being more confident with academic computer
science texts after my encounter with SICP, I'd like to tackle this one.

The algorithms are given as pseudo-code, so much I already figured out. I don't
know yet if there are exercises, and if those require making mathematical
proofs. (I'm writing this during my holidays, away from my bookshelf.) However,
I'd like to focus on the implementation, for which I have to understand the
algorithms.

Besides building up an arsenal of algorithms for different kinds of problems,
I'd like to practice the programming languages from above on those algorithms.
But learning two new things—algorithms _and_ languages—probably will be _too_
hard. So here's my plan:

1. **Reading** and **Understanding**: I'll probably need to read every section
   two or three times until I understand it. And I should also make some
   sketches and play through some examples using pen and paper. With that
   understanding, I'll go ahead with the implementation of the algorithm.
2. **Go**: The pseudo-code descriptions are a good fit for the strucutred
   programming model of Go. I already know Go quite well, and it comes with all
   the tools needed to play with algorithms: Data structures (slices, maps,
   structs), relatively lean syntax, and a built-in test module with
   benchmarking. Using Go, I can focus on the algorithm and play with it, until
   I really understood it.
3. **Rust**: Having a working version in Go, I can just translate it to Rust.
   There will be a few language-specific issues to deal with (immutability, the
   need for smart pointers for recursive data structures), but that is exactly
   what I need to learn. Doing some runtime comparisons with Go will also be
   interesting.
4. **Erlang**: Now towards functional programming: Having implemented an
   algorithm in two languages, I'll now translate it to a different paradigm.
   This will teach me the basics of Erlang. Making the algorithms work
   concurrently would also be a good exercise, for which I probably should first
   go back to Go (and Rust), which also yields interesting comparisons for the
   different concurrency models. (Elixir can wait for later.)
5. **Clojure**: Having grasped the functional implementation of an algorithm,
   re-writing it in Clojure should amount to simplification and more concise
   code.

TODO: finish
