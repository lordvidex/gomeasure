## gomeasure
gomeasure is a CLI tool that measures the lines of code and the count of files in a directory or project

<!--TODO : Create a table for all commands -->
## Installation
### Build yourself
This project is open source. You can simply clone this repo and build using:
```shell
go build -o bin/gomeasure main.go 
```

### Download releases
Checkout the most updated releases [here](https://github.com/lordvidex/gomeasure/releases/) and download the binary that corresponds to your operating system

### MacOS (homebrew)
1. Using homebrew
```bash
$ brew tap lordvidex/lordvidex
$ brew install gomeasure
$ gomeasure --version # to confirm

```

### Debian (amd64) 
1. To **install** on debian distributions copy and paste the following command in your terminal
```bash
curl "https://raw.githubusercontent.com/lordvidex/gomeasure/master/scripts/install.sh" | sh
```
### Docker usage (new)
You can use the docker image to run gomeasure.
```bash
docker run --rm -v $(pwd):/app lordvidex/gomeasure:alpine line /app
```
#### How to Use
- The `-v` argument is used to mount the directory to be measured to the container's `/app` directory. Therefore, to measure any directory instead of the current directory, you can replace `$(pwd)` with the path to the directory you want to measure.
- The `--rm` argument is used to remove the container after the command is executed to prevent dangling containers.
- All arguments after the image name are passed to the gomeasure command.
- The `/app` should not be replaced because source directory is mapped to `/app` in the container.

## Uninstall

### Debian (amd64)
1. To **uninstall** on debian distributions copy and paste the following command in your terminal
```bash
curl "https://raw.githubusercontent.com/lordvidex/gomeasure/master/scripts/uninstall.sh" | sh
```
## Usage
```bash
gomeasure --help
```

### Configuration File
gomeasure uses a configuration file to make it easier to fine tune the CLI flags.
gomeasure looks for a file named `.gomeasure.yaml` in the current directory. If it doesn't find it, it looks for it in the home directory. If no configuration file is found, it uses the default values.

Note that the priority of the configurations are as follows:
1. CLI flags (highest priority)
2. Configuration file (in current directory)
3. Configuration file (in home directory)
4. Default values (lowest priority)

Checkout the [example configuration file](./example/.gomeasure.yaml) to get started.

## DISCLAIMER:
Normally, the program reads all kinds of file and returns the number of lines in each of them. Therefore, it is up to the developer to filter out the files to be scanned with 
flags `-i` and `-I`.
