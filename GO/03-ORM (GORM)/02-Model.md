- https://gorm.io/docs/models.html#Fields-Tags


![[Pasted image 20250306104206.png]]

# Struct to Model
- Database Table
- `gorm field tags
- Receiver func TableName() 
	- table name on database

**model/country.go**
```go
package model

type Country struct {
	Idx  int    `gorm:"column:idx;primary_key;AUTO_INCREMENT"`
	Name string `gorm:"column:name;NOT NULL"`
}

func (m *Country) TableName() string {
	return "country"
}

```


# SQL to Gorm
- https://sql2gorm.mccode.info

![[Pasted image 20250306105710.png]]

## DDL Script

- MySql command
```sql
SHOW CREATE TABLE xxxxxx
```

- PhpMyAdmin 
![[Pasted image 20250306105954.png]]

![[Pasted image 20250306110042.png]]

## Test
```go
db.Find(&model)
```

**main.go**
```go
	dialactor := mysql.Open(dsn)
	db, err := gorm.Open(dialactor) // Declaration of db connection
	if err != nil {
		panic(err)
	}
	fmt.Println("Connection successful")

	countries := []model.Country{}
	result := db.Find(&countries)
	if result.Error != nil {
		panic(result.Error)
	}
	fmt.Println(countries)
```

![[Pasted image 20250306110350.png]]


