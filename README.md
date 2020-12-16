02 - Saludos

En el main crear una función que reciba un nombre y devuelva un saludo con la forma “hola <nombre>”. Si el nombre recibido es vacío debe devolver “hola extraño”. Incluir un test para cada caso utilizando testify.

03 - Una función dos retornos

Crear una función que retorne la suma y la resta de dos números enteros. Desde el main llamar a dicha función recuperando ambos resultados e imprimirlos por pantalla. Luego llamar nuevamente a la misma función pero esta vez ignorando el resultado de la resta

04 - Simula mail

Programa que simule el envío de un mail imprimiendo por pantalla el mensaje 

	enviando correo…
	destinatario: <destinatario>
	asunto: <asunto>
	mensaje: <mensaje>

Para los distintos casos que se van a plantear a continuación la idea es reutilizar la función de “simulación de envio de mail”

Escribir el código dentro de un paquete llamado correo y llamar desde el main 

a - Crear función no exportada llamada simulaEnvio que reciba por parámetros un destinatario, asunto y mensaje del tipo string e imprima por pantalla “enviando correo…” y debajo el destinatario, el asunto y el mensaje. En el caso de que el parámetro sea vacío imprimir un guión “-”

b - Función exportada que me permita enviar un mail con un destinatario (sin asunto y sin mensaje)

c - Función exportada que me permita enviar un mail con un destinatario y asunto (sin mensaje)

d - Función exportada que me permita enviar un mail con un destinatario, asunto y mensaje

05 - Password

Crear una función que permita validar la seguridad de una contraseña retornando un error si no cumple con algún criterio especificado. La contraseña debe tener al menos 8 caracteres, un número y una mayúscula. 

TIP: usar regex para simplificar la detección de las reestricciones ;) 

    regexp.MatchString(".*\\d.*", password) 
    regexp.MatchString(".*[A-Z].*", password)

Test que falle si tiene menos de 8 caracteres

Test que falle si no tiene un número

Test que falle si no tiene una mayúscula

Test que pase la validación y que la contraseña tenga más de un número

06 - Fechas

Crear un paquete util que me permita realizar distintas operaciones con fechas. Para el formateo de fechas no vale usar fmt, usar la función Format de time.Time

Tip pasar como se genera una fecha a partir de un string

a - Hacer una función que formatee una fecha en formato RFC3339
a1 - Escribir test que reciba una fecha y valide que esté en el formato esperado

b - Hacer una función que formatee una fecha en formato DD/MM/AAAA HH:mm
b1 - Escribir test que reciba la fecha de nacimiento de la persona y valide que esté en el formato esperado

c - Hacer una función que permita sumar 1 dia a una fecha y que devuelva el resultado en time.Time
c1 - Escribir test

07 - Persona

Crear una estructura persona y crear una función que permita cambiar nombre, una que reciba puntero y otra que no reciba puntero y mostrar la diferencia entre ambas. (https://play.golang.org/p/RVHbwW4x8oZ )

08 - Mascota

Crear una estructura persona y mascota. Una persona puede tener varias mascotas. 

a - Crear una función que me diga cuantas mascotas tiene una persona

b - Crear una función que me permita agregar una mascota a una persona

09 - Twitter

Escribir un programa que permita a una persona generar tweets. 

a - Crear una estructura Tweet en un paquete llamado datos

b - Crear una estructura Persona en un paquete llamado persona

c - Crear una función que permita a una persona twittear algo, dicho tweet se agrega a la colección de tweets de la persona

d - Dadas varias personas encontrar a la persona que más twitteó

e - Función que devuelva todos los tweets de una persona

f - Dada una persona, contar la cantidad de veces que se repite cada palabra en sus tweets. Devolver palabra y cantidad de veces que se repite (maps) (range)

g - Dada una persona, contar las primeras 3 palabras que más se repiten en sus tweets. Devolver palabra y cantidad de veces que se repite (maps) (range)

10 - Almacén de productos
     
a - A un almacén se pueden agregar productos, cada producto tiene un nombre (único), una descripción (opcional), cantidad de unidades disponibles y categorías (al menos una). Validar y devolver error en caso contrario

b - Se puede consultar la lista de todos los productos que están en el almacén

c - Del almacén se puede consultar el stock de un producto por nombre, si el producto no existe informar con un error.

d - Se puede retirar un producto del almacén (siempre y cuando haya	 stock)
     
     
TIP: si usás maps podés simplificar algunas consultas.

TIP 2: Recordá que si usás punteros las modificaciones a un objeto hechas usando el puntero se reflejan en todos los lugares donde hayas usado el mismo puntero.

