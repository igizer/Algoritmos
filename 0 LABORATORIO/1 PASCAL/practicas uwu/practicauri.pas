program Practicauri;
var
    edad, i, dinero : Integer;
    peso : real;
    nombre : string;
    siono : Char;
    band : boolean;
const
    programa = '"Este eres tu!"';
    version = 1;
begin
    dinero := 3;
    for i := 0 to 30 do
        begin
            writeln()
        end;
    while dinero <> 0 do
        begin
            writeln('Ejecutar codigo (S/N)');
            readln(siono);
            if siono = 'S' then
                begin
                    dinero := dinero - 1;
                    writeln('tu dinero restante es ', dinero);
                    Writeln('Bienvenidos a ', programa, ' esta es la version ', version);
                    Writeln('Escribe tu nombre');
                    Readln(nombre);
                    Writeln('Escribe tu edad');
                    readln(edad);
                    writeln('escribe tu peso');
                    readln(peso);
                    writeln('HOLA ', nombre, ' tenes ', edad, ' años, y pesas ', peso:0:3);
                end
            else if siono = 'N' then
                begin
                    writeln('entonce jodete nene');
                    writeln('encima pelotudo');
                end
            else
                writeln('caracter introducido incorrecto')
        end
end.