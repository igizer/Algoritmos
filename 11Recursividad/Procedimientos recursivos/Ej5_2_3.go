Dada una lista de nombres ordenada en forma ascendente,
construir un procedimiento recursivo que imprima como salida la misma lista, pero en orden descendente
sin modificar la lista original.

procedimiento descendente(p:puntero a nodo) es
	si p <> nil entonces
		descendente(*p.prox)
		esc(*p.dato)
	finsi
finprocedimiento



// a  b  c

 apila 			desapila
//a b c =  |  c b a 