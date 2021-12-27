package main

import (
	"cart/helper"
	"fmt"
)

func main() {
	
	allConnection()
	fmt.Println("App running on " + helper.Config.Address + ":" + helper.Config.Port)

}
