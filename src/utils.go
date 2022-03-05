package main

import (
	"context"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
	"github.com/gin-gonic/gin"
)

type middleware func(gin.HandlerFunc) gin.HandlerFunc

func chainMiddleWareWithDummy(mws ...middleware) gin.HandlerFunc {
	chain := func(ctx *gin.Context) {}
	for j := len(mws) - 1; j >= 0; j-- {
		chain = mws[j](chain)
	}
	return chain
}

type ctxCli struct {
	ctx context.Context
	cli *client.Client
}

func waitForDeath(cc *ctxCli, id string) error {
	statusCh, errCh := cc.cli.ContainerWait(cc.ctx, id, container.WaitConditionNotRunning)
	select {
	case <-statusCh:
	case err := <-errCh:
		if err != nil {
			return err
		}
	}
	return nil
}

func startContainer(cc *ctxCli, id string) error {
	err := cc.cli.ContainerStart(cc.ctx, id, types.ContainerStartOptions{})
	return err
}

func getLogs(cc *ctxCli, id string) (string, string, error) {
	streams, err := cc.cli.ContainerLogs(cc.ctx, id, types.ContainerLogsOptions{ShowStdout: true, ShowStderr: true})
	if err != nil {
		return "", "", err
	}
	defer streams.Close()
	stdOutput, stdErr := new(strings.Builder), new(strings.Builder)
	_, err = stdcopy.StdCopy(stdOutput, stdErr, streams)
	if err != nil {
		return "", "", err
	}
	return stdOutput.String(), stdErr.String(), nil
}

func rmContainer(cc *ctxCli, id string) error {
	err := cc.cli.ContainerRemove(cc.ctx, id, types.ContainerRemoveOptions{})
	return err
}
