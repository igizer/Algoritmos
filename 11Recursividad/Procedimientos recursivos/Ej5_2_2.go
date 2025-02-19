Leer una palabra (una cadena de caracteres) y la cantidad de caracteres y generar su palíndromo.
El palíndromo de “Venezuela” es “aleuzeneV”.

//si en el procedimiento tomamos el arreglo y la cantidad de letras:

Procedimiento palindromo(a: arreglo[1..n]:caracter, n:entero) ES
    si n=1 entonces
        esc(a[n])
    sino
        esc(a[n])
        palindromo(a, n-1)
    finsi
finprocedimiento



//sino si antes de llamar al procedimiento ya tenemos la variable i:=1 y el valor de n en una variable:

Procedimiento Cadena(arreglo[1..N]:caracter, i:entero) ES 
    SI I=N ENTONCES 
        Esc(arreglo[i]) // U
    SINO 
        Cadena(arreglo, i+1); //P E R U 
        Esc(arreglo[i]) // R E P 
    FSI
F-Proc



P E R U 
1 2 3 4