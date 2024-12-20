Ejercicio 3.20¶
Dadas dos matrices A y B, cuadradas, de 5 x 5, con números enteros, cargar una matriz C, de 5 x 5 teniendo en cuenta
las siguientes condiciones: la diagonal principal completar con ceros, en las posiciones que están por encima de la
diagonal principal, copiar los correspondientes elementos de la matriz A y en las posiciones que están por debajo de la
diagonal principal, copiar los elementos correspondientes de la matriz B.


ACCION Ejercicio_3_20 ES

	AMBIENTE

		A: Arreglo [1..5,1..5] de enteros
		B: Arreglo [1..5,1..5] de enteros
		C: Arreglo [1..5,1..5] de enteros

		i,j: entero

	PROCESO

	// MENOS OPTIMO                                             (MÁS OPTIMO ABAJO CON UN SEGUN)
		PARA i:= 1 HASTA 5 HACER
			PARA j:= 1 HASTA 5 HACER

				SI i=j ENTONCES //diagonal principal

					C[i,j]:= 0
				
				SINO

					SI i<j ENTONCES //por ENCIMA de la diagonal principal i<j

						C[i,j]:= A[i,j]
					
					SINO //por DEBAJO de la diagonal principal i>j

						C[i,j]:= B[i,j]

					FINSI
				
				FINSI
			
			FINPARA
		FINPARA
	
FINACCION



//MÁS OPTIMO         PUEDO USAR TAMBIEN UN SEGUN 

PARA i:= 1 HASTA 5 HACER
	PARA j:= 1 HASTA 5 HACER

		SEGUN i HACER //diagonal principal

			i = j: C[i,j]:= 0

			i < j: C[i,j]:= A[i,j]

			i > j: C[i,j]:= B[i,j]
		
		FINSEGUN

	FINPARA
FINPARA






















ACCION EJ_3_20 ES

	AMBIENTE

		A: ARREGLO [1..5,1..5] de entero
		B: ARREGLO [1..5,1..5] de entero
		C: ARREGLO [1..5,1..5] de entero
		i, j: entero

	PROCESO

		PARA i:= 1 HASTA 5 HACER
			PARA j:= 1 HASTA 5 HACER
				SI i > j ENTONCES
					C[i, j]:= B[i,j]
				SINO
					SI i = j ENTONCES
						C[i, j]:= 0
					SINO
						C[i, j]:= A[i,j]
					FINSI
				FINSI
			FINPARA
		FINPARA
		
FINACCION














		ACCION Arreglo_3.20 (A:ARREGLO [1..5] DE ENTERO, B: ARREGLO[1..5] DE ENTERO) ES

		AMBIENTE
	
		C: ARREGLO [1..5] DE ENTERO
		i,j: ENTERO
	
		PROCESO
	
		Para i:=1 a 5 Hacer
			Para j:=1 a 5 Hacer
				Si i=j
				Entonces
					C[i,j]:=0   //Diagonal principal
				sino
					Si i>j
					Entonces
						C[i,j]:=B[i,j]//Abajo de la matriz
					sino
						C[i,j]:=A[i,j]//Arriba de la matriz
					FS
				FS
			FP
		FP
	fin_accion.