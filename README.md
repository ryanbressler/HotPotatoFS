![Potato](potatosmall.png "potato image copyleft David Gibbs")     HotPotatoFS 
===========

```
HotPotato = Playing catch in a group with a FUSE/timelimit
Group Cache + FUSE = Read-Only, Distribute in Memory Filesystem
```

HotPotato is a simple read-only, in memory, FUSE filesystem used to put groupcache in front 
of a slow disk or network mount (nfs, s3, smb etc) on one or many machines to reduce the time
needed to repeatedly read a file.

It was developed primairly for use in parralel data analysis and simulation. You can use it
to speed up a parralel matlab, R, python etc job running via golem, qsub, gnu parralel etc
without needing to rewrite the analysis code.

It is expermental software in an early state of development and may break.

It is written in go (golang) using groupcache and the bazil.org fuse library: 
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

Potato image copyleft David Gibbs.


