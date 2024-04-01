<div align='center'>
  <h1>gorvus</h1>
  <p> 
gorvus is a command-line interface (CLI) tool written in Go that simplifies the process of generating Dockerfiles and docker-compose.yml files for your projects. With gorvus, you can quickly scaffold Docker configurations without manual editing, saving time and effort. </p>
  <img src='https://img.shields.io/github/languages/top/FelipeMCassiano/gorvus' />
  <img src='https://img.shields.io/github/last-commit/FelipeMCassiano/gorvus' />
</div>


## Features

- **Dockerfile Generation**: Generate Dockerfiles for your projects with ease.
- **docker-compose.yml Generation**: Coming Soon
- **Customizable Templates**: Configure templates to suit your project's specific requirements.
- **CLI Interface**: Intuitive command-line interface for easy interaction.

## Installation

To install gorvus, you need to have Go installed on your system. Then, you can install it using the following command:

```bash
go install github.com/FelipeMCassiano/gorvus/gorvus@latest
```

## Usage
Once installed, you can use gorvus to generate Dockerfiles and docker-compose.yml files for your projects.

### Generate Dockerfile
```
gorvus Dockerfile <projectName> <language>
```


<details>
  <summary><strong>ℹ️ Info</strong></summary>
 Currently, only the language Go supports Dockerfile generation.
</details>


### Generate docker-compose.yml (Coming Soon)

The feature to generate a docker-compose.yml file will be available soon. Stay tuned for updates!

## Contruibuing
If you're interested in contributing to this project, consider reading the [Contributing Guide](contributing.md)

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
