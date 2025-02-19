Ejercicio 5.2.4

Para convertir un número decimal a binario,
simplemente debe dividirse sucesivas veces por dos (2)
hasta quedarnos con el cociente cero (0).
Todos los restos de las divisiones, tomados en orden inverso,
forman el número binario objetivo. Escribir un procedimiento recursivo que,
recibiendo como parámetro un número entero positivo,
muestre en pantalla el mismo número codificado en binario.

procedimiento binario(n:entero) es
	si n div 2 <> 0 entonces
		binario(n )
	sino
		esc(n mod 2)
	finsi
finprocedimiento






procedimiento tranformar(n: entero) es
	
	
	si n div 2 <> 0 entonces			caso recursivo
		transformar(n div 2)
		esc(n mod 2)
	sino
		esc(n mod 2)					caso base
	fs
fP

-------------------------------------------------

4 en binario 100

4 / 2 => 2 / 2 => 1 / 2  = > 0			divisiones enteras

0			0		1					restos 

-----------------------------------------------------------


n = 4

apila:

transformar(4) 4 div 2 = 2		4 mod 2 = 0
transformar(2) 2 div 2 = 1		2 mod 2 = 0
transformar(1) 1 div 2 = 0		1 mod 2 = 1

desapila

transformar(1) ejecuta el sino : esc(1)
vueve a transformar(2) ejecuta la siguiente instruccion dentro del si esc(n mod 2) = 0
vuelve a transformar(4) ejecuta el esc( n mod 2) = 0

resultado en pantalla

1 0 0