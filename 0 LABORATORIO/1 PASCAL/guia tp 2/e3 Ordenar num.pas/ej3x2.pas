program ej3x2;
var
    x1, x2, x3: Integer;

procedure intercambio(a, b, c: integer);
begin
    if (a < b) and (a < c) then
    begin
        if b < c then
            writeln('El orden es: ', a, ', ', b, ' y ', c)
        else
            writeln('El orden es: ', a, ', ', c, ' y ', b)
    end;

    if (b < a) and (b < c) then
    begin
        if a < c then
            writeln('El orden es: ', b, ', ', a, ' y ', c)
        else
            writeln('El orden es: ', b, ', ', c, ' y ', a)
    end;

    if (c < a) and (c < b) then
    begin
        if a < b then
            writeln('El orden es: ', c, ', ', a, ' y ', b)
        else
            writeln('El orden es: ', c, ', ', b, ' y ', a)
    end;
end;

begin
    writeln('A continuación, escriba el primer número');
    readln(x1);
    writeln('A continuación, escriba el segundo número');
    readln(x2);
    writeln('A continuación, escriba el tercer número');
    readln(x3);
    intercambio(x1, x2, x3);
end.