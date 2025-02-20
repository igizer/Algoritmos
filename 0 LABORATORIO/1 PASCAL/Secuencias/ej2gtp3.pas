program ej2gtp3;

const
    signos = [' ', ',', '.', '?', '¿', '¡', '!'];  // Caracteres que separan palabras

var
    sec: Text;
    v: char;
    error, palabras: integer;
    dentroPalabra: boolean;

begin
    Assign(sec, 'C:\Codigo\Algoritmos\0 LABORATORIO\1 PASCAL\Secuencias\archivo.txt');

    {$I-}
    Reset(sec);
    error := IOResult;
    {$I+}

    if error <> 0 then
    begin
        writeln('Error: ¿el archivo archivo.txt existe?');
        halt(2);
    end;

    palabras := 0;
    dentroPalabra := false;

    // Mientras no sea el fin de archivo
    while not EOF(sec) do
    begin
        Read(sec, v);  // Leer un carácter

        // Si el carácter no es un separador (es parte de una palabra)
        if not (v in signos) then
        begin
            // Si no estábamos dentro de una palabra, aumentamos el contador
            if not dentroPalabra then
            begin
                palabras := palabras + 1;
                dentroPalabra := true;
            end;
        end
        else
            dentroPalabra := false;  // Encontramos un separador, salimos de la palabra
    end;

    writeln('La cantidad de palabras es: ', palabras);
    Close(sec);
end.