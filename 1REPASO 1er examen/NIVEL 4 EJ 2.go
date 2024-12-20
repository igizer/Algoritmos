Ejercicio 2
Se tiene información de todas las ediciones de la copa américa hasta la actual, con los respectivos estadios en un archivo secuencial con el siguiente formato:

Estadios (Ordenado por País edición, año edición, ciudad sede y Nombre estadio)
País Edición | Año Edición | Ciudad Sede | Nombre estadio | capacidad | cant partidos jugados

Se le solicita:
1) Total de partidos jugados por país y ciudad sede en estadios con capacidad superior a 40mil personas
2) Generar un archivo de salida con el siguiente formato
	País  Edición | Año edicion | Cant partidos jugados
3) Informar por estadio, el país, año de edición, ciudad, la cantidad de partidos jugados.

Por país	>	Por ciudad sede		>	Por estadios	>	Partidos jugados


ACCION

	AMBIENTE

		Estadios= Registro
			Clave= Registro
				Pais_Edicion: AN(15)
				Año_Edicion: N(4)
				Ciudad_Sede: AN(15)
				Nombre_Estadio: AN(15)
			FINREGISTRO
			Capacidad: N(5)
			Cant_Partidos_Jugados: N(3)
		FINREGISTRO

		Reg_Estadios: Estadios
		Arch_Estadios: Archivo de Estadios Ordenado por Clave

		Salida= Registro
			Pais_Edicion: AN(15)
			Año_Edicion: N(4)
			Cant_Partidos_Jugados: N(3)
		FINREGISTRO

		Reg_Salida: Salida
		Arch_Salida: Archivo de Salida // No se ordena por nada

		Resg_Pais_Edicion: AN(15)
		Resg_Año_Edicion: N(4)
		Resg_Ciudad: AN(15)

		ciudad_sup_40, pais_sup_40: entero //consigna 1

		partidos_ciudad, partidos_año, partidos_pais, //consigna 3
		


		PROCEDIMIENTO Corte_Ciudad ES
			ESC("La cantidad de partidos jugados en la ciudad ", Resg_Ciudad, ", son ", partidos_ciudad)
			ESC("Y la cantidad de partidos jugados en estadios con capacidad superior a 40mil personas son ", ciudad_sup_40)
			pais_sup_40:= pais_sup_40 + ciudad_sup_40
			ciudad_sup_40:= 0
			partidos_año:= partidos_año + partidos_ciudad
			partidos_ciudad:= 0
			Resg_Ciudad:= Reg_Estadios.Clave.Ciudad_Sede
		FINPROCEDIMIENTO

		PROCEDIMIENTO Corte_Año ES
			Corte_Ciudad()
			ESC("La cantidad de partidos jugados en la edición del año ", Resg_Año_Edición, " son ", partidos_año)
			Reg_Salida.Pais_Edicion:= Resg_Pais_Edicion
			Reg_Salida.Año_Edicion:= Resg_Año_Edicion
			Reg_Salida.Cant_Partidos_Jugados:= partidos_año
			GRABAR(Arch_Salida, Reg_Salida)
			partidos_pais:= partidos_pais + partidos_año
			partidos_año:= 0
			Resg_Año_Edicion:= Reg_Estadio.Clave.Año_Edicion
		FINPROCEDIMIENTO

		PROCEDIMIENTO Corte_Pais ES
			Corte_Año()
			ESC("La cantidad de partidos jugados en el país ", Resg.País_Edicion, " son ", partidos_pais)
			ESC("Y la cantidad d epartidos jugados en estadios con capacidad superior a 40mil personas son ", pais_sup_40)
			pais_sup_40:= 0
			partidos_pais:= 0
			Resg_Pais_Edicion:= Reg_Estadios.Clave.Pais_Edicion

		FINPROCEDIMIENTO
	
	PROCESO

		ABRIR E/(Arch_Estadios); LEER(Arch_Estadios, Reg_Estadios)
		ABRIR S/(Arch_Salida)

		partidos_ciudad:= 0
		partidos_año:= 0
		partidos_pais:= 0
		ciudad_sup_40:= 0
		pais_sup_40:= 0


		MIENTRAS NFDA(Arch_Estadios) HACER

			SI Resg_Pais_Edicion <> Reg_Estadios.Clave.Pais_Edicion ENTONCES
				Corte_Pais()
			SINO
				SI Resg_Año_Edicion <> Reg_Estadios.Clave.Año_Edicion
					Corte_Año()
				SINO
					SI Resg_Ciudad <> Reg_Estadios.Clave.Ciudad_Sede
						Corte_Ciudad()
					FINSI
				FINSI
			FINSI
			
            SI Reg_Estadios.Capacidad > 40000 Entonces
                ciudad_sup_40:= ciudad_sup_40 + 1
            FINSI
            partidos_ciudad:= partidos_ciudad + 1
			LEER(Arch_Estadios, Reg_Estadios)
            

		FINMIENTRAS

        Corte_Pais()

        

        











SEGUN CATEDRA

			
/* --- EJERCICIO 2 | CORTE DE CONTROL ---- */
ACCION corte_n4 ES
AMBIENTE
    //Definir archivo de entrada
    estadio = REGISTRO
        clave = REGISTRO    
            Pais: AN(100)
            AAEdicion: N(4)
            Ciudad: AN(100)
            Nombre: AN(100)
        FINREG
        Capacidad: N(5)
        Partidos: N(4)
    FINREG
    arch: archivo de estadio ordenado por clave
    reg: estadio


    //Definir archivo de salida --> grabar en corte del anio
    formato_salida = REGISTRO
        Pais: AN(100)
        aa: N(4)
        cant_partidos: N(5)
    FINREG


    finreg
    salida: archivo de formato_salida
    sal: formato_salida


    //Definir totalizadores (contadores o acumuladores)
    tot_pais_c1, tot_ciudad_c1: entero //consigna 1
    tot_partidos: entero //consigna 2
    tot_pais, tot_aa, tot_ciudad: entero //consigna 3


    //Definir procedimientos de corte


    PROCEDIMIENTO corte_ciudad() ES
        //Informar consigna 1 y 3
        ESC()
        //ACUMULAR NIVEL SUPERIOR
        tot_aa := tot_aa + tot_ciudad
        tot_pais_c1:= tot_pais_c1 + tot_ciudad_c1
        //ACTUALIZAR RESGUARDO
	  res_ciudad := reg.clave.ciudad
        //PONER A CERO CONTADORES


	  tot_ciudad:= 0
	  tot_ciudad_c1 :=0
    FinProcedimiento


    
PROCEDIMIENTO corte_aa() ES
        corte_ciudad()
        //Informar consigna 3
        ESC()
        //Grabar salida - consigna 2 -
        sal.pais := res_pais
        sal.aa := res_aa
        sal.cant_partidos := tot_aa
        grabar(salida,sal)
        //ACUMULAR NIVEL SUPERIOR
        tot_pais:= tot_pais + tot_aa
        //ACTUALIZAR RESGUARDO
        res_aa  := reg.clave.aa
        //PONER A CERO CONTADORES
        tot_aa := 0
    FinProcedimiento


    PROCEDIMIENTO corte_pais() ES
        corte_aa()
        //Informar consigna 1 y 3
        ESC()
        //NO ACUMULAR NIVEL SUPERIOR -> No pide tot gralw
        //ACTUALIZAR RESGUARDO
        res_pais := reg.clave.pais
        //PONER A CERO CONTADORES
        tot_pais := 0
    FinProcedimiento
PROCESO
//inicializar archivos
//inicializar resguardos
//inicializar totalizadores
tot_ciudad := 0; tot_aa := 0; tot_pais := 0


    MIENTRAS NFDA(arch) HACER
        //TRATAR CORTE
        SI res_pais <> reg.clave.pais ENTONCES
            corte_pais()
        SINO    
            SI res_aa <> reg.clave.aa ENTONCES
                corte_aa()
            SINO
                SI res_ciudad <> reg.clave.ciudad ENTONCES
                    corte_ciudad()
                FINSI
            FINSI
        FINSI


        //TRATAR REGSITRO
        //CONSIGNA 1
        SI reg.capacidad > 40000 ENTONCES
            tot_ciudad_c1 := tot_ciudad_c1 + reg.cant_partidos
        FINSI
        //CONSIGNA 2 Y 3
        tot_ciudad := tot_ciudad + reg.cant_partidos
	  esc(“mostrar datos del reg p/ punto 3”)
        LEER(arch,reg)
    FINMIENTRAS


    cerrar(arch)
    cerrar(salida)
FINACCION
