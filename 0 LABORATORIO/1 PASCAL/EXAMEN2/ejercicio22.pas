program ejercicio2;

uses math;

var   
    sec: file of Char;
    v, tipo: Char;
    total, plata, compras, i, p: integer;

function convertir(x: char): integer;
begin
    case x of
        '0': convertir := 0;
        '1': convertir := 1;
        '2': convertir := 2;
        '3': convertir := 3;
        '4': convertir := 4;
        '5': convertir := 5;
        '6': convertir := 6;
        '7': convertir := 7;
        '8': convertir := 8;
        '9': convertir := 9;
    end;
end;

begin
    Assign(sec, 'C:\Users\xpoke\Downloads\Actividades en clase\0 LABORATORIO\1 PASCAL\EXAMEN2\LD03.txt');

    {$I-}
    Reset(sec);
    {$I+}

    if IOResult <> 0 then
    begin
        writeln('Error: el archivo existe?');
        halt(2);
    end;

    total := 0;
    compras := 0;
    plata := 0;

    while not EoF(sec) do
    begin
        read(sec, v);
        if EoF(sec) then break;
        
        read(sec, v);
        if EoF(sec) then break;
        tipo := v;
        
        read(sec, v);
        if EoF(sec) then break;
        compras := convertir(v);

        for i := 1 to compras do
        begin
            for p := 2 downto 0 do
            begin
                if EoF(sec) then break;
                read(sec, v);
                if tipo = 'J' then
                    plata := plata + convertir(v) * round(Power(10, p));
            end;
        end;
        total := total + plata;
        plata := 0;
    end;

    writeln('El importe total vendido a clientes de tipo jurídico es de: ', total);

    close(sec);
end.