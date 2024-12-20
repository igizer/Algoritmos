Ejercicio 2.2.24¶
Una obra social lleva el control de los costos que le insume cada afiliado. Considerando como costos, los pagos que la misma debe realizar (en conceptos de honorarios, pagos a farmacias, etc) por los distintos servicios que el afiliado utiliza. Para esto la empresa cuenta con un archivo COSTOS_POR_AFILIADO, en el cual se registran la cantidad de atenciones y/o servicios que solicita el empleado y el costo total (en pesos) que la empresa debe pagar por estos. Este archivo está conformado por registros con el siguiente formato, y está ordenado Cod_Afiliado

COSTOS_POR_AFILIADO Ordenado por Cod_Afiliado

Cod_AfiliadoFecha_AltaFecha_BajaCant_ServiciosCosto
Cada viernes en la organización, se lleva a cabo un proceso de actualización del archivo COSTOS_POR_AFILIADO. Para poder realizarlo se cuenta con un archivo SERVICIOS_SEMANALES, ordenado por cod_afiliado y fecha_at, en el cual se registran los servicios que solicitaron los afiliados durante cada semana. Cada registro de este archivo presenta el siguiente formato,

SERVICIOS_SEMANALES Ordenado por Cod_Afiliado

Cod_AfiliadoFecha_AtCod_ServicioCosto
En el archivo SERVICIOS_SEMANALES puede existir más de un registro por cada afiliado.

Para realizar el proceso de actualización se deben tener en cuenta las siguientes consideraciones.

*Si algún Cod_Afiliado de serviciosSemanales, no se encuentra en ningún registro del archivo COSTOS_POR_AFILIADO, y el primer Cod_Servicio asociado al mismo es igual a 001 entonces se trata de un nuevo afiliado que fue dado de alta.

*Si Cod_Afiliado de serviciosSemanales es igual al de COSTOS_POR_AFILIADO y el Cod_Servicio es igual a 000 se trata de un afiliado que fue dado de baja y se considera como fecha de baja el valor que reside en Fecha_At. Si en cambio, el Cod_Servicio tiene algún otro valor este representa un servicio o atención que solicitó el afiliado; por lo tanto deben contabilizarse la cantidad de servicios que solicitó; como así también los costos que estos insumen.

*Si algún Cod_Afiliado de COSTOS_POR_AFILIADO no se encuentra en el archivo SERVICIOS_SEMANALES, se trata de un afiliado que no utilizó los servicios en la semana que se realiza la actualización.

Cualquier otro caso distinto a los considerados anteriormente se toma como un error, y deben ser informados por pantalla; indicando claramente el motivo del error.


ACCION Ej_2_2_24 ES

	AMBIENTE

		Fecha = REGISTRO
			dd: 1..31
			mm: 1..12
			aaaa: N(4)
		FINREGISTRO

		COSTOS_POR_AFILIADO = REGISTRO

			Cod_Afiliado: N(5)
			Fecha_Alta: Fecha
			Fecha_Baja: Fecha
			Cant_Servicios: N(2)
			Costos: N(6)

		FINREGISTRO

		amae, asal: Archivo de COSTOS_POR_AFILIADO Ordenado por Cod_Afiliado
		rmaer, rsal: COSTOS_POR_AFILIADO

		SERVICIOS_SEMANALES = REGISTRO

			Cod_Afiliado: N(5)
			Fecha_At: Fecha
			Cod_Servicio: N(3)
			Costo: N(6)

		FINREGISTRO

		amov: Archivo de SERVICIOS_SEMANALES Ordenado por Cod_Afiliado, Fecha_At
		rmov: SERVICIOS_SEMANALES

		contservicios, acumservicios: entero // cont es cantidad, acum es costo

		PROCEDIMIENTO leermae ES
			LEER(amae, rmae)
			SI FDA(amae)
				rmae.Cod_Afiliado:= HV
			FINSI
		FINPROCEDIMIENTO

		PROCEDIMIENTO leermov ES
			LEER(amov, rmov)
			SI FDA(amov)
				rmov.Cod_Afiliado:= HV
			FINSI
		FINPROCEDIMIENTO

	PROCESO

		ABRIR E/(amae); ABRIR E/(amov)
		ABRIR /S(asal)
		
		MIENTRAS (rmae.Cod_Afiliado <> HV) v (rmov.Cod_Afiliado <> HV) HACER

			SI rmae.Cod_Afiliado < rmov.Cod_Afiliado ENTONCES

				GRABAR(asal, rmae)
				leermae

			SINO

				SI rmae.Cod_Afiliado = rmov.Cod_Afiliado ENTONCES



					MIENTRAS rmae.Cod_Afiliado = rmov.Cod_Afiliado HACER

						SI rmov.Cod_Servicio = 000 entonces

							rsal:= rmaer
							rsal.Fecha_Baja:= rmov.fecha_At
			
						SINO

							rsal:= rmae
							contservicios:= contservicios + 1
							acumservicios:= acumservicios + rmov.Costo

						FINSI

						rsal.Cant_Servicios:= contservicios
						rsal.Costos:= acumservicios

						leermov

					FINMIENTRAS

					contservicios:= 0
					acumservicios:= 0

					GRABAR(asal, rsal)

					leermae

				SINO

					SI rmae.Cod_Afiliado > rmov.Cod_Afiliado ENTONCES

						SI rmov.Cod_Servicio = 001 ENTONCES

							rsal.Cod_Afiliado:= rmov.Cod_Afiliado
							rsal.Fecha_Alta:= rmov.Fecha_At
							rsal.Costos:= rmov.Costo
							GRABAR(asal, rsal)

						ESC("ERROR, NO SE PUEDE DAR DE ALTA") // El único error a notificar que tenemos, no existe otro error aparentemente

						FINSI
					
						leermov

					FINSI

				FINSI
			
			FINSI

		FINMIENTRAS

		CERRAR(asal); CERRAR(amae); CERRAR(amov)

FINACCION