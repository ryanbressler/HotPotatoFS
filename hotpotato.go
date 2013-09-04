package HotPotatoFS

import (
	"bazil.org/fuse"
	"bazil.org/fuse/fs"
	"github.com/golang/groupcache"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"syscall"
)

var filecache *groupcache.Group

func ServeNfs(mountpoint string, nfsdir string, peerlist []string) {
	// me := "http://localhost"
	// peers := groupcache.NewHTTPPool(me)

	// Whenever peers change:
	//peers.Set("http://10.0.0.1", "http://10.0.0.2", "http://10.0.0.3")

	filecache = groupcache.NewGroup("filecache", 64<<20, groupcache.GetterFunc(
		func(ctx groupcache.Context, key string, dest groupcache.Sink) error {
			contents, err := ioutil.ReadFile(key)
			dest.SetBytes(contents)
			return err
		}))

	c, err := fuse.Mount(mountpoint)
	if err != nil {
		log.Fatal(err)
	}

	fs.Serve(c, NfsDir{nfsdir})

}

type NfsDir struct {
	Path string
}

func (nf NfsDir) Root() (fs.Node, fuse.Error) {
	return Dir{Node{Path: nf.Path}}, nil
}

type Node struct {
	Path string
}

func (n Node) Attr() fuse.Attr {
	s, err := os.Stat(n.Path)
	if err != nil {
		return fuse.Attr{}
	}

	return fuse.Attr{Size: uint64(s.Size()), Mtime: s.ModTime(), Mode: s.Mode()}
}

type Dir struct {
	Node
}

func (d Dir) Lookup(name string, intr fs.Intr) (fs fs.Node, error fuse.Error) {

	path := filepath.Join(d.Path, name)
	s, err := os.Stat(path)
	if err != nil {
		return nil, fuse.ENOENT
	}
	node := Node{path}
	if s.IsDir() {
		fs = Dir{node}
	} else {
		fs = File{node}
	}

	return
}

func (d Dir) ReadDir(intr fs.Intr) ([]fuse.Dirent, fuse.Error) {
	var out []fuse.Dirent
	files, err := ioutil.ReadDir(d.Path)
	if err != nil {
		return nil, fuse.Errno(err.(syscall.Errno))
	}
	for _, d := range files {
		de := fuse.Dirent{Name: d.Name()}
		if d.IsDir() {
			de.Type = fuse.DT_Dir
		}
		out = append(out, de)
	}

	return out, nil
}

type File struct {
	Node
}

func (f File) ReadAll(intr fs.Intr) ([]byte, fuse.Error) {
	var contents []byte
	err := filecache.Get(nil, f.Path, groupcache.AllocatingByteSliceSink(&contents))
	if err != nil {
		return nil, fuse.ENOENT
	}
	return contents, nil
}
