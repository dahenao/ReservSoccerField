package main

import (
	"database/sql"
	"fmt"
	"github.com/dahenao/ReservSoccerField/internal/client"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	fmt.Print("holirtas")

	dbConnection, err := sql.Open("mysql", "root@tcp(localhost:3306)/reservFieldBD?parseTime=true")
	if err != nil {
		panic(err)
	}

	repository := client.NewRepoCliente(*dbConnection)

	fmt.Print(repository.GetById(1))
}
