# Database Server

- MySQL Database
	- Server: 202.28.34.197
	- Port: 3306
	- Username: landmark
	- Password: landmark@csmsu
	- Database name: landmark

# SQL Library
- Package: [https://pkg.go.dev](https://pkg.go.dev)

![[Pasted image 20250227203139.png]]

![[Pasted image 20250227203218.png]]


# Create new project

![[Pasted image 20250227204222.png]]

 - Initial project
```sh
go mod init go-sql
```

![[Pasted image 20250227204255.png]]

**main.go**
![[Pasted image 20250227204551.png]]

- Sql Open Connection
	- Driver is required
![[Pasted image 20250227204750.png]]

# SQL Drivers

- Driver [https://golang.org/s/sqldrivers](https://golang.org/s/sqldrivers)
![[Pasted image 20250227203403.png]]

# Get library (MySQL driver)
- go get [git repo]

```sh
go get github.com/go-sql-driver/mysql
```

- sql.Open("driver", "DSN")
- Mysql DSN
	- user:pass@tcp(id:port)/dbname

**main.go**
```go
package main

import (
	"database/sql"
	"time"
)

func main() {
	db, err := sql.Open("mysql", "landmark:landmark@csmsu@tcp(202.28.34.197:3306)/landmark")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	println("connection successful")

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(10)
}

```

**Go Guard proposal**
```go
	if err != nil {
		panic(err.Error())
	}
```

![[Pasted image 20250227205455.png]]

- Manual MySQL Driver
```go
import (
	"database/sql"
	"time"
	_ "github.com/go-sql-driver/mysql"
)
```

![[Pasted image 20250227210900.png]]

