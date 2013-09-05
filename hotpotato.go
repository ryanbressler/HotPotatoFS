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

func ServeDir(mountpoint string, target string, limit int, me string, peerlist []string) {

	peers := groupcache.NewHTTPPool(me)

	if peerlist != nil && len(peerlist) > 0 {
		peers.Set(peerlist...)
	}

	filecache = groupcache.NewGroup("filecache", int64(limit)<<20, groupcache.GetterFunc(
		func(ctx groupcache.Context, key string, dest groupcache.Sink) error {
			contents, err := ioutil.ReadFile(key)
			dest.SetBytes(contents)
			return err
		}))

	c, err := fuse.Mount(mountpoint)
	if err != nil {
		log.Fatal(err)
	}

	fs.Serve(c, TargetDir{target})

}

type TargetDir struct {
	Path string
}

func (nf TargetDir) Root() (fs.Node, fuse.Error) {
	return Dir{Node{Path: nf.Path}}, nil
}

type Node struct {
	Path string
}

func (n Node) Attr() fuse.Attr {
	s, err := os.Stat(n.Path)
	if err != nil {
		log.Print(err)
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
		log.Print(err)
		return nil, fuse.ENOENT
	}
	node := Node{path}
	switch {
	case s.IsDir():
		fs = Dir{node}
	case s.Mode().IsRegular():
		fs = File{node}
	default:
		fs = node
	}

	return
}

func (d Dir) ReadDir(intr fs.Intr) ([]fuse.Dirent, fuse.Error) {
	var out []fuse.Dirent
	files, err := ioutil.ReadDir(d.Path)
	if err != nil {
		log.Print(err)
		return nil, fuse.Errno(err.(syscall.Errno))
	}
	for _, d := range files {
		de := fuse.Dirent{Name: d.Name()}
		if d.IsDir() {
			de.Type = fuse.DT_Dir
		}
		if d.Mode().IsRegular() {
			de.Type = fuse.DT_File
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
		log.Print(err)
		return nil, fuse.ENOENT
	}
	return contents, nil
}

// func (f File) Read(req *fuse.ReadRequest, resp *fuse.ReadResponse, intr fs.Intr) fuse.Error {
// 	//log.Print("Read Called: ", req, resp, intr)
// 	//fuse.Debugf = log.Printf
// 	var dst groupcache.ByteView
// 	err := filecache.Get(nil, f.Path, groupcache.ByteViewSink(&dst))
// 	if err != nil {
// 		log.Print(err)
// 		return fuse.ENOENT
// 	}

// 	resp = &fuse.ReadResponse{Data: dst.Slice(int(req.Offset), int(req.Offset)+req.Size).ByteSlice()}

// 	return nil
// }
