package storage

import (
	"testing"

	"github.com/uleam/awii/turismo/internal/errs"
	"github.com/uleam/awii/turismo/internal/models"
)

func TestCheckInMemoria_Guardar(t *testing.T) {
	// Para probar el repositorio de check-ins necesitamos tener un repositorio
	// de turistas y otro de negocios, porque el método Guardar hace validación
	// cruzada con ambos.
	turistas := NewTuristaMemoria()
	negocios := NewNegocioMemoria()
	repo := NewCheckInMemoria(turistas, negocios)	

	// Sembramos un turista y un negocio para que existan en el repositorio
	// y así poder probar la validación cruzada.
	turistas.Guardar(models.Turista{ID: 1, Nombre: "Alice", Nacionalidad: "USA", IdiomaPreferido: "inglés"})
	negocios.Guardar(models.Negocio{ID: 1, Nombre: "Café Manta", Tipo: "café", IdiomasHablados: []string{"español", "inglés"}})

	tests := []struct {
		name    string
		input   models.CheckIn
		wantErr error
	}{
		{	
			name:    "fecha vacía",
			input:   models.CheckIn{ID: 1, TuristaID: 1, NegocioID: 1, Fecha: "", Calificacion: 4},
			wantErr: errs.ErrDatosInvalidos,
		},
		{
			name:    "calificación menor a 1",
			input:   models.CheckIn{ID: 2, TuristaID: 1, NegocioID: 1, Fecha: "2024-06-01", Calificacion: 0},
			wantErr: errs.ErrDatosInvalidos,
		},
		{
			name:    "calificación mayor a 5",
			input:   models.CheckIn{ID: 3, TuristaID: 1, NegocioID: 1, Fecha: "2024-06-01", Calificacion: 6},
			wantErr: errs.ErrDatosInvalidos,
		},
		{
			name:    "turista no existe",
			input:   models.CheckIn{ID: 4, TuristaID: 999, NegocioID: 1, Fecha: "2024-06-01", Calificacion: 4},
			wantErr: errs.ErrNoEncontrado,
		},
		{
			name:    "negocio no existe",
			input:   models.CheckIn{ID: 5, TuristaID: 1, NegocioID: 999, Fecha: "2024-06-01", Calificacion: 4},
			wantErr: errs.ErrNoEncontrado,
		},
		{
			name:    "check-in válido",
			input:   models.CheckIn{ID: 9, TuristaID: 1, NegocioID: 1, Fecha: "2024-06-01", Calificacion: 4},
			wantErr: nil,	
		},
		{
			name:    "ID de check-in ya existe",
			input:   models.CheckIn{ID: 6, TuristaID: 1, NegocioID: 1, Fecha: "2024-06-02", Calificacion: 5},	
			wantErr: errs.ErrYaExiste,
		},
	}	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := repo.Guardar(tt.input)
			if err != tt.wantErr {
				t.Errorf("Guardar(%v) error = %v, wantErr %v", tt.input, err, tt.wantErr)
			}
		})
	}
}

func TestCheckInMemoria_BuscarPorTurista(t *testing.T) {
	turistas := NewTuristaMemoria()
	negocios := NewNegocioMemoria()
	repo := NewCheckInMemoria(turistas, negocios)
	turistas.Guardar(models.Turista{ID: 1, Nombre: "Alice", Nacionalidad: "USA", IdiomaPreferido: "inglés"})
	tests := []struct {
		name      string
		turistaID int
		wantErr   error
	}{
		{
			name:      "ID de turista negativo",
			turistaID: -1,
			wantErr:   errs.ErrDatosInvalidos,
		},
		{
			name:      "turista no existe",
			turistaID: 999,
			wantErr:   errs.ErrNoEncontrado,
		},
		{
			name:      "turista existe pero no tiene check-ins",
			turistaID: 1,
			wantErr:   nil,
		},	
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := repo.BuscarPorTurista(tt.turistaID)	
			if err != tt.wantErr {
				t.Errorf("BuscarPorTurista(%d) error = %v, wantErr %v", tt.turistaID, err, tt.wantErr)
			}	
		})
	}
}

func TestCheckInMemoria_Listar(t *testing.T) {
	turistas := NewTuristaMemoria()
	negocios := NewNegocioMemoria()
	repo := NewCheckInMemoria(turistas, negocios)	
	if len(repo.Listar()) != 0 {	
		t.Errorf("Listar() = %v, want empty slice", repo.Listar())
	}
}

