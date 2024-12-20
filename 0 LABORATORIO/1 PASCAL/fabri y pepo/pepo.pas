//3) Escribir un algoritmo donde se ingresan 3 números distintos y se deben ordenar de
//menor a mayor, e imprimir por pantalla.
program Ej3;
var 
    num1, num2, num3 : integer;
procedure intercambio(a,b,c: integer);
    begin
        if (a<b) and (a<c) then 
            begin
                if b<c then
                    writeln('el orden es ', a, ',' b, 'y', c)
                else
                    writeln('el orden es ', a, ',' c, 'y', b)
            end;
        if (b<a) and (b<c) then 
            begin
                if a<c then
                    writeln('el orden es ', b, ',' a,'y', c)
                else
                    writeln('el orden es', b, ',' c, 'y', a)
            end;
        if (c<a) and (c<b) then 
            begin
                if a<b then
                    writeln('el orden es ', c, ',' a, 'y', b)
                else
                    writeln('el orden es ', c, ',' b, 'y', a)
            end;
    end;
begin
    writeln('ingrese el primer numero ');
    readln(num1);
    writeln('ingrese el segundo numero ');
    readln(num2);
    writeln('ingrese el tercer numer ');
    readln(num3);
    intercambio(num1,num2,num3);
end.