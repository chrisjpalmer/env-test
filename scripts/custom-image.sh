dagger -m github.com/shykes/dagger/cmd/engine@local-defaults <<EOF

    docker_sock=\$(host | unix-socket /var/run/docker.sock)

    . | load-to-docker \$docker_sock --name "custom-dagger-engine" --platform "linux/arm64"
EOF