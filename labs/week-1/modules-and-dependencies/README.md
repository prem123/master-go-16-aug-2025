## Modules

- [Using Go modules](https://go.dev/blog/using-go-modules)
- [New module changes](https://go.dev/blog/go116-module-changes)

Go modules are an essential part of dependency management in Go, introduced to handle versioning and distribution of code in a consistent and reproducible way. 

A module is a collection of Go packages stored in a file tree with a `go.mod` file at its root. The `go.mod` file defines the module’s module path, which is also the `import` path used for the root directory, and its dependency requirements, which are the other modules needed for a successful build. Each dependency requirement is written as a module path and a specific semantic version.

Before Go 1.13, Go forced us to keep the source inside `$GOPATH/src` folder. With `Go modules` this is not required any more and we can keep the source folder anywhere in the file system.

Let's create a folder named as `module-examples` in your workspace and create a file `main.go` inside it.

`module-examples\main.go`:

```go
package main

import "fmt"
import "utils"

func main() {
  result := utils.Sum(10, 20)
  fmt.Println("Result is: ", result)
}
```

`module-examples\utils\myutils.go`:

```go
func Sum(a int, b int) int {
  return a + b
}
```

Compile and run the program from the command line and check the error message:

```shell
$ go run main.go

main.go:4:8: cannot find package "utils" in any of:
  /usr/lib64/golang/src/utils (from $GOROOT)
  /home/ashish/go/src/utils (from $GOPATH)
```

To fix, we need to create a go module inside our `module-examples` folder. 

```shell
go mod init learning-modules
```

Make the following changes to `main.go`:

```go
package main

import (
  "fmt"
  "learning-modules/utils"
)

func main() {
  result := utils.Sum(10, 20)
  fmt.Println("Result is: ", result)
}
```

Run again:

```shell
$ go run main.go

Result is:  30
```

## Module name

When initializing module name using `go mod init <MODULE-NAME>`. The module name can be anything but generally it is named after the place where the source is kept. You will see references of `github.com` at several places and it is a standard practice but not mandatory.
For e.g. the module can be named as:

```shell
go mod init github.com/the-code-camp/go-lang-essentials
```

## External dependencies

When your code uses external packages, those packages (distributed as modules) become dependencies. Using Go modules you can add new dependencies and manage the existing ones. Over time, you may need to upgrade them or replace them. Go provides dependency management tools that help you keep your Go applications secure as you incorporate external dependencies.

Let’s add an external library as a dependency. For this example, we'll use the rsc.io/quote package, which provides simple quote-generating functions.


```shell
$ go get -u rsc.io/quote

# Verify the dependency has been added in your go.mod file.
$ cat go.mod
```

```go
package main

import (
    "fmt"
    "rsc.io/quote"
)

func main() {
    fmt.Println("Welcome to Go modules!")
    fmt.Println(quote.Hello())
}
```

```shell
$ go run main.go

# You should see output similar to:
Welcome to Go modules!
Hello, world.

```
## Managing Dependency Versions

Go modules allow you to specify exact versions of dependencies.

Let’s downgrade the `rsc.io/quote` package to an older version, `e.g., v1.4.0`:

```bash

go get rsc.io/quote@v1.4.0
```
The go.mod file will now reflect this version:

```go

module github.com/the-code-camp/go-lang-essentials

go 1.20

require rsc.io/quote v1.4.0
```

If you attempt to run your code now:

```bash

go run main.go
```
The output may change depending on the version of quote that you're using:

```bash

Welcome to Go modules!
Hello, world.
```

## Tidying Up Your Module
After adding and removing dependencies, you can use go mod tidy to clean up your go.mod file, removing any unused dependencies and ensuring all necessary ones are included.

Run:

```bash
go mod tidy
```

This command will remove any dependencies from go.mod and go.sum that are no longer used in your project.