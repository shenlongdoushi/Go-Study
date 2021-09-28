# Go Language Study

## Installation

https://golang.org/doc/install

https://golang.org/dl/

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

## Tips

### All the Packages and Variables defined in the script must be used, otherwise will be errors.

### Use GO official formatter - gofmt

```go
// just print and review the code
gofmt yourcode.go
// format and overwrite the current code
gofmt -w yourcode.go
```
