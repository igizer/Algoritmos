ACCION primer_registro ES
	AMBIENTE
		fecha = registro
			anio: N(4)
			mes: N(2)
			dia: N(2)
		alumno = registrados
			apellido_nombre: AN(50)
			carrera: AN(3)
			legajo: N(5)
			fecha_nacimiento: fecha
			fecha_ingreso: fecha
			dni: N(8)
			sexo: ("M", "F")
			fecha_ult_examen_aprobado: fecha
			nota: 0...100
				anio: N(4)
				mes: N(2)
				dia: n(2)
		FinRegistro

		alumnos: Archivo de alumno
		al: alumno

	Proceso
		Abrir E/(alumnos); Leer(alumnos, al)
		Escribir("Legajo", "Apellido y nombre", "DNI", "Carrera", "Fecha de ingreso")
		tiene_datos:= Falso
		Mientras no FDA(alumnos) HACER
			tiene_datos:= Verdadero
			SI NO tiene_datos ENTONCES
				Escribir("No se encontraron datos")
			FINSI
			Escribir(al.legajo, al.apellido_nombre, al.dni, al.carrera, al.fecha_ingreso.dd, "/, ", 
			al.fecha_ingreso.mes, "/", al.fecha_ingreso.anio)
			Leer(alumnos, al)

		FinMientras



FinAccion

Accion 2.2-3 es 
Ambiente 
