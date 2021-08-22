package docker

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/jsonmessage"
	"github.com/docker/docker/pkg/term"
	"os"
)

func PullImage(ctx context.Context, client *client.Client, source string)  {
	reader, err := client.ImagePull(ctx, source, types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}
	defer reader.Close()

	termFd, isTerm := term.GetFdInfo(os.Stderr)
	err = jsonmessage.DisplayJSONMessagesStream(reader, os.Stderr, termFd, isTerm, nil)
	if err != nil {
		panic(err)
	}
}
