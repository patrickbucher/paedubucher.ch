---
title: Table-Driven Test Design
subtitle: With an Example in C
author: Patrick Bucher
date: 2020-07-22
lang: en
---

Many universities teach programming in Java. Writing unit tests is one of the
subjects being taught. Many professional Java programmers, but also university
professors, suggest to build those test cases according to a pattern. _Given,
When, Then_ is a common pattern, and so is _Arrange, Act, Assert_. Both patterns
prescribe the following structure for a test case:

1. _Given_/_Arrange_: An environment (in the broadest sense) is built up.
2. _When_/_Act_: The function or method to be tested is invoked.
3. _Then_/_Assert_: The result of the function or method is checked against some
   expectation.

Such a test case might look as follows (Java):

    public void testAddition() {
        // Given/Arrange
        Calculator calc = new Calculator();
        int a = 3;
        int b = 5;

        // When/Act
        sum = calc.add(a, b);

        // Then/Assert
        assertEqual(8, sum);
    }

A rule often taught is the so-called _single assert rule_ from Robert C. Martin,
[whom I refuse to call «Uncle
Bob»](http://marmaro.de/apov/txt/2016-04-27_schaedlicher-kult.txt). It states,
that there should be only one assertion per test case. One can argue whether or
not this rule is useful.

## Unclean Code

However, in my experience this rule leads to a consequence I do not like ‒ and
which also doesn't fit into the _Clean Code_ philosophy (or _cult_, I daresay):
The programming language being used to write test code is a small subset of the
implementation language, often degenerating into a sheer sequence of statements
(imperative programming).

Even though using a subset of a language is often a sensible approach (just
think about C++, or `with` and `eval` in JavaScript, or `unsafe` in Go, etc.),
using a subset of a language that doesn't even contain core features from
structured programming (decisions, loops, data structures) does not sound
sensible to me, except when programming in a purely functional style.

How should an additional test case to cover, say, negative numbers, be added to
the one above? The _single assert rule_ wants us to write an additional test
case:

    public void testAdditionWithNegativeNumbers() {
        // Given/Arrange
        Calculator calc = new Calculator();
        int a = -1;
        int b = 3;

        // When/Act
        sum = calc.add(a, b);

        // Then/Assert
        assertEqual(2, sum);
    }

Who would _type_ in that code, which is almost identical to the one above? Such
code is rather _copied_ than written again. (Why don't I hear somebody shouting
_«Clean Code!!!!11»_ now?)

## Structured Programming to the Rescue

Let's violate the _single assert rule_ for a minute and bring back structured
programming. Let's write a unit test in C!

    typedef struct {
        int a;
        int b;
        int expected;
    } addition_test_case;

    void test_addition()
    {
        addition_test_case tests[] = {
            {3, 5, 8},
            {-1, 3, 2},
        };
        int n = sizeof(tests) / sizeof(addition_test_case);
        for (int i = 0; i < n; i++) {
            addition_test_case test = tests[i];
            int actual = add(test.a, test.b);
            if (actual != test.expected) {
                printf("add(%d, %d): expected %d, got %d\n",
                        test.a, test.b, test.expected, actual);
                exit(1);
            }
        }
        printf("test_addition: %d tests passed\n", n);
    }

This test case, which does not make use of any unit testing framework, was
designed in a _table-driven_ manner. I first got to know the concept of
_table-driven test design_ when learning Go by reading [The Go Programming
Language](http://www.gopl.io/) (p. 306) by Alan A. A. Donovan and the great
Brian W. Kernighan.

However, the concept must predate Go, for I can at least remember one article by
Rob Pike, who later designed Go, mentioning table-driven test design.
(Ironically ‒ or not so ironically ‒ that article was a critique of
object-oriented programming, as far as I can remember.)

## Table-Driven Test Design

Let's break down the parts that make up a table-driven test design.

First, a single test case is defined using a structure that contains all the
input parameters, and the expected result of the test:

    typedef struct {
        int a;
        int b;
        int expected;
    } addition_test_case;

Second, an array ‒ the test _table_ ‒ containing all the test definitions is
defined (_Given_/_Arrange_):

    addition_test_case tests[] = {
        {3, 5, 8},
        {-1, 3, 2},
    };

Third, the test table is processed using a _loop_ (structured programming,
remember that?):

    int n = sizeof(tests) / sizeof(addition_test_case);
    for (int i = 0; i < n; i++) {
        // omitted
    }

For every test case, the result is computed (_Act_/_When_):

    addition_test_case test = tests[i];
    int actual = add(test.a, test.b);

Fourth, the result is validated against the definition (_Then_/_Assert_):

    if (actual != test.expected) {
        printf("add(%d, %d): expected %d, got %d\n",
                test.a, test.b, test.expected, actual);
        exit(1);
    }
    printf("test_addition: %d tests passed\n", n);

An error message is printed if the `actual` value is not equal to the `expected`
value (in case `add` was implemented incorrectly):

    add(3, 5): exptected 8, got 666

Note that this test terminates after the first error. No assertions are used.
The lack of a test framework is compensated by manually defined error and
success messages.

Yes, I'm well aware of the fact that there are unit testing libraries in C. The
point is that this C code covering two test cases is only slightly longer than
the Java code to cover the same amount of test cases would be. (Using Python or
Go rather than C would have shaved off some additional lines.)

Now let's add a third and a fourth test case:

    addition_test_case tests[] = {
        {3, 5, 8},
        {-1, 3, 2},
        {13, 17, 30}, // new
        (-100, 100, 0}, // new
    };

No code was copied. No existing code was modified. Only _two_ lines of code were
added to define _two_ additional test cases. The table-driven test is
_extensible_.  Robert C. Martin would love it, wouldn't he?

## Comparing Apples to Rotten Tomatoes

So why isn't everybody writing table-driven tests instead of triple-A copy-paste
tests?

First, some programming languages make it harder to define data structures as
literals. Languages like JavaScript, Python, or Go are quite good at that. Even
C, as shown above, can be quite concise when it comes to defining static data
structures. Java recently got better at that, but up to version 8, defining a
static map structure was done by adding single elements subsequently. (Why don't
I hear _«DRY principle!!!1»_ now?)

Second, the unit testing framework plays an important role. In C, (at least as
shown above), and in Go (as it is done using the standard library), no
assertions are used. The programmer instead performs the checks manually and
reacts with a reasonable error message. The programmer is supposed to _program_
the tests.

Some unit testing frameworks that do make use of assertions also allow to add
custom error messages to every `assert` call. Other frameworks, such as
[Jest](https://jestjs.io/), just will tell you _on which line_ an assertion
failed. This is not very useful when having assertions within a loop, for the
programmer does not know which test case failed. At least for Jest, writing pure
sequential assertion code is a necessity, and the _single assert rule_ looks
quite reasonable from that perspective.

The [PyTest](https://docs.pytest.org/en/latest/) framework, for example, has
table-driven test design built-in, by providing the static test definitions
through a decorator, which is basically an annotation in Java lingo. (Check
`@pytest.mark.parametrize` for details.) However, this approach makes it
impossible to include information into the test table that needs prior
construction within the test function.

More recent versions of JUnit also allow for parametrized tests (check out the
`@ParametrizedTest` and `@ValueSource` annotations). The restricted stated above
for PyTest also apply here. Again, the poor programmer is put into
straightjacket here, for he's not supposed to _program_ here, but only to
_test_.

My favourite test framework is from the Go standard library, which on one hand
gives the programmer total flexibility, and on the other hand provides an useful
API to construct small but powerful test runners. Checkout the
[testing](https://golang.org/pkg/testing/) package for details. (And read [The
Go Programming Language](https://gopl.io) by all means, even if you don't need
to learn Go. You'll pick up a lot about computer science in this book.)

## Single Assert Rule Revisited

The discussion about testing frameworks and programming languages (and text
editors, and tabs vs. spaces) could be extended here ad nauseam. But let's
review the _single assert rule_ instead, which could be interpreted from two
perspectives:

1. Runtime: `assert` should only be called once per execution of every test
   function/method.
2. Code: There should only be one reference to `assert` in every test
   function/method.

While the first interpretation makes table-driven design impossible, the second
interpretation might be closer to the rule's original intention: Each test case
should only verify one aspect of the function/method being called.

I'll therefore continue to happily violate the first interpretation of the rule,
for the advantages of table-driven test design (extensibility, flexibility, more
concise code) outhweigh the indiscriminate application of some hand-wavy
statements about «doing only one thing» by far. Please let me just _program_
those tests…

As an additional example, check out my test cases for some time formatting
routines
([test_timefmt.c](https://github.com/patrickbucher/countdown/blob/master/test_timefmt.c)).
Here, the test table can be used in two directions: One function uses the left
value as input and the right value as the expected outcome, while the other
functions does the opposite. Here, _two_ new test cases are defined by adding
_one_ (very short) line of code.

Am _I_ allowed to shout _«Clean Code!»_ and _«DRY principle!»_ now, by the way?
