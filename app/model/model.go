package model

import(
	"fmt"

	"gorm.io/driver/postgres" 
	"gorm.io/gorm"
)

var db *gorm.DB

type voluntary struct{
	ID uint64 `gorm:"primaryKey"`
	name string	
	discipler string 
	sectorID int64 
	sector sector `gorm: "foreignKey:sectorRefer, references: sectorID"` //ministerio
}

type sector struct{
	ID int64 `gorm:"primaryKey"`
	Name string
}

func Setup(){

	dsn := "host=localhost user=admin password=test dbname=admin port=5432 sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil{
		panic(err)
	}

	err = db.AutoMigrate(&voluntary{},&sector{})
	if err != nil{
		fmt.Println(err)
	}

}