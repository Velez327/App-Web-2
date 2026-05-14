// Este archivo contiene DOS tests RESUELTOS como ejemplo. Sirven como
// modelo de cómo escribir los tests que faltan: los 8 métodos restantes
// del taller no tienen tests todavía y debés escribirlos vos siguiendo
// estos patrones.
//
// Lo que aprendés leyendo este archivo:
//
//   1. TestGuardar_TableDriven — patrón table-driven con t.Run y subtests.
//      Aplicalo a métodos que tienen MÚLTIPLES casos de validación.
//
//   2. TestBuscarPorID_NegocioExiste — test simple de un solo caso.
//      Aplicalo a métodos con UN comportamiento esperado.
//
// IMPORTANTE: este archivo es solo un ejemplo. Vos vas a crear archivos
// nuevos como turista_repository_test.go y checkin_repository_test.go
// para los otros 8 métodos.
package storage

import (
	"errors"
	"testing"

	"github.com/uleam/awii/turismo/internal/errs"
	"github.com/uleam/awii/turismo/internal/models"
)

// TestGuardar_TableDriven cubre los 6 escenarios de Negocio.Guardar usando
// el patrón table-driven idiomático de Go.
//
// Los 6 casos cubren:
//
//   1. Caso feliz — un negocio válido se guarda sin error
//   2. Nombre vacío — debe fallar con ErrDatosInvalidos
//   3. Tipo no válido — debe fallar con ErrDatosInvalidos
//   4. Idiomas vacío — debe fallar con ErrDatosInvalidos
//   5. Idioma no soportado — debe fallar con ErrDatosInvalidos
//   6. ID duplicado — debe fallar con ErrYaExiste
//
// El primer caso siembra el repo. El sexto caso reusa ese mismo repo para
// probar el ID duplicado. Por eso el repo se construye UNA SOLA VEZ fuera
// del bucle, no dentro.
func TestGuardar_TableDriven(t *testing.T) {
	repo := NewNegocioMemoria()

	// Pre-condición: sembramos un negocio para poder probar "ID duplicado".
	negocioBase := models.Negocio{
		ID: 1, Nombre: "Café Manabita", Tipo: "restaurante",
		Ciudad: "Manta", IdiomasHablados: []string{"es", "en"}, Activo: true,
	}
	if err := repo.Guardar(negocioBase); err != nil {
		// t.Fatalf detiene el test inmediatamente. Si el setup falla,
		// no tiene sentido seguir corriendo el resto de los casos.
		t.Fatalf("setup falló: %v", err)
	}

	// La tabla de casos. Cada elemento es un escenario completo:
	// nombre del subtest, datos de entrada, error esperado.
	casos := []struct {
		nombre    string
		entrada   models.Negocio
		esperaErr error
	}{
		{
			nombre: "caso feliz - negocio válido",
			entrada: models.Negocio{
				ID: 100, Nombre: "Hotel Costa", Tipo: "hotel",
				Ciudad: "Manta", IdiomasHablados: []string{"es", "en"}, Activo: true,
			},
			esperaErr: nil,
		},
		{
			nombre: "nombre vacío falla",
			entrada: models.Negocio{
				ID: 101, Nombre: "", Tipo: "hotel",
				IdiomasHablados: []string{"es"}, Activo: true,
			},
			esperaErr: errs.ErrDatosInvalidos,
		},
		{
			nombre: "tipo no válido falla",
			entrada: models.Negocio{
				ID: 102, Nombre: "Negocio X", Tipo: "panaderia",
				IdiomasHablados: []string{"es"}, Activo: true,
			},
			esperaErr: errs.ErrDatosInvalidos,
		},
		{
			nombre: "lista de idiomas vacía falla",
			entrada: models.Negocio{
				ID: 103, Nombre: "Negocio Y", Tipo: "restaurante",
				IdiomasHablados: []string{}, Activo: true,
			},
			esperaErr: errs.ErrDatosInvalidos,
		},
		{
			nombre: "idioma no soportado falla",
			entrada: models.Negocio{
				ID: 104, Nombre: "Negocio Z", Tipo: "tour",
				IdiomasHablados: []string{"es", "ja"}, Activo: true, // ja=japonés no está en la lista
			},
			esperaErr: errs.ErrDatosInvalidos,
		},
		{
			nombre: "ID duplicado falla",
			entrada: models.Negocio{
				ID: 1, Nombre: "Otro Café", Tipo: "restaurante",
				IdiomasHablados: []string{"es"}, Activo: true,
			},
			esperaErr: errs.ErrYaExiste,
		},
	}

	// Iteramos sobre los casos y corremos un subtest por cada uno.
	// t.Run permite que cada subtest se reporte por separado y que se
	// puedan correr individualmente con `go test -run`.
	for _, c := range casos {
		t.Run(c.nombre, func(t *testing.T) {
			err := repo.Guardar(c.entrada)

			// errors.Is es la forma idiomática de comparar errores
			// tipados. NUNCA uses err == c.esperaErr ni
			// err.Error() == "..." — son frágiles.
			if !errors.Is(err, c.esperaErr) {
				t.Errorf("Guardar(%q): esperaba error=%v, obtuvo error=%v",
					c.entrada.Nombre, c.esperaErr, err)
			}
		})
	}
}

// TestBuscarPorID_NegocioExiste verifica el caso feliz de BuscarPorID.
//
// Este es un test SIMPLE de un solo caso. No necesita el patrón
// table-driven porque solo hay un comportamiento esperado a verificar.
//
// Los OTROS casos de BuscarPorID (ID negativo, ID inexistente) deberían
// ir en otro test, posiblemente table-driven, que VOS tenés que escribir.
func TestBuscarPorID_NegocioExiste(t *testing.T) {
	repo := NewNegocioMemoria()

	// Arrange: creamos y guardamos un negocio.
	esperado := models.Negocio{
		ID: 42, Nombre: "Manabita Crafts", Tipo: "artesania",
		Ciudad: "Manta", IdiomasHablados: []string{"es"}, Activo: true,
	}
	if err := repo.Guardar(esperado); err != nil {
		t.Fatalf("setup falló: %v", err)
	}

	// Act: buscamos el negocio por su ID.
	obtenido, err := repo.BuscarPorID(42)

	// Assert: no debe haber error y debe coincidir con lo guardado.
	if err != nil {
		t.Fatalf("no esperaba error: %v", err)
	}
	if obtenido.ID != esperado.ID {
		t.Errorf("ID: esperaba %d, obtuvo %d", esperado.ID, obtenido.ID)
	}
	if obtenido.Nombre != esperado.Nombre {
		t.Errorf("Nombre: esperaba %q, obtuvo %q", esperado.Nombre, obtenido.Nombre)
	}
	if obtenido.Tipo != esperado.Tipo {
		t.Errorf("Tipo: esperaba %q, obtuvo %q", esperado.Tipo, obtenido.Tipo)
	}
}

func TestNegocioMemoria_Eliminar(t *testing.T) {
	// Negocio que sembramos en cada caso para tener algo que eliminar
     base := models.Negocio{
     ID: 1, Nombre: "Café del Mar", Tipo: "restaurante",
     Ciudad: "Manta", IdiomasHablados: []string{"es", "en"},
    }
      casos := []struct {
      nombre string
      idEliminar int // ID que se intenta eliminar
      errEsperado error
    }{
 {
       nombre: "elimina un negocio existente",
       idEliminar: 1,
       errEsperado: nil,
    },
 {
      nombre: "ID inexistente retorna ErrNoEncontrado",
     idEliminar: 999,
     errEsperado: errs.ErrNoEncontrado,
 },
 }
     for _, c := range casos {
      t.Run(c.nombre, func(t *testing.T) {
 // Arrange: cada caso con su propio repo + el negocio sembrado
     repo := NewNegocioMemoria()
     if err := repo.Guardar(base); err != nil {
     t.Fatalf("setup falló: %v", err)
 }
 // Act
     err := repo.Eliminar(c.idEliminar)
 // Assert
     if !errors.Is(err, c.errEsperado) {
     t.Errorf("esperaba error %v, obtuvo %v", c.errEsperado, err)
    }
 })
 }
}

func TestNegocioMemoria_Listar(t *testing.T) {
	repo := NewNegocioMemoria()

	// Arrange: sembramos varios negocios.
	negocios := []models.Negocio{
		{ID: 1, Nombre: "Café del Mar", Tipo: "restaurante", Ciudad: "Manta", IdiomasHablados: []string{"es", "en"}, Activo: true},
		{ID: 2, Nombre: "Hotel Costa", Tipo: "hotel", Ciudad: "Manta", IdiomasHablados: []string{"es"}, Activo: true},
	}
	for _, n := range negocios {
		if err := repo.Guardar(n); err != nil {
			t.Fatalf("setup falló al guardar negocio %d: %v", n.ID, err)
		}
	}

	// Act
	listado := repo.Listar()

	// Assert
	if len(listado) != len(negocios) {
		t.Errorf("esperaba %d negocios, obtuvo %d", len(negocios), len(listado))
	}
	// Verificamos que cada negocio sembrado esté en el listado.
	for _, esperado := range negocios {
		encontrado := false
		for _, n := range listado {
			if n.ID == esperado.ID {
				encontrado = true
				if n.Nombre != esperado.Nombre {
					t.Errorf("ID %d: Nombre esperado %q, obtenido %q", n.ID, esperado.Nombre, n.Nombre)
				}
				if n.Tipo != esperado.Tipo {
					t.Errorf("ID %d: Tipo esperado %q, obtenido %q", n.ID, esperado.Tipo, n.Tipo)
				}
			}
		}
		if !encontrado {
			t.Errorf("negocio con ID %d no encontrado en el listado", esperado.ID)
		}
	}
}

// El test de Listar verifica que el método retorne todos los negocios
// que se han guardado, y que cada uno tenga los datos correctos. Si el
// repositorio está vacío, debería retornar un slice vacío sin error.

func TestNegocioBuscarPorID_IDNegativo(t *testing.T) {
	repo := NewNegocioMemoria()

	// Act
	_, err := repo.BuscarPorID(-1)

	// Assert
	if !errors.Is(err, errs.ErrDatosInvalidos) {
		t.Errorf("esperaba error %v, obtuvo %v", errs.ErrDatosInvalidos, err)
	}
}

func TestNegocioBuscarPorID_IDNoExiste(t *testing.T)	{
	repo := NewNegocioMemoria()

	// Act
	_, err := repo.BuscarPorID(999)

	// Assert
	if !errors.Is(err, errs.ErrNoEncontrado) {
		t.Errorf("esperaba error %v, obtuvo %v", errs.ErrNoEncontrado, err)
	}
}	

// El test de BuscarPorID_IDNegativo verifica que si se busca un ID negativo,
// el método retorne ErrDatosInvalidos. El test de BuscarPorID_IDNoExiste
// verifica que si se busca un ID que no existe en el repositorio, el método
// retorne ErrNoEncontrado.	

