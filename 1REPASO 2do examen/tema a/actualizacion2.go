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

	ambiente

		clientes= registro
		nrocl: N(8)
		ApNom: AN(30)
		dni: N(8)
		idpack: 1..12
		saldo: N(6)
		estado: ('SALDADO', 'SALDO A FAVOR', 'DEUDOR')
		categoria: ('SIMPLE', 'PLATA', 'ORO', 'BRONCE')
		puntos: N(4)
		fecha_baja: fecha
		finregistro

		client, clientsld: archivo de clientes ordenado por nrocl
		cl, aux: clientes

		novedades= registro
		nrocl: N(8)
		nronov: N(3)
		ApNom: AN(30)
		dni: N(8)
		idpack: 1..12
		fechanov: fecha
		monto: N(6)
		finregistro

		noved: archivo de novedades ordenado por nrocl, nronov
		nv: novedades

		paquetur = registro
		idpack: 1..12
		nombre: AN(30)
		costo: N(7)
		destino: AN (30)
		finregistro

		pack: archivo de paquetur indexado por idpack
		pk: paquetur
		clibaja, montret, totcli, clisim, cliplat, clioro, clidiam: entero
		desc: real

		procedimiento leercl() es
		leer(client, cl)
		si FDA (client) entonces
			cl.nrocl:= HV
		finsi
		finprocedimiento

		procedimiento leernv() es
		leer (noved, nv)
		si FDA (noved) entonces
			nv.nrocl:= HV
		finsi
		finprocedimiento


	proceso

		abrir e/ (client); leercl()
		abrir e/ (noved); leernv()
		abrir /s (clientsld); abrir e/ (pack);

		clibaja:= 0; montret:= 0; totcli:= 0; clisim:= 0; cliplat:= 0; clioro:= 0; clidiam:= 0;

		mientras (cl.nrocl <> HV) o (nv.nrocl <> HV) hacer

			si cl.nrocl < nv.nrocl hacer

				aux:= cl
				esc (clientsld, aux)
				leercl()

			sino

				si cl.nrocl > nv.nrocl

					si nv.nronov = 0 entonces

						aux.nrocl:= nv.nrocl
						aux.apnom:= nv.apnom
						aux.dni:= nv.dni
						aux.idpack:= nv.idpack
						pk.idpack:= nv.idpack
						leer (pack, pk)
						aux.saldo:= pk.costo - nv.monto
						aux.estado:= 'DEUDOR'
						aux.categoria:= 'SIMPLE'
						aux.puntos:= 0
						fecha_baja:= ''
						clisim:= clisim + 1
						leernv()

						mientras aux.nrocl = nv.nrocl hacer
						
							segun nv.monto hacer

								>cl.saldo: aux.saldo:= nv.monto - aux.saldo; aux.estado:= 'SALDO A FAVOR'
								<cl.saldo: aux.saldo:= aux.saldo - nv.monto; aux.estado:= 'DEUDOR'
								=cl.saldo: aux.saldo:= 0; aux.estado:= 'SALDADO'

							finsegun

							leernv()

						finmientras

						escribir (clientsld, aux)

					sino

						esc ('Error: Cliente inexistente')

					finsi

					leercl()

				sino

					si nv.nronov = 999 entonces

						clibaja:= clibaja + 1
						aux:= cl
						aux.estado:= 'SALDO A FAVOR'
						aux.fecha_baja:= nv.fechanov
						pk.idpack:= nv.idpack

						segun aux.categoria hacer

							'SIMPLE': desc:= 0
							'PLATA': desc:= 10
							'ORO': desc:= 15
							'DIAMANTE': desc:= 20

						finsegun

						leer (pack, pk)

						segun cl.estado hacer

							'SALDO A FAVOR':aux.saldo:= pk.costo - pk.costo * (desc/100) + aux.saldo
							'DEUDOR': aux.saldo:=pk.costo - pk.costo * (desc / 100) - aux.saldo
							'SALDADO': aux.saldo:= pk.costo - pk.costo * (desc/100)

						finsegun

						montret:= montret + aux.saldo
						esc (clientsld, aux)
						leernv()

					sino

						si nv.nronov = 0 entonces

							esc ('Error: Alta imposible')

						sino

							aux:= cl

							segun aux.categoria hacer

								'SIMPLE': desc:= 0; clisim:= clisim + 1;
								'PLATA': desc:= 10; cliplat:= cliplat + 1; aux.puntos:= aux.puntos + 10
								'ORO': desc:= 15; clioro:= clioro + 1; aux.puntos:= aux.puntos + 20;
								'DIAMANTE': desc:= 20; clidiam:= clidiam + 1; aux.puntos:= aux.puntos + 30

							finsegun

							mientras cl.nrocl = nv.nrocl hacer

								pk.idpack:= nv.idpack
								leer (pack, pk)

								segun nv.monto hacer

									>cl.saldo: aux.saldo:= nv.monto - aux.saldo; aux.estado:= 'SALDO A FAVOR'
									<cl.saldo: aux.saldo:= aux.saldo - nv.monto; aux.estado:= 'DEUDOR'
									=cl.saldo: aux.saldo:= 0; aux.estado:= 'SALDADO'

								finsegun

								escribir (clientsld, aux)
								leernv()

							finmientras

						finsi

						leercl()

					finsi

				finsi

			finsi

		finmientras

		esc ('Se dieron de baja un total de ', clibaja, 'clientes')
		esc ('La empresa debera reintegrar un monto total de $', montret)		
		esc ('Porcentaje de clientes de cada categoria:')
		esc ('Simple:', totcli * (clisim/100))
		esc ('Plata:', totcli * (cliplat/100))
		esc ('Oro:', totcli * (clioro/100))
		esc ('Diamante:', totcli * (clidiam/100))	
		cerrar (client); cerrar (clientsld); cerrar (noved); cerrar (pack);

Finalgoritmo				