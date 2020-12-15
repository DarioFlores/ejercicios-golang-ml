package seguridad

import (
	"errors"
	"fmt"
	"regexp"
)

func ValidarPassword(password string) (bool, error) {
	v1, err := regexp.MatchString(".*\\d.*", password)
	if !v1 {
		fmt.Println(err)
		return v1, errors.New("Debe contener al menos un numero");
	}
	v2, err2 := regexp.MatchString(".*[A-Z].*", password)
	if !v2 {
		fmt.Println(err2)
		return false, errors.New("Debe contener al menos una mayuscula");
	}

	longitud := len(password)
	if longitud < 8 {
		return false, errors.New("Debe tener al menos 8 caracteres");
	}
	return true, nil
}