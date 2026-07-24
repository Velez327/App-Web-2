package services

import (
	"github.com/joancema/examen-farmacia/internal/models"
	"github.com/joancema/examen-farmacia/internal/storage"
)

// TAREA (CP1): Implemente ClienteService.
//
// Reglas:
//   - NO cambie el nombre del tipo, del constructor ni las firmas de los métodos.
//   - Cliente no tiene reglas de negocio complejas: valide lo evidente según
//     las pantallas (campos obligatorios -> ErrDatosInvalidos) y delegue al
//     repository. Guíese por MedicamentoService.
type ClienteService struct {

	repo storage.ClienteRepository
}

func NuevoClienteService(repo storage.ClienteRepository) *ClienteService {
	return &ClienteService{repo: repo}
}

func (s *ClienteService) Crear(c *models.Cliente) error {
	// TODO: implementar.
	if c.Nombre == "" || c.Cedula == "" || c.Telefono == "" {
		return ErrDatosInvalidos
	}
	return s.repo.Crear(c)
}

func (s *ClienteService) ObtenerPorID(id uint) (models.Cliente, error) {
	// TODO: implementar.
	c, ok := s.repo.ObtenerPorID(id)
	if !ok {
		return models.Cliente{}, ErrNoEncontrado
	}
	return c, nil
}

func (s *ClienteService) Listar() ([]models.Cliente, error) {
	// TODO: implementar.
	return s.repo.Listar()
}
