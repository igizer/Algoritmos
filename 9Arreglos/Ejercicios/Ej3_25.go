Ejercicio 3.25¶
Una Fábrica que posee 4 plantas de producción en nuestro país y compra materia prima a 3 proveedores distintos,
desea realizar un control de los montos totales correspondientes a compras realizadas en cada planta durante el
último año, discriminados por proveedor y por mes. Para ello dispone de un archivo con los datos de las facturas
correspondientes. (Aclaración: el archivo no está ordenado por ningún criterio)

COMPRAS

Nro_Factura | Proveedor (A,B,C) | Fecha dd/mm/aaaa | Nro_Planta 1..4 | Importe



ACCION Ej_3_25 ES

	AMBIENTE

		A: Arreglo [1..4,1..13,1..4] de reales // puse 13 para colocar el monto total del año

		Formato_Fecha = REGISTRO
			dd: 1..31
			mm: 1..12
			aaaa: N(4)
		FINREGISTRO

		Formato_Compra = REGISTRO
			Nro_Factura: N(7)
			Proveedor: ("A","B","C")
			Fecha: Formato_Fecha
			Nro_Planta: 1..4
			Importe: N(9)
		FINREGISTRO

		reg:= Formato_Compra
		arch:= Archivo de Formato_Compra

		i, j, j: entero

	PROCESO

        ABRIR E/(arch)
        LEER(arch, reg)

        //pongo todo en 0
        PARA k:= 1 HASTA 4
            PARA j:= 1 HASTA 13 //Total por año es el 13
                PARA i:= 1 HASTA 4 //Fila 4 total por mes
                    A[i,j,k]:= 0
                FINPARA
            FINPARA
        FINPARA

        MIENTRAS NFDA(arch) HACER

            k:= reg.Nro_Planta
            j:= reg.Fecha.mm

            SEGUN reg.Proveedor HACER
                "A": i:= 1
                "B": i:= 2
                "C": i:= 3
            FINSEGUN

            A[i,j,k]:= A[i,j,k] + reg.Importe //importe del proveedor por un mes
            A[i,13,k]:= A[i,13,k] + reg.Importe //el importe anterior ahora sumado al total por año
            A[4,j,k]:= A[4,i,k] + reg.importe //lo anterior pero ahora sumado a su mes
            A[4,13,k]:= A[4,13,k] + reg.importe // total de los dos primeros

            LEER(arch, reg)
        
        FINMIENTRAS

        PARA k:= 1 HASTA 4 HACER

            PARA i:= 1 HASTA 3 HACER
                ESC("Para la planta",k ,"el total por el proveedor ",i, "es de", FABRICA[i,13,k])    //muestra los totales por proveedor de la planta
            FP

            PARA J:= 1 HASTA 12 HACER
                ESC("Para la planta",k ,"el total por el mes ",j, "es de", FABRICA[4,j,k])          //muestra los totales por mes de la planta
            FP

            ESC("El total por año de la planta",k,"es:" FABRICA[4,13,k])        //muestra el total final por año de la planta
            
        FINPARA

FINACCION





























ACCION Arreglo_3.25 ES
    AMBIENTE
    
    COMPRAS = REGISTRO
        NRO_FAC: N(5)
        PROVEEDOR: ("A","B","C")
        FECHA = REGISTRO
            AA: N(4)
            MM: 1..12
            DD: 1..31
        FR
        NRO_PLANTA: 1..4
        IMPORTE: N(7,2)
    FR

    REG: COMPRAS
    ARCH: Archivo de COMPRAS

    FABRICA: ARREGLO [1..4][1..13][1..4] de Reales
    i,j,k: enteros

    PROCESO
    
    ABRIR E/(ARCH)
    LEER(REG)
    //Pongo en cero la matriz
    PARA K:= 1 HASTA 4 HACER                
        PARA j:= 1 HASTA 13 HACER       //columna 13 total por año por provedores
            PARA I:= 1 HASTA 4 HACER    //fila 4 total por cada mes
                FABRICA[i,j,k]:= 0
            FP
        FP
    FP
    MIENTRAS NDFA(ARCH) HACER
        k:= REG.NRO_PLANTA
        j:= REG.FECHA.MM
  
        SEGUN REG.PROVEEDOR HACER
            "A": i:= 1
            "B": i:= 2
            "C": i:= 3
        FS
        FABRICA[i,j,k]:=  FABRICA[i,j,k] + REG.IMPORTE      //importe del provedor por un mes
        FABRICA[i,13,k]:= FABRICA[i,13,k] + REG.IMPORTE     // total por proveedor
        FABRICA[4,j,k]:= FABRICA[4,j,k] + REG.IMPORTE       // total por mes
        FABRICA[4,13,K]:= FABRICA[4,13,K] + REG.IMPORTE     // total del planta por año
        LEER(ARCH,REG)
    FM
    CERRAR(ARCH)
    PARA k:= 1 HASTA 4 HACER

        PARA I:= 1 HASTA 3 HACER
            ESC("Para la planta",k ,"el total por el proveedor ",i, "es de", FABRICA[i,13,k])    //muestra los totales por proveedor de la planta
        FP

        PARA J:= 1 HASTA 12 HACER
            ESC("Para la planta",k ,"el total por el mes ",j, "es de", FABRICA[4,j,k])          //muestra los totales por mes de la planta
        FP

        ESC("El total por año de la planta",k,"es:" FABRICA[4,13,k])        //muestra el total final por año de la planta     
    FP
fin_accion.