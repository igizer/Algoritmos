CONSIGNA 4
Ejercicio 1
La información de los equipos se encuentra en una secuencia de carácter, con el siguiente formato:

EsCabezaDeSerie(1caracter)NombreSeleccion#InformacionJugadores!

EsCabezaDeSerie(1caracter)NombreSeleccion#InformacionJugadores![...] FDS

Donde InformacionJugadores tiene la siguiente estructura:
ApellidoYNombreJugador1-Posicion(2caracter)AñoNacimiento(4caracter)ApellidoYNombreJugador2-Posicion(2caracter)AñoNacimiento(4caracter) [...] FDS

EscabezaDeSerie se indica con un carácter S o N.

Considerar que la Posición puede ser (AR, DF, ME, DE)

Se le solicita:
Cantidad de jugadores que nacieron luego de una fecha ingresada por el usuario.
¿Cuántos jugadores componen cada plantel ?
Generar una secuencia de salida, con el nombre de la selección, y la cantidad de jugadores, de todos los equipos que son cabeza de serie.

se asume que como máximo el plantel podrá tener 99 personas

ACCION EJ1LVL4 ES
	AMBIENTE
		equipos, sal: secuencia de caracter
		eq, cabeza_serie: caracter
		fecha, jugadores_fecha, jugadores_plantel, fecha_cada_jugador, digitos_fecha: entero

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
		CREAR(sal)
		ESC("A continuación, escribe una fecha para analizar jugadores nacidos en ella")
		LEER(fecha)

		jugadores_fecha:= 0
		
		MIENTRAS NFDS(equipos) HACER

			jugadores plantel:= 0
			cabeza_serie:= eq
			AVZ(equipos, eq)

			SI cabeza_serie = "S" ENTONCES
				
				MIENTRAS eq <> "#" HACER
					GRABAR(sal, eq)
					ESC(eq)
					AVZ(equipos, eq)
				FINMIENTRAS

			SINO

				MIENTRAS eq <> "#" HACER
					ESC(eq)
					AVZ(equipos, eq)
				FINMIENTRAS
			FINSI

			AVZ(equipos, eq)

			MIENTRAS eq <> "!" // Tratamos jugadores del plantel
					
				REPETIR
					AVZ(equipos, eq)
				HASTA QUE eq = "-"
				jugadores plantel:= jugadores plantel + 1
				AVZ(equipos, eq)
				AVZ(equipos, eq)

				fecha_cada_jugador:= 0
				digitos_fecha:= 4

				PARA i:= 1 HASTA 4 // Esto convierte "2004" en 2004
					fecha_cada_jugador:= fecha_cada_jugador + convertir(eq) * (10 ** (digitos_fecha-1)
					AVZ(equipos, eq)
					digitos_fecha := digitos_fecha - 1
				FINPARA
					
				SI fecha_cada_jugador = fecha ENTONCES
					jugadores_fecha:= jugadores_fecha + 1
				FINSI

				AVZ(equipos,eq)
			FINMIENTRAS
				
			SI cabeza_serie = "S"
				GRABAR:=(sal, jugadores_plantel)
			FINSI

			ESC("los jugadores para este plantel son ", jugadores_plantel)
		
		FINMIENTRAS

		ESC("La cantidad de jugadores que nacieron en la fecha ", fecha, " son ", jugadores_fecha)
		CERRAR(equipos); CERRAR(sal)

FINACCION



					
