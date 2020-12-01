package sqldb

import "sfgo/core/config"

var DBConfig struct {
	Type   string
	Dsn    string
	DbName string
}

func init() {
	config.Register("sfgo.db.sqldb", &DBConfig)
}
