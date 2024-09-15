package db

type Driver string

const Mysql Driver = "mysql"
const Pgsql Driver = "pgsql"
const Sqlite Driver = "sqlite"
const DM Driver = "dm" // 达梦数据库
const Oracle Driver = "oracle"
const ClickHouse Driver = "clickhouse"
