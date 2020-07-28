package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

var prefix = "BUNKER"
var storagePath = ""

func getKeyName(key string) string {
	if len(prefix) > 0 {
		return fmt.Sprint(prefix, "_", key)
	}
	return key
}

func SetPrefix(s string) {
	prefix = s
}

func GetStoragePath() string {
	if len(storagePath) < 1 {
		storagePath = os.Getenv(getKeyName("STORAGE_PATH"))
		if len(storagePath) < 1 {
			storagePath = "./storage"
		}
		storagePath, _ = filepath.Abs(storagePath)
	}
	return storagePath
}

func GetBindPort() int {
	port, err := strconv.ParseInt(os.Getenv(getKeyName("BIND_PORT")), 10, 32)
	if err != nil {
		port = 3000
	}
	return int(port)
}

//noinspection ALL
func GetDbDialect() string {
	dialect := os.Getenv(getKeyName("DB_DIALECT"))
	if len(dialect) < 1 {
		return "sqlite3"
	}
	return dialect
}

//noinspection ALL
func GetDbArgs() string {
	args := os.Getenv(getKeyName("DB_ARGS"))
	if len(args) < 1 && GetDbDialect() == "sqlite3" {
		return "bunker.db.sqlite"
	}
	return args

}
