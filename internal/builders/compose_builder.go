package builders

import (
	"fmt"
	"os"
	"strings"

	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/manifoldco/promptui"
)

func BuilderComposefile(input ComposeData, template string) error {
	compose := setComposeSettings(input)

	path := fmt.Sprintf("templates/%s.tmpl", strings.ToLower(template))

	datafile, err := templatesContent.ReadFile(path)
	if err != nil {
		return err
	}

	file, err := os.Create("docker-compose.yml")
	if err != nil {
		return err
	}
	defer file.Close()

	applyTemplate(file, string(datafile), compose)

	return nil
}

func setComposeSettings(compose ComposeData) *ComposeData {
	data := ComposeData{}

	prompts := []struct {
		Label   string
		Pointer *string
	}{
		{"Version", &data.Version},
		{"Image Version", &data.ImageVersion},
		{"DB Name", &data.DbName},
		{"DB User", &data.DbUser},
		{"DB Password", &data.DbPass},
		{"Ports", &data.Ports},
		{"CPU", &data.Cpu},
		{"Memory", &data.Memory},
		{"Network Name", &data.NetworkName},
	}

	for _, p := range prompts {
		prompt := promptui.Prompt{
			Label: p.Label,
		}

		result, err := prompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return nil
		}

		*p.Pointer = result
	}

	if len(compose.Version) == 0 && len(data.Version) == 0 {
		data.Version = "3.9"
	}
	if len(compose.ImageVersion) == 0 && len(data.ImageVersion) == 0 {
		data.ImageVersion = "latest"
	}
	if len(compose.DbName) == 0 && len(data.DbName) == 0 {
		data.DbName = "DB"
	}
	if len(compose.DbUser) == 0 && len(data.DbUser) == 0 {
		data.DbUser = "USER"
	}
	if len(compose.DbPass) == 0 && len(data.DbPass) == 0 {
		data.DbPass = "PASS"
	}
	if len(compose.Ports) == 0 && len(data.Ports) == 0 {
		data.Ports = "5432:5432"
	}
	if len(compose.Cpu) == 0 && len(data.Cpu) == 0 {
		data.Cpu = "1"
	}
	if len(compose.Memory) == 0 && len(data.Memory) == 0 {
		data.Memory = "500"
	}
	if len(compose.NetworkName) == 0 && len(data.NetworkName) == 0 {
		data.NetworkName = "network"
	}
	fmt.Println(text.FgBlue.Sprint("If some field is empty, default configs will be applied. To view all flags in this command, type `compose create -h"))

	return &data
}
