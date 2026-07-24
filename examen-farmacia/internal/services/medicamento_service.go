// ARCHIVO BLOQUEADO — NO MODIFICAR
package services

import (
	"github.com/joancema/examen-farmacia/internal/models"
	"github.com/joancema/examen-farmacia/internal/storage"
)

// MedicamentoService contiene la lógica de negocio de la Entidad A.
// Está completo: úselo como ejemplo de cómo un service valida datos,
// devuelve errores de dominio y delega la persistencia al repository.
type MedicamentoService struct {
	repo storage.MedicamentoRepository
}

func NuevoMedicamentoService(repo storage.MedicamentoRepository) *MedicamentoService {
	return &MedicamentoService{repo: repo}
}

func (s *MedicamentoService) Crear(h *models.Medicamento) error {
	if h.Nombre == "" || h.PrecioUnitario <= 0 {
		return ErrDatosInvalidos
	}
	return s.repo.Crear(h)
}

func (s *MedicamentoService) ObtenerPorID(id uint) (models.Medicamento, error) {
	h, ok := s.repo.ObtenerPorID(id)
	if !ok {
		return models.Medicamento{}, ErrNoEncontrado
	}
	return h, nil
}

func (s *MedicamentoService) Listar() ([]models.Medicamento, error) {
	return s.repo.Listar()
}
