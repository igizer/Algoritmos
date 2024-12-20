// "La secuencia de socios del problema anterior tiene el inconveniente de que los
// números están ordenados pero no son correlativos. Se desea generar una secuencia
// que contenga los números de socios que no figuran en la secuencia de socios.""
ACCION Ejercicio_2_1_5 ES
AMBIENTE
	sec, salida: secuencia de entero
	v1, v2, i: entero
PROCESO
	ARR(sec)
	AVZ(sec,v2)
	CREAR(salida)
	v1:= v2
	AVZ(sec,v2)
	MIENTRAS NFDS(sec) HACER
		SI v2 - v1 <> 1 HACER
			PARA i:= 1 HASTA (v2 - v1 - 1) HACER
				v1:= v1 + 1
				GRABAR (salida, v1)
			FINPARA
		FINSI
		v1 := v2
		AVZ(sec, v2)
	FINMIENTRAS
	CERRAR(sec)
	CERRAR(salida)
	ESCRIBIR("Nueva secuencia generada con exito");
FINACCION 