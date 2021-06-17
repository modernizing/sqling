# SQLing

> SQL to PUML

Todo:

 - [x] import sql parser
 - [x] render
    - [ ] mermaid
    - [x] puml
    - [ ] graphviz
 - [ ] align coco struct
 - [ ] use cobra as cli tools
 - [ ] output
    - [ ] json
    - [x] puml
    - [ ] mermaid
 - [ ] connect to database
    - [ ] search from: MySQL information_schema.KEY_COLUMN_USAGE
    - [ ] query select table_name from information_schema.tables where table_schema='csdb' and table_type='base table';

Refs:

 - [MysqlParser](https://github.com/mysql/mysql-server/blob/8.0/sql/sql_yacc.yy)
 - [Pingcap Parser](https://github.com/pingcap/parser/blob/81106e4996bfdaaf5f0ef87ac8280d03b719594d/compatibility_reporter/mysql80_bnf.txt)

License
---

@ 2021 This code is distributed under the MIT license. See `LICENSE` in this directory.
