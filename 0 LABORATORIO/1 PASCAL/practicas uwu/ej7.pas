// En un colegio primario se necesita realizar un algoritmo que en vez de colocar las notas 
// de los alumnos en forma numérica, la misma debe ser con una frase, el algoritmo debe 
// traducir las notas del 0 al 10 en las siguientes frases:
program Ej7;
Var 
nota : integer;
begin 
 nota := 1;
 repeat
 if not(nota in [1..10]) then
 writeln('ERROR, numero fuera de rango. vuelva a cargar la nota');
 writeln('Ingrese la nota del alumno: ');
 read(nota);
 until nota in [1..10];
 case nota of
 0..3: writeln('Mal');
 4, 5: writeln('Insuficiente');
 6..7: writeln('Bien');
 8, 9: writeln('Sobresaliente');
 10: writeln('Perfecto'); 
 end;
end.
