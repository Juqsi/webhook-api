package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"os"
	"os/exec"
)

func deployFromGithub(ctx *fiber.Ctx) error {
	baseURL := "github.com/"
	var data map[string]interface{}

	if err := json.Unmarshal(ctx.Body(), &data); err != nil {
		//TODO send Error Mail
		fmt.Println(err)
		return err
	}
	repository, ok := data["repository"].(map[string]interface{})
	if !ok {
		//TODO send Error Mail
		fmt.Println(ok)
		return errors.New("Fehler 2")
	}
	fullRepoName, ok := repository["full_name"].(string)
	if !ok {
		//TODO send Error Mail
		fmt.Println(ok)
		return errors.New("Fehler 3")
	}
	cloneDirectory, err := httpsCloneRepo(baseURL, fullRepoName)
	if err != nil {
		//TODO send Error Mail
		fmt.Println(err)
		return err
	}
	err = buildAndDeployImage(fullRepoName, cloneDirectory)
	if err != nil {
		//TODO send Error Mail
		fmt.Println(err)
		return err
	}
	return nil
}

func deployFromGitea(ctx *fiber.Ctx) error {
	baseURL := "gitea.justus-siegert.de/"
	var data map[string]interface{}

	if err := json.Unmarshal(ctx.Body(), &data); err != nil {
		//TODO send Error Mail
		fmt.Println(err)
		return err
	}
	repository, ok := data["repository"].(map[string]interface{})
	if !ok {
		//TODO send Error Mail
		fmt.Println(ok)
		return errors.New("Fehler 2")
	}
	fullRepoName, ok := repository["full_name"].(string)
	if !ok {
		//TODO send Error Mail
		fmt.Println(ok)
		return errors.New("Fehler 3")
	}
	cloneDirectory, err := httpsCloneRepo(baseURL, fullRepoName)
	if err != nil {
		//TODO send Error Mail
		fmt.Println(err)
		return err
	}
	err = buildAndDeployImage(fullRepoName, cloneDirectory)
	if err != nil {
		//TODO send Error Mail
		fmt.Println(err)
		return err
	}

	return nil
}

func httpsCloneRepo(baseURL, fullRepoName string) (string, error) {
	cloneDirectory := fullRepoName + "-" + uuid.New().String()

	username := os.Getenv(fullRepoName + "_name")
	accessToken := os.Getenv(fullRepoName + "_token")

	fullURL := fmt.Sprintf("https://%s:%s@%s%s.git", username, accessToken, baseURL, fullRepoName)

	cmd := exec.Command("git", "clone", fullURL, cloneDirectory)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return cloneDirectory, nil
}
