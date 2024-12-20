// Ejercicio_2_1_2 (chequeado)
// "Dada una secuencia de letras que finaliza con una
// marca "*" contar cuantas letras "A" hay en la secuencia"
ACCION Ejercicio_2_1_2 Es
AMBIENTE
    secuencia: secuencia de caracter
    c: caracter
    cont: entero
PROCESO
    ARRANCAR(secuencia)
    AVANZAR(secuencia,c)
    cont:= 0
    Mientras c<> "*" Hacer
        Si c= "A" Entonces
            cont:= cont+ 1
        Fin Si
        AVZ(secuencia,c)
    Fin Mientras
    ESCRIBIR("La cantidad de letras A en la secuencia son",cont)
FINACCION

