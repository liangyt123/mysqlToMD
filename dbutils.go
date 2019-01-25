package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/lunny/log"
)

const FileName = "db.md"

type ParamStruct struct {
	Dasebase string
	Table    string
	Host     string
	User     string
	Password string
	Port     string
	Path     string
}

type TableStruct struct {
	//TableName     string
	ColumnName    string
	ColumnType    string
	ColumnKey     string
	Extra         string
	IsNull        string
	ColumnComment string
}

type Dbutils struct {

	//数据库名->表名
	//map[deparment][list]
	DbInfo map[string][]TableStruct
}

func (c *Dbutils) InitDbutils() {
	c.DbInfo = make(map[string][]TableStruct, 0)
}

//连接数据库并进行获取信息
func (c *Dbutils) SetDataBaseMap(param *ParamStruct) {

	var rows *sql.Rows
	strsql := "select table_name,column_name,column_type,column_key,extra,is_nullable,column_comment from information_schema.columns where table_schema=?"

	//获取数据库连接
	connectsql := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", param.User, param.Password, param.Host, param.Port, param.Dasebase)
	fmt.Println("connect...,", connectsql)
	db, err := sql.Open("mysql", connectsql)
	if err != nil {
		panic(err.Error())
	}

	defer func() {
		db.Close()
	}()

	if param.Table == "" {
		rows, err = db.Query(strsql, param.Dasebase)
		if err != nil {
			panic(err.Error())
		}
	} else {
		strsql += " and table_name = ?"
		rows, err = db.Query(strsql, param.Dasebase, param.Table)
		if err != nil {
			panic(err.Error())
		}
	}

	for rows.Next() {
		tmptpl := TableStruct{}
		tmptplname := ""
		if err := rows.Scan(&tmptplname, &tmptpl.ColumnName, &tmptpl.ColumnType, &tmptpl.ColumnKey, &tmptpl.Extra, &tmptpl.IsNull, &tmptpl.ColumnComment); err != nil {
			panic(err.Error())
		}

		//加入 就这一次赋值
		if _, ok := c.DbInfo[tmptplname]; !ok {
			tmptitle := TableStruct{}
			tmptitle.ColumnName = "名称"
			tmptitle.ColumnType = "类型"
			tmptitle.ColumnKey = "类型特性"
			tmptitle.Extra = "类型行为"
			tmptitle.IsNull = "是否为空"
			tmptitle.ColumnComment = "备注"
			c.DbInfo[tmptplname] = append(c.DbInfo[tmptplname], tmptitle)
		}
		c.DbInfo[tmptplname] = append(c.DbInfo[tmptplname], tmptpl)
	}
	//等到所有值之后进行输出

}

//
func (c *Dbutils) ConvertMapToString() string {
	mdstr := ""
	for k, v := range c.DbInfo {
		mdstr = mdstr + "###	表名:" + k + "\n"

		for i, v2 := range v {

			mdstr += fmt.Sprintf(" | %-15s | %-15s | %-15s | %-15s | %-15s | %-15s |\n", v2.ColumnName, v2.ColumnType, v2.ColumnKey, v2.Extra, v2.IsNull, v2.ColumnComment)
			if i == 0 {
				mdstr += fmt.Sprintf(" | -------------------- |-------------------- |-------------------- |-------------------- |-------------------- |\n")
			}
		}
	}
	return mdstr
}

//创建文件
func (c *Dbutils) PrintMapToFile(param *ParamStruct) {
	file, err := os.Create(FileName)
	if err != nil {
		log.Debugf("open %s file failed", FileName)
		return
	}
	defer file.Close()

	wbuff := bufio.NewWriter(file)
	_, err = wbuff.WriteString(c.ConvertMapToString())
	if err != nil {
		log.Debugf("write %s file fail", FileName)
	}
	wbuff.Flush()
	return
}
