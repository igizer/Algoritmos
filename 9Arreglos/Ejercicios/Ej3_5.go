Ejercicio 3.5¶
Dados 2 vectores:

A: arreglo [1 .. 30] de reales B: arreglo [1 .. 30] de reales

Ambos ordenados de forma creciente, escribir un algoritmo que realice la mezcla de ambos para obtener otro vector tambien ordenado de forma creciente

C: arreglo [1 .. 60] de reales

ACCION Ejercicio3_5 ES

	AMBIENTE

		A: Arreglo [1..30] de reales
		B: Arreglo [1..30] de reales
		C: Arreglo [1..60] de reales
		indice: entero

	PROCESO

		PARA indice:= 1 HASTA 30 HACER
			C[indice]:= A[indice]
			C[indice+30]:= B[indice+30]
		FINPARA







accion cosas es
    Ambiente
        A: Arreglo [1..30] de reales
        B: Arreglo [1..30] de reales
        C: Arreglo [1..60] de reales
        indice: entero
    algoritmo
        // Se supone que A y B están cargador ya				31 ... 60
        Para indice:=1 hasta 30 hacer
            C[indice] := A[indice];				1 2 3 ... 30
            C[indice+30] := B[indice];				62
							
        finpara
        metodo_de_ordenamiento(C);
    finaccion


para i:= 2 hasta 60 hacer
	x:=c[i]
	j:= i-1
	mientras j>0 y x<c[j] hacer
		c[j+1]:= c[j]
		j:= j-1
	fm
	c[j+1]:= x
fp

para I:= 1 hasta 60 hacer
	esc(c[i])
fp




















ACCION Ej_3_5 ES

	AMBIENTE

		A: Arreglo [1..30] de reales
		B: Arreglo [1..30] de reales
		C: Arreglo [1..60] de reales

	PROCESO

		PARA i:= 1 HASTA 30 HACER

			C[i]:= A[i]
			C[i+30]:= B[i]
			
