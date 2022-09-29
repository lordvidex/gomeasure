## gomeasure
gomeasure is a CLI tool that measures the lines of code and the count of files in a directory or project

<!--TODO : Create a table for all commands -->
## Installation
### Build yourself
This project is open source. You can simply clone this repo and build using:
```
go build main.go -o bin/gomeasure
```

### Download releases
Checkout the most updated releases [here](https://github.com/lordvidex/gomeasure/releases/) and download the binary that corresponds to your operating system

### MacOS (homebrew)
1. Using homebrew
```bash
$ brew tap lordvidex/lordvidex
$ brew install lordvidex/lordvidex/gomeasure
$ gomeasure --version # to confirm

```

### Debain (apt)
```shell
$ curl https://github.com/lordvidex/gomeasure/install.sh
$ sudo sh install.sh
```

## Usage
```
gomeasure --help
```

## DISCLAIMER:
Normally, the program reads all kinds of file and returns the number of lines in each of them. Therefore, it is up to the developer to filter out the files to be scanned with 
flags `-i` and `-I`.
