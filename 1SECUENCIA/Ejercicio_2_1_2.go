// Ejercicio_2_1_2 (por profesores)
// "Dada una secuencia de letras del alfabeto que finaliza con
// la letra z contar cuantas consonantes hay en la secuencia"
ACCION Ejercicio_2_1_2 ES
AMBIENTE
	letras: secuencia de caracter
	v: caracter
	cont: entero
	vocales= ("a","e","i","o","u")
PROCESO
	ARR(letras)
	REPETIR
		AVZ(letras, ventana)
		SI v NO EN vocales ENTONCES
			cont := cont + 1
		FINSI
	HASTA QUE ventana = "z"
	ESC("La cantidad de consonantes en la secuencia es", cont)
FIN ACCION