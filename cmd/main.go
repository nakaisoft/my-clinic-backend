package main

import (
	"fmt"
	"log"

	database "github.com/nakaisoft/my-clinic-backend/config"
	"github.com/nakaisoft/my-clinic-backend/internal/models"
)

func main() {
	db := database.NewDatabase()

	// AutoMigrate: Cria ou atualiza as tabelas automaticamente
	err := db.AutoMigrate(&models.PatientModel{}, &models.PersonalAddress{})
	if err != nil {
		log.Fatal("Erro ao rodar migração:", err)
	}

	fmt.Println("Migração concluída com sucesso!")
}
