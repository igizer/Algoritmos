Dado un vector de 10 números enteros, calcular la suma de sus elementos

funcion suma(a: arreglo [1..10] de entero, i : entero): entero
	si i = 1 entonces
		suma:= a[i]
	sino
		suma:= a[i] + suma(a, i-1)
	finsi
finfuncion