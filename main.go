package main

import (
	"github.com/inherd/sqling/cmd"
	_ "github.com/pingcap/tidb/types/parser_driver"
)

func main() {
	cmd.Execute()
}
