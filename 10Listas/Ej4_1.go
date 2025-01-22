Ejercicio 4.1¶
Diseñar un algoritmo para acceder, eliminar o insertar el k-ésimo elemento de una lista (siendo k una posición dada).
Si la lista está vacía o si el valor de k esta fuera del rango del índice de la lista, invocar al procedimiento de ERROR.
En cualquier otro caso, efectuar el procedimiento solicitado.

ACCION Ej4_1 ES (prim: puntero a nodo) ES
	AMBIENTE

		Nodo = REGISTRO
			Dato: Entero
			Prox: Puntero a nodo
		FINREGISTRO;

		q,P, ant: Puntero a Nodo;
		k: Entero;
		posicion
		op:(1,2,3)

		PROCEDIMIENTO Error() ES	
			ESC("Error la posicion está fuera de rango");
		FINPROCEDIMIENTO


		PROCEDIMIENTO Insertar() ES
			nuevo(q);
			ESC("ingrese el valor para la posicion de k");
			LEER(*q.dato);
			SI prim = nil ENTONCES
				prim:= q				//lista vacia
				q.prox:= nil;
			SINO
				p:= prim ;ant:= nil
				MIENTRAS p<>nil y k>1 HACER
					ant:= p
					p:= *p.prox 		
					k:= k-1;
				FINMIENTRAS;
				SI P= prim ENTONCES		//primer lugar
					prim:= q
					*q.prox:=p;
				SINO
					SI p=nil ENTONCES	//ultimo lugar o medio
						*ant.prox:= q
						*q.prox:= p;	
					SI
				FINSI
			FINSI
		FINPROCEDIMIENTO


		PROCEDIMIENTO Acceder() ES
			P:= prim
			SI (prim = nil) ENTONCES
				ESC("lista vacia");
			SINO
			MIENTRAS (P<>nil) y (K>1) HACER
				p:= *p.prox
				K:= K - 1;	
			FINMIENTRAS;
			SI p = nil ENTONCES
				Error();
			SINO
				ESC("el valor de la posicion", k,"es", *p.dato);
			FINSI;
		FINPROCEDIMIENTO

		PROCEDIMIENTO Eliminar()
			SI prim= nil ENTONCES
				ESC("Error lista vacia");
			SINO
				p:=prim, ant:=nil
				MIENTRAS p<>nil y k>1 HACER
					ant:=p
					p:= *p.prox
					k:= k-1
				FINMIENTRAS
				SI P = Nil
					Error();
				SINO
					SI p= prim ENTONCES	
						Prim:=nil
					SINO
						*ant.prox:= *p.prox
					FINSI
				FINSI
				Disponer(p);
			FINSI
		FINPROCEDIMIENTO

	PROCESO

	ESC("Ingrese la posicion K"); leer(k);
	ESC("Elija la que operacion quiere realizar con la posicion:")
	ESC("1- acceder , 2-Insertar , 3-Eliminar:")Leer(op);
	SEGUN op HACER
		1: Acceder();
		2: Insertar();
		3: Eliminar();
	FINSEGUN;