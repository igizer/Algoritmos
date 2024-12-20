Ejercicio 2
Se tiene la información de todos los jugadores de cada uno de los planteles en un archivo secuencial con
el siguiente formato

Jugadores (Ordenado por selección, Posición y Dorsal)
Seleccion | Posicion (AR, DF, ME, DE, CT) | Dorsal | Apellido y Nombre | Equipo 

Donde: Posicion (AR(Arquero), DF(Defensor), ME(Mediocampista), DE(Delantero), CT(Cuerpo Técnico)
Se le solicita:
Obtener total de personas que integran cada Selección, por nombre de selección y posición, y total general.

ACCION EJ2LVL1 ES

	AMBIENTE
		jugadores= registro
			CLAVE= Registro
				Seleccion: AN(12)
				Posicion: ("AR", "DF", "ME", "DE", "CT")
				Dorsal: N(2)
			FINREGISTRO
			ApellidoNombre: AN(20)
			Equipo: AN(25)
		FINREGISTRO

		reg_jugadores= jugadores
		arch_jugadores= archivo de jugadores ordenado por CLAVE

		resg_seleccion: AN(12)
		resg_posicion: ("AR", "DF", "ME", "DE", "CT") 
		resg_dorsal: N(2)

		cont_personas, cont_seleccion, cont_tg: entero

		PROCEDIMIENTO Corte_Posicion ES
			ESC("Para la posición", resg_posicion, ", la cantidad de personas son ", cont_personas)
			cont_seleccion:= cont_seleccion+cont_personas
			cont_personas:= 0
			resg_posicion:= reg_jugadores.clave.posicion
		FINPROCEDIMIENTO

		PROCEDIMIENTO Corte_Seleccion ES
			cont_posicion()
			ESC("Para la selección ", resg_seleccion, ", la cantidad de personas son ", cont_seleccion)
			cont_tg:= cont_tg + cont_seleccion
			cont_seleccion:= 0
			resg_seleccion:= reg_jugadores.clave.seleccion
		FINPROCEDIMIENTO


	PROCESO
		ABRIR E/(arch_jugadores); LEER(arch_jugadores, reg_jugadores)
		resg_posicion:= reg_jugadores.clave.posicion
		resg_seleccion:= reg_jugadores.clave.seleccion
		cont_personas:= 0
		cont_posicion:= 0
		cont_seleccion:= 0
		cont_tg:= 0

		MIENTRAS NFDA(arch_jugadores) HACER

			SI resg_seleccion <> reg_jugadores.clave.seleccion
				Corte_Seleccion()
			SINO
				SI resg_posición <> reg_jugadores.clave.posicion
					Corte_Posicion()
				FINSI
			FINSI
			
			cont_personas:= cont_personas + 1
			LEER(arc_jugadores, reg_jugadores)

		FINMIENTRAS
		Corte_Seleccion()
		ESC("El total general de personas es de: ", cont_tg)
		CERRAR(arch_jugadores)

FINACCION


	

















	Procedimiento CortePosicion es
		Escribir(cont)
		cont_seleccion:=cont_seleccion+cont
		cont:=0
		resg_posicion:= reg.posicion
	FinProcedimiento
	
	
	Procedimiento CorteSeleccion es
		CortePosicion()
		Escribir(cont_seleccion)
		cont_jug:=cont_jug+cont_seleccion
		cont_seleccion:=0
		resg_seleccion:= reg.seleccion
	
	
	FinProcedimiento
	
	
	Mientras NFDA(arch) hacer
		Si reg.seleccion <> resg_seleccion entonces
			CorteSeleccion()
		Sino
			Si reg.posicion <> resg_posicion entonces
				CortePosicion()
			FinSi
		FinSi
	
	
		cont:= cont+1
		Leer(arch, reg)
	
	
	FinMientras
	CorteSeleccion()
	Escribir(cont_jug)
	