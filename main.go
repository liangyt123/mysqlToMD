package main

import (
	"fmt"
)

func main() {
	param, err := ParseCommnd()
	if err != nil {
		return
	}

	c := Dbutils{}
	c.InitDbutils()
	c.SetDataBaseMap(param)
	c.PrintMapToFile(param)
	fmt.Println("ok")
}
