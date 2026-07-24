package storage

import (

	"gorm.io/gorm"

	"github.com/joancema/examen-farmacia/internal/models"
)

// TAREA (CP2): Implemente VentaGORM contra la interfaz VentaRepository.
//
// Reglas:
//   - NO cambie el nombre del tipo, del constructor ni las firmas de los métodos.
//   - Guíese por MedicamentoGORM: es el mismo patrón con una entidad distinta.
//   - Recuerde: aquí NO va lógica de negocio. Solo persistencia.
type VentaGORM struct {
	db *gorm.DB
}

func NuevaVentaGORM(db *gorm.DB) *VentaGORM {
	return &VentaGORM{db: db}
}

func (r *VentaGORM) Crear(a *models.Venta) error {
	// TODO: implementar.
		return r.db.Create(a).Error
}

func (r *VentaGORM) ObtenerPorID(id uint) (models.Venta, bool) {
	// TODO: implementar.
	var a models.Venta
	if err := r.db.First(&a, id).Error; err != nil {
		return models.Venta{}, false
	}
	return a, true
}

func (r *VentaGORM) Listar() ([]models.Venta, error) {
	// TODO: implementar.
	var lista []models.Venta
	err := r.db.Find(&lista).Error
	return lista, err
}

func (r *VentaGORM) Actualizar(a *models.Venta) error {
	// TODO: implementar.
	return r.db.Save(a).Error
}
