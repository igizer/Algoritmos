Aprovechando que arranca una nueva edición de la Copa América de fútbol masculino, los dirigentes de cada uno de 
los equipos y el comité organizativo quiere obtener estadísticas de los jugadores que participarán de la competición.

NIVEL 1

Ejercicio 1

La información de los equipos se encuentra en una secuencia de carácter, con el siguiente formato:
Argentina#Brasil#Canada#Peru# … # Colombia#FDS

Y la edad de cada uno de los jugadores que componen el plantel, se encuentra disponible en una secuencia de enteros
32|22|27|19|...|33|18|FDS
La cantidad de jugadores que componen un plantel es de 26 jugadores.

Se le solicita:
Informar el promedio de edad de los jugadores que participaran este año en la Copa.
¿Cuál es la edad del jugador más joven que participará en la copa? ¿Y la edad del más longevo ?

ACCION ejercicio1lvl1 ES

	AMBIENTE

		paises: secuencia de caracter
		jugadores: secuencia de enteros
		p: caracter
		j, cont, joven, longevo, suma_edades: entero


	PROCESO

		ARR(paises); AVZ(paises, p);
		cont:= 0

		MIENTRAS NFDS(paises) HACER

			REPETIR
				AVZ(paises, p)
			HASTA QUE p="#"
			cont:= cont + 1

		FINMIENTRAS

		CERRAR(paises)

		ARR(jugadores); AVZ(jugadores, j);
		sumas_edades:= 0
		longevo:= 0
		joven:= 100

		MIENTRAS NFDS(jugadores) HACER

			SI j > longevo ENTONCES
				longevo:= j
			FINSI

			SI j < joven ENTONCES
				joven:= j
			FINSI

			suma_edades:= suma_edades + j

			AVZ(jugadores, j)

		FINMIENTRAS

		CERRAR(jugadores)

		ESC("El promedio de edad de los jugadores que participaran este año en la Copa es de ", suma_edades/(cont*26))
		ESC("La edad del jugador más joven es de ", joven)
		ESC("La edad del jugador más longevo es de ", longevo)

FINACCION
























ACCION Ej1Lvl1 ES
	AMBIENTE
		paises: secuencia de caracter
		jugadores: secuencia de entero
		v: caracter
		v2, longevo, joven, suma_edades, cont: entero


	PROCESO
		ARR(paises) //Tratamos los paises
		AVZ(paises, v)
		cont:= 0
		MIENTRAS NFDS(paises) HACER
			REPETIR
				AVZ(paises,v)
			HASTA QUE v="#"
			cont:= cont + 1
		FINMIENTRAS
		CERRAR(paises)

		ARR(jugadores)//Tratamos jugadores
		AVZ(jugadores, v2)
		suma_edades:= 0
		joven:= v2
		longejo:= v2
		MIENTRAS NFDS(jugadores) HACER
			SI v2 > longejo ENTONCES
				longevo:= v2
			FINSI
			SI v2 < joven ENTONCES
				joven= v2
			FINSI
			suma_edades:= suma_edades + v2
			AVZ(jugadores, v)
		FINMIENTRAS
		CERRAR(jugadores)

		ESC("El promedio de edad por jugador es ", suma_edades/(cont*26))
		ESC("La edad más longeva de un jugador es de ", longevo)
		ESC("La edad más joven de un jugador es de ", joven)
FINACCION




			

