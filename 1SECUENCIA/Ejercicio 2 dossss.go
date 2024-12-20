// Ejercicio_2_1_2 (por mí, duda en que significa ventana, si no es lo mismo que v)?)
// "Dada una secuencia de letras del alfabeto que finaliza con
// la letra z contar cuantas consonantes hay en la secuencia"
ACCION Ejercicio_2_1_2 ES
AMBIENTE
	sec: secuencia de caracter
	v: caracter
	cont: entero
	vocales= ("a","e","i","o","u")
PROCESO
	ARR(sec)
	REPETIR
		AVZ(sec,v)
		SI v NO EN vocales ENTONCES
			cont :=  cont + 1
		FINSI
		AVZ(sec,v)
	HASTA QUE v = "z"
	ESC("La cantidad de consonantes que hay en la secuencia es de", cont)
FINACCION