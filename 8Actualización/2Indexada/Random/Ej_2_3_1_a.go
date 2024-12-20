//Ejercicio 2.3.1.a¶

//Dado un fichero secuencial de Facturas, ordenado por Nro. de Cliente y Nro. de Factura, con la siguiente estructura:

//FACTURAS Ordenado por Nro_Cliente, Nro_Factura
//Nro_ClienteNro_FacturaFechaImporte

//Se desea un listado con el siguiente detalle:

//Nro. Cliente	Nombre Cliente	Numero Factura
//Los datos del cliente se encuentran en un fichero indexado por Nro. de Cliente, que tiene la siguiente estructura:

//CLIENTES Indexado por Nro_Cliente
//Nro_ClienteNombreDNICUITDomicilio



ACCION Ejercicio_2_3_1_a ES

	Formato_Fecha= REGISTRO
		dia: 1..31
		mes: 1..12
		año: N(4)
	FINREGISTRO

	Formato_Factura= REGISTRO
		Nro_Cliente: N(5)
		Nro_Factura: N(8)
		Fecha: Formato_FECHA
		Importe: N(7)
	FINREGISTRO

	reg_mae, reg_sal: Formato_Factura
	mae, arch_sal: ARCHIVO de Formato_Factura Ordenado por Nro_CLiente, Nro_Factura

	Formato_Cliente= REGISTRO
		Nro_Cliente: N(5)
		Nombre: AN(15)
		DNI: N(8)
		Domicilio: AN(20)
	FINREGISTRO
