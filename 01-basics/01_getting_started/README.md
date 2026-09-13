5 reasons to choose golang
1. build time
2. fast startup
3. performance and efficiency
4. concurrency model
5. static typing and compilation

## initialise the project
```
go mod init <module path>
```
module path can correspond to a repository you plan to publish your module to 
## packages
All our code must be belong to a package

The first statement in a file must be a package

## where to start the program? where is the entry point?
 main function is the entry point for go program
 A program can have only one main function

 ## go packages
 go program are orgnized in to packages
 go's standard library, provides different core packages for us to use
 "fmt" is one of them, which can be imported and used

 a packages is a containers of source files

 [go packages](https://pkg.go.dev/)

## execution
go run < file name >