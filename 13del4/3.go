ACCION pares_e_impares ES
	AMBIENTE
		sec, pares, impares: secuencia de enteros
		v, numeros, contp, contimp: enteros
	PROCESO
		ARR(sec)
		AVZ(sec, v)
		CREAR(pares)
		CREAR(impares)
		MIENTRAS NFDS(sec) HACER
			numeros:= v
			SI numeros MOD 2 = 0 ENTONCES
				contp= contp +1
			FINSI
			SI numeros MOD 2 = 0 ENTONCES
				contimp= contimp +1
			FINSI
			numeros:= numeros div 10
			SI contp>contimp ENTONCES
				GRABAR(pares, v)
			SINO
				GRABAR(impares, v)
		MIENTRAS
			

			
