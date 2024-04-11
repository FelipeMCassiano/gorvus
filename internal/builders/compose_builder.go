package builders

import (
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/v6/text"
)

func BuilderComposefile(input ComposeData, template string) error {
	applyDefaultConfigs(&input)

	path := fmt.Sprintf("templates/%s.tmpl", template)

	datafile, err := templatesContent.ReadFile(path)
	if err != nil {
		return err
	}

	file, err := os.Create("docker-compose.yml")
	if err != nil {
		return err
	}
	defer file.Close()

	applyTemplate(file, string(datafile), input)

	return nil
}

func applyDefaultConfigs(compose *ComposeData) {
	if len(compose.Cpu) == 0 {
		compose.Cpu = "1"
	}
	if len(compose.Version) == 0 {
		compose.Version = "3.9"
	}
	if len(compose.ImageVersion) == 0 {
		compose.ImageVersion = "latest"
	}
	if len(compose.DbName) == 0 {
		compose.DbName = "DB"
	}
	if len(compose.DbUser) == 0 {
		compose.DbUser = "USER"
	}
	if len(compose.DbPass) == 0 {
		compose.DbPass = "PASS"
	}
	if len(compose.Ports) == 0 {
		compose.Ports = "5432:5432"
	}
	if len(compose.Memory) == 0 {
		compose.Memory = "500"
	}
	if len(compose.NetworkName) == 0 {
		compose.NetworkName = "network"
	}
	fmt.Println(text.FgBlue.Sprint("If some flag is empty, default configs will be applied. To view all flags in this command, type `compose create -h"))
}
