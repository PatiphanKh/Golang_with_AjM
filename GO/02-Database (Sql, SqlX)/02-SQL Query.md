- **db** is a connection to database server
- rows.Next()
- rows.Scan(...*)
	- Return *
	- Receive &
	- All columns  MUST be  handled

**main.go**
```go
	query := "select * from country"
	rows, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
	// defer means always perform when exit program
	defer rows.Close()
	for rows.Next() {
		idx := 0
		name := ""
		err = rows.Scan(&idx, &name)
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("%v %v\n", idx, name)
	}
	// db.Close() is not required
```

![[Pasted image 20250227211705.png]]

# Row to Struct

- Row should be transformed to Struct
- Create struct

![[Pasted image 20250227212802.png]]

**main.go**
```go
type Country struct {
	idx  int
	name string
}
```

**main.go**
```go
		var country Country
		err = rows.Scan(&country.idx, &country.name)
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("%v %v %v\n", country, country.idx, country.name)
```

![[Pasted image 20250227212842.png]]

# Re-Structure of Code
- Extract to functions
**main.go**
```go
package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var dsn = "landmark:landmark@csmsu@tcp(202.28.34.197:3306)/landmark"

type Country struct {
	idx  int
	name string
}

func getConnection() (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err.Error())
	}
	println("connection successful")

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(10)
	return db, nil
}

func GetCountries() ([]Country, error) {
	db, err := getConnection()
	if err != nil {
		panic(err.Error())
	}
	query := "select * from country"
	rows, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}

	countries := []Country{}

	for rows.Next() {
		var country Country
		err = rows.Scan(&country.idx, &country.name)
		if err != nil {
			panic(err.Error())
		}
		countries = append(countries, country)
	}
	return countries, nil
}

func main() {
	countries, err := GetCountries()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(countries)
}

```

![[Pasted image 20250227213859.png]]

# SQL Parameters

- ? is a parameter

**main.go**
```go
func GetCountryById(idx int) (Country, error) {
	db, err := getConnection()
	if err != nil {
		panic(err.Error())
	}
	query := "select * from country where idx = ?"
	row := db.QueryRow(query, idx)
	var country Country
	err = row.Scan(&country.idx, &country.name)
	if err != nil {
		panic(err.Error())
	}
	return country, nil

}

func main() {
	countries, err := GetCountries()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(countries)

	country, err := GetCountryById(1)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(country)

}
```

![[Pasted image 20250227214704.png]]





