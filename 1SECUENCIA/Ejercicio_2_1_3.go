// Ejercicio_2_1_3
// "Se dispone de una secuencia de caracteres y se desea obtener una secuencia de salida
// que resulte de copiar la secuencia de entrada, descartando el caracter $."
ACCION Ejercicio_2_1_3 ES
AMBIENTE
	sec, salida: secuencia de caracter
	v: caracter
PROCESO
	ARR(sec)
	AVZ(sec,v)
	CREAR(salida)
	MIENTRAS NFDS(sec) ENTONCES
		SI v <> "$" ENTONCES
			GRABAR(salida,v)
		FINSI
		AVZ(sec,v)
	FINMIENTRAS
	ESC("La secuencia de salida sin $ es", salida)
	CERRAR(sec)
	CERRAR(salida)
FINACCION