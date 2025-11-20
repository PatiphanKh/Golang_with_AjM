	# Create Landmark model

**model/landmark.go**
```go
package model

type Landmark struct {
	Idx     int    `gorm:"column:idx;primary_key;AUTO_INCREMENT"`
	Name    string `gorm:"column:name;NOT NULL"`
	Country int    `gorm:"column:country;NOT NULL"`
	Detail  string `gorm:"column:detail;NOT NULL"`
	Url     string `gorm:"column:url;NOT NULL"`
}

func (m *Landmark) TableName() string {
	return "landmark"
}

```


![[Pasted image 20250306111221.png]]

```go
package model

type Landmark struct {
	Idx         int     `gorm:"column:idx;primary_key;AUTO_INCREMENT"`
	Name        string  `gorm:"column:name;NOT NULL"`
	Country     int     `gorm:"column:country;NOT NULL"`
	Detail      string  `gorm:"column:detail;NOT NULL"`
	Url         string  `gorm:"column:url;NOT NULL"`
	CountryData Country `gorm:"foreignKey:Country;references:Idx"`
}

func (m *Landmark) TableName() string {
	return "landmark"
}

```
- Create Field *CountryData* (type Country)
	- ForeignKey is the field *Country*
	- Reference to *Idx* in Country

**main.go**
```go
	landmarks := []model.Landmark{}
	result := db.Find(&landmarks)
	if result.Error != nil {
		panic(result.Error)
	}
	fmt.Println(landmarks)
```

- No Country
![[Pasted image 20250306111951.png]]

# Preload Data
- Preload data to the specific fields (*CountryData*)
- Separate query
- `Preload("FieldName")`

**main.go**
```go
result := db.Preload("CountryData").Find(&landmarks)
```

![[Pasted image 20250306113316.png]]

# Joins Data
- Joins SQL
- `Joins("FieldName")`

**main.go**
```go
result := db.Joins("CountryData").Find(&landmarks)
```

![[Pasted image 20250306113316.png]]

# Preload vs Joins

When considering GORM's `Preload` versus `Join` for loading related data, performance is a key factor. Here's a breakdown of their differences and performance implications:

**Understanding the Concepts**

- **Preload (Eager Loading):**
    - GORM executes separate queries to fetch the related data.   
    - It retrieves the main data first, then issues additional queries to load associated records.
    - This results in "N+1 query problem" situations if not used carefully.
- **Join:**
    - GORM uses SQL `JOIN` clauses to retrieve related data in a single query.
    - This can reduce the number of database round trips.

**Performance Considerations**

- **Number of Queries:**
    - `Join` generally results in fewer queries, which can significantly improve performance, especially when dealing with large datasets.
    - `Preload` can lead to many queries, particularly with nested associations, which can degrade performance.
- **Data Volume:**
    - `Join` can retrieve a large amount of data in a single query, which might be inefficient if you only need a subset of the related data.
    - `Preload` allows you to fetch only the necessary related data in separate queries.
- **Data Duplication:**
    - `Join` can result in data duplication in the result set, which GORM then processes to reconstruct the object relationships.
    - `Preload` avoids data duplication by fetching related data separately.
- **Complexity:**
    - Complex `JOIN` queries can be harder to write and maintain.
    - `Preload` often simplifies the code for loading related data.

**When to Use Which**

- **Use `Join` when:**
    - You need to filter or sort data across multiple tables.   
    - You want to minimize the number of database round trips.
    - The related data volume is relatively small.
- **Use `Preload` when:**
    - You need to load related data for multiple objects.
    - You want to avoid data duplication.
    - You need to load only specific related data.
    - When you are worried about very large join result sets.
- **Considerations:**
    - The N+1 query problem is a significant concern with `Preload`. Ensure you understand how GORM handles preloading to avoid this issue.
    - Database indexing plays a crucial role in the performance of both `Join` and `Preload`.

**In summary:**

- `Join` tends to be faster for simple relationships and when you need to filter or sort across tables.
- `Preload` gives you more control over the data being loaded and can prevent data duplication, but requires careful attention to avoid the N+1 query problem.

It's often beneficial to benchmark both approaches to determine the optimal solution for your specific use case.

# Query
- [https://gorm.io/docs/query.html](https://gorm.io/docs/query.html)
	- *Where()* => condition with column name in database
	- *Order()*
	- *Limit()*
	- *Distinct()*
	- Find() => Get records
	- Take() => Get One Record
	- First()
	- Last()


### Where()
- Examples
```go
// Find the first user with name jinzhu
db.Where("name = ?", "jinzhu").First(&user)
// Find the first user with name jinzhu and age 20
db.Where(&User{Name: "jinzhu", Age: 20}).First(&user)
// Find the first user with name jinzhu and age not equal to 20
db.Where("name = ?", "jinzhu").Where("age <> ?", "20").First(&user)
```

**main.go**
```go
	countries := []model.Country{}
	db.Where("Name like ? or Name = ?", "%ไทย%", "อิตาลี").Find(&countries)
	fmt.Println(countries)
```

![[Pasted image 20250306140605.png]]

### Take()
- Take ONE Record to Object or Slice
**main.go**
```go
	countries := []model.Country{}
	db.Where("Name like ? or Name = ?", "%ไทย%", "อิตาลี").Take(&countries)
	fmt.Println(countries)

	country := model.Country{}
	db.Where("Name like ? or Name = ?", "%ไทย%", "อิตาลี").Take(&country)
	fmt.Println(country)
```

![[Pasted image 20250306140925.png]]

### First() and Last()
- Get ONE Record to an Object

**main.go**
```go
	countries := []model.Country{}
	db.Where("Name like ? or Name = ?", "%ไทย%", "อิตาลี").Find(&countries)
	fmt.Println(countries)

	country := model.Country{}
	db.Where("Name like ? or Name = ?", "%ไทย%", "อิตาลี").First(&country)
	fmt.Println(country)

	db.Where("Name like ? or Name = ?", "%ไทย%", "อิตาลี").Last(&country)
	fmt.Println(country)
```

### Order and Limit
- Get ONE Record to an Object

**main.go**
```go
	countries := []model.Country{}
	db.Order("Name desc").Limit(2).Find(&countries)
	fmt.Println(countries)
```

![[Pasted image 20250306143246.png]]

### Distinct()
- user Model() to specific data table

**main.go**
```go

	countries := []model.Country{}
	db.Distinct("Idx").Find(&countries)
	fmt.Println(countries)

	type unique struct {
		Idx int
	}
	
	uniques := []unique{}
	db.Model(model.Country{}).Distinct("Idx").Find(&uniques)
	fmt.Println(uniques)
```

![[Pasted image 20250306143954.png]]