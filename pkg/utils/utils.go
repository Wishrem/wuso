package utils

import (
	"strings"

	"github.com/Wishrem/wuso/config"
)

func GetMysqlDSN() string {
	return strings.Join([]string{config.MySQL.UserName, ":", config.MySQL.Password, "@tcp(", config.MySQL.Host, ":", config.MySQL.Port, ")/", config.MySQL.Database, "?charset=" + config.MySQL.Charset + "&parseTime=true"}, "")
}
