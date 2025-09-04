Ejercicio 2.2.19¶
En una Empresa Farmacéutica se posee un archivo MAE_REMEDIOS (ordenado por Clave: Farmacia + Medicamento),
el que se actualiza semanalmente, a traves de la información que se encuentra cargada en un archivo de 
MOVIMIENTOS(ordenado por Clavem: Farmacia + Medicamento, y Cod_Mov), de la siguiente forma:
	Si Clave (MAE_REMEDIOS) es menor que Clavem (MOVIMIENTOS), simplemente se transfieren los datos del Maestro a la salida y se graban.
	Si Clave (MAE_REMEDIOS) es igual a Clavem (MOVIMIENTOS) y el Codmov es 1, se considera error y se lista un mensaje indicando el tipo de error;
		pero si el Codmov es 2, entonces es un remedio que deja de fabricarse y se transfiere el registro al archivo de Remedios vencidos (REM_VENC);
		pero si el Cod_Mov es 3, se modifica la cantidad actual con la cantidad recibida.
	Si Clave (MAE_REMEDIOS) es mayor que Clavem (MOVIMIENTOS) y el Codmov es 1, se incorpora el remedio a nuestro Vademecum,
		considerando que la cantidad recibida configura la cantidad actual y la Fecha_Vencimiento es 30 días posterior a la fecha actual;
		pero si el Codmov es 2 o 3 se considera error y se deben producir los correspondientes mensajes de error.
Se considera que solo existe un registro de movimiento para cada registro del maestro.

MAE_REMEDIOS Ordenado por Farmacia y Medicamento
FarmaciaMedicamentoCant_ActualFecha_Vencimiento

MOVIMIENTOS Ordenado por Farmacia, Medicamento y Cod_Mov
FarmaciaMedicamentoCod_MovCant_Recibida

REM_VENC Ordenado por Medicamento
MedicamentoCant_Vencida


ACCION MAE_REMEDIOS ES
	AMBIENTE
		reg= registro
			







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
