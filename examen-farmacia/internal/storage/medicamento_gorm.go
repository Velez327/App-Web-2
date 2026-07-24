// ARCHIVO BLOQUEADO — NO MODIFICAR
package storage

import (
	"gorm.io/gorm"

	"github.com/joancema/examen-farmacia/internal/models"
)

// MedicamentoGORM implementa MedicamentoRepository sobre GORM.
// Esta implementación está completa: úsela como plantilla para ClienteGORM
// y VentaGORM, que usted debe implementar.
type MedicamentoGORM struct {
	db *gorm.DB
}

func NuevoMedicamentoGORM(db *gorm.DB) *MedicamentoGORM {
	return &MedicamentoGORM{db: db}
}

func (r *MedicamentoGORM) Crear(h *models.Medicamento) error {
	return r.db.Create(h).Error
}

func (r *MedicamentoGORM) ObtenerPorID(id uint) (models.Medicamento, bool) {
	var h models.Medicamento
	if err := r.db.First(&h, id).Error; err != nil {
		return models.Medicamento{}, false
	}
	return h, true
}

func (r *MedicamentoGORM) Listar() ([]models.Medicamento, error) {
	var lista []models.Medicamento
	err := r.db.Find(&lista).Error
	return lista, err
}

func (r *MedicamentoGORM) Actualizar(h *models.Medicamento) error {
	return r.db.Save(h).Error
}
