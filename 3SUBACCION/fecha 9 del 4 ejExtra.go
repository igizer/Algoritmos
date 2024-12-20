//Desarrollar una subacción que transforme cada dígito de un número, ingresado por el usuario,
//en los caracteres equivalentes. Además, se pide desarrollar un algoritmo, que use la subacción,
//y almacene estos caracteres en una secuencia de salida, al terminar un número separar con el
//carácter “#”. Agregue los parámetros que considere necesario .
//Ejemplo:
//dado el entero 355. Se guarda en la secuencia de salida: 
//“3” “5” “5” #


ACCION ejercicio_extra ES

	AMBIENTE
		digitos: secuencia de entero
		numeritos: entero
		
	funcion convertir(x:entero):caracter
	SEGUN
		= 0: convertir:= "0"
		= 1: convertir:= "1"
		= 2: convertir:= "2"
		= 3: convertir:= "3"
		= 4: convertir:= "4"
		= 5: convertir:= "5"
		= 6: convertir:= "6"
		= 7: convertir:= "7"
		= 8: convertir:= "8"
		= 9: convertir:= "9"
