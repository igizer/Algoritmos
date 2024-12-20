//Inició un nuevo ciclo lectivo en la UTN FRRe, y desde la cátedra de Algoritmos y Estructuras
//de Datos queremos conocer información de los alumnos inscriptos.

//Se requiere que informe:
//Por cada comisión, cantidad de alumnos inscriptos
//Listado de todos los alumnos (apellido y nombre) que ingresaron en el año 2024


ACCION subse_1 ES

	AMBIENTE

		sec: secuencia de caracter
		v, v1_año, v2_año: caracter
		cont, i: entero

	PROCESO

		ARR(sec)
		AVZ(sec)

		MIENTRAS NFDS(sec) HACER

			cont:= 0

			ESC("Para la comision ", v, " Los alumnos 2024: ")
			AVZ(sec, v)

			Para i:= 1 HASTA 12
				AVZ(sec,v)
			FINPARA

			MIENTRAS v <> "%" HACER

				PARA i:= 1 HASTA 4
					AVZ(sec, v)
				FINPARA
				v1_año:= v
				AVZ(sec,v)
				v2_año:= v
				AVZ(sec,v)
				
				MIENTRAS v <> "!" HACER
			
					SI (v1_año = "2") ^ (v2_año = "4") ENTONCES
						ESC(v)
					FINSI
					AVZ(sec, v)
					cont:= cont + 1

				FINMIENTRAS

				PARA i:= 1 HASTA 5
					AVZ(sec, v)
				FINPARA

				ESC("Y la cantidad total de alumnos es ", cont)

			FINMIENTRAS
		
		FINMIENTRAS

		CERRAR(sec)

FINACCION

		




























ACCION EJ1_Subsecuencia ES
	AMBIENTE
		sec: secuencia de caracter
		v, v1_año, v2_año, comision: caracter
		cont, i: entero
	PROCESO
		ARR(sec)
		AVZ(sec, v)
		MIENTRAS NFDS(sec) HACER
			comision:= v
			cont:= 0
			AVZ(sec, v)
			PARA i:= 1 HASTA 12 HACER
				AVZ(sec, v)
			FINPARA

			ESC("A continuación la lista de apellido y nombre de alumnos de esta comisión: ")
			MIENTRAS v <> "%" HACER
				PARA i:= 1 HASTA 4 HACER
					AVZ(sec, v)
				FINPARA
				v1_año:= v
				AVZ(sec, v)
				v2_año:= v
				AVZ(sec, v)
				SI (v1_año="2") ∧ (v2_año="4") ENTONCES
					MIENTRAS v<>"!" HACER
						ESC(v)
						AVZ(sec, v)
					FINMIENTRAS
				SINO
					MIENTRAS v<>"!" HACER
						AVZ(sec, v)
					FINMIENTRAS
				FINSI
				cont:=cont+1
				AVZ(sec, v)
				PARA i:= 1 HASTA 4
					AVZ(sec, v)
				FINPARA
			FINMIENTRAS
			ESC("La cantidad de alumnos de la comisón ", comision, "es de ", cont)
		FINMIENTRAS
		CERRAR(sec)
FINACCION


























ACCION ejercicio1_subsecuencias ES
	AMBIENTE
		entrada, salida: secuencia de caracter
		v: caracter
		c1, c2, c3, c4, c5, c6, año:entero
	PROCESO
		ARR(entrada)
		AVZ(entrada, v)
		CREAR(salida)
		c1:=0
		c2:=0
		c3:=0
		c4:=0
		c5:=0
		c6:=0
		MIENTRAS NFDS(entrada) HACER
					SI v = "1" ENTONCES
						c1:=c1+1
					FINSI
					SI v = "2" ENTONCES
						c2:=c2+1
					FINSI
					SI v = "3" ENTONCES
						c3:=c3+1
					FINSI
					SI v = "4" ENTONCES
						c4:=c4+1
					FINSI
					SI v = "5" ENTONCES
						c5:=c5+1
					FINSI
					SI v = "6" ENTONCES
						c6:=c6+1
					FINSI
					PARA:i 1 HASTA 17 HACER // para llegar hasta el primer dígito de año
						AVZ(sec, v)
					FINPARA
						año:=v*10
						AVZ(entrada,v) // para avanzar el segundo dígito del año
						año:=año+v
						AVZ(entrada,v) // para avanzar a primer caracter del apellido y nombre
						SI año=24 ENTONCES
							MIENTRAS v<>! HACER // para grabar y avanzar aplldo&nbre hasta marca
								GRABAR(salida,v)
								AVZ(entrada,v)
							FINMIENTRAS
						FINSI
						REPETIR
							AVZ(entrada,v)
						HASTA QUE v = "%"
		ESC("Listado de todos los ingresantes 2024 realizado con éxito")
		ESC("La cantidad de alumnos inscriptos por cada comisión es de:")
		ESC("c1: " c1)
		ESC("c2: " c2)
		ESC("c3: " c3)
		ESC("c4: " c4)
		ESC("c5: " c5)
		ESC("c6: " c6)



		Accion Alumnos es
		AMBIENTE
			sec: secuencia de caracter
			v,comision,aa_ingreso_1,aa_ingreso_2: caracter
			cont_alumnos,i: entero
		
		PROCESO
		
			ARR (sec)
			AVZ (sec,v)
			MIENTRAS nofds (sec) hacer 
				comision := v
				cont_alumnos := 0
				AVZ (sec,v)
				PARA i := 1 a 12 hacer // avanzo toda la info de teoria, practica y horarios
					AVZ(sec,v)
				finPara
		
				ESC("Se listan a continuacion el nombre y apellido de los estudiantes que ingresaron en el anio 2024: ")
				MIENTRAS v <> "%" hacer // trato alumnos
					cont_alumnos := cont_alumnos + 1 // cuento al alumno
					PARA i := 1 a 4 hacer // salteo dia y mes de ingreso, ya que me interesa el año
						AVZ(sec,v)
					finPara // Al salir del ciclo estoy en el primer caracter del anio
					aa_ingreso_1 := v  // resguardo primer digito del año
					AVZ(sec,v)
					aa_ingreso_2 := v  // resguardo segundo digito del año
					AVZ(sec,v)   // me posiciono en la primera letra del nombre
					SI (aa_ingreso_1 = "2") ^ (aa_ingreso_2 = "4") ENTONCES 
						MIENTRAS v <> "!" hacer // listarlo al alumno
							ESC (v)
							AVZ (sec,v)
						finMientras	
					SINO 
						MIENTRAS v <> "!" hacer 
							AVZ (sec,v)
						finMientras
					fin si
		
					AVZ (sec,v) // salteo el "!"
					
					PARA i := 1 a 4 HACER //salteo legajo
						avz (sec,v)
					finPara 
					
					// aca puedo estar en otro alumno, o directamente en un %, que significa que ya no hay mas alumnos de esa comisión
		
				FinMientras // v <> “%”
		
				// al salir de este mientras, estoy en una nueva comision, o tambien podria estar en un fds 
		
				// informar cantidad de alumnos por comision
		
				ESC("Comision:", comision, "Cantidad de alumnos:", cont_alumnos)
		
			finMientras // NFDS(sec)
		
			cerrar (sec)
		finAccion
		