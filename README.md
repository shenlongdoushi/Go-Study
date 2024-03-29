# Go Language Study

## Reference Website
- https://www.liwenzhou.com/posts/Go/golang-menu/
- https://golang.org/doc/install
- https://golang.org/dl/

## Installation

Here is the step about installation on ubuntu server:

### 1. Download installation pkg on server and Extract Pkg to /usr/local.

- Try the latest stable one
- Remove previous version if any.

```shell
sudo apt update -y
sudo apt upgrade -y
sudo apt install wget
sudo wget https://golang.org/dl/go1.17.1.linux-amd64.tar.gz
rm -rf /usr/local/go && tar -C /usr/local -xzf go1.17.1.linux-amd64.tar.gz
```

### 2. Add /usr/local/go/bin to the PATH environment variable and verify Go version.

- Set Go env to user

```shell
cat >> $HOME/.bashrc << EOF
export GOROOT=/usr/local/go
export PATH=$PATH:$GOROOT/bin
EOF
source $HOME/.bashrc
```

- Check Go Version

```
go version
# go version go1.17.1 linux/amd64
```

### 3. First Script - Hello World!

#### hello-world.go

```go
package main
import 'fmt'

func main(){
    fmt.Println("Hello, World!")
}
```

1. Execution without Compile
   - Easy to execute but slow to get output

```shell
# go run hello-world.go
Hello, World!
```

2.  Execution with Compile
    - Build/compile source to executable file.
    - Windows -> .exe file; Linux/Mac -> ELF 64-bit LSB executable file.
    - After build, the file can be run anywhere without setting go env.
    - The compiled file is much bigger than source file because it includes all the dependency.

```shell
# go build hello-world.go
# ls -l
-rwxr-xr-x 1 root root 1766310 Sep 26 18:17 hello-world
-rw-r--r-- 1 root root      72 Sep 26 05:32 hello-world.go
# ./hello-world
Hello, World!
# file hello-world
hello-world: ELF 64-bit LSB executable, x86-64, version 1 (SYSV), statically linked, Go BuildID=tdA533oGZ9DAVBU2Mc-3/QaURQoqz7nlFIKU1zJiP/a1n0XjGLqnuMixhqdNIG/EHa6VwU4ZBny6DT8K6V8, not stripped
```

3. Execution with Different Name Compile
   - variable `-o` modify the output name.

```shell
# go build -o hello-go hello-world.go
# ls -l
-rwxr-xr-x 1 root root 1766310 Sep 26 18:22 hello-go
-rw-r--r-- 1 root root      72 Sep 26 05:32 hello-world.go
# ./hello-go
Hello, World!
# file hello-go
hello-go: ELF 64-bit LSB executable, x86-64, version 1 (SYSV), statically linked, Go BuildID=tdA533oGZ9DAVBU2Mc-3/QaURQoqz7nlFIKU1zJiP/a1n0XjGLqnuMixhqdNIG/EHa6VwU4ZBny6DT8K6V8, not stripped
```

## Go Module

A module is a collection of Go packages stored in a file tree with a go.mod file at its root. The go.mod file defines the module’s module path, which is also the import path used for the root directory, and its dependency requirements, which are the other modules needed for a successful build. Each dependency requirement is written as a module path and a specific semantic version.

### Go Environment

Set up environment correctly is the most critical step.
The latest Go installation has default `GOPATH` value, but can modify if needed. (Usually, GOPATH will be set per workspace/project.)

```shell
# export GOPATH=/root/go
[optional if no need permanent]# source $HOME/.bashrc
# tree -fph -L 2 /root/go
/root/go
├── [drwxr-xr-x 4.0K]  /root/go/bin
│   ├── [-rwxr-xr-x  14M]  /root/go/bin/dlv
│   ├── [-rwxr-xr-x  14M]  /root/go/bin/dlv-dap
│   ├── [-rwxr-xr-x 1.7M]  /root/go/bin/echo2
│   ├── [-rwxr-xr-x 3.0M]  /root/go/bin/go-outline
│   ├── [-rwxr-xr-x 4.2M]  /root/go/bin/gopkgs
│   ├── [-rwxr-xr-x  22M]  /root/go/bin/gopls
│   ├── [-rwxr-xr-x 1.7M]  /root/go/bin/hello-world
│   └── [-rwxr-xr-x  11M]  /root/go/bin/staticcheck
└── [drwxr-xr-x 4.0K]  /root/go/pkg
    ├── [drwxr-xr-x 4.0K]  /root/go/pkg/mod
    └── [drwxr-xr-x 4.0K]  /root/go/pkg/sumdb

4 directories, 8 files
# go env
GO111MODULE=""
GOARCH="amd64"
GOBIN=""
GOCACHE="/root/.cache/go-build"
GOENV="/root/.config/go/env"
GOEXE=""
GOEXPERIMENT=""
GOFLAGS=""
GOHOSTARCH="amd64"
GOHOSTOS="linux"
GOINSECURE=""
GOMODCACHE="/root/go/pkg/mod"
GONOPROXY=""
GONOSUMDB=""
GOOS="linux"
GOPATH="/root/go"
GOPRIVATE=""
GOPROXY="https://proxy.golang.org,direct"
GOROOT="/usr/local/go"
GOSUMDB="sum.golang.org"
GOTMPDIR=""
GOTOOLDIR="/usr/local/go/pkg/tool/linux_amd64"
GOVCS=""
GOVERSION="go1.17.1"
GCCGO="gccgo"
AR="ar"
CC="gcc"
CXX="g++"
CGO_ENABLED="1"
GOMOD="/home/coder/project/go-study/projects/testfolder/go.mod"
CGO_CFLAGS="-g -O2"
CGO_CPPFLAGS=""
CGO_CXXFLAGS="-g -O2"
CGO_FFLAGS="-g -O2"
CGO_LDFLAGS="-g -O2"
PKG_CONFIG="pkg-config"
GOGCCFLAGS="-fPIC -m64 -pthread -fno-caret-diagnostics -Qunused-arguments -fmessage-length=0 -fdebug-prefix-map=/tmp/go-build437493066=/tmp/go-build -gno-record-gcc-switches"
```

#### Example

https://golang.org/doc/tutorial/getting-started

1. Create project folder and `go mod init` module

```shell
# mkdir testmod
# cd testmod/
# go mod init testmod
go: creating new go.mod: module testmod
# tree
.
└── go.mod

0 directories, 1 file
```

2. Create Go package files.

```go
// testmod.go
package main

import (
	"fmt"
    // This external package collects pithy sayings.
    // https://pkg.go.dev/rsc.io/quote
	"rsc.io/quote"
)

func main() {
    fmt.Println(quote.Hello())
    fmt.Println(quote.Go())
}
```

3. Use `go mod tidy` to manage module/packages dependency. The external dependency has to be added first before run it.

```shell
# ls
go.mod  testmod.go
# go run testmod.go
testmod.go:6:2: no required module provides package rsc.io/quote; to add it:
        go get rsc.io/quote
# go mod tidy
go: finding module for package rsc.io/quote
go: found rsc.io/quote in rsc.io/quote v1.5.2
# tree -fph
.
├── [-rw-r--r--  167]  ./go.mod
├── [-rw-r--r--  499]  ./go.sum
└── [-rw-r--r--  121]  ./testmod.go
# go run testmod.go
Hello, world.
Don't communicate by sharing memory, share memory by communicating.
```

4. After execute, the go mod files contents.

- `go.mod`: A module is defined by a UTF-8 encoded text file named go.mod in its root directory. The go.mod file is line-oriented. Each line holds a single directive, made up of a keyword followed by arguments.

```go
module testmod

go 1.17

require rsc.io/quote v1.5.2

require (
	golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c // indirect
	rsc.io/sampler v1.3.0 // indirect
)

```

- `go.sum`: This file lists down the checksum of direct and indirect dependency required along with the version. It is to be mentioned that the go.mod file is enough for a successful build. Then why go.sum file is needed?. The checksum present in go.sum file is used to validate the checksum of each of direct and indirect dependency to confirm that none of them has been modified.

```go
golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c h1:qgOY6WgZOaTkIIMiVjBQcw93ERBE4m30iBm00nkL0i8=
golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c/go.mod h1:NqM8EUOU14njkJ3fqMW+pc6Ldnwhi/IjpwHt7yyuwOQ=
rsc.io/quote v1.5.2 h1:w5fcysjrx7yqtD/aO+QwRjYZOKnaM9Uh2b40tElTs3Y=
rsc.io/quote v1.5.2/go.mod h1:LzX7hefJvL54yjefDEDHNONDjII0t9xZLPXsUe+TKr0=
rsc.io/sampler v1.3.0 h1:7uVkIFmeBqHfdjD+gZwtXXI+RODJ2Wc4O7MPEh/QiW4=
rsc.io/sampler v1.3.0/go.mod h1:T1hPZKmBbMNahiBKFy5HrXp6adAjACjK9JXDnKaTXpA=

```

### Import Local Package

While developing, local packages are often used temporarily for testing purpose instead of updating `GOPATH`. The key to use local package is directory hierarchy and package management.

1.  Create Root Directory and Put Local Directory under Root.

```shell
# tree -fph leetcode/
leetcode
├── [-rw-r--r--   39]  leetcode/go.mod
├── [-rw-r--r--  167]  leetcode/main.go
└── [drwxr-xr-x 4.0K]  leetcode/twoSum
    ├── [-rw-r--r--  196]  leetcode/twoSum/twoSum.go
    └── [-rw-r--r-- 1.1K]  leetcode/twoSum/twoSum_test.go

1 directory, 4 files
```

2. `go mod init` at root level and name package at child level

```go
// go.mod
// Best practice to name a git source for future operation
module github.com/nl/leetcode

go 1.17

```

```go
// twoSum.go
// package name twoSum, which is under module github.com/nl/leetcode/twoSum
// func name has to be first letter capital to be called external, e.g., TwoSum()
package twoSum

func TwoSum(nums []int, target int) ([]int)  {
	m := make(map[int]int)
	for k,v := range nums {
		if idx,ok:=m[target-v]; ok {
			return []int{idx, k}
		}
		m[v]=k
	}
	return nil
}
```

3. Import local packages into main function `main.go`. `go run` command has to run in the same directory with go file.

```go
// main.go
// Import local package to call functions
package main

import (
	"fmt"

	"github.com/nl/leetcode/twoSum"
)

func main() {
	fmt.Println("Import Local Package")
	fmt.Println(twoSum.TwoSum([]int{2,7,11,15},9))
}
```

```shell
# go env | grep GOROOT
GOROOT="/usr/local/go"
# pwd
/home/coder/project/go-study/leetcode
# go run main.go
Import Local Package
[0 1]
```

## Tips

### All the Packages and Variables defined in the script must be used, otherwise will be errors.

### Use GO official formatter - gofmt

```go
// just print and review the code
gofmt yourcode.go
// format and overwrite the current code
gofmt -w yourcode.go
```
