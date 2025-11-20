
- Enhanced original sql function
```sh
go get github.com/jmoiron/sqlx
```
- Struct with Exposed fields
- **Recommended**!
- **sqlx**.Open(...,...)


- using of struct with exposed fields
```go
// SqlX
type CountryX struct {
	// Exposed fields
	Idx  int
	Name string
}
```

- use sqlx connection
```go
// SqlX
func getConnectionX() (*sqlx.DB, error) {
	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		panic(err.Error())
	}
	println("connection successful")

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(10)
	return db, nil
}
```


### .Select()
- Select Multiple row
- Return pointer to array of struct
```go
// SqlX
func GetCountriesX() ([]CountryX, error) {
	db, err := getConnectionX()
	if err != nil {
		return nil, err
	}
	query := "select * from country"
	countries := []CountryX{}
	err = db.Select(&countries, query)
	if err != nil {
		return nil, err
	}
	return countries, nil
}


func main() {

	countries, err := GetCountriesX()
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("%v\n", countries[0])

}
```


![[Pasted image 20250227225601.png]]

### .Get()
- Select One row
- Return pointer to array of struct

```go
// SqlX
func GetCountryXById(idx int) (CountryX, error) {
	db, err := getConnectionX()
	if err != nil {
		panic(err.Error())
	}
	query := "select * from country where idx = ?"
	var country CountryX
	err = db.Get(&country, query, idx)
	if err != nil {
		return country, err
	}
	return country, nil
}

func main() {

	country, err := GetCountryXById(1)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("%v\n", country)

}

```

### NamedExec()
- Pass struct and auto mapping

```go
// SqlX
func AddCountryX(country CountryX) (int64, int64, error) {
	db, err := getConnectionX()
	if err != nil {
		return -1, -1, err
	}
	query := "insert into country(name) values(:name)"
	result, err := db.NamedExec(query, country)
	if err != nil {
		return -1, -1, err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return -1, -1, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return -1, -1, err
	}

	return affected, id, nil
}
func main() {

	country := CountryX{Name: "Asgard"}
	affected, idx, err := AddCountryX(country)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(affected, idx)

}
```

![[Pasted image 20250227231431.png]]

