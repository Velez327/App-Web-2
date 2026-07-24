// ARCHIVO BLOQUEADO — NO MODIFICAR
package storage

import (
	"sync"

	"github.com/joancema/examen-farmacia/internal/models"
)

// VentaMemoria implementa VentaRepository en memoria.
// Se usa en los tests de reglas de negocio como fake del repositorio real.
type VentaMemoria struct {
	mu     sync.Mutex
	datos  map[uint]models.Venta
	nextID uint
}

func NuevaVentaMemoria() *VentaMemoria {
	return &VentaMemoria{datos: make(map[uint]models.Venta), nextID: 1}
}

func (r *VentaMemoria) Crear(a *models.Venta) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	a.ID = r.nextID
	r.nextID++
	r.datos[a.ID] = *a
	return nil
}

func (r *VentaMemoria) ObtenerPorID(id uint) (models.Venta, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()
	a, ok := r.datos[id]
	return a, ok
}

func (r *VentaMemoria) Listar() ([]models.Venta, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	lista := make([]models.Venta, 0, len(r.datos))
	for _, a := range r.datos {
		lista = append(lista, a)
	}
	return lista, nil
}

func (r *VentaMemoria) Actualizar(a *models.Venta) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.datos[a.ID]; !ok {
		return ErrRegistroNoExiste
	}
	r.datos[a.ID] = *a
	return nil
}
