# Functions in Gorm

## Create
- value is an address of the object based on a model

**main.go**
```go
	country := model.Country{Name: "USA"}
	result := db.Create(&country)
	if result.Error != nil {
		panic(result.Error)
	}
	if result.RowsAffected > 0 {
		fmt.Println("Success")
	}
	countries := []model.Country{}
	result = db.Find(&countries)
	if result.Error != nil {
		panic(result.Error)
	}
	fmt.Println(countries)
```

![[Pasted image 20250306131644.png]]

## Delete

`Delete(value interface{}, conds ...interface{}) (tx *gorm.DB)`
- Primary key is required
- value is an address of the object based on a model

Using of a model and PK
- *Don't need &*
**main.go**
```go
	result := db.Delete(model.Country{}, 234)
	if result.Error != nil {
		panic(result.Error)
	}
	if result.RowsAffected > 0 {
		fmt.Println("Success")
	}

	countries := []model.Country{}
	result = db.Find(&countries)
	if result.Error != nil {
		panic(result.Error)
	}
	fmt.Println(countries)
```

Using of an object with PK
- *Need &*
**main.go**
```go
	country := model.Country{Idx: 232}
	result := db.Delete(&country)
	if result.Error != nil {
		panic(result.Error)
	}
	if result.RowsAffected > 0 {
		fmt.Println("Success")
	}

	countries := []model.Country{}
	result = db.Find(&countries)
	if result.Error != nil {
		panic(result.Error)
	}
	fmt.Println(countries)
```

## Delete with condition

- `where ("condition (with Model fields)", params)`

**main.go**
```go
	name := "%" + "New" + "%"
	result := db.Delete(model.Country{}, "Name like ?", name)
	if result.Error != nil {
		panic(result.Error)
	}
	if result.RowsAffected > 0 {
		fmt.Println("Success")
	}

	countries := []model.Country{}
	result = db.Find(&countries)
	if result.Error != nil {
		panic(result.Error)
	}
	fmt.Println(countries)
```

## Update 

- Update existing data 
	- Use func Save()
```go
	country := model.Country{}
	db.Model(model.Country{}).Where("Idx = ?", 1).Take(&country)
	fmt.Println(country)
	country.Name = "Thailand"
	db.Save(&country)

	db.Model(model.Country{}).Where("Idx = ?", 1).Take(&country)
	fmt.Println(country)
```

- Update specific fields
	- Use func Where() and Update()
```go
	country := model.Country{}
	db.Model(model.Country{}).Where("Idx = ?", 1).Update("Name", "ประเทสไทย")

	db.Model(model.Country{}).Where("Idx = ?", 1).Take(&country)
	fmt.Println(country)
```

