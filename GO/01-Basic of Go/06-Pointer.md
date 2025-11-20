# What is a Pointer?
- Reference into memory
# Symbol

- `&var` = Address of var
- `*var` = Value pointed by address var
- `*type` = Pointer of type (store an address)
![[Pasted image 20250220200536.png]]

**main.go**
```go
package main

import "fmt"

func main() {
	var x int = 10
	// Value of x
	fmt.Printf("%v\n", x)

	// Address of x
	fmt.Printf("%v\n", &x)
}
```

![[Pasted image 20250220201626.png]]

```go
package main

import "fmt"

func main() {
	var x int = 10
	// Value of x
	fmt.Printf("%v\n", x)

	// Address of x
	fmt.Printf("%v\n", &x)

	var y *int = &x
	// Value of y (address)
	fmt.Printf("%v\n", y)

	// Value of pointer y
	fmt.Printf("%v\n", *y)

}
```

![[Pasted image 20250220201855.png]]

# Pass by reference of primitive type
- Primitive type can be referenced

```go
package main

import "fmt"

func main() {
	var a int = 0
	fmt.Println(a)
	fmt.Println(&a)
	change(&a)
	fmt.Println(a)
}

func change(x *int) {
	fmt.Println(x)
	*x = 10
}

```

![[Pasted image 20250220202312.png]]




