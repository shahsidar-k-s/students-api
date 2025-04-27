package main

import (
	"fmt"

	"github.com/shahsidar-k-s/students-api/internal/config"
)

func main() {
	//loading the config files
	configs := config.MustLoad()
	fmt.Println("configs==>", configs)
}
