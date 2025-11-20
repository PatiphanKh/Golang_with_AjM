# Exec()
- Return Result
- LastInsertId
- Row Affected

## Modify Country Struct
- SetName(name)

**main.go**
```go
type Country struct {
	idx  int
	name string
}

func (c *Country) SetName(name string) {
	c.name = name
}
```

## Create data
- func AddCountry(country)

**main.go**
```go

func AddCountry(country Country) (int64, int64, error) {
	db, err := getConnection()
	if err != nil {
		return -1, -1, err
	}
	query := "insert into country(name) values(?)"
	result, err := db.Exec(query, country.name)
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
	country := Country{}
	country.SetName("New Country")
	affected, idx, err := AddCountry(country)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(affected, idx)
	}
```

![[Pasted image 20250227220816.png]]

## Update Data


**main.go**
```go
func UpdateCountry(country Country) (int64, error) {
	db, err := getConnection()
	if err != nil {
		return -1, err
	}
	query := "update country set name = ? where idx = ?"
	result, err := db.Exec(query, country.name, country.idx)
	if err != nil {
		return -1, err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return -1, err
	}

	return affected, nil
}

func main() {
	country := Country{idx: 175, name: "Super New Country"}
	affected, err := UpdateCountry(country)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(affected)

}

```

![[Pasted image 20250227222417.png]]

## Delete Data

**main.go**
```go
func DeleteCountry(idx int) (int64, error) {
	db, err := getConnection()
	if err != nil {
		return -1, err
	}
	query := "delete from country where idx = ?"
	result, err := db.Exec(query, idx)
	if err != nil {
		return -1, err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return -1, err
	}

	return affected, nil
}

func main() {
	affected, err := DeleteCountry(175)
	if err != nil {
		panic(err.Error())
	}
	println(affected)
}

```

![[Pasted image 20250227222740.png]]






