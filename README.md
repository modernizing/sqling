# SQLing

> build domain model from SQL file

## Usage

1. download sqling from [release](https://github.com/inherd/sqling/releases) or

```bash
go get github.com/inherd/sqling
```

2. dump sql file

```bash
mysqldump -u root -p -h localhost --no-data mall > database.sql
```

3. run `sqling -i database.sql`


### CLI Usage

CLI

```bash
Sqling is a modeling tool to build from SQL file.

Usage:
  sqling [flags]

Flags:
  -h, --help                 help for sqling
  -i, --input string         input file
  -t, --output_type string   output file type, support for puml, json (default "puml")
```

## Todo

Todo:

 - [x] import sql parser
 - [x] use cobra as cli tools
 - [x] output
    - [x] json
    - [x] puml
 - [ ] web runtime
    - [ ] web assembly

Thinks:

 - [ ] connect to database
    - [ ] search from: MySQL information_schema.KEY_COLUMN_USAGE
    - [ ] query select table_name from information_schema.tables where table_schema='csdb' and table_type='base table';

Refs:

 - [MysqlParser](https://github.com/mysql/mysql-server/blob/8.0/sql/sql_yacc.yy)
 - [Pingcap Parser BNF](https://github.com/pingcap/parser/blob/81106e4996bfdaaf5f0ef87ac8280d03b719594d/compatibility_reporter/mysql80_bnf.txt)

Related projects:

 - [SOAR](https://github.com/XiaoMi/soar) - 是一个对 SQL 进行优化和改写的自动化工具。
 - [SQLFlow](https://github.com/sql-machine-learning/sqlflow)  is a compiler that compiles a SQL program to a workflow that runs on Kubernetes.

License
---

@ 2021 This code is distributed under the MIT license. See `LICENSE` in this directory.
