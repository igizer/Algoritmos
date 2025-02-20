//2) Contar todas las palabras de una secuencia
//Para poder ejecutar este código es necesario crear un archivo (con nombre: archivo.txt) en
//el mismo lugar donde se encuentra el código del programa.
//Colocar en el archivo el siguiente texto:
//Lorem ipsum dolor sit amet, consectetur adipiscing elit. Mauris eu eros odio. Nunc venenatis
//nibh et dolor commodo, eu dictum libero malesuada. Nullam ac urna nisl. Quisque condimentum
//augue sagittis.
//Código:

program intento2;
var
sec: file of char;
v: Char;
error, palabras: integer;

BEGIN

    Assign(sec, 'C:\Codigo\Algoritmos\0 LABORATORIO\1 PASCAL\Secuencias\archivo.txt');
    {$I-}
    reset(sec);
    error:= IOResult;
    {$I+}
    if error <> 0 then
    begin
        writeln('error: el archivo archivo.txt exite?');
        halt(2);
    end;

    read(sec, v);
    palabras:= 0;

    while not EOF(sec) do
    begin
        while v<>'.' do
        begin
            repeat
                read(sec,v);
            until v=' ';
            palabras:= palabras + 1;
        end;
        read(sec, v);
    end;

    WriteLn('La cantidad de palabras en el texto son de: ', palabras);
    close(sec);

END.