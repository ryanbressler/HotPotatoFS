package main

import (
	"flag"
	"github.com/ryanbressler/HotPotato"
)

func main() {
	var mountpoint string
	flag.StringVar(&mountpoint, "mountpoint", "hotpotato", "Where to mount the fuse dir.")

	var target string
	flag.StringVar(&target, "target", "", "The target dir")

	flag.Parse()

	HotPotato.ServeNfs(mountpoint, target)

}
