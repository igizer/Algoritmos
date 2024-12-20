//Se tiene una secuencia de enteros que contiene todos los CUIT
//de los empleados de una empresa, la misma termina con el CUIT 0.
//Para evitar largas esperas en los días de pago, la empresa
//necesita organizarlos de acuerdo al último dígito de su documento,
//generar una secuencia con los CUIT de las personas que su número
//de documento termine con 0, 1, 2 o 3.
ACCION
AMBIENTE
	entrada, salida: secuencia de entero
	v: entero
	rango = (0, 1, 2, 3)
PROCESO
	ARR(entrada); CREAR(salida);
	AVZ(entrada, v);
	MIENTRAS v <> 0 HACER
		SI ((v MOD 100) DIV 10) EN RANGO ENTONCES
			GRABAR(salida, v);
		FINSI
		AVZ(sec, v);
	FIN MIENTRAS
	CERRAR(entrada); CERRAR(salida);
	ESCRIBIR("Nueva secuencia generada con exito");
FINACCION