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

	// type Config struct {
	// 	Hostname        string
	// 	Domainname      string
	// 	User            string
	// 	Memory          int64
	// 	MemorySwap      int64
	// 	CpuShares       int64
	// 	AttachStdin     bool
	// 	AttachStdout    bool
	// 	AttachStderr    bool
	// 	PortSpecs       []string
	// 	ExposedPorts    map[Port]struct{}
	// 	Tty             bool
	// 	OpenStdin       bool
	// 	StdinOnce       bool
	// 	Env             []string
	// 	Cmd             []string
	// 	Dns             []string  // For Docker API v1.9 and below only
	// 	Image           string
	// 	Volumes         map[string]struct{}
	// 	VolumesFrom     string
	// 	WorkingDir      string
	// 	Entrypoint      []string
	// 	NetworkDisabled bool
	// }

	name := "test-green"
	create := true

	if create {
		// Create a container
		options := docker.CreateContainerOptions{
			Name: name,
			Config: &docker.Config{
				Image:        "charliek/docker-testing:green",
				AttachStdin:  true,
				AttachStdout: true,
				AttachStderr: true,
				ExposedPorts: map[docker.Port]struct{}{
					"9090/tcp": struct{}{},
				},
				Env: []string{"TEST4=TEST2", "FOO=BAR"},
			},
		}

		container, err := client.CreateContainer(options)
		if err != nil {
			fmt.Printf("Error creating container: %v\n\n", err)
		} else {
			fmt.Printf("Created the container %+\n\n", container)
		}
	}

	// inspect the container we just created
	c, err := client.InspectContainer(name)
	if err != nil {
		fmt.Printf("Error inspecting container: %v\n\n", err)
	} else {
		fmt.Printf("Found the container %+v\n\n", c)
	}

	// type HostConfig struct {
	// 	Binds           []string
	// 	ContainerIDFile string
	// 	LxcConf         []KeyValuePair
	// 	Privileged      bool
	// 	PortBindings    map[Port][]PortBinding
	// 	Links           []string
	// 	PublishAllPorts bool
	// 	Dns             []string  // For Docker API v1.10 and above only
	// }

	host := &docker.HostConfig{
		PortBindings: map[docker.Port][]docker.PortBinding{
			"9090/tcp": []docker.PortBinding{docker.PortBinding{
				HostIp:   "",
				HostPort: "9090",
			}},
		},
	}

	if err := client.StartContainer(name, host); err != nil {
		fmt.Printf("Error starting container %v\n\n", err)
	}

}
