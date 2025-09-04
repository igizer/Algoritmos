Ejercicio 4.1¶
Diseñar un algoritmo para acceder, eliminar o insertar el k-ésimo elemento de una lista (siendo k una posición dada).
Si la lista está vacía o si el valor de k esta fuera del rango del índice de la lista, invocar al procedimiento de ERROR.
En cualquier otro caso, efectuar el procedimiento solicitado.

ACCION ej4_1 (prim: puntero a nodo) es
	ambiente
		q, p, ant: puntero a nodo
		k: entero
		op(1, 2, 3)

		nodo = registro
			dato: entero
			prox: puntero a nodo
		fr

		procedimiento error() es
			esc("Error la posicion está fuera de rango")
		finprocedimiento

		procedimiento insertar()es
			nuevo(q)
			esc("ingrese el valor para la posicion de k")
			leer(*q.dato)
			si prim = nil entonces
				prim:= q
				*q.prox:= nil
			finsi
			p:=prim; ant:=nil
			mientras p<>nil ^ k>1 hacer
				ant:=p
				p:=*p.prox
				k:=k-1
			fm
			si p=prim entonces // primer lugar
				prim:= q
				*q.prox:= p
			sino
				si p=nil entonces// ultimo lugar o en medio
					*ant.prox:= q
					*q.prox:= p
				fs
			fs
		fp

		procedimiento acceder() es
			prim:= p
			si prim=nil entonces
				esc("error, lista vacía")
			fs
			mientras p<>nil ^ k>1 hacer
				p:= p.prox
				k:= k-1
			fm
			si p=nil entonces
				error()
			sino
				esc("el valor de la posicion", k,"es", *p.dato)
			fs
		fp

		procedimiento eliminar() es
			prim:= p; ant:= nil
			si prim=nil entonces
				esc("lista vacía")
			fs
			mientras p<>nil ^ k>1 hacer
				ant:= p
				p:= p.prox
				k:= k-1
			fm
			si p=nil entonces
				error()
			sino
				si p=prim entonces
					prim:=nil
				sino
					*ant.prox:= p.prox
				fs
			fs
			disponer(p)
		fp
	
	PROCESO

		esc("Escriba a continuación la posición k de la lista"); leer(k)
		esc("Escriba a continuación la operación a realizar")
		esc("1- Insertar")
		esc("2- Acceder")
		esc("3- Eliminar")
		leer(op)
		segun op hacer
			1: insertar()
			2: acceder()
			3: eliminar()
		fs
	
FINACCION






		

Accion Ejer_4.1 ES (prim: puntero a nodo) Es
	ambiente

	Nodo = registro
		Dato: Entero
		Prox: Puntero a nodo
	fr;

	q,P, ant: Puntero a Nodo;
	k: Entero;
	posicion
	op:(1,2,3)

	Procedimiento Error() Es	
		Esc("Error la posicion esta fuera de rango");
	FS


	Procedimiento Insertar()Es
		nuevo(q);
		esc("ingrese el valor para la posicion de k");
		leer(*q.dato);
		si prim = nil entonces
			prim:= q				//lista vacia
			q.prox:= nil;
		sino
			p:= prim ;ant:= nil
			Mientras p<>nil y k>1 Hacer
				ant:= p
				p:= *p.prox 		
				k:= k-1;
			fm;
			si P= prim entonces		//primer lugar
				prim:= q
				*q.prox:=p;
			sino
				si p=nil entonces	//ultimo lugar o medio
					*ant.prox:= q
					*q.prox:= p;	
				fs
			fs
		fs
	Fin_procedimiento


	Procedimiento Acceder()Es
		P:= prim
		si (prim = nil) entonces
			esc("lista vacia");
		sino
		Mientras (P<>nil) y (K>1) Hacer
			p:= *p.prox
			K:= K - 1;	
		FM;
		si p = nil entonces
			Error();
		sino
			esc("el valor de la posicion", k,"es", *p.dato);
		fs;
	Fin_procedimiento

	Procedimiento Eliminar()
		SI prim= nil entonces
			esc("Error lista vacia");
		sino
			p:=prim, ant:=nil
			mientras p<>nil y k>1 hacer
				ant:=p
				p:= *p.prox
				k:= k-1
			fm
			si P = Nil
				Error();
			sino
				si p= prim entonces	
					Prim:=nil
				sino
					*ant.prox:= *p.prox
				fs
			fs
			Disponer(p);
		fs
	Fin_procedimiento


	proceso

	Esc("Ingrese la posicion K"); leer(k);
	Esc("Elija la que operacion quiere realizar con la posicion:")
	Esc("1- acceder , 2-Insertar , 3-Eliminar:")Leer(op);
	Segun op hacer
		1: Acceder();
		2: Insertar();
		3: Eliminar();
	Fs;
