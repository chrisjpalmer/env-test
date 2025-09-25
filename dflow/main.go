// A generated module for DaggerFlowModule functions
//
// This module has been generated via dagger init and serves as a reference to
// basic module structure as you get started with Dagger.
//
// Two functions have been pre-created. You can modify, delete, or add to them,
// as needed. They demonstrate usage of arguments and return types using simple
// echo and grep commands. The functions can be called from the dagger CLI or
// from one of the SDKs.
//
// The first line in this comment block is a short description line and the
// rest is a long description with more detail on the module's purpose or usage,
// if appropriate. All modules should have a short description.

package main

import (
	"context"
	"dagger/dagger-flow-module/internal/dagger"
	"fmt"
)

type DFlow struct {
	SSH      *dagger.Socket
	Secret   *dagger.Secret
	StrValue string
	File     *dagger.File
}

func New(
	ssh *dagger.Socket,
	secret *dagger.Secret,
	strvalue string,
	file *dagger.File,
) *DFlow {
	return &DFlow{
		SSH:      ssh,
		Secret:   secret,
		StrValue: strvalue,
		File:     file,
	}
}

// Returns a container that echoes whatever string argument is provided
func (m *DFlow) Do(ctx context.Context) (*dagger.Container, error) {
	const ghHost = "github.com ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIOMqqnkVzrm0SdG6UOoqKLsabgH5C9okWi0dh2l9GKJl"

	sec, err := m.Secret.Plaintext(ctx)
	if err != nil {
		return nil, err
	}

	fmt.Println("the secret is", sec)

	return dag.Container().
		From("alpine:latest").
		WithExec([]string{"apk", "add", "git", "openssh"}).
		WithWorkdir("/app").
		WithFile("file", m.File).
		WithSecretVariable("SECRET", m.Secret).
		WithEnvVariable("STRVALUE", m.StrValue).
		WithUnixSocket("/var/ssh.sock", m.SSH).
		WithEnvVariable("SSH_AUTH_SOCK", "/var/ssh.sock").
		WithNewFile("/root/.ssh/known_hosts", ghHost).
		WithExec([]string{"git", "clone", "git@github.com:nine-digital/library-ci-workflows"}).
		Terminal(), nil
}
