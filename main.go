package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/calavera/docker-volume-api"
)

type garbageDriver struct{}

func (g garbageDriver) Create(r volumeapi.VolumeRequest) volumeapi.VolumeResponse {
	return volumeapi.VolumeResponse{}
}

func (g garbageDriver) Remove(r volumeapi.VolumeRequest) volumeapi.VolumeResponse {
	return volumeapi.VolumeResponse{}
}

func (g garbageDriver) Path(r volumeapi.VolumeRequest) volumeapi.VolumeResponse {
	return volumeapi.VolumeResponse{Mountpoint: filepath.Join(r.Root, r.Name)}
}

func (g garbageDriver) Mount(r volumeapi.VolumeRequest) volumeapi.VolumeResponse {
	p := filepath.Join(r.Root, r.Name)

	if err := os.MkdirAll(p, 0755); err != nil {
		return volumeapi.VolumeResponse{Err: err}
	}

	if err := ioutil.WriteFile(filepath.Join(p, "test"), []byte("TESTTEST"), 0644); err != nil {
		return volumeapi.VolumeResponse{Err: err}
	}

	return volumeapi.VolumeResponse{Mountpoint: p}
}

func (g garbageDriver) Umount(r volumeapi.VolumeRequest) volumeapi.VolumeResponse {
	p := filepath.Join(r.Root, r.Name)

	err := os.RemoveAll(p)
	return volumeapi.VolumeResponse{Err: err}
}

func main() {
	d := garbageDriver{}
	h := volumeapi.NewVolumeHandler(d)
	fmt.Println("Listening on :7878")
	fmt.Println(h.ListenAndServe("tcp", ":7878", ""))
}
