package db

import (
	"gapi/models"
	nanoid "github.com/matoous/go-nanoid/v2"
	db "github.com/sonyarouje/simdb"
	"time"
)

func Database() *db.Driver {
	driver, err := db.New("data")
	if err != nil {
		panic(err)
	}
	return driver
}

func AddLog(s string) {
	driver := Database()
	id, _ := nanoid.New()
	logModel := models.RequestLOG{
		CID:  id,
		Data: s,
		Time: time.Now().String(),
	}
	err := driver.Insert(logModel)
	if err != nil {
		panic(err)
	}
}
