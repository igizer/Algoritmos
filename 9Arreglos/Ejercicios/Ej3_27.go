Ejercicio 3.27¶
Se desea efectuar una estadística de ventas. Se cuenta para ello con una secuencia de las facturas emitidas,
las cuales son identificadas por un Número; dicha secuencia contiene información relativa al cliente: su Número y
Zona a la cual pertenece, como así también el Tipo de mercadería entregada, la Cantidad de unidades, el total gravado,
el total exento de IVA y el valor del IVA. Construya un algoritmo que emita por zona, y dentro de la zona por tipo de
mercadería el total de unidades vendidas, el total gravado, el total exento y el total de IVA, y además un total general
con la misma información. Hay 9 zonas y 4 tipos de mercadería.


ACCION Ej_3_27 ES

	AMBIENTE

		Tipo_Mercaderia = REGISTRO
			Num: N(7)
			Zon: 1..9
			Merc: 1..4
			Cant: N(4)
			Tot_Gravado: N(4)
			Tot_ExIVA: N(4)
			Tot_CIVA: N(4)
		FINREGISTRO

		reg: Tipo_Mercaderia
		arch: Archivo de Tipo_Mercaderia

		regArreglo = REGISTRO
			Tot_Gravado: N(4)
			Tot_ExIVA: N(4)
			Tot_CIVA: N(4)
		FINREGISTRO

		A: Arreglo [1..10,1..5] de regArreglo

	
	PROCESO

		ABRIR E/(arch); LEER(arch, reg)

		//Pongo todo en 0
		PARA i:= 1 HASTA 10 HACER // 10 Total por tipo de mercaderia
			PARA j:= 1 HASTA 5 HACER // 5 Total por Zona
				A[i,j].Tot_Gravado:= 0
				A[i,j].Tot_ExIVA:= 0
				A[i,j].Tot_CIVA:= 0
			FINPARA
		FINPARA

		MIENTRAS NFDA(arch) HACER

			i:= reg.Zon
			j:= reg.Merc

			//Total Gravado
			A[i,j].Tot_Gravado:= A[i,j].Tot_Gravado + reg.Tot_Gravado
			A[10,j].Tot_Gravado:= A[10,j].Tot_Gravado + reg.Tot_Gravado //Total por Tipo de Mercaderia
			A[i,5].Tot_Gravado:= A[i,5].Tot_Gravado + reg.Tot_Gravado //Total por Zona
			A[10,5].Tot_Gravado:= A[10,5].Tot_Gravado + reg.Tot_Gravado //Total General por Tipo de Mercaderia y zona

			//Total Exedente IVA
			A[i,j].Tot_ExIVA:= A[i,j].Tot_ExIVA + reg.Tot_ExIVA
			A[10,j].Tot_ExIVA:= A[10,j].Tot_ExIVA + reg.Tot_ExIVA //Total por Tipo de Mercaderia
			A[i,5].Tot_ExIVA:= A[i,5].Tot_ExIVA + reg.Tot_ExIVA //Total por Zona
			A[10,5].Tot_ExIVA:= A[10,5].Tot_ExIVA + reg.Tot_ExIVA //Total General por Tipo de Mercaderia y zona

			//Total Con IVA
			A[i,j].Tot_CIVA:= A[i,j].Tot_CIVA + reg.Tot_CIVA
			A[10,j].Tot_CIVA:= A[10,j].Tot_CIVA + reg.Tot_CIVA //Total por Tipo de Mercaderia
			A[i,5].Tot_CIVA:= A[i,5].Tot_CIVA + reg.Tot_CIVA //Total por Zona
			A[10,5].Tot_CIVA:= A[10,5].Tot_CIVA + reg.Tot_CIVA //Total General por Tipo de Mercaderia y zona

			LEER(arch,reg)
		
		FINMIENTRAS

		CERRAR(arch)

		PARA i:= 1 HASTA 9 HACER
			PARA j:= 1 HASTA 4 HACER

				ESC("Para la Zona ", i, " y mercadería de tipo ", j)
				ESC("El Total Gravado es ", A[i,j].Tot_Gravado)
				ESC("El Total exedente de IVA es ", A[i,j].Tot_ExIVA)
				ESC("El Total con IVA es ", A[i,j].Tot_CIVA)

			FINPARA
		FINPARA

		PARA j:= 1 HASTA 4 HACER

			ESC("Para el tipo de mercadería ", j, " hay un total de Gravado de ", A[10,j].Tot_Gravado)
			ESC("Para el tipo de mercadería ", j, " hay un total de exedente de IVA de ", A[10,j].Tot_ExIVA)
			ESC("Para el tipo de mercadería ", j, " hay un total con IVA de ", A[10,j].Tot_CIVA)
		
		FINPARA

		PARA i:= 1 HASTA 9 HACER

			ESC("Para la zona ", i, " hay un total de Gravado de ", A[i,5].Tot_Gravado)
			ESC("Para la zona ", i, " hay un total de exedente de IVA de ", A[i,5].Tot_ExIVA)
			ESC("Para la zona ", i, " hay un total con IVA de ", A[i,5].Tot_CIVA)
		
		FINPARA

		ESC("El Total General entre Zona y Tipo de Mercadería en cuento a Total de Gravado es de ", A[10,5].Tot_Gravado)
		ESC("El Total General entre Zona y Tipo de Mercadería en cuento a Total de exedente de IVA es de ", A[10,5].Tot_ExIVA)
		ESC("El Total General entre Zona y Tipo de Mercadería en cuento a Total con IVA es de ", A[10,5].Tot_CIVA)
	
	CERRAR(arch)

FINACCION