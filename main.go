package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/calavera/docker-volume-api"
)

var (
	root = flag.String("root", volumeapi.DefaultDockerRootDirectory, "Docker volumes root directory")
)

type garbageDriver struct {
	root string
}

func (g garbageDriver) Create(r volumeapi.VolumeRequest) volumeapi.VolumeResponse {
	fmt.Printf("Create %v\n", r)
	return volumeapi.VolumeResponse{}
}

func (g garbageDriver) Remove(r volumeapi.VolumeRequest) volumeapi.VolumeResponse {
	fmt.Printf("Remove %v\n", r)
	return volumeapi.VolumeResponse{}
}

func (g garbageDriver) Path(r volumeapi.VolumeRequest) volumeapi.VolumeResponse {
	fmt.Printf("Path %v\n", r)
	return volumeapi.VolumeResponse{Mountpoint: filepath.Join(g.root, r.Name)}
}

func (g garbageDriver) Mount(r volumeapi.VolumeRequest) volumeapi.VolumeResponse {
	p := filepath.Join(g.root, r.Name)
	fmt.Printf("Mount %s\n", p)

	if err := os.MkdirAll(p, 0755); err != nil {
		return volumeapi.VolumeResponse{Err: err}
	}

	if err := ioutil.WriteFile(filepath.Join(p, "test"), []byte("TESTTEST"), 0644); err != nil {
		return volumeapi.VolumeResponse{Err: err}
	}

	return volumeapi.VolumeResponse{Mountpoint: p}
}

func (g garbageDriver) Unmount(r volumeapi.VolumeRequest) volumeapi.VolumeResponse {
	p := filepath.Join(g.root, r.Name)
	fmt.Printf("Unmount %s\n", p)

	err := os.RemoveAll(p)
	return volumeapi.VolumeResponse{Err: err}
}

func main() {
	d := garbageDriver{*root}
	h := volumeapi.NewVolumeHandler(d)
	fmt.Println("Listening on :7878")
	fmt.Println(h.ListenAndServe("tcp", ":7878", ""))
}
