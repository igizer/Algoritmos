// Dada una secuencia de caracteres que posee las ventas realizadas de un determinado producto.
// como máximo se pueden comprar 9. La información de las ventas en la secuencia sigue el
// siguiente formato:
// fechaVentaCantidadVendida
// Ejemplo:
// 010220236210320242150220247230220241
// La fecha siempre tiene 8 caracteres y un solo carácter para las unidades vendidas,
// no hay una marca que separe cada venta. 

//FILA 2
//Contabilizar unidades vendidas para un mes en particular (que ingresa un usuario).

ACCION subse_2 ES

























	AMBIENTE
		sec: secuencia de caracter
		v: caracter
		productos, mes, resg_mes, i: entero

		FUNCION convertir(x:caracter):entero
			SEGUN
				= "0" : convertir:= 0
				= "1" : convertir:= 1
				= "2" : convertir:= 2
				= "3" : convertir:= 3
				= "4" : convertir:= 4
				= "5" : convertir:= 5
				= "6" : convertir:= 6
				= "7" : convertir:= 7
				= "8" : convertir:= 8
				= "9" : convertir:= 9
			FINSEGUN
		FINFUNCION

	PROCESO

		ARR(sec); AVZ(sec, v)
		resg_mes:= 0
		ESC("A continuación escribe en números un mes para contabilizar unidades vendidas")
		LEER(mes)

		MIENTRAS NFDS(sec) HACER

			productos:= 0

				AVZ(sec,v)
				AVZ(sec,v)
				resg_mes:= convertir(v) * 10
				AVZ(sec,v)
				resg_mes:= resg_mes + convertir(v)

		FINMIENTRAS

FINACCION






























ACCION fila2 ES
	AMBIENTE
		ventas: secuencia de caracter
		v: caracter
		digito1, digito2, i, mes, cantidad: entero
	FUNCION transformar(x:caracter):entero
		SEGUN x HACER
			= "0": transformar:=0
			= "1": transformar:=1
			= "2": transformar:=2
			= "3": transformar:=3
			= "4": transformar:=4
			= "5": transformar:=5
			= "6": transformar:=6
			= "7": transformar:=7
			= "8": transformar:=8
			= "9": transformar:=9
		FINSEGUN
	FINFUNCION
	PROCESO
		ARR(ventas)
		AVZ(ventas, v)
		cantidad:= 0
		ESC("Ingrese un mes para contabilizar sus unidades vendidas")
		LEER(mes)
		MIENTRAS NFDS(ventas) HACER
			AVZ(ventas, v)
			AVZ(ventas, v)
			digito1:= transformar(v)*10
			AVZ(ventas, v)
			digito2:= transformar(v)
			PARA i:= 1 HASTA 5 HACER
				AVZ(ventas, v)
			FINPARA
			SI (digito1+digito2) = mes ENTONCES
				cantidad_mes:= cantidad_mes + transformar(v)
			FINSI
			AVZ(ventas, v)
		FINMIENTRAS
		CERRAR(sec)
		ESC("La cantidad de unidades vendidas para el mes ", mes, "en total es de", cantidad_mes)
FINACCION


