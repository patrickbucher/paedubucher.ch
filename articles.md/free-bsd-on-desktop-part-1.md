---
title: FreeBSD on the Desktop (Part I)
subtitle: Basic Setup
author: Patrick Bucher
date: 2020-08-11T11:00:00
lang: en
---

I'm a happy user of [Arch Linux](https://www.archlinux.org/) both on my private
computers and on my work laptop. I even managed to get through four years of
university with my setup, and only had to bring a Windows machine on some rare
occasions, even though some professors are openly hostile towards a Linux setup.
(It doesn't run Microsoft Project and the real Excel, after all…)

Recently, I got interested in the BSDs, especially
[OpenBSD](https://www.openbsd.org/) and [FreeBSD](https://www.freebsd.org/).
Even though OpenBSD with its minimalistic appeal is suited to my taste, I'm
currently looking at FreeBSD for a couple of reasons. First, I'll have to
maintain a storage server (a FreeBSD setup using ZFS) at work. Second, I've also
built up a storage server at home. FreeBSD gives me the things I need mostly out
of the box (ZFS with redundancy on really cheap hardware). And third, I just
like to learn new things.

# Why FreeBSD?

However, when it comes to learning new things in my spare time, I'd rather spend
my time on something that will be useful in the long run. I use the [Lindy
Effect](https://en.wikipedia.org/wiki/Lindy_effect) as a guide: Technologies
like Kubernetes, the latest JavaScript framework, or Web Assembly have only been
around for a couple of years, and it's possible that those will go as fast as
they came. Older technologies _that are still around nowadays_, on the other
hand, can be expected to be around for many more years. Examples are the C
programming language, various Unix shells, and ‒ FreeBSD. (I also make
exceptions to this rule now and then. For example, I learned Go and Rust in the
summers of 2018 and 2019, respectively. Go, which had its 1.0 release earlier
than Rust, proofed to be the more stable choice.)

FreeBSD is now more than 25 years old. Its roots, however, go back to AT&T's
original Unix from the 1970s. (All the code has been replaced or rewritten
since.) Being old is not enough, of course; a technology is only worthwhile
learning if it is still alive. Even though FreeBSD is rarely on the front page
of [Hacker News](https://news.ycombinator.com/), it is still widely used.
[Netflix](https://papers.freebsd.org/2019/fosdem/looney-netflix_and_freebsd/)
streams videos through FreeBSD systems, and the operating system of the
PlayStation 4, [Orbis
OS](https://en.wikipedia.org/wiki/PlayStation_4_system_software), is based on
FreeBSD.

FreeBSD is not only likely to stay around for a long time, it probably also
won't undergo fundmental changes very soon or very often. One example is the
startup system of FreeBSD, which is still based on `init` and `rc` scripts.
Ubuntu, on the other side, switched from SysVinit to Upstart in 2006, and again
from Upstart to systemd in 2015. That is one init system to learn for a lifetime
(FreeBSD) vs. three in less than a decade (Ubuntu).

Documentation is another advantage of FreeBSD. Having a system that rarely
introduces breaking changes makes it easier and more worthwile to provide good
documentation. The FreeBSD team not only provides good [manual
pages](https://www.freebsd.org/cgi/man.cgi), but also a well curated
[FAQs](https://www.freebsd.org/doc/en_US.ISO8859-1/books/faq/) and the very
useful [Handbook](https://www.freebsd.org/doc/en_US.ISO8859-1/books/handbook/).
FreeBSD material has such a long shelf life so that it is even worthwhile to
print books about that operating system. I'm reading [Absolute
FreeBSD](https://nostarch.com/absfreebsd3) now (from front to back, that is).
Michael Warren Lucas, the author of this book, has written [even more
books](https://www.tiltedwindmillpress.com/product-category/tech/) on FreeBSD
(and OpenBSD), which are not only useful and of high-quality, but also
well-written and fun to read.

As I started to work with FreeBSD, I suddenly realized what the term
«distribution» is probably supposed to mean: Not just a bunch of software
cobbled together with more or less frequent upgrades, but an entire operating
system that not only provides working software, but also the means to build that
very software on a standard installation. A kernel can be compiled and installed
with a single command. Thanks to the ports tree, the packages can be compiled
easily and in a consistent way. I haven't tried _all_ Linux distributions, of
course, but quality standards that high are certainly not the rule in the
GNU/Linux world. (Debian and Arch, the Linux distributions I use the most and
know the best, are still absolutely great operating systems.)

# FreeBSD on the Desktop?

It's safe to say that FreeBSD is a good choice for servers in a Unix
environment. But does it also work well on a desktop computer? And is it even an
option for the kind of desktop I like to run: not GNOE, KDE, or Xfce, but a
minimalistic setup based on [dwm](http://dwm.suckless.org/), which probably
isn's used by many. Since dwm can only be configured by modifying the `config.h`
file, I won't be able to use the version from the ports tree. I use my desktop
computer mainly for work (programming, reading, writing, researching information
on the internet) and some entertainment mostly provided through the web browser
(reading, videos).

I generally use mid-range commodity hardware with on-board graphics, so hardware
compatibility should not be an issue. My desktop computer is a small-form Dell
OptiPlex 7040 from 2016 with an Intel Core i5 CPU. I replaced the original 128
TB SSD with a 500 GB model last year, and upgraded the original 8 GB of memory
with a ridiculous amount of 24 GB. (Some RAM bars just happened to lay around
here…) The computer has a WiFi card, of course, but since my router is just next
to my desk, I rather use a stable ethernet connection.

I already tried out OpenBSD once on that computer, and didn't have any issues
getting it to run. So FreeBSD having access to the same code base under the BSD
license is likely to suppport this hardware as well. Let's just try it out!

# Preparations

FreeBSD supports multiple versions at any given point in time. At the time of
this writing, 12.1 and 11.4 are the current versions intended for production.
Let's pick the most recent version 12.1. The [download
page](https://download.freebsd.org/ftp/releases/amd64/amd64/ISO-IMAGES/12.1/)
for the `amd64` architecture offers various options. The compressed
`mini-memstick` archive weighs the least and provides everything that is needed
for an installation on a computer with an internet connection. I download it to
my laptop runnng Arch Linux, and then verify the checksum:

    $ wget https://download.freebsd.org/ftp/releases/amd64/amd64\
    /ISO-IMAGES/12.1/FreeBSD-12.1-RELEASE-amd64-mini-memstick.img.xz
    $ wget https://download.freebsd.org/ftp/releases/amd64/amd64/\
    ISO-IMAGES/12.1/CHECKSUM.SHA512-FreeBSD-12.1-RELEASE-amd64
    $ sha512sum -c CHECKSUM.SHA512-FreeBSD-12.1-RELEASE-amd64 --ignore-missing
    FreeBSD-12.1-RELEASE-amd64-mini-memstick.img.xz: OK

The archive (389 MBs) needs to be unpacked and then is copied to a USB dongle
(`dev/sda`):

    $ unxz FreeBSD-12.1-RELEASE-amd64-mini-memstick.img.xz
    # dd if=FreeBSD-12.1-RELEASE-amd64-mini-memstick.img of=/dev/sda bs=1M
    $ sync

# Initial Setup

The BIOS is setup to use UEFI rather than legacy boot. I plug in the USB dongle
and start the FreeBSD installer. These are the settings I use during setup:

- Keymap: Swiss-German (`ch.kdb`)
- Hostname: `optiplex` (I'm not very imaginative when it comes to naming
  things.)
- Components: just `lib32`, `ports`, and `src`
- Network: interface `em0` with DHCP and IPv4
- Mirror: Main Site
- Partitioning: Auto (UFS, I'm going to learn about ZFS later) on the entire
  disk `ada0` using GPT and the following partitions (device, space, type,
  label, mount point):
    - `ada0p1`: 200 MB `efi` boot (no mount point)
    - `ada0p2`: 24 GB `freebsd-swap` swap0 (no mount point, size of physical memory)
    - `ada0p3`: 16 GB `freebsd-ufs` root `/`
    - `ada0p4`: 4 GB `freebsd-ufs` temp `/tmp`
    - `ada0p5`: 4 GB `freebsd-ufs` var `/var`
    - `ada0p6`: 32 GB `freebsd-ufs` usr `/usr`
    - `ada0p7`: 386 GB `freebsd-ufs` home `/home` (remainder of the space)
- Root password: I won't tell you, but a strong one!
- CMOS clock: UTC
- Time Zone: Europe/Switzerland (`CEST`)
- Services: `sshd`, `moused`, `ntpd`, `powerd`, and `dumpdev`
- Security Hardening Options: everything
- User: `patrick` with additional group `wheel` (to become root), the `tcsh`
  shell, a strong password and, otherwise, suggested settings

It can be argued whether or not the chosen partition sizes are reasonable.
However, it is always a good idea to use separate `/tmp` and `/var` partitions
to make sure that no process can fill up the entire disk.

Make sure to use `efi` as the type for the boot partition, not `freebsd-boot` as
suggested in _Absolute FreeBSD_ (3rd Edition on page 36).

# First Boot, First Issues

After the installation, I choose to shutdown the system. I unplug my USB dongle
as soon as the screen turns black.

Before the first boot, I have to change my boot options in the BIOS so that the
computer boots from the SSD on which FreeBSD was just installed.

The system boots and even shows my mouse on the terminal! The network is up and
running. However, there is a message warning me that the leapsecond file is
expired. The
[solution](https://forums.FreeBSD.org/threads/leapseconds-file-expired.56645/post-322290) suggested in the FreeBSD Forum

    # service ntpd onefetch

fails with a certificate verification error. A [bug
report](https://bugs.freebsd.org/bugzilla/show_bug.cgi?id=230017) suggest to
install the package `ca_root_nss`:

    # pkg install ca_root_nss

Which not only installs the package management software, but also solves the
issue above: The warning concerning the leapsecond file doesn't appear after the
next boot. Now that the basic system is up and running, let's tackle the GUI!

# Installing the GUI

Since _Absolute FreeBSD_ doesn't cover graphical user interfaces, I have to
resort to the handbook. In [Chapter
5.3](https://www.freebsd.org/doc/handbook/x-install.html) it says, that the
easiest way to setup the X Window System would be by installing the `xorg`
package. Since I prefer a minimalistic setup, I opt for the `x11/xorg-minimal`
package:

    # pkg install x11/xorg-minimal

Which depends on Python 3.7, Perl 5, and Wayland, among others, and weighs
roughly 1 GB, which is not exactly minimalistic, in my opinion. On the other
hand, it is notable that the base setup works without Perl or Python.

## Compiling `dwm`

Since I like to keep my `dwm` version up to date, I fetch the sources using
`git`, which first needs to be installed:

    # pkg install git
    # git clone https://git.suckless.org/dwm

A first naive compilation attempt fails:

    # cd dwm
    # make
    drw.c:5:10: fatal error: 'X11/Xlib.h' file not found
    #include <X11/Xlib.h>

The `config.mk` file expects the header files to be located under
`/usr/X11R6/include`. However, FreeBSD has those files stored under a different
location:

    # find / -type f -name Xlib.h
    /usr/local/include/X11/Xlib.h

So in `config.mk`, the lines

    X11INC = /usr/X11R6/include
    X11LIB = /usr/X11R6/lib

need to be replaced with

    X11INC = /usr/local/include
    X11LIB = /usr/local/lib

The next compilation fails with another error (different error message, yay!):

    # make
    drw.c:6:10: fatal error: 'X11/Xft/Xft.h' file not found

That's the price I have to pay for minimalism, I guess. Executing `pkg search
Xft` reveals the package `libXft`, which I install:

    # pkg install libXft

This shows to be a good idea, because now I'm getting a different error message:

    # make
    Xft.h:39:10: fatal error: 'ft2build.h' file not found

It turns out that the file is on the system, but cannot be found:

    # find / -type f -name ft2build.h
    /usr/local/include/freetype2/ft2build.h

Again, the `local` path segment is missing in `config.mk`:

    FREETYPEINC = /usr/include/freetype2

Which is changed as follows:

    FREETYPEINC = /usr/local/include/freetype2

Retry, fresh error again:

    # make
    dwm.c:40:10: fatal error: 'X11/extensions/Xinerama.h' file not found

The `config.mk` contains the following section:

    # Xinerama, comment if you don't want it
    XINERAMALIBS  = -lXinerama
    XINERAMAFLAGS = -DXINERAMA

So what is _Xinerama_, after all? According to
[Wikipedia](https://en.wikipedia.org/wiki/Xinerama):

    > Xinerama is an extension to the X Window System that enables X
    > applications and window managers to use two or more physical displays as
    > one large virtual display. 

Since I have only one screen, I can do without Xinerama, so I comment out those
lines:

    # Xinerama, comment if you don't want it
    # XINERAMALIBS  = -lXinerama
    # XINERAMAFLAGS = -DXINERAMA

Now `dwm` compiles, and I can install it:

    # make dwm
    # make install

## Starting Xorg with `dwm`

I switch to my personal account and create the file `/home/patrick/.xinitrc`
with the following content:

    exec dwm

Now I run `startx`, which unfortunately fails:

    Fatal server error:
    (EE) no screens found(EE)

The error log `/var/log/Xorg.0.log` does not offer any additional information
that seems helpful to me. It turns out that `/etc/X11` is empty. [Section
5.4](https://www.freebsd.org/doc/handbook/x-config.html) of the handbook is
about Xorg configuration. I create a minimalistic configuration for my graphics
card in `/etc/X11/xorg.conf`:

    Section "Device"
        Identifier "Card0"
        Driver     "intel"
    EndSection

I also need to install the display driver with the matching kernel module,
because my choice of `xorg-minimal` from before.

    # pkg install xf86-video-intel drm-kmod

The kernel modul can be activated on startup by adding it to the `rc.conf` as
follows:

    # echo 'lkd_list="/boot/modules/i915kms.ko"' >> /etc/rc.conf

After a restart, the console is shown in a much higher resolution. However,
`startx` now complains about a missing font. Let's install the `xorg-fonts` meta
package, which should provide a monospace font needed for `dwm`:

    # pkg install xorg-fonts

Now, finally, `dwm` works! Since `startx` is long to type, I define the alias
`x` for it in `~/.cshrc`:

    alias x     startx

And start `dwm`:

    $ x

## Configure `dwm`

By default, `dwm` uses the `Alt` as the modifier key (`MODKEY`). I prefer to use
the «Windows» or «Super» Key), for it has no purpose on my system. (`Alt` is
for some emacs-style readline commands.) To do this, the `MODKEY` variable has
to be changed in `config.h` as follows:

    #define MODKEY Mod4Mask

The default rules make Firefox to appear on the last tag, and Gimp to be used
with floating layout, which makes no sense with more recent versions of Gimp.
Let's just undefine those rules:

    static const Rule rules[] = {
        {NULL, NULL, NULL, 0, 0, -1},
    };

I also like my windows to be split evenly:

    static const float mact = 0.50;

As a terminal, let's use `qterminal` instead of `st`, for the latter does not
support scrollback buffers:

    static const char *termcmd[] = {"qterminal", NULL};

`qterminal` and `dmenu` need to be installed:

    # pkg install qterminal dmenu

## Status Line

`dwm` can display status information using the `xsetroot` command. The text to
be displayed is computed in a background task that can be defined in `.xinitrc`.
On laptops, I usually print the battery status. On desktops, the current date
and time suffices. Here's the `.xinitrc` that displays this information
(surrounded by spaces) in five second intervals:

    while true
    do
        xsetroot -name " $(date +'%Y-%m-%d %H:%M') "
    done &
    setxkbmap ch
    exec dwm

The keymap is also set to the `ch` (i.e. Swiss German) variant just before
executing `dwm`. The `xsetroot` and `setxkbmap` utilities need to be installed
for this:

    # pkg install xsetroot setxkbmap

## Volume Control

In order to test audio, let's download the Free Software Song:

    $ curl https://www.gnu.org/music/free-software-song.ogg > fss.ogg

I prefer `mplayer`, which needs to be installed:

    # pkg install mplayer

Make sure to include `/usr/local/bin` in your `$PATH` variable in order to run
`mplayer` without further path specification (`.cshrc`):

    export PATH="$PATH:/usr/local/bin"

Playing the song as follows works if i plug in a headphone into one of the front
plugs:

    $ mplayer fss.ogg

The devices are listed in `/dev/sndstat` and switched by setting the respective
device number:

    # sysctl hw.snd.dfault_unit=1

The default volume is set to 85, which is quite loud for Richard Stallman's
singing voice. The volume can be changed relatively or absolutely using the
`mixer` command:

    $ mixer vol -10
    Setting the mixer from 85:85 to 75:75
    $ mixer vol 50
    Setting the mixer from 75:75 to 50:50

I don't always want to type that command, but rather use the volume keys on my
keyboard. So let's add a couple of commands to the `dwm` config (`config.h`,
just before the `keys[]` section):

    static const char *upvol[] = {"mixer", "vol", "+5"});
    static const char *downvol[] = {"mixer", "vol", "-5"});
    static const char *mutevol[] = {"mixer", "vol", "0"});

For the key mapping, I first need to figure out the key codes for my volume
keys, which can be done using `xev`:

    # pkg install xev
    $ xev > xev.out

Just press the volume down, volume up, and mute button in that order. Then close
the `xev` window and inspect `xev.out`:

TODO: the keys are not detected
