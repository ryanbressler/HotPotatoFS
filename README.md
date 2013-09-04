![Potato](potatosmall.png "potato image copyleft David Gibbs")     HotPotatoFS    ![Potato](potatosmall.png "potato image copyleft David Gibbs")
===========


HotPotato = Playing catch in a group with a FUSE/timelimit
Group Cache + FUSE = Read Only Distribute in Memory Filesystem

HotPotato is a simple read only FUSE filesystem used to put groupcache in front of a slow 
network mount (nfs, s3 etc) to reduce data acces time.

It is expermental software in an early state of development and may break.

It is written in golang and combines groupcache and the bazil.org fuse library: 
https://github.com/golang/groupcache
http://bazil.org/fuse/


QuickStart 
-----------

With go and go path set up:

```bash
go get github.com/ryanbressler/HotPotatoFS
go install github.com/ryanbressler/HotPotatoFS/hotpot

# single machine
hotpot -mountpoint /hotpotato -target /nfsmount

# config file for multple peers coming soon.

```

Administrive
------------

HotPotatofFS was developed by members of the Shumelevich lab at the Institute for Systems Biology to support distributed
computing in cancer and biomedical research as part of our work on The Cancer Genome Atlas and other projecs.

Code is under a 3 clause BSD. 

HotPotato image copyleft David Gibbs.


