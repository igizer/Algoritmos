Dados dos números: a y b. Calcule la potencia ab, usando sólo multiplicaciones sucesivas.

funcion potencia(a,b:entero):entero es
	si b=1 entonces
		potencia:= a
	sino
		potencia:= a * potencia(a,b-1)
	finsi
finfuncion


	prueba de escritorio:


	potencia(2,5)=2*potencia(2,4)
	potencia(2,5)=2*2*potencia(2,3)
	potencia(2,5)=2*2*2*potencia(2,2)
	potencia(2,5)=2*2*2*2*potencia(2,1)
	potencia(2,5)=2*2*2*2*2
	potencia(2,5)=32
