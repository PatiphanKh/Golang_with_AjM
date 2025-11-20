# Installation

- Official website: [https://go.dev](https://go.dev)
![[Pasted image 20250220143758.png]]

![[Pasted image 20250220143835.png]]

- Software Spec: [https://go.dev/ref/spec](https://go.dev/ref/spec)
- Go package: [https://pkg.go.dev](https://pkg.go.dev)
- Libs suggestion: [https://go.libhunt.com](https://go.libhunt.com)

# VSCode Extensions

- GO
![[Pasted image 20250220144131.png]]

- Error lens
![[Pasted image 20250220144200.png]]

- Code Runner
![[Pasted image 20250220150318.png]]

# Go CLI (Command Line Interface)

- Check go version
```sh
go version
```

![[Pasted image 20250220144713.png]]

- Check environment variables
```sh
go env
```

![[Pasted image 20250220144830.png]]

# Hello World!!!
- Create a folder /**basic** for the project
		- go/basic
```sh
go mod init <project name>

go mod init go-basic
```

![[Pasted image 20250220145810.png]]

![[Pasted image 20250220145847.png]]

## Entrypoint
- Package main
- Func main()

**main.go**
```go
package main

func main() {
	println("Hello, world!")
}

```

- Built-in function
	- println()

```sh
go run main.go
go run .
```

![[Pasted image 20250220150237.png]]

- Code Runner
![[Pasted image 20250220150349.png]]