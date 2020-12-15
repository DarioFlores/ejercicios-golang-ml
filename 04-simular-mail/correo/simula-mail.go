package correo

import "fmt"

func simulaEnvio(destinatario, asunto, mensaje string)  {
	fmt.Println("enviando correo…")
	fmt.Println("Destinatario:", destinatario);
	fmt.Println("Asunto:", asunto);
	fmt.Println("Mensaje:", mensaje);
}


func MailConDestinatarioSinAsuntoYMensaje(destinatario string)  {
	simulaEnvio(destinatario,"-", "-");
}

func MailConDestinatarioYAsuntoSinMensaje(destinatario, asunto string)  {
	simulaEnvio(destinatario,asunto, "-");
}

func MailConDestinatarioYAsuntoYMensaje(destinatario, asunto, mensaje string)  {
	simulaEnvio(destinatario,asunto, mensaje);
}