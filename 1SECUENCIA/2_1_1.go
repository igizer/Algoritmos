// Ejercicio_2_1_1
// "Dada una secuencia de letras que finaliza con una
// marca "*" contar cuantas letras "A" hay en la secuencia"
ACCION ejercicio_2_1_1 ES

	AMBIENTE

		sec: secuencia de caracter
		v: caracter
		cont: entero

	PROCESO

		ARR(sec)
		AVZ(sec, v)
		cont:= 0

		MIENTRAS v <> "*" HACER

			SI v = "a" ENTONCES
				cont:= cont + 1
			FINSI

			AVZ(sec, v)
		
		FINMIENTRAS

		CERRAR(sec)

		ESC("La cantidad de a es ", cont)

FINACCION





















ACCION Ejercicio_2_1_1 ESC
	AMBIENTE
		sec: secuencia de caracter
		v: caracter
		cont: entero
		A_Decir = "La cantidad de letra a son: "

	PROCESO
		ARR(sec)
		AVZ(sec, v)
		cont:= 0

		MIENTRAS v <> "*" HACER
			
			SI v = "a" ENTONCES
				cont:= cont + 1
			FINSI

			AVZ(sec, v)
		
		FINMIENTRAS

		ESC("La cantidad de letras 'a' son: ", cont)

		CERRAR(sec)
FINACCION














ACCION
	AMBIENTE
		sec: secuencia de caracter
		v: caracter
		cont: entero
		
	PROCESO
		ARR(sec);
		AVZ(sec, v);
		cont:= 0;
		MIENTRAS v <> "*" HACER
			SI v = "A" ENTONCES
				cont:= cont + 1;
			FINSI
			AVZ(sec, v);
		FINMIENTRAS
		ESC("Hay ", cont, " letras A en la secuencia.");
		CERRAR(sec);
FINACCION