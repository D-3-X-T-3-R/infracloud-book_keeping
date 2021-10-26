package config

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"io/ioutil"
	"os"
)

type Config struct {
	Username string `json:"username"`
	Password string `json:"password"`
	DBName   string `json:"db_name"`
	DBType   string `json:"db_type"`
}

var db *gorm.DB

func GetConn(config_path string) {
	jsonFile, err := os.Open(config_path)
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var config Config
	json.Unmarshal(byteValue, &config)
	Conn(config)
	fmt.Println("conn established")
}

func Conn(config Config) {
	conn_str := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		config.Username,
		config.Password,
		config.DBName)
	conn, err := gorm.Open(config.DBType, conn_str)
	if err != nil {
		panic(err)
	}
	db = conn
}

func GetDB() *gorm.DB {
	return db
}
