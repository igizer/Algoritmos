// Ejercicio_2_1_3
// "Se dispone de una secuencia de caracteres y se desea obtener una secuencia de salida
// que resulte de copiar la secuencia de entrada, descartando el caracter $."

accion r2_1_3 es
	ambiente
		sec, sal: secuencia de caracter
		c: caracter

	proceso
		arr(sec)
		avz(sec, c)
		crear(sal)

		mientras nfds(sec) hacer
			si sec <> "$" entonces
				grabar(sal, c)
			finsi
			avz(sec)
		finmientras

		esc("La secuencia fue realizada con éxito")
		cerrar(sec)
		cerrar(sal)
finaccion