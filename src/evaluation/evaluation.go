package evaluation

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
)

// Convenience struct
// Holds context and a docker client
type ctxCli struct {
	ctx context.Context
	cli *client.Client
}

// Initialise the container
func initContainer(args evaluationArgs, cc *ctxCli) (string, error) {
	container_config := &container.Config{
		Image: lang_image_map[args.lang],
		Cmd: []string{"bash", script, args.id, lang_extension_map[args.lang], args.problemId,
			strconv.FormatInt(int64(args.timeLimit), 10), mnt_dir, submissions_dir, testcases_dir},
		User: unp_user,
		WorkingDir: "/home/" + unp_user,
	}
	host_config := &container.HostConfig{
		Binds:     []string{"judge-submissions" + ":" + container_mnt_path},
		Resources: container.Resources{Memory: int64(args.memoryLimit * 1e6)},
	}
	cont, err := cc.cli.ContainerCreate(cc.ctx, container_config, host_config, nil, nil, "")
	fmt.Println(cont.ID)
	return cont.ID, err
}

// Start evaluation
func startContainer(cc *ctxCli, id string) error {
	err := cc.cli.ContainerStart(cc.ctx, id, types.ContainerStartOptions{})
	return err
}

// Wait till the container finishes
// evaluation
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

// Get stdout and stderr
func getLogs(cc *ctxCli, id string) (string, string, error) {
	streams, err := cc.cli.ContainerLogs(cc.ctx, id, types.ContainerLogsOptions{ShowStdout: true})
	if err != nil {
		return "", "", err
	}
	defer streams.Close()
	std_output, std_err := new(strings.Builder), new(strings.Builder)
	_, err = stdcopy.StdCopy(std_output, std_err, streams)
	if err != nil {
		return "", "", err
	}
	fmt.Println(std_output.String(), std_err.String())
	return std_output.String(), std_err.String(), nil
}

// Remove the container after use
func rmContainer(cc *ctxCli, id string) error {
	err := cc.cli.ContainerRemove(cc.ctx, id, types.ContainerRemoveOptions{})
	return err
}

// Master evaluation function
// An error at any stage is returned
// Otherwise, the verdict and another 
// parameter (stderr from container) 
// are returned
func evaluate(args evaluationArgs) (string, string, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	cc := ctxCli{
		ctx: context.Background(),
		cli: cli,
	}
	if err != nil {
		return "", "", err
	}
	cid, err := initContainer(args, &cc)
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
	verdict, verdictError, err := getLogs(&cc, cid)
	if err != nil {
		return "", "", err
	}
	err = rmContainer(&cc, cid)
	if err != nil {
		return "", "", err
	}
	return verdict, verdictError, nil
}
