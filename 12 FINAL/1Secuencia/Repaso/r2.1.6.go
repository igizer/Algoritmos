// Dada una secuencia de enteros que almacena la cantidad de habitantes de las ciudades
// capitales de las 23 provincias de la República Argentina, discriminados 4 categorías:
// menores de 18 años (varones y mujeres) y mayores de 18 años (varones y mujeres).
// Se pide determinar la población total y los siguientes porcentajes: masculinos,
// femeninos, mayores de 18 y menores de 18.

accion repaso2_1_6 es
	ambiente
		sec: secuencia de entero
		vmen, vmay, mmen, mmay, total, v, i: entero // varones menores, varones mayores, mujeres menores, mujeres mayores

	proceso
		arr(sec) // abrir secuencia de entrada
		avz(sec, v) // leer primer valor de la secuencia
		vmen:= 0; vmay:= 0; mmen:= 0; mmay:= 0; // inicializar acumuladores

		para i:= 1 hasta 23 // recorrer las 23 capitales de provincia
			vmen:= vmen + v // sumar cantidad de varones menores
			avz(sec, v) // leer siguiente valor
			mmen:= mmen + v // sumar cantidad de mujeres menores
			avz(sec, v) // leer siguiente valor
			vmay:= vmay + v // sumar cantidad de varones mayores
			avz(sec, v) // leer siguiente valor
			mmay:= mmay + v // sumar cantidad de mujeres mayores
			avz(sec, v) // leer siguiente valor (queda listo para la próxima vuelta)
		finpara

		cerrar(sec) // cerrar secuencia

		total:= vmen + vmay + mmen + mmay // calcular total de población

		// mostrar resultados con porcentajes
		esc("la población total del pais es de ", total, " personas")
		esc("el porcentaje de mujeres menores es de ", (mmen/total)*100, "%")
		esc("el porcentaje de mujeres mayores es de ", (mmay/total)*100, "%")
		esc("el porcentaje de varones menores es de ", (vmen/total)*100, "%")
		esc("el porcentaje de varones mayores es de ", (vmay/total)*100, "%")

finaccion


//----------la version de hace un año no tiene coherencia ni sentido-------------
ACCION Ejercicio_2_1_6 ES

	AMBIENTE

		sec: secuencia de enteros
		v, i, fmen, vmen, fmay, vmay, total: entero
		pfmen, pvmen, pfmay, pvmay: real
	
	PROCESO

		ARR(sec); AVZ(sec, v)
		fmen:= 0; fmay:= 0; vmen:= 0; vmay:= 0



		MIENTRAS NFDS(sec) HACER // tratar paises

			PARA i:= 1 HASTA 23 // primero provincias de cada pais

				vmay:= vmay + v
				AVZ(sec, v)
				fmay:= fmay + v
				AVZ(sec, v)
				vmen:= vmen + v
				AVZ(sec, v)
				fmen:= fmen + v
				AVZ(sec, v)

			FINPARA

			MIENTRAS v <> "*" HACER //segundo trato todos los politicos de ese pais
				GRABAR(salida, v)

		FINMIENTRAS

		CERRAR(sec)

		total:= vmar + vmen + fmay + fmen

		pfmen:= fmen*100/total
		pfmay:= fmay*100/total
		pvmen:= vmen*100/total
		pvmay:= vmay*100/total

		ESC("La población total es de ", total)

		ESC("El porcentaje de femeninos menores es de", pfmen)
		ESC("El porcentaje de femeninos mayores es de", pfmay)
		ESC("El porcentaje de varones menores es de", pvmen)
		ESC("El porcentaje de varones mayores es de", pvmay)

FINACCION