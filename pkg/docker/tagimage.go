package docker

import (
	"context"
	"github.com/docker/docker/client"
)

func TagImage(ctx context.Context, cli *client.Client, source string, target string)  {
	err := cli.ImageTag(ctx, source, target)
	if err != nil {
		panic(err)
	}
}
