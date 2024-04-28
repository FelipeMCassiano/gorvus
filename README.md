<div align='center'>
  <h1>Gorvus (Beta)</h1>
  <p>gorvus is a command-line interface (CLI) tool written in Go that simplifies the process of generating Dockerfiles and docker-compose.yml files for your projects. With gorvus, you can quickly scaffold Docker configurations without manual editing, saving time and effort.</p>
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
go install github.com/FelipeMCassiano/gorvus/gorvus@v1.2.0
```

## Usage

Once installed, you can use gorvus to generate Dockerfiles and docker-compose.yml files for your projects.

### Generate Dockerfile
* You can use Gorvus interactively without flags:
```bash
gorvus gend
```
* OR specify language and project name:
```bash
gorvus createDockerfile --language<language> --projectName<projectName>
```

> [!NOTE]
> Currently, only the languages Go, Rust, Node(js and ts), Java(gradle), Dotnet and Bun supports Dockerfile generation.

### Generate docker-compose.yml
```bash
gorvus compose 
```

> [!NOTE]
> Currently, only Postgres, Mysql and MongoDb have support for docker-compose with template generation.

### Add Services into docker-compose.yml
* You can use Gorvus interactively without flags, or specify service details:
```bash
gorvus compose add
```
* OR provide details using flags:
```bash
gorvus compose add --image<image> --service<serviceName> --ports<ports> --envs<environment> --networks<networkName> --hostname<hostname>
```
### Add Networks into docker-compose.yml
* You can use Gorvus interactively without flags:
```bash
gorvus compose add-net
```
* OR specify network details:
```bash
gorvus compose add-net --name<network name> --driver<network driver> --name-docker<reference this network when you're connecting containers>
```

### Remove Services or Networks in docker-compose.yml

* To remove a service
```bash
gorvus compose rm -s<service name>
```
* To remove a network
```bash
gorvus compose rm -n<network name>
```

## Contributing

If you're interested in contributing to this project, consider reading the [Contributing Guide](CONTRIBUTING.md)

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
