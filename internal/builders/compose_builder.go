package builders

import (
	"fmt"
	"os"
	"strings"

	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/manifoldco/promptui"
)

func BuilderComposefile(template string) error {
	compose := setComposeSettings()

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

	applyTemplate(file, string(datafile), *compose)

	return nil
}

func setComposeSettings() *ComposeData {
	data := new(ComposeData)

	prompts := []struct {
		Label   string
		Pointer *string
	}{
		{"docker-compose Version", &data.Version},
		{"Image Version", &data.ImageVersion},
		{"DB Name", &data.DbName},
		{"DB User", &data.DbUser},
		{"DB Password", &data.DbPass},
		{"Ports", &data.Ports},
		{"CPU", &data.Cpu},
		{"Memory (MB)", &data.Memory},
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

	if len(data.Version) == 0 {
		data.Version = "3.9"
	}
	if len(data.ImageVersion) == 0 {
		data.ImageVersion = "latest"
	}
	if len(data.DbName) == 0 {
		data.DbName = "DB"
	}
	if len(data.DbUser) == 0 {
		data.DbUser = "USER"
	}
	if len(data.DbPass) == 0 {
		data.DbPass = "PASS"
	}
	if len(data.Ports) == 0 {
		data.Ports = "5432:5432"
	}
	if len(data.Cpu) == 0 {
		data.Cpu = "1"
	}
	if len(data.Memory) == 0 {
		data.Memory = "500"
	}
	if len(data.NetworkName) == 0 {
		data.NetworkName = "network"
	}
	fmt.Println(text.FgBlue.Sprint("If some field is empty, default configs will be applied."))

	return data
}
