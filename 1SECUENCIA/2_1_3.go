// Ejercicio_2_1_3
// "Se dispone de una secuencia de caracteres y se desea obtener una secuencia de salida
// que resulte de copiar la secuencia de entrada, descartando el caracter $."
ACCION Ejercicio_2_1_3 ES
	AMBIENTE
		entrada, salida: secuencia de caracter
		v: caracter
	PROCESO
		ARR(entrada)
		AVZ(entrada, v)
		CREAR(salida)
		MIENTRAS NFDS(entrada) HACER
			SI v <> "$" ENTONCES
				GRABAR(salida, v)
			FINSI
			AVZ(entrada, v)
		FINMIENTRAS
		ESC("La secuencia fue realizada con éxito")
		CERRAR(entrada)
		CERRAR(salida)
FINACCION