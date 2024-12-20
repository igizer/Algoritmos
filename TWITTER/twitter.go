//Twitter ha detectado muchos accesos no autorizados a cuentas de más de 1 millón de followers.
//Dada una cuenta y los accesos no autorizados detectados en el mes de Marzo/2023,
//determinar su nivel de criticidad teniendo en cuenta:
//- Si tuvo más de 100 accesos no autorizados en el mes, nivel: GRAVE
//- Si tuvo más de 35 accesos no autorizados en el mes, nivel: OBSERVABLE
//- Si tuvo 35 o menos accesos no autorizados en el mes, nivel: LEVE
ACCION TWITTER ES
	AMBIENTE
		cuenta: entero
		PROCEDIMIENTO criticidad(x:entero) ES
			SI x > 0 ENTONCES
				SI x > 100 ENTONCES
					ESC("Nivel GRAVE")
				SINO
					SI x > 35 ENTONCES
						ESC("Nivel OBSERVABLE")
					SINO
						ESC("Nivel LEVE")
					FINSI
				FINSI
			FINSI
		FINPROCEDIMIENTO
as
	PROCESOS
		ESC("A continuación, ingrese la cantidad de accesos no autorizados de la cuenta en el mes de marzo")
		LEER(cuenta)
		ESC("La críticidad de esta cuenta es: ", criticidad(cuenta))
FINACCION