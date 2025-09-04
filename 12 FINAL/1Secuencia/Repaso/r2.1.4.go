// "Se dispone de una secuencia de números de socios, y se
// desea saber la cantidad de socios que están registrados."

accion r2_1_4_ es
	ambiente
		sec: secuencia de entero
		e, cont: entero
	
	proceso
		arr(sec)
		avz(sec, e)
		cont:= 0

		mientras nfds(sec) hacer
			cont:= cont+1
			avz(sec)
		finmientras

		esc("El total de socios registrados son: ", cont)
finaccion