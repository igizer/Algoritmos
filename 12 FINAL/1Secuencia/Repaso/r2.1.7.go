//Se tiene una secuencia de enteros que contiene todos los CUIT
//de los empleados de una empresa, la misma termina con el CUIT 0.
//Para evitar largas esperas en los días de pago, la empresa
//necesita organizarlos de acuerdo al último dígito de su documento,
//generar una secuencia con los CUIT de las personas que su número
//de documento termine con 0, 1, 2 o 3.

//los cuit tienen 11 digitos XX-12.345.678-X en este caso el 8 es el digito que necesito analizar

accion repaso2_1_7 es
	ambiente
		sec, sal: secuencia de enteros // secuencia de entrada y secuencia de salida
		v, numero: entero // v: CUIT actual; numero: dígito que se analiza
		rango = (0, 1, 2, 3) // conjunto de dígitos permitidos
	proceso
		crear(sal) // inicializar secuencia de salida
		arr(sec) // abrir secuencia de entrada
		avz(sec, v) // leer primer CUIT

		mientras v <> 0 hacer // mientras no sea el fin de datos
			numero := ((v div 10) mod 10) // extraer el anteúltimo dígito del CUIT
			si numero en rango entonces // si el dígito está entre 0 y 3
				grabar(sal, v) // guardar CUIT en la secuencia de salida
			finsi
			avz(sec, v) // avanzar a siguiente CUIT
		finmientras
	
		cerrar(sec) // cerrar secuencia de entrada
		cerrar(sal) // cerrar secuencia de salida

		esc("secuencia generada con exito.") // mensaje final
fin accion

//----------------respuesta vieja------

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