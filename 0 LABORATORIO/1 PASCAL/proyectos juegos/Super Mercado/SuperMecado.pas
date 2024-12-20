program Super_Mercado;
    var
    Dinero, i, x, z, y, cantidad, Facil, Normal, Dificil : integer; 
    Play: char;
    band : Boolean;
    seleccion : string;
begin
    for x := 1 to 5 do
        begin
            for y := 1 to 5 do
                begin
                    write( '(', x, ',', y, ') ');
                end;
            writeln();
        end;
    writeln();
    writeln();
    writeln();
    writeln();
    writeln();
    WriteLn('¿Dar Play? (S o N)');
    ReadLn(Play);
    if Play = 'S' Then
        begin
            Facil := 1500;
            Normal := 1000;
            Dificil := 500;
            cantidad := 0;
            WriteLn('');
            WriteLn('');
            WriteLn('');
            WriteLn('');
            WriteLn('');
            WriteLn('Bienvenido al super!');
            WriteLn('Instrucciones:');
            WriteLn('');
            WriteLn('Elije un modo de juego');
            WriteLn('1- Facil');
            WriteLn('2- Normal');
            WriteLn('3- Dificil');
            ReadLn(seleccion);
            if seleccion = '1' Then
                begin
                    seleccion := 'Facil';
                    Dinero := Facil;
                end;
            if seleccion = '2' Then
                begin
                    seleccion := 'Normal';
                    Dinero := Normal;
                end;
            if seleccion = '3' Then
                begin
                    seleccion := 'Dificil';
                    Dinero := Dificil;
                end;
            WriteLn('');
            WriteLn('');
            WriteLn('');
            WriteLn('Tu modo de juego seleccionado es ', seleccion);
            WriteLn('Dispones de $', Dinero);
            while seleccion <> '0' do
                begin
                    WriteLn('Que queres hacer a continuacion?');
                    WriteLn('1- Comprar');
                    WriteLn('2- Vender');
                    WriteLn('3- Apostar');
                    ReadLn(seleccion);
                    if seleccion = '1' then
                        begin
                            if dinero > 49 then
                                begin
                                    cantidad := cantidad + 1;
                                    dinero := dinero - 50;
                                    Writeln('ahora tienes $', dinero);
                                    Writeln('ahora tienes ', cantidad, ' unidades');
                                end;
                        end;
                    if seleccion = '2' then
                        begin
                            if cantidad > 0 then
                                begin
                                    cantidad := cantidad - 1;
                                    dinero := dinero + 60;
                                    Writeln('ahora tienes $', dinero);
                                    Writeln('ahora tienes ', cantidad, ' unidades');
                                end;
                        end;
                    if seleccion = '3' then
                        begin
                            dinero := 0;
                            Writeln('ahora tienes $', dinero);
                            Writeln('pero aun tienes ', cantidad, ' unidades');
                        end;
                end;
        end;
    if Play = 'N' Then
        WriteLn('Adios y hasta la proxima!');
end.