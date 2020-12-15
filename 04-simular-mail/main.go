package main

import (
	"fmt"
	"github.com/DarioFlores/04-simular-mail/correo"
)

func main() {
	fmt.Println()
	correo.MailConDestinatarioSinAsuntoYMensaje("Poche")
	fmt.Println()
	correo.MailConDestinatarioYAsuntoSinMensaje("Gisella", "Cara de culo!")
	fmt.Println()
	correo.MailConDestinatarioYAsuntoYMensaje("Dario", "Carta de amor", "Hola bebe")
}
