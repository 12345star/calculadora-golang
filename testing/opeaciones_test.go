package testing_test

// ---------------------------------------------------------------------------------------------
// Paquete que contiene las operaciones básicas de una calculadora:
// suma, resta, multiplicación y división.
// ----------------------------------------------------------------------------------------------

import (
	"calculadora/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

// -------------------------------------------------------
// TestsSuma prueba la función Suma del paquete model.
// Verifica que la suma de 2 y 3 sea igual a 5.
// Es metodo publico porque empieza con mayuscula
// Usa la libreria testify para las aserciones
// Retorna un error si la aserción falla.
// -------------------------------------------------------
func testsSuma(t *testing.T) {
	result := model.Suma(2, 3)
	assert.Equal(t, 5.0, result, "La suma de 2 y 3 debe ser 5")
}

// -------------------------------------------------------------------------------------
// TestSuma es un envoltorio para llamar a testsSuma.
// Permite ejecutar testsSuma como una subprueba.
// -------------------------------------------------------------------------------------
func TestSuma(t *testing.T) {
	t.Run("TestSuma", testsSuma)
}

// -------------------------------------------------------------------------------------
// TestResta prueba la función Resta del paquete model.
// Verifica que la resta de 5 y 3 sea igual a 2.
// Usa la libreria testify para las aserciones
// Retorna un error si la aserción falla.
// ------------------------------------------------------------------------------------
func TestResta(t *testing.T) {
	t.Run("TestResta", func(t *testing.T) {
		result := model.Resta(5, 3)
		assert.Equal(t, 2.0, result, "La resta de 5 y 3 debe ser 2")
	})
}

// -------------------------------------------------------------------------------------
// TestMultiplicacion prueba la función Multiplicacion del paquete model.
// Verifica que la multiplicación de 2 y 3 sea igual a 6.
// Usa la libreria testify para las aserciones
// Retorna un error si la aserción falla.
// ------------------------------------------------------------------------------------
func TestMultiplicacion(t *testing.T) {
	t.Run("TestMultiplicacion", func(t *testing.T) {
		result := model.Multiplicacion(2, 3)
		assert.Equal(t, 6.0, result, "La multiplicación de 2 y 3 debe ser 6")
	})
}

// -------------------------------------------------------------------------------------
// TestDivision prueba la función Division del paquete model.
// Verifica que la división de 6 entre 3 sea igual a 2.
// También verifica que dividir por cero retorne un error.
// Usa la libreria testify para las aserciones
// Retorna un error si alguna aserción falla.
// ------------------------------------------------------------------------------------
func TestDivision(t *testing.T) {
	t.Run("TestDivision", func(t *testing.T) {
		result, err := model.Division(6, 3)
		assert.Nil(t, err, "No debe haber error en la división de 6 entre 3")
		assert.Equal(t, 2.0, result, "La división de 6 entre 3 debe ser 2")
	})

	t.Run("TestDivisionPorCero", func(t *testing.T) {
		_, err := model.Division(6, 0)
		assert.NotNil(t, err, "Debe haber un error al dividir por cero")
	})
}

// -------------------------------------------------------------------------------------
// TestSumaParametrizada prueba la función Suma del paquete model
// con múltiples casos de prueba utilizando una tabla de pruebas.
// Cada caso de prueba incluye dos números de entrada y el resultado esperado.
// Usa la libreria testify para las aserciones
// Retorna un error si alguna aserción falla.
// ----------------------------------------------------------------------------------------------
func TestSumaParametrizada(t *testing.T) {
	tests := []struct {
		name     string
		a        float64
		b        float64
		expected float64
	}{
		{"SumaPositivos", 2, 3, 5},
		{"SumaNegativos", -2, -3, -5},
		{"SumaMixtos", -2, 3, 1},
		{"SumaCero", 0, 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := model.Suma(tt.a, tt.b)
			assert.Equal(t, tt.expected, result, "La suma de %.2f y %.2f debe ser %.2f", tt.a, tt.b, tt.expected)
		})
	}
}

// -------------------------------------------------------------------------------------
// TestRestaParametrizada prueba la función Resta del paquete model
// con múltiples casos de prueba utilizando una tabla de pruebas.
// Cada caso de prueba incluye dos números de entrada y el resultado esperado.
// Usa la libreria testify para las aserciones
// Retorna un error si alguna aserción falla.
// ----------------------------------------------------------------------------------------------
func TestRestaParametrizada(t *testing.T) {
	tests := []struct {
		name     string
		a        float64
		b        float64
		expected float64
	}{
		{"RestaPositivos", 5, 3, 2},
		{"RestaNegativos", -5, -3, -2},
		{"RestaMixtos", -5, 3, -8},
		{"RestaCero", 0, 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := model.Resta(tt.a, tt.b)
			assert.Equal(t, tt.expected, result, "La resta de %.2f y %.2f debe ser %.2f", tt.a, tt.b, tt.expected)
		})
	}
}

// -------------------------------------------------------------------------------------
// TestMultiplicacionParametrizada prueba la función Multiplicacion del paquete model
// con múltiples casos de prueba utilizando una tabla de pruebas.
// Cada caso de prueba incluye dos números de entrada y el resultado esperado.
// Usa la libreria testify para las aserciones
// Retorna un error si alguna aserción falla.
// ----------------------------------------------------------------------------------------------
func TestMultiplicacionParametrizada(t *testing.T) {
	tests := []struct {
		name     string
		a        float64
		b        float64
		expected float64
	}{
		{"MultiplicacionPositivos", 2, 3, 6},
		{"MultiplicacionNegativos", -2, -3, 6},
		{"MultiplicacionMixtos", -2, 3, -6},
		{"MultiplicacionCero", 0, 5, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := model.Multiplicacion(tt.a, tt.b)
			assert.Equal(t, tt.expected, result, "La multiplicación de %.2f y %.2f debe ser %.2f", tt.a, tt.b, tt.expected)
		})
	}
}

// -------------------------------------------------------------------------------------
// TestDivisionParametrizada prueba la función Division del paquete model
// con múltiples casos de prueba utilizando una tabla de pruebas.
// Cada caso de prueba incluye dos números de entrada y el resultado esperado.
// Usa la libreria testify para las aserciones
// Retorna un error si alguna aserción falla.
// ----------------------------------------------------------------------------------------------
func TestDivisionParametrizada(t *testing.T) {
	tests := []struct {
		name        string
		a           float64
		b           float64
		expected    float64
		expectError bool
	}{
		{"DivisionPositivos", 6, 3, 2, false},
		{"DivisionNegativos", -6, -3, 2, false},
		{"DivisionMixtos", -6, 3, -2, false},
		{"DivisionPorCero", 5, 0, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := model.Division(tt.a, tt.b)
			if tt.expectError {
				assert.NotNil(t, err, "Se esperaba un error al dividir %.2f entre %.2f", tt.a, tt.b)
			} else {
				assert.Nil(t, err, "No se esperaba un error al dividir %.2f entre %.2f", tt.a, tt.b)
				assert.Equal(t, tt.expected, result, "La división de %.2f entre %.2f debe ser %.2f", tt.a, tt.b, tt.expected)
			}
		})
	}
}
