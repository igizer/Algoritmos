
Se dispone de un archivo con datos de los alumnos de la U.T.N. con la siguiente información para cada alumno:

PRODUCTOS Ordenado por Sexo, Carrera, Nro_Legajo

| Sexo | Carrera | Nro_Legajo |/ Fecha_Ingreso / Total_Materias_Aprobadas /
Se desea un listado con el siguiente detalle por sexo, carrera y total general

Total de alumnos que ingresaron en el año 2009 con más de 20 materias aprobadas.

Total de alumnos que ingresaron en el año 2009 con menos de 20 materias aprobadas


ACCION Corte_2_2_11 ES
    AMBIENTE
        PRODUCTO = REGISTRO
            Sexo: ("F", "M")
            Carrera: ("ISI", "IQ", "IEM")
            Nro_Legajo: N(5)
            Fecha_Ingreso = REGISTRO
                Dia: 1..31
                Mes: 1..12
                Año: N(4)
            FINREGISTRO
            Total_Materias_Aprobadas: N(2)
        FINREGISTRO
        Reg_Prod: PRODUCTO
        Arch_Prod: Archivo de PRODUCTO Ordenado por Sexo, Carrera, Nro_Legajo

        Resg_Sexo: ("F", "M")
        Resg_Carrera: ("ISI", "IQ", "IEM")
        TOTC_MAS, TOTC_MEN, TOTS_MAS, TOTS_MEN, TOTG_MAS, TOTG_MEN: N(5)

        PROCEDIMIENTO CORTE_SEXO ES
            ESC("Para el sexo", Resg_Sexo)
            ESC("Total de alumnos que ingresaron en 2009 con más de 20 materias aprobadas: ", TOTS_MAS)
            ESC("Total de alumnos que ingresaron en 2009 con menos de 20 materias aprobadas: ", TOTS_MEN)
            TOTC_MAS:= TOTC_MAS + TOTS_MAS
            TOTC_MEN:= TOTC_MEN + TOTS_MEN
            TOTS_MAS:= 0
            TOTS_MEN:= 0
        FINPROCEDIMIENTO

        PROCEDIMIENTO CORTE_CARRERA ES
            ESC("Para la carrera", Resg_Carrera)
            ESC("Total de alumnos que ingresaron en 2009 con más de 20 materias aprobadas: ", TOTC_MAS)
            ESC("Total de alumnos que ingresaron en 2009 con menos de 20 materias aprobadas: ", TOTC_MEN)
            TOTG_MAS:= TOTG_MAS + TOTC_MAS
            TOTG_MEN:= TOTG_MEN + TOTC_MEN
            TOTC_MAS:= 0
            TOTC_MEN:= 0
        FINPROCEDIMIENTO

        PROCESO
            ABRIR E/(Arch_Prod)
            LEER(Arch_Prod, Reg_Prod)
            Resg_Sexo:= Reg_Prod.Sexo
            Resg_Carrera:= Reg_Prod.Carrera
            TOTS_MAS:= 0
            TOTS_MEN:= 0
            TOTC_MAS:= 0
            TOTC_MEN:= 0
            TOTG_MAS:= 0
            TOTG_MEN:= 0
            MIENTRAS NFDA(Arch_Prod) HACER
            SI Resg_Sexo<>Reg_Prod.Sexo ENTONCES
                CORTE_SEXO
            SINO
                SI Resg_Carreta<>Reg_Prod.Carrera ENTONCES
                CORTE_CARRERA
            FINSI
            SI Reg_Prod.Fecha_Ingreso <> 2009 ENTONCES
                SI Reg_Prod.Total_Materias_Aprobadas > 20 ENTONCES
                    TOTS_MAS:= TOTS_MAS + 1
                FINSI
                SI Reg_Pro.Total_Materias_Aprobadas > 20 ENTONCES // En un examen me van a decir que debía de haber usado un sino, pero dice explícitamente en la consigna mayores de 20 y menores de 20, un sino haría que sea menor o IGUAL a 20, nunca me pide IGUAL a 20
                    TOTS_MEN:= TOTS_MEN + 1
                FINSI
            SI Resg_Carreta<>Reg_Prod.Carrera ENTONCES
                CORTE_CARRERA
            FINSI
            LEER(Arch_Prod, Reg_Prod)
            ESC("El total general de estudiantes con más de 20 materias aprobadas son", TOTG_MAS)
            ESC("El total general de estudiantes con menos de 20 materias aprobadas son", TOTG_MEN)

















ACCION Corte_2.2.11  ES
    AMBIENTE 

    FACULTAD = REGISTRO
        SEXO: ("F","M");
        CARRERA : ("ISI","IEM","IQ");
        NRO_LEGAJO:N(5);
        FECHA_INGRESO = REGISTRO
            AA:N(4);
            MM: 1..12;
            DD: 1..31;
        FR;
        TOTAL_MATERIAS_APROBADAS: N(2);
    FR;
    ARCHIVO1 : Archivo de Facultad ordenado por SEXO,CARRERA,NRO_LEGAJO;
    REGISTRO1 : FACULTAD;

    RES_SEXO: ("F","M");
    RES_CARRERA:("ISI","IEM","IQ");
    TOTS_MAS, TOTS_MEN, TOTC_MAS, TOTC_MEN,TOTG_MAS,TOTG_MEN: N(5)

    SUBACCION CORTE_CARRERA ES  
        ESC("Para la carrera :",RES_CARRERA);
        ESC("El total de alumnos que ingresaron en el 2009 con mas de 20 materias aprobadas es ", TOTC_MAS);
        ESC("El total de alumnos que ingresaron en el 2009 con menos de 20 materias aprobadas es:",TOTC_MEN);
        TOTS_MAS := TOTS_MAS + TOTC_MAS;
        TOTS_MEN := TOTS_MEN + TOTC_MEN;
        RES_CARRERA: REGISTRO1.CARRERA;
        TOTC_MAS:= 0;
        TOTC_MEN:= 0;
    FSUB;

    SUBACCION CORTE_SEXO ES
        CORTE_CARRERA; 
        ESC ("Para el sexo:",RES_SEXO);
        ESC("El total de alumnos que ingresaron en el 2009 con mas de 20 materias aprobadas es ", TOTS_MAS);
        ESC("El total de alumnos que ingresaron en el 2009 con menos de 20 materias aprobadas es:",TOTS_MEN);
        TOTG_MAS := TOTG_MAS + TOTS_MAS;
        TOTG_MEN := TOTG_MEN + TOTS_MEN;
        RES_SEXO: REGISTRO1.SEXO;
        TOTS_MAS:= 0;
        TOTS_MEN:= 0;
    FR;

    PROCESO
    ABRIR E/(ARCHIVO1); LEER(ARCHIVO1,REGISTRO1);
    RES_CARRERA:= REGISTRO1.CARRERA;
    RES_SEXO:= REGISTRO1.SEXO;
    TOTC_MAS:= 0; TOTC_MEN:=0; TOTS_GEN:=0; TOTS_MAS:=0; TOTS_MEN:=0; TOTG_MAS:=0; TOTG_MEN:=0;
    MIENTRAS NFDA(ARCHIVO1) HACER
        SI REGISTRO1.SEXO <> RES_SEXO ENTONCES
            CORTE_SEXO
        SINO 
            SI REGISTRO1.CARRERA <> RES_CARRERA ENTONCES
                CORTE_CARRERA;
            FS;
        FS;
        SI REGISTRO1.FECHA_INGRESO.AA = 2009 ENTONCES   
            SI REGISTRO1.TOTAL_MATERIAS_APROBADAS > 20 ENTONCES
                TOTC_MAS:= TOTC_MAS + 1;
            SINO
                TOTC_MEN:= TOTC_MEN + 1;
            FS;
        FS;
        LEER(ARCHIVO1,REGISTRO1);
    FM;
    CORTE_SEXO;
    ESC("El total general de alumno que ingresaron en el 2009 con mas de 20 materias aprobadas es:", TOTG_MAS);
    ESC("El total general de alumno que ingresaron en el 2009 con menos de 20 materias aprobadas es:", TOTG_MEN);
    CERRAR(ARCHIVO1);
FIN_ACCION.