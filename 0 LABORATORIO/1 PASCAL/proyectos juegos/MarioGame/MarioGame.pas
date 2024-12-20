
Program MarioGame;
Const Saludo = 'Bienvenido a MarioGame!';
Const Version = 1.0;
Var Play : char;
Var Tiempo : Integer;
Var Vidas : Integer;
Var Band : Boolean;
var Dinero, i, x, z, y : integer;
var seleccion : string;
var Facil : integer;
var cantidad : integer;
var Normal : integer;
var Dificil : integer;
Begin
    WriteLn('¿Dar play? (S o N)');
    ReadLn(Play);
    if Play = 'S' Then
        begin
        Tiempo := 280;
        Vidas := 3;
        WriteLn(Saludo);
        WriteLn('Version ', Version:0:1);
        WriteLn('Te quedan ', Tiempo, ' Segundos');
        WriteLn('Te quedan ', Vidas, ' Vidas');
        for x := 1 to 5 do
            begin
                for y := 1 to 5 do
                        if (x = 4) and (y = 3) then
                            write('  O   ')
                        else
                            write( '(', x, ',', y, ') ');
                writeln();
            end;
        end;
    if Play = 'N' Then
        begin
        WriteLn('Ok, goodbye!');
        end;
End.