package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// TODO: とりあえずでまずはinfra層に実装を置いとくが、ドメイン層にインタフェースを持たせたい
func Connect() (*gorm.DB, error) {
	return gorm.Open("mysql", "root:root@tcp(localhost:3306)/todos?charset=utf8&parseTime=True&loc=Local")
}
