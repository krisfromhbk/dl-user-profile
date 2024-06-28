package main

import (
	"dl-user-profile/internal/server"
	"fmt"
)

func main() {
	srv := server.New()
	err := srv.Start()
	if err != nil {
		fmt.Println(err)
	}
}
