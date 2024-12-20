Ejercicio 2 (2 ptos)
Luego del proceso de actualización del archivo de CLIENTES, se requiere un informe de totales.
Es necesario conocer estadísticas de cantidad de paquetes adquiridos por categoria de socio,
teniendo en cuenta que la empresa ofrece 12 paquetes turísticos distintos (id_paquete: 1..12).
Considerando además todo lo mencionado en el Ejercicio 1, escribir un algoritmo que permita informar por pantalla

a. Cantidad de paquetes adquiridos por cada categoría de cliente y nombre de paquete
(considerar sólo paquetes saldados, es decir cuyo saldo del cliente sea SALDADO o con SALDO A FAVOR).

b. Cantidad total de paquetes saldados.

c. El paquete con más ventas





ACCION ejercicio_2() ES
	AMBIENTE
		formato_fecha = Registro
			aaa: N(4)
			mm: (1...12)
			dd: (1...31)
		FinRegistro
		formato_clientes = Registro
			numero_cliente:
			apellido_nombre:
			DNI: N(8)
			ID_paquete: (1....12)
			saldo:
			estado: {"SALDADO", "SALDO A FAVOR", "DEUDOR"}
			categoria: {"SIMPLE", "PLATA", "ORO", "DIAMANTE"}
			puntos:
			fecha_baja: formato_fecha
		FinRegistro
		formato_paquetes = Registro
			ID_paquete: (1...12)
			nombre:
			costo:
			destino:
		FinRegistro
		matriz_paquetes: Arreglo de [1...5, 1...13] de enteros
		archivo_maestro: archivo de formato_clientes ordenado por numero_cliente
		registro_maestro: formato_clientes
		archivo_indexado: Archivo de formato_paquetes indexado por ID_paquete
		registro_indexado: formato_paquetes
		contador_saldados, mayor: entero
		Procedimiento Inicializar() Es
			AbrirE/(archivo_maestro)
			AbrirE/(archivo_indexado)
			Para (x:= 1 hasta 5) Hacer
				Para (y:= 1 hasta 13) Hacer
					matriz_paquetes[x, y]:= 0
				FinPara
			FinPara
			contador_saldados:= 0
			mayor:= LV
		FinProcedimiento
		Funcion ConvertirCategoria(variable: alfanumerico): entero Es
			Segun (variable) Hacer
				="SIMPLE": ConvertirCategoria:= 1
				="PLATA": ConvertirCategoria:= 2
				="ORO": ConvertirCategoria:= 3
				="DIAMANTE": ConvertirCategoria:= 4
			FinSegun
		FinFuncion
	PROCESO
		Inicializar()
		Mientras (NFDA(archivo_maestro)) Hacer
			Para (x:= 1 hasta 4) Hacer
				Para (y:= 1 hasta 12) Hacer
					i:= ConvertirCategoria(registro_maestro.categoria)
					j:= registro_maestro.ID_paquete
					Si (registro_maestro.estado <> "DEUDOR") Entonces
						Si (registro_maestro.estado = "SALDADO") Entonces
							contador_saldados:= contador_saldados + 1
						FinSi
						matriz_paquetes[i, j]:= matriz_paquetes[i, j] + 1
						matriz_paquetes[5, j]:= matriz_paquetes[5, j] + 1
						matriz_paquetes[i, 13]:= matriz_paquetes[i, 13] + 1
					FinSi
				FinPara
			FinPara
			Leer(archivo_maestro, registro_maestro)
		FinMientras
		Para (x:= 1 hasta 4) Hacer
			Escribir("El total de paquetes para la categoria ", ConvertirCategoria(x), " es de", matriz_paquetes[x, 13], " paquetes.")
		FinPara
		Para (y:= 1 hasta 12) Hacer
			registro_indexado.ID_paquete:= y
			Leer(archivo_indexado, registro_indexado)
			Escribir("El total de paquetes de ID ", registro_indexado.nombre, " es de ", matriz_paquetes[5, y], " paquetes.")
		FinPara
		Para (y:= 1 hasta 12) Hacer
			Si (matriz_paquetes[4, y] > mayor) Entonces
				mayor:= matriz_paquetes[4, y]
				x:= y
			FinSi
		FinPara
		registro_indexado.ID_paquete:= x
		Leer(archivo_indexado, registro_indexado)
		Escribir("El paquete con más ventas fue el de ", registro_indexado.nombre, " con un total de ", mayor, " ventas.")
		Cerrar(archivo_maestro)
		Cerrar(archivo_indexado)
FINACCION