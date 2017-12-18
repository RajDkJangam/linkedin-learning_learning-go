# Compile a complete application with `go install`

- First create the proper file structure.

```
md gowork
cd gowork
md src
md bin
md pkg
```

- Set your GOPATH variable to point to the application directory

```
set GOPATH=C:\gowork
go env
cd src
md palindrome
md palindrome
```

- Put palindrome.go in `gowork\src\palindrome\`.

Here, we are matching source code name to directory name, but this is not required. What's important is for the file to have a package declaration `main`, and a function `main()` with no arguments. That's what defines it as the startup application.

> Packages should be all lowercase and a single world. Don't mix the cases and do not include spaces. See https://blog.golang.org/package-names.

- Create your "gopath" application.

`go install`

- Check for your executable in the \bin directory

`cd ..\..\bin`
`dir`

- Generate executables for other platforms by adding another environment variable

`cd ..\src\palindrome`

- Generate an executable that will work on Mac and other OS's with the same ancestors
> You may require admin privledges because this will create a directory

`set GOOS=darwin`
`go install`

- Check for your executable in the \bin directory

`cd ..\..\bin`
`dir`
`cd darwin_amd64`
`dir`

There are third-party tools that handle compilation in bulk.

> For more instruction, see https://golang.org/doc/code.html