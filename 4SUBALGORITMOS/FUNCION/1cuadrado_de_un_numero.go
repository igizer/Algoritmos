//Crear un algoritmo que retorne el cuadrado de un número
ACCION cuadrado_de_un_numero ES
	AMBIENTE
		numero:entero
	FUNCION cuadrado(x:entero):entero
		cuadrado:= x**2
	FINFUNCION
	PROCESO
		ESC("Este algoritmo calcula el cuadrado de un número")
		ESC("Escribe a continuación un número para calcular su cuadrado")
		LEER(numero)
		ESC("El cuadrado del número ", numero, " es: ", cuadrado(numero))





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