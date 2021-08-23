package docker

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/jsonmessage"
	"github.com/docker/docker/pkg/term"
	"os"
)

func PushImage(ctx context.Context, cli *client.Client, dockerAddress string, target string, username string, password string) {
	auth := types.AuthConfig{}
	if dockerAddress != "" {
		auth = types.AuthConfig{
			Username:      username,
			Password:      password,
			ServerAddress: dockerAddress,
		}
	} else {
		auth = types.AuthConfig{
			Username: username,
			Password: password,
		}
	}

	authBytes, _ := json.Marshal(auth)
	authBase64 := base64.URLEncoding.EncodeToString(authBytes)
	push, err := cli.ImagePush(ctx, target, types.ImagePushOptions{All: true,
		RegistryAuth: authBase64})
	if err != nil {
		panic(err)
	}
	termFd, isTerm := term.GetFdInfo(os.Stderr)
	err = jsonmessage.DisplayJSONMessagesStream(push, os.Stderr, termFd, isTerm, nil)
	if err != nil {
		panic(err)
	}
}
