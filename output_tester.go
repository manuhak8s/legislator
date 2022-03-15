package main

import (
	"fmt"
    "github.com/manuhak8s/legislator/pkg/config"
)

func main() {
	var c config.Config
	c.ReadConfig()

	fmt.Println(c)
	fmt.Println(c.GetDefaultPolicy())
	fmt.Println(c.GetConnectedSets())

	defaultPolicy, _ := c.GetDefaultPolicy()
	fmt.Println(defaultPolicy)

	connectedSets, _ := c.GetConnectedSets()
	fmt.Println(connectedSets[2])
}
