Ejercicio 3.22¶
Dada una matriz cuadrada de 5 x 5 de números, sumar filas y columnas y guardar en una matriz de 2 x 5,
de modo que la fila 1 contenga la suma de cada fila y la fila 2, la suma de cada columna)

a b c d e
f g h i j
q r s t u
i v w x y
z A B C D

a b c d e
f g h i j

//opcion mia
ACCION Ejercicio_3_22 ES

	AMBIENTE

		A: Arreglo [1..5,1..5] de entero
		B: Arreglo [1..2,1..5] de entero
		i, j, acum: entero

	PROCESO

		acum:= 0

		//inicializo B
		PARA i:=1 HASTA 5 HACER
			PARA j:= 1 HASTA 2 HACER
				B[i]:= 0
			FINPARA
		FINPARA


		PARA i:=1 HASTA 5 HACER
			PARA j;= 1 HASTA 5 HACER

				B[1,j]:= B[1,j]+ A[i,j] //Sumo para fila 1 la suma de cada fila
				acum:= acum + A[i,j]

			FINPARA

			B[2,i]:= acum
			acum:= 0
		FINPARA
	
FINACCION

//opcion resuelto
Acción ejercicio22 (mimatriz:arreglo[1..5,1..5] de entero) es
Ambiente
	otramatriz:arreglo[1..2,1..5] de entero
	i,j:1..5
Algoritmo
	Para i:=1 a 2 Hacer
		Para j:=1 a 5 Hacer
			otramatriz[i,j]:=0
		FP
	FP
	Para i:=1 a 5 hacer
		Para j:=1 a 5 Hacer
			otramatriz[1,i]:=otramatriz[1,i]+mimatriz[i,j]
			otramatriz[2,j]:=otramatriz[2,j]+mimatriz[i,j]
		FP
	FP
FA




ACCION Ej_3_22 (A: Arreglo [1..5,1..5] de entero) ES

	AMBIENTE

		B: Arreglo [1..2,1..5] de entero

		i, j: entero

	PROCESO

		PARA i:=1 HASTA 2 HACER
			PARA j:=1 HASTA 5 HACER
				B[i,j]:=0
			FINPARA
		FINPARA

		PARA i:= 1 HASTA 5 HACER

			PARA j:= 1 HASTA 5 HACER

				B[1,i]:= B[1,] + A[i,j]////// MAL, ME BUGUEÉ LA CABEZA
				B[2,j]:= B[2,j] + A[j,i]////// MAL, ME BUGUEÉ LA CABEZA

FINACCION
