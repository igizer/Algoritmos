





00450
12345

ACCION Ej_1 ES

	AMBIENTE //NO SÉ SI ESTÁ COMPLETO, EL PROCESO SI	

		m: Arreglo [1..3, "A"..."E"] de entero
		i: entero
		j: "A"..."E"

		r= REGISTRO
			nota: N(2)
			legaho: N(5)
			comision: ("A", "B", "C", "D", "E") // cabe aclarar que la com E no existe, pero es para ponerle la suma total
		FINREGISTRO                             // de todos los otros. Ejemplo: los aprobados de todas las coms es [1, E] sumados en E
												// y los desaprobados van sumados en [2, E]
												// luego el total que rindieron es la suma tanto de aprobados como desaprobados, y eso va en
												// [3, E]

	
	PROCESO

		ABRIR E/(arch); LEER(arch,reg)
		PARA i:= 1 HASTA 3 HACER
			PARA j:= "A" hasta "E" HACER
				m[i, j]:= 0
			FINPARA
		FINPARA

		MIENTRAS NFDA(arch) HACER
			//indice "i" nos indica la fila (Aprobado/Desaprobado)
			i:= 2
			SI reg.nota >= 6 ENTONCES
				i:= 1
			FINSI

			j:= reg.comision

			m[i, j]:= m[i, 3] + 1
			m[3, j]:= m[i, 3] + 1
			m[i, "E"]:= m[i, 3] + 1
			m[3, "E"]:= m[3, "E"] + 1

			SEGUN reg.comision HACER
				"A": j:= 1
			FINSEGUN

		FINMIENTRAS