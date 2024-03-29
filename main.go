package main

import (
	"log"

	"./database"
	. "./model"
	. "./route"
	"github.com/jinzhu/gorm"
)

func main() {
	dbUser, dbPwd, dbName := "root", "limitless", "api"

	db, err := database.Connect(dbUser, dbPwd, dbName)
	defer db.Close()

	if err != nil {
		log.Fatalln(err)
	}

	err = db.DB().Ping()

	if err != nil {
		log.Fatalln(err)
	}

	Migration(db)

	router := SetupRoutes(db)
	router.Run(":8000")
}

func Migration(db *gorm.DB) {
	db.AutoMigrate(&Structure{})
	db.AutoMigrate(&Abonnement{})
	db.AutoMigrate(&Position{})
	db.AutoMigrate(&Fonction{})
	db.AutoMigrate(&Acteur{})
	db.AutoMigrate(&CategorieArticle{})
	db.AutoMigrate(&Article{})
	db.AutoMigrate(&Livraison{})
	db.AutoMigrate(&BonLivraison{})
	db.AutoMigrate(&Facture{})
	db.AutoMigrate(&BonCommande{})
}
