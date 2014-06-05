package main

import (
	"fmt"
	"github.com/fsouza/go-dockerclient"
	"log"
	"os"
)

func main() {
	endpoint := os.Getenv("DOCKER_HOST")
	if endpoint == "" {
		endpoint = "unix:///var/run/docker.sock"
	}
	client, err := docker.NewClient(endpoint)
	if err != nil {
		log.Fatalf("Error connecting to docker: %v", err)
	}

	app, err := AppFromFile("/Users/cknudsen/projects/play/updoc/test/dtesting.json")

	if err != nil {
		log.Fatalf("Error loading test app file %v", err)
	}

	if err := pullImage(client, app, os.Stdout); err != nil {
		log.Fatalf("Error pulling the image %v", err)
	}

	// Create the specified container
	if _, err := createContainer(client, app); err != nil {
		log.Fatalf("Error creating container %s : %v", app.Name, err)
	}

	// inspect the container we just created
	c, err := client.InspectContainer(app.Name)
	if err != nil {
		log.Fatalf("Error inspecting container %s : %v", app.Name, err)
	} else {
		fmt.Printf("Found the container %+v\n\n", c)
	}

	// Start the container
	if err := startContainer(client, app); err != nil {
		fmt.Printf("Error starting container %v\n\n", err)
	}
}
