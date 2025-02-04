program corte_control2;
type
EMPLEADOS= RECORD
  PLANTA: STRING [50];
  CATEGORIA: 1..4;
  NOMBRE: STRING [50];
  END;
var
arch: file of empleados;
reg: empleados;
error: integer;
cat_emple,planta_emple,total_emple:integer;
resg_planta:STRING [50];
resg_cat:1..4;
procedure corte_categoria();
 begin
 writeln ('Cantidad de empleados',cat_emple);
 planta_emple:= planta_emple + cat_emple;
 cat_emple:=0;
 resg_cat:= reg.CATEGORIA;
  end;
procedure corte_planta();
begin
 writeln ('Cantidad de empleados',planta_emple);
 total_emple:= total_emple + planta_emple;
 planta_emple:=0;
 resg_planta:= reg.PLANTA;
end;
Begin
 Assign (arch, 'C:\Codigo\Algoritmos\0 LABORATORIO\1 PASCAL\CORTES DE CONTROL\1\empleados.dat');
{$I-}
Reset (arch);
error:= IOResult;
{$I+}
if error <> 0 then
begin
writeln('Error: el archivo empleados.dat exite?');
halt(2);
end;
 read (arch,reg);
 cat_emple:=0;
 planta_emple:=0;
 total_emple:=0;
 resg_planta:=reg.PLANTA;

resg_cat:=reg.CATEGORIA;
while not EOF(ARCH) do
begin
   writeln('Nombre del empleado',reg.NOMBRE);
   cat_emple:= cat_emple + 1;
  if  resg_planta <> reg.PLANTA THEN
  corte_planta

  else
  begin
      if resg_cat <> reg.CATEGORIA THEN
       corte_categoria;

   end;
   read (arch,reg);
end;
corte_planta;
writeln('Total de cantidad de empleados es de', total_emple);
close(arch);
readln;
END.