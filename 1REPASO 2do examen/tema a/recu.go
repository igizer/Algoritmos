//Terminado, pero no supe usar el arreglo para valor agregado (01:13:00.32)





// Consigna: la Secretaría de Turismo de la ciudad de Resistencia habilitó el
// programa "EcoBici" para promover el uso de la bici con fines turísticos. Se
// trata de 6 circuitos para recorrer 14 barrios de la ciudad. Los paseos tienen
// un costo dependiendo de su duración:
//
// (*) Paseos intensivos: superan las 6 hs. Tienen un costo fijo inicial de $1500
//     más un costo variable dependiendo del circuito y cantidad de horas utilizadas.
// (*) Paseos recreativos: duran 6 hs. o menos. Tienen un costo fijo inicial de
//     $1000 más un costo variable dependiendo del circuito y cantidad de horas
//     utilizadas.
//
// Los costos variables de los paseos se encuentran en un arreglo de dos dimensiones
// donde cada posición corresponde al número de circuito y tipo de paseo, respectivamente.
// Se cuenta con un archivo secuencial con la información de todas las bicicletas
// disponibles en la ciudad:
//
// BICICLETAS (ordenado por nro_serie y modelo)
// | Nro_serie | Modelo | Fecha_adquisición | Fecha_ult_mantenimiento | Disponibilidad |
//
// Además se cuenta con información de todas las novedades para una bicicleta.
// Las mismas pueden tratarse de movimientos de préstamos, o bien de una novedad
// respecto a un mantenimiento realizado a la unidad.
//
// NOVEDADES (ordenado por nro_serie, modelo, tipo_novedad y fecha_novedad)
// | Nro_serie | Modelo | Tipo_novedad | Fecha_novedad | Hora_inicio (00..23) | Hora_fin (00..23) | Circuito_nro | ID_usuario |
//
// Tipo de novedad puede ser:
// (*) 1 (Alta de una nueva unidad): cuando se da de alta, la fecha de último
//     mantenimiento debe estar vacía.
// (*) 2 (Préstamo): no afecta los datos de las bicicletas, la secretaría lo
//     utiliza con fines estadísticos.
// (*) 3 (Mantenimiento): implica que deberá actualizarse la fecha de último
//     mantenimiento.
// (*) 4 (Baja de una unidad).
//
// Se pide:
// (a) Desarrollar un algoritmo que permita mantener actualizado el archivo de
//     bicicletas, considerando que pueden existir múltiples novedades para un
//     mismo número de serie. Considerar cualquier tipo de error informando por
//     pantalla al usuario.
// (b) Indicar la cantidad total de préstamos realizados para un circuito
//     particular que ingresa el usuario, y el total recaudado.

ACCION Modelo1_Actualizacion (costos_var: arreglo de [1..6, 1..2] de entero) ES

	Formato_Fecha= REGISTRO
		dia: 1..31
		mes: 1..12
		año: entero(4)
	FINREGISTRO

	Formato_Bicicletas
		Clave= REGISTRO
			Nro_serie: entero(5)
			Modelo: entero(3)
		FINREGISTRO
		Fecha_adquisición: Formato_Fecha
		Fecha_ult_mantenimiento: Formato_Fecha
		Disponibilidad: ("si", "no") // puede ser que sea cantidad
	FINREGISTRO

	regbici, regsal: Formato_Bicicletas
	archbici: Archivo de Formato_Bicicleta Ordenado por Clave
	archsal: Archivo de Formato_Bicicletas

	Formato_Novedades
		Clavem= REGISTRO
			Nro_serie: entero(5)
			Modelo: entero(3)
		FINREGISTRO
		Tipo_novedad: 1..4
		Fecha_novedad: Formato_Fecha
		Hora_inicio: 00..23
		Hora_fin: 00..23
		Circuito_nro: 1..6
		ID_usuario: N(4)
	FINREGISTRO

	regnov: Formato_Novedades
	archnov: Archivo de Formato_Novedades Ordenado por Clavem, Tiplo_novedad, Fecha_novedad

	PROCEDIMIENTO Leer_bici ES
		LEER(archbici, regbici)
		SI FDA(archbici) ENTONCES
			regbici.Clave:= HV
		FINSI
	FINPROCEDIMIENTO

	PROCEDIMIENTO Leer_nov ES
		LEER(archnov, regnov)
		SI FDA(archnov) ENTONCES
			regnov.Clavem:= HV
		FINSI
	FINPROCEDIMIENTO

	PROCEDIMIENTO Chequear_Prestamo ES
		SI regnov.Circuito_nro = circuito ENTONCES
			contprestamos:= contprestamos + 1
			SI (regnov.Hora_fin - regnov.Hora_inicio) > 6 HACER
				recaudo:= 1500
			SINO
				recaudo:= 1000
			FINSI
			totrecaudo:= totrecaudo + recaudo
		FINSI
		
		FINSI

	contprestamos, totrecaudado, recaudo, circuito: entero
	baja: logico
	rescircuito: entero(6)
	





PROCESO

	ABRIR E/(archbici); ABRIR E/(archnov)
	ABRIR S/(archsal)
	Leer_bici(), Leer_nov()

	contprestamos:= 0
	totrecaudo:= 0

	ESC("Ingrese un circuito")
	LEER(circuito)

	MIENTRAS (regbici <> HV) o (regnov <> HV) HACER

		SI regbici.Clave.Nro_serie < regnov.Clavem.Nro_serie ENTONCES
			regsal:= regbici
			GRABAR(archsal, regsal)
			Leer_bici()
		
		SINO

			SI regbici.Clave.Nro_serie = regnov.Clavem.Nro_serie ENTONCES

				regsal:= regbici

				MIENTRAS regbici.Clave.Nro_serie = regnov.Clavem.Nro_serie HACER
					SEGUN regnov.Tipo_Novedad HACER
						= 1: ESC("ERROR: No se puede dar de alta una unidad ya previamente adquirida")
						= 2: Chequear_Prestamo()
						= 3: regsal.Fecha_ult_mantenimiento:= regnov.Fecha_novedad
						= 4: baja:= V 
					FINSEGUN
					Leer_nov()
				FINMIENTRAS

				SI NO baja ENTONCES
					GRABAR(archsal, regsal)
				FINSI

				Leer_bici()

			SINO

				regsal.Clave:= regnov.Clavem
				regsal.Fecha_adquisición:= regnov.Fecha_novedad
				Leer_nov()

				MIENTRAS regsal.Clave.Nro_serie = regnov.Clavem.Nro_serie HACER
					SEGUN regnov.Tipo_Novedad HACER
						= 1: ESC("ERROR: No se puede dar de alta una unidad ya previamente adquirida")
						= 2: ESC("No hay prestamo existente")
						= 3: regsal.Fecha_ult_mantenimiento:= regnov.Fecha_novedad
						= 4: baja:= V 
					FINSEGUN
					Leer_nov()
				FINMIENTRAS

				SI NO baja ENTONCES
					GRABAR(archsal, regsal)
				FINSI

			FINSI

		FINSI

	FINMIENTRAS
	
	ESC("La cantidad total de préstamos realizado para el circuito ", circuito, " es de ", contprestamos)
	ESC("La cantidad total recaudada para el circuito ", circuito, " es de ", totrecaudado)

	CERRAR(archsal); CERRAR(archbici); CERRAR(archnov)

FINACCION
	
	








regsal.Fecha_ult_mantenimiento:= " "


Acción ActualizaciónM2(costos_var: arreglo de [1..6, 1..2] de entero) ES

Ambiente

    Fecha = Registro
        Año: N(4)
        Mes: (1..12)
        Día: (1..31)
    Fin Registro

    Bicicletas = Registro
        Clave = Registro
            NroSerie: N(5)
            Modelo: N(4)
        Fin Registro
        FechaAdquisición: Fecha
        FechaÚltimoMantenimiento: Fecha
    Fin Registro

    ARCH_BIC, ARCH_SAL: archivo de Bicicletas ordenado por Clave 
    REG_BIC, REG_SAL: Bicicletas

    Novedades = Registro
        Clave = Registro
            NroSerie: N(5)
            Modelo: N(4)
        Fin Registro
        TipoNovedad: (1..4)
        FechaNovedad: Fecha
        HoraInicio: (0..23)
        HoraFin: (0..23)
        NroCircuito: N(1)
        IDUsuario: N(5)
    Fin Registro

    ARCH_NOV: archivo de Novedades ordenado por Clave, TipoNovedad y FechaNovedad
    REG_NOV: Novedades

    circ_us, horas_paseo, costo_paseo, cant_circ_us, tot_circ_us: entero
    baja_bicicleta: lógico

    Procedimiento Inicializar() ES
        cant_circ_us := 0
        tot_circ_us := 0
    Fin Procedimiento

    Procedimiento LEER_BIC() ES
        LEER(ARCH_BIC, REG_BIC)
        Si FDA(ARCH_BIC) Entonces
            REG_BIC.Clave := HV
        Fin Si
    Fin Procedimiento

    Procedimiento LEER_NOV() ES
        LEER(ARCH_NOV, REG_NOV)
        Si FDA(ARCH_NOV) Entonces
            REG_NOV.Clave := HV
        Fin Si
    Fin Procedimiento

    Procedimiento ObtenerCircuitoUsuario() ES
        Escribir("Introduzca un circuito: ")
        Leer(circ_us)
    Fin Procedimiento

    Procedimiento ChequearPréstamo() ES
        Si REG_NOV.NroCircuito = circ_us Entonces
            cant_circ_us := cant_circ_us + 1
            horas_paseo := REG_NOV.HoraFin - REG_NOV.HoraInicio

            Si horas_paseo > 6 Entonces
                costo_paseo := 1500 + costos_var[REG_NOV.NroCircuito, 1]
            Contrario
                costo_paseo := 1000 + costos_var[REG_NOV.NroCircuito, 2]
            Fin Si

            tot_circ_us := tot_circ_us + costo_paseo
        Fin Si
    Fin Procedimiento

Proceso

    ABRIR E/ (ARCH_BIC)
    ABRIR E/ (ARCH_NOV)
    ABRIR S/ (ARCH_SAL)

    LEER_BIC()
    LEER_NOV()
    Inicializar()
    ObtenerCircuitoUsuario()

    Mientras REG_BIC.Clave <> HV o REG_NOV.Clave <> HV Hacer
        Si REG_BIC.Clave < REG_NOV.Clave Entonces
            GRABAR(ARCH_SAL, REG_BIC)
            LEER_BIC()
        Contrario
            Si REG_BIC.Clave > REG_NOV.Clave Entonces
                REG_SAL.Clave := REG_NOV.Clave
                REG_SAL.FechaAdquisición := REG_NOV.FechaNovedad
                REG_SAL.FechaÚltimoMantenimiento.Año := ""
                REG_SAL.FechaÚltimoMantenimiento.Mes := ""
                REG_SAL.FechaÚltimoMantenimiento.Día := ""
                LEER_NOV()

                Mientras REG_SAL.Clave = REG_NOV.Clave Hacer
                    Según REG_NOV.TipoNovedad Hacer
                        = 1: Escribir("Error. No se puede volver a dar de alta la bicicleta.")
                        = 2: ChequearPréstamo()
                        = 3: REG_SAL.FechaÚltimoMantenimiento := REG_NOV.FechaNovedad
                        = 4: baja_bicicleta := V
                    Fin Según
                    LEER_NOV()
                Fin Mientras

                Si NO baja_bicicleta Entonces
                    GRABAR(ARCH_SAL, REG_SAL)
                Fin Si
            Contrario
                REG_SAL := REG_BIC

                Mientras REG_SAL.Clave = REG_NOV.Clave Hacer
                    Según REG_NOV.TipoNovedad Hacer
                        = 1: Escribir("Error. No se puede volver a dar de alta la bicicleta.")
                        = 2: ChequearPréstamo()
                        = 3: REG_SAL.FechaÚltimoMantenimiento := REG_NOV.FechaNovedad
                        = 4: baja_bicicleta := V
                    Fin Según
                    LEER_NOV()
                Fin Mientras

                Si NO baja_bicicleta Entonces
                    GRABAR(ARCH_SAL, REG_SAL)
                Fin Si
                LEER_BIC()
            Fin Si
        Fin Si
    Fin Mientras

    Escribir("Para el circuito" circ_us ", indicado por el usuario, se realizaron" cant_circ_us "préstamos.")
    Escribir("Con dichos préstamos se recaudó un total de" tot_circ_us "pesos.")

    CERRAR(ARCH_BIC)
    CERRAR(ARCH_NOV)
    CERRAR(ARCH_SAL)

Fin Acción