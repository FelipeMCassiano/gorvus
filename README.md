<div align='center'>
  <h1>Gorvus v2</h1>
  <p>gorvus is a command-line interface tool written in Go that simplifies the process of generating Dockerfiles and docker-compose.yml files for your projects. With gorvus, you can quickly scaffold Docker configurations without manual editing, saving time and effort.</p>
  <img src='https://img.shields.io/github/languages/top/FelipeMCassiano/gorvus' alt='GitHub top language' />
  <img src='https://img.shields.io/github/last-commit/FelipeMCassiano/gorvus' alt='GitHub last commit' />

</div>

> [!TIP]
> Looking for a dead-simple automation tool? [ruke](https://github.com/kauefraga/ruke) is waiting for you!

## Features

- **Dockerfile Generation**: Generate Dockerfiles for your projects with ease.
- **docker-compose.yml Generation**: Generate docker-compose.yml files to orchestrate multi-container Docker applications.
- **Customizable Templates**: Configure templates to suit your project's specific requirements.
- **Command-line Interface**: Use an intuitive and fancy interface to maximize your experience.

## Installation

To install gorvus, you need to have Go installed on your system. Then, you can install it using the following command:

```bash
go install github.com/FelipeMCassiano/gorvus/v2/gorvus@v2.0.2
```
## Usage

Once installed, you can use gorvus to generate Dockerfiles and docker-compose.yml files for your projects.

| Command | Flags | Description | Interactive Mode|
| :---: | :--- | :--- | :---:|
| `gorvus createDockerfile` | `-p --projectName <PROJECT-NAME>`, `-l --language <LANGUAGE-TEMPLATE>` |  Create Dockerfile based on input language and project name | yes |
| `gorvus compose` | `--cd <DIRECTORY>` | Manages directory's docker-compose.yml | no |


> **_NOTE:_** All subcommands inherits flags from their parent.
> e.g: `gorvus compose --cd add ...`

| Subcommand | Flags | Description | Interactive Mode |
| :-----: | :--- | :--- | :---: |
| `gorvus compose new` | doens't have flags | Create a new docker-compose.yml file with or without a template | yes |
| `gorvus compose add` | `-s --service <SERVICE-NAME>`, `-i --image <IMAGE>`, `-p --ports <PORTS>`, `-e --envs <ENVS>`, `-n --networks <NETWORK>`, `--hs <HOSTNAME>` | Adds a new service into docker-compose.yml | yes |
| `gorvus compose add-net` | `-n --name<NETWORK-NAME>`, `-d --driver<NETWORK-DRIVE>` ,`-x --name-network<reference this network when you're connecting containers>` | Adds a new network into docker-compose.yml | yes |
| `gorvus compose rm ` | `-s --service <SERVICE-NAME>`, `-n --network<NETWORK-NAME>` | Remove services or networks in docker-compose.yml | yes |

### Templates
- Currently, only the languages Go, Rust, Node(js and ts), Java(gradle), Dotnet and Bun supports Dockerfile generation.
-  Currently, only Postgres, Mysql and MongoDb have support for docker-compose with template generation.

## Contributing

If you're interested in contributing to this project, consider reading the [Contributing Guide](CONTRIBUTING.md)

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
