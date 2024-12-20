// Ejercicio_2_1_2 (por profesores)
// "Dada una secuencia de letras del alfabeto que finaliza con
// la letra z contar cuantas consonantes hay en la secuencia"
ACCION 2_1_2 ES

	AMBIENTE

		sec: secuencia de caracter
		v: caracter
		cont, edad: entero
		vocales= {"a", "e", "i", "o", "u"}

	PROCESO

		ARR(sec)
		cont:= 0

		REPETIR
			AVZ(sec, v)
			SI v NO EN vocales ENTONCES
				cont:= cont + 1
			FINSI
		HASTA v = "z"

		ESC("Las consonantes en total son ", cont)
		CERRAR(sec)

FINACCION

















REPETIR

	TAL COSA

HASTA v  

























ACCION
	AMBIENTE
		sec: secuencia de caracter
		v: caracter
		cont: entero
		vocales= ("a", "e", "i", "o", "u")
	PROCESO
		ARR(sec)
		AVZ(sec, v)
		cont:= 0
		REPETIR
			SI v NO EN vocales ENTONCES
				cont:= cont + 1
			FINSI
			AVZ(sec, v)
		HASTA QUE v = "Z"
		ESC("Las consonantes en total son ", cont)
		CERRAR(sec)
FINACCION

