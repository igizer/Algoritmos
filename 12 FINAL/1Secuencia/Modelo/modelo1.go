La información de los equipos se encuentra en una secuencia de carácter, con el siguiente formato:
Argentina#Integrantes(2caracteres)EsCabezaDeSerie(1caracter)Brasil#34SCanada#33N … # Colombia#FDS

Integrantes indica la cantidad de personas que conforman el equipo (jugadores y cuerpo técnico) EsCabezaDeSerie indica S o N
si la selección fue considerada como cabeza de serie en el sorteo.

Y la informacion de la composicion del plantel en una secuencia de caracteres con el siguiente formato
Posicion(2caracter)NombreYApellidoJugador#Dorsal(2caracter)Edad(2caracter)|FDS

Posición: AR (Arquero), DF (Defensa), ME (Mediocampista), DE (Delantero), CT (Cuerpo Técnico)

Se le solicita:
a) Generar una secuencia de salida con la información de todos los Jugadores (no conforman el cuerpo técnico) (Nombre y apellido)
que forman parte de equipos que solo son cabeza de serie. Separar la información de los Jugador con el carácter “%”
b) Indicar por equipo, cuál es la edad del jugador más joven.

ACCION MODELO1 ES
	AMBIENTE
		equipos, plantel, salida: secuencia de caracter
		eq, pl, men: caracter //eq: ventana de equipos/ pl: ventana de plantel/ men: menor edad
		cont, i: entero

		funcion convertir(x:caracter):entero
			segun
				= "0": convertir:= 0
				= "1": convertir:= 1
				= "2": convertir:= 3
				= "3": convertir:= 3
				= "4": convertir:= 4
				= "5": convertir:= 5
				= "6": convertir:= 6
				= "7": convertir:= 7
				= "8": convertir:= 8
				= "9": convertir:= 9
			finsegun
		finfuncion

	PROCESO
		arr(equipos); avz(equipos, eq)
		arr(plantel); avz(plantel, pl)
		crear(salida)

		mientras nfds(equipos) hacer
			

FINACCION









































ACCION Nivel3Ej1 ES
	AMBIENTE
		equipos, plantel, jugadores_c_s: secuencia de caracteres
		eq, pl: caracteres
		i, paises, integrantes, joven, edad: entero

		FUNCION convertir(x:caracter):entero
			SEGUN
				= "0": convertir:= 0
				= "1": convertir:= 1
				= "2": convertir:= 2
				= "3": convertir:= 3
				= "4": convertir:= 4
				= "5": convertir:= 5
				= "6": convertir:= 6
				= "7": convertir:= 7
				= "8": convertir:= 8
				= "9": convertir:= 9
			FINSEGUN
		FINFUNCION

	PROCESO
		ARR(equipos); AVZ(equipos, eq)
		ARR(plantel); AVZ(plantel, pl)
		CREAR(jugadores_c_s)
		paises:= 0
		integrantes:= 0
		MIENTRAS NFDS(equipos) HACER
			joven:= 100
			MIENTRAS eq <> "#" HACER
				GRABAR(eq)
				AVZ(equipo, eq)
			FINMIENTRAS
			AVZ(equipos, eq)
			integrantes:= convertir(eq)*10
			AVZ(equipos, eq)
			integrantes:= integrantes + convertir(eq)
			AVZ(equipos, eq)
			PARA i:=1 HASTA integrantes
				SI eq = "S" ^ pl <> "C" ENTONCES
					AVZ(plantel, pl)
					AVZ(plantel, pl)
					MIENTRAS pl <> "#"
						GRABAR(jugadores_c_s, pl)
						AVZ(plantel, pl)
					FINMIENTRAS
				SINO
					MIENTRAS pl <> "#"
						AVZ(plantel, pl)
					FINMIENTRAS
				FINSI
				PARA i:= 1 HASTA 3
					AVZ(plantel, pl)
				FINPARA
				edad:= convertir(pl)*10
				AVZ(plantel, pl)
				edad:= edad + convertir(pl)
				SI edad < joven ENTONCES
					joven:= edad
				FINSI
			FINPARA
			ESC("El más joven de este equipo tiene ", joven, " años")
			AVZ(equipos, eq)
		FINMIENTRAS
		Corte
	
























	PROCESO
		Mientras NFDS(equipos) hacer
		Mientras eq <> "#" hacer
			Escribir(eq)
			Avanzar(equipos, eq)
		FinMientras

		// Avanzamos # 
		Avanzar(equipos, eq)
		ca_integrantes:=Convertir(eq)*10; Avanzar(equipos, eq)
		ca_integrantes:=ca_integrantes+Convertir(eq); Avanzar(equipos, eq)
		
		es_cabeza_serie:=eq;

		edad_mas_joven:=100

		// Recorremos informacion del plantes
		Para i:= 1 hasta ca_integrantes hacer
			posc_1:= pl; Avanzar(plantel, pl);
			posc_2:= pl; Avanzar(plantel, pl);

			// Nombre y apellido del jugador
			Si es_cabeza_serie = "S" y posc_1 <> "C" y posc_2 <> "T" entonces 
				Mientras pl<>"#" hacer 
					Escribir(salida, pl)
					Avanzar(plantel, pl);
				FinMientras

				Escribir(salida, "%")
			Sino 
				Mientras pl<>"#" hacer 
					Avanzar(plantel, pl);
				FinMientras	
			FinSi

			// Avanzamos # 
			Avanzar(plantel, pl);
			
			// Avanzamos Dorsal
			Avanzar(plantel, pl); Avanzar(plantel, pl);
			edad:=Convertir(pl)*10;  Avanzar(plantel, pl);
			edad:=edad+Convertir(pl);  Avanzar(plantel, pl);

			Si edad < edad_mas_joven entonces
				edad_mas_joven:= edad
			FinSi

		FinPara

		Escribir("Edad Jugador mas joven:", edad_mas_joven)

		FinMientras





























	
		ARR(equipos); AVZ(equipos, eq) // Arrancamos la secuencia de todos los países que son
		ARR(plantel); AVZ(plantel, pl) // Arrancamos la secuencia de todos los integrantes de cada equipo/país
		CREAR(jugadores_c_s) // Creamos ahora la secuencia de salida, para grabar todos los jugadores de cada equipo cabeza de serie
		MIENTRAS NFDS(equipos) HACER
			paises:= paises + 1 // Cuento cuantos países son, uno por uno
			REPETIR // Ya contado el país, escribiremos el nombre en pantalla para luego informar el más joven de cada equipo
				ESC(eq)
				AVZ(equipos, eq) 
			HASTA QUE eq = "#"
			AVZ(equipos, eq) // Salimos de la marca "#", y pasamos a ver cuantos integrantes son
			integrantes:= convertir(eq)*10 // Como el número es un caracter y no un digito, lo convertiremos a uno
										   // para poder sumar primero la decena
			AVZ(equipos, eq)
			integrantes:= integrantes + convertir(eq) //una vez guardada la decena, ahora avanzamos y guardamos la unidadde jugadore
			AVZ(equipos, eq) // Ahora avanzamos a ver si es cabeza se serie según una "S" de un sí, en caso contrario será "N" de un no
			SI eq = "S" ENTONCES // Entonces sí es cabeza de serie
				PARA i:=1 HASTA integrantes HACER //Trato cada integrante del equipo cabeza de serie, para ver su es jugador o ct
					SI pl <> "C" ENTONCES // si empieza con C la posición es cuerpo tecnico, sino es jugador
						AVZ(plantel, pl)
						AVZ(plantel, pl)
						Mientras pl <> "#" HACER // Llegamos a la primera letra del jugador para grabar todo el nombre y apellido
							GRABAR(jugadores_c_s, pl) // Grabamos y avanzamos caracter a caracter del nombre y apellido hasta la marca# 
							AVZ(plantel, pl)
						FINMIENTRAS
						GRABAR(jugadores_c_s, "%") // La info de cada jugador pide separado por "%", entonces se lo escribe así
					SINO
						Mientras pl <> "#" HACER // Como es cuerpo técnico
						AVZ(plantel, pl) // Avanzamos hasta la marca#, sin grabar ya qué la consigna no nos pide a los cuerpos técnicos
						FINMIENTRAS
					FINSI
					PARA i:=1 HASTA 5 // Una vez llegamos a # en ambos casos, avanzamos 5 lugares hasta la primera letra de posicón
							AVZ(plantel, pl) 
					FINPARA
			SINO // Entonces no es cabeza de serie como único caso contrario del anterior
				PARA i:=1 HASTA integrantes HACER // A diferencia de antes, ahora no grabamos en ningún caso, entonces salteamos a todos
					Mientras pl <> "#" HACER // Lo mismo que hicimos antes para saltear cuerpos técnicos, pero ahora con todos
						AVZ(plantel, pl)
					FINMIENTRAS
					Para i:=1 HASTA 5 // Avanzamos 5 lugares hasta info del siguiente integrante a saltear
							AVZ(plantel, pl) 
					FINPARA
				FINPARA
			FINSI
					

						

							


