- **;** is optional
- 1 line 1 command
	- Use ; to separate command line
- **fmt** package
	- Built-in api, more powerful
	- fmt.Println() => Print 1 line
	- fmt.Printf(“%v”, var) => Print by String format
		- %v = value
		- %s = string
		- %d = number
	- Auto import using Go extension

**main.go**
```go
package main

import "fmt"

func main() {
	println("Hello, world!")
	fmt.Println("Hello World!!!")
	fmt.Printf("%v %v", "Hello", 999)
}

```

![[Pasted image 20250220151351.png]]

# Data Types and Variable Declaration

- **=** Assignation operator

```go
var <name> <type> = <value>
```

- Declared variables MUST be used
- Zero value is automatically added to a variable
- String to Int
	- strconv.Itoa(var)

**main.go**
```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var name string
	var age int
	println(name + " " + strconv.Itoa(age))
	fmt.Printf("%v %v\n", name, age)
}

```

![[Pasted image 20250220153522.png]]

- **:=** Declare and Assign (Short declaration)

```go
<name> := <value>
department := "ComSci"
```

- Auto detect data type
- Can be used only in function
- *Recommended!!!*

**main.go**
```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var name string
	var age int
	println(name + " " + strconv.Itoa(age))
	fmt.Printf("%v %v\n", name, age)
	department := "ComSci"
	fmt.Print(department)
}
```

![[Pasted image 20250220160821.png]]

- Basic data types
	- (u)int`<len>`
	- float`<len>`
	- bool
	- string
	- array
	- slice (~list)
	- struct (~class)

![[Pasted image 20250220162038.png]]

**main.go**
```go
package main

import (
	"fmt"
)

func main() {
	gpa := 3.56
	year := -4
	fmt.Printf("%T %T", gpa, year)

}
```

![[Pasted image 20250220162221.png]]

# Commons Errors

![[Pasted image 20250220162505.png]]

# Control Statement

- IF
	- ( ) are optional
	- { } are required

```go
if expr {
}
```

- Go extension will format source code automatically

**main.go**
```go
package main

func main() {
	isAdmin := false
	if isAdmin {
		println("Admin")
	} else {
		println("User")
	}
}
```

