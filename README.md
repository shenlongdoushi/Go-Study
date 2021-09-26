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

```shell
export PATH=$PATH:/usr/local/go/bin
go version
# go version go1.17.1 linux/amd64
```

### 3. First Script - Hello World!

```go
package main
import 'fmt'

func main(){
    fmt.Println("Hello, World!")
}

// go run hello-world.go
// Hello, World!
```
