// Consigna 2 ---> tiempo total: 00:39:17
// Agregrar Consigna 1 ---> tiempo total: 00:46:44


Ejercicio 2.3.2¶
Una empresa dispone de un fichero secuencial con datos de sus empleados, ordenado por número de sucursal, y categoría, con los siguientes datos:

EMPLEADOS Ordenado por Nro_Sucursal, Categoria

||Nro_Sucursal||Categoria (A,B,C)|| Nombre_Empleado | Cod_Curso | Tecnico (si,no) |
Y un fichero con datos de cursos, indexado por código de curso:

CURSO Indexado por Cod_Curso

||Cod_Curso|| Nombre | Fecha | Cant_Horas |
Emitir un listado informando:

1. Para cada empleado: sucursal, categoría, nombre del empleado y nombre del curso que debe realizar.
2. Por sucursal, categoría y toda la empresa:
	a. Total empleados técnicos
	b. Total empleados no técnicos
	c. Total empleados


ACCION Ejercicio_2_3_2 ES

	AMBIENTE

		Formato_Empleado = REGISTRO
			Clave= REGISTRO
				Nro_Sucursal: entero(5)
				Categoria: ("A", "B", "C")
			FINREGISTRO
			Nombre_Empleado: AN(25)
			Cod_Curso: entero(5)
			Tecnico: ("si", "no")
		FINREGISTRO

		regempleado: Formato_Empleado
		archempleado: Archivo de Empleado Ordenado por Clave

		Formato_Curso = REGISTRO
			Cod_Curso: entero(5)
			Nombre: AN(15)
			Fecha= REGISTRO
				dia: (1..31)
				mes: (1..12)
				año: entero(4)
			FINREGISTRO
			Cant_Horas: entero(0..24)
		FINREGISTRO

		regcurso: Formato_Curso
		archcurso: Archivo de Formato_Curso Indexado por Cod_Curso

		//totales categoria
		ctottecnico, ctotnotecnico, ctotempleados: entero 

		//totales sucursal
		stottecnico, stotnotecnico, stotempleados: entero 
		
		//totales empresa
		etottecnico, etotnotecnico, etotempleados: entero

		//resguardos
		rescategoria, ressucursal: caracter


		//procedimientos de corte
		PROCEDIMIENTO Corte_Categoria ES
			//Informo
			ESC("Para la categoría: ", rescategoria)
			ESC("Total empleados técnicos: ", ctottecnico)
			ESC("Total empleados no técnicos: ", ctotnotecnico)
			ESC("Total empleados: ", ctotempleados)

			//sumo a totales más grandes
			stottecnico:= stottecnico + ctottecnico
			stotnotecnico:= stotnotecnico + ctotnotecnico
			stotempleados:= stotempleados + ctotempleados

			//reinicio contadores totales
			ctottecnico:= 0
			ctotnotecnico:= 0
			ctotempleados:= 0

			//actualizo el resguardo
			rescategoria:= regempleado.Clave.Categoria
		FINPROCEDIMIENTO

		PROCEDIMIENTO Corte_Sucursal ES
			//Informo
			ESC("Para la sucursal: ", ressucursal)
			ESC("Total empleados técnicos: ", stottecnico)
			ESC("Total empleados no técnicos: ", stotnotecnico)
			ESC("Total empleados: ", stotempleados)

			//sumo a totales más grandes
			etottecnico:= etottecnico + stottecnico
			etotnotecnico:= etotnotecnico + stotnotecnico
			etotempleados:= etotempleados + stotempleados

			//reinicio contadores totales
			stottecnico:= 0
			stotnotecnico:= 0
			stotempleados:= 0

			//actualizo el resguardo
			ressucursal:= regempleado.Clave.Nro_Sucursal
		FINPROCEDIMIENTO



	PROCESO

		ABRIR E/(archempleado); ABRIR E/(archcurso)
		LEER(archempleado, regempleado)

		//totales
		ctottecnico:= 0; ctotnotecnico:= 0; ctotempleados:= 0
		stottecnico:= 0; stotnotecnico:= 0; stotempleados:= 0
		etottecnico:= 0; etotnotecnico:= 0; etotempleados:= 0

		//resguardos
		rescategoria:= regempleado.Clave.Categoria
		ressucursal:= regempleado.Clave.Nro_Sucursal



		MIENTRAS NFDA(archempleado) HACER

			//Cortes
			SI ressucursal <> regempleado.Clave.Nro_Sucursal ENTONCES

				Corte_Sucursal()
			
			SINO

				SI rescategoria <> regempleado.Clave.Categoria ENTONCES

					Corte_Categoria()

				FINSI

			FINSI


			//Sumo empleados a las categorias
			ctotempleados:= ctotempleados + 1

			SEGUN regempleado.Tecnico Hacer
				"si": ctottecnico:= ctottecnico + 1
				"no": ctotnotecnico:= ctotnotecnico + 1
			FINSEGUN
		

		//Consigna 1
			ESC("Empleado:", regempleado.Nombre_Empleado)
			ESC("Sucursal nro: ", regempleado.Clave.Nro_Sucursal)
			ESC("Categoria: ", regempleado.Clave.Categoria)

			//Proceso SI EXISTE el indexado
			regcurso.Cod_Curso:= archcurso.Cod_Curso
			LEER(archcurso, regcurso)
			SI EXISTE ENTONCES
			ESC("Nombre Curso a Realizar: ", regcurso.Nombre)
			FINSI

			//Avanzo a los siguientes registros para continuar informando datos y cortes
			LEER(archempleado, regempleado)

		FINMIENTRAS

		//llamo al corte más grnade para procesar los últimos datos
		Corte_Sucursal()

		//corte general de la empresa
		ESC("Para la empresa ")
		ESC("Total general empleados técnicos: ", etottecnico)
		ESC("Total general empleados no técnicos: ", etotnotecnico)
		ESC("Total general empleados: ", etotempleados)
		
		//cierro los archivos
		CERRAR(archempleado); CERRAR(archcurso)









//Ahora como hizo el de la guía resuelta


ACCION actualizaCorte ES
    AMBIENTE

    EMPLEADOS = REGISTRO
        NRO_SUCURSAL: N(5)
        CATEGORIA: AN(1)
        NOMBRE_EMPLEADO: AN(30)
        COD_CURSO: N(5)
        TECNICO: AN(2)
    FR;

    AEMPLEADOS: Archivo ordenado por NRO_SUCURSAL,CATEGORIA;
    REMPLEADOS: EMPLEADOS;

    CURSO = REGISTRO    
        COD_CURSO: N(5)
        NOMBRE: AN(20)
        FECHA = REGISTRO
            AA:N(4)
            MM:1..12
            DD:1..31
        FR;
        CANT_HORAS: N
    FR;

    ACURSO: Archivo de CURSO indexado por COD_CURSO;
    RCURSO: CURSO;

    RES_SUCU: N(5)
    RES_CATE: AN(1)

    TOTSITEC_CAT,TOTNOTEC_CAT,TOTEMP_CAT: N
    TOTSITEC_SUC,TOTNOTEC_SUC,TOTEMP_SUC: N 
    TOTSITEC_GEN,TOTNOTEC_GEN,TOTEMP_GEN: N

    SUBACCION CORTE_CATEGORIA ES
        ESC("Para la categoria:",RES_CATE)
        ESC("Total de empleados tecnicos", TOTSITEC_CAT)
        ESC("Total de empleados  no tecnicos", TOTNOTEC_CAT)
        ESC("Total empleados:", TOTEMP_CAT)
        TOTSITEC_SUC:= TOTSICTEC_SUC + TOTSITEC_CAT;
        TOTNOTEC_SUC:= TOTNOTEC_SUV + TOTNOTEC_CAT;
        TOTEMP_SUC:= TOTEMP_SUC + TOTEMP_CAT;
        TOTSITEC_CAT:=0;
        TOTNOTEC_CAT:=0;
        TOTEMP_CAT:=0;
        RES_CATE:= REMPLEADOS.CATEGORIA;
    FSUB     

    SUBACCION CORTE_SUCURSAL ES
        CORTE_CATEGORIA;
        ESC("Para la sucursal:",RES_SUCU)
        ESC("Total de empleados tecnicos", TOTSITEC_SUC)
        ESC("Total de empleados  no tecnicos", TOTNOTEC_SUC)
        ESC("Total empleados:", TOTEMP_SUC)
        TOTSITEC_GEN:= TOTSICTEC_GEN + TOTSITEC_SUCT;
        TOTNOTEC_GEN:= TOTNOTEC_GEN + TOTNOTEC_SUCT;
        TOTEMP_GEN:= TOTEMP_GEN + TOTEMP_SUC;
        TOTSITEC_SUC:=0;
        TOTNOTEC_SUC:=0;
        TOTEMP_SUC:=0;
        RES_SUCU:= REMPLEADOS.NRO_SUCURSAL;
    FSUB

    PROCESO
    ABRIR E/(AEMPLEADOS) ABRIR E/(ACURSO)
    LEER(AEMPLEADOS,REMPLEADOS)
    TOTSITEC_CAT:=0;TOTNOTEC_CAT:=0;TOTEMP_CAT:=0;
    TOTSITEC_SUC:=0;TOTNOTEC_SUC:=0;TOTEMP_SUC:=0;
    TOTSITEC_GEN:=0;TOTNOTEC_GEN:=0;TOTEMP_GEN:=0;
    RES_CATE:= REMPLEADOS.CATEGORIA;
    RES_SUCU:= REMPLEADOS.NRO_SUCURSAL;
    MIENTRAS NFDA(AEMPLEADOS) HACER
        SI RES_CSUCU <> REMPLEADOS.NRO_SUCURSAL ENTONCES
            CORTE_SUCURSAL;
        SINO 
            SI RES_CATE <> REMPLEADOS.CATEGORIA ENTONCES
                CORTE_CATEGORIA;
            FS;
        FS;
        RCURSO.COD_CURSO:= REMPLEADOS.COD_CURSO;
        LEER(ACURSO,RCURSO)
        SI EXISTE ENTONCES
            ESC(REMPLEADOS.NRO_SUCURSAL, REMPLEADOS.CATEGORIA, REMPLEADO.NOMBRE_EMPLEADO, RCURSO.NOMBRE);
        SINO
            ESC(REMPLEADOS.NRO_SUCURSAL, REMPLEADOS.CATEGORIA, REMPLEADO.NOMBRE_EMPLEADO, "nombre no disponible");
        FS;
        TOTEMP_CAT:= TOEMP_CAT + 1;
        SI REMPLEADOS.TECNICO = "SI" ENTONCES
            TOTSITEC_CAT:= TOTSITEC_CAT +1
        SINO
            TOTNOTEC_CAT:= TOTNOTEC_CAT + 1;
        FS;
        LEER(AEMPLEADOS,REMPLEADOS)
    FM;
    CORTE_SUCURSAL;
    ESC("Total de empleados tecnicos de la empresa", TOTSITEC_GEN)
    ESC("Total de empleados  no tecnicos de la empresa", TOTNOTEC_GEN)
    ESC("Total empleados de la empresa:", TOTEMP_GEN)
    CERRAR(AEMPLEADOS); CERRAR(ACURSO);
Fin_accion.