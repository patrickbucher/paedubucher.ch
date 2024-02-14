---
title: Distracted by Tools
subtitle: The Sorry State of Computer Science Tuition
author: Patrick Bucher
date: 2018-10-18T12:00:00+0100
lang: en
---

I’m doing a bachelor’s degree at a university of applied sciences with its own
computer science department. I work part-time besides and already have a couple
of years of practice in the field. In fact, I’ll soon reach the point of being
longer in IT than I have spent time outside of it, for I’m already 31 years old
and started my apprenticeship with the age of 16.

I’m mostly interested in programming. This term, I picked two rather big
courses on that topic:

- _Programming Concepts and Paradigms_: This course has four parts. First, we
  got an overview of different fields of languages and looked at the
  differences between structural and object-oriented programming languages.
  Then we went over to declarative programming languages and spent three weeks
  on logical programming with Prolog, which was a really rewarding experience.
  The last couple of weeks we’ll get the chance to dig into a programming
  language of our own choice.  But today, we started with
  declarative-functional programming, which we’ll do in Racket/Scheme.
- _Microcontrollers_: This course has two parts: C Programming in general, and
  working on a specific microcontroller platform with assembly language and C
  in particular.

I already glanced at LISP a few times, and am somewhat familiar with the C
programming language (I worked through the K&R book and wrote a couple of tiny
helper programs for my personal use). Therefore, I don’t feel anxious about
failing and can look forward to do the things well with reasonable effort. (I’d
also like to spend some time on Go besides, if possible.) I am also willing to
do the things the way it suits my working style best (working alone, not in a
group; skip classes and read about the subjects on my own if, I feel like it).

However, the professors sometimes give me a hard time.

In the microcontrollers course, we’re supposed to use Windows. All the example
projects are for the Codewarrior IDE, which is based on Eclipse. I asked a
couple of Linux users, who did the course before, and all of them sticked to
Windows for that particular part. Because working with microcontrollers is a
total new field for me (except some Arduino toy projects), I’d rather stick to
the professor’s instructions and use the proposed environment. 

For the C part of the course, however, which takes place on a different day, I
use my usual Arch Linux setup consisting of `vim`, `gcc` and `Makefile`s. The
professor told us to use NetBeans for that purpose, which is a bit weird in my
opinion.  If we use the Eclipse-based IDE Codewarrior for C programming on
microcontrollers, we could certainly also use it for C programming on our local
machine. But I won’t use an IDE for the C exercises anyway.

So I did fine with the C setup of my choice so far. But today, I had to present
my solution to a mandatory exercise. It was about refactoring the
reverse-polish calculator from the K&R book. The one big `main.c` file needed
to be broken down into a header file and a couple of `.c` files. I also wrote a
`Makefile` and used some additional compiler flags, which I both quickly showed
and commented on. I demonstrated the program and showed the code in my terminal
using `vim`.

The professor considered the exercise being well done and gave me a good hint
that I could handle some variables differently. I realized that I should make
them `static`, which I did on the fly. I re-compiled the program and ran it
again, everything worked.

When I went back to my desk, the professor made an additional remark. He asked
me, why I use `vim`. I simply stated that it is a very capable text editor, and
that I’m extremely efficient working with it. He then told me a story about his
apprenticeship, where apprentices sometimes used the wrong tools for a job,
say, using pliers instead of a wrench for tightening a nut (it was mechanics,
not IT). So he suggested using the _best_ tool for the task at hand, which is a
modern IDE like NetBeans, and not `vim`, according to him. Besides that,
`Makefile`s were something the IDE generates, nothing to write on one’s own. So
I should really get into using a “real and professional IDE”, and not hacky old
`vim`.

I didn’t want to start an argument, because it was already late in the evening,
and three more students needed to demonstrate their exercises within the next
couple of minutes left. I just said that I use IDEs at work (mostly Eclipse and
IntelliJ for Java, and sometimes RStudio for R), but that I prefer _simple_
tools for simple tasks, like the exercise at  hand.

A fellow student, who also mostly uses Linux, used Visual Studio Code for both
C programming and for the presentation. (The professor didn’t comment on his
choice of tools.) After the course, this fellow student and I walked to the
train station, and we made some ironic remarks about the professor’s comments
to blow off some steam. (As a Linux user, they really give you a hard time at
that university.) The discussion got a bit more serious, and we both agreed
that we should spend more time on really understanding the basics (compiler
flags, the structure of a `Makefile`) and not waste our time with big and
bloaty tools, where we spend time figuring out the difference between two
slightly different icons showing a hammer, which are used to trigger a build
with slightly different options.

On my way home, I had a couple of other thoughts:

- Aren’t we supposed to learn the basics instead of some workflows of specific
  products, which are going to change every couple of years? In the first
  couple of lessons, the professor walked us through the build output, which
  was produced by an IDE-generated `Makefile`. He spent about ten minutes on
  that, because the output was rather big. The program being compiled was just
  a “Hello World” program, which I ran twice during that time by typing `make
  hello && ./hello` and `tcc -run hello.c`. (I even had enough time for
  installing `tcc` during his explanations.)
- An IDE is not a _tool_, it’s rather an entire _production line_, if the
  manufacturing metaphor should be preserved. And one specific production line
  (NetBeans) isn’t probably the best fit for the construction of every C
  programm. (If it were so, why are we supposed to use _two_ different IDEs for
  the same language in the same course?)
- Do Linux kernel hackers use an IDE other tham `vi`/`vim` or `emacs` with a
  couple of plugins and customisations?  I can’t imagine that, at least not
  something bloaty and Java-based such as NetBeans (or Eclipse). So the point
  that simple tools fail for big code bases, and that an IDE is better suited,
  seems to fail. (I don’t have any hard evidence for my claim, but I read that
  Linus Torvalds uses a customized MicroEMACS, for example. Having watched
  videos featuring other great hackers like Eric Raymond, Richard Stallman and
  Rob Pike, I cannot remember having seen an IDE once.)
- Aren’t computer science professors supposed to motivate students to
  experiment and try things out, at least during their studies? Aren’t we going
  to be told to use language X with IDE Y on platform Z for decades _after_ our
  studies? And aren’t young people the ones that come up with new and better
  ways of doing things once in a while?
- The professor sees the merits in the K&R book, but could he imagine its
  authors having used anything like NetBeans for writing the examples? Unix is
  about simplicity and frugality, not about unconditional use of the most
  powerful tools and concepts, no matter what the task at hand is.
- An IDE uses a lot of resources. One could argue that this is hardly an issue
  with our cheap and fast computers nowadays. Using a rather new Lenovo
  ThinkPad, memory, disk space and CPU load really aren’t the problem. However,
  the space on the screen is very limited, and my minimalistic desktop
  environment based on [dwm](http://dwm.suckless.org/) uses this particular
  resource very efficiently. So does `vim`, which doesn’t distract me with
  dozens of menu items, but gives me a simple and effective interface for
  dealing with text files. One could compare the IDE to an SUV and my setup to
  a bicycle, but that bicycle would run with the speed of a Formula 1 car
  consuming an ounce of olive oil for an entire race. (My setup also works
  great on a Rasperry Pi, even on the single-core version 1.)
- The professor also objected that `vim` doesn’t offer a debugger facility. I
  neither know nor care if there’s a `gdb` plugin for `vim`, but when I had to
  debug a program, I did just well with using `gdb` outside of `vim`. Hell,
  NetBeans also uses `gdb` as a C debugger, and so does Eclipse, but when I use
  raw `gdb`, I at least do not need to remember whether it was Eclipse or
  NetBeans that uses F6 for stepping over a line of code, I just learn the
  self-explanatory `step` command for `gdb` once and use it forever. (Using
  `gdb` productively requires debug symbols, which are included with the `gcc`
  flag `-g`, which would be a great opportunity for “joined-up thinking”—a
  notion to be found on every education expert’s slide desks.)
- An IDE is not only a tool that you can pick up and lay down again, it shapes
  the way your work is organized (folder structures, configuration files).
  Multiple programmers using different IDEs for the same project will certainly
  cause trouble sooner or later, and so will the transition of one IDE to
  another. (I personally don’t see a bright future for NetBeans, which I only
  saw being used _once_ outside of my school.) Working collaboratively with
  `vim` and `emacs`, or switching from one editor to another doesn’t give you
  that amount of pain.

Then I also thought about the other course in the morning, which was about
getting in touch with Racket/Scheme. We saw very little code (not even a
stereotypical Fibonacci or factorial function), talked only a bit about
concepts, but spent a good deal of the two lessons learning about options and
menu points in DrRacket, which is the IDE to be used during the course. I
didn’t figure out yet how to get the previous command again in the REPL
(`Ctrl-P` or up-arrow in most shells), but heard about the option to deactivate
case-sensitivity for identifiers. (I even found out how to switch the user
interface language to Russian, so that for once _I_ could give the professor a
hard time.) Wouldn’t I’ve been motived enough yet to learn Scheme, those two
lessons wouldn’t have done it either.

There are probably dozens of freely available LISP and Scheme implementations
with capabilities exceeding our needs for a three-week introductory course by
far. Pick one, get to learn how to use it, and spend the rest of the time
learning about the concepts (functional programming) and the language (be it
LISP, Scheme, Arc or whatever). But why wasting time installing and fine-tuning
tools we’re only going to use for a couple of weeks and never will see again?

And _if_ time is assigned to getting in touch with tools, why is it wasted on
things as ephemeral as menu items, configuration forms, and not spent more
wisely to get to know the command line flags (which most certainly must be
known in production) or the syntax of a `Makefile`?

So much time is wasted on irrelevant details. (I’m not even talking about CPU
time wasted during startup of those tools, nor about download traffic, memory
and hard disk usage.) For tuition, one should pick the most simple tools that
get a job done, learn how to work with them—and then focus on the subject
matter at hand.

And, _please_, if a student does an effort learning things on his own and
teaching them to his fellow students, be happy that the classroom gets to see a
new perspective on the subject matter—including the professor, who thinks he
already figured out the one and only way.
