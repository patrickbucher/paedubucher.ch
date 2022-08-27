---
title: Bumbling Boomer Bob and the Klean Kode Klan
subtitle: niladic, monadic, dyadic, triadic, polyadic, shmyadic
author: Patrick Bucher
date: 2022-08-27T10:00:00
lang: en
---

If you are around my age, work as a programmer, and took some classes on the
subject since 2009, you probably were subjected to Robert C. Martin's book
_Clean Code: A Handbook of Agile Software Craftsmanship_. (As of this day, the
[first
edition](https://www.informit.com/store/clean-code-a-handbook-of-agile-software-craftsmanship-9780132350884)
from 2008 hasn't been updated yet. Clearly, _Clean Code_ must be a—nay,
_the_—perfect book, which does not require improvements or adjustments).

I bought the book out of personal interest, probably in 2009. As many others, I
first was fascinated by it, but never managed to read more than one or two
chapters. I also rarely bothered to read specific sections offering solutions to
problems I faced on a daily basis. Probably the solutions offered weren't too
helpful, even though I worked as a Java programmer during that time, and _Clean
Code_ is very much about the kind of Java we wrote back then. So _Clean Code_
collected dust in my bookshelf.

# Clean Code and the Klean Kode Kult

When I studied computer science at the local technical college, I was exposed to
_Clean Code_ again. Or rather it was shoved down the students throats as a
gospel by disciples of the _Klean Kode Kult_, as I like to call them.

The members of the _Klean Kode Kult_ strengthen their belief by joining up for a
ritual called _Clean Code Shaming_, where they superficially look at a piece of
code they don't understand on first sight, and then just yell _«Clean Code!!!1»_
at its author in order to give proof of their superiority and sophistication.

Remember: Code you initially don't understand is _always_ just bad code and
certainly not a chance to improve your understanding of programming, especially
if pointless techniques like _Memoization_ or _Lexical Closures_ are used, i.e.
techniques you haven't been exposed to yet.

My friend [meillo](http://marmaro.de/) pointed out the [cult-like nature of
_Clean Code_](http://marmaro.de/apov/txt/2016-04-27_schaedlicher-kult.txt)
roughly at that time when its disciples came after me. A leader being called
«Uncle Bob», a scripture that doesn't require a second edition after many years
(but spawns sequels such as _The Clean Coder_, _Clean Architecture_, _Clean Agile_,
and _Clean Craftsmanship_), disciples willing to align themselves into
[grades](https://clean-code-developer.com/grades/) and wear
[bracelets](https://clean-code-developer.com/die-initiative/bracelets/) for self
castigation: If you still don't think that _Clean Code_ is a cult, just talk to
one of its disciples, point out a contradiction in the _Clean Code_ book, and
witness the angry reaction caused by your blasphemous remark.

But wait, a contradiction in _Clean Code_? That's impossible! Or maybe not?

# Bumbling Boomer Bob on Function Arguments

Working as a programmer for almost twenty years, I am still a layperson when it
comes to the exegesis of _Clean Code_, because I obviously still haven't been
englightened by this masterpiece yet. But trying to slay a straw man—and
failing to do so!—will hopefully bring me back to the Right Path, so that I can
finally abondon my wrongthink and give up on my hellish ends.

Let's hear what The Englightened has to say about function arguments (_Clean
Code_, Chapter 3, p. 40):

> The ideal number of arguments for a function is zero (niladic). Next comes one
> (monadic), followed closely by two (dyadic). Three arguments (triadic) should
> be avoided where possible. More than three (polyadic) requires very special
> justification—and then shouldn’t be used anyway.

This is misleading on so many levels that I need to dissect it in multiple
paragraphs.

Uncle Bob™ uses the terms _niladic_, _monadic_, _dyadic_, _triadic_, and
_polyadic_ for functions with arities of 0, 1, 2, 3, and n, respectively. There
certainly is a _qualitative_ difference betweeen 0, 1, and n, the difference
between, 2, 3, and n is only of a _quantitative_ nature. But judging by the
terms used, the author sees a _qualitative_ difference between all those
arities, to wit (_Clean Code_, p. 42):

> A function with two arguments is harder to understand than a monadic function.

And:

> Functions that take three arguments are significantly harder to understand
> than dyads. The issues of ordering, pausing, and ignoring are more than
> doubled. I suggest you think very carefully before creating a triad.

This difference is clearly of a _quantitative_ nature, because adding another
argument only makes the function «harder to understand», no matter if you go
from one to two of from two to three arguments.

Having been exposed to Haskell for a couple of hours, I'd expect to read about
_Curried Functions_ here: functions of arity `n` that return a function of arity
`n-1` when being invoked with a single argument. But obviously those Haskell
guys must be stupid, because they also bother with _Partial Function
Application_, which only makes sense when you have multiple arguments, i.e.
_diadic_, _triadic_, or even—Bob forbid—_polyadic_ functions!

This must also the reason why
[SICP](https://mitpress.mit.edu/9780262510875/structure-and-interpretation-of-computer-programs/)
makes for such a bad introductory textbook, because the Professors Abelson and
Sussman clearly haven't read _Clean Code_ when coming up with this abomination
(_Structure and Interpretation of Computer Programs_, Exercise 1.32, Chapter 1,
p. 61):

    (accumulate combiner null-value term a next b)

Six arguments, are you kidding me? Polyadic _ad nauseam_! Bob hates it.

SICP is such a horrible book that it even required a second edition. It even was
[adapted](https://mitpress.mit.edu/9780262543231/structure-and-interpretation-of-computer-programs/)
from MIT Scheme to JavaScript recently, whereas _Clean Code_ clearly would
withstand such blasphemous attempts, being firmly grounded in the Java culture
of the mid-2000s.

(If you didn't figure out where my exegesis went from serious to sarcastic, stop
reading this text and just forget about it. Put on the _Clean Code_ bracelet of
the day and refactor that cryptic `this.x += 3;` statement to a _clean_
`increaseXByThree();` niladic method instead.)

But Bumbler Bob is here to help (_Clean Code_, p. 43):

> When a function seems to need more than two or three arguments, it is likely
> that some of those arguments ought to be wrapped into a class of their own.
> (p. 43)

I wonder how many arguments a constructor for such a class might require.
Certainly, the introduction of the Builder Pattern would be The Right Solution™
for this issue. Much clearer than having a function with four arguments. (My
healing progress is starting!)

## Niladic (Im)purity

Boomer Bob is clearly familiar with the concept of _Pure Functions_, otherwise
he wouldn't object so strongly against side effects (_Clean Code_, p. 44):

> Side effects are lies. Your function promises to do one thing, but it also
> does other hidden things. Sometimes it will make unexpected changes to the
> variables of its own class. Sometimes it will make them to the parameters
> passed into the function or to system globals. In either case they are devious
> and damaging mistruths that often result in strange temporal couplings and
> order dependencies.

Unless Bobby-O considers lies as good, he clearly speaks out _in favour_ of
functions without arguments and _against_ side effects. So we should all be
writing side-effect free functions without arguments. But what can such a
function return?

- Nothing
- A constant value
- A random value

_Unless_ the function also operates on global variables or an object's
properties. But then those functions (or methods) are _not_ really side-effect
free, because their semantics is influenced by side-effects of other
functions/methods.

In order for a function to do something useful without side-effects, function
arguments are needed. The amount of arguments needed is determined by the
_domain_ of the function, i.e. by what the function is actually supposed to
do. Shall a function compute the solutoins to a quadratic equation, the
arguments `a`, `b`, and `c` are needed. Shall a function draw an arc on a
canvas? You need to define the coordinates of the circle's center (`x` and `y`),
its radius (`r`), start and end angle, and whether or not the arc shall be drawn
clockwise or counter-clockwise.

### Complexity: Inherent and Accidental

Admittedly, _Clean Code_ offers some useful advice to make such APIs easier to
understand, e.g.:

- Instead of having two paramters `x` and `y`, a `Point` abstraction might come
  in handy. (see _Argument Objects_ on p. 43)
- Instead of passing flag arguments (clockwise/counter-clockwise), provide two
  functions. (see _Flag Arguments_, p. 41)

But Boomer Bob entirely misses the point: The difference between _inherent_ and
_accidental_ complexity. Solving a square equation _requires_ three arguments.
An arc is _defined_ by its mid-point, radius, angles, and curve orientation.
This is complexity _inherent_ to the problem at hand.

The thickness, colour, and opacity of an arc being drawn on a canvas has to do
with a specific application of the concept. So while I consider it good advice
to wrap drawing details (thickness, colour, opacity) into an argument object, or
to use a `Point` abstraction instead of two loose `x` and `y` arguments, there's
no reasonable way to deal with arcs using _niladic_ or _monadic_ functions, save
for curried functions, which clearly aren't on Babbling Bob's mind here.

You might also separate the computation of an arc from actually drawing it.
Here, the computation returns the coordinates to be drawn, which you can pass
into the `draw` method of the canvas object, maybe together with the drawing
details. A pair of _monadic_ methods for setting coordinates and drawing details
on that object won't make anything clearer, but only introduce more side
effects: _accidental_ side effects this time, which change the state of the
object without any palpable benefit. Calling the `draw` method to actually draw
on the canvas is the only _desired_ side-effect: the complexity introduced
thereby being of an _inherent_ nature.

### More Arguments Can Make For Better Abstractions

Another example: Consider a `reduce` function. (This is a higher-order function,
but obviously not at the height of _Clean Code_, for such concepts are not
mentioned in The Masterpiece.) Consider the following interfaces:

    reduce(combine, values)

and

    reduce(combine, values, initialValue)

Where `combiner` is itself a function with the following interface:

    combine(accumulator, x)

The _dyadic_ `reduce` function must assume a value out of the given `values`
(usually its first element) as the initial value to be used for the
`accumulator`. So the _dyadic_ `reduce` can only return something of the same
type as the elements of `values` have. (E.g.  `values` is the integer array `[1,
2, 3]`, and `combine` sums up the `accumulator` with the current value argument
`x`, then `reduce` must return an integer.)

The _triadic_ `reduce` function can accept any initial value, as long as the
`combine` function is capable of dealing properly with that type of value. (E.g.
`reduce` can be used to partition the integer array of `[1, 2, 3, 4]` into two
arrays of odd and even numbers. The `initialValue` then could be a tuple of two
empty arrays: `([],[])`. Those arrays are filled by the `combine` function: odd
numbers in the first array, even numbers in the second array: `([1,3],[2,4])`.)

Not only is a _triadic_ `reduce` function more powerful than the _dyadic_ one,
it is also more _general_, i.e. a higher abstraction. If you want to do repeated
modifications to a vector in Clojure, you need something like a triadic reduce
function! But I guess that _Clean Coder_ Bob figured this out in recent years,
judging by his [delight in
Clojure](https://blog.cleancoder.com/uncle-bob/2019/08/22/WhyClojure.html).

# Conclusion

I ranted away half of my Saturday morning on roughly half a page of Bob's
Timeless Scripture. The problem is not that _Clean Code_ is a book with advice
that aged poorly, because it was written from a mid 2000s-Java perspective. The
problem is its uncritical fellowship taking this advice at face value, because
their perspective is too narrow.

I'm convinced that Robert C. Martin would write a totally different book on the
subject nowadays than he did back in 2008. But a second edition of _Clean Code_
would have very little in common with the first edition still being in print.
Rewriting the entire book probably would be less work than re-editing it. And
its fellowship would feel cheated by reading a book full of advice contrary to
its original.

If you feel offended by this text, please take an hour to watch Brian
Kernighan's lecture on [The Elements of Programming
Style](https://www.youtube.com/watch?v=8SUkrR7ZfTA) and let his advice sink in.
Reconsider your habit of yelling _«Clean Code!!!1»_, _«Train Wreck!!!1»_, or
even _«SOLID!!!1»_ (what doess the «L» stand for, again?) at other programmers
without first having tried to understand their code and familiarized yourself with
the concepts being used therein. Try out a functional programming language or
two, e.g. Haskell and Scheme, and consider their up- and downsides compared
to, say, Java or C#. Then read _Clean Code_ again (or: _actually_ read it), but
with the grain of salt extracted from your recent encounters with different
ideas and concepts. Read it critically, not as a gospel, and you'll extract some
real value out of it: by carefully considering each advice and its proper area
of application—as limited that might be.
