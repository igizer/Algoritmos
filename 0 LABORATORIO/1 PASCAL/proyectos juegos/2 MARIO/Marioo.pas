program Marioo;

Uses Crt;

var x, y, xply, yply : integer;
var tecla : char;
begin
    xply := 3;
    yply := 10;
    while true do
        begin
            for x := 1 to 6 do
                begin
                    for y := 1 to 19 do
                        if (x = xply) AND (y = yply) then
                            write('O ')
                        else
                            write('* ');
                    writeln();
                end;
            writeln();
            writeln('Press a key');
            tecla:= ReadKey;
            if tecla = 'w' then
                xply := xply-1;
            if tecla = 's' then
                xply := xply + 1;
            if tecla = 'd' then
                yply := yply + 1;
            if tecla = 'a' then
                yply := yply - 1;
            ClrScr;
        end;
end.