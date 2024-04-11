<div align='center'>
  <h1>gorvus (beta)</h1>
  <p>gorvus is a command-line interface (CLI) tool written in Go that simplifies the process of generating Dockerfiles and docker-compose.yml files for your projects. With gorvus, you can quickly scaffold Docker configurations without manual editing, saving time and effort.</p>
  <img src='https://img.shields.io/github/languages/top/FelipeMCassiano/gorvus' alt='GitHub top language' />
  <img src='https://img.shields.io/github/last-commit/FelipeMCassiano/gorvus' alt='GitHub last commit' />
</div>

## Features

- **Dockerfile Generation**: Generate Dockerfiles for your projects with ease.
- **docker-compose.yml Generation**: Generate docker-compose.yml files to orchestrate multi-container Docker applications.
- **Customizable Templates**: Configure templates to suit your project's specific requirements.
- **Command-line Interface**: Use an intuitive and fancy interface to maximize your experience.

## Installation

To install gorvus, you need to have Go installed on your system. Then, you can install it using the following command:

```bash
go install github.com/FelipeMCassiano/gorvus/cmd@latest
```

## Usage

Once installed, you can use gorvus to generate Dockerfiles and docker-compose.yml files for your projects.

### Generate Dockerfile

```bash
gorvus createDockerfile --language<language> --projectName<projectName>
```

> [!NOTE]
> Currently, only the languages Go, Rust, Node(js and ts) and Bun supports Dockerfile generation.

### Generate docker-compose.yml

```bash
gorvus compose create --template<template>
```

> [!NOTE]
> Currently, only Postgres have support for docker-compose with template generation.

#### Add Services into docker-compose.yml

```bash
gorvus compose add --image<image> --service<serviceName> --ports<ports> --env<environment> --networks<networkName> --hostname<hostname>
```
#### Add Networks into docker-compose.yml
```bash
gorvus compose add-net --name<network name> --driver<network driver> --name-docker<reference this network when you're connecting containers>


## Contributing

If you're interested in contributing to this project, consider reading the [Contributing Guide](CONTRIBUTING.md)

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
