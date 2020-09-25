---
title: Semi-Automated Arch Linux Setup with Disk Encryption
subtitle: Using dm-crypt, LUKS, and some Shell Scripts
author: Patrick Bucher
date: 2020-09-25T14:00:00
lang: en
---

I've been using Arch Linux since 2016. I got to understand the system better
since then, and installing Arch nowadays is rather a strain on my fingers than
on my brain. I automated my personal setup procedure to some extend with a
couple of shell scripts I'm adjusting as time goes and hardware changes.

Even though I've often taken my laptop computer with me, especially for
university, I never encrypted my disk or a single partitoin thereon. Sensitive
data, such as passwords, are stored encryted using GPG. And my GPG private key
is protected with a strong passphrase; so strong, that I mistype it once in a
while.

Having stored SSH keys on my laptop, however, is a risk in terms of data
protection. I do not protect most of my SSH keys with passwords. (_Not_ having
to use passwords is one of the conveniences SSH gives you, after all.) But some
those SSH keys are gateways to quite some important data, and therefore I need
to re-consider my security practices.

Protecting SSH keys with a passphrase is very inconvenient, especially if you do
a lot of `git push` operations via SSH. It's possible to use a password
protected SSH key for `git` only, but having to deal with multiple SSH keys is
inconvenient, too. (My rule is to use one SSH key per client machine.) So
encrypting the hard disk or at least some partitions of it might be the better
solution, because it only requires me to enter an additional password once a
day.

# dm-crypt and LUKS

The Arch wiki provides an article on [how to encrypt an entire
system](https://wiki.archlinux.org/index.php/Dm-crypt/Encrypting_an_entire_system).
Such a setup is based on LUKS (Linux Unified Key Setup), a disk encryption
specification, and dm-crypt, the Linux kernel's device mapper for encrypted
devices.

The wiki mentions different
[scenarios](https://wiki.archlinux.org/index.php/Dm-crypt/Encrypting_an_entire_system#Overview)
for disk encryption. The first option, [LUKS on a
partition](https://wiki.archlinux.org/index.php/Dm-crypt/Encrypting_an_entire_system#LUKS_on_a_partition)
is very simple, but not ideal for me, because I'm not just using one giant root
partition, but separate partitions for `/var`, `/tmp` and `/home`. With this
setup, a process filling up my `/var/log` directory won't prevent me from
writing to `/home` or the root partition, which might be necessary to deal with
the issue.

One could argue that it's sufficient to just encrypt the `/home` partition,
which contains the most sensitive data, after all. However, it's possible that
some program logs sensitive data to `/var/log` or `/opt`, so that an encrypted
`/home` partition won't prevent sensitive data to be leaked. An unencrypted root
partition also makes it possible to tamper programs, so that data is leaked
later.

It's also possible to encrypt the partitions mentioned seperately, which would
require me to enter my password upon booting for every single partition.

The second option, [LVM on
LUKS](https://wiki.archlinux.org/index.php/Dm-crypt/Encrypting_an_entire_system#LVM_on_LUKS),
is a better fit for my purpose. Here, LUKS provides one single block device, on
top of which the encrypted partitions are created.

The other, arguably more sophisticated scenarios, are mostly useful when dealing
with multiple disks, or if an encrypted boot partition is required. So let's
stick to scenario two: LVM on LUKS.

## Partitioning

My Lenovo Thinkpad X1 has a 500 GB SSD and 16 GB memory. This is important when
considering partition sizes. I usually create my partitions as follows:

1. A 256 MB `/boot` partition for the bootloader (FAT-32).
2. A 16 GB (the size of my memory) swap partition.
3. A 64 GB `/` (root) partition (ext4).
4. A 8 GB `/var` partition (ext4).
5. A 4 GB `/tmp` partition (ext4).
6. And a `/home` partition with the remainder of the space, i.e. roughly 400 GB
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
continue with the partitioning.

## Partitioning

My Thinkpad X1 has a US keyboard, which I [prefer to
use](https://paedubucher.ch/articles/2020-09-16-compose-key-on-x.html) nowadays.
However, it's possible to change the keyboard layout as follows (for Swiss
German), if wanted:

    # loadkeys de_CH-latin1

Next, I establish a WiFi connection with `frzbxpdb5` being my network SSID, 
`wlan0` being my WiFi device, and `[topsecret]` being my password:

    # iwctl --passphrase `[topsecret]` station wlan0 connect frzbxpdb5

In order to partition the hard disk, one first needs to know the device name,
which can be found using `lsblk`:

    # lsblk

This lists `sda`, the USB dongle, and `nvme0n1`, the solid state drive
aforementioned. For the actual partitioning, I have created a
[script](https://github.com/patrickbucher/docs/blob/master/arch-setup/partition-encrypted.sh),
which I download:

    # curl https://github.com/patrickbucher/docs/blob/master/arch-setup/\
        partition-encrypted.sh > partition-encrypted.sh

Let's go through that script. First, some definitions:

    #!/usr/bin/bash

    set -e

    # define a disk and set partition sizes
    disk='nvme0n1'
    init_mibs='1'
    boot_mibs='256'
    swap_mibs='16000'
    root_mibs='64000'
    var_mibs='8000'
    tmp_mibs='4000'

With `set -e` activated, the script is exited as soon as an error happens. Next,
I define the partition sizes in megabytes. Notice that I add a gap of one
megabyte at the very beginning, so that the sector allocation will be correct,
no matter what block size my SSD is using.

Next, the existing partitions are deleted, and the disk is overwritten with
random junk data:

    # delete all partitions and overwrite disk with random data
    disk_dev="/dev/${disk}"
    shred --random-source=/dev/urandom --iterations=1 "/dev/${disk}"
    parted -s "$disk_dev" mklabel gpt

I haven't used this Thinkpad X1 productively, so there's no sensitive data to be
overwritten. However, filling the disk with random data is a good idea, so that
an attacker cannot distinguish between encrypted data and mere junk. (Using more
than a single iteration is reasonable if there is existing data to be
overwritten.) The GPT partition scheme is then created using `parted`.
