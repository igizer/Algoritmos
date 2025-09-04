//2) Contar todas las palabras de una secuencia
//Para poder ejecutar este código es necesario crear un archivo (con nombre: archivo.txt) en
//el mismo lugar donde se encuentra el código del programa.
//Colocar en el archivo el siguiente texto:
//Lorem ipsum dolor sit amet, consectetur adipiscing elit. Mauris eu eros odio. Nunc venenatis
//nibh et dolor commodo, eu dictum libero malesuada. Nullam ac urna nisl. Quisque condimentum
//augue sagittis.
//Código:

program ej2gtp3;

const
signos = [' ', ',', '.', '?', '¿', '¡', '!'];

var
sec: text;
v: char;
error, palabras: integer;

begin

Assign(sec, 'C:\Codigo\Algoritmos\0 LABORATORIO\1 PASCAL\Secuencias\archivo.txt');

{$I-}
Reset(sec);
error:= IOResult;
{$I+}

if error <> 0 then
    begin
    writeln('Error: el archivo archivo.txt exite?');
    halt(2);
    end;

palabras:=0;

while not EOF(sec) do
    begin

        repeat
            read(sec,v)
        until not (v in signos);

        repeat
            read(sec,v)
        until (v in signos) or EOF(sec);

        palabras:= palabras+1;
        
    end;

    writeln('La cantidad de palabras son: ', palabras);
    close(sec);

end.