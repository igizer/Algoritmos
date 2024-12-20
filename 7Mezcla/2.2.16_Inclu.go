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

		PELICULA = REGISTRO
			Nro_Pelicula: N(6)
			Titulo: AN(20)
			Género: AN(15)
			Cant_Copias: N(5)
			Fecha_Estreno = REGISTRO
				AA: N(4)
				MM: 1..12
				DD: 1..31
			FINREGISTRO
		FINREGISTRO

		RegNuev, RegPel, RegSal: PELICULA

		ArchNuev, ArchPel, ArchSal: Archivos de PELICULA Ordenados por Nro_Pelicula

		PROCEDIMIENTO LeerNuev ES
			Leer(ArchNuev,RegNuev)
			SI FDA(ArchNuev)
				RegNuev.Nro_Pelicula:= HV
			FINSI
		FINPROCEDIMIENTO

		PROCEDIMIENTO LeerPelis ES
			Leer(ArchPel, RegPel)
			SI FDA(ArchPel) HACER
				RegPel.Nro_Pelicula:= HV
			FINSI
		FINPROCEDIMIENTO

	PROCESO

		ABRIR E/(ArchNuev); ABRIR E/(ArchPel); ABRIR S/(ArchSal);
		LEER(ArchNuev, RegNuev); LEER(ArchPel, RegPel);

		MIENTRAS RegNuev.Nro_Pelicula <> HV o RegPel.Nro_Pelicula <> HV HACER
			SI RegNuev.Nro_Pelicula < RegPel.Nro_Pelicula ENTONCES
				GRABAR(ArchSal, RegNuev)
				LeerNuev()
			SINO
				SI RegNuev.Nro_Pelicula > RegPel.Nro_Pelicula ENTONCES
					GRABAR(ArchSal, RegPel)
					LeerPel()
				SINO
					SI RegNuev.Nro_Pelicula = RegPel.Nro_Pelicula ENTONCES
						RegSal:= RegPel
						GRABAR(ArchSal, RegSal)
						LeerNuev(); LeerPel();
				FINSI
			FINSI
		FINMIENTRAS

		CERRAR(ArchNuev); CERRAR(ArchPel); CERRAR(ArchSal);

FINACCION

