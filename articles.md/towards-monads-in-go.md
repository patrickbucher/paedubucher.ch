---
title: "Towards Monads (in Go)"
subtitle: "If I'd got them, I couldn't explain them…"
author: "Patrick Bucher"
date: 2022-11-05T15:23:00
lang: "en"
---

> As soon as you understand Monads, you're no longer able to explain Monads.

— Somebody on the Internet

Having read `n` tutorials on Monads, one must come to the conclusion that
tutorial `n+1` needs to be written. Not so much in order to explain monads to
others, but rather to gain more clarity. So here I write down my current
understanding of Monads, and my few readers get the chance to correct my
misconceptions about them. I'll be using example code in Go for this purpose,
because Haskell programmers won't bother to visit my site anyway.

# Basic Arithmetic in Go

Functors are about composing functions, Monads are about composing _lifting_
functions. For a start, we need some functions to compose. Everybody knows
function composition from basic arithmetic:

```
3 - 2 + 1 * 4
```

We have _numbers_ that are chained together by _operators_.

First, let's define a type alias for numbers:

```go
type Number float64
```

Second, we need a type for operators:

```go
type Operator func(Number) Number
```

But wait, why does the Operator only take a single number? Because we use
closures to implement them:

```go

func Adder(i Number) Operator {
	return func(n Number) Number {
		return n + i
	}
}

func Subtractor(i Number) Operator {
	return func(n Number) Number {
		return n - i
	}
}

func Multiplier(i Number) Operator {
	return func(n Number) Number {
		return n * i
	}
}

func Divider(i Number) Operator {
	return func(n Number) Number {
		return n / i // NOTE: divide by zero possible!
	}
}
```

An `Operator` function is defined as some operation to be done on another
number. For example, a `Multiplier` for the value `3` can be created. This
function then can be applied to another number, say, `2`, which then is
multiplied with `3` (`2 * 3 = 6`).

Notice how elegantly the `NOTE` comment in the function `Divider` points to a
possible problem to be resolved later on in the text. Yes: we could divide by
zero, which is the problem to be solved using Monads.

Here's how to use those operators:

```go
func main() {
	var n Number = 1
	addTwo := Adder(2)
	subOne := Subtractor(1)
	mulThree := Multiplier(3)
	divFour := Divider(4)

	// ((((1 + 2) * 3) - 1) / 4) = 2
	fmt.Println(divFour(subOne(mulThree(addTwo(n)))))
}
```

We start with `n=1`, add `2` to it (`=3`), multiply it with `3` (`=9`), subtract
`1` from it (`=8`), and divide it by `4`, which produces the result `2`.

# Composing Functions

Nested function can be tedious to read. If we'd like to re-use a formula, it's
better to _compose_ those functions, which is achieved by another function
called `Compose`:

```go
// f(g(n))
func Compose(f, g Operator) Operator {
	return func(n Number) Number {
		x := g(n)
		y := f(x)
		return y
	}
}
```

Two functions `f` and `g` can be composed as `f(g(x))`. The composition of those
functions is itself a function, let's call it `h`. Calling `h(x)` is the
equivalent of calling `f(g(x))`. (I call `g` the _inner_ and `f` the _outer_
function, and `f` is called with the result from the inner function call.)

Let's compose the functions frome before:

```go
func main() {
	var n Number = 1
	addTwo := Adder(2)
	subOne := Subtractor(1)
	mulThree := Multiplier(3)
	divFour := Divider(4)

	f := Compose(mulThree, addTwo)
	g := Compose(subOne, f)
	h := Compose(divFour, g)

	// ((((1 + 2) * 3) - 1) / 4) = 2
	fmt.Println(h(n))
}
```

Which produces the same result as before: `2`.

## Bad Composition

Let's mess up this example by introducing some nefarious operation: divide
something by zero (I'll be leaving out the boiler-plate `main` code from here):

```go
divZero := Divider(0)
f = Compose(mulThree, addTwo)
g = Compose(subOne, f)
h = Compose(divZero, g) // NOTE: this would break everything
fmt.Println(h(n))
```

This, of course, does not produce a numeric result, but `+Inf`. One could make a
mathematical argument on the correctness of this result (which I can't), or
refer to the [IEEE 754](https://standards.ieee.org/ieee/754/6210/) standard for
floating-point arithmetic (which I did), but further composition is clearly hurt
by such a result.

If the definition of `Number` is changed to `int32`, for example, the programm
will break:

```
panic: runtime error: integer divide by zero
```

Which definitively hurts composition.

So we need a safe way to divide numbers, which either produces a proper result
to be used for further computations (or as the final result), or some kind of an
error, so that further computations downstream won't be attempting to compute
with.

# Wrapping Results

A new type `Result` is created, which can represent the result of a successful
computation (`Val`) or the reason why this computation failed (`Err`):

```go
type Result struct {
	Val Number
	Err error
}
```

The operators are supposed to deal with such results. More precisely: The
operators still accept values of type `Number`, but produce values of type
`Result`:

```go
type LiftingOperator func(Number) Result
```

Those operators are supposed to _lift_ an ordinary `Number` into another
context: the `Result`.

A `Result` is used like a union in C: either `Val` is set (indicating a properly
performed computation), or `Err` is set (indicating some problem). The
implementations of the operations addition, subtraction, and multiplication are
straightforward:

```go
func LiftingAdder(i Number) LiftingOperator {
	return func(n Number) Result {
		return Result{n + i, nil}
	}
}

func LiftingSubtractor(i Number) LiftingOperator {
	return func(n Number) Result {
		return Result{n - i, nil}
	}
}

func LiftingMultiplier(i Number) LiftingOperator {
	return func(n Number) Result {
		return Result{n * i, nil}
	}
}
```

The result of the computation is wrapped in `Result` with no error and returned.

The division, however, might produce either a `Result` with an actual value
(`Val`), or a `Result` only consisting of an error (`Err`):

```go
func LiftingDivider(i Number) LiftingOperator {
	return func(n Number) Result {
		if i == 0 {
			return Result{0.0, errors.New("divide by zero")}
		}
		return Result{n / i, nil}
	}
}
```

## Composing Lifting Functions

Those lifting functions can be composed, too, but require a modified
implementation of the `Compose` function from before, called `ComposeLifting`:

```go
func ComposeLifting(f, g LiftingOperator) LiftingOperator {
	return func(n Number) Result {
		x := g(n)
		if x.Err != nil {
			return Result{0.0, fmt.Errorf("call inner function: %w", x.Err)}
		}
		y := f(x.Val)
		if y.Err != nil {
			return Result{0.0, fmt.Errorf("call outer function: %w", y.Err)}
		}
		return Result{y.Val, nil}
	}
}
```

If the result of the inner function `g` turns out be be erroneous, the outer
function `f` is never called, but the error returned from `g` is wrapped and
returned. The same is done for the result of the outer functon `f`. Only if both
`g` and `f` produce proper results (represented by a value) can the composed
function itself return a result with a value set.

Let's compose some _lifting_ functions:

```go
addLiftingTwo := LiftingAdder(2)
subLiftingOne := LiftingSubtractor(1)
mulLiftingThree := LiftingMultiplier(3)
divLiftingFour := LiftingDivider(4)

fl := ComposeLifting(mulLiftingThree, addLiftingTwo)
gl := ComposeLifting(subLiftingOne, fl)
hl := ComposeLifting(divLiftingFour, gl)
fmt.Println(hl(n))
```

This program produces the output `{2 <nil>}`, i.e. the result `2` was computed,
and no error (`<nil>`) occured.

Let's compose some computations including the nefarious division by zero:

```go
addLiftingTwo := LiftingAdder(2)
subLiftingOne := LiftingSubtractor(1)
mulLiftingThree := LiftingMultiplier(3)
divLiftingZero := LiftingDivider(0)

fl := ComposeLifting(mulLiftingThree, addLiftingTwo)
gl := ComposeLifting(subLiftingOne, fl)
hl := ComposeLifting(divLiftingZero, gl)
fmt.Println(hl(n))
```

Which produces the following output:

```
{0 call outer function: divide by zero}
```

Since there is an error, `0` is _not_ the result, but a placeholder. The error
message tells us pretty well what went wrong.

The function composition is also not hurt if we put the nefarious division by
zero in the middle:

```go
addLiftingTwo := LiftingAdder(2)
divLiftingZero := LiftingDivider(0)
subLiftingOne := LiftingSubtractor(1)
mulLiftingThree := LiftingMultiplier(3)

fl := ComposeLifting(divLiftingZero, addLiftingTwo)
gl := ComposeLifting(subLiftingOne, fl)
hl := ComposeLifting(mulLiftingThree, gl)
fmt.Println(hl(n))
```

However, a longer error message is produced:

```
{0 call inner function: call inner function: call outer function: divide by zero}
```

The code can be found on
[GitHub](https://gist.github.com/patrickbucher/70b8f1ebe91ea2f50cb829c7bad657a7).

# Conclusion

Functions operate on values and can be composed to new functions. Lifting
functions operate on values, too, but lift them in a context for some purpose
like error handling. This lifted context is called a Monad. Lifting functions
can be composed and make the aspect, for which the lifting happens, transparent
to the user of an API. Monads are about composing lifting functions.

Note: I didn't bother to mention [Monad
Laws](https://wiki.haskell.org/Monad_laws).

