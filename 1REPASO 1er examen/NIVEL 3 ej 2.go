Ejercicio 2
Se tiene la información de todos los jugadores de cada uno de los planteles en un archivo secuencial con el siguiente formato

Jugadores (Ordenado por selección, Posición y Dorsal)
Seleccion | Posicion (AR, DF, ME, DE, CT) | Dorsal | Apellido y Nombre | Participaciones

Donde: Posicion (AR (Arquero), DF (Defensa), ME (Mediocampista), DE(Delantero), CT (Cuerpo Técnico)
El campo Participaciones indica cuántas veces participó un jugador en una copa américa. 
Se le solicita:
1)	Obtener total de personas que integran cada Selección, por nombre de selección y posición, y total general discriminando
por jugadores con experiencia y sin experiencia. Excluyendo a los integrantes de la comisión técnica.
2)	Generar una secuencia de salida con la siguiente información:
Selección |  Cant Jugadores nuevos | Cant jugadores con experiencia

ACCION EJ2LVL3 ES

	AMBIENTE

		Jugadores = Registro
			Clave= Registro
				Seleccion: AN(20)
				Posicion: ("AR", "DF", "ME", "DE", "CT")
				Dorsal: N(2)
			FINREGISTRO
			ApellidoNombre: AN(20)
			Participaciones: N(2)
		FINREGISTRO

		Reg_Jugadores: Jugadores
		Arch_Jugadores: Archivo de Jugadores Ordenado por Clave


		Salida = Registro
			Seleccion: AN(20)
			Jugadores_Nuevos: N(3)
			Jugadores_Experimentados: N(3)
		FINREGISTRO

		Reg_Sal: Salida
		Arch_Sal: Archivo de Salida


		Resg_Seleccion: AN(20)
		Resg_Posicion: ("AR", "DF", "ME", "DE", "CT")

		sec_sal: secuencia de caracter ///////////////////////////////////////

		cont_jugadoresCExp, cont_jugadoresSExp, cont_seleccionCExp, cont_seleccionSExp, cont_tgCExp, cont_tgSExp, i, Integrantes: entero

		PROCEDIMIENTO Corte_Posicion ES
			ESC("La cantidad de jugadores para la posicion ", Resg_Posicion, "CON EXPERIENCIA es de ", cont_jugadoresCExp)
			ESC("La cantidad de jugadores para la posicion ", Resg_Posicion, "SIN EXPERIENCIA es de ", cont_jugadoresSExp)
			cont_seleccionCExp:= cont_seleccionCExp + cont_jugadoresCExp
			cont_seleccionSExp:= cont_seleccion + cont_jugadoresSExp
			cont_posicionCExp:= 0
			cont_posicionSExp:= 0
			Resg_Posicion:= Reg_Jugadores.Clave.Posicion
		FINPROCEDIMIENTO

		PROCEDIMIENTO Corte_Seleccion ES
			Corte_Posicion()
			ESC("La cantidad de jugadores para la seleccion ", Resg_Seleccion, "CON EXPERIENCIA es de ", cont_seleccionCExp)
			ESC("La cantidad de jugadores para la seleccion ", Resg_Seleccion, "SIN EXPERIENCIA es de ", cont_seleccionSExp)
			cont_seleccionCExp:= cont_seleccionCExp + cont_jugadoresCExp
			cont_seleccionSExp:= cont_seleccion + cont_jugadoresSExp
			
			Reg_Sal.Seleccion:= Resg_Seleccion
			Reg_Sal.Jugadores_Nuevos:= cont_seleccionCExp
			Reg_Sal.Jugadores_Experimentados:= cont_seleccionSExp
			GRABAR(Arch_Sal, Reg_Sal)

			cont_posicionCExp:= 0
			cont_posicionSExp:= 0
			Resg_Seleccion:= Reg_Jugadores.Seleccion
		FINPROCEDIMIENTO

	PROCESO

		ABRIR E/(Arch_Jugadores); LEER(Arch_Jugadores, Reg_Jugadores)
		ABRIR S/(Arch_Sal)
		cont_jugadoresCExp:= 0
		cont_jugadoresSExp:= 0
		cont_seleccionCExp:= 0
		cont_seleccionSExp:= 0
		cont_tgCExp:= 0
		cont_tgSExp:= 0
		
		Resg_Seleccion:= Reg_Jugadores.Clave.Seleccion
		Resg_Posicion:= Reg_Jugadores.Clave.Posicion

		MIENTRAS NFDA(Arch_Jugadores) HACER

			SI Resg_Seleccion <> Reg_Jugadores.Clave.Seleccion ENTONCES
				Corte_Seleccion()
			SINO
				SI Resg_Posicion <> Reg_Jugadores.Clave.Posicion ENTONCES
					Corte_Posicion()
				FINSI
			FINSI

			SI Reg_Jugadores.Clave.Posicion <> "CT" ENTONCES
				SI Reg_Jugadores.Clave.Participaciones > 0 ENTONCES
					cont_jugadoresCExp:= cont_jugadoresCExp + 1
				SINO
					cont_jugadoresSExp:= cont_jugadoresSExp + 1
				FINSI
			FINSI

			LEER(Arch_Jugadores, Reg_Jugadores)
		
		FINMIENTRAS

		Corte_Seleccion()
		ESC("El total general de jugadores con experiencia es de ", cont_tgCExp)
		ESC("El total general de jugadores sin experiencia es de ", cont_tgSExp)
		
		CERRAR(Arch_Jugadores); CERRAR(Arch_Sal)
		ESC("Archivo de Salida creado con éxito")

FINACCION