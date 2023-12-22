package main

import (
	"fmt"
	"os"

	"github.com/sohail-9098/ms-user-auth/router"
)

func main() {
	fmt.Println(os.Getwd())
	router.StartApplication()
}
