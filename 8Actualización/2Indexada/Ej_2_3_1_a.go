Ejercicio 2.3.1.a¶
Dado un fichero secuencial de Facturas, ordenado por Nro. de Cliente y Nro. de Factura, con la siguiente estructura:

FACTURAS Ordenado por Nro_Cliente, Nro_Factura

Nro_ClienteNro_FacturaFechaImporte
Se desea un listado con el siguiente detalle:

Nro. Cliente	Nombre Cliente	Numero Factura
Los datos del cliente se encuentran en un fichero indexado por Nro. de Cliente, que tiene la siguiente estructura:

CLIENTES Indexado por Nro_Cliente

Nro_ClienteNombreDNICUITDomicilio

ACCION Ej_2_3_1_a_Index ES

	AMBIENTE

		Formato_Fecha = registro
			dia: 1..31
			mes: 1..12
			año: N(4)
		FINREGISTRO

		Formato_Factura = registro
			Clave= registro
				Nro_Cliente: N(5)
				Nro_Factura: N(9)
			FINREGISTRO
			Fecha: Formato_Fecha
			Importe: N(9)
		FINREGISTRO

		rfac: Formato_Factura
		afac: Archivo de Formato_Factura Ordenado por Clave

		Formato_Clientes = registro
			Nro_Cliente: N(5)
			Nombre: AN(25)
			DNI: N(8)
			CUIT: N(11)
			Domicilio: AN(25)
		FINREGISTRO

		rcli: Formato_Cliente
		acli: Archivo de Formato_Cliente Indexado por Nro_Cliente
		resgnrocliente: N(5)

	PROCESO

		ABRIR E/(afac); ABRIR E/(acli)
		LEER(afac, rfac)
		ESC("|Nro. Cliente| |Nombre Cliente| |Nro. Factura|")

		MIENTRAS NFDA(afac) HACER

			resgnrocliente:= rfac.Clave.Nro_Cliente
			rcli.Nro_Cliente:= rfac.Clave.Nro_Cliente
			LEER(acli, rcli.Nro_Cliente)

			SI EXISTE ENTONCES

				MIENTRAS NFDA(afac) y (resgnrocliente = rfac.Clave.Nro_Cliente) HACER
					ESC(rfac.Clave.Nro_Cliente, " ", rcli.Nombre_Cliente, " ", rfac.Clave.Nro_Factura)
					LEER(afac, rfac)
				FINMIENTRAS

			SINO

				MIENTRAS NFA(afac) y (resgnrocliente = rfac.Clave.Nro_Cliente) HACER
					LEER(afac, rfac)
				FINMIENTRAS

			FINSI
		
		FINMIENTRAS

		CERRAR(afac); CERRAR(acli)

FINACCION







ACCION Ej_2_3_1_a_Index ES

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
			Nro_Factura: N(7)

		FINREGISTRO

		acliente: Archivo de F_Cliente Indexado por Nro_Cliente
		rcliente: F_Cliente

		resg_nrocliente: N(5)

	
	PROCESO

		ABRIR E/(afactura); ABRIR E/(acliente)
		LEER(afactura, rfactura)
		ESC("|Nro_Cliente| |Nombre Cliente| |Numero Factura|");

		MIENTRAS NFDA(afactura) HACER

			resg_nrocliente:= rfactura.Nro_Cliente
			rcliente.Nro_Cliente:= rfactura.Nro_Cliente
			LEER(acliente, rcliente)

			SI EXISTE ENTONCES

				MIENTRAS NFDA(afactura) y (resg_nrocliente = rfactura.Nro_Cliente)
					ESC(rfactura.Nro_Cliente, " ", rcliente.Nombre_Cliente, " ", rfactura.Nro_Factura)
					LEER(afactura, rgfactura)
				FINMIENTRAS
			
			SINO

				MIENTRAS NFDA(afactura) y (resg_nrocliente = rfactura.Nro_Cliente)
					LEER(afactura, rgfactura)
				FINMIENTRAS

			FINSI

		FINMIENTRAS

		CERRAR(afactura); CERRAR(acliente)

FINACCION













ACCION Act_Indexada_2.3.1.a ES
    AMBIENTE

    FACTURAS = REGISTRO
        NRO_CLIENTE: N(5)
        NRO_FACTURA: N(5)
        FECHA = REGISTRO    
            AA:N(4)
            MM: 1..12
            DD: 1..31
        FR;
        IMPORTE: N(8)
    FR;

    AFACTURAS: Archivo de FACTURAS ordenado por NRO_CLIENTE, NRO_FACTURA;
    RFACTURAS: FACTURAS;

    CLIENTES = REGISTRO
        NRO_CLIENTE: N(5)
        NOMBRE: AN(20)
        DNI: N(8)
        CUIT: N(11)
        DOMICILIO: AN(30)
    FR;
    
    ACLIENTES: Archivo indexado por NRO_CLIENTE;
    RCLIENTES: CLIENTES

    RES_NROCLIENTE: N(5)
    

    PROCESO
    ABRIR E/(AFACTURAS); ABRIR E/(ACLIENTES);
    LEER(AFACTURAS,RFACTURA)
    ESC("|Nro_Cliente| |Nombre Cliente| |Numero Factura|");
    MIENTRAS NFDA(AFACTURAS) HACER
        
        RES_NROCLIENTES:= RFACTURAS.NRO_CLIENTE
        RCLIENTES.NRO_CLIENTE:= RFACTURAS.NRO_CLIENTE
        LEER(ACLIENTES,RCLIENTES);

        SI EXISTE ENTONCES
           
            MIENTRAS NFDA (AFACTURAS) Y (RES_NROCLIENTE =  RFACTURA.NRO_CLIENTE) HACER
                ESC(RFACTURAS.NRO_CLIENTE, RCLIENTES.NOMBRE, RFACTURAS.NRO_FACTURA);
                LEER(AFACTURA,RFACTURA);
            FM;
        SINO
             MIENTRAS NFDA (AFACTURAS) Y (RES_NROCLIENTE =  RFACTURA.NRO_CLIENTE) HACER
                LEER(AFACTURAS,RFACTURAS);
            FM 
        FS;      
    FM;
    CERRAR(AFACTURAS); CERRAR(ACLIENTES)
Fin_accion.
