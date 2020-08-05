---
title: '«Four in a Row» in Haskell (Part I)'
subtitle: Background and General Considerations
author: Patrick Bucher
date: 2020-08-03T12:00:00
lang: en
---

In a [recent interview](https://youtu.be/O9upVbGSBFo?t=3741), Brian W. Kernighan
said that he always re-implements the same program when he's learning a new
programming language. In his case, it's a programm to process a text file
containing tabular data. In this task, his programming language AWK (Kernighan
is the «K» in «AWK») shines, for it was designed for that kind of a task.

Such a _stock program_ allows to evaluate a programming language from a certain
perspective. Different programs offer different perspectives. I personally
didn't have such a stock program yet, but there is at least one program I have
already implemented in multiple programming languages: the board game _Four in a
Row_.

# Four in a Row: My Stock Program

This game is played by two players, usually on a 7x6 grid (seven columns, six
rows). The grid is setup to be perpendicular to the table, so that the stones
fall to the lowest free field of the chosen column. The players take turns
setting their stones (red for one player, yellow for the other one). The player
that first can set four of stones into a horizontal, vertical, or diagonal row
wins the game.

I first implemented this game as a program towards the end of my first year as
an apprentice. The task was an optional assignment in an introductory
programming class. C was used as the implementation language. A more recent
re-implementation of that program is available on
[GitHub](https://github.com/patrickbucher/prog/blob/master/vier_gewinnt/vier_gewinnt.c).
The hardest part was to get the winning detection right, especially for the
diagonal rows. Since the grid was implemented as a two-dimensional array,
diagonals clipping the edge would erroneously also be detected as a winning row.
Some additional checks for index boundaries fixed the issue.

16 years later, my apprenticeship already was far in the past. I was studying
computer science in my eight and last term. For a _Game Design_ class, I had to
write a case study on improving an existing game. I picked _Four in a Row_ and
extended it with a couple of new game mechanics. The case study, written in
German, and the source code, can be found on
[GitHub](https://github.com/patrickbucher/v13r93w1nn7), too. This time, I used
Python as the implementation language. The [NumPy](https://numpy.org/) library
made this task very comfortable, and I was able to implement the board logic
with rather few lines of Python code. The unit tests, implemented using
[PyTest](https://docs.pytest.org/en/stable/), took up far more lines than the
actual code.

Both versions were implemented for the command line. However, the latter
version was implemented in a way that would also support graphical frontends.

## Building Blocks

Having implemented the same program with much more programming experience and
using a different programming language, the resulting code looked quite
different. However, I was able to detect some common patterns.

On a very high level, there are two parts for such a program: First, the _board
logic_ that deals with the grid, its manipulations and validations (Is a row not
full yet?  What is the bottom-most empty row in a given column? Are four stones
of the same color in a row?). Second, the _game logic_, which consists of a big
loop that lets the players take turns setting their stones, prints the grid, and
ends the game upon a win or draw.

The board logic can be taken further apart into the following components:

1. **Creating an Empty Grid**: At the beginning of a game, an empty grid with
   given dimensions has to be created. (The physical game is played on a 7x6
   grid, but a computer game can offer additional flexibility with the number of
   rows and columns given as arguments.)
2. **Validating a Move**: As soon as all fields of a column are filled, the
   column must no longer be chosen by players. A function is needed that checks
   which columns still have at least one empty field.
3. **Setting a Stone**: If a stone is to be set into a non-full column, the
   bottom-most empty row of that column has to be figured out. Then, the field
   is modified by setting the player's stone into that position.
4. **Detecting a Win**: After every move, it has to be checked whether or not
   the grid contains four stones of the same color laying in the same
   horizontal, vertical, or diagional row, without any gaps in between. If the
   detection gets to know which player did the last move, and into what
   coordinates that stone was put, the algorithm has to do less work, as opposed
   to an approach where the whole grid is evaluated for both players. (For the
   case-study, I had to use the latter approach, for one of the additional game
   mechanics allowed to flip the grid, which required a full evaluation of the
   whole grid afterwards.)
5. **Formatting the Grid**: This part could also be implemented in the game
   logic.  However, offering the capability to print the current grid from the
   board component (be it a module or a class) in a nicely formatted way is a
   good design decision in terms of cohesion. This function can be made very
   flexible by accepting formatting parameters, such as the characters to be
   used to display fields that are empty, or contain a stone of either player.

A function to format the current grid makes an important separation between the
inner state of the grid and its textual representation on the command line. It
is a good idea to represent the state of the fields as _numbers_ internally,
but to use _characters_ in order to display them nicely on the command line.
Internally, `0` can used for empty fields. For fields holding a stone of player
one or two, the values `1` and `2`, respectively, can be used. The empty field
can be displayed using a whitespace character, an underscore, or a dash. The
stones of the players can be easily distinguished when using `x` and `o` for
their output.

# Towards Haskell

The programming language _Haskell_, which has been mentioned in this article's
title, but not in the text ever since, shall be used to create an additional
implementation of the _Four in a Row_ game. But why Haskell?

First, I'm currently learning Haskell. It turns out that writing useful programs
in Haskell is not that easy, because advanced concepts like Monads have to be
understood in order to perform input/output operations. I'm working through the
rather dense book [Programming in Haskell (Second
Edition)](https://www.cs.nott.ac.uk/~pszgmh/pih.html) (by Graham Hutton) at the
moment, and I've almost finished the first part. The knowledge acquired from
those first nine chapters allows me to implement the board logic. The
interactive part then has to wait until I (nearly) finished the book.

Second, I'm interested in functional programming. I consider Haskell as a
stepping stone into that programming paradigm. I have some minor experience in
Prolog, and I'd like to learn Erlang later on. Knowledge about functional
programming also helps when programming in Python and JavaScript, which also
support features like lambda expressions, higher-order functions, and, in case
of Python, list comprehensions.

Implementing _Four in a Row_ in Haskell gives me a couple of challenges.
Unlike an implementation in C or Python, the grid must not be modified during
gameplay. A new grid, representing the fresh state, has to be build up based on
the previous state and the player's action, instead. I also need to figure out
how to detect a winning row in a declarative way, i.e. without loops and
counter variables. The input/output of the actual game logic will probably be
the biggest challenge later on. The game logic, implemented as a loop in both C
and Python, needs to be implemented using a different mechanism.

My plan is to implement the board logic, consisting of the five components
stated above, in the next couple of days in Haskell. I'll write an article
describing my approach and containing the code for the board logic as soon as I
have a decent solution for the problems stated. The game logic has to wait for a
couple of weeks of even months, depending on my progress with _Programming in
Haskell_.

Stay tuned, and feel free to put (maybe needed?) pressure on me, when those
articles do not appear any time soon…
