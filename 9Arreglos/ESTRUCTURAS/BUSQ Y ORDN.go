// Los ejercicios me los inventé


//Buscar el dorsal 10 en el plantel de un equipo de 24 personas

ACCION Busqueda_Lineal ES

	AMBIENTE

		A: Arreglo [1..24] de Jugadores

		Jugadores= REGISTRO
			Nombre: AN(20)
			Posicion: AN(2)
			Dorsal: N(2)
		FINREGISTRO

		Busco= 10
		i: Entero

	PROCESO

		PARA i:= 1 HASTA 24 HACER

			SI BUSCO = A[i].Dorsal ENTONCES
				ESC("El Dorsal 10 que buscaba está en la posición ", i)
			FINSI

		FINPARA
	
	FINACCION



// El mismo ejercicio que el anterior, pero ahora que el algoritmo se detenga cuando encuentre el dorsal 10


ACCION Busqueda_Lineal_Con_Centinela ES

	AMBIENTE

		A: Arreglo [1..24] de Jugadores

		Jugadores= REGISTRO
			Nombre: AN(20)
			Posicion: AN(2)
			Dorsal: N(2)
		FINREGISTRO

		Busco= 10
		i: Entero

	PROCESO

		i:= 1

		MIENTRAS (i < 24) v (A[i].Dorsal <> Busco) HACER

			i:= i + 1

		FINMIENTRAS

		SI A[i].Dorsal = Busco ENTONCES
			ESC("El dorsal 10 que buscaba, está en la posición ", i)
		SINO
			ESC("No hay ninguna persona con Dorsal 10 en el plantel")
		FINSI

FINACCION


















	LEER(X)
	i:= 1

	MIENTRAS (A[i] <> X) y (i < N) HACER
		i:= i + 1
	FINPARA

	SI A[i] = X ENTONCES
		ESC("El valor ", X, ", se encuentra en la posición ", i)
	SINO
		ESC("Valor no encontrado")
	FINSI









	LEER(X)
	i:=1
	Posiciones:=N

	centro:=(i+Posiciones) DIV 2

	MIENTRAS (i < Posiciones) y (A[centro] <> X) HACER
		SI (A[centro] > X) ENTONCES
			Posiciones:=centro-1
		SINO
			i:=centro+1
		FIN SI
		centro:=(i+Posiciones) div 2;
	FIN MIENTRAS

	SI (A[centro] = X) ENTONCES
		ESCRIBIR ("El elemento buscado está en la posición ",centro)
	SINO
		ESCRIBIR ("El elemento buscado no está en el arreglo");
	FIN SI





Inserción directa¶


// Ordena de menor a mayor

CARGAR_ARREGLO

PARA i:= 2 a n HACER
    x:=a[i]
    j:=i-1
    MIENTRAS (j > 0) y (x < a[j]) HACER
        a[j+1]:=a[j]
        j:=j-1
    FIN MIENTRAS
    a[j+1]:=x
FIN PARA



	CARGAR_ARREGLO

	PARA i:= 2 HASTA N HACER

		X:= A[i]
		j:= i-1

		MIENTRAS (j > 0) y (X < A[j]) HACER
			A[j+1]










	//Ordena de MENOR a MAYOR
	
	CARGAR_ARREGLO

	PARA i:=2 HASTA N HACER // TE ORDENA PERFECTO no importa que tan desordenado

		X:= A[i]
		j:= i-1

		MIENTRAS (j>0) y (X<A[j]) HACER
			A[j+i]:= A[j]
			j:=j-1
		FINMIENTRAS
		A[j+1]:= X

	FINPARA








	Seleccion directa¶


	// Ordena de mayor a menor

	CARGAR_ARREGLO

	PARA i:= 1 a n-1 HACER
		x:=a[i]
		max:=i
		PARA j:= i+1 a n HACER
			SI (x < a[j]) ENTONCES
				max:=j
				x:=a[j]
			FIN SI
		FINPARA
		a[max]:=a[i]
		a[i]:=x
	FINPARA




	CARGAR_ARREGLO

	PARA i:= 1 HASTA N-1 HACER // TE ORDENA PERFECTO no importa que tan desordenado

		X:= A[i]
		max:= i

		PARA j:= i+1 HASTA N HACER

			SI X<A[j] ENTONCES
				X:= A[j]
				max:= j
			FINSI

		FINPARA

		A[max]:= A[i]
		A[i]:=X

	FINPARA




	CARGAR_ARREGLO

	band:= falso // TE ORDENA PERFECTO no importa que tan desordenado

	MIENTRAS no band HACER

		band:= verdadero

		PARA j:= 1 HASTA n-1 //Ordena solo 2 nums, pero por el mientras, terminamos haciendo varios- 
							 //-PARA que ordenan todo a la larga de 2 en 2, no es tan eficiente/rapido
			SI A[j] < A[j+1] ENTONCES
				X:= A[j]
				A[j]:= A[j+1]
				A[j+1]:= x
				band:= falso
			FINSI

		FINPARA

	FINMIENTRAS