// Teniendo en cuenta el ejercicio anterior, se solicita  que la
// secuencia de salida sea una secuencia de caracteres y los CUIT
// se separen unos de otros con el caracter "-".

accion repaso_2_1_8 es

	ambiente
		sec: secuencia de entero // secuencia de entrada (CUITs)
		sal: secuencia de caracter // secuencia de salida (CUITs como caracteres con guiones)
		ant, act, numero: entero // ventana de dos CUITs para evitar guion final
		rango = (0, 1, 2, 3) // conjunto de dígitos válidos (anteúltimo del CUIT)

	proceso
		arr(sec) // abrir secuencia de entrada
		avz(sec, act) // leer primer CUIT
		ant := act
		avz(sec, act) // leer segundo CUIT
		crear(sal) // inicializar secuencia de salida

		mientras act <> 0 hacer // mientras no sea el fin (último CUIT)
			numero := ((ant div 10) mod 10) // extraer anteúltimo dígito del CUIT
			si numero en rango entonces // si está dentro del rango permitido
				grabar(sal, ant) // grabar CUIT
				grabar(sal, "-") // grabar guion porque aún hay al menos un CUIT más
			finsi
			ant := act
			avz(sec, act) // avanzar a siguiente CUIT
		finmientras

		// tratar el último CUIT leído fuera del bucle (para no poner guion al final)
		numero := ((ant div 10) mod 10)
		si numero en rango entonces
			grabar(sal, ant)
		finsi

		cerrar(sec) // cerrar secuencia de entrada
		cerrar(sal) // cerrar secuencia de salida

		esc("secuencia generada con exito.") // mensaje final
fin accion



//---------------versión vieja---------------



// Teniendo en cuenta el ejercicio anterior, se solicita  que la
// secuencia de salida sea una secuencia de caracteres y los CUIT
// se separen unos de otros con el caracter "-".
ACCION Ejercicio_2_1_8 ES
AMBIENTE
	entrada, ordenado: secuencia de entero
	salida: secuencia de caracter
	v, v2: entero
	rango = (0, 1, 2, 3)
PROCESO
	ARR(entrada); CREAR(ordenado);
	AVZ(entrada, v);
	MIENTRAS v <> 0 HACER
		num:= v MOD 10;
		SI ((v MOD 100) DIV 10) EN rango ENTONCES
			GRABAR(ordenado, v);
			GRABAR("-")
		FINSI
		AVZ(entrada, v);
	FIN MIENTRAS
	CERRAR(entrada);
	CERRAR(ordenado);
	ARR(ordenado);
	AVZ(ordenado,v);
	CREAR(salida);
	GRABAR(salida,v);
	MIENTRAS NDFS(ordenado) HACER
		AVZ(ordenado,v)
		GRABAR(salida,"-");
		GRABAR(salida,v);
		
	FINMIENTRAS
FINACCION