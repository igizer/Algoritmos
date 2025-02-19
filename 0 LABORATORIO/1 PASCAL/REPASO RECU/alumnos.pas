// En un colegio primario se necesita realizar un algoritmo que en vez de colocar las notas 
// de los alumnos en forma numérica, la misma debe ser con una frase, el algoritmo debe 
// traducir las notas del 0 al 10 en las siguientes frases:
program alumnos;
var
nota: integer;
begin
    writeln('escribe la nota del alumno');
    read(nota);
    case nota of
        0..3: writeln('pibe como vas a sacar esa nota de mierda pelotudo');
        4..5: writeln('ahí noma trankate');
        6..7: writeln('uu ahí noma safaste');
        8..10: writeln('WEALE');
    end;
end.