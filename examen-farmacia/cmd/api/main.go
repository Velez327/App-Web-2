// ARCHIVO BLOQUEADO — NO MODIFICAR
package main

import (
	"log"
	"net/http"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"

	"github.com/joancema/examen-farmacia/internal/handlers"
	"github.com/joancema/examen-farmacia/internal/models"
	"github.com/joancema/examen-farmacia/internal/services"
	"github.com/joancema/examen-farmacia/internal/storage"
)

func main() {
	db, err := gorm.Open(sqlite.Open("farmacia.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("no se pudo abrir la base de datos: %v", err)
	}

	if err := db.AutoMigrate(
		&models.Medicamento{},
		&models.Cliente{},
		&models.Venta{},
	); err != nil {
		log.Fatalf("error en la migración: %v", err)
	}

	sembrarMedicamentos(db)

	// Repositories (GORM)
	medicamentoRepo := storage.NuevoMedicamentoGORM(db)
	clienteRepo := storage.NuevoClienteGORM(db)
	ventaRepo := storage.NuevaVentaGORM(db)

	// Services
	medicamentoSvc := services.NuevoMedicamentoService(medicamentoRepo)
	clienteSvc := services.NuevoClienteService(clienteRepo)
	ventaSvc := services.NuevaVentaService(ventaRepo, medicamentoRepo, clienteRepo)

	// Handlers + Router
	router := handlers.NuevoRouter(
		handlers.NuevoMedicamentoHandler(medicamentoSvc),
		handlers.NuevoClienteHandler(clienteSvc),
		handlers.NuevaVentaHandler(ventaSvc),
	)

	log.Println("API de la farmacia escuchando en http://localhost:8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}

// sembrarMedicamentos carga el catálogo inicial solo si la tabla está vacía.
// Los clientes y ventas se crean vía API.
func sembrarMedicamentos(db *gorm.DB) {
	var total int64
	db.Model(&models.Medicamento{}).Count(&total)
	if total > 0 {
		return
	}
	iniciales := []models.Medicamento{
		{Nombre: "Paracetamol 500mg", PrecioUnitario: 8.50, Stock: 10, Activo: true},
		{Nombre: "Ibuprofeno 400mg", PrecioUnitario: 6.00, Stock: 4, Activo: true},
		{Nombre: "Suero oral", PrecioUnitario: 5.00, Stock: 2, Activo: true},
		{Nombre: "Jarabe descontinuado", PrecioUnitario: 15.00, Stock: 3, Activo: false},
	}
	for i := range iniciales {
		db.Create(&iniciales[i])
	}
}
