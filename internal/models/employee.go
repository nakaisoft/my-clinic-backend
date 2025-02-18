package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Definição dos tipos de função que um funcionário pode ter
type EmployeeRole string

const (
	RoleManager   EmployeeRole = "manager"   // Acesso total
	RoleDoctor    EmployeeRole = "doctor"    // Apenas agendamentos e pacientes
	RoleEmployee  EmployeeRole = "employee"  // Acesso padrão
)

// EmployeeModel representa um usuário da aplicação
type EmployeeModel struct {
	ID       string       `gorm:"type:string;primaryKey" json:"id"`
	Name     string       `gorm:"not null" json:"name"`
	Email    string       `gorm:"unique;not null" json:"email"`
	Password string       `gorm:"not null" json:"-"` // Não expor senha na API
	Role     EmployeeRole `gorm:"type:varchar(20);not null" json:"role"`
	ClinicID *string      `gorm:"type:string;index" json:"clinic_id,omitempty"` // FK opcional para a clínica
	Clinic   *ClinicModel `gorm:"foreignKey:ClinicID" json:"clinic,omitempty"`

	CreatedAt time.Time    `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time    `gorm:"autoUpdateTime" json:"updated_at"`
}

// Create insere um novo funcionário no banco de dados e gera um UUID se necessário
func (e *EmployeeModel) Create(db *gorm.DB) error {
	if e.ID == "" {
		e.ID = uuid.New().String()
	}
	return db.Create(e).Error
}
