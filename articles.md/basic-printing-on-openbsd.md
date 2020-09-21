---
title: Basic Printing on OpenBSD
subtitle: Brother Network Printer with PostScript 
author: Patrick Bucher
date: 2020-09-20T18:00:00
lang: en
---

I have a roughly ten year old Brother HL-5370DW printer on the shelf next to me.
This printer is mostly used by my wife to print sewing patterns. When I was
studying computer science, I sometimes printed documents I've written for
proofreading. I often was able to find typos that I didn't see on the screen
even after proofreading the document two or three times. However, I didn't
bother to print out my bachelor thesis. Printing 120 pages just for proofreading
just seemed a waste to me. I did my proofreading on the screen extra carefully,
and nobody complained about typos. (Which doesn't mean that there were none.)

Having finished my studies, I hardly ever print out documents. However, I still
prefer to read long texts on paper rather than on the screen. Therefore I often
buy technical books as paperbacks or hardcovers rather than ebooks. And if I buy
an ebook with demanding content, I print out those sections for offline reading.

Having switched to OpenBSD for my private computing shifted my reading habits
more towards manpages. When I need to figure out how something works on OpenBSD,
`apropos(1)` beats Google as a starting point in many cases. Some manpages are
really long, for example `ksh(1)`. I have a book on the Korn Shell in my
basement, which covers `ksh93`. However, there are some differences between
`ksh93` and OpenBSD's `pdksh`. So reading the manpage not only gives me more
accurate information, but also _less_ to read.

So why not printing out the manpage `ksh(1)`? I can do so even nicely formatted
using PostScript:

    $ man -T ps -O paper=a4 ksh >ksh.1.ps

Now `ksh.1.ps` can be read with `zathura(1)`, given that the package
`zathura-ps` is installed:

    # pkg_add zathura zathura-ps
    $ zathura ksh.1.ps

But why using PostScript and not PDF like anybody else for the last twenty five
years? Because PostScript is the least common denominator and, thus, supported
out of the box by OpenBSD. (For fancier printing options, check out `cups`, but
I'd like to keep it minimalistic for the moment.)

# Printer Setup

I figured out how to configure my printer by reading the section _The lpd
Printing Daemon_ in the 16th chapter of [Absolute OpenBSD (2nd
Edition)](https://nostarch.com/obenbsd2e) (p. 306-307) by [Michael W
Lucas](https://mwl.io/). This is how I applied the configuration to my local
setup.

First, I created the file `/etc/printcap` with the following content:

    lp|brother:\
        :sh=:\
        :rm=192.168.178.52:\
        :sd=/var/spool/output/brother:\
        :lf=/var/log/lpd-errs:\
        :rp=brother

There must be a newline at the end of the file. The line breaks are escaped
using backslashes, except for the last line. The options are defined as follows:

- The first line defines two names for my printer: `lp`, which should always be
  there, and `brother`, which is my arbitrary name for the printer.
- The second line (`sh`) defines that no _burst page_ (summarizing the last
  print job on a special page) should be printed.
- The third line (`rm`) refers to the printer on the network. My FritzBox always
  gives the same IP to my printer. It's also possible to use the printer's
  hostname.
- The fourth line (`sd`) defines the spooler directory for this printer. Print
  jobs are written into that directory.
- The fifth line (`lf`) defines a log file for error messages, which you hopefully
  never need to check.
- The sixth line (`rp`) defines the remote printer name.

Next, the spooler directory needs to be created. It must be owned by the user
`root` and the group `daemon`. Regular users need write access to this directory
in order to print documents:

    # mkdir /var/spool/output/brother
    # chown -R root:daemon /var/spool/output/brother
    # chmod 777 /var/spool/output/brother

Now the printer daemon `lpd` needs to be activated. To do so on system startup,
add the following line to `/etc/rc.conf/local`:

    lpd_flags=""

Then start the service:

    # /etc/rc.d/lpd restart

**Update (2020-09-21)**: As one reader on
[Hackernews](https://news.ycombinator.com/item?id=24535357#24538879) pointed
out, the last two steps can be performed using `rcctl(8)`:

    # rcctl enable lpd
    # rcctl restart lpd

The manpage says that `rcctl(8)` was introduced in OpenBSD 5.7 back in 2015.
_Absolute OpenBSD (2nd Edition)_ is from 2013 and, thus, older than that.

# Printing Documents

Now the printer is ready to accept jobs. In order to print the PostScript file
generated before, just run `lpr` on the file:

    $ lpr ksh.1.ps

It's also possible to send the PostScript output directly to the printer (this
is Unix, after all), if no preview is needed:

    $ man -T ps -O paper=a4 ksh | lpr

Printing plain text files behaved strange on my setup, but could to using the
`pr` formatter with `lpr` as follows:

    $ lpr -p plain.txt

Instead, I also convert plain text files to PostScript, which looks quite nice
on paper. I use `enscript(1)` for this task:

    # pkg_add enscript
    $ enscript plain.txt -o plain.ps
    $ lpr plain.ps

PDFs can also be converted to PostScript using `pdf2ps(1)`, which comes with
GhostScript, i.e. the `ghostscript` package:

    $ pdf2ps document.pdf document.ps

Unfortunately, this doesn't work with all PDFs. But for the time being, I have
enough manpages to read. Printing PostScript works extremely fast, by the way.
When I press return at the end of a `lpr` command, I can see the status LED on
my printer start blinking almost immediately.
