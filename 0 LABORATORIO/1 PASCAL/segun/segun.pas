program segun;
Var 
nota : integer;
begin 
 writeln('Ingrese la nota del alumno: ');
 read(nota);
 case nota of
 0..3: writeln('Mal');
 4, 5: writeln('Insuficiente');
 6..7: writeln('Bien');
 8, 9: writeln('Sobresaliente');
 10: writeln('Perfecto');
 else
 writeln('ERROR, numero fuera de rango.');
 end;
end.
