Ejercicio 2.2.12¶
Dada una secuencia con información de población de un PAIS:

POBLACION Ordenado por Provincia, Departamento, Ciudad, Barrio, ID_Casa

Provincia/Departamento/Ciudad/Barrio/ID_Casa/Cantidad_Habitantes
Generar una secuencia con información de los departamentos de esa provincia, cuyo registro tenga el siguiente formato:

POBLACION_SALIDA

Provincia/Departamento/Cantidad_Habitantes


ACCION Ejercicio2_2_12 ES
    AMBIENTE
        Poblacion= REGISTRO
            Provincia: AN(15)
            Departamento: AN(15)
            Ciudad: AN(15)
            Barrio: AN(30)
            ID_Casa: N(4)
            Cantidad_Hab: N(5)
        FinRegistro
        Arch_Pob:= Archivo de Poblacion Ordenado Por Provincia, Departamento, CIudad, Barrio, ID_Casa
        Reg_Pob:= Poblacion

        Salida= REGISTRO
            Provincia: AN(15)
            Departamento: AN(15)
            Cantidad_Hab: N(5)
        FinRegistro
        Arch_Sal:= Archivo de Salida
        Reg_Sal:= Salida

        Resg_Prov: AN(15)
        Resg_Dept: AN(15)
        Total_H: N(5)

        PROCEDIMIENTO Corte_Prov ES
            Reg_Sal.Provincia:= Resg_Prov
            Reg_Sal.Departamento:= Resg_Dept
            Reg_Sal.Cantidad_Hab:= Total_Hab
            Escribir(Arch_Sal, Reg_Sal)
            Total_Hab:= 0
            Reg_Dept:= Reg_Pob.Departamento
        FINPROCEDIMIENTO

        PROCEDIMIENTO Corte_Departamento ES
            Corte_Prov
            Resg_Provincia:= Reg_Pob.Provincia
        FINPROCEDIMIENTO

    PROCESO
        ABRIR E/(Arch_Pob)
        ABRIR S/(Arch_Sal)
        LEER(Arch_Pob, Reg_Pob)
        Resg_Prov:= Reg_Pob.Provincia
        Resg_Dept:= Reg_Pob.Departamento
        Total_H:=0
        MIENTRAS NoFDA HACER
            SI Resg_Prov <> Reg_Pob.Provincia ENTONCES
                Corte_Provincia
            SINO
                SI Resg_Dept <> Reg_Pob.Departamento ENTONCES
                Corte_Departamento
                FINSI
            FINSI
            Total_H:= Total_H + Cantidad_Hab
            LEER(Arch_Pob, Reg_Pob)
        FINMIENTRAS
        Corte_Provincia
        CERRAR(Arch_Pob)
        CERRAR(Arch_Sal)
FINACCION