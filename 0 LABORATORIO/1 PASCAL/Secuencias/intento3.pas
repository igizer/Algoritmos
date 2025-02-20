program intento3;
const
    signos = [' ', ',', '.', '?', '¿', '¡', '!'];  // Definir separadores de palabras
var
    sec: Text;
    v: Char;
    error, palabras: Integer;
    dentroPalabra: Boolean;
BEGIN
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
    dentroPalabra := False;

    while not EOF(sec) do
    begin
        Read(sec, v);

        // Si encontramos una letra (o cualquier cosa que no sea un signo)
        if not (v in signos) then
        begin
            if not dentroPalabra then
            begin
                palabras := palabras + 1;
                dentroPalabra := True;
            end;
        end
        else
            dentroPalabra := False; // Si encontramos un separador, marcamos el final de la palabra
    end;

    WriteLn('La cantidad de palabras en el texto es: ', palabras);
    Close(sec);
END.