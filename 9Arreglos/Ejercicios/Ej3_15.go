Ejercicio 3.15¶
Se precisa ordenar un arreglo de enteros de menor a mayor, eliminando los números repetidos


ACCION Ejercicio_3_15 ES

	AMBIENTE

		A: Arreglo [1..n] de enteros
		i, j, x: entero

	PROCESO

		x:= A[i]

		PARA i:= 2 HASTA n HACER

			X:=A[i]
			j= i-1

			MIENTRAS (x < A[j]) y (j > 0) ENTONCES

				A[i]:= A[j]
				A[j]:= x
				j:= j-1

			FINSI

			A[j+1]:= x



			NO ENTIENDO NATI


3 2 6 1
2 3 6 1
 
x= 6
  j i
2 3 6 1
2 3 6 1

x= 1
j     i
2 3 6 1
2 3 1 6
2 1 1 3













ACCION Ej_3_15 ES

	AMBIENTE

		A: Arreglo [1..N] de enteros
		i, j, x: entero

	PROCESO

		PARA i:= 2 HASTA N HACER

			x:= A[i]
			j:= i-1

			MIENTRAS (x < A[j]) y (j > 0) ENTONCES

				A[i]:= A[j]
				A[j]:= x
				j:= j-1

			FINSI

			A[j+1]:= x

		FINPARA

FINACCION

   x=2
   j   i  			
0- 3 1 2 4 5
1- 1 3 2 4 5
2- 1 2 3 4 5