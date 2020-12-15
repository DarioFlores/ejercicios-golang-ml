package seguridad_test

import (
	seguridad "github.com/DarioFlores/05-password"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Test que falle si tiene menos de 8 caracteres
func TestFallaSiTieneMenosDe8Caracteres(t *testing.T) {
	validado, err := seguridad.ValidarPassword("ds2dDd2");
	assert.Equal(t, false, validado, "Debe ser false porque se envia un password con menos de 8 caracteres")
	assert.Equal(t, "Debe tener al menos 8 caracteres", err.Error(), "Debe hacer referencia de que se envia un password con menos de 8 caracteres")
}

// Test que falle si no tiene un número
func TestFallaSiNoTieneAlMenosUnaMayuscula(t *testing.T) {
	validado, err := seguridad.ValidarPassword("lalaa2lallalal");
	assert.Equal(t, false, validado, "Debe ser false porque debe contener al menos una mayuscula")
	assert.Equal(t, "Debe contener al menos una mayuscula", err.Error(), "Debe hacer referencia de que debe contener al menos una mayuscula")
}

// Test que falle si no tiene una mayúscula
func TestFallaSiNoTieneAlMenosUnNumero(t *testing.T) {
	validado, err := seguridad.ValidarPassword("lalaaLLlallalal");
	assert.Equal(t, false, validado, "Debe ser false porque debe contener al menos una numero")
	assert.Equal(t, "Debe contener al menos un numero", err.Error(), "Debe hacer referencia de que debe contener al menos una numero")
}

// Test que pase la validación y que la contraseña tenga más de un número
func TestPasarSiElPasswordEsValido(t *testing.T) {
	validado, _ := seguridad.ValidarPassword("lalaaLLl2allalal");
	assert.Equal(t, true, validado, "Debe ser true porque el password que se pasa es valido")
}