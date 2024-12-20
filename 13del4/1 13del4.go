ACCION ej1 ES
	AMBIENTE
		entrada: secuencia de caracter
		salida: secuencia de entero
		v: caracter
	FUNCION convertir(x:caracter):entero
		SEGUN
			= "0": convertir:= 0
			= "1": convertir:= 1
			= "2": convertir:= 2
			= "3": convertir:= 3
			= "4": convertir:= 4
			= "5": convertir:= 5
			= "6": convertir:= 6
			= "7": convertir:= 7
			= "8": convertir:= 8
			= "9": convertir:= 9
		FINSEGUN
	FINFUNCION
	PROCESO
		ARR(entrada)
		AVZ(entrada, v)
		CREAR(salida)
		MIENTRAS NFDS(entrada) HACER
			MIENTRAS v <> " "
				GRABAR(salida, convertir(v))
				AVZ(entrada, v)
			FINMIENTRAS
			AVZ(entrada, v)
		FINMIENTRAS
		CERRAR(entrada)
		CERRAR(salida)
FINACCION
MALLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLL