# Function in GO

- Declaration
```go
func <name>(<params>) <type>{

}
```

**main.go**
```go
func main() {
	x := add(10, 5)
	fmt.Println(x)
}

func add(x int, y int) int {
	return x + y
}
```

![[Pasted image 20250220193434.png]]

# Variadic parameters
- using ... to extract parameter to array

```go
func <name>(name …type) <type>{

}
```

**main.go**
```go
func main() {
	x := variadic("h", "e", "l", "l", "o")
	fmt.Println(x)
}

func variadic(x ...string) string {
	str := ""
	fmt.Println(len(x))
	for _, v := range x {
		str += v
	}
	return str
}

```

![[Pasted image 20250220194152.png]]

# Tuple
- Return multiple values

**main.go**
```go
func main() {
	x := add(10, 5)
	fmt.Println(x)
	x, y := cal(10, 5)
	fmt.Println(x, y)
}

func add(x int, y int) int {
	return x + y
}

func cal(x int, y int) (int, int) {
	return x + y, x - y
}
```

![[Pasted image 20250220194425.png]]

# Function Signature (Type)
- Function as a variable

**main.go**
```go
package main

func main() {
	// Declare function as a variable
	// Signature type = (int, int, int) float64
	// f is a fuction
	f := func(x int, y int, z int) float64 {
		return (float64(x+y+z) / 3.0)
	}
	result := f(1, 2, 3)
	println(result)
	resultSomething := something(f)
	println(resultSomething)

}

// Parameter as a function
func something(fx func(int, int, int) float64) float64 {
	return fx(5, 6, 7)
}

```

![[Pasted image 20250220195342.png]]

# Package
- Scope
	- Variable, Functions
		- var [varName] = ……..
	- Cannot use := in package
	- Encapsulation => Expose
		- Capital Letter

![[Pasted image 20250220195759.png]]

**scope/scope.go**
```go
package scope

var id = 0

func getId() int {
	return id
}
```


**main.go**
```go
package main

func main() {
	scope.getId()
}
```

![[Pasted image 20250220200124.png]]


## Fix Encapsulation (Capital letter)

**scope/scope.go**
```go
package scope

var id = 0

func GetId() int {
	return id
}
```

**main.go**
```go
package main

import "go-basic/scope"

func main() {
	scope.GetId()
}
```

