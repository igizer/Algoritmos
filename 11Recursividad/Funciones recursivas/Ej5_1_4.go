Construir un algoritmo recursivo que permita determinar si los dígitos de un número n dado son todos pares.

funcion digitos(n:entero):entero es
	si (n mod 2 <> 0) entonces
		digitos:= falso
	sino
		si n=0 entonces//caso base
			digitos:= verdadero
		sino
			digitos:= digitos(n div 10)//caso recursivo

		finsi
	finsi
finfuncion
