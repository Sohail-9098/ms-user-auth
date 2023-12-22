package utils

import (
	"time"

	"github.com/sohail-9098/ms-user-auth/router"
)

func StartApp() {
	go router.StartApplication()
	time.Sleep(time.Second * 2)
}
