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

func PushImage(ctx context.Context, cli *client.Client, dockerAddress string, target string, username string, password string)  {
	auth := types.AuthConfig{}
	if dockerAddress != "" {
		auth = types.AuthConfig{
			//Username: "AWS",
			Username: username,
			//Password: "eyJwYXlsb2FkIjoiU0Z4ZlpGcG5FT3FOdFA1eWJJem9jL2tXSHhwL2ZBa0FpTTNMZjJqRVBjT0ROcXdQTWdSWGNFOSt6aTN5NUJxQWZOZEVwZi9qL2pjMXl0T3BrcWdRc290NTNXTStMZG9xdWhiVmpjdEcvWlB4SG1reFliVllTNTE1VlBPNTFCS2pYTWMwT3p6M0VHMU5XTjFsMU9xSWFhdVFsbDVrQmZQUHJZV1Frd2ZkQmxvd2c5NWFPK3VqaWcwYXRzSWZreUtMeTZ5b0QrT3NkaWszbjB6QWNpNnVscjBRZDZKVHl6NGFTSkZpVDlSMXJVWS9TMU1oWEZ3cWhEeHhSNTU2YWxtN2EvL3I1enhJUHRNVE1LWWVtWWxjejIzckQvL0JUWWtVT21aWjU2SE4zK0dOR0RpcGpab1l3RXlndjFGVGFKTzdUMUhWS1owUS9KWGhsdEc3b1E1UlY0SFk4dzQyNldXTEh6M3h3ZVZNZWRPNkRJUEJ6YmZLckVtTHp3bEg2VGx2NXVLOTRxd0NuYWtIWkJtK0VEMGQxekYwd3VyaVgvN2ExSWszZEVLdEUyWDU4dG1naDZJNzRIQm1kbk9JRmhiTEs3ZktKcDRNSEpoWFNmcU1IZS9UZmdWenU4c2lldUFXZ3BxS0tmRTdvNGhLOVZWdnNKbmxCVVhvTnRFVlFsbWk1KzE1Q0R5dUZ5NlJuQ0ZmbkJNN1NVOGtlSmZJTjdldHVYM1RHQlY2bXJPQ1RKdkU1ZHptUGtPaHZ5azRGOXA3VXNacU03N2YzN2lSWjNnVVpZUjJsQ0F4QlRKZjdyTGgxQVBiaFhZTTR0V0NSV0p2N1NQb0hObzlKdEpmMHI4cTZZWnBHRWhFdUJPTVZsUGlRVHk1UVhYTUo5QXdCS25jS1dUbFRjb1RsVlRLeElGamdkbXJ2amxNbXNKNmtGOVdXMVk1dnhhMHd0a1l1NEJYdFVDbjUvdkZPczdkb3JrUFNtYmtETlRaK0RoMy92eUVNemQzcWZnWGk5ajFtbGxwRHd3REFZRktvd2lCd3NQdTB1NUtJeHM5c0NpN1kwaHRsRWMyLzFMdUtIS0E2MDM3SG5hZmxMYXNDNWlFVFF0TjA1NTBXdTZmc2xDWUY3dXI4VXl6UjcvSjJaQjFEMG9WZ0ZtaG14WDZVUGx3TXBWcjdZenFZN0JXaDVzMlN5cWg2eWpwOHhlbjl6amJVUHZWOWxJTUlwblRIakw3NmVMazV2QlRKR2VlT0I0eldWRmxzK2FEVEtYLyt3QXQ1VFRETk15Z1NYU29DS29JbVIrbWJLdmxGTUxuZDY5a2hUOWhoUmpGbzNaTlk2aG9LdEtIb0FaQzdoejh2MXpjWGF4N29PV3FLbnU5YWZjbCt6MXArK1lMN3ZiZGtzNFpHeW45cmp3SHJ0bkErWTZpcjRIdFdoZUxrUzRVVUY3SVBqQTZTTDFndW4ySkExRFFRTHRtdnNRSTBZd2dKZW8wQlJES0tlcWNlcG9TSmtORE1HV3pZOFU3NUFBbUt2ck1nUDVzZ2NuOWNtMzQ2UXc2VGZPbkx1QmlIQkw2bVVGKy9mYlE4anppRWRKSC9GYUFtUE9FT0RhcXByYjQ0Y1hGZHhmall0NklNMVFBNkhaRW1na2J3NHU0U0huL2N6dXp1eDZ3enI3VEplSEMyc1o2MnZ0S2pBeDRlVVUrT21SdHRjSXlRS0hiUXZrZWNBY3ROcmRNcEhUa0ZWeFJseTBiMWE4NGluNCtFK0FLa215d3JZSVhyUHQrZTJOOU93WGNIcW1MbURsTkQ0OGZnVWdpTkwzS0ZXR0ZEVk91L2dWd0ZDeFpyNjVlMVFSQ1dsVmQrdlFTZzA3TGdnanI4UGpvUVFjQjNWMnpoOElHZ1J6YVFiUDdIVk4xaTY4eEF5ZFNwM2JmbVlJTmRnRUxTMkV5ODM1NlZmaFNIbGJzUkNrSjhndW5hTTM0M3NyNEVsT05YTnZCcVhVVFB5ZWlTS1R5bzNJdHZQMWJHb2Y3bzJoNVZ3c1FQQThKdmNhTHdwV1VJQ0ZiU1hDZWFEMVZ6N2hYNHNKWDNIb21nUEVUMXZWU1FsN2M5RU9PZTN0STNmVmRKbWZhRzhMeVNRU0o1U2dNbzQ4NFNWTmk1OXRBVlU4cENLOVdVL3pFcUpBb1IyaUlqWkxJTmlVVmJkQWl6YUFFcmhZY0RjRVNwMUE9IiwiZGF0YWtleSI6IkFRSUJBSGdWVEN5a1ZSRWdjTXZqNG1EUWt0eVc0eWFrNVFDQlZZQ09IenM4cVRWZHZRRzRGR2NKeGVQSWFsRlJlL0Y3VUVvdEFBQUFmakI4QmdrcWhraUc5dzBCQndhZ2J6QnRBZ0VBTUdnR0NTcUdTSWIzRFFFSEFUQWVCZ2xnaGtnQlpRTUVBUzR3RVFRTXRBNXFVckhHeERuRWhpUllBZ0VRZ0R2MVNVNUJ0L0RKUUdXWlQ4QzArejlmcS8xdmhwSW9yRlVRVEttZHcrRWxrZDRDcS94TjJTWU52TnBxTjR1eE9Ydm5pcEl3VGpwTWdJK2I4dz09IiwidmVyc2lvbiI6IjIiLCJ0eXBlIjoiREFUQV9LRVkiLCJleHBpcmF0aW9uIjoxNjI2MTI0MTYwfQ==",
			Password: password,
			ServerAddress: dockerAddress,
		}
	} else {
		auth = types.AuthConfig{
			//Username: "AWS",
			Username: username,
			//Password: "eyJwYXlsb2FkIjoiU0Z4ZlpGcG5FT3FOdFA1eWJJem9jL2tXSHhwL2ZBa0FpTTNMZjJqRVBjT0ROcXdQTWdSWGNFOSt6aTN5NUJxQWZOZEVwZi9qL2pjMXl0T3BrcWdRc290NTNXTStMZG9xdWhiVmpjdEcvWlB4SG1reFliVllTNTE1VlBPNTFCS2pYTWMwT3p6M0VHMU5XTjFsMU9xSWFhdVFsbDVrQmZQUHJZV1Frd2ZkQmxvd2c5NWFPK3VqaWcwYXRzSWZreUtMeTZ5b0QrT3NkaWszbjB6QWNpNnVscjBRZDZKVHl6NGFTSkZpVDlSMXJVWS9TMU1oWEZ3cWhEeHhSNTU2YWxtN2EvL3I1enhJUHRNVE1LWWVtWWxjejIzckQvL0JUWWtVT21aWjU2SE4zK0dOR0RpcGpab1l3RXlndjFGVGFKTzdUMUhWS1owUS9KWGhsdEc3b1E1UlY0SFk4dzQyNldXTEh6M3h3ZVZNZWRPNkRJUEJ6YmZLckVtTHp3bEg2VGx2NXVLOTRxd0NuYWtIWkJtK0VEMGQxekYwd3VyaVgvN2ExSWszZEVLdEUyWDU4dG1naDZJNzRIQm1kbk9JRmhiTEs3ZktKcDRNSEpoWFNmcU1IZS9UZmdWenU4c2lldUFXZ3BxS0tmRTdvNGhLOVZWdnNKbmxCVVhvTnRFVlFsbWk1KzE1Q0R5dUZ5NlJuQ0ZmbkJNN1NVOGtlSmZJTjdldHVYM1RHQlY2bXJPQ1RKdkU1ZHptUGtPaHZ5azRGOXA3VXNacU03N2YzN2lSWjNnVVpZUjJsQ0F4QlRKZjdyTGgxQVBiaFhZTTR0V0NSV0p2N1NQb0hObzlKdEpmMHI4cTZZWnBHRWhFdUJPTVZsUGlRVHk1UVhYTUo5QXdCS25jS1dUbFRjb1RsVlRLeElGamdkbXJ2amxNbXNKNmtGOVdXMVk1dnhhMHd0a1l1NEJYdFVDbjUvdkZPczdkb3JrUFNtYmtETlRaK0RoMy92eUVNemQzcWZnWGk5ajFtbGxwRHd3REFZRktvd2lCd3NQdTB1NUtJeHM5c0NpN1kwaHRsRWMyLzFMdUtIS0E2MDM3SG5hZmxMYXNDNWlFVFF0TjA1NTBXdTZmc2xDWUY3dXI4VXl6UjcvSjJaQjFEMG9WZ0ZtaG14WDZVUGx3TXBWcjdZenFZN0JXaDVzMlN5cWg2eWpwOHhlbjl6amJVUHZWOWxJTUlwblRIakw3NmVMazV2QlRKR2VlT0I0eldWRmxzK2FEVEtYLyt3QXQ1VFRETk15Z1NYU29DS29JbVIrbWJLdmxGTUxuZDY5a2hUOWhoUmpGbzNaTlk2aG9LdEtIb0FaQzdoejh2MXpjWGF4N29PV3FLbnU5YWZjbCt6MXArK1lMN3ZiZGtzNFpHeW45cmp3SHJ0bkErWTZpcjRIdFdoZUxrUzRVVUY3SVBqQTZTTDFndW4ySkExRFFRTHRtdnNRSTBZd2dKZW8wQlJES0tlcWNlcG9TSmtORE1HV3pZOFU3NUFBbUt2ck1nUDVzZ2NuOWNtMzQ2UXc2VGZPbkx1QmlIQkw2bVVGKy9mYlE4anppRWRKSC9GYUFtUE9FT0RhcXByYjQ0Y1hGZHhmall0NklNMVFBNkhaRW1na2J3NHU0U0huL2N6dXp1eDZ3enI3VEplSEMyc1o2MnZ0S2pBeDRlVVUrT21SdHRjSXlRS0hiUXZrZWNBY3ROcmRNcEhUa0ZWeFJseTBiMWE4NGluNCtFK0FLa215d3JZSVhyUHQrZTJOOU93WGNIcW1MbURsTkQ0OGZnVWdpTkwzS0ZXR0ZEVk91L2dWd0ZDeFpyNjVlMVFSQ1dsVmQrdlFTZzA3TGdnanI4UGpvUVFjQjNWMnpoOElHZ1J6YVFiUDdIVk4xaTY4eEF5ZFNwM2JmbVlJTmRnRUxTMkV5ODM1NlZmaFNIbGJzUkNrSjhndW5hTTM0M3NyNEVsT05YTnZCcVhVVFB5ZWlTS1R5bzNJdHZQMWJHb2Y3bzJoNVZ3c1FQQThKdmNhTHdwV1VJQ0ZiU1hDZWFEMVZ6N2hYNHNKWDNIb21nUEVUMXZWU1FsN2M5RU9PZTN0STNmVmRKbWZhRzhMeVNRU0o1U2dNbzQ4NFNWTmk1OXRBVlU4cENLOVdVL3pFcUpBb1IyaUlqWkxJTmlVVmJkQWl6YUFFcmhZY0RjRVNwMUE9IiwiZGF0YWtleSI6IkFRSUJBSGdWVEN5a1ZSRWdjTXZqNG1EUWt0eVc0eWFrNVFDQlZZQ09IenM4cVRWZHZRRzRGR2NKeGVQSWFsRlJlL0Y3VUVvdEFBQUFmakI4QmdrcWhraUc5dzBCQndhZ2J6QnRBZ0VBTUdnR0NTcUdTSWIzRFFFSEFUQWVCZ2xnaGtnQlpRTUVBUzR3RVFRTXRBNXFVckhHeERuRWhpUllBZ0VRZ0R2MVNVNUJ0L0RKUUdXWlQ4QzArejlmcS8xdmhwSW9yRlVRVEttZHcrRWxrZDRDcS94TjJTWU52TnBxTjR1eE9Ydm5pcEl3VGpwTWdJK2I4dz09IiwidmVyc2lvbiI6IjIiLCJ0eXBlIjoiREFUQV9LRVkiLCJleHBpcmF0aW9uIjoxNjI2MTI0MTYwfQ==",
			Password: password,
		}
	}

	authBytes, _ := json.Marshal(auth)
	authBase64 := base64.URLEncoding.EncodeToString(authBytes)
	push, err := cli.ImagePush(ctx, target, types.ImagePushOptions{All: true,
		RegistryAuth:authBase64})
	if err != nil {
		panic(err)
	}
	termFd, isTerm := term.GetFdInfo(os.Stderr)
	err = jsonmessage.DisplayJSONMessagesStream(push, os.Stderr, termFd, isTerm, nil)
	if err != nil {
		panic(err)
	}
}
