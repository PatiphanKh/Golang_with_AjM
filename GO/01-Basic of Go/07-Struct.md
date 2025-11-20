# Struct

- **No Class in Go**
- Struct is a type
	- Value type
	- Only fields
	- No methods
	- No var

**main.go**
```go
package main

import "fmt"

type Person struct {
	name    string
	age     int
	address Address
}

type Address struct {
	houseNo  int
	city     string
	province string
	postcode string
}


func main() {
	ajm := Person{
		name: "Potchara",
		age:  22,
		address: Address{
			houseNo:  123,
			city:     "Meaung",
			province: "Maha Sarakham",
			postcode: "44000",
		},
	}
	fmt.Println(ajm)
}

```

![[Pasted image 20250220203602.png]]


# Package Scope (again)

![[Pasted image 20250220203948.png]]

**package/demostruct/demo.go**
```go
package demostruct

type Student struct {
	name string
	gpa  float32
}

```


**main.go**
```go
package main

import (
	"go-basic/package/demostruct"
)

func main() {
	s1 := demostruct.Student{}
	s1.name = "Potchara"
	s1.gpa = 3.56
}
```

![[Pasted image 20250220204122.png]]

# Fix Error

**package/demostruct/demo.go**
```go
package demostruct

type Student struct {
	Name string
	Gpa  float32
}

```

**main.go**
```go
package main

import (
	"go-basic/package/demostruct"
)

func main() {
	s1 := demostruct.Student{}
	s1.Name = "Potchara"
	s1.Gpa = 3.56
}
```

# Encapsulation
- Setter & Getter

## Receiver function
- Method (member) of struct

![[Pasted image 20250220204654.png]]

**package/demostruct/demo.go**
```go
package demostruct

type Student struct {
	name string
	gpa  float32
}

func (s *Student) SetName(name string) {
	s.name = name
}

func (s *Student) SetGpa(gpa float32) {
	s.gpa = gpa
}

func (s *Student) GetName() string {
	return s.name
}

func (s *Student) GetGpa() float32 {
	return s.gpa
}
```

**main.go**
```go
package main

import (
	"go-basic/package/demostruct"
)

func main() {
	s1 := demostruct.Student{}
	s1.SetName("Potchara")
	s1.SetGpa(3.56)
	fmt.Println(s1)
	fmt.Println(s1.GetName())
	fmt.Println(s1.GetGpa())
}
```

![[Pasted image 20250220205024.png]]



