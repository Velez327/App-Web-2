package storage

import (
    
	"testing"	

	"github.com/uleam/awii/turismo/internal/errs"
	"github.com/uleam/awii/turismo/internal/models"
)
func TestTuristaMemoria_Guardar_TableDriven(t *testing.T) {
	repo := NewTuristaMemoria()			
	casos := []struct {
		nombre    string
		entrada   models.Turista	
		esperaErr error
	}{
		{
			nombre: "caso feliz - turista válido",
			entrada: models.Turista{
				ID: 1, Nombre: "Alice", Nacionalidad: "USA", IdiomaPreferido: "en",	
			},
			esperaErr: nil,
		},	
		{
			nombre: "nombre vacío falla",
			entrada: models.Turista{
				ID: 2, Nombre: "", Nacionalidad: "USA", IdiomaPreferido: "en",
			},
			esperaErr: errs.ErrDatosInvalidos,
		},
		{
			nombre: "nacionalidad vacía falla",
			entrada: models.Turista{
				ID: 3, Nombre: "Bob", Nacionalidad: "", IdiomaPreferido: "en",
			},
			esperaErr: errs.ErrDatosInvalidos,
		},
	}
	for _, caso := range casos {
		t.Run(caso.nombre, func(t *testing.T) {
			err := repo.Guardar(caso.entrada)	
			if err != caso.esperaErr {
				t.Errorf("esperaba error %v, pero got %v", caso.esperaErr, err)
			}	
		})
	}
}

func TestTuristaMemoria_BuscarPorID_TableDriven(t *testing.T) {	
	repo := NewTuristaMemoria()
	// Sembramos un turista para probar la búsqueda.
	repo.Guardar(models.Turista{ID: 1, Nombre: "Alice", Nacionalidad: "USA", IdiomaPreferido: "en"})	
	casos := []struct {
		nombre    string
		entrada   int
		esperaErr error
	}{
		{
			nombre: "caso feliz - ID existe",
			entrada: 1,
			esperaErr: nil,
		},
		{
			nombre: "ID negativo falla",
			entrada: -1,
			esperaErr: errs.ErrDatosInvalidos,
		},
		{
			nombre: "ID no existe falla",
			entrada: 999,
			esperaErr: errs.ErrNoEncontrado,
		},
	}
	for _, caso := range casos {
		t.Run(caso.nombre, func(t *testing.T) {
			_, err := repo.BuscarPorID(caso.entrada)
			if err != caso.esperaErr {
				t.Errorf("esperaba error %v, pero got %v", caso.esperaErr, err)
			}
		})
	}
}

func TestTuristaMemoria_Listar(t *testing.T) {
	repo := NewTuristaMemoria()
	// Sembramos varios turistas.		
	turistas := []models.Turista{
		{ID: 1, Nombre: "Alice", Nacionalidad: "USA", IdiomaPreferido: "en"},
		{ID: 2, Nombre: "Bob", Nacionalidad: "UK", IdiomaPreferido: "en"},
	}
	for _, p := range turistas {
		repo.Guardar(p)
	}
	resultado := repo.Listar()
	if len(resultado) != len(turistas) {
		t.Fatalf("esperaba %d turistas, pero got %d", len(turistas), len(resultado))
	}	
	// Opcional: validar que los turistas listados coinciden con los sembrados.
	for _, p := range turistas {
		encontrado := false
		for _, r := range resultado {
			if r.ID == p.ID && r.Nombre == p.Nombre && r.Nacionalidad == p.Nacionalidad && r.IdiomaPreferido == p.IdiomaPreferido {
				encontrado = true
				break
			}	
		}
		if !encontrado {
			t.Errorf("turista ID %d no encontrado en resultado", p.ID)
		}
	}
}