Ejercicio 2.2.23¶

Una empresa de distribución de energía eléctrica posee un archivo maestro con los siguientes datos de sus clientes
(la fecha de última lectura son las mediciones realizadas hasta el mes de mayo del 2014):

CLIENTE Ordenado por Id_Casa
Id_Casa Fecha_Ult_Lectura Cant_Lecturas Prom_Lecturas Tipo_Consumidor
El campo tipo de consumidor puede ser:
A (cuando el promedio de lectura es menor a 20K)
B (cuando el promedio de lectura es menor a 40K)
C (cuando el promedio de lectura es igual o supera los 40K)

Cuenta además con el siguiente archivo con datos de las mediciones realizadas en los dos últimos años (desde junio del 2014 hasta junio del 2016)
Mediciones Ordenado por Id_Casa
Id_Casa Fecha_Medicion Consumo

Se pide actualizar el archivo maestro con la información de las lecturas hasta el mes de enero del año 2015 inclusive. Los campos que se deben
actualizar son: la fecha de ultima lectura, cantidad de lecturas, el promedio y modificar (en caso de ser necesario) el tipo de consumidor.
En caso en que no exista el ID casa (una conexión nueva) se tiene que crear en el maestro y actualizar los otros campos.




















ACCION Ejercicio_2_2_23 ES
	AMBIENTE
		FECHA= REGISTRO
			dia: 1..31
			mes: 1..12
			año: N(4)
		FINREGISTRO
		CLIENTE= REGISTRO
			Id_Casa: N(5)
			Fecha_Ult_Lectura: FECHA
			Cant_Lecturas: N(4)
			Prom_Lecturas: N(10)
			Tipo_Consumidor: ("A", "B", "C")
		FINREGISTRO
		Reg_cliente, Reg_sal: CLIENTE
		Arch_cliente, Arch_sal: Archivo de CLIENTE Ordenado por Id_Casa
		MEDICIONES= REGISTRO
			Id_Casa: N(5)
			FechaMedicion: FECHA
			Consumo: N(6)
		FINREGISTRO
		Reg_mediciones: MEDICIONES
		Arch_mediciones: Archivo de MEDICIONES Ordenado por Id_Casa
		PROCEDIMIENTO Leer_cliente ES
			LEER(Arch_cliente, Reg_cliente)
			SI FDA(Arch_cliente) ENTONCES
				Reg_cliente.Id_Casa:= HV
			FINSI
		FINPROCEDIMIENTO
		PROCEDIMIENTO Leer_mediciones ES
			LEER(Arch_mediciones, Reg_mediciones)
			SI FDA(Arch_mediciones) ENTONCES
				Reg_mediciones.Id_Casa:= HV
			FINSI
		FINPROCEDIMIENTO
	PROCESO
		ABRIR E/(Arch_cliente); ABRIR E/(Arch_mediciones)
		ABRIR /S(Arch_sal)
		Leer_cliente; Leer_mediciones
		MIENTRAS (Reg_cliente.Id_Casa <> HV) v (Reg_mediciones.Id_Casa <> HV) HACER
			SI Arch_cliente.Id_Casa < Arch_Mediciones.Id_Casa ENTONCES
				GRABAR(Arch_sal, Reg_cliente)
				Leer_cliente
			SINO
				SI Arch_cliente.Id_Casa = Arch_Mediciones.Id_Casa ENTONCES
					MIENTRAS Arch_cliente.Id_Casa = Arch_Mediciones.Id_Casa HACER
						SI (Reg_mediciones.Fecha_Medicion.año < 2015) v ((Reg_mediciones.Fecha_Medicion.año = 2015) ^
						(Reg_mediciones.Fecha_Medicion.mes < 2))
							Reg_cliente.Prom_Lecturas:= (Reg_cliente.Cant_Lecturas * Reg_cliente.Prom_Lecturas) + Reg_mediciones.Consumo / (Reg_cliente.Cant_Lecturas + 1) 
							Reg_clientes.Cant_Lecturas:= Reg_clientes.Cant_Lecturas + 1
							Reg_clientes.Fecha_Ult_Lectura:= Reg_mediciones.Fecha_Medicion						
						FINSI
						Leer_mediciones						
					FINMIENTRAS
					SEGUN Rec_cliente.Prom_Lecturas HACER
							<20: Reg_cliente.Tipo_Consumidor:= "A"
							20..39: Reg_cliente.Tipo_Consumidor:= "B"
							39>: Reg_cliente.Tipo_Consumidor:= "C" 					
					FINSEGUN
					GRABAR(Arch_sal, Reg_cliente)
					Leer_cliente
				FINSI















ACCION 2.2.23 ES
AMBIENTE
	FECHA:REGISTRO
		AÑO:N(4)
		MES:1..12
		DIA:1..31
	FIN REGISTRO
	FORMATO_CLIENTE=REGISTRO
		ID_CASA:N(4)
		FECHA_ULT_LECTURA:FECHA
		CANT_LECTURAS:N(4)
		PROM_LECTURAS:N(10)
		TIPO_CONSUMIDOR:("A","B","C")
	FIN REGISTRO
	FORMATO_MEDICIONES=REGISTRO
		ID_CASA:N(4)
		FECHA_MEDICION:FECHA
		CONSUMO:N(6)
	FIN REGISTRO
	CLIENTE:ARCHIVO DE FORMATO_CLIENTE ORDENADO POR ID_CASA
	SALIDA:ARCHIVO DE FORMATO_CLIENTE ORDENADO POR ID_CASA
	MEDICIONES:ARCHIVO DE FORMATO_MEDICIONES ORDENADO POR ID_CASA
	MAE:FORMATO_CLIENTE
	MOV:FORMATO_MEDICIONES

	PROCEDIMIENTO LEER_CLIENTE ES:
		LEER(CLIENTE, MAE)
		SI FDA(CLIENTE) ENTONCES
			MAE.ID_CASA:=HV
		FIN SI
	FIN PROCEDIMIENTO
	PROCEDIMIENTO LEER_MEDICIONES ES:
		LEER(MEDICIONES, MOV)
		SI FDA(MEDICIONES) ENTONCES
			MOV.ID_CASA:=HV
		FIN SI
	FIN PROCEDIMIENTO
INICIO
	ABRIR/E(CLIENTE)
	ABRIR/E(MEDICIONES)
	ABRIR/S(SALIDA)
	LEER_MEDICIONES
	LEER_CLIENTE
	MIENTRAS (MAE.ID_CASA <> HV) o (MOV.ID_CASA) HACER
		SI (MAE.ID_CASA < MOV.ID_CASA) ENTONCES
			GRABAR(SALIDA,MAE)
			LEER_CLIENTE
		SiNO
			SI (MAE.ID_CASA = MOV.ID_CASA) ENTONCES
				MIENTRAS (MAE.ID_CASA = MOV.ID_CASA) HACER
					SI (MOV.FECHA_MEDICION.AÑO < 2015) O 
						((MOV.FECHA_MEDICION.AÑO = 2015) Y (MOV.FECHA_MEDICION.MES = 1) ENTONCES
						MAE.PRMEDIO := (MAE.CANT_LECTURAS * MAE.PROMEDIO) + MOV.CONSUMO/  MAE.CANT_LECUTRAS + 1
						MAE.CANT_LECUTRAS := MAE.CANT_LECUTRAS + 1
						MAE.FECHA_ULT_LECTURA := MOV.FECHA_MEDICION
					FIN SI
					LEER_MEDICIONES
				FIN MIENTRAS
				SEGUN MAE.PROMEDIO HACER
					<20: MAE.TIPO_CONSUMIDOR:= 'A'
					<40: MAE.TIPO_CONSUMIDOR:= 'B'
					OTRO: MAE.TIPO_CONSUMIDOR:= 'C'
				FIN SEGUN
				GRABAR(SALIDA,MAE)
				LEER_CLIENTE
			SiNO
				MAE.ID_CASA := MOV.ID_CASA
				MAE.FECHA_ULT_LECTURA := MOV.FECHA_MEDICION
				MAE.CANT_LECUTRAS := 1
				MAE.PROMEDIO := MOV.CONSUMO
				LEER_MEDICIONES
				MIENTRAS (MAE.ID_CASA = MOV.ID_CASA) HACER
					SI (MOV.FECHA_MEDICION.AÑO < 2015) O 
						((MOV.FECHA_MEDICION.AÑO = 2015) Y (MOV.FECHA_MEDICION.MES = 1) ENTONCES
						MAE.PRMEDIO := (MAE.CANT_LECTURAS * MAE.PROMEDIO) + MOV.CONSUMO/  MAE.CANT_LECUTRAS + 1
						MAE.CANT_LECUTRAS := MAE.CANT_LECUTRAS + 1
						MAE.FECHA_ULT_LECTURA := MOV.FECHA_MEDICION
					FIN SI
					LEER_MEDICIONES
				FIN MIENTRAS
				SEGUN MAE.PROMEDIO HACER
					<20: MAE.TIPO_CONSUMIDOR:= 'A'
					<40: MAE.TIPO_CONSUMIDOR:= 'B'
					OTRO: MAE.TIPO_CONSUMIDOR:= 'C'
				FIN SEGUN
				GRABAR(SALIDA,MAE)
			FIN SI
		FIN SI
	FIN MIENTRAS

	CERRAR

FIN ACCION