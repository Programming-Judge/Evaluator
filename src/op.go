package main

import (
	"context"
	"os"
	"path/filepath"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

func initOpContainer(cc *ctxCli, p string, w int) (string, error) {
	cmd := []string{"true"}
	switch w {
	case 2:
		cmd = []string{"bash", "work.sh", "2", opMntPath, submissionsDir, p}
	case 4:
		cmd = []string{"bash", "work.sh", "4", opMntPath, testCasesDir, p}
	}
	containerConfig := &container.Config{
		Image:      opImage,
		WorkingDir: "/tmp",
		Cmd:        cmd,
	}
	hostConfig := &container.HostConfig{
		Binds: []string{volume + ":" + opMntPath},
	}
	cont, err := cc.cli.ContainerCreate(cc.ctx, containerConfig, hostConfig, nil, nil, "")
	return cont.ID, err
}

func copyToContainer(fileName string, w int, id string, cc *ctxCli) error {
	currDir, err := os.Getwd()
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}
	wd := submissionsDir
	if w == 3 {
		wd = testCasesDir
	}
	super := filepath.Dir(currDir)
	path := filepath.Join(super, iface, wd, fileName)
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	err = cc.cli.CopyToContainer(cc.ctx, id, opMntPath+"/"+wd, file, types.CopyToContainerOptions{})
	return err
}

func op(p string, w int) error {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	cc := ctxCli{
		ctx: context.Background(),
		cli: cli,
	}
	if err != nil {
		return err
	}
	id, err := initOpContainer(&cc, p, w)
	if err != nil {
		return err
	}
	switch w {
	case 1:
		err = copyToContainer(p, w, id, &cc)
		if err != nil {
			return err
		}
	case 3:
		err = copyToContainer(p, w, id, &cc)
		if err != nil {
			return err
		}
	case 2:
		err = startContainer(&cc, id)
		if err != nil {
			return err
		}
		err = waitForDeath(&cc, id)
		if err != nil {
			return err
		}
	case 4:
		err = startContainer(&cc, id)
		if err != nil {
			return err
		}
		err = waitForDeath(&cc, id)
		if err != nil {
			return err
		}
	}
	return rmContainer(&cc, id)
}
