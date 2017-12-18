# Go command-line toods
There are three (3) commands in the Go\bin directory

## `godoc`
Extract and display documentation for packages
example: `godoc fmt`

## `gofmt`
Automatically formats a file properly with indentation, but this does not update the file you run it on, it only outputs the formatted version.
example: `gofmt badformatting.go`

### write the changes to file
Include the `-w` flag
`gofmt -w palindrome.go`

## `go`
Incorporates > 12 subcommands. Type go and you will see those commands. Mainly, we will work with:
- `run`: compiles and runs a single source code file
- `build`: compile a file into an executable
- `install`: compiles and packages a more complex Go application that has a complete package structure

### Create a temporary executable file
`go run palindrome.go`
This compiles the application and creates a temporary executable on your disk, and then runs the application. This new executable is not in the same directory. You can see this with `dir` or `ls`

### Create a persistent executable file
`go build palindrome.go`
You can see this with `dir` or `ls`. Then, type the name of the executable for it to run. This executable is operating system-specific. To use it on multiple platforms, you must compile them on multiple platforms.

### Create a complete application
`go install` is recommended for creating more complete applications, but requires a more complex directory structure and an enviornment variable.