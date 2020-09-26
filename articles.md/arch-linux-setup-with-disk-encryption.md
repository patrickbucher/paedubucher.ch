---
title: Arch Linux Setup with Disk Encryption
subtitle: Using dm-crypt, LUKS, and the systemd Boot Loader
author: Patrick Bucher
date: 2020-09-26T13:00:00
lang: en
---

I've been using Arch Linux since 2016. I got to understand the system better
since then, and installing Arch nowadays is rather a strain on my fingers than
on my brain. I automated my personal setup procedure to some extent with a
couple of shell scripts, which I'm adjusting as time goes and hardware changes.
I collect those scripts and setup instructions in a [GitHub
repo](https://github.com/patrickbucher/docs/tree/master/arch-setup).

Even though I've often taken my laptop with me, especially for
university, I never made the effort to encrypt my disk or a single partition
thereon. Sensitive data, such as passwords, are stored encrypted using GPG. And
my GPG private key is protected with a strong passphrase; so strong, that I
mistype it once in a while.

Having stored SSH keys on my laptop, however, is a risk in terms of data
protection. I do not protect most of my SSH keys with passwords. (_Not_ having
to use passwords is one of the conveniences SSH gives you, after all.) But some
of those SSH keys are gateways to quite some important data, and therefore I
need to re-consider my security practices.

Protecting SSH keys with a passphrase is very inconvenient, especially if you do
a lot of `git push` operations via SSH. It's possible to use a password
protected SSH key for `git` only, but having to deal with multiple SSH keys is
inconvenient, too. (My rule is to use one SSH key per client machine.) So
encrypting the hard disk or at least some partitions of it might be the better
solution, because it only requires me to enter an additional password once
during startup, which is once or twice a day.

# dm-crypt and LUKS

The Arch wiki provides an article on [how to encrypt an entire
system](https://wiki.archlinux.org/index.php/Dm-crypt/Encrypting_an_entire_system).
Such a setup is based on LUKS (Linux Unified Key Setup), a disk encryption
specification, and dm-crypt, the Linux kernel's device mapper for encrypted
devices.

The wiki mentions different
[scenarios](https://wiki.archlinux.org/index.php/Dm-crypt/Encrypting_an_entire_system#Overview)
for disk encryption. The first option, [LUKS on a
partition](https://wiki.archlinux.org/index.php/Dm-crypt/Encrypting_an_entire_system#LUKS_on_a_partition),
is very simple, but not ideal for me, because I'm not just using one giant root
partition, but separate partitions for `/var`, `/tmp` and `/home`. With this
setup, a process filling up my `/var/log` directory won't prevent me from
writing to `/home` or the root partition, which might be necessary to deal with
the issue.

One could argue that it's sufficient to just encrypt the `/home` partition,
which contains the most sensitive data, after all. However, it's possible that
some program logs sensitive data to `/var/log` or stores it under `/opt`, so
that an encrypted `/home` partition won't prevent sensitive data from being
leaked. An unencrypted root partition also makes it possible to tamper
programs, so that data is leaked later.

It's also possible to encrypt the partitions mentioned seperately, which would
require me to enter my password upon booting for every single partition. This is
highly impractical, and therefore I won't do it that way.

The second option, [LVM on
LUKS](https://wiki.archlinux.org/index.php/Dm-crypt/Encrypting_an_entire_system#LVM_on_LUKS),
is a better fit for my purpose. Here, LUKS provides one single block device, on
top of which the encrypted partitions are created. The other, arguably more
sophisticated scenarios, are mostly useful when dealing with multiple disks, or
if an encrypted boot partition is required. So let's stick to scenario two: LVM
on LUKS.

# Partitioning

My Lenovo Thinkpad X1 has a SSD with roughly 475 GB and 16 GB memory. This is
important when considering partition sizes. I usually create my partitions as
follows:

1. A 256 MB `/boot` partition for the bootloader (FAT 32).
2. A 16 GB (the size of my memory) swap partition.
3. A 128 GB `/` (root) partition (ext4).
4. A 64 GB `/var` partition (ext4).
5. A 8 GB `/tmp` partition (ext4).
6. And a `/home` partition with the remainder of the space, i.e. roughly 260 GB
   in my case (ext4).

The choice of partition sizes is subject to debate. Some prefer to have a swap
partition with twice the size of the physical memory. Some prefer to make the
root, `/var/`, and `/tmp` partitions smaller or bigger. But the above
partitioning scheme was never the cause of any trouble for me, so far.

With those decisions taken, let's proceed to the setup.

# Setup Procedure

So let's go through the setup step by step.

## Prepare the Installation Medium

First, I download the latest Arch amd64 image via BitTorrent [Magnet
link](magnet:?xt=urn:btih:db56a13a6555179990837759ca27274d0be49aca&dn=archlinux-2020.09.01-x86_64.iso)
from the [download](https://www.archlinux.org/download/) page and verify its
checksum against the one on the website:

    $ sha1sum archlinux-2020.09.01-x86_64.iso
    95ebacd83098b190e8f30cc28d8c57af0d0088a0

Then I copy the image on a USB dongle with a capacity of at least 700 MB:

    $ sudo dd if=archlinux-2020.09.01-x86_64.iso of=/dev/sda bs=4M
    $ sync

Then I unplug my USB dongle and plug it into the Thinkpad X1 I'd like to setup,
and boot from it (UEFI boot). Once the installation environment is loaded, let's
continue with the setup.

My Thinkpad X1 has a US keyboard, which I [prefer to
use](https://paedubucher.ch/articles/2020-09-16-compose-key-on-x.html) nowadays.
However, it's possible to change the keyboard layout as follows (for Swiss
German), if wanted:

    # loadkeys de_CH-latin1

Next, I establish a WiFi connection with `frzbxpdb5` being my network SSID, 
`wlan0` being my WiFi device, and `[topsecret]` being my password:

    # iwctl --passphrase '[topsecret]' station wlan0 connect frzbxpdb5

Let's continue with the partitioning of the disk.

## Partitioning

In order to partition the hard disk, one first needs to know the device name,
which can be found using `lsblk`:

    # lsblk

This lists `sda`, the USB dongle, and `nvme0n1`, the 475 GB solid state drive
aforementioned.

### Filling the Disk with Junk Data

As a first step, the disk shall be overwritten with random data.  Even though I
haven't used this laptop productively yet, and so there's no sensitive data
stored on it, it is still a good practice to fill it up once with random data.
This makes it harder for an attacker to distinguish between encrypted data and
random junk.

    # shred --random-source=/dev/urandom --iterations=1 /dev/nvme0n1

In my case, using only one single iteration is sufficient. If there was actual
data stored on the disk, multiple iterations should be considered.

### Creating a Boot Partition

Next, I create a new GPT partition scheme on that disk:

    # parted -s /dev/nvme0n1 mklabel gpt

I'm not going to encrypt my boot partition, so I create it as I do for a regular
setup without encryption (as a FAT 32 boot partition, that is):

    # parted -s /dev/nvme0n1 mkpart boot fat32 1MiB 257MiB
    # parted -s /dev/nvme0n1 set 1 esp on
    # mkfs.fat -F 32 /dev/nvme0n1p1

I leave a gap of 1 MB at the beginning, so no matter what block size my SSD
uses, the boot partition will always be properly aligned. The `esp` flag
identifies the partition as a UEFI system partition.

### Creating a Partition for Encryption

Now comes the crucial part. All the encrypted data is put into a single big
partition (`nvmen1p2`), taking up all of the remaining disk space, which is then
partitioned using a volume manager:

    # parted -s /dev/nvme0n1 mkpart cryptlvm 257MiB '100%'

The encryption is set up on this partition:

    # cryptsetup luksFormat /dev/nvme0n1p2

Enter "YES" if asked for confirmation, and pick a strong passphrase to be
entered twice. (At this point, remember exactly which keyboard layout you're
on!).

To further work with the encrypted partition, let's open it, which requires to
enter the password chosen before:

    # cryptsestup open /dev/nvme0n1p2 cryptlvm

Now a physical volume for the volume mapping needs to be created:

    # pvcreate /dev/mapper/cryptlvm

The partitions are going to be managed in a volume group, which I simply call
`volgrp` for the sake of brevity:

    # vgcreate volgrp /dev/mapper/cryptlvm

Now everything is set up to create the remaining, i.e. the encrypted partitions.

### Creating the Remaining Partitions

Those are created using `lvcreate` by setting a size (`-L`/`-l` parameter) and a
name (`-n` parameter):

    # lvcreate -L 16G volgrp -n swap
    # lvcreate -L 128G volgrp -n root
    # lvcreate -L 64G volgrp -n var
    # lvcreate -L 8G volgrp -n tmp
    # lvcreate -l '100%FREE' volgrp -n home

The partitions are going to be formatted using the `swap` and `ext4` format,
respectively:

    # mkswap /dev/volgrp/swap
    # mkfs.ext4 -F /dev/volgrp/root
    # mkfs.ext4 -F /dev/volgrp/var
    # mkfs.ext4 -F /dev/volgrp/tmp
    # mkfs.ext4 -F /dev/volgrp/home

For the actual setup, those partitions (and the `/boot` partition created
before) need to be mounted to `/mnt`:

    # mount /dev/volgrp/root /mnt

    # mkdir /mnt/boot
    # mount /dev/nvme0n1p1 /mnt/boot

    # swapon /dev/volgrp/swap

    # mkdir /mnt/var
    # mount /dev/volgrp/var /mnt/var

    # mkdir /mnt/tmp
    # mount /dev/volgrp/tmp /mnt/tmp

    # mkdir /mnt/home
    # mount /dev/volgrp/home /mnt/home

Now the partitions are ready for a regular bootstrap installation. (Setting up
the boot loader will require some more specific instructions to disk encryption
later on.)

## Bootstrap Installation

Now let's install the base system, together with the `lvm2` package:

    # pacstrap /mnt base linux linux-firmware lvm2

In order to get the mounting done automatically upon restart, let's save it in
the `fstab` file:

    # genfstab -U /mnt >> /mnt/etc/fstab

When this is done, let's switch into the installed environment:

    # arch-chroot /mnt

A password needs to be set for the `root` user:

    # passwd

In order to have a WiFi connection after rebooting, let's install a couple of
networking packages (some of those specific to my Intel hardware):

    # pacman -S iw wpa_supplicant dialog intel-ucode netctl dhcpcd

I set the time zone to Zurich (Europe), update and save the system time:

    # ln -sf /usr/share/zoneinfo/Europe/Zurich /etc/localtime
    # timedatectl set-ntp true
    # hwclock --systohc

For language and locale, I simply use `en_US.UTF-8` and `en_US-UTF-8 UTF-8`,
respectively:

    # echo 'en_US.UTF-8 UTF-8' >> /etc/locale.gen
    # locale-gen
    # echo 'LANG=en_US.UTF-8' > /etc/locale.conf

Due to the lack of imagination, I call my laptop simply _carbon_:

    # echo carbon > /etc/hostname

This is a very basic setup. Now let's make sure it can be booted by installing
the boot loader:

## Configuring the Boot Loader

I've always been using the systemd boot loader on Arch Linux, which is quite
simple to configure. First, the computer needs to get a unique id, then the boot
loader can be installed into the `/boot` partition:

    # systemd-machine-id-setup
    # bootctl --path=/boot install

Now comes the tricky part: The UUID of the boot partition needs to be figured
out. `blkid` lists various partitions, but which one to choose? It's the
LUKS partition containing the encrypted volume: `/dev/nvme0n1p2`. Its UUID can
be extracted as follows, and shall be saved into a variable for later use:

    # uuid=$(blkid | grep 'crypto_LUKS' | egrep -o ' UUID="[^"]+"')
    # uuid=$(echo $uuid | awk -F '=' '{ print $2 }' | tr -d '"')

The first line lists the partitions (`blkid`), extracts the line with the
encrypted partition (`grep`), and further extracts the part of the line defining
the UUID (`egrep`). Don't forget the space in front of `UUID`, otherwise the
`PARTUUID` is extracted, too. In the second line, the definition
(`UUID="abc..."`) is split at the equal sign, of which the second part
(`"abc..."`) is taken using `awk`. Then the surrounding double quotes are
removed with `tr`. The variable `$uuid` now contains the UUID of the encrypted
partition.

**Update**: As the user _fra-san_ [pointed
out](https://unix.stackexchange.com/a/611507/223188), there's an easier way to
extract the UUID:

    # uuid=$(blkid --match-tag UUID -o value /dev/nvme0n1p2)

Having this information, the entry for the boot loader can be created as
follows:

    # cat <<EOF >/boot/loader/entries/arch.conf
    title   Arch Linux
    linux   /vmlinuz-linux
    initrd  /initramfs-linux.img
    options cryptdevice=UUID=${uuid}:cryptlvm root=/dev/volgrp/root
    EOF

Now a simple bootloader configuration needs to be created:

    # cat <<EOF >/boot/loader/loader.conf
    default arch
    timeout 0
    editor  0
    EOF

## Creating the Ramdisk Environment

As mentioned earlier, there are some further adjustments for disk encryptions to
be made. Let's open `/etc/mkinitcpio.conf` and go to the `HOOKS` definition:

    HOOKS=(base udev autodetect modconf block filesystems keyboard fsck)

Which needs to be extended with `encrypt` and `lvm2`:

    HOOKS=(base udev autodetect modconf block filesystems keyboard fsck encrypt lvm2)

Then the ramdisk environment can be created:

    # mkinitcpio -P

Which generated the `initramfs-linux.img` file referred to from the boot loader
entry created before, as well as a fallback `initramfs-linux-fallback.img`.

## Finishing

Now that everything is set up, leave the `chroot` environment, unmount the new
system's partitions, and shut down the computer:

    # exit
    # umount -R /mnt
    # shutdown -h now

After removing the USB dongle, start the system again.

Early in the boot process, you'll be asked to enter the passphrase for
`/dev/nvme0n1p2`. After doing so, the system will boot. If not, you did
something wrong in the process, because it's working just fine on my Thinkpad X1
Carbon.

I also had issues getting everything right in the first place, but fortunately
got help on the [Linux & Unix
StackExchange](https://unix.stackexchange.com/q/611421/223188). Thanks to the
user _frostschutz_ for pointing out the core issue. The user _Cbhihe_ pointed
out that my partition sizes weren't sustainable, so that I adjusted it
accordingly for this article.

# Conclusion

We set up Arch Linux with encrypted partitions using the variant LUKS on LVM.
The system is split up into multiple partitions, which are all encrypted using
the same password over one common encrypted volume.

The `/boot` partition remains unencrypted, which would allow an attacker to
tamper with my system boot, and, potentially, with my entire setup. So if my
laptop gets lost, there is no 100% guarantee that my system hasn't been
manipulated. However, having all the other partitions encrypted makes
manipulation or data leaks extremely unlikely.

When my laptop gets stolen or lost for an extended period of time, it is still a
good idea to revoke all the SSH and GPG keys. But thanks to encryption, that
procedure won't be an act of emergency, but one that could be carefully planned
for and executed at a convenient time.

The system set up so far is really just a base installation. In order to work on
this computer, further tasks have to be performed, as roughly described in my
[Arch
Setup Notes](https://github.com/patrickbucher/docs/blob/master/arch-setup/arch-setup.md)
on GitHub. Maybe I'll write a follow-up article in the future to cover those
steps more in detail.
