//2) Contar todas las palabras de una secuencia
//Para poder ejecutar este código es necesario crear un archivo (con nombre: archivo.txt) en
//el mismo lugar donde se encuentra el código del programa.
//3)Mostrar el promedio de palabras que hay por oración.
//Se requiere ocupar el mismo archivo dado anteriormente
//Colocar en el archivo el siguiente texto:
//Lorem ipsum dolor sit amet, consectetur adipiscing elit. Mauris eu eros odio. Nunc venenatis
//nibh et dolor commodo, eu dictum libero malesuada. Nullam ac urna nisl. Quisque condimentum
//augue sagittis.
//Código:

program ej2secuencia;
var
    sec: file of char;
    v: char;
    palabras, error, oracion: integer;

BEGIN
    Assign(sec, 'C:\Users\xpoke\Downloads\Actividades en clase\0 LABORATORIO\1 PASCAL\guia tp3\ej2secuencia\archivo.txt');

    {$I-}
    Reset(sec);
    {$I+}

    if IOResult <> 0 then
        begin
            Writeln('Error: el archivo archivo.txt exite?');
            halt(2);
        end;
    
    read(sec, v);

    palabras:=0;
    oracion:=0;

    While Not EoF(sec) do
        begin
            while Not Eof(sec) and (v <> '.') do
                begin
                    While v = ' ' do
                        begin
                            read(sec, v);
                        end;

                    palabras:= palabras + 1;
                    
                    while (v <> ' ') and (v <> '.') and Not EoF(sec) do
                        begin
                        read(sec,v);
                        end;

                end;
            oracion:= oracion + 1
        end;

    Writeln('El promedio de palabras por oracion es de:  ', palabras div oracion);
    close(sec);

END.