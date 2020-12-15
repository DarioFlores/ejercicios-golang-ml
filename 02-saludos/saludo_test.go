package saludo_test

import (
	saludo "github.com/DarioFlores/02-saludos"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestPorcentajeConTestify usa el paquete assert de testify para probar,
// esta dependencia la manejamos con go modules.
func TestPorcentajeConTestify(t *testing.T) {
	elSaludo := saludo.Saludo("Dario")
	assert.Equal(t, "Hola Dario", elSaludo, "Si le pasamos un nombre el saludo debe ser Hola <nombre>")
}

// TestPorcentajeConTestify usa el paquete assert de testify para probar,
// esta dependencia la manejamos con go modules.
func TestPorcentajeConTestify2(t *testing.T) {
	elSaludo := saludo.Saludo("")
	assert.Equal(t, "Hola extraño", elSaludo, "Si el nombre es un string vacio debe decir Hola extraño")
}