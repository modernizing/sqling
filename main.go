package main

import (
	. "github.com/inherd/sqling/parser"
	. "github.com/inherd/sqling/render"
	_ "github.com/pingcap/tidb/types/parser_driver"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		return
	}

	filename := args[1]
	dat, err := ioutil.ReadFile(filename)
	Check(err)
	sql := string(dat)

	Convert(sql)
}

func Convert(sql string) {
	Write(ParseSql(sql))
}
