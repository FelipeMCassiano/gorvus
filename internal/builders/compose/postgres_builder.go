package compose

import (
	"fmt"
	"os"

	"github.com/FelipeMCassiano/gorvus/internal/builders"
	"github.com/manifoldco/promptui"
)

func PostgreSQLComposeFileBuilder() error {
	compose := setPostgresSettings()

	path := fmt.Sprintf("templates/%s.tmpl", "postgres")

	datafile, err := templatesContent.ReadFile(path)
	if err != nil {
		return err
	}

	file, err := os.Create("docker-compose.yml")
	if err != nil {
		return err
	}
	defer file.Close()

	builders.ApplyTemplate(file, string(datafile), *compose)

	return nil
}

func setPostgresSettings() *ComposeData {
	data := new(ComposeData)

	prompts := []struct {
		Label   string
		Pointer *string
	}{
		{"docker-compose Version (Default: 3.9)", &data.Version},
		{"Image Version (Default: latest)", &data.ImageVersion},
		{"DB Name (Default: DB)", &data.DbName},
		{"DB User (Default: USER)", &data.DbUser},
		{"DB Password (Default: PASS)", &data.DbPass},
		{"Restart (Default: no)", &data.Restart},
		{"Ports (Default: 5432:5432)", &data.Ports},
		{"CPU (Default: 1)", &data.Cpu},
		{"Memory (MB) (Default: 500)", &data.Memory},
		{"Network Name (Default: network)", &data.NetworkName},
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
	if len(data.Restart) == 0 {
		data.Restart = "no"
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

	return data
}
