package view

// ---------------------------------------------------------------------------------------------
// Paquete view que contiene funciones para mostrar mensajes en la consola.
// ---------------------------------------------------------------------------------------------
import "fmt"

// ---------------------------------------------------------------------------------------------
// Función para mostrar el mensaje de inicio de la calculadora.
// ---------------------------------------------------------------------------------------------
func MensajeIncio() {
	fmt.Println("----------------------------------------------")
	fmt.Println("Calculadora Sencilla que realiza operaciones")
	fmt.Println("usand el lenguaje Go.")
	fmt.Println("----------------------------------------------")
}

// ---------------------------------------------------------------------------------------------
// Función para mostrar mensajes de error en la división por cero.
// ---------------------------------------------------------------------------------------------
func MensajeErrorDivisionPorCero() {
	fmt.Println("Erro: Division por zero no es permitida.")
}

// ---------------------------------------------------------------------------------------------
// Función para mostrar mensajes de error en los argumentos.
// ---------------------------------------------------------------------------------------------
func MensajeErrorArgumentos() {
	fmt.Println("Erro: Ambos argumentos deben números válidos.")
}

// ---------------------------------------------------------------------------------------------
// Función para mostrar mensajes de error en el tipo de operador.
// ---------------------------------------------------------------------------------------------
func MensajeErrorOperadorInvalido() {
	fmt.Println("Erro: Operador inválido. Use +, -, * ou /.")
}

// ---------------------------------------------------------------------------------------------
// Función para mostrar el resultado de la operación.
// ---------------------------------------------------------------------------------------------
func MensajeResultado(result float64) {
	fmt.Printf("Resultado: %.2f\n", result)
}
