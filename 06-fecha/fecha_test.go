package fechas_test

import (
	fechas "github.com/DarioFlores/06-fecha"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Test que falle si tiene menos de 8 caracteres
func TestFallaSiTieneMenosDe8Caracteres(t *testing.T) {
	validado, err := fechas.FormatoRFC3339("2020-01-02T15:04:05-03:00")
	assert.Equal(t, "2020-01-02 15:04:05 -0300 -03", validado.String(), "Debe ser false porque se envia un password con menos de 8 caracteres")
	assert.Equal(t, "Debe tener al menos 8 caracteres", err.Error(), "Debe hacer referencia de que se envia un password con menos de 8 caracteres")
}
