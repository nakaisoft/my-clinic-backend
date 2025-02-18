package models

import "time"

type ClinicModel struct {
	ID    string `gorm:"type:string;primaryKey" json:"id"`
	Title string

	CreatedAt time.Time       `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time       `gorm:"autoUpdateTime" json:"updated_at"`

	Employees []EmployeeModel `gorm:"foreignKey:EmployeeID;constraint:OnDelete:CASCADE" json:"Employees,omitempty"`
}
