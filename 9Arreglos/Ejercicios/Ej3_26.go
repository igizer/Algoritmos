Ejercicio 3.26¶

Se dispone de un archivo secuencial de facturas de una empresa de energía eléctrica, correspondientes a un año, con el siguiente formato:
Nro_Factura| Nro_Usuario| Zona| Fecha| Consumo


Se desea obtener un cuadro estadístico que informe los consumos y los importes totales
facturados por mes discriminados por zona, con el siguiente formato:


		Zona A	Zona B	Zona C	Zona D	Total por mes
Enero	...		...		...		...		...
Febrero	...		...		...		...		...
...		...		...		...		...		...
Diciembr...		...		...		...		...
TotXzona...		...		...		...		Total General


La ciudad está dividida en 4 Zonas (A,B,C,D), la tarifa por Kw. que cobra la empresa en cada una de ellas es:
Zona A:0.05
Zona B:0.07
Zona C:0.09
Zona D:0.13



ACCION Ejercicio_3_26 ES

	AMBIENTE

		i, j: entero
		Tarifa: real

		Formato_Fecha = REGISTRO
			dia: 1..31
			mes: 1..12
			año: entero(4)
		FINREGISTRO

		Formato_Facturas
			Nro_Factura: entero(10)
			Nro_Usuario: entero(8)
			Zona: ("A", "B", "C", "D")
			Fecha: Formato_Fecha
			Consumo: N(7,2)
		FINREGISTRO

		regfac: Formato_Facturas
		archfac: Archivo de Facturas

		Formato_Arreglo= REGISTRO
			Consumo: N(7,2)
			Importe: entero(9)
		FINREGISTRO

		A: Arreglo [1..13,"A".."E"] de Formato_Arreglo



	PROCESO

		ABRIR E/(archfac); LEER(archfac, regfac)

		PARA i:= 1 HASTA 13 HACER
			PARA j:= "A" HASTA "E" HACER
				A[i,j].Consumo:= 0
				A[i,j].Importe:= 0
			FINPARA
		FINPARA


		MIENTRAS NFDA(archfac) HACER

			PARA i:= 1 HASTA 13 HACER
				PARA j:= "A" HASTA "E" HACER

					SEGUN j=
						"A": Tarifa:= regfac.Consumo*0.05
						"B": Tarifa:= regfac.Consumo*0.07
						"C": Tarifa:= regfac.Consumo*0.09
						"D": Tarifa:= regfac.Consumo*0.13
					FINSEGUN

					A[i,j].Consumo:= A[i,j].Consumo + regfac.Consumo
					A[i,"E"].Consumo:= A[i,"E"].Consumo + regfac.Consumo //Total por mes
					A[13,j].Consumo:= A[13,j].Consumo + regfac.Consumo //Total por Zona
					A[13,"E"].Consumo:= A[13,"E"].Consumo + regfac.Consumo //Total General

					A[i,j].Importe:= A[i,j].Importe + Tarifa
					A[i,"E"].Importe:= A[i,"E"].Importe + Tarifa //Total por mes
					A[13,j].Importe:= A[13,j].Importe + Tarifa //Total por Zona
					A[13,"E"].Importe:= A[13,"E"].Importe + Tarifa //Total General

				FINPARA
			FINPARA

		FINMIENTRAS

		PARA i:= 1 HASTA 12 HACER
			PARA j:= "A" HASTA "D" HACER

				ESC("Para el mes nro ", i, " en la zona ", j)
				ESC("El consumo total en Kw fue de ", A[i,j].Consumo)
				ESC("El importe total fue de ", A[i,j].Importe)
			
			FINPARA
		FINPARA

		PARA i:= 1 HASTA 12 HACER
			ESC("Para el mes nro", i)
			ESC("El total general entre todas las zonas en Consumo (Kw) fue de ", A[i,"E"].Consumo)
			ESC("El total general entre todas las zonas en Importe fue de ", A[i,"E"].Importe)
		FINPARA

		PARA i:= "A" HASTA "E" HACER
			ESC("Para la zona", j)
			ESC("El total general entre todos los meses en Consumo (Kw) fue de ", A[13,j].Consumo)
			ESC("El total general entre todas los meses en Importe fue de ", A[13,j].Importe)
		FINPARA

		ESC("El total general entre todas las zonas durante el año:")
		ESC("En cuanto a Consumo (Kw) total fue de ", A[13,"E"].Consumo)
		ESC("En cuanto a Importe total fue de ", A[13,"E"].Importe)

		CERRAR(archfac)
	
FINACCION










ACCION Ej_3_26__Proceso_Estadístico_1 ES

	AMBIENTE

		

		i,j: enteros
		Tarifa: real


		Formato_Fecha= REGISTRO
			dd: 1..31
			mm: 1..12
			aaaa: N(4)
		FINREGISTRO

		Empresa= REGISTRO
			Nro_Factura: N(10)
			Nro_Usuario: N(7)
			Zona: ("A","B","C","D")
			Fecha: Formato_Fecha
			Consumo: N(7,2)
		FINREGISTRO

		reg: Empresa
		arch: Archivo de Empresa

		REGIS = REGISTRO
			Consumo: N(7,2)
			Importe: N(7,2)
		FINREGIS

		A: Arreglo[1..13,"A".."E"] de REGIS


	PROCESO

		ABRIR E/(arch)
		LEER(arch, reg)

		PARA i:= 1 HASTA 13 HACER // 13 ES TOTAL ZONA
			PARA j:= "A" HASTA "D" HACER // D ES TOTAL MES
				A[i,j]:= 0
			FINPARA
		FINPARA

		MIENTRAS NFDA(arch) HACER

			PARA i:= 1 HASTA 13 HACER
				PARA j:= 1 HASTA 5 HACER

					SEGUN j HACER
						"A": Tarifa:= reg.Consumo * 0.05
						"B": Tarifa:= reg.Consumo * 0.07
						"C": Tarifa:= reg.Consumo * 0.09
						"D": Tarifa:= reg.Consumo * 0.13
					FINSEGUN

					//Consumo
					A[i,j].Consumo:= A[i,j].Consumo + reg.Comsumo
					A[i,"E"].Consumo:= A[i,"E"].Consumo + reg.Consumo //Total por mes
					A[13,j].Consumo:= A[13,j].Consumo + reg.Comsumo //Total por zona
					A[13,"E"].Consumo:= A[13,"E"].Consumo + reg.Comsumo //Total General

					//Importe
					A[i,j].Importe:= A[i,j].Importe + reg.Comsumo * Tarifa
					A[i,"E"].Importe:= A[i,"E"].Importe + reg.Consumo * Tarifa //Total por mes
					A[13,j].Importe:= A[13,j].Importe + reg.Comsumo * Tarifa //Total por zona
					A[13,"E"].Importe:= A[13,"E"].Importe + reg.Comsumo * Tarifa //Total General

				FINPARA
			FINPARA

		FINMIENTRAS

		//Ahora Informo

		PARA i:= 1 HASTA 12 HACER
			PARA j:= "A" HASTA "D" HACER

				ESC("Para el mes: ", i, " en la zona ", j)
				ESC("El consumo en kw fue de: ", A[i,j].Consumo)
				ESC("Y el importe total fue de: ", A[I,J].Importe)
			

			FINPARA
		FINPARA	

		PARA j:= "A" HASTA "D" HACER

			ESC("El total del consumo de la zona ", j, " es de: ", A[13,j].Consumo)
			ESC("El total de importe de la zona ", j, " es de: ", A[13,j].Immporte)

		FINPARA

		PARA i:= 1 HASTA 12 HACER

			ESC("El total del consumo del mes ", i " es de: ", A[i,"E"].Consumo)
			ESC("El total de importe del mes ", i " es de: ", A[i,"E"].Immporte)

		FINPARA

		CERRAR(arch)

FINACCION