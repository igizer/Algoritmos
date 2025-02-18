program repasoej5;
uses math; // Para poder usar la función Power
var 
  sec, sal: file of char;
  v : char;
  i, error, cont: integer;
  esMantenimiento, esSistemas: boolean;
  long: integer;

function aEntero(caracter: char): integer;
begin
  case caracter of
    '0': aEntero := 0;
    '1': aEntero := 1;
    '2': aEntero := 2;
    '3': aEntero := 3;
    '4': aEntero := 4;
    '5': aEntero := 5;
    '6': aEntero := 6;
    '7': aEntero := 7;
    '8': aEntero := 8;
    '9': aEntero := 9;
  end;
end;

BEGIN
  Assign(sal, 'C:\Codigo\Algoritmos\0 LABORATORIO\1 PASCAL\guia tp3\ej5funcion\salidaej5.txt');
  Rewrite(sal);
  Assign(sec, 'C:\Codigo\Algoritmos\0 LABORATORIO\1 PASCAL\guia tp3\ej5funcion\entradaej5.txt');

  {$I-}
  Reset(sec);
  {$I+}

  if IOResult <> 0 then
  begin
    writeln('Error: el archivo archivo.txt existe?');
    halt(2); 
  end;

  cont := 0;
  while not Eof(sec) do
  begin
    esMantenimiento := false;
    esSistemas := false;

    read(sec, v); // Lee destino (primer caracter)
    if v = '1' then
      esMantenimiento := true
    else if v = '2' then
      esSistemas := true;

    if esMantenimiento then
      write(sal, v);  // Escribe destino

    read(sec, v); // Lee origen
    if esMantenimiento then
      write(sal, v);  // Escribe origen

    read(sec, v); // Lee cantidad de caracteres en el mensaje
    long := aEntero(v); // Convierte cantidad de caracteres a entero

    // Lee y escribe los caracteres del mensaje
    for i := 1 to long do
    begin
      if not Eof(sec) then
      begin
        read(sec, v);
        if esMantenimiento then
          write(sal, v);
      end;
    end;

    if esMantenimiento then
      write(sal, '#');

    // Si el destino es Sistemas, contamos el mensaje
    if esSistemas then
      cont := cont + 1;
  end;

  writeln('Cantidad de mensajes para Sistemas: ', cont);

  Close(sal); 
  Close(sec); 
END.  { Este END. es el final correcto del programa }