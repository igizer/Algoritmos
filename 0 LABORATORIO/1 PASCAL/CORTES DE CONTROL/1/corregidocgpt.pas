program corregidocgpt;
type
  EMPLEADOS = RECORD
    PLANTA: STRING[50];
    CATEGORIA: 1..4;
    NOMBRE: STRING[50];
  END;

var
  arch: file of empleados;
  reg: empleados;
  error, cat_emple, planta_emple, total_emple: integer;
  resg_planta: STRING[50];
  resg_cat: 1..4;

procedure corte_categoria();
begin
  writeln('Cantidad de empleados en categoría ', resg_cat, ': ', cat_emple);
  planta_emple := planta_emple + cat_emple;
  cat_emple := 0;
  resg_cat := reg.CATEGORIA;
end;

procedure corte_planta();
begin
  writeln('Cantidad de empleados en planta ', resg_planta, ': ', planta_emple);
  total_emple := total_emple + planta_emple;
  planta_emple := 0;
  resg_planta := reg.PLANTA;
end;

begin
  Assign(arch, 'C:\Codigo\Algoritmos\0 LABORATORIO\1 PASCAL\CORTES DE CONTROL\1\empleados.dat');
  {$I-}
  Reset(arch);
  error := IOResult;
  {$I+}
  if error <> 0 then
  begin
    writeln('Error: el archivo empleados.dat existe?');
    halt(2);
  end;

  if EOF(arch) then
  begin
    writeln('El archivo está vacío.');
    Close(arch);
    Halt(0);
  end;

  read(arch, reg);
  cat_emple := 0;
  planta_emple := 0;
  total_emple := 0;
  resg_planta := reg.PLANTA;
  resg_cat := reg.CATEGORIA;

  while not EOF(arch) do
  begin
    writeln('Nombre del empleado: ', reg.NOMBRE);
    cat_emple := cat_emple + 1;

    read(arch, reg);

    if resg_planta <> reg.PLANTA then
    begin
      corte_categoria();
      corte_planta();
    end
    else if resg_cat <> reg.CATEGORIA then
    begin
      corte_categoria();
    end;
  end;

  // Últimos cortes para el último grupo
  corte_categoria();
  corte_planta();
  
  writeln('Total de empleados: ', total_emple);
  Close(arch);
  Readln;
end.