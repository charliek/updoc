package main

import (
	"fmt"
	"github.com/fsouza/go-dockerclient"
	"io"
	"strings"
)

func pullImage(client *docker.Client, app *UpdocApp, out io.Writer) error {
	ops := updocToPullOpts(app)
	ops.OutputStream = out
	return client.PullImage(*ops, docker.AuthConfiguration{})
}

func updocToPullOpts(app *UpdocApp) *docker.PullImageOptions {
	return &docker.PullImageOptions{
		Registry:   app.Registry,
		Repository: app.Image,
		Tag:        app.Tag,
	}
}

// TODO consider using the below images methods ported over from the python client
// TODO write tests before doing so

// Figure out the repository url form the image name. Ported from the python client.
func resolveRepositoryName(repoName string) (*docker.PullImageOptions, error) {
	if strings.Contains(repoName, "://") {
		return nil, fmt.Errorf("Repository name cannot contain a scheme (%s)", repoName)
	}
	parts := strings.SplitN(repoName, ":", 2)
	if !strings.Contains(parts[0], ".") && !strings.Contains(parts[0], ":") && parts[0] != "localhost" {
		// This is a docker index repo (ex: foo/bar or ubuntu)
		return &docker.PullImageOptions{
			Registry:   indexUrl,
			Repository: repoName,
			Tag:        "latest",
		}, nil
	}
	if len(parts) < 2 {
		return nil, fmt.Errorf("Invalid repository name (%s)", repoName)
	}
	if strings.Contains(parts[0], "index.docker.io") {
		return nil, fmt.Errorf("Invalid repository name, try (%s) instead", parts[1])
	}
	return &docker.PullImageOptions{
		Registry:   expandRegistryUrl(parts[0]),
		Repository: parts[1],
		Tag:        "latest",
	}, nil
}

func expandRegistryUrl(hostname string) string {
	if strings.HasPrefix(hostname, "http:") || strings.HasPrefix(hostname, "https:") {
		if strings.LastIndex(hostname, "/") < 9 {
			hostname = fmt.Sprintf("%s/v1/", hostname)
		}
		return hostname
	}
	if pingRegistry(fmt.Sprintf("https://%s/v1/_ping", hostname)) {
		return fmt.Sprintf("https://%s/v1/", hostname)
	}
	return fmt.Sprintf("http://%s/v1/", hostname)
}

func pingRegistry(url string) bool {
	// TODO need to hit a url with a timeout to check for https

	// def ping(url):
	//     try:
	//         res = requests.get(url)
	//     except Exception:
	//         return False
	//     else:
	//         return res.status_code < 400

	return false
}
