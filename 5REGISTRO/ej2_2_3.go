accion 2.2.1 es 
ambiente 
	fecha= registro
		dia: N(2)
		mes: N(2)
		aa: N(2)
	socios= registro 
		n_socio: N(8)
		n_telefono:N(10)
		apellido_nombre:AN(50)
		carrera:("ISI,IQ,IEM")
		domicilio: AN(50)
		Cantidad_Unidades_Prestadas: N(2)
		campo_f1: fecha // caso hipotético
		campo_fecha2 // caso hipotético
	FINREGISTRO 
	reg: socios
	Socios_arch, arch_s: archivo de socios 
	

Proceso	
	ABRIR e/(Socios_arch) ; LEER(socios,s)
	ABRIR s/(arch_s)
	MIENTRAS no FDA (socios) HACER
		SI reg.carrera = ISI ENTONCES y reg.Cantidad_Unidades_Prestadas> 4 entonces
			ESC(arch_s, reg)
		FINSI
		FINSI
	FIN MIENTRAS 
	CERRAR(socios_arch)
	CERRAR(arch_s)
FINACCION

