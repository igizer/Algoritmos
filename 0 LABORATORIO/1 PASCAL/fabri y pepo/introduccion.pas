program Introduccion;
var
    edad : Integer;
    peso : Real;
    nombre : string;
    VoF : char;
const
    juego = '"Este sos vos"';
    version = 1;
begin
    writeln('Hola, bienvenido a ', juego, ' version: ', version);
    writeln('Cual es tu nombre?');
    readln(nombre);
    writeln('Ingrese su edad');
    readln(edad);
    writeln('ingrese su peso');
    readln(peso);
    Writeln('Hola ', nombre);
    writeln('Tenes ', edad, ' años, y pesas ', peso:0:2);
end.

