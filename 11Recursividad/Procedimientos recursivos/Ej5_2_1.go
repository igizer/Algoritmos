Imprimir las cifras de un número n (siendo n≥0) en orden inverso al original. Por ej.: el inverso de 254 es 452. 

procedimiento inverso(n) ES

	si n < 10 entonces
		esc(n)
	sino
		esc(n mod 10)
		inverso(n div 10)
	finsi

finprocedimiento