---
title: Compose Key on X
subtitle: Say Goodbye to the Swiss Keyboard Layout
author: Patrick Bucher
date: 2020-09-16T21:00:00
lang: en
---

I've been using the Swiss keyboard layout for most of my life. Things changed
when I first had to work on a Mac Book. If the Swiss keyboard layout is not
great for programming on a usual keyboard, because it requires combinations with
Alt-Gr to type braces and brackets, I consider the Mac version of it outright
horrible, because braces and brackets are located on the digit's row of the
keyboard. Those symbols are not even painted onto their respective keys, which
makes the transition for non-Mac users even harder.

# Entering US Keyboard Layout

So instead of learning an additional inefficient way of typing special
characters, I decided to adopt to the US keyboard layout. This layout gives you
most characters required for programming with a single key or a combination of
Shift and another key.

One issue of the US layout is that it doesn't provide characters such as `ö` or
`é`, which I still needed to type in emails and for documentation. Mac OS
provides a composing mechanism that let's you type a double quote followed by a
letter like `o`, and then combines those characters to the German umlaut `ö`. If
you just want to type a double quotation mark, you need to type Space right
after it, so that the quotation mark is not attempted to be combined with the
next character entered. This is annoying of course when programming. So I never
really got warm with Mac keyboards. An external keyboard with US layout didn't
help much in that respect, but at least allowed me to type those cumbersome
combinations faster and more precisely.

# Switching Keyboard Layouts

Another solution to the goal conflict of typing German umlauts and special
characters programming as fast as possible is to switch the keyboard layout as
needed. I'm very familier with this approach, because I use the Russian keyboard
layout once in a while. I modified my [dwm](http://dwm.suckless.org/)
configuration so that pressing Super+Tab changes my keyboard layout using the
following script (`switchkb`):

    #!/bin/sh

    layout="$(setxkbmap -query | grep layout | egrep -o [a-z]{2}$)"
    if [ "$layout" == "ch" ]; then
        setxkbmap ru
    elif [ "$layout" = "ru" ]; then
        setxkbmap ch
    elif [ "$layout" = "ch" ]; then
        setxkbmap us
    fi

This lets me cycle through the keyboard layouts I need very quickly. I also
display the keyboard layout in my status bar using
[slstatus](http://tools.suckless.org/slstatus/), which I described in a
[previous article](/articles/2020-09-05-openbsd-on-the-desktop-part-i.html).
Even though this is a feasible solution, I still have to switch mentally between
the Swiss German and the US keyboard layout. Typing a double quote in a German
email requires me to press another button than typing that very same character
when programming. This additional mental burden is wearing me down on a usual
working day which consists of 70% programming and 30% communicating. The
communication part might actually be bigger, but some of the communication also
takes part in English, so that I'd be better off using the US rather than the
Swiss German keyboard layout.

Typing cyrillic letters still requires the Russian keyboard layout, so I won't
get rid of my `switchkb` script and Super+Tab anytime soon. However, since
typing punctuation marks using the Russian keyboard layout is almost the same as
typing those characters on the US keyboard layout, I rather ditch my native
Swiss German layout and embrace the imperialistic option.

# Composing Characters

While being able to write source code faster and not having to distinguish
between English, German, and Russian when typing punctuation marks sounds great
in terms of _efficiency_, not typing the typographically correct representations
of German umlauts (_Ae_, _Oe_, _Ue_ instead of _Ä_, _Ö_, _Ü_) or misspelling
french words with accent marks (_depecherent_ instead of _dépêchèrent_) is bad
in terms of communicating _effectively_. So there must be at least a way to type
those characters somewhat efficiently.

One option would be to type those characters as hexadecimal Unicode code points.
I've already learnt of few by heart, such as U+2012, U+2013, and U+2014 for
dashes of different lenghts, or simple to remember ones like U+00ab and U+00bb
for guillemets (_«»_). In `vim`, those sequences can be entered by presing
Ctrl+V followed by the part after U+. In GTK applications, Ctrl+U enters a mode
to enter those codes to be finished with a Space (or maybe some other character
not representing a hexadecimal digit). However, remembering a lot of Unicode
code points is not a very intuitive way to type and probably leads to a lot of
lookups on [FileFormat.Info](https://www.fileformat.info), which is at least a
more sophisticated way than just googling those character code points, which
usually ends up on FileFormat.Info anyway, or, worse, on a big Wikipedia page
discussing the history and cultural significance of the `LATIN SMALL LETTER C
WITH CEDILLA` (ç) before giving me the code point U+00E7.

## Entering the Compose Key

A better option is certainly the _Compose Key_. (Check out this [Wikipedia
Article](https://en.wikipedia.org/wiki/Compose_key) for a discussion of its
history and cultural significance). Even though not even the nerdiest of nerd
keyboards come with a physical compose key nowadays, the X Window System still
supports the underlying mechanism. And since X is running on both my Linux and
OpenBSD computers, using its compose key mechanism helps me both at work and for
my private computing, the latter consisting mostly of writing articles about
entering special characters on special computing setups and the like.

A good reference for this mechanism is the manual page `Compose(5)` or
`XCompose(5)` on both GNU/Linux and OpenBSD. (I guess it also applies to
FreeBSD, but I didn't check yet). However, there's some additional information
required to do a basic setup, so here's a quick guide.

First, a compose key needs to be picked. There are various options, which can be
looked up as follows:

    # OpenBSD
    $ grep 'compose:' /usr/X11R6/share/X11/xkb/rules/base.lst

    # Arch Linux
    $ grep 'compose:' /usr/share/X11/xkb/rules/base.lst

Which shows a list of keys that could be used for the compose key:

      compose:ralt         Right Alt
      compose:lwin         Left Win
      compose:lwin-altgr   3rd level of Left Win
      compose:rwin         Right Win
      compose:rwin-altgr   3rd level of Right Win
      compose:menu         Menu
      compose:menu-altgr   3rd level of Menu
      compose:lctrl        Left Ctrl
      compose:lctrl-altgr  3rd level of Left Ctrl
      compose:rctrl        Right Ctrl
      compose:rctrl-altgr  3rd level of Right Ctrl
      compose:caps         Caps Lock
      compose:caps-altgr   3rd level of Caps Lock
      compose:102          &lt;Less/Greater&gt;
      compose:102-altgr    3rd level of &lt;Less/Greater&gt;
      compose:paus         Pause
      compose:prsc         PrtSc
      compose:sclk         Scroll Lock

Caps Lock is a good choice, but would make it harder to type anonymous hate mail
in all-caps to people not agreeing with my operating system preferences.
Therefore I pick the Menu key, which I didn't even use once in my dark Windows
days. The compose key can be defined in `~/.xinitrc`, ideally together with the
choice of the aforementioned imperialistic keyboard layout:

    setxkbmap us
    setxkbmap -option compose:menu

On my Thinkpad, I prefer the PrtSc key, which is placed between AltGr and Ctrl
on the right hand side of the Space bar. I don't often take screenshots, since the
content of my screen is mostly text, which better captured using primary
selection. (Did I already mention that I use Arch, btw.?)

    setxkbmap us
    setxkbmap -option compose:prtsc

After restarting X, the compose key is already ready for action. Now I can type
`Menu " O` to enter _Ö_ on my PC (OpenBSD), or `PrtSc ' e` to enter _é_. Not all
together at the same time, but as a sequence, which is much more convenient.
Pre-defined sequences can be looked up in the compose files under
`/usr/X11/share/X11/locale/[locale]/Compose` for OpenBSD, or
`/us/share/X11/locale/[locale]/Compose` on Arch Linux, respectively. As locale,
I simply use `en_US.UTF-8`, which gives me a wide range of sequences
(hand-picked examples):

    <Multi_key> <C> <o> 			: "©"   copyright # COPYRIGHT SIGN
    <Multi_key> <R> <o> 			: "®"   registered # REGISTERED SIGN
    <Multi_key> <plus> <minus>     	: "±"   plusminus # PLUS-MINUS SIGN
    <Multi_key> <s> <s>            	: "ß"   ssharp # LATIN SMALL LETTER SHARP S
    <Multi_key> <E> <equal>        	: "€"   EuroSign # EURO SIGN
    <Multi_key> <u> <slash> 		: "µ"   mu # MICRO SIGN
    <Multi_key> <quotedbl> <A>     	: "Ä"   Adiaeresis # LATIN CAPITAL LETTER A WITH DIAERESIS
    <Multi_key> <acute> <E>        	: "É"   Eacute # LATIN CAPITAL LETTER E WITH ACUTE
    <Multi_key> <asciitilde> <N>   	: "Ñ"   Ntilde # LATIN CAPITAL LETTER N WITH TILDE
    <Multi_key> <comma> <c>        	: "ç"   ccedilla # LATIN SMALL LETTER C WITH CEDILLA
    <Multi_key> <v> <z> 			: "ž"   U017E # LATIN SMALL LETTER Z WITH CARON

Those sequences are very intuitive to type, so looking them up will hardly be
needed. This configuration also allows me to type scientific transliterations of
Russian words, such as Č in Čechov (`Menu < C`) instead of the German Tschechow
or the English Chekov. Same with š in Puškin and Ž in Dr. Živago.

If those sequences do not suffice for one's particular needs, more sequences can
be defined in the file `~/.XCompose`. When doing so, it is important to also
include the original definitions mentioned above as follows, so that it's not
needed to re-define them all:

    include "%S/en_US.UTF-8/Compose"

The `%S` resolves to the path `/usr/X11R6/share/X11/locale` for OpenBSD and
`/usr/share/X11/locale` for GNU/Linux. (`%L` would resolve to the current
locale.) Additional rules can be defined as follows:

    <Multi_key> <colon> <minus> <parenright> : "☺" U263A # WHITE SMILING FACE

This allows to type the smiley using `Menu : - )` on my OpenBSD machine. It is
also possible to enter whole strings in that manner:

    <Multi_key> <m> <a> <c> : "fanboy"

Whose effect I leave for you to figure out.
