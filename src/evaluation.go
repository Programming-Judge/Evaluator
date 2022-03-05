package main

import (
	"context"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"strconv"
)

func initEvaluationContainer(args evaluationArgs, cc *ctxCli) (string, error) {
	containerConfig := &container.Config{
		Image: langImageMap[args.lang],
		Cmd: []string{"bash", evaluationScript, args.ID, langExtensionMap[args.lang], args.problemID,
			strconv.FormatInt(int64(args.timeLimit), 10), evalMntDir, submissionsDir, testCasesDir},
		User:       unpUser,
		WorkingDir: "/home/" + unpUser,
	}
	hostConfig := &container.HostConfig{
		Binds:     []string{volume + ":" + evalMntPath},
		Resources: container.Resources{Memory: int64(args.memoryLimit * 1e6)},
	}
	cont, err := cc.cli.ContainerCreate(cc.ctx, containerConfig, hostConfig, nil, nil, "")
	return cont.ID, err
}

func evaluate(args evaluationArgs) (string, string, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	cc := ctxCli{
		ctx: context.Background(),
		cli: cli,
	}
	if err != nil {
		return "", "", err
	}
	cid, err := initEvaluationContainer(args, &cc)
	if err != nil {
		return "", "", err
	}
	err = startContainer(&cc, cid)
	if err != nil {
		return "", "", err
	}
	err = waitForDeath(&cc, cid)
	if err != nil {
		return "", "", err
	}
	stdout, stderr, err := getLogs(&cc, cid)
	if err != nil {
		return "", "", err
	}
	err = rmContainer(&cc, cid)
	if err != nil {
		return "", "", err
	}
	return stdout, stderr, nil
}
