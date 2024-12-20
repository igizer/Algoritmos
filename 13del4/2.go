ACCION ej2 ES
	AMBIENTE
		sec1, sec2: secuencia de caracter
		v1, v2: caracter
		edad, i: entero
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
		ARR(sec1)
		ARR(sec2)
		AVZ(sec1)
		AVZ(sec2)
		MIENTRAS (NFDS(sec1)) y (NFDS(sec2)) HACER
			edad:= convertir(v1)*10
			AVZ(sec1, v1)
			edad:= edad + convertir(v1)
			AVZ(sec1, v1)
			SI (edad>=18) y (edad<=30) ENTONCES
				MIENTRAS v1<>"," HACER
					ESC(v1)
					AVZ(sec1,v1)
				FINMIENTRAS
				PARA i:= 1 HASTA 14 HACER
					AVZ(sec1, v1)
				FINPARA
				ESC("/")
				PARA i:= 1 HASTA 4
					ESC(v2)
					AVZ(sec2, v2)
				FINPARA
				ESC("*")
			SINO
				MIENTRAS v1<>"," HACER
					AVZ(v1)
				FINMIENTRAS
				PARA i:= 1 HASTA 14 HACER
					AVZ(sec1, v1)
				FINPARA
				PARA i:= 1 HASTA 4
					AVZ(sec2, v2)
				FINPARA
		FINMIENTRAS
		CERRAR(sec1)
		CERRAR(sec2)
FINACCION
			

			