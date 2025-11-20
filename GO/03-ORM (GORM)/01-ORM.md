# Object Relational Mapping (ORM)
- Mapping Object to Data in Database
- Entity(Model) => Table

![[Pasted image 20250306100701.png]]

# GORM
- Main website
	- https://gorm.io
## Get library
- [https://pkg.go.dev/gorm.io/gorm](https://pkg.go.dev/gorm.io/gorm)

![[Pasted image 20250306101134.png]]

- Gorm Library
```sh
go get gorm.io/gorm
```

- Dialactor (Driver)
	- [https://pkg.go.dev/gorm.io/driver/mysql](https://pkg.go.dev/gorm.io/driver/mysql)
```sh
go get gorm.io/driver/mysql
```

- dsn: MySql
	- user:pwd@tcp(ip:port)/dbname?collation=utf8mb4_unicode_ci&parseTime=true

- Database drivers
	- https://gorm.io/docs/connecting_to_the_database.html

# New Project
- go-gorm

![[Pasted image 20250306102307.png]]

**main.go**
```go
package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "landmark:landmark@csmsu@tcp(202.28.34.197:3306)/landmark?collation=utf8mb4_unicode_ci&parseTime=true"
	dialactor := mysql.Open(dsn)
	_, err := gorm.Open(dialactor)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connection successful")
}

```


![[Pasted image 20250306102236.png]]

# Configuration (Viper)
- Organize configuration of the project

```sh
go get github.com/spf13/viper 
```

**config.yaml**
```yaml
mysql:
  dsn: "landmark:landmark@csmsu@tcp(202.28.34.197:3306)/landmark?collation=utf8mb4_unicode_ci&parseTime=true"
```

**main.go**
```go
func main() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	fmt.Println(viper.Get("mysql.dsn"))
	dsn := viper.GetString("mysql.dsn")

	dialactor := mysql.Open(dsn)
	_, err = gorm.Open(dialactor)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connection successful")
}
```

![[Pasted image 20250306103816.png]]