![Potato](potatosmall.png "potato image copyleft David Gibbs")     HotPotatoFS 
===========

```
Groupcache + FUSE 
= Read-Only, Distribute, in-Memory-Caching Filesystem in Go
= Playing catch in a group with a FUSE
= HotPotato
```

HotPotato is a simple, read-only, in-memory-Cacheing, FUSE filesystem used to put groupcache in front 
of a slow disk or network mount (nfs, s3, smb etc) on one or many machines to reduce the time
needed to repeatedly read a file.

It was developed primairly for use in parralel data analysis and simulation. You can use it
to speed up a parralel R, python, matlab or whatever job running via golem, qsub, gnu parralel etc
without needing to rewrite the analysis code.

It is expermental software in an early state of development and may break.

It is written in go (golang) using groupcache and the bazil.org fuse library:

https://github.com/golang/groupcache

http://bazil.org/fuse/


QuickStart 
-----------

With go and go path set up:

Instalation

```bash
go get github.com/ryanbressler/HotPotatoFS
go install github.com/ryanbressler/HotPotatoFS/hotpot
```

Single Machine

```bash
hotpot -mountpoint /hotpotato -target /nfsmount
```

Multi Machine

peerfile.txt:
```
http://host1:8080
http://host2:8080
```

host1:
```bash
hotpot -mountpoint /hotpotato -target /nfsmount -me http://host1:8080 -peers peerfile.txt
```

host2:
```bash
hotpot -mountpoint /hotpotato -target /nfsmount -me http://host2:8080 -peers peerfile.txt
```


Credit
------------

HotPotatofFS was developed by members of the Shumelevich lab at the Institute for Systems Biology to support distributed
computing in cancer and biomedical research as part of our work on The Cancer Genome Atlas and other projecs.

Code is under a 3 clause BSD. 

Potato image copyleft David Gibbs.


