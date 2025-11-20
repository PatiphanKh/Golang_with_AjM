# Array

- Declaration always with zero value
- [len], [...] type{}

**main.go**
```go
package main

import "fmt"

func main() {
	var array1 = [3]int{}
	array2 := [3]string{}
	var array3 = [...]float32{}
	var array4 = [3]bool{true, true}
	fmt.Printf("%v\n", array1)
	fmt.Printf("%v\n", array2)
	fmt.Printf("%v\n", array3)
	fmt.Printf("%v\n", array4)
}

```

![[Pasted image 20250220165715.png]]

- Array Indexing
	- [index]
	- [start:end]
		- End is not included

**main.go**
```go
package main

import "fmt"

func main() {
	var array1 = [3]int{}
	array2 := [3]string{}
	var array3 = [...]float32{}
	var array4 = [3]bool{true, true}
	fmt.Printf("%v\n", array1)
	fmt.Printf("%v\n", array2)
	fmt.Printf("%v\n", array3)
	fmt.Printf("%v\n", array4)
	fmt.Printf("=======\n")
	fmt.Printf("%v\n", array4[0])
	fmt.Printf("%v\n", array4[1:2])
	fmt.Printf("%v\n", array4[1:])
	fmt.Printf("%v\n", array4[:2])
}
```

![[Pasted image 20250220171833.png]]

# Slice (List)
- **[ ]**
- len(slice)
```go
	slice1 := []string{}
	slice1 = append(slice1, "aj.m")
	slice2 := append(slice1, "potchara")
	slice2 = append(slice2, "csmsu")

	fmt.Printf("%v\n", slice1[0])
	fmt.Printf("%v\n", slice2[2])
	fmt.Printf("%v\n", len(slice2))
```

![[Pasted image 20250220172935.png]]

# Map
- Name := map[key-type]value-type
- Value, exist := map[key]

**main.go**
```go
func main() {
	fullname := map[string]string{}
	fullname["65011290001"] = "Potchara"
	fullname["65011290002"] = "Aj.M"
	fmt.Printf("%v\n", fullname["65011290001"])
	value, exist := fullname["65011290002"]
	fmt.Printf("%v %v\n", value, exist)
	value, exist = fullname["65011290003"]
	fmt.Printf("%v %v\n", value, exist)
}
```

![[Pasted image 20250220191819.png]]

# Control Statement (Loop)

- For
	- *For i*
	- *For range* return (Index and value)

**main.go**
```go
func main() {
	fullnames := [5]string{
		"aj.M", "aj.A"}
	for i := 0; i < len(fullnames); i++ {
		fmt.Println(fullnames[i])
	}
	for index, value := range fullnames {
		fmt.Printf("%v. %v\n", index, value)
	}
}
```

![[Pasted image 20250220192403.png]]

