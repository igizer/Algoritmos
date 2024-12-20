//Ejercicio 2.2.19¶
//En una Empresa Farmacéutica se posee un archivo MAE_REMEDIOS (ordenado por Clave: Farmacia + Medicamento),
//el que se actualiza semanalmente, a traves de la información que se encuentra cargada en un archivo de 
//MOVIMIENTOS(ordenado por Clavem: Farmacia + Medicamento, y Cod_Mov), de la siguiente forma:
	//Si Clave (MAE_REMEDIOS) es menor que Clavem (MOVIMIENTOS), simplemente se transfieren los datos del Maestro a la salida y se graban.
	//Si Clave (MAE_REMEDIOS) es igual a Clavem (MOVIMIENTOS) y el Codmov es 1, se considera error y se lista un mensaje indicando el tipo de error;
		//pero si el Codmov es 2, entonces es un remedio que deja de fabricarse y se transfiere el registro al archivo de Remedios vencidos (REM_VENC);
		//pero si el Cod_Mov es 3, se modifica la cantidad actual con la cantidad recibida.
	//Si Clave (MAE_REMEDIOS) es mayor que Clavem (MOVIMIENTOS) y el Codmov es 1, se incorpora el remedio a nuestro Vademecum,
		//considerando que la cantidad recibida configura la cantidad actual y la Fecha_Vencimiento es 30 días posterior a la fecha actual;
		//pero si el Codmov es 2 o 3 se considera error y se deben producir los correspondientes mensajes de error.
//Se considera que solo existe un registro de movimiento para cada registro del maestro.

//MAE_REMEDIOS Ordenado por Farmacia y Medicamento
//FarmaciaMedicamentoCant_ActualFecha_Vencimiento

//MOVIMIENTOS Ordenado por Farmacia, Medicamento y Cod_Mov
//FarmaciaMedicamentoCod_MovCant_Recibida

//REM_VENC Ordenado por Medicamento
//MedicamentoCant_Vencida

ACCION Ejercicio_2_2_19 ES

	AMBIENTE
		Fecha = REGISTRO
			dia: 1..31
			mes: 1..12
			año: N(4)
		FINREGISTRO

		Formato_mae= REGISTRO
			Clave= REGISTRO
				Farmacia: N(5)
				Medicamento: N(5)
			FINREGISTRO
			Cant_Actual: N(4)
			Fecha_Vencimiento: Fecha
		FINREGISTRO

		Formato_mov= REGISTRI
			Clavem= REGISTRO
				Farmacia: N(5)
				Medicamento: N(5)
			FINREGISTRO
			Codmov: 1..3
			Cant_Recibida: N(4)
		FINREGISTRO

		Formato_Vencido= REGISTRO
			Medicamento: N(5)
			Cant_Vencida: N(4)
		FINREGISTRO

		rmae, rsal: Formato_mae
		rmov: Formato_mov
		rvenc: Formato_Vencido
		archmae, archsal: Archivo de rmae Ordenado por Clave
		archmov: Archivo de rmov ordenado por Clavem y Codmov
		archvenc: Archivo de rvenc ordenado por Medicamento

		PROCEDIMIENTO Leer_mae ES
			LEER(archmae, rmae)
			SI FDA(archmae) ENTONCES
				rmae.Clave:= HV
			FINSI
		FINPROCEDIMIENTO

		PROCEDIMIENTO Leer_mov ES
			LEER(archmov, rmov)
			SI FDA(archmov) ENTONCES
				rmov.Clavem:= HV
			FINSI
		FINPROCEDIMIENTO

	PROCESO

		ABRIR E/(archmae); ABRIR E/(archmov); ABRIR /S(archsal); ABRIR /S(archvenc)
		Leermae; Leermov

		MIENTRAS (rmae.Clave <> HV) v (rmov.Clavem <> HV) HACER

			SI rmae.Clave < rmov.Clavem ENTONCES
				GRABAR(archmae, rmae)
				Leermae
			SINO
				SI rmae.Clave = rmov.Clavem ENTONCES
					SEGUN rmov.Codmov HACER
						1: 
							ESC("ERROR, Código de movimiento 1")
						2: 
							rvenc.Medicamento:= rmae.Clave.Medicamento
							rvenc.Cant_Recibida:= rmae.Clave.Cant_Actual
							GRABAR(archvenc, rvenc)
						3:
							rsal.Clave:= rmae.Clave
							rsal.Fecha_Vencimiento:= rmae.Fecha_Vencimiento
							rsal.Cant_Actual:= rmov.Cant_Recibida
							GRABAR(archsal, rsal)
					FINSEGUN
					Leermae; Leermov
				SINO
					SI rmae.Clave > rmov.Clave ENTONCES
						SI rmov.Codmov = 1 entonces
								rsal.Clave:= rmov.Clavem
								rsal.Cant_Actual:= rmov.Cant_Recibida
								rsal.Fecha_Vencimiento:= Fecha_Actual() + 30
								GRABAR(archsal, rsal)
						SINO
							ESC("ERROR EN CÓDIGO DE MOVIMIENTO")
						FINSI
						Leermov
					FINSI
				FINSI
			FINSI

		FINMIENTRAS

		CERRAR(archmae); CERRAR(archmov); CERRAR(archvenc); CERRAR(archsal)

FINACCION


































ACCION Ej_2_2_19Unitaria ES //Creo que es Por Lotes, por definición noto eso

	AMBIENTE

		Fecha = REGISTRO
			dd: 1..31
			mm: 1..12
			aaaa: N(4)
		FINREGISTRO

		formato_mae = REGISTRO
			Clave = REGISTRO
				Farmacia: N(5)
				Medicamento: N(5)
			FINREGISTRO
			Cant_Actual: N(4)
			Fecha_Vencimiento: Fecha
		FINREGISTRO

		regmae, regsal: formato_mae
		archmae, archsal: Archivo de formato_mae Ordenado por Clave

		formato_mov = REGISTRO
			Clave = REGISTRO
				Farmacia: N(5)
				Medicamento: N(5)
				Cod_Mov: 1..3
			FINREGISTRO
			Cant_Recibida: N(4)
		FINREGISTRO

		regmov: formato_mov
		archmov: Archivo de formato_mov Ordenado por Clave

		formato_vencido = REGISTRO
			Medicamento: N(5)
			Cant_Vencida: N(4)
		FINREGISTRO

		regvenc: formato_vencido
		archvenc: Archivo de formato_vencido Ordenado por Medicamento

		PROCEDIMIENTO Leermae ES
			LEER(archmae, regmae)
			SI FDA(archmae) ENTONCES
				regmae.Clave:= HV
			FINSI
		FINPROCEDIMIENTO

		PROCEDIMIENTO Leermov ES
			LEER(archmov, regmov)
			SI FDA(archmov) ENTONCES
				regmov.Clave:= HV
			FINSI
		FINPROCEDIMIENTO


	PROCESO

		ABRIR E/(archmae); ABRIR E/(archmov); ABRIR /S(archsal); ABRIR /S(archvenc)
		Leermae; Leermov

		MIENTRAS (regmae.Clave <> HV) v (regmov.Clave <> HV) HACER

			SI regmae.Clave < regmov.Clave ENTONCES

				GRABAR(archsal, regmae)
				Leermae

			SINO
				SI regmae.Clave = regmov.Clave ENTONCES

					SEGUN regmae.Clave.Cod_Mov HACER

						1: ESC("ERROR Código de Movimiento: 1")
							
						2: regvenc.Medicamento:= Regmae.Clave.Medicamento
							regvenc.Cant_Vencida:= Regmaer.Cant_Actual
							GRABAR(archvenc, regvenc)
						
						3: regsal.Clave:= regmae.Clave
							regsal.Cant_Actual:= regmov.Cant_Recibida
							regsal.Fecha_Vencimiento:= regmae.Fecha_Vencimiento
							GRABAR(archsal, regsal)
					
					FINSEGUN
					Leermae; Leermov
				
				SINO
					SI regmae.Clave > regmov.Clave ENTONCES

						SI regmae.Clave.Cod_Mov = 1 ENTONCES
							regsal.Clave.Farmacia:= regmov.Clave.Farmacia
							regsal.Clave.Medicamento:= regmov.Clave.Medicamento
							regsal.Cant_Actual:= regmov.Cant_Recibida
							regsal.Fecha_Vencimiento:= Fecha_Actual() + 30 //Existe una función en el sistema que devuelve la fecha actual
						SINO
							ESC("ERROR Código de Movimiento: ", regmov.Clave.Cod_Mov)
						FINSI
						Leermov
					FINSI
				FINSI
			FINSI
		FINMIENTRAS

		CERRAR(archmae); CERRAR(archmov); CERRAR(archsal); CERRAR(archvenc)

FINACCION





























ACCION ACTUALIZACION_UNITARIA_2_19 ES

	AMBIENTE
	
		FECHA = REGISTRO
			dia: 1..31
			mes: 1..12
			año: N(4)
		FINREGISTRO

		REMEDIOS = REGISTRO
			CLAVE = REGISTRO
				Farmacia: N(2)
				Medicamento: N(5)
			FINREGISTRO
			Cant_Actual: N(4)
			Fecha_Vencimiento: FECHA
		FINREGISTRO

		Reg_mae, Reg_sal: REMEDIOS
		Arch_mae, Arch_sal: Archivo de REMEDIOS Ordenado por Clave

		MOVIMIENTOS = REGISTRO
			CLAVE = REGISTRO
				Farmacia: N(2)
				Medicamento: N(5)
			FINREGISTRO
			Cod_Mov: 1..3
			Cant_Recibida: N(4)
		FINREGISTRO

		Reg_mov: MOVIMIENTOS
		Arch_mov: Archivo de MOVIMIENTOS Ordenado por Clave, Cod_Mov

		VENCIDOS = REGISTRO
			Medicamento: N(5)
			Cant_Vencida: N(4)
		FINREGISTRO

		Reg_vencidos: VENCIDOS
		Arch_vencidos: Archivo de VENCIDOS Ordenado por Medicamento

		PROCEDIMIENTO Leer_mae ES
			LEER(Arch_mae, Reg_mae)
			SI FDA(Arch_mae) ENTONCES
				Reg_mae.Clave:= HV
			FINSI
		FINPROCEDIMIENTO

		PROCEDIMIENTO Leer_mov ES
			LEER(Arch_mov, Reg_mov)
			SI FDA(Arch_mov) ENTONCES
				Reg_mov.clave:= HV
			FINSI
		FINPROCEDIMIENTO


	PROCESO


		ABRIR E/(Arch_mae); ABRIR E/(Arch_mov)
		ABRIR /S(Arch_sal); ABRIR /S(Arch_vencidos)
		Leer_mae; Leer_mov


		MIENTRAS (Arch_mae <> HV) v (Arch_mov <> HV) HACER

			SI Reg_mae.Clave < Reg_mov.Clave ENTONCES
				GRABAR(Arch_sal, Reg_mae)
				Leer_mae
			SINO
				SI Reg_mae.Clave = Reg_mov.Clave ENTONCES
					SEGUN Reg_mov.Cod_Mov HACER
						1:
							ESCRIBIR("ERROR")
							GRABAR(Arch_sal, Reg_mae)

						2:
							Reg_vencidos.Medicamento:= Reg_mae.Clave.Medicamento
							Reg_vencidos.Cant_Vencida:= Reg_mae.Cant_Actual
							GRABAR(Arch_vencidos, Reg_vencidos)
						
						3: 
							Reg_sal:= Reg_mae
							Reg_sal.Cant_Actual:= Reg_mov.Cant_Recibida
							GRABAR(Arch_sal, Reg_sal)
					FINSEGUN
					Leer_mae, Leer_mov
				SINO
					SI reg_mae.Clave > Reg_mov.Clave ENTONCES
						SI Reg_mov.Cod_Mov = 1 ENTONCES
							Reg_sal.Clave:= Reg_mov.Clave
							Reg_sal.Cant_Actual:= Reg_mov.Cant_Recibida
							Reg_sal.Fecha_Vencimiento:= Fecha_Actual() + 30 //Existe una función en el sistema que devuelve la fecha actual
							GRABAR(Arch_sal, Reg_sal)
						SINO
							ESCRIBIR("ERROR")
						FINSI
						Leer_mov
					FINSI
				FINSI
			FINSI
		FINMIENTRAS

		CERRAR(Arch_mae); CERRAR(Arch_mov); CERRAR(Arch_sal); CEERAR(Arch_vencidos)

FINACCION
		
					








Accion Act_Unitaria_2_19 Es

	AMBIENTE

	FECHA = REGISTRO
		AA: N(4);
		MM: 1..12;
		DD: 1..31;
	FR;

	MAE= REGISTRO
		CLAVE: REGISTRO
			FARMACIA: N(2);
			MEDICAMENTO: N(5);
		FR;
		CANT_ACTUAL: N(6);
		FECHA-VENCE:FECHA;		
	FR;

	ARCH_MAE, ARCH_SAL: Archivo de MAE ordenado por CLAVE;
	REG_MAE, REG_SAL : MAE;

	MOV = REGISTRO
		CLAVE: REGISTRO
			FARMACIA: N(2);
			MEDICAMENTO: N(5);
		FR;
		CODMOD: 1..3;
		CANT-RECIBIDA: N(6);
	FR;

	ARCH_MOV: Archivo de MOV ordenado por CLAVE,CODMOD;
	REG_MOV: MOV;

	VENCIDOS = REGISTRO
		MEDICAMENTO: N(5);
		CANT-VENCIDA: N(6);
	FR;

	ARCH_VEN: Archivo de VENCIDOS ordenado por MEDICAMENTO;
	REG_VEN: VENCIDOS;

	SUBACCION LEER_MAE ES
		LEER(ARCH_MAE,REG_MAE);
		SI FDA(ARCH_MAE) ENTONCES
			REG_MAE.CLAVE:= HV;
		FS;
	FSUB;

	SUBACCION LEER_MOV ES
		LEER(ARCH_MOV, REG_MOV)
		SI FDA(ARCH_MOV) ENTONCES
			REG_MOV.CLAVE:= HV;
		FS;
	FSUB;

	PROCESO
	ABRIR E/(ARCH_MAE); ABRIR E/(ARCH_MOV);
	ABRIR /S(ARCH_SAL); ABRIR /S(ARCH_VEN);
	LEER_MAE; LEER_MOV;
	
	MIENTRAS (ARCH_MAE <> HV) v (ARCH_MOV <> HV) HACER			
		SI REG_MAE.CLAVE < REG_MOV.CLAVE ENTONCES					
			REG_SAL:= REG_MAE;
			ESCRIBIR(ARCH_SAL,REG_SAL);
			LEER_MAE;
		SINO
			SI REG_MAE.CLAVE = REG_MOV.CLAVE ENTONCES
				SEGUN REG_MOV.COD_MOD HACER
					1: ESCRIBIR("Error alta invalida");
						REG_SAL:= REG_MAE;
						ESCRIBIR(ARCH_SAL,REG_SAL);

					2: 	
						REG_VEN.MEDICAMENTO:= REG_MAE.CLAVE.MEDICAMENTO;
						REG_VEN.CANT_VENCIDA:= REG_MAE.CANT-ACTUAL;
						ESCRIBIR(ARCH_VEN,REG_VEN);

					3:
						REG_SAL:= REG_MAE;
						REG_SAL.CANT-ACTUAL:= REG_MOV.CANT-RECIBIDA:
						ESCRIBIR(ARCH_SAL,REG_SAL);
				FS;
				LEER_MAE; LEER_MOV;
			SINO

				SI REG_MAE.CLAVE > REG_MOV.CLAVE ENTONCES
					SEGUN REG_MOV.CODMOD HACER					
						1:
							REG_SAL.CLAVE.FARMACIA:= REG_MOV.CLAVE.FARMACIA;
							REG_SAL.CLAVE.MEDICAMENTO:= REG_MOV.CLAVE.MEDICAMENTO;
							REG_SAL.CANT-ACTUAL:= REG_MOV.CANT-RECIBIDA;
							REG_SAL.FECHA-VENCE:= FECHA_ACTUAL() + 30;					//Funcion por defecto de la pc
							ESCRIBIR(ARCH_SAL,REG_SAL);

						2: ESCRIBIR("Error baja invalida");

						3: ESCRIBIR("Error modificacion invalida");
					FS;
					LEER_MOV;
				FS;
			FS;
		FS;
	FM;
	CERRAR(ARCH_MAE); CERRAR(ARCH_MOV); CERRAR(ARCH_SAL); CERRAR(ARCH_VEN);

Fin Accion.










AMBIENTE
FECHA= REGISTRO
	dia: 1..31
	mes: 1..12
	año: N(4)
FINREGISTRO

REMEDIOS= REGISTRO
	Clave= REGISTRO
		Farmacia: N(5)
		Medicamento: N(3)
	FINREGISTRO
	Cant_Actual
	Fecha_Vencimiento: FECHA
FINREGISTRO

Reg_mae, Reg_sal: REMEDIOS
Arch_mae, Arch_sal: ARCHIVO de REMEDIOS Ordenado por clave

MOVIMIENTOS= REGISTRO
	Clave= REGISTRO
		Farmacia: N(5)
		Medicamento: N(3)
	FINREGISTRO
	Cod_Mov: 1..3
	Cant_Recibida
FINREGISTRO

Reg_mov: MOVIMIENTOS
Arch_mov:: ARCHIVO de MOVIMIENTOS Ordenado por CLave, Cod_Mov

VENCIDOS= REGISTRO
	Medicamento: N(3)
	Cant_Vencida: N(3)
FINREGISTRO

Reg_vencidos: VENCIDOS
Arch_vencidos: ARCHIVO de VENCIDOS Ordenado por Medicamento

PROCEDIMIENTO Leer_mae ES
	LEER(Arch_mae, Reg_mae)
	SI FDA(Arch_mae) ENTONCES
		Reg_mae.Clave:= HV
	FINSI
FINPROCEDIMIENTO

PROCEDIMIENTO Leer_mov ES
	LEER(Arch_mov, Reg_mov)
	SI FDA(Arch_mov) ENTONCES
		Reg_mov.Clave:= HV
	FINSI
FINPROCEDIMIENTO

PROCESO

ABRIR E/(Arch_mae); ABRIR E/(Arch_mov)
ABRIR /S(Arch_sal); ABRIR /S(Arch_vencidos)
Leer_mae; Leer_mov

MIENTRAS (Arch_mae <> HV) v (Arch_mov <> HV) HACER
	




