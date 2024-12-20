//3) Escribir un algoritmo donde se ingresan 3 números distintos y se deben ordenar de
//menor a mayor, e imprimir por pantalla.
program ej3;
var
    x1, x2, x3 : integer;
procedure intercambio(a, b, c: integer);
    begin
        if (a < b) and (a < c) then //a<b<c y a<c<b
            begin
                if b < c then
                    writeln('el orden es: ', a, ', ', b, ' y ', c)
                else
                    writeln('el orden es: ', a, ', ', c, ' y ', b)
            end;
        
        if (a > b) and (b < c) then //b<a<c y b<c<a
            begin
                if a < c then
                    writeln('el orden es: ', b, ', ', a, ' y ', c)
                else
                    writeln('el orden es: ', b, ', ', c, ' y ', a)
            end;
        
        if (a > c) and (c < b) then //c<a<b y c<b<a
            begin
                if b < a then
                    writeln('el orden es: ', c, ', ', b, ' y ', a)
                else
                    writeln('el orden es: ', c, ', ', a, ' y ', b)
            end;
    end;
begin
    writeln(' a continuacion escribe el primer numero');
    readln(x1);
    writeln(' a continuacion escribe el segundo numero');
    readln(x2);
    writeln(' a continuacion escribe el tercer numero');
    readln(x3);
    intercambio(x1, x2, x3);
end.