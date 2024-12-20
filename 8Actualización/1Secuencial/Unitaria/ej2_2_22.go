








     NO ES DIFICIL PERO SALTEO NMS, ALGO INTENTÉ IGUAL










Ejercicio 2.2.22
IDEM ejercicio 2.21 pero agregando al enunciado: “Al finalizar el proceso se desea conocer: el usuario que posee más amigos.”

//Ejercicio 2.2.21¶
//En un práctico para la Facultad un grupo de alumnos debe implementar una Red Social llamada UTNBook.
//Para lo cual debe utilizar los siguientes archivos:
//AMIGOS Ordenado por Cod_Usuario y Cod_Amigo
//||Cod_Usuario||Cod_Amigo||Fecha_Amistad|Mensaje_Muro|
//Cada registro indica la fecha desde que los usuarios son amigos y el último mensaje que un amigo ha escrito en el muro del usuario.
//NOTIFICACIONES Ordenado por Cod_Usuario y Cod_Amigo
//||Cod_Usuario||Cod_Amigo||Cod_Movimiento|Fecha_Amistad|Mensaje_Muro|
//Periódicamente se debe actualizar el archivo Amigos, con el fin de que refleje las amistades que posee cada usuario.
//En el archivo notificaciones pueden existir tres tipos de acciones: la solicitud de una amistad (Cod_mov = 'A');
//la eliminación de una amistad (Cod_mov = 'B'); o la información de un mensaje que un amigo ha escrito en el muro del usuario (Cod_mov = 'M').
//Considerar que la eliminación de una amistad implica la baja física de un registro y que hay un solo registro de Notificaciones por
//cada registro de Amigo.

ACCION EJ2_2_22 ES

	AMBIENTE

		FECHA = REGISTRO
			dia: 1..31
			mes: 1..12
			año: AN(4)
		FINREGISTRO

		AMIGO = REGISTRO
			CLAVE = REGISTRO
				Cod_Usuario: N(5)
				Cod_Amigo: N(5)
			FINREGISTRO
			Fecha_Amistad: FECHA
			Mensaje_Muro: AN(250)
		FINREGISTRO

		Reg_amigo, Reg_sal: AMIGO
		Arch_amigo, Arch_sal: Archivo de AMIGO Ordenado por Clave

		NOTIFICACIONES = REGISTRO
			CLAVE = REGISTRO
				Cod_Usuario: N(5)
				Cod_Amigo: N(5)
			FINREGISTRO
			Cod_Movimiento: ("A", "B", "M")
			Fecha_Amistad: FECHA
			Mensaje_Muro: AN(250)
		FINREGISTRO

		Reg_notif: NOTIFICACIONES
		Arch_notif: Archivo de NOTIFICACIONES Ordenado por CLAVE

		PROCEDIMIENTO Leer_Amigo ES
			Leer(Arch_amigo, Reg_amigo)
			SI FDA(Arch_amigo) ENTONCES
				Reg_amigo.CLAVE:= HV
			FINSI
		FINPROCEDIMIENTO

		PROCEDIMIENTO Leer_notif ES
			Leer(Arch_notig, reg_notif)
			SI FDA(Arch_notif) ENTONCES
				Reg_notif.CLAVE:= HV
			FINSI
		FINPROCEDIMIENTO

		RESG_USUARIO, RESG_MAS_AMIGOS: N(5)
		Cont_amigos, Mayor_Cantidad_Amigos: entero

		PROCEDIMIENTO Mas_amigos ES
			SI Reg_amigo.Cod_Usuario = RESG_USUARIO ENTONCES
				Cont_amigos:= Cont_amigos + 1
				SI Cont_amigos > Mayor_Cantidad _Amigos ENTONCES
						Mayor_Cantidad _Amigos:= Cont_amigos
						
						
					FINSI
			SINO
				SI Reg_amigo.Cod_Usuario <> HV ENTONCES
					RESG_USUARIO:= Reg_amigo.Cod_Usuario
				FINSI
			FINSINO
			Leer_amigo
		FINPROCEDIMIENTO

					



	PROCESO

		ABRIR E/(Arch_amigo); ABRIR E/(Arch_notif)
		ABRIR /S(Arch_sal)

		Cont_amigos:= 0
		RESG_USUARIO:= Reg_amigo.Cod_Usuario
		Leer_amigo; Leer_notif

		MIENTRAS (Reg_amigo.CLAVE <> HV) v (Reg_notif.CLAVE <> HV) HACER

			SI Reg_amigo.CLAVE < Reg_notif.CLAVE ENTONCES//Tratamos a los amigos
				GRABAR(Arch_sal, Reg_amigo)
				Leer_amigo
			
			SINO
				SI Reg_amigo.CLAVE = Reg_Notif.CLAVE ENTONCES
					SEGUN Reg_amigo.CLAVE HACER
						"A":
							ESC("Error, el usuario", Reg_notif.CLAVE, " ya lo sigue de antemano")
							GRABAR(Arch_sal, Reg_amigo)
						
						"B":
							ESC("El usuario", Reg_notif.CLAVE, " ha eliminado la amistad")
							//en actualizacion secuencial no se puede borrar fisicamente. solo se avanza
							//simplemente no lo voy a grabar en el nuevo archivo para actualizar así la lista de amigos sin él

						"C":
							ESC("El usuario", Reg_notif.CLAVE, " ha escrito un nuevo mensaje en tu muro")
							GRABAR(Arch_sal, Reg_amigo)
					FINSEGUN

					Mas_amigos; Leer_amigo; Leer_notif

				SINO
					SI Reg_amigo.CLAVE > Reg_Notif.CLAVE ENTONCES
						SEGUN Reg_amigo.CLAVE HACER
							"A":
								ESC("Tiene una solicitud de amistad de: ", Reg_notif.CLAVE)
								Reg_sal.CLAVE.Cod_Amigo:= Reg_notif.CLAVE.Cod_Amigo
								Reg_sal.CLAVE.Cod_Usuario:= Reg_notif.CLAVE.Cod_Usuario
								Reg_sal.Fecha_Amistad:= Reg_notif.Fecha_Amistad
								Reg_sal.Mensaje_Muro:= Reg_notif.Mensaje_Muro
								GRABAR(Arch_sal, Reg_sal)

							"B":
								ESC("Solicitud inválida porque el usuario no es su amigo como para dejarlo de seguir")
							
							"M":
								ESC("Error, el usuario no envió un mensaje en su muro, no es su amigo")
						FINSEGUN
						Leer_notif
					FINSI
				FINSI
			FINSI
		FINMIENTRAS