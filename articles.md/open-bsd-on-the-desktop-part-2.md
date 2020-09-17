---
title: OpenBSD on the Desktop (Part II)
subtitle: GUI Tweaks et cetera
author: Patrick Bucher
date: 2020-09-12T15:00:00
lang: en
---

A week ago, I've installed [OpenBSD on my
Thinkpad](https://paedubucher.ch/articles/2020-09-05-openbsd-on-the-desktop-part-i.html).
I've been using it now and then, and already have changed a couple of things in
respect to the original setup described in the article. I also installed OpenBSD
on the Dell Optiplex on which I [previously installed
FreeBSD](https:///paedubucher.ch/articles/2020-08-11-freebsd-on-the-desktop.html)
a month before. This means that I'm no longer using FreeBSD on the desktop, at
least not for the moment. However, FreeBSD is running on a disk station I built
earlier this summer. Maybe I'll describe that particular setup (using ZFS) in a
later article.

Except for that storage server, I'd like to use OpenBSD for most of my private
computing. In this article, I describe some GUI tweaks and additional setup
tasks I perfmormed in order to feel more at home on my OpenBSD machines. Some of
the tasks performed are _not_ specific to OpenBSD, but could also be applied to
a Linux setup.

# doas

`sudo` originally came from the OpenBSD community. It is almost as widely used
in the Unix world as SSH, which is the most prominent OpenBSD project.  However,
`sudo` became bigger and harder to configure. Therefore, Ted Unangst came up
with a simpler alternative called `doas`, which stands for _Dedicated OpenBSD
Application Subexecutor_. `doas` is less powerful than `sudo`, but much smaller,
easier to configure, and, thus, more secure. The full rationale can be read in
[Ted Unangst's Blog](https://flak.tedunangst.com/post/doas).

A basic `doas` setup requires to login as root for one last time. The
configuration shall be kept extremely simple. I'd like to permit all users from
the `wheel` group (which is just me on my computers) to use `doas` without
entering the password every time but only once when executing a command that
requires `root` permissions. This is only a single line in `/etc/doas.conf`:

    permit persist :wheel    

Let's check this setup by logging in as a user of the wheel group and trying to
update the packages:

    $ doas pkg_add -u

This works, so bye bye `root` account.

# Fonts for dwm, dmenu, and st

By default, `dwm`, `dmenu`, and `st` use a monospace font of size 10, or
pixelsize 12, respectively, which is hard to read on a screen with a high
resolution. On Linux, I use the the TrueType font DejaVu Sans Mono. For OpenBSD,
I'd rather use something more minimalistic: the [Terminus bitmap
font](http://terminus-font.sourceforge.net/).

As `pkg_info -Q terminus` shows, this font comes in different versions. I prefer
the version with the centered tilde, which I install:

    $ doas pkg_add terminus-font-4.47p0-centered_tilde

Let's reconfigure `st` first, for testing changes doesn't require a restart of
the window manager. I stored my suckless sources in `~/suckless`, so the
font for `st` can be configured in `~/suckless/config.h`. I replace the existing
font configuration

    static char *font = "Liberation Mono:pixelsize=12:antialias=true:autohint=true";

with

    static char *font = "Terminus:pixelsize=24";

The options `antialias` and `autohinting` are not needed for a bitmap font, so I
left them away. 24 pixels is rather big, but my screen is big enough to show two
text editors with more than 80 characters per line next to each other, so let's
keep it this way. I rebuild and reinstall `st`, then switch to `dwm`:

    $ doas make install
    $ cd ../dwm

The font configuration in the `config.h` file looks a bit different here:

    static const char *fonts =      { "monospace:size=10" };
    static const char dmenufont =   "monospace:size=10";

Let's just use the same font as for `st` here:

    static const char *fonts =      { "Terminus:pixelsize=24" };
    static const char dmenufont =   "Terminus:pixelsize=24";

Note that I'm using `pixelsize` instead of `size` here. (24pt would be much
bigger than 24px.) Then I rebuild and reinstall `dwm`.

    # make install

This configuration appllies also to `dmenu` and `slstatus`, so we're done with
the fonts.

# X Background

By default, the desktop background is a pattern of black and grey dots, which is
a strain to the eye. Even though I rarely look at an empty desktop for long, I'd
rather change this to a solid color. This can be done by adding a command to
`~/.xinitrc`:

    xsetroot -solid black

Right before `dwm` is executed.

# USB Flash Drive

Even though SSH is almost ubiquitous nowadays, a USB flash drive is still useful
when it comes to exchanging data between computers, especially if Windows is
involved, or if the network does not allow SSH.

Block storage devices are accessible through the device nodes `/dev/sd*`,
whereas `*` stands for the number of the disk. The disks can be listed as
follows:

    $ sysctl hw.disknames
    hw.disknames=sd0:ef0268c97ae7a246

Only `sd0` is active, even though I already plugged in my USB dongle. However,
the system already figured out that there is a second disk:

    $ sysctl hw.diskcount
    hw.diskcount=2

The next free disk would have the name `sd1`. The device nodes can be created by
running the `MAKEDV` script in `/dev`:

    $ cd /dev
    $ doas sh MAKEDEV sd1

Let's initialize a new MBR partition schema on `sd1`:

    $ doas fdisk -iy sd1

The new disk layout can be checked using `disklabel`:

    $ doas disklabel sd1
    # /dev/rsd1c
    ...

The first line of the output tells us that there's a partition under
`/dev/rsd1c`. (The `r` refers to «raw», as opposed to «block».) The partition
can be formatted using `newfs` by referring to that partition name:

    $ doas newfs sd1c

This creates a default FFS (Fast File System) partition, which is useful to
exchange data between BSD operating systems. The formatted partition is then
ready to be mounted:

    $ doas mount /dev/sd1c /mnt

## Other Partition Types

Other partition types are available under other utilities.

### FAT32

The following command creates a FAT32 partition:

    $ doas newfs_msdos -F 32 sd1c

The `-F 32` parameter specifies FAT32 (as opposed to FAT16 or FAT8). To mount
the partition, use the according `mount` command:

    $ doas mount_msdos /dev/sd1c /mnt

### EXT2

In order to create an `ext2fs` file system, the partition type needs to be
specified accordingly. First, you might consider a GPT partition schema instead
of MBR (additional `-g` parameter):

    $ doas fdisk -igy sd1

Then use `disklabel` interactively to define a new partition:

    $ doas disklabel -E sd1

First, delete all the partitions with `z`. Then, create a new partition with
`a`, and make sure to specify the type as `ext2fs` instead of the default
`4.2BSD`. Notice that the new partition has a different letter (say, `a`), so
you need to use `sd1a` instead of `sd1c` for the next steps. Write the changes
by typing `w`, then exit with `q`. Now you can format and mount the partition:

    $ doas newfs_ext2fs sd1a
    $ doas mount_ext2fs /dev/sd1a /mnt

# SSH Key (GitHub)

In order to access my GitHub repositories, I first create a new SSH key:

    $ ssh-keygen -t rsa -b 4096

Since I manage my passwords with `pass` (of which more later), I don't know most
of them by heart. So I can't just login to GitHub and add my public key.
Therefore, I copy my public key to my laptop, on which I'm already logged in to
GitHub.

This can be either done using `scp`, for which `sshd` has to be running on my
laptop (which currently has the IP `192.168.178.53`):

    $ scp ~/.ssh/id_rsa.pub 192.168.178.53:/home/patrick

Or using the USB flash drive formatted with `ext2` from before:

    $ doas newfs_ext2fs -I sd1a
    $ doas mount_ext2fs /dev/sd1a /mnt
    $ doas cp ~/.ssh/id_rsa.pub /mnt/

Then `id_rsa.pub` can be copied into the according [GitHub Settings
Page](https://github.com/settings/ssh/new), after which cloning GitHub
repositories should work on the OpenBSD machine:

    $ git clone git@github.com:patrickbucher/conf

# GPG Key

My passwords are encrypted using GPG. To encrypt them, I need to copy my private
key from my other machine. First, I list my private keys:

    $ gpg --list-keys --keyid-format SHORT
    pub   rsa2048/73CE6620 2016-11-11 [SC]
          22F91EE20D641CBCF5B8678E82B7FE3A73CE6620
    uid         [ultimate] Patrick Bucher <patrick.bucher@mailbox.org>
    sub   rsa2048/AF6246E3 2016-11-11 [E]

Then I export both public and private key to an according file using the armored
key format:

    $ gpg --export --armor 73CE6620 > public.key
    $ gpg --export-secret-key --armor 73CE6620 > private.key

The two key files can be copied via SSH or the USB flash disk again, which I
won't show here.

Back on my OpenBSD machine, I need to install GnuPG first, because OpenBSD only
has `signify` installed by default:

    $ doas pkg_add gnupg

I pick the 2.2 version. Now I can import my keys:

    $ gpg2 --import private.key
    $ gpg2 --import public.key

The key is not trusted so far, so I need to change that:

    $ gpg2 --edit-key 73CE6620
    > trust
    > 5
    > y
    > quit

5 stands for ultimate trust, which seems appropriate.

# Password Manager

I use `pass` as a password manager, which can be installed as the
`password_store` package in OpenBSD:

    $ doas pkg_add password-store

Now that I have both my GPG private key and a working SSH key for GitHub, I can
clone my passwords stored on a private GitHub repository:

    $ git clone git@github.com:patrickbucher/pass .password-store

Now I can copy my GitHub password to the clipboard as follows:

    $ pass -c github

# Aliases

I use a lot of aliases, such as `gcl` as a shortcut for `git clone`, and `gad`
for `git add`, etc. Since OpenBSD uses a Public Domain Korn Shell by default,
the `.bashrc` configuration from my Linux machines won't work here, unless I
switch to `bash`, which is not exactly the point of using OpenBSD.

I define my aliases in `~/.kshrc` (excerpt):

    alias gcl='git clone'
    alias gad='git add'

In order to load those settings, an according `ENV` parameter needs to be
defined in `~/.profile` (see `man 1 ksh` for details):

    export ENV=$HOME/.kshrc

After the next login, `~/.profile` is reloaded, and the aliases are ready to be
used.

# Conclusion

Not only is my enhanced setup now ready to do some serious work, but I also
increased my understanding of some OpenBSD subjects. There are still things to
be improved and to be understood, but my setup is now good enough so that I no
longer need a Linux machine running next to it. I'm looking forward to use and
learn about OpenBSD in the time to come. I'll write additional articles on the
subject as soon as I have enough subject material ready.
