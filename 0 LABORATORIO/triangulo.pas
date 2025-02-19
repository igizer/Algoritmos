program triangulo;
var
niveles, i, j: integer;

BEGIN
    Writeln('Ingresa una cantidad de niveles para la piramide');
    Readln(niveles);
    For i:=1 to niveles do
        begin
            for j:=1 to i do
                begin
                    write('*');
                end;
            writeln;
        end;
END.
