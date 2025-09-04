//ACTUALIZACION CON CORTE DE CONTROL

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


ACCION EMPLEADOS ES

	AMBIENTE

		//primer archivo
		formato_empleado = registro
			clave= registro
				nro_sucursal: n(5)
				categoria: ("A", "B", "C")
			fr
			nombre_empleado: an(25)
			cod_curso: n(5)
			tecnico: ("si", "no")
		fr

		//resguardos, registro y archivo
		resgnro_sucursal:n(5)
		resgcategoria: ("A", "B", "C")
		regemp: formato_empleado
		archemp: archivo de regemp ordenado por clave

		//segundo archivo
		formato_curso = registro
			cod_curso: n(5)
			nombre: an(25)
			fecha= registro
				dia: 1..31
				mes: 1..12
				año: n(4)
			fr
		fr

		//registro y archivo
		regcur: formato_curso
		archcur: Archivo de regcur Indexado por cod_curso

		//contadores de total de empleados tecnicos/no_tecnicos y en total sumados
		//por sucursal:
		stot_tecnicos, stot_notecnicos, stot_empleados: entero
		//por categoria:
		ctot_tecnicos, ctot_notecnicos, ctot_empleados: entero
		//total general:
		tot_tecnicos, tot_notecnicos, tot_empleados: entero

		//cortes de control
		procedimiento corte_categoria es
			//paso 1: corte anterior
				//ninguno
			//paso 2: informar
			esc("Para la categoria: ", resgcategoria)
			esc("la cantidad de empleados tecnicos son: ", ctot_tecnicos)
			esc("la cantidad de empleados no tecnicos son: ", ctot_notecnicos)
			esc("la cantidad de empleados en total son: ", ctot_empleados)
			//paso 3: sumo a totales más grandes
			stot_tecnicos:= stot_tecnicos + ctot_tecnicos
			stot_notecnicos:= stot_notecnicos + ctot_notecnicos
			stot_empleados:= stot_empleados +  ctot_empleados
			//paso 4: actualizar resguardos
			resgcategoria:= regemp.categoria
			//paso 5: contadores en 0
			ctot_tecnicos:= 0
			ctot_notecnicos:= 0
			ctot_empleados:= 0
		fp

		procedimiento corte_sucursal es
			//paso 1: corte anterior
			corte_categoria
			//paso 2: informar
			esc("Para la sucursal nro: ", resg_nro_sucursal)
			esc("la cantidad de empleados tecnicos son: ", stot_tecnicos)
			esc("la cantidad de empleados no tecnicos son: ", stot_notecnicos)
			esc("la cantidad de empleados en total son: ", stot_empleados)
			//paso 3: sumo a totales más grandes
			tot_tecnicos:= tot_tecnicos + stot_tecnicos
			tot_notecnicos:= tot_notecnicos + stot_notecnicos
			tot_empleados:= tot_empleados +  stot_empleados
			//paso 4: actualizar resguardos
			resgnro_sucursal:= regemp.nro_sucursal
			//paso 5: contadores en 0
			stot_tecnicos:= 0
			stot_notecnicos:= 0
			stot_empleados:= 0
		fp


	PROCESO

		//antes del bucle:
		//1- abrir archivos y leer si es secuencial
		ABRIR E/(archemp); LEER(archemp, regemp)//secuencial
		ABRIR E/(archcur);//indexado por cod_curso

		//2- contadores en 0
		stot_tecnicos:= 0; stot_notecnicos:= 0; stot_empleados:= 0
		ctot_tecnicos:= 0; ctot_notecnicos:= 0; ctot_empleados:= 0
		tot_tecnicos:= 0; tot_notecnicos:= 0; tot_empleados:= 0

		//3-reguardos
		resgnro_sucursal:= regemp.nro_sucursal
		resgcategoria:= regemp.categoria

		//4- bucle para cortes y manejar archivos 
		MIENTRAS NFDA(archemp) HACER

			//punto 2 cortes y contadores:
			//trato cortes
			si resgnro_sucursal <> regemp.nro_sucursal entonces
				corte_sucursal
			sino
				si resgcategoria <> regemp.categoria entonces
					corte_categoria
				fs
			fs

			//trato contadores
			ctot_empleados:= ctot_empleados + 1 //sin importar si son tecnico o no, es un empleado mas
			segun regemp.tecnico hacer
				"si": ctot_tecnicos:= ctot_tecnicos + 1 // si el empleado es tecnico
				"no": ctot_notecnicos:= ctot_notecnicos + 1 // si el empleado no es tecnico
			fs

			//punto 1
			//Para cada empleado: sucursal, categoria, nombre del empleado y nombre del curso que debe realizar.
			esc("Nombre del empleado: ", regemp.nombre_empleado)
			esc("Sucursal nro: ", regemp.nro_sucursal)
			esc("Categoria: ", regemp.categoria)
			
			//TRATAR INDEXADO/PROCESO SI EXISTE
			regcur.cod_curso:= regemp.cod_curso
			LEER(archcur, regcur.cod_cur)
			SI EXISTE ENTONCES
			ESC("Nombre del curso a realizar: ", regcur.nombre)
			FS

			//TRATAR SECUENCIAL
			LEER(archemp, regemp)
		
		FINMIENTRAS

		//5- Llamar manualmente al corte más grande porque no se lo llega a llamar automáticamente
		corte_sucursal

		//6- Corte total general
		esc("Total general de empleados en toda la empresa:")
		esc("la cantidad de empleados tecnicos son: ", tot_tecnicos)
		esc("la cantidad de empleados no tecnicos son: ", tot_notecnicos)
		esc("la cantidad de empleados en total son: ", tot_empleados)

		//7- Cierro los archivos
		CERRAR(archemp); CERRAR(archcur)

FINACCION














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