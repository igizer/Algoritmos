// "La secuencia de socios del problema anterior tiene el inconveniente de que los
// números están ordenados pero no son correlativos. Se desea generar una secuencia
// que contenga los números de socios que no figuran en la secuencia de socios.""

accion r2_1_5 es
	ambiente
		sec, sal: secuencia de entero
		act, ant: entero// ventana actual y ventana anterior, socios es un num que va a escribir los numeros de socios faltantes
	
	proceso
		arr(sec)
		avz(sec, act)
		ant:= act
		avz(sec, act)
		crear(sal)

		mientras nfds(sec) hacer
			si (act-ant)>1 entonces
				repetir
					ant:=ant+1
					grabar(sal, ant)
				hasta (act-ant)=1
			finsi
			ant:= act
			avz(sec,act)
		fin mientras
		cerrar(sec); cerrar(sal);
		esc("secuencia de salida creada con exito.")

finaccion


----------------------SINO TAMBIEN FUNCIONA Y ERA LA IDEA DE HACER:------------------

accion r2_1_5 es
	ambiente
		sec, sal: secuencia de entero
		act, ant, i: entero
proceso
		arr(sec)
		avz(sec, act)
		ant:= act
		avz(sec, act)
		crear(sal)

	mientras nfds(sec) hacer
			si (act-ant)<> 1 hacer
				para i:= 1 hasta (act-ant-1)
					ant:=ant+1
					grabar(sal, ant)
				finpara
			finsi
			ant:= act
			avz(sec,act)
		fin mientras
		cerrar(sec); cerrar(sal);
		esc("secuencia de salida creada con exito.")

finaccion