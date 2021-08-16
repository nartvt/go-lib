.env file
```dotenv
    DB_DRIVER=postgres
    DB_HOST=127.0.0.1
    DB_PORT=5432
    DB_NAME=test
    DB_USERNAME=postgres
    DB_PASSWORD=root
```

```postgresql
create database test;
create table user
(
    id     int,
    number int
)

```
````go

func initPostgresql() {
    config := loadConfigFileName(".", ".env")
    fmt.Println(config)
    ConnectInstance = NewInstance(config)
    if ConnectInstance.GetConnectInstance().(*sql.DB) != nil {
        fmt.Println("Connected")
    } else {
    	fmt.Println("Not Connect")
    }
}

type User struct {
    Id     int8
    Number uint64
}

func main() {
    initPostgresql()
    sqlString := "select id, number from test_table"
    sqlQuery := ConnectInstance.GetConnectInstance().(*sql.DB)
    row := sqlQuery.QueryRow(sqlString)
    users := make([]User, 2)
    if row != nil {
        var user User
        if err := row.Scan(&user.Id, &user.Number); err != nil {
            fmt.Println(err.Error())
        } else {
            users = append(users, user)
        }    
    }
    fmt.Println(users)
}
````