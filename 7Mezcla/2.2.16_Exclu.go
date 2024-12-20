Ejercicio 2.2.16¶
Construir un algoritmo que a partir de un fichero de películas nuevas conteniendo:

PELICULAS_NUEVAS Ordenado por Nro_Pelicula
||Nro_Pelicula||Titulo|Genero|Cant_Copias|Fecha_Estreno

y otro fichero de peliculas existentes, ambos ordenados por película,
PELICULAS Ordenado por Nro_Pelicula
||Nro_Pelicula||Titulo|Genero|Cant_Copias|Fecha_Estreno

Genere un único archivo (con el mismo formato de los ficheros de entrada) que contenga
todas las peliculas. Considerar que hay un solo registro por películay no se repiten entre ficheros.



ACCION Mezcla1_Inclu ES

	AMBIENTE
	
		PELICULA= REGISTRO
			Nro_Pelicula: N(6)
			Titulo: AN(20)
			Genero: AN(15)
			Cant_Copias: N(5)
			Fecha_Estreno= REGISTRO
				DIA: 1..31
				MES: 1..12
				AÑO: AN(4)
			FINREGISTRO
		FINREGISTRO

		RegNuev, RegSal, RegPel: PELICULA
		ArchNuev, ArchPel, ArchSal: ARCHIVOS de PELICULA Ordenados por Nro_Pelicula

	PROCESO

		ABRIR E/(ArchNuev); ABRIR E/(ArchPel); ABRIR S/(ArchSal);
		LEER(ArchNuev, RegNuev); LEER(ArchPel, RegPel);

		MIENTRAS NFDA(ArchNuev) y NFDA(ArchPel) HACER
			SI RegPel.Nro_Pelicula < RegNuev.Nro_Pelicula
				GRABAR(ArchSal, RegPel)
				LEER(ArchPel, RegPel)
			SINO
				GRABAR(ArchSal, RegNuev)
				LEER(ArchNuev, RegNuev)
			FINSI
		FINMIENTRAS

		MIENTRAS NFDA(ArchPel) HACER
			GRABAR(ArchSal, RegPel)
			LEER(ArchPel, RegPel)
		FINMIENTRAS

		MIENTRAS NFDA(ArchNuev) HACER
			GRABAR(ArchSal, RegNuev)
			LEER(ArchNuev, RegNuev)
		FINMIENTRAS

		CERRAR(ArchPel); CERRAR(ArchNuev); CERRAR(ArchSal)

FINACCION








Cerrar(salida)