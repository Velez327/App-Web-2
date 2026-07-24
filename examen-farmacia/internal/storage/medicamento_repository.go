// ARCHIVO BLOQUEADO — NO MODIFICAR
package storage

import "github.com/joancema/examen-farmacia/internal/models"

// MedicamentoRepository define el contrato de persistencia de la Entidad A.
type MedicamentoRepository interface {
	Crear(h *models.Medicamento) error
	ObtenerPorID(id uint) (models.Medicamento, bool)
	Listar() ([]models.Medicamento, error)
	Actualizar(h *models.Medicamento) error
}
