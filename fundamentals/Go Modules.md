# Go Modules

`go mod init github.com/aditya109/project-test`

This will create `go.mod` and `go.sum` file inside the `project-test` directory

To update dependencies, type `go mod tidy`.

# Installing Go in Windows

1. Install go from its official site.
2. Under Environment Variables > User Variables, 
   - in ` Path`, add your installation directory `bin` path. `C:\Program Files\Go\bin`.
   - add `GOPATH` as another variable, provide the value of the your go-workspace. Mine is `D:\Work\go-workspace`. 
3. Close all open `Powershell` windows and code editors.
4. Now try running `go version`.