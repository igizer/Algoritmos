//5) Se posee una secuencia con los mensajes internos enviados entre los empleados de una
//empresa, la secuencia está formada de la siguiente manera: 3 caracteres para el destino, 3
//caracteres para el origen, 3 caracteres que indican la cantidad de caracteres que tiene el
//mensaje y el mensaje. No existe ningún carácter que separe un mensaje de otro.
//se pide:
//a) Escribir una secuencia de salida con todos los mensajes (incluyendo origen y
//destino) que van hacia el área de Mantenimiento (que destino comience con 1).
//b) Contar la cantidad de mensajes que se dirigen hacia el área de sistemas (que
//destino comience con 23).
program ej5guia3;
uses math; // Para poder usar el "Power(numero_a_elevar, exponente)"" que es la potencia
var 
 sec, sal: file of char;
 v : char;
 i, error, cont: integer;
 esMantenimiento, esSistemas: boolean;
 long: integer;
function aEntero(caracter: char): integer;
begin
 case caracter of
 '0': aEntero:= 0;
 '1': aEntero:= 1;
 '2': aEntero:= 2;
 '3': aEntero:= 3;
 '4': aEntero:= 4;
 '5': aEntero:= 5;
 '6': aEntero:= 6;
 '7': aEntero:= 7;
 '8': aEntero:= 8;
 '9': aEntero:= 9;
 end;
end;
BEGIN

    Assign (sal, 'C:\Codigo\Algoritmos\0 LABORATORIO\1 PASCAL\guia tp3\ej5funcion\salidaej5.txt');
    rewrite(sal);

    Assign (sec, 'C:\Codigo\Algoritmos\0 LABORATORIO\1 PASCAL\guia tp3\ej5funcion\entradaej5.txt');
 
    {$I-}
    Reset(sec);
    {$I+}

    if IOResult <> 0 then
    begin
        writeln('Error: el archivo archivo.txt existe?');
        halt(2); 
    end;

    read(sec, v);

    cont := 0;

    While Not Eof(sec) do
    begin
        esMantenimiento := false;
        esSistemas := false;
        if v = '1' then
        esMantenimiento := true
        else
        begin
        if v = '2' then
        esSistemas := true; 
        end;

        if esMantenimiento then
        write(sal, v);

        read(sec, v);

        if (v = '3') and esSistemas then
        cont := cont + 1;

        for i:= 1 to 4 do
        begin
            if esMantenimiento then
            write(sal, v);

            read(sec, v);
        end;

        long:= 0;

        for i := 2 downto 0 do
        begin
            read(sec, v);
            if esMantenimiento then
                write(sal, v);
            long := long + aEntero(v) * round(Power(10, i));// Acá está lo del uses math
        end;

        for i:= 1 to long do
        begin
            if esMantenimiento then
                write(sal, v);
            read(sec, v);
        end;

        if esMantenimiento then
            write(sal, '#');
    end;

    writeln('cantidad de mensajes para Sistemas: ', cont);

    close(sal); 
    close(sec); 
END.