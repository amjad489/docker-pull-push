# docker-pull-push
If you have faced when pulling images from docker hub due to api rate limits or if you want to migrate for any other reason. This cli will help to migrate docker images from docker hub to AWS ECR.

https://www.docker.com/increase-rate-limits#:~:text=The%20rate%20limits%20of%20100,the%20six%20hour%20window%20elapses.

## Run
Download the latest release: [https://github.com/amjad489/docker-pull-push/releases/tag/v0.1.0](https://github.com/amjad489/docker-pull-push/releases/tag/v0.1.0).
```
$ ./docker-pull-push
For example:

docker-pull-push -s docker.elastic.co/elasticsearch/elasticsearch:7.13 -t AWS_ACCOUNT_ID.dkr.ecr.us-east-1.amazonaws.com/elasticsearch:7.13.3

Usage:
  docker-pull-push [flags]

Flags:
  -p, --awsProfile string     AWS profile (default "default")
  -l, --awsRegion string      AWS Region (default "us-east-1")
  -h, --help                  help for docker-pull-push
  -r, --registryType string   ECR (default "docker")
  -s, --sourceImage string    docker.elastic.co/elasticsearch/elasticsearch:7.13
  -t, --targetImage string    AWS_ACCOUNT_ID.dkr.ecr.us-east-1.amazonaws.com/elasticsearch:7.13.3

```

# TODO List
- support migration to other repositories.