package main

import (
	"reflect"
	"testing"
)

var test1Expected = &UpdocApp{
	Name:     "dtesting",
	Image:    "charliek/docker-testing",
	Tag:      "green",
	Registry: indexUrl,
	Stdin:    false,
	Stdout:   true,
	Stderr:   true,
	Ports: []PortMapping{
		PortMapping{
			Protocol:      "tcp",
			HostPort:      9090,
			ContainerPort: 9090,
		},
	},
	Env: map[string]string{
		"FOO":  "BAR",
		"VAR2": "value",
	},
	Command: []string{"/bin/echo", "'12345' '6789'"},
}

var test2Expected = &UpdocApp{
	Name:     "pullreq",
	Image:    "charliek/pullreq",
	Tag:      "latest",
	Registry: indexUrl,
	Stdin:    false,
	Stdout:   true,
	Stderr:   false,
	Ports: []PortMapping{
		PortMapping{
			Protocol:      "udp",
			HostPort:      22,
			ContainerPort: 22,
		},
		PortMapping{
			Protocol:      "tcp",
			HostPort:      80,
			ContainerPort: 443,
		},
	},
	Env:     nil,
	Command: nil,
}

func TestAppConfigLoading(t *testing.T) {
	app, err := AppFromFile("examples/test1.json")
	if err != nil {
		t.Errorf("Error loading app config: %v", err)
	}

	if !reflect.DeepEqual(app, test1Expected) {
		t.Errorf("Configs are not equal")
	}
}

func TestLoadFromDir(t *testing.T) {
	apps, err := LoadFromDir("examples")
	if err != nil {
		t.Errorf("Error loading app config: %v", err)
	}

	if len(apps) != 2 {
		t.Errorf("Invalid number of apps %d", len(apps))
	}
	if !reflect.DeepEqual(apps[0], test1Expected) {
		t.Errorf("test1 is not valid")
	}
	if !reflect.DeepEqual(apps[1], test2Expected) {
		t.Errorf("test2 is not valid")
	}
}
