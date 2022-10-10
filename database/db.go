package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {
	connectString := "host=172.19.0.2 user=root password=root dbname=root port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(connectString))
	if err != nil {
		log.Panic("Erro ao conectar com o banco de dados: " + err.Error())
	}
}
