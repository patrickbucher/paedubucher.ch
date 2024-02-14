---
title: Cargo Cult
subtitle: Getting it Wrong by Wanting it Right
author: Patrick Bucher
date: 2018-10-05T12:00:00+0100
lang: en
---

Accepting interfaces instead of requiring specific implementation is a good
practice. Take Java collections as an example. A method accepting an `ArrayList`
argument will only be useful if the caller is already using an `ArrayList`.
Otherwise he’d need to convert the implementation he’s using, say, `LinkedList`,
to an `ArrayList` first. A method accepting the `List` or even the `Collection`
interface is far more likely to accept what the caller already has.

Many programmers recite this practice as “program on interfaces, not on
implementations”. However, this rule is misleading, because the somewhat blurry
notion of “programming on interfaces” does not mean the same as “accepting
interfaces”.

Once in a while, a unexperienced Java programmer tries to follow the “program
on interfaces” rule and stumbles upon a problem. He wrote:

    List<String> names = new ArrayList<>();

Where’s the problem? The unexperienced Java programmer objects that he’s not
“programming on an interface”, because he wrote `new ArrayList<>()` and hence
has a reference to a specific implementation.

Maybe he’ll try to write a `ListFactory` (he’s a Java Programmer, after all)
that provides a method with the signature `List<T> createList()` that will do
the “dirty work” for him.  But guess what: this method will contain the
instantiation of `ArrayList` as well. The problem is not solved, but wrapped in
an additional layer of pointlessnes.

What will the unexperienced Java programmer do now? He might search the issue
on the web. Maybe he finds an article about dependency injection.  But then
he’ll be wondering where to get that dependency injected from. So he’ll either
give up and stick to his original solution, or post the “issue” on a chat or
forum.

What’s wrong with the original code? Nothing. Interfaces are for references,
classes are for instances. At one point, somebody has to decide on an
implementation. One could wonder, if the compiler or a framework might figure
out the right implementation, optimized for performance according to what the
programmer uses the collection for. A linked list is faster than an array list
if items have to be inserted at the beginning. An array list is faster than a
linked list for random access. But the compiler is unable to figure that out,
at least at this point.

The issue is not only about performance, it’s also about semantics. Whereas the
programmer doesn’t have to care about whether a list is implemented using a
linked list or an array, semantics matter a lot to him in other cases. Take
maps for example. The Java interface `Map<K, V>` has the implementations
`TreeMap<K, V>` and `LinkedHashMap<K, V>`, among others.  When adding an
element to a `TreeMap`, the element will be inserted at a specific position, so
that the map’s element will be sorted in their natural order. A
`LinkedHashMap`, however, preserves the insertion order of the elements.

Of course, methods dealing with maps should still use the `Map` interface. It’s
entirely up to the caller, if and how the elements of the map passed in are
sorted. Maybe he doesn’t care, maybe he cares a lot. But he has to take the
decision. Such decisions not only affect the runtime properties of a program
(performance, memory consumption), but possibly also its semantics and
therefore its results.

So why the misunderstanding of the unexperienced Java programmer that insisted
on “programming on interfaces”? He’s guilty of _Cargo Cult Programming_. He
heard the general rule to “accept interfaces instead of implementations” and
understood it as “use interfaces instead of implementations” or even
interpreted it as “replace class names with interface names” (because he only
used the former in his code). He translated the rule “use A instead of B in
context C” and forgot about the context. He just remembered the “A instead of
B” part.

So how could the unexperienced programmer avoid such misunderstandings? Instead
of literally following some rules he heard of, or following the part of the
rule he understood, he should ask himself, if there’s an actual problem in his
code. If he’s unable to explain that problem in a couple of sentences, he
either doesn’t understand the problem―or maybe there is no problem at all.
