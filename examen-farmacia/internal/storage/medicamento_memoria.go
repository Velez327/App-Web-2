// ARCHIVO BLOQUEADO — NO MODIFICAR
package storage

import (
	"sync"

	"github.com/joancema/examen-farmacia/internal/models"
)

// MedicamentoMemoria implementa MedicamentoRepository en memoria.
// Se usa en los tests de reglas de negocio como fake del repositorio real.
type MedicamentoMemoria struct {
	mu     sync.Mutex
	datos  map[uint]models.Medicamento
	nextID uint
}

func NuevoMedicamentoMemoria() *MedicamentoMemoria {
	return &MedicamentoMemoria{datos: make(map[uint]models.Medicamento), nextID: 1}
}

func (r *MedicamentoMemoria) Crear(h *models.Medicamento) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	h.ID = r.nextID
	r.nextID++
	r.datos[h.ID] = *h
	return nil
}

func (r *MedicamentoMemoria) ObtenerPorID(id uint) (models.Medicamento, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()
	h, ok := r.datos[id]
	return h, ok
}

func (r *MedicamentoMemoria) Listar() ([]models.Medicamento, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	lista := make([]models.Medicamento, 0, len(r.datos))
	for _, h := range r.datos {
		lista = append(lista, h)
	}
	return lista, nil
}

func (r *MedicamentoMemoria) Actualizar(h *models.Medicamento) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.datos[h.ID]; !ok {
		return ErrRegistroNoExiste
	}
	r.datos[h.ID] = *h
	return nil
}
