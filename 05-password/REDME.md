Crear un método que permita validar la seguridad de una contraseña retornando un error si no cumple con algún criterio especificado. La contraseña debe tener al menos 8 caracteres, un número y una mayúscula. 

TIP: usar regex para simplificar la detección de las reestricciones ;) 

``
regexp.MatchString(".*\\d.*", password)
``

``
regexp.MatchString(".*[A-Z].*", password)
``

Test que falle si tiene menos de 8 caracteres
Test que falle si no tiene un número
Test que falle si no tiene una mayúscula
Test que pase la validación y que la contraseña tenga más de un número

``
/^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[$@$!%*?&])[A-Za-z\d$@$!%*?&]{8,15}[^'\s]/
``