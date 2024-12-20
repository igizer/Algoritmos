// Dada una secuencia de enteros que almacena la cantidad de habitantes de las ciudades
// capitales de las 23 provincias de la República Argentina, discriminados 4 categorías:
// menores de 18 años (varones y mujeres) y mayores de 18 años (varones y mujeres).
// Se pide determinar la población total y los siguientes porcentajes: masculinos,
// femeninos, mayores de 18 y menores de 18.
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













ACCION Ejercicio_2_1_6 ES
	AMBIENTE
		capitales, salida: secuencia de entero;
		i, v, total, vmay18, vmen18, mmay18, mmen18: entero;
		p_vmay18, p_vmen18, p_mmay18, p_mmen18: real;
	PROCESO
	ARR(capitales);
	AVZ(capitales, v);
	CREAR(salida)
	v:= 0, total:=0, vmay18:= 0, vmen18:= 0, mmay18:= 0, mmen18:= 0;
	MIENTRAS NFDS(capitales) HACER;
		PARA i:= 1 HASTA 23 HACER;
			vmay18:= vmay18 + v;
			GRABAR(salida, vmay18);
			AVZ(capitales, v);
			vmen18:= vmen18 + v;
			GRABAR(salida, vmen18);
			AVZ(capitales, v);
			mmay18:= mmay18 + v;
			GRABAR(salida, mmay18);
			AVZ(capitales, v);
			mmen18:= mmen18 + v;
			GRABAR(salida, men18);
			AVZ(capitales, v);
		FINPARA
	FINMIENTRAS
	total:= vmay18 + vmen18 + mmay18 + mmen18;
	p_vmay18:= vmay18 * 100 / total;
	p_vmen18:= vmen18 * 100 / total;
	p_mmay18:= mmay18 * 100 / total;
	p_mmen18:= mmen18 * 100 / total;
	ESC("La población total de las capitales de cada provincia del país son ", total);
	ESC("La cantidad porcentual de varones mayores de 18 es de ", p_vmay18, "%");
	ESC("La cantidad porcentual de varones menores de 18 es de ", p_vmen18, "%");
	ESC("La cantidad porcentual de mujeres mayores de 18 es de ", p_mmay18, "%");
	ESC("La cantidad porcentual de mujeres menores de 18 es de ", p_mmen18, "%");
	CERRAR(capitales)
	CERRAR(salida)
FIN ACCCION