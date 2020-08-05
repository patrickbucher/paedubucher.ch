---
title: '«Four in a Row» in Haskell (Part II)'
subtitle: Implementation of the Board Logic
author: Patrick Bucher
date: 2020-08-05T08:00:00
lang: en
---

In my [last
article](https://paedubucher.ch/articles/2020-08-03-four-in-a-row-in-haskell-part-i.html),
I outlined the purpose of a _stock program_: a non-trivial coding exercise to
be done in every new programming language somebody is learning. I also stated
that «Four in a Row» is my personal stock program, and that I'd like to
implement it in Haskell.

The main challenge in Haskell is the functional programming paradigm.
Immutability is the main difference between an implementation of «Four in a
Row» in a functional programming language compared to rather structured
programming languages such as C or Python. The object-oriented aspect of an
implementation in Python makes hardly a difference, for OOP equally allows for
mutable an immutable programming beneath the surface. (In introductory courses
on OOP, hidden mutability is rather praised as a virtue; the disadvantages of
mutability are only taught in advanced courses by showing the advantages of
constructs like immutable classes. Learn and unlearn, but I'm digressing…)

A later re-implementation of my stock program in Python might profit from the
experiences made in Haskell. Structured programming also allows for
immutability, and list comprehensions allow for compact code to produce new
state based on older state, without modifying existing state. (This might be
matter for a fourth article, but let's not get ahead of ourselves.)

In this article, I'm going to show how the board logic of «Four in a Row» can
be implemented in Haskell.

# Let There Be Code

As analyzed in my previous article, the board logic consists of five building blocks:

1. Create an empty grid with given dimensions.
2. Validate if a move (i.e. the choice of a column) is allowed for a given board.
3. Set a player's stone in the right place on the grid based on the choice of a column.
4. Detect if a player's won the game by checking if four of the player's stones
   lay in a horizontal, vertical, or diagonal line.
5. Format the grid as a string in order to display it on the command line.

The last building block, formatting, won't be covered in this article. I first
have to learn more about strings, formatting, and IO in Haskell, but I don't
like to wait to cover the other parts, which I'm already capable of
implementing with my current knowledge.

## Type Glossary

Before implementing the actual logic, let's define a couple of type aliases:

    type Grid = [Row]
    type Row = [Int]
    type Col = [Int]
    type Stone = Int

A grid (type `Grid`) is a list of rows. A row (type `Row`) itself is a list of
integers. As discussed in my previous article, `0` is going to be used for
empty fields. The fields occupied by player one and two shall be represented by
the numbers `1` and `2`, respectively.

Just like a row, a column (type `Col`) is a list of integers. It is an
alternative way to express the relationships between individual fields. The
`Grid`, however, uses the `Row` type as its building blocks.

A `Stone` is an integer, too. It represents the player's number for fields
occupied by a particular player.

Those types won't add powerful abstractions to the program, but make the
signature of certain functions a bit clearer. (It's also possible to limit the
scope of the types declared to certain values, but let's focus on the program
logic instead.)

## Creating a Grid

The function `new_grid` accepts two integer parameters (number of rows and
columns) and produces a grid of those dimensions:

    new_grid :: Int -> Int -> Grid
    new_grid r c = [new_row c | _ <- [1..r]]

A list comprehension is used to build up the grid as a list of `r` rows. The
row itself is created by a function `new_row`:

    new_row :: Int -> Row
    new_row c = [0 | _ <- [1..c]]

Again, a list comprehension is used to build a single row consisting of `c`
elements: one per column.

The `new_grid` function can be used as follows (`>` indicates the REPL, the
output has been wrapped for better readability):

    > new_grid 6 7
    [[0,0,0,0,0,0,0],
     [0,0,0,0,0,0,0],
     [0,0,0,0,0,0,0],
     [0,0,0,0,0,0,0],
     [0,0,0,0,0,0,0],
     [0,0,0,0,0,0,0]]

## Validating a Move

A move solely consists of a column index. Let's assume a 6x7 grid (6 rows, 7
columns) if nothing else is stated. A valid move must be in the range of 0
(leftmost column) to 6 (inclusive, rightmost column).

For a move to be valid, the column must have an empty field, i.e. it must
contain the value `0`. Since the columns are filled up from the bottom, a
column is not full if its the top-most field is equal to `0`. So this
validation seems trivial.

However, in order to deal with _columns_ rather than _rows_ (remember, the grid
is defined in terms of rows, now the other way around), we first need a way to
gather the fields of a column. The function `get_column` expects a grid and a
column index and returns the fields belonging to that particular column:

    get_column :: Grid -> Int -> Col
    get_column g c = [row !! c | row <- g]

A list comprehension is used to select the element at index `c` in every grid
row using the index operator (`!!`).

The function `is_valid_move` simply extracts the column chosen by the player
and checks its topmost field to be empty (equals `0`, that is):

    is_valid_move :: Grid -> Int -> Bool
    is_valid_move g c = (get_column g c) !! 0 == 0

(Notice that no boundary checks are implemented throughout the program, unless
absolutely necessary for getting the logic right.)

This function can be used as follows:

    > g = new_grid 6 7
    > is_valid_move g 0
    True

## Setting a Stone

The first two building blocks were easy to write without modifying state.
Performing a move on the grid by setting a stone into a certain column,
however, is a step that requires a modification of some sort. The solution is
to not mutate the given grid, but to produce a new grid based on the given grid
by accounting for a player's move.

The function `apply_move` expects a grid, a column (chosen by the player and
validated using `is_valid_move`), and the player's number (to set the right
value in the new grid):

    apply_move :: Grid -> Int -> Int -> Grid

Because only a column is given, the row coordinate has to be figured out. Since
stones played are falling down the grid in the physical version of the game,
the bottom-most free field of a column has to be found:

    bottom_most :: Grid -> Int -> Int -> Int
    bottom_most g v c = length (takeWhile (\x -> x == v) col) - 1
                        where col = get_column g c

The lowest free position is found by extracting a subsequent list of a given
value `v`, which can be handed in as an argument. (The value `0` has to be used
for this particular use case by the caller.) The built-in function `takeWhile`
is used to extract a list. The bottom-most position of a column with the given
value `v` is simply the length of the extracted sub list minus one (indexes are
zero-based). Again, the `get_column` function is used to get access to a
particular column. 

Now `apply_move` can be implemented as follows:

    apply_move g c p = replace_value g r c p
                       where r = bottom_most g 0 c

Another function is needed: `replace_value`, which creates a new grid based on
the existing grid `g`, by setting the player's stone value `p` to the
coordinate `r`/`c`. (The row coordinate is figured out using `bottom_most`, as
shown above.)

The function `replace_move` is implemented as follows:

    replace_value :: Grid -> Int -> Int -> Stone -> Grid
    replace_value g r c p = take r g ++ [new_row] ++ drop (r + 1) g
                            where new_row = replace_row_value (g !! r) c p

Given the row index `r`, the first `r` rows are taken. (This excludes the row
to be transformed, because the index is zero-based.) The row at index `r` is
computed as `new_row` in a further step. The remaining rows are extracted from
the existing grid by dropping the first `r + 1` rows from it. Those three
components are concatenated using the `++` operator to a new grid.

The `new_row` looks like the old row at index `r`, expect that a single value
at index `c` (the column) has to be replaced with the player's value `v`. The
function `replace_row_value` performs this transformation:

    replace_row_value :: Row -> Int -> Stone -> Row
    replace_row_value r c p = take c r ++ [p] ++ drop (c+1) r

The same logic using `take` and `drop` can be implemented for the column's
field like for the grid's rows before. The empty field at column index `c` can
simply be replaced by a list solely consisting of the player's value `v`. List
concatenation is used again to produce the tranformed column.

A move can be applied as follows:

    > g = new_grid 6 7
    > g1 = apply_move g 3 1
    > g1
    [[0,0,0,0,0,0,0],
     [0,0,0,0,0,0,0],
     [0,0,0,0,0,0,0],
     [0,0,0,0,0,0,0],
     [0,0,0,0,0,0,0],
     [0,0,0,1,0,0,0]]

    > g2 = apply_move g1 4 2
    > g2
    [[0,0,0,0,0,0,0],
     [0,0,0,0,0,0,0],
     [0,0,0,0,0,0,0],
     [0,0,0,0,0,0,0],
     [0,0,0,0,0,0,0],
     [0,0,0,1,2,0,0]]

`apply_move` could also be invoking `is_valid_move` for validation. But this
task should be left for the client to be implemented later on.

## Detecting a Win

TODO: Code must be improved first, for the resulting row of a move is _not_
returned, and the client cannot know which row is to be evaluated. Extend
`is_win` to check all the fields of the player that just performed a move.

# Conclusion

TODO: logic works, no boundary checks, some things are not really nice
