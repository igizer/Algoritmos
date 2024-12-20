Ejercicio 2 (2 ptos)
Luego del proceso de actualización del archivo de CLIENTES, se requiere un informe de totales.
Es necesario conocer estadísticas de cantidad de paquetes adquiridos por categoria de socio,
teniendo en cuenta que la empresa ofrece 12 paquetes turísticos distintos (id_paquete: 1..12).
Considerando además todo lo mencionado en el Ejercicio 1, escribir un algoritmo que permita informar por pantalla

a. Cantidad de paquetes adquiridos por cada categoría de cliente y nombre de paquete
(considerar sólo paquetes saldados, es decir cuyo saldo del cliente sea SALDADO o con SALDO A FAVOR).

b. Cantidad total de paquetes saldados.

c. El paquete con más ventas // interpreto como que no pidió darse de baja, hayan terminado de pagar o no



// i: categorias de socio (4 + 1)
// j: paquetes turísticos (12 + 1)


ACCION Ejercicio_Arreglo ES

	AMBIENTE

		consignas_arreglo = REGISTRO

			adquirido: N(3)
			saldado: N(3)
			nodeudores: N(3)

		FINREGISTRO

		A: Arreglo [1..5, 1..13] de consignas_arreglo

		//a es arreglo nodeudores
		
		//b es arreglo 5,13
		

		//c es paquete con más ventas, tomo los no dados de baja de cualquier estado
		masventas: 1..12
		cantventas: entero
		
		

		Fecha = REGISTRO
			dd: 1..31
			mm: 1..12
			aaaa: N(4)
		FINREGISTRO

		CLIENTES = REGISTRO
			Nro_Cliente: N(3)
			apellido_nombre: AN(30)
			dni: N(8)
			id_paquete: N(3)
			saldo: N(7)
			estado: ("SALDADO", "DEUDOR", "SALDO A FAVOR")
			categoría: ("SIMPLE", "PLATA", "ORO", "DIAMANTE")
			puntos: N(3)
			fecha_baja: Fecha
		FINREGISTRO

		acl: Archivo de CLIENTES Ordenado por Nro_Cliente
		rcl: CLIENTES

		PAQUETES_TURISTICOS = REGISTRO
			Id_paquete: N(2)
			nombre: AN(15)
			costo: N(7)
			destino: AN(25)
		FINREGISTRO

		apaquete: Archivo de PAQUETES_TURISTICOS Indexado por Id_paquete
		rpaquete: PAQUETES_TURISTICOS

		PAQUETES_TURISTICOS = REGISTRO
			Id_paquete: 1..12
			nombre: AN(15)
			costo: N(7)
			destino: AN(25)
		FINREGISTRO

		apaquete: Archivo de PAQUETES_TURISTICOS Indexado por Id_paquete
		rpaquete: PAQUETES_TURISTICOS

		adquirido: ("SALDADO", "SALDO A FAVOR")
		i: rcl.estado
		j: rcl.id_paquete
	
	PROCESO

		ABRIR E/(acl); ABRIR E/(apaquete)
		LEER(acl, rcl)
		
		cantventas:= 0
		


		PARA i:= 1 HASTA 5 HACER
			PARA j:= 1 HASTA 13 HACER
				A[i,j]:= 0
			FINPARA
		FINPARA

		MIENTRAS NFDA(acl) HACER

			PARA i:= 1 HASTA 5 HACER
				PARA i:= 1 HASTA 13 HACER



					SI cl.Fecha_baja.dd <> " " ENTONCES //gente que no se dió de baja

						SI rcl.estado = "SALDADO" ENTONCES
							A[i,j].saldado:= A[i,j].saldado + 1
							A[5,j].saldado:= A[5,j].saldado + 1 
							A[i,13].saldado:= A[i,13].saldado + 1
							A[5,13].saldado:= A[5,13].saldado + 1 // <- Lo importante para b
						FINSI

						A[i,j].adquirido:= A[i,j].adquirido + 1
						A[5,j].adquirido:= A[5,j].adquirido + 1
						A[i,13].adquirido:= A[i,13].adquirido + 1
						A[5,13].adquirido:= A[5,13].adquirido + 1

						SI cl.estado <> "DEUDORES" ENTONCES
							A[i,j].nodeudores:= A[i,j].nodeudores + 1
							A[5,j].nodeudores:= A[5,j].nodeudores + 1
							A[i,13].nodeudores:= A[i,13].nodeudores + 1
							A[5,13].nodeudores:= A[5,13].nodeudores + 1
						FINSI

						
					FINSI

					LEER(acl, rcl)

				FINPARA
			FINPARA
		FINMIENTRAS


		//informo todo
		PARA i:= 1 HASTA 4 HACER // consigna a
			PARA j:= 1 HASTA 12 HACER
				rpaquete.Id_paquete:= j
				LEER(apaquete, rpaquete)
				ESC("Para la categoría de socio ", i, " paquete ", rpaquete.nombre, " la cantidad de paquetes adquiridos es de ", A[i,j].nodeudores)
			FINPARA
		FINPARA

		

		//abajo, b) total general de paquetes saldados
		ESC("La cantidad total de paquetes SOLO SALDADOS (no saldo a favor por lo que interpreto) es ", A[5,13].saldado)



		PARA i:= 1 HASTA 4
			ESC("Para la categoría de socio ", i, " el total de paquetes vendido es ", A[i,13].adquirido)
		FINPARA


		//abajo, el paquete con más ventas
		PARA j:= 1 HASTA 12 HACER
			SI A[5,j].adquirido > cantventas ENTONCES
				cantventas:= A[5,j].adquirido
				masventas:= j
			FINSI
		FINPARA
		ESC("El paquete con más ventas entre todos los estados es el numero ", masventas)

		CERRAR(acl); CERRAR(apaquete)

FINACCION
