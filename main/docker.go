package main

import (
	"errors"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
	"os"
	"os/exec"
	"strings"
)

func buildAndDeployImage(fullRepoName, cloneDirectory string) error {
	tmp := strings.Split(fullRepoName, "/")
	if len(tmp) < 2 {
		return errors.New("kein Repo name bzw zu kurz")
	}
	var imageName string
	imageName = tmp[1] + ":latest"
	cmd := exec.Command("docker", "build", "./"+cloneDirectory, "-t", imageName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return err
	}
	ports, err := getExposedPorts(imageName + ":latest")
	if err != nil {
		return err
	}
	parameter := []string{}
	parameter = append(parameter, "run")
	for _, port := range ports {
		parameter = append(parameter, "-p")
		parameter = append(parameter, port+":"+port)
	}
	parameter = append(parameter, imageName)
	cmd = exec.Command("docker", parameter...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func getExposedPorts(imageName string) ([]string, error) {
	cli, err := client.NewClientWithOpts()
	if err != nil {
		return nil, err
	}
	defer cli.Close()

	imageInspect, _, err := cli.ImageInspectWithRaw(context.Background(), imageName)
	if err != nil {
		return nil, err
	}

	var exposedPorts []string
	for port := range imageInspect.Config.ExposedPorts {
		exposedPorts = append(exposedPorts, port.Port())
	}

	return exposedPorts, nil
}
