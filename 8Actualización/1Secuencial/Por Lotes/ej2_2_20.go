Ejercicio 2.2.20¶

Se desea generar una secuencia, a partir de la secuencia del problema 2.01,
donde se actualice la fecha de aprobación del último exámen, para ello se cuenta
con información del último turno de exámen, según el siguiente detalle:

EXAMENES Ordenado por Nro_Legajo
||Nro_Legajo||Carrera|Cod_Materia|Fecha_Examen|Nota|

Se debe considerar que no necesariamente todos los alumnos se presentan en cada
turno de exámen, y que un alumno puede presentarse a rendir más de una materia en un mismo turno.
Supónganse ambas secuencias ordenadas por Número de Legajo.

//2.01
//Se dispone de una secuencia con los datos de los alumnos de la facultad:

//Apellido y Nombre
//Carrera (ISI - IEM - IQ)
//Nº de Legajo
//Fecha de Nacimiento
//Fecha de Ingreso
//D.N.I.
//Sexo (M o F)
//Fecha de último examen aprobado
//Nota
//Se desea un listado de los mismos, con el siguiente formato:

//Nro_LegajoApellido_NombreDocumentoCarrera|

ACCION Ejercicio2_2_20 ES
	
	AMBIENTE

		FECHA= REGISTRO
			dia: 1..31
			mes: 1..12
			año: AN(4)
		FINREGISTRO

		examen= REGISTRO
			Nro_Legajo: N(5)
			Carrera: ("ISI","IQ","IEM")
			Cod_Materia: N(6)
			Fecha_Examen: Fecha
			Nota: 1..10
		FINREGISTRO

		rex: examen
		aex: Archivo de examen Ordenado por Nro_Legajo

		facu= REGISTRO
			ApNom: AN(30)
			Carrera: ("ISI","IQ","IEM")
			Nro_Legajo: N(5)
			Fecha_Nacimiento: Fecha
			Fecha_Ingreso: Fecha
			DNI: N(8)
			Sexo: ("M", "F")
			Fecha_ult_ex_aprobado: Fecha
			Nota: 1..10
		FINREGISTRO

		rfacu, rsal: facu
		afacu, asal, aux: Archivo de facu Ordenado por Nro_Legajo

		PROCEDIMIENTO Leerfacu ES
			Leer(afacu, rfacu)
			SI FDA(afacu) ENTONCES
				rfacu:= HV
			FINSI
		FINPROCEDIMIENTO

		PROCEDIMIENTO Leerexamen ES
			Leer(aex, rex)
			SI FDA(aex) ENTONCES
				rex:= HV
			FINSI




















ACCION Ej_2_2_20 ES

	AMBIENTE

		FECHA = REGISTRO
			dia: 1..31
			mes: 1..12
			año: N(4)			
		FINREGISTRO

		EXAMEN = REGISTRO
			Nro_Legajo: N(5)
			Carrera: ("ISI", "IEM", "IQ")
			Cod_Materia: N(6)
			Fecha_Examen: FECHA
			Nota: 1..10
		FINREGISTRO

		Reg_ex: EXAMEN
		Arch_ex: Archivo de EXAMEN Ordenado por Nro_Legajo

		FACULTAD = REGISTRO
			ApNom: AN(15)
			Carrera: ("ISI", "IEM", "IQ")
			Nro_Legajo: N(5)
			Fecha_Nacimiento: FECHA
			Fecha_Ingreso: FECHA
			DNI: AN(8)
			Sexo: ("M", "F")
			Fecha_UEA: FECHA
			Nota: 1..10
		FINREGISTRO

		Reg_facu, Reg_sal: FACULTAD
		Arch_facu, Arch_sal, AUX: Archivo de FACULTAD Ordenado por Nro_Legajo

		PROCEDIMIENTO Leer_facultad ES
			LEER(Arch_facu, Reg_facu)
			SI FDA(Arch_facu) ENTONCES
				Reg_facu:= HV
			FINSI
		FINPROCEDIMIENTO

		PROCEDIMIENTO Leer_ex ES
			LEER(Arch_ex, reg_ex)
			SI FDA(Arch_ex) ENTONCES
				Reg_ex:= HV
			FINSI
		FINPROCEDIMIENTO


	PROCESO

		ABRIR E/(Arch_facu); ABRIR E/(Arch_ex)
		ABRIR /S(Arch_sal)
		Leer_facu; Leer_ex

		MIENTRAS (Reg_facu <> HV) v (Reg_ex <> HV) HACER
			SI Reg_facu.Nro_Legajo < Arch_ex.Nro_Legajo ENTONCES
				GRABAR(Arch_sal,Reg_facu.Nro_Legajo)
				Leer_facu
			SINO
				SI Reg_facu.Nro_legajo = Arch_ex.Nro_legajo ENTONCES
					AUX:= Reg_facu
					MIENTRAS Reg_facu.Nro_legajo = Arch_ex.Nro_legajo HACER
						SI Reg_ex.Nota > 5 ENTONCES
							AUX.Nota:= Reg_ex.Nota
							AUX.Fecha_UEA:= Reg_ex.Fecha_UEA
						FINSI
						Leer_ex
					FINMIENTRAS
					GRABAR(Arch_sal, AUX)
					Leer_facu
				SINO
					SI Reg_facu.Nro_Legajo > Reg_ex.Nro_Legajo ENTONCES
						ESC("El legajo no EXISTE error")
						Leer_ex
					FINSI
				FINSI
			FINSI
		FINMIENTRAS

		CERRAR(Arch_ex); CERRAR(Arch_facu); CERRAR(Arch_sal)

FINACCION
				
					
				



















Accion Act_Porlotes_2.20 Es

	AMBIENTE

	FECHA = REGISTRO
		AA: N(4);
		MM: 1..12;
		DD: 1..31;
	FR;

	FACULTAD = REGISTRO	
		N_LEG: N(6);	
		APYNOM: AN(50);
		FECHA_ING: FECHA;			
		SEXO: ("M","F");
		CARRERA: ("ISI","IEM","IQ");
		FECHA_NAC: FECHA;
		DNI: N(8):
		FECHA_UEXAPR: FECHA;
		NOTA: N(2);
	FR;

	ARCH_FAC, ARCH_SAL: Archivo de FACULTAD ordenado por N_LEG;			
	REG_FAC, REG_SAL, AUX: FACULTAD;

	EXAMENES = REGISTRO
		N_LEG: N(6);
		CARRERA: ("ISI","IEM","IQ");
		COD_MATERIA: N(6);
		FECHA_EX: FECHA;
		NOTA: N(2);
	FR;

	ARCH_EX: Archivo de EXAMENES ordenado por N_LEG;
	REG_EX: EXAMENES;

	SUBACCION LEER_FAC ES
		LEER(ARCH_FAC,REG_FAC);
		SI FDA(ARCH_FAC) ENTONCES
			REG_FAC.N_LEG:= HV;
		FS;
	FSUB;

	SUBACCION LEER_EX ES
		LEER (ARCH_EX,REG_EX);
		SI FDA(ARCH_EX) ENTONCES
			REG_EX.N_LEG:= HV;
		FS;
	FSUB;

	PROCESO
	ABRIR E/(ARCH_FAC); ABRIR E/(ARCH_EX);
	ABRIR /S(ARCH_SAL);
	LEER_FAC; LEER_EX;
	MIENTRAS (REG_FAC.N_LEG <>  HV) v (REG_EX.N_LEG <> HV) HACER
        SI REG_FAC.N_LEG < REG_EX.N_LEG ENTONCES
            REG_SAL:= REG_FAC;
            ESC(ARCH_SAL,REG_SAL);
            LEER_FAC;
        SINO 
            SI REG_FAC.N_LEG = REG_EX.N_LEG ENTONCES
                AUX:= REG_FAC;
                MIENTRAS REG_FAC.N_LEG = REG_EX.N_LEG HACER
                    SI REG_EX.NOTA > 5 ENTONCES
                        AUX.NOTA:= REG_EX.NOTA
                        AUX.FECHA_UEXAPR:= REG_EX.FECHA_EX
                    FS;
                    LEER_EX;
                FM;
                REG_SAL:= AUX;
                ESC(ARCH_SAL,REG_SAL);
                LEER_FAC;
            SINO
                SI REG_FAC.N_LEG > REG_EX.N_LEG ENTONCES
                    LEER_EX;
                    ESC("Legajo Erroneo no esta en la base de datos:", REG_EX.N_LEG) 
                FS;
            FS;
        FS;
    FM;
    CERRAR(ARCH_FAC); CERRAR(ARCH_EX); CERRAR(ARCH_SAL);
Fin_Accion.
