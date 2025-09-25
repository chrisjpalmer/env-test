dagger -m github.com/shykes/dagger@local-defaults <<EOF

    docker_sock=\$(host | unix-socket /var/run/docker.sock)

    . | load-to-docker $docker_sock "custom-dagger-engine" "linux/arm64"
EOF