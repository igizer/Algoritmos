Program AdivinarJugadores;
uses CRT;

const
    c1 = 'franco armani';
    c25 = 'jeremias ledesma';
    c37 = 'lucas lavagnino';
    c41 = 'santiago beltran';
    c2 = 'federico gatonni';
    c3 = 'ramiro funes mori';
    c4 = 'gonzalo montiel';
    c6 = 'german pezzella';
    c14 = 'leandro gonzalez pirez';
    c16 = 'fabricio bustos';
    c17 = 'paulo diaz';
    c20 = 'milton casco';
    c21 = 'marcos acuna';
    c28 = 'lucas martinez quarta';
    c5 = 'matias kranevitter';
    c7 = 'matias rojas';
    c8 = 'maximiliano meza';
    c10 = 'manuel lanzini';
    c18 = 'gonzalo martinez';
    c23 = 'rodrigo villagra';
    c24 = 'enzo perez';
    c26 = 'ignacio fernandez';
    c29 = 'rodrigo aliendro';
    c30 = 'franco mastantuono';
    c31 = 'santiago simon';
    c34 = 'giuliano galoppo';
    c39 = 'santiago lencinas';
    c9 = 'miguel angel borja';
    c11 = 'facundo colidio';
    c15 = 'sebastian driussi';
    c19 = 'gonzalo tapia';
    c32 = 'agustin ruberto';
    c38 = 'ian subiabre';

var
    respuesta: string;
    puntaje, total: integer;

procedure BorrarLinea;
begin
    GotoXY(1, WhereY - 1);  { Mueve el cursor a la línea anterior }
    ClrEol;                 { Borra la línea actual }
end;

procedure Preguntar(dorsal: integer; jugador: string);
begin
    repeat
        write('Dorsal ', dorsal, ': ');
        readln(respuesta);
        respuesta := LowerCase(respuesta); { Convierte la entrada a minúsculas }
        
        if respuesta = 'nose' then 
        begin
            BorrarLinea;
            writeln('Dorsal ', dorsal, ': ', respuesta, '       || ',puntaje, '/', total, '   era: ', jugador);
            exit; { Sale de la función sin sumar punto }
        end
        else if respuesta <> jugador then
        begin
            BorrarLinea; { Borra la línea del intento incorrecto }
        end;
    until respuesta = jugador;

    puntaje := puntaje + 1;
    BorrarLinea;
    writeln('Dorsal ', dorsal, ': ', respuesta, '                                               >> ',puntaje, '/', total, '   correcto!');
end;

begin
    ClrScr;
    writeln('------- ADIVINE EL JUGADOR DE RIVER SEGUN SU DORSAL -------');
    writeln;
    writeln('cosas a tener en cuenta:');
    writeln('   - escribir nombre completo de los jugadores');
    writeln('   - escribir todo en minuscula');
    writeln('   - en la "enie" esribir la "n"');
    writeln('   - si se ingresa un nombre incorrecto se elimina y reintenta');
    writeln('   - si se desea rendir en un dorsal escribir la palabra "nose"');
    writeln;


    puntaje := 0;
    total := 32;

    writeln('   --- ARQUEROS ---');
    Preguntar(1, c1);
    Preguntar(25, c25);
    Preguntar(37, c37);
    Preguntar(41, c41);

    writeln;
    writeln('   --- DEFENSORES ---');
    Preguntar(2, c2);
    Preguntar(3, c3);
    Preguntar(4, c4);
    Preguntar(6, c6);
    Preguntar(14, c14);
    Preguntar(16, c16);
    Preguntar(17, c17);
    Preguntar(20, c20);
    Preguntar(21, c21);
    Preguntar(28, c28);

    writeln;
    writeln('   --- MEDIOCAMPISTAS ---');
    Preguntar(5, c5);
    Preguntar(7, c7);
    Preguntar(8, c8);
    Preguntar(10, c10);
    Preguntar(18, c18);
    Preguntar(23, c23);
    Preguntar(24, c24);
    Preguntar(26, c26);
    Preguntar(29, c29);
    Preguntar(30, c30);
    Preguntar(31, c31);
    Preguntar(34, c34);
    Preguntar(39, c39);

    writeln;
    writeln('   --- DELANTEROS ---');
    Preguntar(9, c9);
    Preguntar(11, c11);
    Preguntar(15, c15);
    Preguntar(19, c19);
    Preguntar(32, c32);
    Preguntar(38, c38);

    writeln;
    writeln('Tu puntaje final es: ', puntaje, '/', total, ' puntos.');
    writeln('¡Gracias por jugar!');

    readln;
    Delay(60000);
end.