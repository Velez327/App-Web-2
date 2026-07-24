package storage

import (
	"gorm.io/gorm"

	"github.com/joancema/examen-farmacia/internal/models"
)

// TAREA (CP1): Implemente ClienteGORM contra la interfaz ClienteRepository.
//
// Reglas:
//   - NO cambie el nombre del tipo, del constructor ni las firmas de los métodos:
//     los tests de acceptance/ compilan contra ellos.
//   - Guíese por MedicamentoGORM (medicamento_gorm.go): es el mismo patrón.
type ClienteGORM struct {
	db *gorm.DB
}

func NuevoClienteGORM(db *gorm.DB) *ClienteGORM {
	return &ClienteGORM{db: db}
}

func (r *ClienteGORM) Crear(c *models.Cliente) error {
	// TODO: implementar.
	return r.db.Create(c).Error
}

func (r *ClienteGORM) ObtenerPorID(id uint) (models.Cliente, bool) {
	// TODO: implementar.
	var c models.Cliente
	if err := r.db.First(&c, id).Error; err != nil {
		return models.Cliente{}, false
	}
	return c, true
}

func (r *ClienteGORM) Listar() ([]models.Cliente, error) {
	// TODO: implementar.
	var lista []models.Cliente
	err := r.db.Find(&lista).Error
	return lista, err
}
