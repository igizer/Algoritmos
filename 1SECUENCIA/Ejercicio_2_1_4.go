// "Se dispone de una secuencia de números de socios, y se
// desea saber la cantidad de socios que están registrados."
ACCION Ej_2_1_4
AMBIENTE
	sec: secuencia de entero
	v, cont: entero
PROCESO
	ARR(sec)
	AVZ(sec,v)
	cont := 0
	MIENTRAS NFDS(sec) HACER
		cont = cont + 1
		AVZ(sec,v)
	FINMIENTRAS
	ESC("La cantidad de socios que están registrados son", cont)
	CERRAR(sec)
FINACCION