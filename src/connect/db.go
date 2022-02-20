package connect

import (
	"fmt"
	"query_generator/src/tools"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConDB(ini *tools.Conn) (conn *gorm.DB, err error) {
	conStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", ini.UserName, ini.Password, ini.Host, ini.Port, ini.Schema)
	db, err := gorm.Open(mysql.Open(conStr), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect database")
	}
	cn, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to connect database")
	}
	defer cn.Close()
	return db, err
}
