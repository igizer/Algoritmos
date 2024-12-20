Ejercicio N° 1 (3 ptos)

Una importante empresa de Turismo de la Ciudad de Resistencia ha realizado a comienzos de año la venta de PAQUETES TURISTICOS para
la temporada de VACACIONES 2024. A todos los clientes que adquirieron paquetes se les permitió la flexibilidad para realizar pagos
durante todo este año, ya que en el último mes se correría un proceso que actualizará el estado de cuenta de cada UTNBook

La información de todos los clientes se encuentra dentro de un archivo secuencial con el siguiente formato:

CLIENTES, ordenado por nro_cliente
Nro_Cliente | apellido_nombre | dni | id_paquete | saldo | estado | categoría | puntos | fecha_baja


El estado podrá ser: 'SALDADO", si el paquete se encuentra pagado en su totalidad; "SALDO A FAVOR' si el cliente, luego del
proceso de recuentos de pagos, tiene un saldo a favor, o bien "DEUDOR" en caso de que el cliente no haya llegado a pagar la
totalidad del costo del paquete turístico

Además, los clientes de le empresa están agrupados por categorias, lo cual les permite acceder a importantes descuentos.
Las categorías pueden ser: SIMPLE, PLATA, ORO, DIAMANTE.

Todos los clientes que se dan de alta por primera vez en la empresa acceden a la categoria SIMPLE.
Los clientes en categoria PLATA tienen una bonificación del 10% y suman 10 puntos por el paquete adquirido.
Los clientes en categoría ORO tienen una bonificación del 15% y suman 20 puntos por el paquete adquirido.
Los clientes en categoría DIAMANTE tienen una bonificación del 20% y además suman 30 puntos por el paquete adquirido.

La información de todos los movimientos de los clientes se encuentra dentro de un archivo secuencial, con el siguiente formato.

NOVEDADES, ordenado por nro_cliente, nro_novedad
Nro_cliente | nro_novedad | apellido_nombre | dni | id _paquete] fecha_novedad | monto

Si el nro_novedad = 0 implica un nuevo cliente que adquirió un paquete, en este caso la cantidad de puntos arrancará en 0, y su saido será DEUDOR

Si el nro novedad = 999, entonces significa que el cliente desea cancelar (dar de baja) su paquete turístico,
y por ende se deberá actualizar además, el campo de fecha_baja (el estado quedará con SALDO A FAVOR).

Los valores intermedios de 1...998 indican los pagos realizados por los clientes. Puede haber mas de una o ninguna novedad para cada cliente.

Además se cuenta con la información de los PAQUETES TURISTICOS en un archivo indexado.

PAQUETES_TURISTICOS, indexado por id_paquete
Id _paquete | nombre | costo | destino

Desarrollar un algoritmo que permita actualizar el archivo de clientes,
informando al usuario por pantalla cualquier tipo de error, e informar al final.

a. La cantidad de clientes que se dieron de baja, y el monto total que debería reintegrar la empresa.
b. El porcentaje de clientes de cada categoria.


ACCION Ejercicio_Actualizacion ES

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

	acl, asal: Archivo de CLIENTES Ordenado por Nro_Cliente
	rcl, aux: CLIENTES

	NOVEDADES = REGISTRO
		Clavem = REGISTRO
			Nro_cliente: N(3)
			nro_novedad: 0..999
		FINREGISTRO
		apellido_nombre: AN(30)
		dni: N(8)
		id _paquete: N(3)
		fecha_novedad: Fecha
		monto: N(7)
	FINREGISTRO

	anov: Archivo de NOVEDADES Ordenado por Clavem
	rnov: NOVEDADES

	PAQUETES_TURISTICOS = REGISTRO
		Id_paquete: N(2)
		nombre: AN(15)
		costo: N(7)
		destino: AN(25)
	FINREGISTRO

	apaquete: Archivo de PAQUETES_TURISTICOS Indexado por Id_paquete
	rpaquete: PAQUETES_TURISTICOS

	PROCESO leercl ES
		LEER(acl, rcl)
		SI FDEA(acl) ENTONCES
			rcl.Nro_Cliente:= HV
		FINSI
	FINPROCESO

	PROCESO leernov ES
		LEER(anov, rnov)
		SI FDEA(anov) ENTONCES
			rnov.Clave.Nro_Cliente:= HV
		FINSI
	FINPROCESO

	//a-
	contclientesbaja, desc: entero
	acummontoreintegrar, costofinal: real

	//b-
	contsimple, contplata, contoro, contdiamante, contcategoria: entero

	PROCESO

		contclientesbaja:= 0
		acummontoreintegrar:= 0

		contsimple:= 0
		contplata:= 0
		contoro:= 0
		contdiamante:= 0
		contcategoria:= 0

		ABRIR E/(acl); ABRIR E/(anov); ABRIR E/(apaquete)
		ABRIR /S(asal)

		leercl(); leernov()

		MIENTRAS (rcl.Nro_Cliente <> HV) v (rnov.Clave.Nro_Cliente <> HV)

			SI rcl.Nro_Cliente < rnov.Clave.Nro_Cliente ENTONCES

				aux:= rcl
				GRABAR(asal, aux)

				SEGUN rcl.categoría HACER
					"SIMPLE": contsimple:= contsimple + 1
					"PLATA": contplata:= contplata + 1
					"ORO": contoro:= contoro + 1
					"DIAMANTE": contdiamante:= contdiamante + 1
				FINSEGUN

				leercl()

			SINO

				SI rcl.Nro_Cliente > rnov.Clave.Nro_Cliente ENTONCES
			
					SI rnov.nro_novedad = 0 ENTONCES

						aux.Nro_Cliente:= rnov.Clavem.Nro_cliente
						aux.apellido_nombre:= rnov.apellido_nombre
						aux.dni:= rnov.dni
						aux.Id_paquete:= rnov.Id_paquete
						rpaquete.Id_paquete:= rnov.Id_paquete
						leer(apaquete, rpaquete)

							aux.saldo:= rpaquete.costo - rnov.monto
							aux.puntos:= 0
							aux.estado:= "DEUDOR"
							aux.categoría:= "SIMPLE"
							aux.Fecha_baja:= " "
							contsimple:= contsimple + 1
						
						leernov()

						MIENTRAS rnov.Nro_Cliente = aux.Nro_Cliente HACER

							SI rpaquete.costo < (aux.saldo + rnov.monto)
								aux.saldo:= rnov.monto - aux.saldo; aux.estado:= "SALDO A FAVOR"; 
								acummontoreintegrar:= acummontoreintegrar + aux.saldo // supongo que hay que reintegrar tmb aun que no sea dado baja
							SINO
								SI rpaquete.costo = (aux.saldo + rnov.monto)
									aux.saldo:= 0; aux.estado:= "SALDADO"
								SINO
									SI rpaquete.costo < (aux.saldo + rnov.monto)
										aux.saldo:= aux.saldo + rnov.monto
									FINSI
								FINSI
							FINSI

							leernov()

						FINMIENTRAS

						GRABAR(asal, aux)

					SINO
						ESC("ERROR CODIGO NOVEDAD ", rnov.nro_novedad , " (cliente inexistente) ")
					FINSI

					leercl()

				SINO

					SI rnov.nro_novedad = 999 ENTONCES

						rpaquete.Id_paquete:= aux.Id_paquete
						leer(apaquete, rpaquete)
					  si exite entonces

						SEGUN aux.categoría HACER

							"SIMPLE": desc:= 0
							"PLATA": desc:= 0.10
							"ORO": desc:= 0.15
							"DIAMANTE": desc:= 0.20

						FINSEGUN


						SEGUN aux.estado HACER

							"DEUDOR": acummontoreintegrar:= acummontoreintegrar + (aux.saldo - rpaquete.costo * desc)
															
							"SALDADO": acummontoreintegrar:= acummontoreintegrar + rpaquete.costo - rpaquete.costo * desc)
										aux.saldo:=	rpaquete.costo - rpaquete.costo * desc							
								
							"SALDO A FAVOR": acummontoreintegrar:= acummontoreintegrar + (aux.saldo + rpaquete.costo - rpaquete.costo * desc)
										aux.saldo:=	aux.saldo + rpaquete.costo - rpaquete.costo * desc

						FINSEGUN

						aux:= rcl
						aux.estado:= "SALDO A FAVOR"
						contclientesbaja:= contclientesbaja + 1
						aux.Fecha_baja:= rnov.Fecha_baja
						leernov()
                     fs
					SINO

						SI rnov.nro_novedad = 0 ENTONCES
							ESC("Error no se puede dar de alta")
							leernov()
						
						SINO

							aux:= cl
							rpaquete.Id_paquete:= aux.Id_paquete
							leer(apaquete, rpaquete)

							SEGUN aux.categoría HACER

								"SIMPLE": desc:= 0; contsimple:= contsimple + 1
								"PLATA": desc:= 0.10; contplata:= contplata + 1
								"ORO": desc:= 0.15; contoro:= contoro + 1
								"DIAMANTE": desc:= 0.20; contdiamante:= contdiamante + 1

							FINSEGUN

							MIENTRAS rcl.Nro_Cliente = rnov.Clave.Nro_Cliente HACER

								aux.saldo:= aux.saldo + rnov.monto

								SI aux.saldo = rpaquete.Costo
									aux.saldo:= 0
									aux.estado:= "SALDADO"
								SINO
									SI aux.saldo > rpaquete.Costo
										aux.saldo:= aux.saldo - rpaquete.Costo
										aux.estado:= "SALDO A FAVOR"
										acummontoreintegrar:= acummontoreintegrar + aux.saldo
									FINSI
								FINSI

								leernov()

							FINMIENTRAS

						FINSI
					FINSI
					GRABAR(asal, aux)
					leercl()
				FINSI
			FINSI
		FINMIENTRAS

		//a-
		ESC("La cantidad a reintegrar entre personas que se dieron a baja y saldos a favor es ", acummontoreintegrar)
		ESC("La cantidad de personas que se dieron de baja son ", contclientesbaja)

		//b- 
		contcategoria:= contsimple + contplata + contoro + contdiamante
		ESC("Porcentaje de Categoría Simple: ", contsimple/contcategoria)
		ESC("Porcentaje de Categoría Plata: ", contplata/contcategoria)
		ESC("Porcentaje de Categoría Oro: ", contoro/contcategoria)
		ESC("Porcentaje de Categoría Diamante: ", contdiamante/contcategoria)

		CERRAR(acl); CERRAR(anov); CERRAR(apaquete)

FINACCION





				
								

								




					

						
						




							





