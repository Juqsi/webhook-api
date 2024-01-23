package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func buildAndDeployImage(fullRepoName, cloneDirectory string) error {
	tmp := strings.Split(fullRepoName, "/")
	if len(tmp) < 2 {
		return errors.New("kein Repo name bzw zu kurz")
	}
	stackName := tmp[1]
	var imageName string
	imageName = tmp[1] + ":latest"

	renameCmd := exec.Command("docker", "tag", imageName, tmp[1]+":backup")
	renameCmd.Stdout = os.Stdout
	renameCmd.Stderr = os.Stderr
	err := renameCmd.Run()
	if err != nil {
		return err
	}

	cmd := exec.Command("docker-compose", "-p", stackName, "down")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		fmt.Println("stack existiert nicht")
	}

	cmd = exec.Command("docker-compose", "-p", stackName, "-f", "./"+cloneDirectory+"/docker-compose.yml", "up", "-d")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
