Ejercicio 3.10¶
//Dado un arreglo de 50 elementos, cada uno de los cuales tiene los siguientes datos: Código de localidad y
//Lluvia caída en un año. Escribir un algoritmo que permita saber dada una localidad, cuanto llovió en el año.
//Aplicar el método más adecuado considerando que el arreglo esta ordenado por Código de localidad.

ACCION Ejercicio_3_10 (A: Arreglo [1..50] de Datos) ES

    AMBIENTE

        datos= REGISTRO
            Cod_Loc: N(5)
            Lluvia_Caida: N(5)
        FINREGISTRO

        i, localidad: entero
        band: logico
        
    
    PROCESO

        band:= f
        ESC("Escriba codigo de localidad a analizar")
        LEER(localidad)
        

        Para i:= 1 HASTA 50 HACER
            SI A[i].Cod_Loc = localidad ENTONCES
                ESC("La cantidad de lluvia caida en el año para la localidad: ", localidad, " es de: ", A[i].Lluvia_Caida)
                band:= v
            FINSI
        FINPARA

        SI band = f ENTONCES
            ESC("Código de localidad inválido")
        FINSI

FINACCION



















ACCION Ejercicio_3_10 (A: Arreglo (1..50) de DATOS) ES

    AMBIENTE

        DATOS= REGISTRO
            Cod_Loc: N(5)
            Lluvia_Caida: N(5)
        FINREGISTRO

        localidad: N(5)
        i: Entero
        band: logico
    
    PROCESO

        ESC("Escriba la localidad a analizar")
        LEER(localidad)
        band: false

        PARA i:= 1 HASTA 50 HACER
            SI localidad = A[i].Cod_Loc ENTONCES
                ESC("Para la localidad ", localidad, ", la cantidad de lluvia que cayó es de: ", A[i].Lluvia_Caida)
                band: true // No es necesario, pero esto es para saber que se encontró y hacer lo que hago a continuación abajo
            FINSI
        FINPARA

            SI band = false ENTONCES
              ESC("La localidad ingresada no se encuentra entre las localidades informadas")
            FINSI

FINACCION

































ACCION Arreglo_3.10 (A: arreglo [1..50] de Datos) Es
    AMBIENTE

    Datos = REGISTRO
        cod_localidad: n(5)
        Lluvia: n(5)
    FR

    localidad: n(5)
    i: Entero

    PROCESO

    ESC("ingrese el codigo de localidad que desea buscar:")
    LEER(localidad)
    PARA i:= 1 hasta 50 HACER
        SI localidad = A[i].cod_localidad ENTONCES
            ESC("Para la localidad",localidad,"lluvia anual:",A[i].Lluvia)
        SINO
            ESC("no exite la localidad ingresada")
        FS
    FP
F_ACCION.