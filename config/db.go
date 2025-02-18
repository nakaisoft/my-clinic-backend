package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// NewDatabase cria e retorna uma conexão com o banco de dados usando GORM.
func NewDatabase() *gorm.DB {
	dsn := "user:password@tcp(localhost:3306)/mydb?charset=utf8mb4&parseTime=True&loc=Local" // Para MySQL
	// dsn := "host=localhost user=myuser password=mypassword dbname=mydb port=5432 sslmode=disable" // Para PostgreSQL

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Erro ao conectar ao banco:", err)
	}

	fmt.Println("Conectado ao banco de dados com sucesso!")
	return db
}
