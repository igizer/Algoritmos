Ejercicio 3.24¶
Se cuenta con información acerca de los cajeros de un Supermercado, el cual se halla estructurado en 10 cajas registradoras.
Dicha información está registrada en una secuencia que contiene Apellido y Nombre del empleado,
número de caja que tiene asignada, importe facturado y horario de facturación; la secuencia está
ordenada alfabéticamente por Apellido y Nombre. Se solicita una estadística de los importes facturados,
discriminado por número de caja y franja de horas y además los montos totales, según el siguiente formato:

Cajas			8-10	10-12	12-16	16-18	18-20	Total x Cajas
1				...		...		...		...		...		...
..				...		...		...		...		...		...
10				...		...		...		...		...		...
Total x horas	...		...		...		...		...		...




//rehacer
ACCION Ejercicio_3_24 ES

	AMBIENTE

		Formato_Empleado = REGISTRO
			ApNom: AN(25)
			NroCaja: 1..10
			Facturado: N(6)
			Horario: 8..20
		FINREGISTRO

		remp: Formato_Empleado
		aemp: Archivo de Formato_Empleado Ordenado por ApNom
		i, j: entero
		HORARIOS: AN(15)

		Cajeros: Arreglo[1..11,1..6] de reales

	PROCESO

		ABRIR E/(aemp)
		LEER(aemp, remp)

		PARA i:=1 HASTA 11 HACER
			PARA j:=1 HASTA 6 HACER
				Cajeros[i,j]:= 0
			FINPARA
		FINPARA

		MIENTRAS NFDA(aemp) HACER

			SEGUN remp.NroCaja:
				1: i:= 1
				2: i:= 2
				3: i:= 3
				4: i:= 4
				5: i:= 5
				6: i:= 6
				7: i:= 7
				8: i:= 8
				9: i:= 9
				10: i:= 10
			FINSEGUN

			SEGUN remp.Horario:
				8..9: j:= 1
				10..11: j:= 2
				12..15: j:= 3
				16..17: j:= 4
				18..20: j:= 5
			FINSEGUN

			Cajero[i,j]:= Cajero[i,j] + remp.Facturado

			LEER(aemp, remp)

		FINMIENTRAS

		PARA i:= 1 HASTA 10 HACER //totales por caja
			Cajeros[i,6]:= Cajeros[i,1]+Cajeros[i,2]+Cajeros[i,3]+Cajeros[i,4]+Cajeros[i,5]
		FINPARA

		PARA j:= 1 HASTA 5 HACER //totales por caja
			Cajeros[11,j]:= Cajeros[i,j]+Cajeros[i,j]+Cajeros[i,j]+Cajeros[i,j]+Cajeros[i,j]
		FINPARA

		PARA j:= 1 HASTA 5 HACER
			SEGUN j HACER
				1: HORARIOS:= "8-10"
				2: HORARIOS:= "10-12"
				3: HORARIOS:= "12-16"
				4: HORARIOS:= "16-18"
				5: HORARIOS:= "18-20"
			FS
			ESC("Para el horario:", HORARIOS, "el total recaudado fue:", Cajeros [11,j])
		FINPARA

		ESC("El total general es de: $", Cajeros[11,6])













ACCION Arreglo_3.24 ES

    AMBIENTE

    EMPLEADOS = REGISTRO
        NOMYAPE: AN(20)
        N_CAJA: N(5)
        IMPORTE: N(7,2)
        HORARIO: N(2)
    FR

    ARCH: Archivo de EMPLEADOS ordenado por NOMYAPE
    REG: EMPLEADOS
    i,j: Enteros
    CAJEROS: arreglo [1..11][1..6] de Reales
    HORARIOS: AN(15)

    PROCESO
    ABRIR E/(ARCH)
    LEER(ARCH,REG)
    PARA i:= 1 HASTA 11 HACER
        PARA j:= 1 HASTA 6 HACER
            CAJEROS[i,j]:= 0
        FP
    FP
    MIENTRAS NFDA(ARCH) HACER
        i:= REG.N_CAJA
        SEGUN .HREGORARIO HACER
            8..9:   j:= 1
            10..11: j:= 2
            12..15: j:= 3
            16..17: j:= 4
            18..20: j:= 5
        FS           
        CAJEROS[i,j]:=  CAJEROS[i,j] + REG.IMPORTE
        CAJEROS[i,6]:=  CAJEROS[i,6] + REG.IMPORTE
        CAJEROS[11,J]:= CAJEROS[11,j] + REG.IMPORTE
        CAJEROS[11,6]:= CAJEROS[11,6] + REG.IMPORTE
        LEER(ARCH,REG)
    FM
    CERRAR(ARCH)
    PARA i:= 1 hasta 10 HACER
        ESC("para la caja",i,"importe total:", CAJEROS[i,6])
    FP
    PARA J:= 1 HASTA 5 HACER
        SEGUN J HACER
            1: HORARIOS:= "8-10"
            2: HORARIOS:= "10-12"
            3: HORARIOS:= "12-16"
            4: HORARIOS:= "16-18"
            5: HORARIOS:= "18-20"
        FS
        ESC("Para el horario:", HORARIOS, "el total recaudado fue:", CAJEROS[11,J])
    FP
    ESC("El total general fue de:", CAJEROS[11,6])
    
Fin_accion.



ACCION Ej_3_24 ES

	AMBIENTE

		Formato_Empleado= REGISTRO
			NomAp: AN(25)
			Facturado: N(7)
			NumCaja: 1..10
			Horario: 8..20
		FINREGISTRO

		Reg_Empleado:= Formato_Empleado
		Arch_Empleado:= Archivo de Formato_Empleado Ordenado por NomAp

		A: Arreglo [1..11,1..6] de reales

		i, j: entero
		totalxcaja, totalxhora: real
	
	PROCESO

		ABRIR E/(Arch_Empleado)
		LEER(Arch_Empleado, Reg_Empleado)

		MIENTRAS NFDA(Arch_Empleado) HACER

			PARA i:= 1 HASTA 11 HACER

				totalxcaja:= 0

				SI i < 11 ENTONCES

					PARA j:= 1 HASTA 6 HACER

						SI j < 6 ENTONCES
							A[i,j]:= Reg_Empleado.Facturado
							totalxcaja:= totalxcaja + Reg_Empleado.Facturado
						SINO
						A[i,j]:= totalxcaja

				SINO



				FINSI

				FINPARA

			FINPARA

		FINMIENTRAS

