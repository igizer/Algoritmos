Calcular el factorial de un número positivo n. Tener en cuenta la definición matemática de n!:

funcion factorial (n: entero) : entero es
	si n=0 entonces
		factorial:= 1
	sino
		factorial:= n * factorial(n-1)
	finsi
finfuncion