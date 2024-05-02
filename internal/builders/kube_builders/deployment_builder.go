package kubebuilders

import (
	"errors"
	"fmt"
	"os"

	"github.com/FelipeMCassiano/gorvus/internal/builders"
	"github.com/manifoldco/promptui"
)

func DeploymentKubeBuilder() error {
	depl := setDeploySettings()
	datafile, err := templatesContent.ReadFile("templates/kube_deployment.tmpl")
	if err != nil {
		return err
	}
	prompt := promptui.Prompt{
		Label:    "File Name",
		Validate: validatedPrompt,
	}
	fileName, _ := prompt.Run()

	fileYaml := fmt.Sprintf("%s.yaml", fileName)
	file, err := os.Create(fileYaml)
	if err != nil {
		return err
	}

	builders.ApplyTemplate(file, string(datafile), *depl)

	return nil
}

func setDeploySettings() *KubeDatabuilder {
	data := new(KubeDatabuilder)

	prompts := []struct {
		Label   string
		Pointer *string
	}{
		{"Api version", &data.ApiVersion},
		{"Deploy name", &data.Name},
		{"App name", &data.AppName},
		{"Number of replicas", &data.Replicas},
		{"Image", &data.Image},
		{"Ports", &data.Ports},
	}

	for _, p := range prompts {
		prompt := promptui.Prompt{
			Label:    p.Label,
			Validate: validatedPrompt,
		}
		result, err := prompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return nil
		}

		*p.Pointer = result

	}

	return data
}

func validatedPrompt(input string) error {
	if len(input) < 1 {
		return errors.New("This field is required")
	}
	return nil
}
