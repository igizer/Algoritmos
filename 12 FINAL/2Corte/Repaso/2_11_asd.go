Se dispone de un archivo con datos de los alumnos de la U.T.N. con la siguiente información para cada alumno:

PRODUCTOS Ordenado por Sexo, Carrera, Nro_Legajo

| Sexo | Carrera | Nro_Legajo |/ Fecha_Ingreso / Total_Materias_Aprobadas /
Se desea un listado con el siguiente detalle por sexo, carrera y total general

Total de alumnos que ingresaron en el año 2009 con más de 20 materias aprobadas.

Total de alumnos que ingresaron en el año 2009 con menos de 20 materias aprobadas


ACCION CORTE_2_2_11 ES
AMBIENTE

	cont_sex_mas, cont_carrera_mas, cont_tot_mas: entero // contadores por sexo, carrera y total con MÁS de 20 aprobados
	cont_sex_menos, cont_carrera_menos, cont_tot_menos: entero // contadores por sexo, carrera y total con MENOS de 20 aprobados

	PRODUCTOS = registro
		sexo: ("Femenino", "Masculino")
		carrera: ("isi", "iq", "iem")
		nro_legajo: n(5)
		fecha_ingreso = registro
			dia: 1..31
			mes: 1..12
			año: n(4)
		fr
		total_materias_aprobadas: n(2)
	fr

	resg_sexo: ("Femenino", "Masculino"); resg_carrera: ("isi", "iq", "iem"); // resguardos de sexo y carrera
	reg: PRODUCTOS // registro de productos
	arch: Archivo de PRODUCTOS Ordenado por sexo, carrera, nro_legajo // archivo de productos


	procedimiento corte_carrera es
		//paso cero, llamar primero a un corte mas chico englobado en este, si hay

		//primero, informar corte
		esc("la cantidad de personas en la carrera  ", reg.carrera, "que ingresaron en 2009 CON MAS DE 20 MATERIAS APROBADAS son: ", cont_sexo_mas)
		esc("y con MENOS de 20 materias son: ", cont_sexo_menos)

		//segundo, sumar los contadores al contador siguiente más grande que él
		cont_sexo_mas:= cont_carrera_mas + cont_sexo_mas
		cont_sexo_menos:= cont_carrera_menos + cont_sexo_menos

		//tercero, actualizar los resguardos del registro al nuevo registro
        resg_carrera:= reg.carrera

		//cuarto, poner los contadores de este corte en 0 para seguir contando por donde quedó
		cont_carrera_mas:= 0
		cont_carrera_menos:= 0
	fp


	procedimiento corte_sexo es
		//paso cero, llamar primero a un corte mas chico englobado en este, si hay
		corte_carrera

		//primero, informar corte
		esc("la cantidad de personas de sexo ", reg.sexo, "que ingresaron en 2009 CON MAS DE 20 MATERIAS APROBADAS son: ", cont_sexo_mas)
		esc("y con MENOS de 20 materias son: ", cont_sexo_menos)

		//segundo, sumar los contadores al contador siguiente más grande que él
		cont_tot_mas:= cont_tot_mas + cont_sexo_mas
		cont_tot_menos:= cont_tot_menos + cont_sexo_menos

        //tercero, actualizar los resguardos del registro al nuevo registro
        resg_sexo:= reg.sexo

		//cuarto, poner los contadores de este corte en 0 para seguir contando por donde quedó
		cont_sexo_mas:= 0
		cont_sexo_menos:= 0
	fp

PROCESO

    ABRIR(arch); leer(arch, reg)
    resg_carrera:= reg.carrera
    resg_sexo:= reg.sexo
    cont_sex_mas:= 0; cont_carrera_mas:= 0; cont_tot_mas:= 0;
    cont_sex_menos:= 0; cont_carrera_menos:= 0; cont_tot_menos:= 0;

    mientras nfda(arch) hacer
        SI reg.sexo <> resg_sexo ENTONCES
            corte_sexo
        sino
            si reg.carrera <> resg_carrera entonces
                corte_carrera
            finsi
        finsi
        si reg.fecha_ingreso.año = 2009 entonces
            si reg.total_materias_aprobadas > 20 entonces
                cont_carrera_mas := cont_carrera_mas + 1
            sino
                cont_carrera_menos := cont_carrera_menos + 1
            fs
        fs
        leer(arch, reg)
    fm
    esc("EL TOTAL GENERAL DE PERSONAS QUE INGRESARON EN EL 2009 CON MAS DE 20 MATERIAS APROBADAS SON: ", cont_tot_mas)
    esc("EL TOTAL GENERAL DE PERSONAS QUE INGRESARON EN EL 2009 CON MENOS DE 20 MATERIAS APROBADAS SON: ", cont_tot_menos)
    CERRAR(arch)

finaccion


	


		











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