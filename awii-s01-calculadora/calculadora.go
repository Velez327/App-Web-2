package main

import "fmt"

// func main() {

// 	fmt.Printf("==== CALCULADORA CIENTIFICA v1.0 ==== \n")

// 	var a = 50
// 	var b = 20

// 	suma := a + b
// 	resta := a - b
// 	multiplicacion := a * b
// 	divicion := float64(a) / float64(b)

// 	fmt.Printf("%d + %d = %d\n", a, b, suma)
// 	fmt.Printf("%d - %d = %d\n", a, b, resta)
// 	fmt.Printf("%d * %d = %d\n", a, b, multiplicacion)
// 	fmt.Printf("%d / %d = %.2f\n", a, b, divicion)

// }

func main() {
	for {
		fmt.Println("Infinite loop")
		fmt.Println("==== CALCULADORA CIENTÍFICA v1.0 ====")
		var numero1 float64
		fmt.Println("Digite el primer número:")
		fmt.Scanln(&numero1)
		var numero2 float64
		fmt.Println("Digite el segundo número:")
		fmt.Scanln(&numero2)
		var operacion string
		fmt.Println("Ingrese la operacion (+, -, *, /, ^, !):")
		fmt.Scanln(&operacion)

		for numero2 == 0 && operacion == "/" {
			fmt.Println("No se puede dividir por cero. Por favor, ingrese un número diferente de cero para el segundo número:")
			fmt.Scanln(&numero2)
		}

		for operacion != "+" && operacion != "-" && operacion != "*" && operacion != "/" && operacion != "^" && operacion != "!" {
			fmt.Println("Operación no válida. Por favor, ingrese una operación válida (+, -, *, /, ^, !):")
			fmt.Scanln(&operacion)
		}

		var calculo float64
		switch operacion {
		case "+":
			calculo = numero1 + numero2
		case "-":
			calculo = numero1 - numero2
		case "*":
			calculo = numero1 * numero2
		case "/":
			calculo = numero1 / numero2
		case "^":
			if numero2 == 0 {
				calculo = 1
			} else {
				calculo = 1
				for i := 1; i < int(numero2); i++ {
					calculo *= numero1
				}
			}
		case "!":
			if numero1 == 0 {
				calculo = 1
			} else {
				calculo = 1
				for i := 1; i <= int(numero1); i++ {
					calculo *= float64(i)
				}
			}
		}
		fmt.Println("El resultado es:", calculo)
	}
}
