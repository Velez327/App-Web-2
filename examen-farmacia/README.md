# Examen Práctico Final (C4) — Temática: FARMACIA

**TDI-601 Aplicaciones Web II — 2026-1**

Usted debe construir el backend que necesitan las pantallas de `pantallas/`,
siguiendo la arquitectura de este repositorio. Lea el enunciado en PDF antes
de empezar.

## Qué hay en este repositorio

| Carpeta / archivo | Estado | Descripción |
|---|---|---|
| `pantallas/` | BLOQUEADO | Su única especificación funcional. Mírelas primero. |
| `acceptance/` | BLOQUEADO | Los tests que definen su nota. |
| `internal/models/medicamento.go` | BLOQUEADO | Modelo de ejemplo (Entidad A). |
| `internal/models/estados.go` | BLOQUEADO | Constantes de estado dla venta. |
| `internal/models/cliente.go` | **TAREA CP1** | Complete los campos. |
| `internal/models/venta.go` | **TAREA CP1** | Complete los campos. |
| `internal/storage/*_repository.go` | BLOQUEADO | Las interfaces (contratos). |
| `internal/storage/*_memoria.go` | BLOQUEADO | Implementaciones en memoria de referencia. |
| `internal/storage/medicamento_gorm.go` | BLOQUEADO | Implementación GORM de ejemplo. |
| `internal/storage/cliente_gorm.go` | **TAREA CP1** | Impleméntela. |
| `internal/storage/venta_gorm.go` | **TAREA CP2** | Impleméntela. |
| `internal/services/errores.go` | BLOQUEADO | Errores de dominio. |
| `internal/services/medicamento_service.go` | BLOQUEADO | Service de ejemplo. |
| `internal/services/cliente_service.go` | **TAREA CP1** | Impleméntelo. |
| `internal/services/venta_service.go` | **TAREA CP2** | Las 5 reglas de negocio viven aquí. |
| `internal/handlers/respuestas.go` | BLOQUEADO | Helpers RespondJSON / RespondError. |
| `internal/handlers/routes.go` | BLOQUEADO | El contrato HTTP. |
| `internal/handlers/medicamento_handler.go` | BLOQUEADO | Handler de ejemplo (mapeo de errores). |
| `internal/handlers/cliente_handler.go` | **TAREA CP1** | Impleméntelo. |
| `internal/handlers/venta_handler.go` | **TAREA CP3** | Impleméntelo. |
| `cmd/api/main.go` | BLOQUEADO | Cableado completo + seeder. |
| `DECISIONES.md` | **TAREA CP1** | Máx. 10 líneas. |

Los archivos marcados como TAREA contienen comentarios `TODO` con las
indicaciones. **No cambie nombres ni firmas**: los tests compilan contra ellos.

## Primeros pasos

```bash
go mod tidy
go build ./...        # debe compilar desde el minuto cero
go test ./acceptance/... -v
```

**Importante:** hasta que complete los campos de los modelos (CP1), el paquete
`acceptance` no compila. Los errores de compilación le indican exactamente qué
campos faltan; las pantallas le dicen qué significan y qué tipo deben tener.

## Autoevaluación por checkpoint

```bash
go test ./acceptance/... -v -run TestCP1
go test ./acceptance/... -v -run TestCP2
go test ./acceptance/... -v -run TestCP3
```

Verde = puntos. Cuando un checkpoint esté en verde, llame al docente.

## Gate de entrega

```bash
go build ./... && go vet ./... && go test ./...
```

Un proyecto que no compila no se califica.

## Correr la API (opcional, para probar con Postman)

```bash
go run ./cmd/api
```

El seeder carga 4 medicamentos (una inactiva) la primera vez.
