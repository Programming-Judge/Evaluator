package main

import (
	"context"
	"fmt"
	"os"

	"io"
	"path/filepath"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/client"
)

// Refer to https://docs.docker.com/engine/api/sdk/examples/
func execute(id, lang, timelimit string) (string, error) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {

		// Error in creation of docker client
		return "", err
	}

	image_name := lang_image_map[lang]

	// Image Pull
	// reader, err := cli.ImagePull(ctx, "docker.io/library/alpine", types.ImagePullOptions{})
	// if err != nil {
	//     panic(err)
	// }
	// io.Copy(os.Stdout, reader)

	// Path to the directory where the
	// files are mounted (in the host)
	location, _ := os.Getwd()

	// Modified to work on windows as well
	location = filepath.Join(filepath.Dir(location), "interface", bind_mnt_dir)

	// Container creation
	resp, err := cli.ContainerCreate(
		ctx,
		&container.Config{
			Image: image_name,
			Cmd:   []string{id, lang_extension_map[lang], bind_mnt_dir, timelimit},
		},
		&container.HostConfig{
			Mounts: []mount.Mount{
				{
					Type:   mount.TypeBind,
					Source: location,
					Target: "/home/" + unp_user + "/" + bind_mnt_dir,
				},
			},
		},
		nil,
		nil,
		"")
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	// Container start
	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {

		// Some error has occurred while starting the container
		fmt.Println(err)
		return "", err
	}

	// Wait till the container stops
	statusCh, errCh := cli.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			fmt.Println(err)
			// Some error has occurred while the
			// container was executing the code
			panic(err)
		}
	case <-statusCh:
	}

	// Read output
	out, err := cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true})
	if err != nil {

		// Some error has occurred while reading
		// the output from the container
		return "", err
	}
	defer out.Close()

	buf := new(strings.Builder)
	_, err = io.Copy(buf, out)

	if err != nil {
		return "", err
	}

	return buf.String(), err
}
