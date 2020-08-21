---
title: Virtual Machines with libvirt and Networking
subtitle: Using Debian 10 («Buster») as Host and Guest
author: Patrick Bucher
date: 2020-08-01T22:30:00
lang: en
---

I'd like to dig deeper into system administration tasks. At work, I have to
manage a fleet of Linux servers with Puppet. And in my spare time, I'd like to
manage the servers I run with Ansible or Puppet in the future.

Virtual Machines are easily obtained nowadays. Cloud providers such as Digital
Ocean or Exoscale offer virtual machines with various operating systems at
rather moderate prices. You only have to pay for the time the virtual machines
are actually running, so you can save money by shutting those hosts down when
not needed.

However, running those virtual machines locally costs even less. No additional
public IPv4 addresses are wasted, and, most importantly, a local setup allows
you to test changes to be applied to your productive environment locally
beforehand.

This article shows how to set up three virtual machines ‒ `master`, `node1`, and
`node2`, which later could be used for a Puppet setup with a Puppetmaster ‒
using [libvirt](https://libvirt.org/) on top of
[KVM](https://www.linux-kvm.org/page/Main_Page). [Debian 10
(«Buster»)](https://www.debian.org/releases/buster/) is going to be used both as
the host and guest operating system. The host operating system is installed on a
Dell Latitude E6430 from 2013 with 8 GB or RAM, which is just laying around
here. (This also proofs that you don't need a whole lot of hardware resources
for such a setup.)

# Setting up the Virtualization

Given a fresh Debian setup with the lightweight LXQt desktop, a couple of
packages need to be installed in order to get virtualization to work:

    # apt-get install \
        qemu-kvm \
        libvirt-clients \
        libvirt-daemon-system \
        virtinst \
        bridge-utils

Make sure to activate virtualization in the BIOS. Check if the `kvm` kernel
module is activated:

    $ lsmod | grep ^kvm
    kvm                 835584  1 kvm_intel

If there is a number not equal to 0 in the third column, `kvm` is up and
running.

# Setting up the Virtual Network

Usually a `default` network is pre-defined, which can be checked as follows:

    # virsh net-list --all
     Name      State      Autostart   Persistent
    ----------------------------------------------
     default   inactive   no          yes

The `default` network can be configured to be started up automatically:

    # virsh net-autostart default
    Network default marked as autostarted

Until the next system restart, it is started up manually:

    # virsh net-start default
    Network default started

A bridge interface `virbr0` should have been created:

    # brctl show
    bridge name     bridge id               STP enabled     interfaces
    virbr0          8000.5254005f4e6b       yes             virbr0-nic

Make sure that NAT is activated:

    # sudo sysctl -a | grep 'net.ipv4.ip_forward ='
    net.ipv4.ip_forward = 1

The value of the above property must be `1`.

## Possible Issues

If `iptables` is in use, make sure to forward the traffic from the guests over
the bridge `virbr0`, so that the guests also have internet access:

    # iptables -I FORWARD -i virbr0 -o virbr0 -j ACCEPT

# Setting up the Virtual Machines

Since networking over the bridge interface requires `root` privileges, all
virtual machine files are put into the `/opt/vms` directory, which first needs
to be created:

    # mkdir /opt/vms
    # cd /opt/vms

The network installer for Debian Buster can be downloaded from the official
website:

    # wget https://cdimage.debian.org/debian-cd/current/amd64/iso-cd/\
    debian-10.4.0-amd64-netinst.iso

The `master` virtual machine is now setup using `virt-install`:

    # virt-install \
        --name master \
        --memory 1024 \
        --vcpus=1,maxvcpus=2 \
        --cpu host \
        --cdrom debian-10.4.0-amd64-netinst.iso \
        --disk /opt/vms/master.qcow2,size=8,format=qcow2 \
        --network network=default \
        --virt-type kvm

The machine gets 1 GB of memory and a 8 GB disk. Most importantly, the network
is set to the `default` network.

A window showing the Debian installer appears. Just install the standard system
utilities and the SSH server. The following users and passwords shall be
configured:

- `root`: `topsecret`
- `user`: `secret`

After the setup is finished, just let the system boot, and login as `root`. Then
shut the virtual machine down:

    # shutdown -h now

The two additional guest nodes can be created by cloning the `master` virtual
machine just set up:

    # virt-clone --original master --name node1 --file node1.qcow2
    # virt-clone --original master --name node2 --file node2.qcow2

Now start up all the nodes:

    # virsh --connect qemu:///session start master
    # virsh --connect qemu:///session start node1
    # virsh --connect qemu:///session start node2

# Configuring the Virtual Network

In order to conveniently access the guests, static IPs should be assigned to
them. The network configuration can be edited as follows:

    # virsh net-edit default

An editor showing an XML configuration appears:

    <network>
      <name>default</name>
      <uuid>fecb90d5-9b46-48f6-8b93-e57032f8ba6a</uuid>
      <forward mode='nat'/>
      <bridge name='virbr0' stp='on' delay='0'/>
      <mac address='52:54:00:63:d3:70'/>
      <ip address='192.168.122.1' netmask='255.255.255.0'>
        <dhcp>
          <range start='192.168.122.2' end='192.168.122.254'/>
        </dhcp>
      </ip>
    </network>

The `dhcp` section needs to be extended with static IP definitions, which map
the MAC addresses of the guest's virtual network interfaces to the static IP
addresses to be used.

The MAC addresses of the virtual machines can be extracted from their
configuration as follows:

    # virsh dumpxml master | grep -i '<mac'
        <mac address='52:54:00:db:07:7c'/>
    # virsh dumpxml node1 | grep -i '<mac'
        <mac address='52:54:00:a4:77:a9'/>
    # virsh dumpxml node2 | grep -i '<mac'
        <mac address='52:54:00:51:e8:ef'/>

Using those MAC addresses, new static host definitions can be created as
follows:

    <host mac='52:54:00:db:07:7c' name='master' ip='192.168.122.2'/>
    <host mac='52:54:00:a4:77:a9' name='node1' ip='192.168.122.3'/>
    <host mac='52:54:00:51:e8:ef' name='node2' ip='192.168.122.4'/>

The XML network definition should now look as follows (the `uuid` and `mac
address` of the host will vary):

    <network>
      <name>default</name>
      <uuid>fecb90d5-9b46-48f6-8b93-e57032f8ba6a</uuid>
      <forward mode='nat'/>
      <bridge name='virbr0' stp='on' delay='0'/>
      <mac address='52:54:00:63:d3:70'/>
      <ip address='192.168.122.1' netmask='255.255.255.0'>
        <dhcp>
          <range start='192.168.122.2' end='192.168.122.254'/>
          <host mac='52:54:00:db:07:7c' name='master' ip='192.168.122.2'/>
          <host mac='52:54:00:a4:77:a9' name='node1' ip='192.168.122.3'/>
          <host mac='52:54:00:51:e8:ef' name='node2' ip='192.168.122.4'/>
        </dhcp>
      </ip>
    </network>

After saving the configuration, the network `default` needs to be restarted:

    # virsh net-destroy default
    # virsh net-start default

The guest virtual machines must also be restarted so that they will get the new
IP addresses assigned:

    # virsh shutdown master
    # virsh shutdown node1
    # virsh shutdown node2

    # virsh --connect qemu:///session start master
    # virsh --connect qemu:///session start node1
    # virsh --connect qemu:///session start node2

The virtual machines should now be accessible through SSH:

    $ ssh user@192.168.122.2
    $ ssh user@192.168.122.3
    $ ssh user@192.168.122.4

Make sure that the network communication is working between the guests:

    [user@master]$ ping node1
    [user@master]$ ping node2

Also make sure to define the proper hostname in `/etc/hostname`, for it is still
`master` for the two guests that have been cloned from the initial image:

    [root@node1]# echo 'node1' > /etc/hostname
    [root@node2]# echo 'node2' > /etc/hostname

## Adding Some Comfort

Consider adding the following definitions to `/etc/hosts`:

    192.168.122.2   master
    192.168.122.3   node1
    192.168.122.4   node2

So that you can access your virtual machines by their host names:

    $ ssh user@master
    $ ssh user@node1
    $ ssh user@node2

In order to login to the guests without typing a password, create an SSH key
locally without any passphrase:

    $ ssh-keygen -t rsa -b 4096 -f ~/.ssh/id_vms_rsa

Make sure that your `~/.ssh` folder has the access mode `700`, and the contained
files all have the access mode `600` (thanks to [meillo](http://marmaro.de/) for
pointing that out):

    $ chmod 700 ~/.ssh
    $ chmod 600 ~/.ssh/*

Copy the public key to the hosts using `ssh-copy_id` (thanks to meillo again for
hinting that utility to me):

    $ ssh-copy-id -i ~/.ssh/id_vms_rsa user@master
    $ ssh-copy-id -i ~/.ssh/id_vms_rsa user@node1
    $ ssh-copy-id -i ~/.ssh/id_vms_rsa user@node2


Check that the SSH connection now works without any password:

    $ ssh -i ~/.ssh/id_vms_rsa user@master
    $ ssh -i ~/.ssh/id_vms_rsa user@node1
    $ ssh -i ~/.ssh/id_vms_rsa user@node2


# Conclusion

Three virtual machines running Debian GNU/Linux have been installed on a
rather old laptop running Debian GNU/Linux itself. Those virtual machines can be
comfortably accessed without any passwords through SSH, and are able to
communicate with one another over a virtual network.

It took me almost a day ‒ and gave me some additional grey hair ‒ to get all
this information together from various sources. After I figured out how to
create the setup described above, it only took me about two hours to reproduce 
everything on another laptop (including the setup of the laptop itself) and to
write this article.

Since I did the try-and-error part on Arch Linux, this article can also be used
on that distribution, and probably many others as well. Only the packages to be
installed will probably vary on other distributions.

I plan to describe the setup of a local Puppet environment based on the setup
described above in a forthcoming article.
