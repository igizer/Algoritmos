ACCION 2.1.14 ES
    AMBIENTE
    	se,ss:secuencia de caracteres;
        ve:caracter;
        total,pares,impares:enteros;
        numerospares= ("0","2","4","6","8");
        numerosimpares= ("1","3","5","7","9");

	PROCESO

		pares:=0;
		impares:=0;
		CREAR(ss);
		ARR(se);
		AVZ(se,ve)
		MIENTRAS NFDS(se) HACER
			SI ve EN numerospares ENTONCES
				pares:= pares + 1;
			SINO 
				SI ve EN numerosimpares ENTONCES
					impares:= impares + 1;
					GRABAR(ss,ve);
				FINSI
			FINSI
			AVZ(se,ve);
		FINMIENTRAS

		CERRAR(se); CERRAR(ss);
		total:= impares + pares;
		ESC("El total de digitos impares que se encontraron son: ", impares)
		ESC("El total de digitos pares que se encontraron son: ", pares)
		ESC("El total de caracteres numericos (dígitos) son: ", total)
		ESC("El porcentaje de caracteres impares es: ", (impares/(total * 0,01)) "%")
		ESC("El porcentaje de caracteres pares es: ", (pares/(total * 0,01)) "%")
FINACCION
       