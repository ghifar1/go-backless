package databases

import "os"

func Init() {
	connType := os.Getenv("DB_CONNECTION")
	switch connType {
	case "mysql":
		MysqlInit()
	}
}

func MysqlInit() {

}
