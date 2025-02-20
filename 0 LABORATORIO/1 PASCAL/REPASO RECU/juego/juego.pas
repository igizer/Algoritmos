//14) Se pide realizar un juego en donde cada jugador ingresa un número entre 1 y 100, el
//jugador que ingresa el número más grande entre los dos gana la cantidad de punto de (100
//- el número elegido).

//13:00

program juego;
var
jugador1, jugador2, seleccion1, seleccion2, turno: integer;

begin
 
jugador1:= 0;
jugador2:= 0;
turno:=1;

writeln('INSTRUCCIONES:');
writeln('Cada jugador ingresa un número entre 1 y 100,');
writeln('el jugador que ingresa el número más grande entre los dos');
writeln('gana la cantidad de puntos de (100 - el número elegido).');
writeln('EMPECEMOS!');
writeln();

    While (jugador1 < 100) and (jugador2 < 100) do
    begin

        writeln('PUNTUACION');
        writeln('JUGADOR 1: ', jugador1);
        writeln('JUGADOR 2: ', jugador2);
        writeln();

        writeln('turno: ', turno);
        writeln('Jugador1');
        writeln('Escriba un número del 1 al 100');
        readln(seleccion1);

        writeln('Jugador2');
        writeln('Escriba un número del 1 al 100');
        readln(seleccion2);

        if seleccion1>seleccion2 then
        begin
            jugador1:= jugador1 + (100-seleccion1)
        end
        else
        begin
            jugador2:= jugador2 + (100-seleccion2)
        end;

        turno:= turno+1;

        //simulo clear screen
        writeln();
        writeln();
        writeln();
        writeln();
        writeln();
        writeln();
        writeln();
        writeln();
        writeln();
        writeln();

    end;

    if jugador1 > 99 then
    begin
        writeln('Ganador Jugador 1, en un total de ', turno, ' turnos.');
    end
    else  // El else debe ir directamente después del "end" del if
    begin
        writeln('Ganador Jugador 2!, en un total de ', turno, ' turnos.');
    end;
end.