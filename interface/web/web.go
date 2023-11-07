package web

import (
	"fmt"
)

func New() error {
	fmt.Println("Create, Share, Move with Motion.")
	NewServer()
	err := server.GetGinEngine().Run(server.addr)
	return err
}
