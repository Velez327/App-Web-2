package main

import (
	"fmt"
	"semana03-taller-relaciones/internal/cafeteria"
)

func main() {
	c := cafeteria.Cliente{ID: 1, Nombre: "Ana"}
	fmt.Println(c)

	var repo cafeteria.Repository = cafeteria.NewRepoMemoria()
	repo.GuardarC(cafeteria.Cliente{ID: 1, Nombre: "Ana", Carrera: "TI", Saldo: 15.00, Telefono: 1234567890, Email: "ana@example.com"})
	repo.GuardarC(cafeteria.Cliente{ID: 2, Nombre: "Sam", Carrera: "TI", Saldo: 30.00, Telefono: 1561156165, Email: "sam@example.com"})
	repo.GuardarP(cafeteria.Producto{ID: 1, Nombre: "Café", Precio: 1.25, Stock: 20, Categoria: "Bebidas", Detalles: "Bebida caliente"})
	repo.GuardarP(cafeteria.Producto{ID: 2, Nombre: "Capuchino", Precio: 1.75, Stock: 30, Categoria: "Bebidas", Detalles: "Bebida con leche vaporizada"})
	repo.GuardarP(cafeteria.Producto{ID: 3, Nombre: "Expreso", Precio: 1.55, Stock: 40, Categoria: "Bebidas", Detalles: "Bebida fuerte de café"})

	// Buscamos un producto existente
	p, err := repo.BuscarPorIDC(101)
	if err != nil {
		fmt.Printf("Error al buscar cliente: %s\n", err.Error())
		return
	}
	fmt.Printf("\nEncontrado: %s\n", p.Nombre)

	_, err = repo.BuscarPorIDC(999)
	if err != nil {
		fmt.Printf("Error al buscar cliente: \n", err)
	}
}

// 1. ¿Tuviste que poner Cliente, Producto y Pedido en el mismo paquete? ¿Por qué sí o por qué no?
//Josselyn: Sí, porque si los separáramos en paquetes distintos, tendríamos que importar ambos paquetes en el paquete de Pedido,
// lo que complicaría la organización del código y crearía dependencias circulares.
//Pablo: Además, al estar relacionados, es más sencillo mantenerlos juntos para facilitar la gestión de las relaciones entre
// ellos.

// 2. ¿Qué problema aparecería si intentaras separar Producto en un paquete aparte cuando Pedido lo tiene anidado?
//Josselyn: Si Producto está en un paquete separado, el paquete de Pedido tendría que importar el paquete de Producto para usar la
// estructura de Producto, al mismo tiempo, si el paquete de Producto necesita referenciar a Pedido y importar el paquete
// de Pedido, además, esta separación dificultaría la gestión de las relaciones entre las entidades.
//Pablo: crearía una dependencia circular, lo que complicaría la compilación y el mantenimiento del código.
// Además, al estar relacionados, es más practico mantenerlos juntos para facilitar la gestión de las
// relaciones entre ellos.

// 3. Comparando con el Día A (donde usamos IDs): ¿qué ventaja tiene el modelo con IDs para organizar el código en paquetes?
//Josselyn: El modelo con IDs permite una mayor modularidad y separación de responsabilidades y se ve mas limpio al momento
// y organizado al momento de codificar
//Pablo: la ventaja del modelo con IDs es que permite una mayor modularidad y separación de responsabilidades,
// lo que facilita la organización del código en paquetes y la gestión de dependencias entre ellos.
