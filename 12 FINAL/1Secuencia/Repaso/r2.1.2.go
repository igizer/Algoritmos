// Ejercicio_2_1_2 (por profesores)
// "Dada una secuencia de letras del alfabeto que finaliza con
// la letra z contar cuantas consonantes hay en la secuencia"

ACCION 2_1_2_repaso ES

	AMBIENTE
		sec: secuencia de caracter
		c: caracter
		cont: entero
		vocales= ("a", "e", "i", "o", "u")

	PROCESO
		cont:= 0
		arr(sec)
		avz(sec, v)

		repetir 
			si v no en vocales entonces
				cont:= cont+1
			finsi
			avz(sec,c)
		hasta v="z"

		esc("la cantidad de consonantes son: ", cont)
		cerrar(sec)
finaccion





ACCION 2_1_2_original ES

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
