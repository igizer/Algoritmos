//Se tiene una secuencia de enteros que contiene los números de teléfono de los clientes
//de una empresa,cada número tiene 10 dígitos. La empresa necesita organizar los clientes
//de acuerdo al último dígito de su número telefónico, generar una secuencia con los números
//de las personas que termine con 5, 6 o 7.
ACCION fecha19_03_ej1 ES
	AMBIENTE
		entrada, salida: secuencia de entero
		v, u_d, i: entero
		terminados = (5, 6, 7)
	PROCESO
		ARR(entrada)
		AVZ(entrada, v)
		CREAR(salida)
		MIENTRAS NFDS(entrada) HACER
			u_d:= v MOD 10
			SI u_d EN terminados ENTONCES
				GRABAR(salida, v)
			FINSI
			AVZ(entrada, v)
		FINMIENTRAS
		ESCRIBIR("Los clientes fueron organizados con éxito")
		CERRAR(entrada)
		CERRAR(salida)
	FINPROCESO
