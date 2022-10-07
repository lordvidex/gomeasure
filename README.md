## gomeasure
gomeasure is a CLI tool that measures the lines of code and the count of files in a directory or project

<!--TODO : Create a table for all commands -->
## Installation
### Build yourself
This project is open source. You can simply clone this repo and build using:
```shell
go build main.go -o bin/gomeasure
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

## Uninstall

### Debian (amd64)
1. To **uninstall** on debian distributions copy and paste the following command in your terminal
```bash
curl "https://raw.githubusercontent.com/lordvidex/gomeasure/master/scripts/uninstall.sh" | sh
```
## Usage
gomeasure --help
```

## DISCLAIMER:
Normally, the program reads all kinds of file and returns the number of lines in each of them. Therefore, it is up to the developer to filter out the files to be scanned with 
flags `-i` and `-I`.
