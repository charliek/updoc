package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

type PortMapping struct {
	Protocol      string
	HostPort      uint16
	ContainerPort uint16
}

type UpdocApp struct {
	Name    string
	Image   string
	Stdin   bool
	Stdout  bool
	Stderr  bool
	Ports   []PortMapping
	Env     map[string]string
	Command []string
}

type appSlice []*UpdocApp

func (apps appSlice) findByName(name string) *UpdocApp {
	for _, app := range apps {
		if app.Name == name {
			return app
		}
	}
	return nil
}

func LoadFromDir(dirPath string) (appSlice, error) {
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	apps := make(appSlice, 0, 3)
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".json") {
			app, err := AppFromFile(fmt.Sprintf("%s/%s", dirPath, file.Name()))
			if err != nil {
				return nil, fmt.Errorf("Error loading app file %v: %v", file.Name(), err)
			}
			apps = append(apps, app)
		}
	}
	return apps, nil
}

func AppFromFile(appFile string) (*UpdocApp, error) {
	var app UpdocApp
	appBytes, err := ioutil.ReadFile(appFile)
	if err != nil {
		return nil, fmt.Errorf("Error reading file %s: %v", appFile, err)
	}

	if err := json.Unmarshal(appBytes, &app); err != nil {
		return nil, fmt.Errorf("Error unmarshaling json: %v", err)
	}
	return &app, nil
}
