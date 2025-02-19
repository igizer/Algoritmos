Dados dos números enteros, divídalos (división entera) y muestre el resultado, usando sólo la operación resta.
//el de la guia estaba mal, mi respuesta bien que es esta
funcion division(a,b:entero):entero es

	ambiente
		contador:entero
	
	proceso
		si (a < b) entonces
			division:= 0
		sino
			division:= 1 + division(a-b, b)
			contador:= contador + 1
		finsi
finfuncion
	

		4				3				2			1					0
divisino(20,5)= division(15,5) = division(10,5) =division (5,5) =division (0,5) 


	//caso recursivo

	divisor			5		b
	dividendo		20 		a

	contador


	20