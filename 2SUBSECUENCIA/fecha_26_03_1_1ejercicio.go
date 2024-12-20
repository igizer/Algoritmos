// Dada una secuencia de caracteres que posee las ventas realizadas de un determinado producto.
// como máximo se pueden comprar 9. La información de las ventas en la secuencia sigue el
// siguiente formato:
// fechaVentaCantidadVendida
// Ejemplo:
// 010220236210320242150220247230220241
// La fecha siempre tiene 8 caracteres y un solo carácter para las unidades vendidas,
// no hay una marca que separe cada venta. 

// FILA 1
// Generar una secuencia de salida de entero, con todas las cantidades vendidas, en meses pares.

ACCION fila1 ES
	AMBIENTE
		ventas_e: secuencia de caracter
		ventas_s: secuencia de entero
		v, resguardo: caracter
		total, i: entero
		pares = ("0","2","4","6","8")
	
	PROCESO
		ARR(ventas_e)
		AVZ(ventas_e, v)
		CREAR(ventas_s)
		total:= 0
		MIENTRAS NFDS(ventas_e)
			PARA i:= 1 HASTA 3
				AVZ(ventas_e, v)
			FINPARA
			resguardo:= v
			PARA i:= 1 HASTA 5
				AVZ(ventas_e)
			FINPARA
			SI resguardo EN pares ENTONCES
				SEGUN v HACER
					= "0": GRABRAR(ventas_s, 0)
					= "1": GRABRAR(ventas_s, 1)
					= "2": GRABRAR(ventas_s, 2)
					= "3": GRABRAR(ventas_s, 3)
					= "4": GRABRAR(ventas_s, 4)
					= "5": GRABRAR(ventas_s, 5)
					= "6": GRABRAR(ventas_s, 6)
					= "7": GRABRAR(ventas_s, 7)
					= "8": GRABRAR(ventas_s, 8)
					= "9": GRABRAR(ventas_s, 9)
				FINSEGUN
			FINSI
			AVZ(ventas_s, v)
		FINMIENTRAS
		CERRAR(ventas_e)
		CERRAR(ventas_s)
FINACCION