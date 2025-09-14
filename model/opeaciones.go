package model

// ---------------------------------------------------------------------------------------------
// Paquete que contiene las operaciones básicas de una calculadora:
// suma, resta, multiplicación y división.
// ----------------------------------------------------------------------------------------------

import "fmt"

// Suma realiza la suma de dos números.
// Es metodo publico porque empieza con mayuscula
// Retorna la suma de a y b.
// Es metodo publico porque empieza con mayuscula
func Suma(a, b float64) float64 {
	return a + b
}

// Resta realiza la resta de dos números.
// Es metodo publico porque empieza con mayuscula
// Retorna la resta de a y b.
// Es metodo publico porque empieza con mayuscula
func Resta(a, b float64) float64 {
	return a - b
}

// Multiplicacion realiza la multiplicación de dos números.
// Es metodo publico porque empieza con mayuscula
// Retorna la multiplicación de a y b.
// Es metodo publico porque empieza con mayuscula
func Multiplicacion(a, b float64) float64 {
	return a * b
}

// Division realiza la división de dos números y maneja el caso de división por cero.
// Retorna un error si se intenta dividir por cero.
// Es metodo publico porque empieza con mayuscula
func Division(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("división por cero no es permitida")
	}
	return a / b, nil
}
