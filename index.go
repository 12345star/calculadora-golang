// ---------------------------------------------------------------------------------------------
// Programa.....: Calculadora Sencilla
// Archivo......: calculadora/index.go
// version......: 1.0
// Autor........: David Salas
// Fecha........: 2024-06-20
// Descripción..: Esta es una calculadora simple que realiza operaciones básicas como suma,
//                resta, multiplicación y división.
// Uso..........: El usuario debe proporcionar dos números y un operador (+, -, *, /)
//                como argumentos de línea de comandos.
// Comando......: go run calculadora/index.go 5 + 3
// ---------------------------------------------------------------------------------------------

package main

// go mod init calculadora
// go mod tidy
// go get github.com/stretchr/testify/assert
// go build -o calculadora
// ./calculadora 5 + 3
// go run calculadora/index.go 5 + 3
// ---------------------------------------------------------------------------------------------
// Importar los paquetes necesarios para que funcione adecuadmente el codigo de la calculadora
// como fmt para imprimir en consola, os para manejar argumentos de linea de comandos
// y strconv para convertir cadenas a números.
// ---------------------------------------------------------------------------------------------
import (
	"calculadora/model"
	"calculadora/view"
	"os"
	"strconv"
)

// ---------------------------------------------------------------------------------------------
// Función principal donde se ejecuta el programa.
// ---------------------------------------------------------------------------------------------
func main() {

	// variable
	var result float64
	// estructura
	type Calculadora struct {
		Num1     float64
		Num2     float64
		Operator string
	}

	// funcion a llamar para mostrar el mensaje de inicio
	view.MensajeIncio()

	// Verificar si se proporcionaron los argumentos correctos
	if len(os.Args) < 4 {
		view.MensajeErrorArgumentos()
		return
	}

	num1, err1 := strconv.ParseFloat(os.Args[1], 64) // primer numero
	operator := os.Args[2]                           // operador(+, -, *, /)
	num2, err2 := strconv.ParseFloat(os.Args[3], 64) // segundo numero

	// Crear una instancia de Calculadora
	var calculadora Calculadora
	calculadora.Num1 = num1
	calculadora.Num2 = num2
	calculadora.Operator = operator
	operator = calculadora.Operator

	// Verificar si los argumentos son válidos

	if err1 != nil || err2 != nil {
		view.MensajeErrorArgumentos()
		return
	}

	// Realizar la operación según el operador proporcionado
	switch operator {
	case "+": // suma
		result = model.Suma(calculadora.Num1, calculadora.Num2)
	case "-": // resta
		result = model.Resta(num1, num2)
	case "*": // multiplicacion
		result = model.Multiplicacion(num1, num2)
	case "/": // division
		if num2 == 0 {
			view.MensajeErrorDivisionPorCero()
			return
		} else {
			result, err2 = model.Division(num1, num2)
			if err2 != nil {
				view.MensajeErrorDivisionPorCero()
				return
			}
		}
	default:
		view.MensajeErrorOperadorInvalido()
		return
	}

	// Mostrar el resultado
	view.MensajeResultado(result)
}
