Ejercicio 1
La información de los equipos se encuentra en una secuencia de carácter, con el siguiente formato:
Argentina#Brasil#Canada#Peru# … # Colombia#FDS

Además se encuentra con información individual de cada jugador: dorsal y edad en una secuencia de enteros con el siguiente formato
10|36|11|36|19|...|33|18|FDS
Esto significa que, el primer jugador de la Selección Argentina que usará el dorsal 10 tiene 36 años (🐐 🏆), el siguiente usará
el dorsal 11 y también tiene 36 años, y así sucesivamente.
La cantidad de jugadores que componen un plantel es de 26 jugadores.

Se le solicita:
Informar el promedio de edad de los jugadores de un determinado dorsal, que lo indicará el usuario por única vez.
¿Cuál es la edad del jugador más joven por equipo? ¿Y la edad del más longevo de cada equipo?

ACCION NIVEL2 ES
	AMBIENTE
		equipos: secuencia de caracter
		jugadores: secuencia de enteros
		v1: caracter
		v2, dorsal, joven, longevo, suma_edades, resg_dorsal, jugadores_c_dorsal, paises, i: entero

	PROCESO
		ARR(paises); AVZ(paises, v1)
		ARR(jugadores); AVZ(jugadores, v2)
		ESC("A continuación, escriba el número de dorsal a sacar edad media por jugador")
		LEER(dorsal)
		paises:= 0
		jugadores_c_dorsal:= 0
		suma_edades:= 0
		MIENTRAS NFDS(paises) HACER // Tratamos equipo por equipo
			joven:= 100
			longevo:= 0
			REPETIR
				AVZ(paises, v1)
			HASTA QUE v1="#"
			paises:= paises + 1
			ESC("Para equipo N°", paises)
			PARA i:=1 HASTA 26 HACER
				resg_dorsal:= v2 // resguardo numero de dorsal
				AVZ(jugadores, v2) // ahora habiendo resguardado, voy a comparar, y guardar si me sirve, además de definir longevo y joven
				SI resg_dorsal = dorsal ENTONCES
					suma_edades:= suma_edades + v2
					jugadores_c_dorsal:= jugadores_c_dorsal + 1
				FINSI
				SI joven > v2 ENTONCES
					joven:= v2
				FINSI
				SI longevo < v2 ENTONCES0


				
					longevo:= v2
				FINSI
				AVZ(jugadores, v2)
			FINPARA
			ESC("La edad de jugador más joven del equipo es ", joven, ".")
			ESC("La edad de jugador más longevo del equipo es es ", longevo, ".")
			AVZ(paises, v1)
		FINMIENTRAS
		CERRAR(paises)
		CERRAR(jugadores)
		ESC("Para el dorsal ", dorsal, " el promedio de edad por jugador es de: ", (suma_edades/jugadores_c_dorsal), "años.")
FINACCION
