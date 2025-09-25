dagger -m github.com/shykes/dagger@local-defaults <<EOF

    . | cli | binary --platform "darwin/arm64" | export "dagger-cli"

EOF