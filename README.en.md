# Calculator in Go ![Logo de Go](https://go.dev/images/favicon-gopher.svg)

This project is a **simple calculator** developed in **Golang**, which allows you to perform basic operations between two numbers from the console:

- ➕ Addition
- ➖ Subtraction
- ✖️ Multiplication
- ➗ Division

## 📌 Description

The calculator is designed to run in the terminal and receive user input, processing operations through separate functions to keep the code clean and modular.

Version in [Spanish](./README.md)

## Project Page 📋

You can find the project page at [Golang Calculator](https://12345star.github.io/calculadora-golang).

## 🚀 Getting Started

To start using the calculator, follow these steps:

1. Clone the repository to your local machine:
   ```bash
   git clone https://github.com/12345star/calculadora-golang.git
   ```

2. Navigate to the project directory:
   ```bash
   cd calculadora-golang
   ```

## 📋 Prerequisites

Make sure you have **Go** installed on your system.

**GO** 1.18 or higher

* Official [Golang page](https://go.dev)
* [Testify](https://github.com/stretchr/testify)
* [mod tidy documentation](https://go.dev/ref/mod#go-mod-tidy)

## 📦 Dependencies

**Run the following commands in your terminal:**

```bash
go mod init calculadora
```

```bash
go mod tidy
```

```bash
go get github.com/stretchr/testify/assert
```

## ⚙️ Running the Calculator

1. Command to run the calculator:
   ```bash
   go run index.go <num1> <operator> <num2>
   ```
   1. Addition:
   ```bash
   go run index.go 5 + 3
   ```
   2. Subtraction:
   ```bash
   go run index.go 5 - 3
   ```
   3. Multiplication:
   ```bash
   go run index.go 5 * 3
   ```
   4. Division:
   ```bash
   go run index.go 6 / 3
   ```

---
## 🧪 Unit Tests

To run the tests, use the following command:

```bash
go test ./...
```

---
## 📄 License

This project is under the **MIT License**.

You can freely use, modify, and distribute it.

See the [MIT](./LICENSE) file.

---
## ✨ Author

Created with ❤️ by [@David](https://github.com/12345star)  
* **David Salas** - *Initial Work* - [David](https://github.com/12345star)
* **David Salas** - *Documentation* - [David](https://github.com/12345star)

---
## 📂 Project Structure

```plaintext
.
├── index.go                    # Main entry point
├── README.md                   # Project documentation
├── LICENSE                     # Project license
├── .gitignore                  # Files and folders ignored by Git
├── img/                        # Images folder
│   └── calculadora-golang.png  # Project structure image
├── go.mod                      # Go module
├── go.sum                      # Dependencies
├── model/
│   └── operaciones.go          # Mathematical operations logic
├── testing/
│   └── operaciones_test.go     # Unit tests
└── view/
    └── view.go                 # Console input/output interface

```
---
## 🚀 Project Structure

  ![Example structure](./img/calculadora-golang.png)