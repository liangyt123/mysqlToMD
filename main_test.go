package main

import (
	"testing"
)

func TestToPrint(t *testing.T) {
	p := ParamStruct{}
	p.Host = "192.168.126.70"
	p.User = "root"
	p.Password = "123456"
	p.Dasebase = "riskauth"
	p.Port = "3306"

	c := Dbutils{}
	c.InitDbutils()
	c.SetDataBaseMap(p)
	c.PrintMapToFile(p)
}
