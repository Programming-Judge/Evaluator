package main

import (
	"context"
	"io"
	"os"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/client"
)

//Refer to https://docs.docker.com/engine/api/sdk/examples/

func execute(code_path , input_path , output_path, time_limit string) (string , error) {

	ctx := context.Background()
    cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
    if err != nil {
        return "" , err
    }

	//Container image name
	image_name := "cpp/test"

	//Image Pull
    // reader, err := cli.ImagePull(ctx, "docker.io/library/alpine", types.ImagePullOptions{})
    // if err != nil {
    //     panic(err)
    // }
    // io.Copy(os.Stdout, reader)

	//Folder Path to mount
	 
	//TODO Find Folder Path by parsing
	location , _ := os.Getwd() 
	location += "/submissions"

	//Container creation
    resp, err := cli.ContainerCreate(
		ctx,
		&container.Config{
        Image: image_name,
		Cmd: []string{code_path , input_path , output_path, time_limit},
    	},
		&container.HostConfig{
			Mounts: []mount.Mount{
				{
					Type:   mount.TypeBind,
					Source: location,
					Target: "/submissions",
				},
			},
		},
		nil,
		nil,
		"")
    if err != nil {
        return  "" , err
    }

	//Container start
    if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
        return "" , err
    }

	//Wait till container stops
    statusCh, errCh := cli.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
    select {
    case err := <-errCh:
        if err != nil {
            panic(err)
        }
    case <-statusCh:
    }

	//Read output 
    out, err := cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true})
	if err != nil {
        return  "" , err
    }
	defer out.Close()

	// Convert io.ReaderCloser to string
	buf := new(strings.Builder)
	_, err = io.Copy(buf, out)

	if err != nil {
		return  "" , err
	}

	s := buf.String()
	

	return  s , err
}