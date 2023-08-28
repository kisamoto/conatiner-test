package main

import (
	"context"
	"fmt"
	"log"

	"github.com/containers/podman/v4/pkg/bindings"
	"github.com/containers/podman/v4/pkg/bindings/containers"
	"github.com/containers/podman/v4/pkg/specgen"
)

const socket = "unix://run/podman/podman.sock"

func main() {
	conn, err := bindings.NewConnection(context.Background(), socket)
	if err != nil {
		log.Fatalf("unable to connect to podman: %v", err)
	}
	if err := CreateContainer(conn, "hello-world"); err != nil {
		log.Fatalf(err.Error())
	}
	if err := StartContainer(conn, "hello-world"); err != nil {
		log.Fatalf(err.Error())
	}
}

func CreateContainer(conn context.Context, id string) error {
	s := specgen.NewSpecGenerator("hello-world:latest", false)
	s.Name = id

	_, err := containers.CreateWithSpec(conn, s, nil)
	if err != nil {
		return fmt.Errorf("create container: %w", err)
	}
	log.Print("podman container created")
	return nil
}

func StartContainer(conn context.Context, id string) error {
	return containers.Start(conn, id, nil)
}
