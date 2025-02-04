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

ACCION EJ_3_2 ES
	AMBIENTE
		Jugadores = REGISTRO
			Clave = REGISTRO
				Seleccion: AN(10)
				Posicion: ("AR", "DF", "ME", "DE", "CT")
				Dorsal: N(2)
			FR
			Ap_Nom: AN(20)
			Participaciones: N(3)
		FR

		regjug: Jugadores
		archjug: Archivo de Jugadores Ordenado por Clave

		Salida = REGISTRO
			Seleccion: AN(10)
			JugadoresCE: N(4)
			JugadoresSE: N(4)
		FR

		regsal: Salida
		archsal: Archivo de Salida

		
		resg_posicion: ("AR", "DF", "ME", "DE", "CT")
		resg_seleccion: AN(10)

		contexp, contsinexp, contseleccionCE, contseleccionSE, contposicionCE, contposicionSE: entero// 1

		PROCEDIMIENTO Corte_posicion ES
			ESC("El total de jugadores para la posicion ", resg_posicion, " CON EXPERIENCIA es ", contposicionCE)
			ESC("El total de jugadores para la posicion ", resg_posicion, " SIN EXPERIENCIA es ", contposicionSE)
			//CON EXPERIENCIA
			contseleccionCE:= contseleccionCE + contposicionCE
			contposicionCE:= 0
			resg_posicion:= regjug.Clave.Posicion
			//SIN EXPERIENCIA
			contseleccionSE:= contseleccionSE + contposicionSE
			contposicionSE:= 0
			resg_posicion:= regjug.Clave.Posicion
		FP

		PROCEDIMIENTO Corte_seleccion ES
			ESC("El total de jugadores para la seleccion ", resg_seleccion, " CON EXPERIENCIA es ", contseleccionCE)
			ESC("El total de jugadores para la seleccion ", resg_seleccion, " SIN EXPERIENCIA es ", contseleccionSE)
			regsal.Seleccion:= regjug.Clave.Seleccion
			//CON EXPERIENCIA
			contexp:= contexp + contseleccionCE
			regsal.JugadoresCE:= contseleccionCE
			contseleccionCE:= 0
			//SIN EXPERIENCIA
			contsinexp:= contsinexp + contseleccionSE
			regsal.JugadoresSE:= contseleccionSE
			contseleccionSE:= 0

			GRABAR(archsal, regsal)
			resg_seleccion:= regjug.Clave.Seleccion
		FP

	PROCESO

		ABRIR /E(archjug); LEER(archjug,regjug)
		ABRIR S/(archsal)

		contexp:= 0
		contsinexp:= 0
		contseleccionCE:= 0
		contseleccionSE:= 0
		contposicionCE:= 0
		contposicionSE:= 0

		MIENTRAS NFDA(archjug) HACER

			SI regjug.posicion <> resg_posicion ENTONCES
				Corte_Posicion()
			SINO
				SI regjug.seleccion <> resg_seleccion ENTONCES
				Corte_seleccion()
				FINSI
			FINSI

			SI regjug.Clave.Posicion <> "CT" ENTONCES
				SI regjug.Participaciones > 0 ENTONCES
					contposicionCE:= contposicionCE + 1
				SINO
					contposicionSE:= contposicionSE + 1
				FINSI
			FINSI

			LEER(archjug, regjug)
		
		FINMIENTRAS

		Corte_seleccion()

		ESC("El total general de jugadores con experiencia es de ", contexp)
		ESC("El total general de jugadores sin experiencia es de ", contsinexp)

		CERRAR(archjug); CERRAR(archsal)

		ESC("Archivo de salida creado con éxito")

FINACCION


		
