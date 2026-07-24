package models

import "gorm.io/gorm"

// TAREA (CP1): Complete los campos de Venta según lo que muestran las pantallas.
//
// Pistas de trabajo:
//   - Un Venta referencia a un Medicamento y a un Cliente (claves foráneas).
//   - Recuerde el campo de estado (use las constantes de estados.go) y el total.
//   - Los tests de acceptance/ compilan contra los nombres EXACTOS de los campos.
type Venta struct {
	gorm.Model
	// TODO: agregue aquí los campos.
	MedicamentoID uint        `gorm:"not null" json:"medicamento_id"`
	Medicamento   Medicamento `gorm:"foreignKey:MedicamentoID" json:"medicamento"`
	ClienteID     uint        `gorm:"not null" json:"cliente_id"`
	Cliente       Cliente     `gorm:"foreignKey:ClienteID" json:"cliente"`
	Cantidad      uint        `gorm:"not null" json:"cantidad"`
	Estado        string      `gorm:"size:20;not null" json:"estado"`
	Total         float64     `gorm:"not null" json:"total"`
}
