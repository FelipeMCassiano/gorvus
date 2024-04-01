# gorvus - Dockerfile Generator

gorvus is a command-line interface (CLI) tool written in Go that simplifies the process of generating Dockerfiles and docker-compose.yml files for your projects. With GoCLI, you can quickly scaffold Docker configurations without manual editing, saving time and effort.

## Features

- **Dockerfile Generation**: Generate Dockerfiles for your projects with ease.
- **docker-compose.yml Generation**: Generate docker-compose.yml files to orchestrate multi-container Docker applications.
- **Customizable Templates**: Configure templates to suit your project's specific requirements.
- **CLI Interface**: Intuitive command-line interface for easy interaction.

## Installation

To install gorvus, you need to have Go installed on your system. Then, you can install it using the following command:

```bash
go install github.com/FelipeMCassiano/gorvus
```

## Usage
Once installed, you can use gorvus to generate Dockerfiles and docker-compose.yml files for your projects.

### Generate Dockerfile
```
gorvus Dockerfile <projectName> <language>
```
### Generate docker-compose.yml (Coming Soon)

The feature to generate a docker-compose.yml file will be available soon. Stay tuned for updates!

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
