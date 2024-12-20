EXCLUYENTE

Mientras nfda(arch1) y nfda(arch2) hacer
	Si r.nro_pelicula < r2.NRO_PELICULA Entonces
		Grabar(salida, r)
		leer(arch_1, r)
	SINO
		grabar(salida, r2)
		leer( arch_2, r2)
finMientras

mientras nfda(arch1) hacer
	GRABAR(salida, r2)
	leer(arch2, r2)
finMientras

cerrar(salida)

-----------------------------------------------------------------

INCLUYENTE

procedimiento LeerArch1() Entonces
	leer(arch, r)
	Si FDA(arch) hacer
		r1.nro_pelicula:= HV
	finhacer
FINPROCEDIMIENTO

procedimiento LeerArch2() Entonces
	leer(arch2, r2)
	Si FDA(arch) hacer
		r1.nro_pelicula:= HV
	finhacer
FINPROCEDIMIENTO

mientras r.nro_pelicula <> HV o r2.nro_pelicula <> HV hacer
	Si r.nro_pelicula < r2.nro_pelicula Entonces
		grabar(salida, r)
		leerArch1()
	SINO
		grabar(salida, r2)
		leerArch2()
	FINSI
FINMIENTRAS

CERRAR(salida)


-------------------------------------------------------------