package main

import (
	"fmt"
	"github.com/xwb1989/sqlparser"
)

func main() {
	sql := "CREATE TABLE `cms_help` (\n  `id` bigint(20) NOT NULL AUTO_INCREMENT,\n  `category_id` bigint(20) DEFAULT NULL,\n  `icon` varchar(500) DEFAULT NULL,\n  `title` varchar(100) DEFAULT NULL,\n  `show_status` int(1) DEFAULT NULL,\n  `create_time` datetime DEFAULT NULL,\n  `read_count` int(1) DEFAULT NULL,\n  `content` text,\n  PRIMARY KEY (`id`)\n) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='帮助表';\n"
	stmt, err := sqlparser.Parse(sql)
	if err != nil {
		// Do something with the err
	}

	// Otherwise do something with stmt
	switch stmt := stmt.(type) {
	case *sqlparser.Select:
		_ = stmt
	case *sqlparser.Insert:
	case *sqlparser.DDL:
		switch stmt.Action {
		case "create":
			fmt.Println("..")
		}
	default:
		fmt.Println(stmt)
	}
}