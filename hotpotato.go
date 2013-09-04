package HotPotatoFS

import (
	"bazil.org/fuse"
	"bazil.org/fuse/fs"
	"io/ioutil"
	//"github.com/golang/groupcache"
	"log"
	"os"
	"syscall"
)

func ServeNfs(mountpoint string, nfsdir string) {
	c, err := fuse.Mount(mountpoint)
	if err != nil {
		log.Fatal(err)
	}

	fs.Serve(c, NfsDir{nfsdir})

}

// FS implements the hello world file system.
type NfsDir struct {
	Path string
}

func (nf NfsDir) Root() (fs.Node, fuse.Error) {
	return Dir{Node{Path: nf.Path}}, nil
}

// Dir implements both Node and Handle for the root directory.
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

// func (Dir) Lookup(name string, intr fs.Intr) (fs.Node, fuse.Error) {
// 	if name == "hello" {
// 		return File{}, nil
// 	}
// 	return nil, fuse.ENOENT
// }

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

// File implements Handle
type File struct {
	Node
}

func (File) ReadAll(intr fs.Intr) ([]byte, fuse.Error) {
	return []byte("hello, world\n"), nil
}
