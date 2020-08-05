---
title: '«Four in a Row» in Haskell (Part II)'
subtitle: Implementation of the Board Logic
author: Patrick Bucher
date: 2020-08-05T23:00:00
lang: en
---

In my [last
article](https://paedubucher.ch/articles/2020-08-03-four-in-a-row-in-haskell-part-i.html),
I outlined the purpose of a _stock program_: a non-trivial coding exercise to
be done in every new programming language somebody is learning. I also stated
that «Four in a Row» is becoming my personal stock program, and that I'd like to
implement it in Haskell.

The main challenge in Haskell is the functional programming paradigm.
Immutability is the main difference between an implementation of «Four in a
Row» in a functional programming language compared to rather structured
programming languages such as C or Python. The object-oriented aspect of an
implementation in Python makes hardly a difference, for OOP equally allows for
mutable an immutable programming beneath the surface. (In introductory courses
on OOP, hidden mutability is rather praised as a virtue than frowned upon; the
disadvantages of mutability are only taught in advanced courses by showing the
advantages of constructs like immutable classes. Learn and unlearn, but I'm
digressing…)

A later re-implementation of my stock program in Python might profit from the
experiences made in Haskell. Structured programming also allows for
immutability, and list comprehensions allow for compact code to produce new
state based on older state, without modifying existing state. (This
re-implementation could be matter for a fourth article, but let's not get ahead
of ourselves.)

In this article, I'm going to show how the board logic for the game «Four in a
Row» can be implemented in Haskell.

# Let There Be Code

As analyzed in my previous article, the board logic consists of five building blocks:

1. Create an empty grid with given dimensions.
2. Validate if a move (i.e. the choice of a column) is allowed for a given board.
3. Set a player's stone in the right place on the grid based on the choice of a column.
4. Detect if a player has won the game by checking if four of the player's
   stones lay in a horizontal, vertical, or diagonal line.
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

A `Stone` is an integer, too. It represents a player's number for fields
occupied by his or her stones.

Those types won't add powerful abstractions to the program, but make the
signature of certain functions a bit clearer. (It's also possible to limit the
scope of the types declared to certain values, but let's focus on the program
logic instead.)

## Creating a Grid

The function `new_grid` accepts two integer parameters (number of rows and
columns), and produces a grid of those dimensions:

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
is defined in terms of rows, not the other way around), we first need a way to
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
is used to extract a list based on a lambda expression: Elements are taken from
the column as long as the lambda expression holds true. The bottom-most position
of a column with the given value `v` is simply the length of the extracted sub
list minus one (indexes are zero-based). Again, the `get_column` function is
used to get access to the fields of a particular column. 

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
components are concatenated to a new grid using the `++` operator.

The `new_row` looks like the old row at index `r`, expect that a single value
at index `c` (the column) has to be replaced with the player's value `v`. The
function `replace_row_value` performs this transformation:

    replace_row_value :: Row -> Int -> Stone -> Row
    replace_row_value r c p = take c r ++ [p] ++ drop (c+1) r

The same logic using `take` and `drop` can be implemented for the column's
fields like for the grid's rows before. The empty field at column index `c` can
simply be replaced by a list solely consisting of the player's stone value `v`.
List concatenation is used again to produce the tranformed column.

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

Figuring out whether or not a player's most recent move leads to a win is the
hardest part of this program, no matter what implementation language is used.
(However, I didn't try Prolog _yet_ for this.) Let's analyze the problem.

First, what do we know? The player with a number (`1` or `2`) just picked a
column (between `0` and `5` in our 6x7 grid). A stone was set in the bottom-most
empty field of that column. The actual row where the stone landed in is unknown.
However, this information can be found out: it is the top-most row of the chosen
column holding the player's stone value. All the fields above must be empty.

Second, what do we need to find out? Starting from the coordinates (given
column, row figured out as described above), there are three possibilities to
build a row of four values: horizontal, vertical, and diagonal lines. A
horizontal line is a row, and a vertical row is a column. Diagonal lines can
occur in two directions: ascending or descending. So we actually need to account
for four kinds of rows, which need to be extracted from the row/column
coordinates.

Third, once the horizontal, the vertical, and the two diagonal lines going
through the player's stone most recently set are established, a simple check can
be done: Does the line, which can be represented as a list, contain a list of
four of the player's stones? If that's the case, the player just won the game.

Let's implement that algorithm in a top-down manner!

The function `is_win` expects a grid, a column, and a player's stone value, and
returns a boolean value indicating if the player just won the game:

    is_win :: Grid -> Int -> Stone -> Bool
    is_win g c p = horizontal_win g row p ||
                   vertical_win g c p ||
                   diagonal_win g row c p
                   where row = top_most g p c

Three predicate functions `horizontal_win`, `vertical_win`, and `diagonal_win`
handle the three different shapes of winning rows. To check for a vertical win,
the row is irrelevant. For the other wins, the row where the player's stone just
landed in is figured out using the `top_most` function:

    top_most :: Grid -> Stone -> Int -> Int
    top_most g v c = length (takeWhile (\x -> x /= v) col)
                     where col = get_column g c

This function expects a grid, a player's stone value, a column, and returns the
top-most row containing the player's stone. Going through the column from top to
bottom, values are read into a list as long as they are not equal to the
player's stone value. The length of that list is the row coordinate of the
player's top-most stone in that column. Again, the column is extracted using the
`get_column` function explained further above.

### Vertical and Horizontal Win

A vertical and horizontal winning row can be detected in the same manner. The
only difference is that the former works on columns, and the latter on rows:

    horizontal_win :: Grid -> Int -> Stone -> Bool
    horizontal_win g r p = contained fiar (g !! r)
                           where fiar = [p | _ <- [1..4]]

    vertical_win :: Grid -> Int -> Stone -> Bool
    vertical_win g c p = contained fiar (get_column g c)
                         where fiar = [p | _ <- [1..4]]

In both cases, a grid, an index (row or column, respectively), and a player's
stone value is expected. The boolean return value indicates whether or not the
row or column contains a sub-list consisting of four of the player's stone
values: `fiar`, which is built using a list comprehension.

For the horizontal win, the row can be directly accessed from the grid using the
row index (`g !! r`). For the vertical win, the `get_column` function is used
once again.

The function `contained` is the tricky part. This function checks whether or not
a smaller list (first argument) is part of a larger list (second argument). A
possible implementation looks as follows:

    contained :: Eq a => [a] -> [a] -> Bool
    contained [] []                     = True
    contained [] ys                     = True
    contained xs []                     = False
    contained (x:xs) (y:ys) | x == y    = and [x == y | (x,y) <- zip xs ys]
                                          && length xs <= length ys
                                          || contained (x:xs) ys
                            | otherwise = contained (x:xs) ys

The lists processed can be of any type that supports the comparison operator
(`Eq a`). A boolean value is returned indicating whether or not the first list
is contained in the second list. The function is implemented using pattern
matching, which covers the following cases:

1. An empty list is contained in another empty list (first base case).
2. An empty list is contained in any non-empty list (second base case).
3. A non-empty list is not contained in an empty list (negative base case).
4. A non-emtpy list is _possibly_ contained in another non-empty list (complex
   case).

The «possibly» in the fourth case can be resolved as follows: If the first
elements of the two lists do match, the remainders of the two lists need to be
checked for a match. A list comprehension zipping those tails together and
comparing the corresponding elements creates a list of booleans indicating
matches. If all those booleans are `True`, the first list must be contained in
the second list, _if the second list is at least as long as the first list_.
(Notice that the `zip` function only picks values until the shorter of the two
zipped lists is exhausted. The length check ensures that the comparison of the
lists does not end prematurely.)

The `otherwise` case is processed when the two list's heads to not match. In
this case, the `contained` function is invoked again with the full first list
and the second's list tail: It shall be checked whether or not the first list is
contained in the second's list tail.

### Diagonal Win

Detecting a diagonal win works in the same manner as detecting a horizontal or
vertical win. However, there are two subtle details that make the implementation
more complicated:

First, there are _two_ kinds of diagonal lines: ascending and descending. This
can be handled by implementing two different functions.

Second, extracting a diagonal line as a list from the two-dimensional grid is
much more complicated than extracting a horizontal line (row) or a vertical line
(column).

Let's start with the `diagonal_win` function, which accounts for both winning
rows in ascending or descending order:

    diagonal_win :: Grid -> Int -> Int -> Stone -> Bool
    diagonal_win g r c p = contained fiar (diag_asc g r c) ||
                           contained fiar (diag_desc g r c)
                           where fiar = [p | _ <- [1..4]]

The function expects a grid, both row and column indication, and the player's
stone value. As always, a boolean value indicating a win is returned. A win is
detected, if the list of four player's stone values is contained in either the
ascending or the descending diagonal line.

Those lines are extracted from the grid using the `diag_asc` and `diag_desc`
functions, respectively. The two functions look quite similar, but have subtle
differences in the way they process the grid:

- An _ascending_ row starts at the bottom of the grid, i.e. with the highest row
  index. It starts at the left, i.e. with the lowest column index.
- A _descending_ row starts at the top of the grid, i.e. with the lowest row
  index. It also starts at the left, and, thus, with the lowest column index.

The function `diag_asc` expects a grid and both row and column indices. It
returns the ascending diagonal row containing that coordinate:

    diag_asc :: Grid -> Int -> Int -> [Int]
    diag_asc g r c = [g !! i !! j | (i,j) <- zip (reverse [0..max_row]) [min_col..cols-1]]
                     where
                       max_row = r + offset
                       min_col = c - offset
                       offset  = max (min (rows - r - 1) (cols - c - 1)) 0
                       rows    = length g
                       cols    = length (g !! 0)

The function is implemented using a list comprehension. The variable `i` is the
row index, `j` the column index. Those indices are obtained by zipping a list of
row indices with a list of column indices. The starting and end point of those
lists are the tricky part.

Consider this grid, in which `-` stands for an empty field, and the upper-case
`F` for the field played most recently (with the `r` and `c` arguments as
indices). All the fields indicated with a lower-case `f` are to be extracted for
the ascending diagonal holding the upper-case `F`:

        !
    0 1 2 3 4 5 6
    - - - - - - f 0
    - - - - - f - 1
    - - - - f - - 2
    - - - f - - - 3
    - - F - - - - 4 !
    - f - - - - - 5

The row and column indices of `F` are given as `4` and `2`. The starting point
at the bottom-left can be figured out by shifting the coordinates by an
_offset_. This offset is the smaller value of the following two differences:

- `rows - r - 1`: the number of rows minus the row index (minus one to account
  for the zero-based row index)
- `cols - c - 1`: the number of columns minus the column index (minus one;
  zero-based index again)

The offset is set to `0`, if either difference becomes negative (boarder
clipping). The offset is calculated as follows:

    offset = max (min (rows - r - 1) (cols - c - 1)) 0
    offset = max (min (6 - 4 - 1) (7 - 2 - 1)) 0
    offset = max (min 1 4) 0
    offset = max 1 0
    offset = 1

And the starting points `max_row`/`min_col` (bottom left) are calculated based
on the given indices of `F` as follows:

    max_row = r + offset
    max_row = 4 + 1
    max_row = 5

    min_col = c - offset
    min_col = 2 - 1
    min_col = 1

The diagonal line can be drawn up to the row index `0` and the column index `6`.
Here, it is possible to always use the maximum value, because the `zip` function
will stop picking values once the shorter list is exhausted.

The number of rows and columns can simply be figured out using the `length`
function applied on the grid as a whole and on a single row thereof:

    rows = length g
    cols = length (g !! 0)

Notice that in order to create a list containing the _falling_ values from
`max_row` down to `0`, a rising list from `0` to `max_row` has to be created and
reversed:

    > reverse [0..max_row]
    [0,1,2,3,4,5]

The other way around, an empty list would be created:

    > [max_row..0]
    []

The somewhat easier to understand function `diag_desc` is simply pasted here
without any further comments.  Figuring out how it works is left to the reader.
The extensive comments above on `diag_asc` certainly help for this purpose:

    diag_desc :: Grid -> Int -> Int -> [Int]
    diag_desc g r c = [g !! i !! j | (i,j) <- zip [min_row..rows-1] [min_col..cols-1]]
                      where
                        min_row = r - offset
                        min_col = c - offset
                        offset  = min r c
                        rows    = length g
                        cols    = length (g !! 0)

# Conclusion

The complete board logic required to implement a basic «Four in a Row» game has
been implemented in Haskell. The whole code described, plus some additional
attempts to format the grid as a string, can be found on
[GitHub](https://github.com/patrickbucher/programming-in-haskell/blob/master/four-in-a-row/Board.hs).

The linked code also defines a module `Board` which exports the public interface
of the board consisting of the four building blocks discussed in this article
and its predecessor. The file
[BoardTest.hs](https://github.com/patrickbucher/programming-in-haskell/blob/master/four-in-a-row/BoardTest.hs)
defines a couple of unit tests written in HUnit for basic verification of the
logic.

The actual board logic requires a little less than 100 SLOC. Comparable
implementations I've written in Python and C only take up slightly more lines. I
could have made some functions _shorter_, but probably not _clearer_ with my
limited knowledge of Haskell.

The `contained` function, for example, looks a bit bulky, but actually contains
very little logic. It is possible that the negative base case could be
eradicated, because a length check is already performed in the complex case.
However, I rather have a clear statement of the base cases than saving an easy
to understand line of code.

I might revisit this code and improve it as my knowledge of Haskell improves.
But the next step in my journey is to implement an interactive game based on
this board, which will be the subject of an article to be published in weeks or
maybe months.
