Ejercicio 4.2¶
Se dispone de una lista simplemente encadenada de números enteros,
diseñar un algoritmo que a partir de ella genere otra lista conteniendo los nodos cuyos datos terminan en cero;
dichos elementos, deberán ser eliminados de la lista original.
Se asume que la lista está cargada, y que el algoritmo recibe como parámetro de entrada la dirección del primer elemento.

ACCION Ej4_2 (prim:puntero a nodo) es
	AMBIENTE

		nodo = registro
			dato: entero
			prox: puntero a nodo
		fr

		q, p, ant: puntero a nodo
		prim2, ant2: puntero a nodo

	PROCESO

		p:=prim; prim2:=nil

		si p=nil entonces // primero, veo que existan elementos en la lista
			esc("error, lista vacía")
		sino // si existen elementos en la lista, seguimos
			mientras p<>nil hacer // tratamos los elementos desde el primero hasta que terminemos con el último
				si (p mod 10) = 0 entonces // cumple la condicion de la actividad
					nuevo(q)
					*q.dato:=p.dato; *q.prox:= nil

					// 1- ESCRIBIR EN LA LISTA 2
					si prim2=nil entonces // si es el primero de la lista
						prim2:= q
						ant2:= q
					sino // o sea que no es el primero
						ant2.prox:= q
						ant2:=q
					fs

					// 2- AVANZAMOS
					p:= p.prox

					// 3- ELIMINAR EN LA LISTA 1
					si prim=p entonces // si es el primero de la lista
						disponer(prim)
						prim:= p
					sino // o sea que no es el primero
						disponer(ant.prox)
						ant.prox:= p
					fs
				sino
					ant:= p
					p:= p.prox
				fs
			fm
		fs

FINACCION








accion Lista_4.2 (Prim: puntero a nodo) ES

	AMBIENTE

	Nodo = registro
		dato: Entero
		prox: Puntero a Nodo
	FR	
	p,ant,q: puntero a Nodo
	prim2: puntero a nodo
	ant2: puntero a nodo
	PROCESO

	p:= prim
	prim2:= nil
   	si  (P = nil) entonces
		esc("error lista vacia")
   	sino
		mientras (p<>nil) hacer

			si (*p.dato mod 10 = 0) entonces
				nuevo(q); *q.dato:= *p.dato;
				*q.prox:= nil;
		
				Si prim2 = nil entonces
					ant2:= q
					prim2:= q;
				sino

					ant2.prox:= q
					ant2:=q
				fs

				si (p= prim) entonces
					p:= *P.prox
					disponer(prim)
					prim:= p
				sino
					p:=*p.prox
					disponer(*ant.prox)
					ant.prox:=p
				fs
			sino
				ant:=p
				p:=*p.prox
			fs
		fm
 	fs
Fin_accion.