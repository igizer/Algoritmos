En un sector de un hospital, donde se encuentran internados 50 pacientes,
se toma la temperatura de cada paciente 4 veces al día durante una semana.

Lectura	Dom	Lun	Mar	Mie	Jue	Vie	Sab
1	  	XX	XX	XX	XX	XX	XX	XX
2	  	XX	XX	XX	XX	XX	XX	XX
3	  	XX	XX	XX	XX	XX	XX	XX
4	  	XX	XX	XX	XX	XX	XX	XX
Se dispone de un arreglo con la información recopilada de todos los pacientes.

Construir un algoritmo que:

-Liste por pantalla las temperaturas máxima y mínima de cada paciente, indicando el día y lectura en la que ocurrieron.
-Genere un nuevo arreglo que contenga la temperatura promedio por día de cada paciente.


ACCION Ej_3_23 (Pacientes: Arreglo [1..4,1..7,1..50] de entero) ES




	AMBIENTE

		Salida: Arreglo [1..50, 1..7] de entero
		i, j, k, tmax, tmin, lmax, lmin, dmax, dmin, tpromedio : entero




	PROCESO

		PARA i:= 1 HASTA 50 HACER //PRIMERO INICIALIZO
			PARA j:= 1 HASTA 7 HACER
				Salida[i,j]:= 0
			FINPARA
		FINPARA

		PARA k:= 1 HASTA 50 HACER //paciente por paciente

			tmax:= LV
			tmin:= HV
			lmax:= 0
			lmin:= 0
			dmax:= 0
			dmin:= 0

			PARA j:= 1 HASTA 7 HACER //día por día

				tpromedio:=0

				PARA i:= 1 HASTA 4 HACER //nro de lectura para este día

					tpromedio:= tpromedio + Paciente[i,k,j]

					SI tmax < Paciente[i,k,j] ENTONCES
						tmax:= Paciente[i,k,j]
						lmax:= i
						dmax:= j
					FINSI

					SI tmin > Paciente[i,k,j] ENTONCES
						tmin:= Paciente[i,k,j]
						lmin:= i
						dmin:= j
					FINSI
				
				FINPARA

				Salida[k,j]:= tpromedio/4

			FINPARA

			ESC("Para el paciente nro: ", k, " la temperatura máxima registrada es de: ", tmax, " día de semana ", dmax, " lectura ", lmax)
			ESC("y su temperatura mínima registrada es de: ", tmin, " día de semana ", dmin, " lectura ", lmin)
		
		FINPARA

FINACCION







		

















ACCION Ej_3_23 (A: Arreglo[1..4,1..7,1..50] de entero) ES

	AMBIENTE

		B: Arreglo [1..50, 1..7] de entero
		i, j, k, tmax, tmin, dmax, dmin, tsuma, tprom: entero
	
	PROCESO

		PARA i:= 1 HASTA 50 HACER
			PARA j:= 1 HASTA 7 HACER
				B[k,j]:= 0
			FINPARA
		FINPARA


		PARA k:= 1 HASTA 50 HACER

			dmax:= 0
			dmin:= 0
			tmax:= 0
			tmin:= 0

			PARA j:= 1 HASTA 7 HACER

				tsuma:=0
				
				PARA i:= 1 HASTA 4 HACER
					tsuma:= A[i,j,k] + tsuma
				FINPARA

				tprom:= tsuma div 4

				B[k,j]:= tprom
				
				SI B[k,j] > tmax ENTONCES
					tmax:= B[k,j]
					dmax:= j
				SINO
					SI B[k,j] < tmin ENTONCES
						tmin:= B[k,j]
						dmin:= j
					FINSI
				FINSI

			FINPARA

			ESC("Para el paciente Nro. ", k, ", se registró en su día ", dmax, " de estadía una temperatura maxima de: ", tmax)
			ESC("Y se registró en su día ", dmin, " de estadía una temperatura mínima de: ", tmin)

		FINPARA

FINACCION																																									