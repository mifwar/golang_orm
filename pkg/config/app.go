package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"

	yaml "gopkg.in/yaml.v2"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

type config struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DBname   string `yaml:"dbname"`
}

func getConfig(c *config) {

	conf_dir, _ := filepath.Abs("pkg/config/config/config.yml")
	conf, err := ioutil.ReadFile(conf_dir)

	if err != nil {
		log.Fatal("error when reading config: ", err.Error())
	}

	err = yaml.Unmarshal(conf, c)

	if err != nil {
		log.Fatal("error at json marshal: ", err.Error())
	}
}

func Connect() {

	dbConfig := config{}
	additionalConfig := "?charset=utf8&parseTime=True&loc=Local"
	getConfig(&dbConfig)
	gormCommand := fmt.Sprintf("%s:%s@/%s%s", dbConfig.Username, dbConfig.Password, dbConfig.DBname, additionalConfig)

	d, err := gorm.Open("mysql", gormCommand)
	if err != nil {
		log.Fatal(err.Error())
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
