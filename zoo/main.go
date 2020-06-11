package main

import (
	"fmt"

	"github.com/kengaku-eiichi/go-trainning/zoo/animals"
)

func main() {
	fmt.Println(AppName())
	fmt.Println(animals.ElephantFeed())
	fmt.Println(animals.MonkeyFeed())
	fmt.Println(animals.RabbitFeed())
}
