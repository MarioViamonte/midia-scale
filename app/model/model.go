package model

import(
	"fmt"

	"gorm.io/driver/postgres" 
	"gorm.io/gorm"
)

var db *gorm.DB

type voluntary struct{
	ID uint64 `json: "id" gorm:"primaryKey"`
	name string	`json: "name"`
	discipler string `json: "discipler"`  //discipulador
	sectorID int64 
	sector sector `json: "sector" gorm: "foreignKey:sectorRefer, references: sectorID"` //ministerio
}

type sector struct{
	ID int64 `json:"id" gorm:"primaryKey"`
	Name string`json: "name"`
}

func Setup(){

	dsn := "host=localhost user=admin password=test dbname=admin port=5432 sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil{
		panic(err)
	}

	err = db.AutoMigrate(&voluntary{}, &sector{})
	if err != nil{
		fmt.Println(err)
	}

}