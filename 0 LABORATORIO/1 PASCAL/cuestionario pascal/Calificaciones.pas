program Calificaciones;

uses
  SysUtils;

var
  archivo: TextFile;
  calificacion: Integer;
  suma, contador, contadorMayorIgual: Integer;
  promedio: Real;
  calificaciones: array of Integer;

begin
  AssignFile(archivo, 'C:\Users\xpoke\Downloads\Actividades en clase\0 LABORATORIO\1 PASCAL\cuestionario pascal\PROBLEMA1.txt');
  Reset(archivo);

  suma := 0;
  contador := 0;
  contadorMayorIgual := 0;

  // Leer calificaciones desde el archivo y calcular la suma y el contador
  while not Eof(archivo) do
  begin
    ReadLn(archivo, calificacion);
    suma := suma + calificacion;
    contador := contador + 1;
    SetLength(calificaciones, contador);
    calificaciones[contador - 1] := calificacion;
  end;

  CloseFile(archivo);

  // Calcular el promedio
  if contador > 0 then
    promedio := suma / contador
  else
    promedio := 0;

  // Contar cuántos estudiantes tienen calificación mayor o igual al promedio
  for calificacion in calificaciones do
  begin
    if calificacion >= promedio then
      contadorMayorIgual := contadorMayorIgual + 1;
  end;

  // Mostrar el promedio y el número de estudiantes que superaron o igualaron el promedio
  WriteLn('Promedio de calificaciones: ', promedio:0:2);
  WriteLn('Número de estudiantes con calificación mayor o igual al promedio: ', contadorMayorIgual);
end.