Ejercicio 2
Se tiene la información de todos los jugadores de cada uno de los planteles en un archivo secuencial con
el siguiente formato

Jugadores (Ordenado por selección, Posición y Dorsal)
Seleccion | Posicion (AR, DF, ME, DE, CT) | Dorsal | Apellido y Nombre | Edad | Equipo 

Donde: Posicion (AR(Arquero), DF(Defensor), ME(Mediocampista), DE(Delantero), CT(Cuerpo Técnico)
Se le solicita:
Obtener total de personas que integran cada Selección, por nombre de selección y posición, y
total general. Excluyendo a los integrantes de la comisión técnica.

Por selección indicar, promedio de edad de los jugadores.


ACCION EJLVL2 ES

	AMBIENTE
		jugadores= registro
			Seleccion: AN(12)
			Posicion: ("AR", "DF", "ME", "DE", "CT")
			Dorsal: N(2)
			ApellidoNombre: AN(20)
			Edad: N(2)
			Equipo: AN(25)
		FINREGISTRO

		reg_jugadores: jugadores
		arch_jugadores: Archivo de jugadores ordenado por Seleccion, Posicion, Dorsal

		resg_seleccion: AN(12)
		resg_posicion: ("AR", "DF", "ME", "DE", "CT")

		acumulador_edad, cont_jugadores, cont_seleccion, cont_tg: entero

		PROCEDIMIENTO Corte_Posicion ES
			ESC("Para la posición ", resg_posicion, ", la cantidad de personas es de ", cont_jugadores)
			cont_seleccion:= cont_seleccion + cont_jugadores
			cont_jugadores:= 0
			resg_posicion:= reg_jugadores.posicion
		FINPROCEDIMIENTO

		PROCEDIMIENTO Corte_Seleccion ES
			Corte_Posicion()
			ESC("Para la selección ", resg_seleccion, ", la cantidad de personas es de ", cont_seleccion)
			ESC("La edad promedio de los jugadores de esta selección es de ", acumulador_edad/cont_seleccion)
			// Arriba yo usé para sacar el promedio diviendo con "/", nos va a dar de edad posiblemente un promedio
			//Con coma, o sea un real, pero si queremos una edad exacta, usamos "a DIV bS"
			cont_tg:= cont_tg + cont_seleccion
			cont_seleccion:= 0
			acumulador_edad:= 0
			resg_seleccion:= reg_jugadores.seleccion
		FINPROCEDIMIENTO

	PROCESO
		ABRIR E/(arch_jugadores); LEER(arch_jugadores, reg_jugadores)
		cont_jugadores:= 0
		cont_posicion:= 0
		cont_seleccion:= 0
		acumulador_edad:= 0
		
		MIENTRAS NFDA(arch_jugadores) HACER

			SI resg_seleccion <> reg_jugadores.seleccion ENTONCES
				Corte_Seleccion()
			SINO
				SI resg_posicion <> reg_posicion ENTONCES
					Corte_Posicion()
				FINSI
			FINSI

			SI reg_jugadores.Posicion <> "CT" ENTONCES
				cont_jugadores:= cont_jugadores + 1
			FINSI

			LEER(arch_jugadores, reg_jugadores)

		FINMIENTRAS

		Corte_Seleccion()
		ESC("El total general de personas es de ", cont_tg)
	CERRAR(arch_jugadores)
FINACCION


		




























ACCION EJ2LVL1 ES

	AMBIENTE
		jugadores= registro
			Seleccion: AN(12)
			Posicion: ("AR", "DF", "ME", "DE", "CT")
			Dorsal: N(2)
			ApellidoNombre: AN(20)
			Equipo: AN(25)
		FINREGISTRO
		
		reg_jugadores= jugadores
		arch_jugadores= archivo de jugadores ordenado por Seleccion, Posicion, Dorsal

		resg_seleccion: AN(12)
		resg_posicion: ("AR", "DF", "ME", "DE", "CT") 
		resg_dorsal: N(2)

		cont_personas, cont_seleccion, cont_tg: entero

		PROCEDIMIENTO Corte_Posicion ES
			ESC("Para la posición", resg_posicion, ", la cantidad de personas son ", cont_personas)
			cont_seleccion:= cont_seleccion+cont_personas
			cont_personas:= 0
			resg_posicion:= reg_jugadores.posicion
		FINPROCEDIMIENTO

		PROCEDIMIENTO Corte_Seleccion ES
			cont_posicion()
			ESC("Para la selección ", resg_seleccion, ", la cantidad de personas son ", cont_seleccion)
			cont_tg:= cont_tg + cont_seleccion
			cont_seleccion:= 0
			resg_seleccion:= reg_jugadores.seleccion
		FINPROCEDIMIENTO


	PROCESO
		ABRIR E/(arch_jugadores); LEER(arch_jugadores, reg_jugadores)
		resg_posicion:= reg_jugadores.posicion
		resg_seleccion:= reg_jugadores.seleccion
		cont_personas:= 0
		cont_posicion:= 0
		cont_seleccion:= 0
		cont_tg:= 0

		MIENTRAS NFDA(arch_jugadores) HACER

			SI resg_seleccion <> reg_jugadores.seleccion
				Corte_Seleccion()
			SINO
				SI resg_posición <> reg_jugadores.posicion
					Corte_Posicion()
				FINSI
			FINSI
			
			cont_personas:= cont_personas + 1
			LEER(arc_jugadores, reg_jugadores)

		FINMIENTRAS
		Corte_Seleccion()
		ESC("El total general de personas es de: ", cont_tg)

FINACCION