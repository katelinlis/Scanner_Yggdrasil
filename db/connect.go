package db

import (
	"YggdrasilMap/db/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"time"
)

func Database() *gorm.DB {
	//open a db connection
	db, err := gorm.Open("postgres", "  user=postgres password=admin dbname=Peers  sslmode=disable ")

	if err != nil {
		panic(err)
	}

	//TODO Не работает нужно реализовать подключение к бд подругому
	db.DB().SetMaxOpenConns(10)
	db.DB().SetMaxIdleConns(2)
	db.DB().SetConnMaxLifetime(time.Hour)

	db.AutoMigrate(&model.Peer{}, &model.PeerLinks{})
	return db
}
