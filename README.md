# NothingsBland.com

## Development Tasks

Setup

The following commands will spin up the docker environment
```
$ source ./docker/docker_env.sh
$ docker-wipe && docker-build && docker-run
```

- docker-wipe   # Stop, kill and remove all docker processes and images
- docker-build  # Build the development docker images
- docker-run    # Run the development images with proper volumes
- app-build-run   # Will build and execute application

## Notes

- Only use Docker engine for local development
- Kubernetes will be attempted for server orchestration