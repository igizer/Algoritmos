Ejercicio 2.3.1.b¶
Dado un fichero secuencial de Facturas, ordenado por Nro. de Cliente y Nro. de Factura, con la siguiente estructura:

FACTURAS Ordenado por Nro_Cliente, Nro_Factura

Nro_ClienteNro_FacturaFechaImporte
Se desea un listado con el siguiente detalle:

Nro. Cliente	Nombre Cliente	Total Facturado	Cantidad de Facturas
Los datos del cliente se encuentran en un fichero indexado por Nro. de Cliente, que tiene la siguiente estructura:

CLIENTES Indexado por Nro_Cliente

Nro_ClienteNombreDNICUITDomicilio


ACCION Ejercicio2_3_1_b ES

	AMBIENTE

		Formato_Fecha = REGISTRO
			dia: 1..31
			mes: 1..12
			año: N(4)
		FINREGISTRO

		Formato_Factura= REGISTRO
			Clave= REGISTRO
				Nro_Cliente: N(5)
				Nro_Factura: N(5)
			FINREGISTRO
			Fecha: Formato_Fecha
			Importe: N(9)
		FINREGISTRO

		rfac: Formato_Factura
		afac: Archivo de Formato_Factura Ordenado por Clave

		Formato_Cliente= REGISTRO
			Nro_Cliente: N(5)
			Nombre: AN(25)
			DNI: N(8)
			CUIT: N(11)
			Domicilio: AN(30)
		FINREGISTRO

		rcli: Formato_Cliente
		acli: Archivo de Formato_Cliente Indexado por Nro_Cliente
		resgnrocliente: N(5)

		contfac, acumfac: entero

	PROCESO

		ABRIR E/(afac); ABRIR E/(acli)
		LEER(afac, rfac)

		ESC("|Nro. Cliente| |Nombre Cliente| |Total Facturado| |Cantidad de Facturas|")

		MIENTRAS NFDA(afac)
		
			contfac:= 0
			acumfac:= 0

			resgnrocliente:= rfac.Clave.Nro_Cliente
			rcli.Nro_Cliente:= rfac.Clave.Nro_Cliente
			LEER(acli, rcli)

			SI EXISTE ENTONCES

				MIENTRAS NFDA(afac) y (resgnrocliente = rfac.Clave.Nro_Cliente) HACER

					contfac:= contfac + 1
					acumfac:= acumfac + rfac.Importe

					LEER(afac, rfac)
				
				FINMIENTRAS

				ESC(resgnrocliente, " ", rcli.Nombre, " ", acumfac, " ", contfac)

			SINO

				MIENTRAS NFDA(afac) y (resgnrocliente = rfac.Clave.Nro_Cliente) HACER

					LEER(afac, rfac)

				FINMIENTRAS
			
				ESC("El número de cliente no existe")

			FINSI

		FINMIENTRAS

		CERRAR(afac); CERRAR(acli)

FINACCION























ACCION Ej_2_3_1b ES

	AMBIENTE

		F_Fecha = REGISTRO

			dd: 1..31
			mm: 1..12
			aaaa: N(4)

		FINREGISTRO

		F_Factura = REGISTRO

			Clave = REGISTRO
				Nro_Cliente: N(5)
				Nro_Factura: N(7)
			FINREGISTRO
			Fecha: F_Fecha
			Importe: N(6)

		FINREGISTRO

		afactura: Archivo de F_Factura Ordenado por Clave
		rfactura: F_Factura

		F_Cliente = REGISTRO

			Nro_Cliente: N(5)
			Nombre_Cliente: AN(25)
			DNI: N(8)
			CUIT: N(11)
			Domicilio: AN(30)

		FINREGISTRO

		acliente: Archivo de F_Cliente Indexado por Nro_Cliente
		rcliente: F_Cliente

		resg_nrocliente: N(5)
		cont_factura, tot_facturado: entero

		
	PROCESO

		ABRIR E/(afactura); ABRIR E/(acliente)
		LEER(afactura, rfactura)
		ESC("|Nro_Cliente| |Nombre Cliente| |Total Facturado| |Cantidad de Facturas|");

		MIENTRAS NFDA(afactura) HACER

			cont_factura:= 0
			tot_facturado:= 0

			resg_nrocliente:= rfactura.Nro_Cliente
			rcliente.Nro_Cliente:= rfactura.Nro_Cliente
			LEER(acliente, rcliente)

			SI EXISTE ENTONCES

				MIENTRAS NFDA(afactura) y (resg_nrocliente = rfactura.Nro_Cliente)
					cont_factura:= cont_factura + 1
					tot_facturado:= tot_facturado + rfactura.Importe
					LEER(afactura, rgfactura)
				FINMIENTRAS

				ESC(rfactura.Nro_Cliente, " ", rcliente.Nombre_Cliente, " ", tot_facturado, " ", cont_factura)
				
			SINO

				MIENTRAS NFDA(afactura) y (resg_nrocliente = rfactura.Nro_Cliente)
					LEER(afactura, rgfactura)
				FINMIENTRAS

				ESC("El número del cliente no existe")

			FINSI



		FINMIENTRAS

		CERRAR(afactura); CERRAR(acliente)

FINACCION