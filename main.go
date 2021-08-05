package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/shahzaibaziz/go-migrate/cmd"
)

func main() {
	cmd.Execute()
}
