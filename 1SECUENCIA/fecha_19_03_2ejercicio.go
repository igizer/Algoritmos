//Se tiene una secuencia de caracteres que contiene los numeros de telefono
//de los clientes de una empresa,cada número tiene 10 dígitos. La empresa
//necesita organizar los clientes de acuerdo al último dígito de su número
//telefónico, generar una secuencia con los números de las personas que inician
//en un número ingresado por usuario.
ACCION fecha19_03_ej2 ES
	AMBIENTE
		entrada, salida: secuencia de entero
		digito, seleccionado, i: entero
	PROCESO
		ARR(entrada)
		AVZ(entrada, digito)
		CREAR(salida)
		ESC("Elige un dígito por el que organizar")
		Leer(seleccionado)
		MIENTRAS NFDS(entrada) Hacer
			SI digito EN seleccionado ENTONCES
				PARA:i 1 HASTA 10 HACER
					GRABAR(salida, digito)
					AVZ(entrada, digito)
				FINPARA
			SINO
				PARA:i 1 HASTA 10 HACER
					AVZ(entrada, digito)
				FINPARA
		FINMIENTRAS
		ESC("Los clientes fueron organizados con éxito")
		CERRAR(salida)
		CERRAR(entrada)
FINACCION
