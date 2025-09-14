# 🧮 Calculadora en Go

Este proyecto es una **calculadora sencilla** desarrollada en **Golang**, que permite realizar operaciones básicas entre dos números desde la consola:

- ➕ Suma
- ➖ Resta
- ✖️ Multiplicación
- ➗ División

---

## 📌 Descripción

La calculadora está diseñada para ejecutarse en la terminal y recibir la entrada del usuario, procesando las operaciones mediante funciones separadas para mantener un código limpio y modular.

---

## 📂 Estructura del proyecto

```plaintext
.
├── index.go                # Punto de entrada principal
├── go.mod                  # Módulo de Go
├── go.sum                  # Dependencias
├── model/
│   └── operaciones.go      # Lógica de operaciones matemáticas
├── testing/
│   └── operaciones_test.go # Pruebas unitarias
└── view/
    └── view.go              # Interfaz de entrada/salida en consola


