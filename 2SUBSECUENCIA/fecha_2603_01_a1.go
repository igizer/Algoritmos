//Finalizando el ciclo lectivo 2023 de algoritmos,
//se necesitan estadísticas sobre cuántos aprobaron de manera directa la materia,
//aprobaron la cursada (regularizar), o quedaron libres. para definir la situación
//de cada estudiante se debe contabilizar la cantidad de aprobados.
//De seis instancias evaluativas, se define:
//Aprobación directa (promoción) 6 instancias aprobadas
//Aprobación de cursada (regular) 4 instancias aprobadas
//Libre menos de 4 instancias aprobadas

ACCION ejercicio_2_1 ES
	AMBIENTE
		entrada, salida: secuencia de caracter
		v: caracter
		i, promo, reg, lib: entero
		numeros= (0,1,2,3,4,5,6, 7, 8, 9)

	PROCESO
		ARR(entrada)
		AVZ(entrada, v)
		CREAR(salida)
		promo:=0
		reg:=0
		lib:=0
		MIENTRAS NFDS(entrada) HACER
			MIENTRAS v <> "#" HACER
				MIENTRAS v <> numeros HACER
					GRAABAR(salida, v)
					AVZ(entrada, v)
				FINMIENTRAS
				PARA i:= 1 HASTA 5 HACER
					GRABAR(salida, v)
					AVZ(entrada, v)
				FINPARA
				SEGUN v HACER


