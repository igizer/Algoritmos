// Ejercicio_2_1_1
// "Dada una secuencia de letras que finaliza con una
// marca "*" contar cuantas letras "A" hay en la secuencia"
ACCION 2_1_1 ES

	AMBIENTE
		sec: secuencia de caracter
		v: caracter
		cont: entero

	PROCESO
		cont:=0
		ARR(sec)
		AVZ(sec, v)
		
		Mientras v<>"*" hacer
			si v = "A" entonces
				cont:= cont + 1
			finsi

			avz(sec, v)
		finmientras

		cerrar(sec)

		esc("la cantidad de letras 'A' son: ", cont)

finaccion