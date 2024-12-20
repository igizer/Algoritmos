Ejercicio 2.2.21¶
En un práctico para la Facultad un grupo de alumnos debe implementar una Red Social llamada UTNBook.
Para lo cual debe utilizar los siguientes archivos:

AMIGOS Ordenado por Cod_Usuario y Cod_Amigo
||Cod_Usuario||Cod_Amigo||Fecha_Amistad|Mensaje_Muro|

Cada registro indica la fecha desde que los usuarios son amigos y el último mensaje que un amigo ha escrito en el muro del usuario.

NOTIFICACIONES Ordenado por Cod_Usuario y Cod_Amigo
||Cod_Usuario||Cod_Amigo||Cod_Movimiento|Fecha_Amistad|Mensaje_Muro|

Periódicamente se debe actualizar el archivo Amigos, con el fin de que refleje las amistades que posee cada usuario.
En el archivo notificaciones pueden existir tres tipos de acciones: la solicitud de una amistad (Cod_mov = 'A');
la eliminación de una amistad (Cod_mov = 'B'); o la información de un mensaje que un amigo ha escrito en el muro del usuario (Cod_mov = 'M').
Considerar que la eliminación de una amistad implica la baja física de un registro y que hay un solo registro de Notificaciones por
cada registro de Amigo.

ACCION EJ2_2_21 ES

	AMBIENTE

		Fecha= REGISTRO
			dd: 1..31
			mm: 1..12
			aaaa: N(4)
		FINREGISTRO

		UTNBook = REGISTRO
			Clave = REGISTRO
				Cod_Usuario: N(5)
				Cod_Amigo: N(5)
			FINREGISTRO
			Fecha_Amistad: Fecha
			Mensaje_Muro: AN(250)
		FINREGISTRO

		rmae, rsal: UTNBook
		amae, asal: Archivo de UTNBook Ordenado por Clave

		Notificaciones = REGISTRO
			Clavem = REGISTRO
				Cod_Usuario: N(5)
				Cod_Amigo: N(5)
			FINREGISTRO
			Cod_Movimiento: ("A","B","M")
			Fecha_Amistad: Fecha
			Mensaje_Muro: AN(250)
		FINREGISTRO

		rmov: Notificaciones
		amov: Archivo de Notificaciones Ordenado por Clavem

		PROCEDIMIENTO leermae ES
			LEER(amae, rmae)
			SI FDA(amae) ENTONCES
				rmae.Clave:=HV
			FINSI
		FIPROCEDIMIENTO

		PROCEDIMIENTO leermov ES
			LEER(amov, rmov)
			SI FDA(amov) ENTONCES
				rmov.Clave:=HV
			FINSI
		FIPROCEDIMIENTO

		solicitud: caracter //esto es por una flayada

	PROCESO

		ABRIR E/(amae); ABRIR E/(amov); ABRIR /S(asal)
		leermae; leermov

		MIENTRAS (regmae.Clave <> HV) v (regmov.Clave <> HV)

			SI rmae < rmov ENTONCES

				GRABAR(asal, rmae)
				leermae
			
			SINO
				SI rmae = rmov ENTONCES

					SEGUN rmov.Cod_Movimiento HACER

						"A": ESC("ERROR, CODIGO DE MOVIMIENTO: A") //Notifico error, si te sigue no te va a volver a seguir
						GRABAR(asal, rmae)
						
						"B": ESC("Se ha eliminado una amistad") //No se puede hacer baja física, solo no lo grabo en la salida y leo

						"M": ESC("Tu amigo ", rmae.Clave.Cod_Amigo, ", ha escrito un mensaje en tu muro ")
						GRABAR(asal, rmae)

					FINSEGUN
					leermae; leermov

				SINO
					SI rmae > rmov ENTONCES

						SI rmov.Cod_Movimiento = "A" ENTONCES

							ESC("Tiene una nueva solicitud del amistad del usuario ", rmov.Clave.Cod_Usuario)
							//Acá a continuación voy a flayar, puede que esté mal poner un condicional porque es actualización peor bue
							ESC("Desea aceptar la solicitud de amistad? (Escribir S/N)")
							LEER(solicitud)
							SI solicitud = "S" ENTONCES
								rsal.Clave:= rmov.Clave
								rsal.Fecha_Amistad:= rmov.Fecha_Amistad
								rsal.Mensaje_Muro:= rmov.Mensaje_Muro
								GRABAR(asal, rsal)
							SINO
								ESC("No se ha aceptado la solicitud de amistad")
							FINSI
						
						SINO

							ESC("ERROR, CODIGO DE MOVIMIENTO: ", rmov.Cod_Movimiento) // Notifico el error, no te puede dejar de seguir quien no te sigue
																					  //y tampoco escribir en tu muro
						FINSI

						leermov

					FINSI
				FINSI
			FINSI
		FINMIENTRAS
	
	CERRAR(amae); CERRAR(amov); CERRAR(asal)

FINACCION

						













ACCION EJ2_2_21 ES

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

		PROCEDIMIENTO Leer_notif
			Leer(Arch_notig, reg_notif)
			SI FDA(Arch_notif) ENTONCES
				Reg_notif.CLAVE:= HV
			FINSI
		FINPROCEDIMIENTO

	PROCESO

		ABRIR E/(Arch_amigo); ABRIR E/(Arch_notif)
		ABRIR /S(Arch_sal)

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
					Leer_amigo; Leer_notif

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
















			SI Arch_amigo.Cod_Movimiento = "A" ENTONCES
				ESC("¡Tiene una nueva solicitud de amistad!")
				Leer_

			SINO
				SI Arch_amigo.Cod_Movimiento = "B" ENTONCES
					ESC("Se ha eliminado una amistad :c")
					//en actualizacion secuencial no se puede borrar fisicamente. solo se avanza

				SINO
					SI Arch_amigo.Cod_Movimiento = "M" ENTONCES
						ESC("¡Nuevo mensaje en su muro!")
					FINSI

				FINSI

			FINSI

		FINMIENTRAS


















Accion Act_Unitaria_2.21 Es

	AMBIENTE

	FECHA = REGISTRO
		AA: N(4);
		MM: 1..12;
		DD: 1..31;
	FR;

	AMIGOS = REGISTRO
		CLAVE : REGISTRO
			COD_USUARIO: N(6);
			COD_AMIGO: N(5);
		FR;
		FECHA_AMISTAD:FECHA;
		MENSAJE_MURO: AN(50);
	FR;

	A_AMI,A_SAL: Archivo de AMIGOS ordenado por CLAVE;
	R_AMI,R_SAL: AMIGOS;

	NOTIFICACIONES = REGISTRO
		CLAVE : REGISTRO
			COD_USUARIO: N(6);
			COD_AMIGO: N(5);
		FR;
		COD_MOV: ("A","B","M");
		FECHA_AMISTAD:FECHA;
		MENSAJE_MURO: AN(50);
	FR,

	A_NOT: Archivo de NOTIFICACIONES ordenado por CLAVE;
	R_NOT: NOTIFICACIONES;

	SUBACCION LEER_AMI ES
		LEER(A_AMI,R_AMI);
		SI FDA(A_AMI) ENTONCES
			R_AMI.CLAVE.COD_USUARIO:= HV;
			R_AMI.CLAVE.COD_AMIGO:= HV;
		FS;
	FSUB;

	SUBACCION LEER_NOT ES
		LEER(A_NOT,R_NOT)
		SI FDA(A_NOT) ENTONCES
			R_NOT.CLAVE.COD_USUARIO:= HV;
			R_NOT.CLAVE.COD_AMIGO:= HV;
		FS;
	FSUB;

	PROCESO

	ABRIR E/(A_AMI); ABRIR E/(A_NOT);
	ABRIR /S(A_SAL);
	LEER_AMI; LEER_NOT;
	MIENTRAS (R_AMI.CLAVE <> HV) v (R_NOT.CLAVE <> HV) HACER
		SI R_AMI.CLAVE < R_NOT.CLAVE ENTONCES
			R_SAL:= R_AMI;
			ESCRIBIR(A_SAL,R_SAL);
			LEER_AMI;
		SINO
			SI R_AMI.CLAVE = R_NOT.CLAVE ENTONCES
				SEGUN R_NOT.COD_MOV HACER
					"A": ESCRIBIR ("ERROR SOLICITUD INVALIDA");
						R_SAL:= R_AMI;
						ESCRIBIR(A_SAL,R_SAL);


					"B": ESC("SE A ELIMINADO AL AMIGO", R_AMI.COD_AMIGO )  
                        //en actualizacion secuencial no se puede borrar fisicamente. solo se avanza

					"M": 
						R_AMI.MENSAJE_MURO:= R_NOT.MENSAJE_MURO;
						R_SAL:= R_AMI;
						ESCRIBIR(A_SAL,R_SAL);
				FS;
				LEER_AMI,LEER_NOT;
			SINO
				SI R_AMI.CLAVE > R_NOT.CLAVE ENTONCES
					SEGUN R_NOT.COD_MOV HACER
						"A": 
							R_SAL.CLAVE.COD_USARIO := R_NOT.CLAVE.COD_USUARIO;
							R_SAL.CLAVE.COD_AMIGO :=  R_NOT.CLAVE.COD_AMIGO;
							R_SAL.FECHA_AMISTAD:= R_NOT.FECHA_AMISTAD;
							R_SAL.MENSAJE_MURO := R_NOT.MENSAJE_MURO;
							ESCRIBIR(A_SAL, R_SAL);

						"B": ESCRIBIR("BAJA INVALIDA");

						"M": ESCRIBIR("MODIFICACION INVALIDA");
					FS;
					LEER_NOT;
				FS;
			FS;
		FS;	
	FM;
	CERRAR(A_MI); CERRAR(A_SAL); CERRAR(A_NOT);
Fin_Accion.