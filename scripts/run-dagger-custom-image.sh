docker run --rm -d \
  -v /var/lib/dagger \
  -v $HOME/workspace/ssl/:/usr/local/share/ca-certificates/ \
  --name "dagger-engine-custom" \
  --privileged \
  "dagger-engine-custom"

export _EXPERIMENTAL_DAGGER_RUNNER_HOST=docker-container://dagger-engine-custom