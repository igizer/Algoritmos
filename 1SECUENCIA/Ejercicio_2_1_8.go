// Teniendo en cuenta el ejercicio anterior, se solicita  que la
// secuencia de salida sea una secuencia de caracteres y los CUIT
// se separen unos de otros con el caracter "-".
ACCION Ejercicio_2_1_8 ES
AMBIENTE
	entrada, ordenado: secuencia de entero
	salida: secuencia de caracter
	v, v2: entero
	rango = (0, 1, 2, 3)
PROCESO
	ARR(entrada); CREAR(ordenado);
	AVZ(entrada, v);
	MIENTRAS v <> 0 HACER
		num:= v MOD 10;
		SI ((v MOD 100) DIV 10) EN rango ENTONCES
			GRABAR(ordenado, v);
			GRABAR("-")
		FINSI
		AVZ(entrada, v);
	FIN MIENTRAS
	CERRAR(entrada);
	CERRAR(ordenado);
	ARR(ordenado);
	AVZ(ordenado,v);
	CREAR(salida);
	GRABAR(salida,v);
	MIENTRAS NDFS(ordenado) HACER
		AVZ(ordenado,v)
		GRABAR(salida,"-");
		GRABAR(salida,v);
		
	FINMIENTRAS
FINACCION