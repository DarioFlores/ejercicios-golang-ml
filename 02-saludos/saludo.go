package saludo

func Saludo(nombre string) string {
	if nombre == "" {
		return "Hola extraño";
	}
	return "Hola " + nombre;
}
