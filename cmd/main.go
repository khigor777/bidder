package main

import (
	"try/bidder"

	"fmt"
)
func main()  {

	c, e := bidder.ReadConfig("config.json")
	if e != nil{
		fmt.Println(e)
	}
	//pool := bidder.NewRedisPool(c)

	fmt.Println(c.Server)

}
