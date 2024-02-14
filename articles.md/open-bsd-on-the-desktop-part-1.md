---
title: OpenBSD on the Desktop (Part I)
subtitle: Basic Setup with Xorg and dwm
author: Patrick Bucher
date: 2020-09-05T20:00:00+0100
lang: en
---

Let's install OpenBSD on a Lenovo Thinkpad X270. I used this computer for my
computer science studies. It has both Arch Linux and Windows 10 installed as
dual boot. Now that I'm no longer required to run Windows, I can ditch the dual
boot and install an operating system of my choice.

# Preparation

First, I grab my work Thinkpad running Arch Linux and some USB dongle big enough
for the [amd64 miniroot
image](https://cdn.openbsd.org/pub/OpenBSD/6.7/amd64/miniroot67.fs) (roughly
five megabytes, that is). This small image does not include the file sets, which
will be downloaded during installation instead. I also download the [SHA256
checksums](https://mirror.ungleich.ch/pub/OpenBSD/6.7/amd64/SHA256) from the
Swiss mirror, and verify the downloaded image, before I copy it on my dongle:

    $ sha256sum -c --ignore-missing SHA256 
    miniroot67.fs: OK
    $ sudo dd if=miniroot67.fs of=/dev/sda bs=1M

# Installation

The Thinkpad X270 is connected to my network through Ethernet. The WiFi firmware
usually needs to be installed separately, so only Ethernet will work out of the
box. The BIOS has UEFI activated. OpenBSD and UEFI has issues on older hardware
(at least on a 2014 Dell laptop I have), but let's try it on this laptop,
anyway.

I plug in the dongle prepared before, and start the computer. I interrupt
the regular boot with Enter and pick an alternative boot method by pressing F12.
Now I pick my USB dongle. After roughly a minute, the installer has been
started. Now I follow these steps:

- I choose the option `I` to install OpenBSD.
- For the keyboard layout, I pick `sg`, for Swiss German.
- As a hostname, I simply pick `x270`, because it's a Thinkpad X270, and I'm not
  very creative when it comes to naming things.
- From the available network options (`iwm0`: WiFi, `em0`: Ethernet, and
  `vlan0`: Virtual LAN), I pick `em0`.
- I try to get an IPv4 address over DHCP, which seems to work very quickly.
- Next, I type in my very secret root password twice.
- I do _not_ start `sshd` by default, because I don't need to connect to this
  machine through SSH. It's supposed to be a workstation, not a server.
- The X Window System should not be started by `xnodm(1)`, so I leave it to
  `no`.
- Neither do I want to change the default to `com0`.
- I set up my user `patrick` with my proper name `Patrick Bucher`, and a decent
  password.
- The time zone has been detected properly as `Europe/Zurich`, which I just
  leave the way it is.
- The installer detected two disks: `sd0` and `sd1`. Since `sd0` is the detected
  SSD in my laptop, the UEFI issue from my Dell laptop doesn't exist on this
  computer. I pick `sd0` for the root disk, since `sd1` is my USB dongle.
- I choose to use the whole disk with a GPT partitioning schema, because it's
  2020.
- An auto-allocated layout for `sd0` is presented. It looks decent to me, so I
  just go with that auto layout.
- I don't want to initialize another disk, so I just press Enter (`done`).
- Since the miniroot image does not come with the file sets, I pick `http` as
  the location for the sets.
- I don't use a proxy, and use the mirror `mirrog.ungleich.ch` and the server
  directory `pub/OpenBSD/6.7/amd64` as proposed.
- Next, I unselect the game sets by entering `-game*`. (I heard that they're not
  much fun to play.) I leave all the other sets activated, including the `x`
  sets, which will be required for the GUI later on.
- After those sets are installed, I press Enter (`done`). Now the installer
  performs various tasks, after which I choose to `halt` the computer. This
  gives me time to remove the USB dongle.

# First Boot

I now restart my laptop, and OpenBSD boots. This takes more time than booting
Arch Linux, which uses `systemd`, whereas OpenBSD uses `rc`, which performs the
startup tasks sequentially.

There's a message showing up that various firmware (`intel-firmware`,
`iwm-firmware`, `inteldrm-firmware`, `uvideo-firmware`, and `vmm-firmware`) has
been installed automatically. Very nice, indeed.

## WiFi Connection

Now that the `iwm-firmware` has been installed, I can connect right away to my
WiFi network `frzbxpdb5`. I create a file called `/etc/hostname.iwm0`, wich
`hostname` being a literal string, and `iwm0` being the WiFi network card. The
connection to my WiFi network consists of a single line:

    dhcp nwid frzbxpdb5 wpakey [my-wpakey]

Whereas `frzbxpdb5` is my WiFi network's ESSID, and `[my-wpakey]` needs to be
replaced by the actual WPA key.

Then the networking can be restarted for that device:

    # sh /etc/netstart iwm0

This script is kind enough to set the file permissions of `/etc/hostname.iwm0`
to `640`, and then connects to my WiFi network.

I unplug the Ethernet cable and `ping openbsd.org`, which works fine, even after
a restart.

# Installing the GUI

My GUI on Unix-like systems is based on the Dynamic Window Manager (`dwm`) and a
couple of other tools, such as `dmenu`, `st`, `slstatus`, `slock`, all created and
maintained by the [Suckless](http://suckless.org/) community.

This software doesn't come with configuration facilities, but needs to be
configured in the respective C header file `config.h`, and then re-compiled.
Even though OpenBSD offers `dwm` as a package, customizing and configuring that
window manager requires to build it from source.

## Building `dwm` and Friends

First, I need to install `git` to fetch the source code:

    # pkg_add git

Then I fetch the source code for `dwm`, `dmenu`, `st`, and `slstatus` from [Suckless](http://suckless.org/):

    $ git clone https://git.suckless.org/dwm
    $ git clone https://git.suckless.org/dmenu
    $ git clone https://git.suckless.org/st
    $ git clone https://git.suckless.org/slstatus

### Building `dwm`

Next, I try to build `dwm`:

    $ cd dwm
    $ make

This fails with an error message (`'ft2build.h' file not found`), which reminds
me of building `dwm` on FreeBSD roughly a month before. Since I can finde the
header file at another location:

    # find / -type f -name ft2build.h
    /usr/X11R6/include/freetype2/ft2build.h

I simply can modify the `config.mk` accordingly by changing

    FREETYPEINC = /usr/include/freetype2

to

    FREETYPEINC = $(X11INC}/freetype2

Actually, I only need to comment the above line, and uncomment the line below

    # OpenBSD (uncomment)

The Suckless folks obviously are friendly towards OpenBSD, which is also
noticable in other places (more evidence to be shown further below).

The next compilation attempt succeeds:

    $ make

So let's install `dwm`, too:

    # make install

By default, and as to be seen in `config.h`, the keyboard combination
`[Alt]+[Shift]+[Enter]` (deeply engraved into the muscle memories of many `dwm`
users) starts the `st` terminal. This will be built in a while. However, I
prefer to use the _Super_ or _Windows_ key instead of `Alt`, since the former
is of no use in OpenBSD, and the latter still comes in handy when working with
the emacs readline mode. Therefore, I change the `MODKEY` from

    #define MODKEY Mod1Mask

to

    #define MODKEY Mod4Mask

Then I rebuild and reinstall `dwm`:

    # make install

### Building `st`

Let's switch over to the `st` source directory and just try to compile it:

    $ cd ../st
    $ make

Here, we get a warning that the function `pledge` (an OpenBSD mitigation, which
is built into the `master` branch, but surrounded by an `ifdef` preprocessor
statement, so that it will only be compiled for OpenBSD) is imported implicitly.
Let's just ignore this warning for now.

What's worse, the compilation fails with the error message:

    ld: error: unable to find library -lrt

Here, the FAQ comes in handy, stating that

    If you want to compile st for OpenBSD you have to remove -lrt from
    config.mk, ...

Having done so in `config.mk`, `st` compiles without any further issues, and,
thus, can be rebuilt and installed:

    # make install

### Building `dmenu`

Even OpenBSD users with Suckless tools have to open another GUI application than
a terminal emulator once in a while. For this purpose, Suckless offers `dmenu`.
Let's switch over to it and compile it:

    $ cd ../dmenu
    $ make

Again, we have the issue with `ft2build.h`, which can be resolved as above with
`dwm`: by using the proper path for `FREETYPEINC` in `config.mk`. Afterwards,
the build succeeds, and `dmenu` can be installed:

    # make install

### Building `slstatus`

`dwm` has a status bar on the top right, which can be used to show various
information. I used to write some shell commands in `.xinitrc` to compose such a
status line, and then set it by `xset -b` once every five seconds or so. This
approach generates a multitude of processes every couple of seconds.

`slstatus` is a C programm that is capable of showing various kinds of more or
less useful information. Let's switch over to `slstatus` and see, what is
available in `config.def.h`:

    $ cd ../slstatus
    $ less config.def.h

The comments section lists different functions (`battery_perc` for the battery
percentage, `datetime` for date and time information, `temp` for thermal
information, etc.). I usually display the CPU load, the battery percentage, the
memory usage, the current keyboard layout, and the current date and time.

Before configuring those, let's try to compile `slstatus`:

    $ make

This worked fine, so let's configure the information to be displayed in
`config.h`:

    static const struct arg args[] = {
        /* function    format    argument */
        { datetime,    "%s",     "%F %T" },
    };

This renders the current date as follows:

    $ date +"%F %T"
    2020-09-05 19:26:38

I also like to have the weekday included, but not the seconds, so I define a
different argument string:

    $ date +"%a %Y-%m-%d %H:%M"
    Sat 2020-09-05 19:27

That's better, so let's use it in `config.h` (surrounded with some spaces in the
format string):

    static const struct arg args[] = {
        /* function    format    argument */
        { datetime,    " %s ",   "%a %Y-%m-%d %H:%M" },
    };

The other settings I like to have do not require any arguments, at least not on
OpenBSD, so I only need to define a decent format string (with `|` as a
seperator) for those:

    static const struct arg args[] = {
        /* function    format           argument */
        { cpu_perc,     " cpu: %s%% |", NULL },
        { battery_perc, " bat: %s%% |", NULL },
        { ram_used,     " mem: %s |",   NULL },
        { keymap,       " %s |"         NULL },
        { datetime,     " %s ",         "%a %Y-%m-%d %H:%M" },
    };

This actually compiles, so let's install it:

    # make install

## Configuring X Startup

Now that all software is compiled and installed, let's run X. To do so, a file
`.xinitrc` in the user's directory is required (`/home/patrick/.xinitrc`):

    setxkbmap ch
    slstatus &
    exec dwm

This sets the keyboard map for X to Swiss German, starts `slstatus` in the
background, and then executes `dwm`.

X can now be started by typing `startx`. This is a bit cumbersome to type every
time, so let's define a symbolic link to it:

    # ln -s "$(which startx)" /usr/bin/x

Now let's start X:

    $ x

If everything was configured properly, `dwm` shows up, and the status line says
that the whole system only uses roughly 60 megabytes of RAM. That's slim. The
keyboard combinations to open `st` and `dmenu` work, too.

# Conclusion

Installing a basic GUI with Suckless software was a rather smooth experience on
OpenBSD. (For FreeBSD, I deliberately have chosen a rather fine-grained approach
to installing X packages, which caused some additional work.) However, various
settings require additional tweaking. I also didn't use audio yet, which require
the volume buttons to be configured accordingly in `dwm`.

I'll also need to setup `sudo` or `doas`. As a regular Linux user, I'm used to
`sudo`, of course, but the simplicity of `doas` is a good argument to try it as
an alternative.

But those are things I'd like to cover in an upcoming article.
