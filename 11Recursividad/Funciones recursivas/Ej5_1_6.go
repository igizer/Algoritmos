Determine recursivamente si un número dado es par o impar, usando sólo la operación resta.


funcion determinar(n: entero) : logico ES
	
	
	si n = 0 entonces
		determinar:= verdadero    //par
	sino
		si n = 1 entonces
			determinar:= falso    //impar
		sino
			determinar:= determinar(n-2)
		fs
	fs
F_funcion.



prueba de escritorio

determinar(6) = determinar(4) =determinar(2) = determinar(0)


------------propongo esta versión----------

funcion paridad(n:entero):logico es
	segun n hacer
	
		=0: paridad:= verdadero
		=1: paridad:= falso
		De otro modo: paridad:= paridad(n-2)

	finsegun
finfuncion