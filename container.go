package main

import (
	"fmt"
	"github.com/fsouza/go-dockerclient"
)

func createContainer(client *docker.Client, app *UpdocApp) (*docker.Container, error) {
	options := updocToContainerOpts(app)
	return client.CreateContainer(*options)
}

func updocToContainerOpts(app *UpdocApp) *docker.CreateContainerOptions {
	exposedPorts := make(map[docker.Port]struct{})
	for _, p := range app.Ports {
		exposedPorts[docker.Port(fmt.Sprintf("%d/%s", p.ContainerPort, p.Protocol))] = struct{}{}
	}

	env := []string{}
	for k, v := range app.Env {
		env = append(env, fmt.Sprintf("%s=%s", k, v))
	}

	println(fmt.Sprintf("%s:%s", app.Image, app.Tag))

	return &docker.CreateContainerOptions{
		Name: app.Name,
		Config: &docker.Config{
			Image:        fmt.Sprintf("%s:%s", app.Image, app.Tag),
			AttachStdin:  app.Stdin,
			AttachStdout: app.Stdout,
			AttachStderr: app.Stderr,
			ExposedPorts: exposedPorts,
			Env:          env,
		},
	}
}

func updocToHostConfig(app *UpdocApp) *docker.HostConfig {
	portBindings := make(map[docker.Port][]docker.PortBinding)
	for _, p := range app.Ports {
		k := docker.Port(fmt.Sprintf("%d/%s", p.ContainerPort, p.Protocol))
		portBindings[k] = []docker.PortBinding{docker.PortBinding{
			HostIp:   "",
			HostPort: fmt.Sprintf("%d", p.HostPort),
		}}
	}

	return &docker.HostConfig{
		PortBindings: portBindings,
	}
}

func startContainer(client *docker.Client, app *UpdocApp) error {
	host := updocToHostConfig(app)
	return client.StartContainer(app.Name, host)
}

func containerExists(client *docker.Client, name string) (bool, error) {
	if _, err := client.InspectContainer(name); err != nil {
		switch err.(type) {
		case *docker.NoSuchContainer:
			return false, nil
		default:
			return false, err
		}
	}
	return true, nil
}
