program yea;
type
formato_empleado = record
 plan: string[20];
 cate: 1..4;
 nomb: string[50];
end;
var
 arch: file of formato_empleado;
 reg: formato_empleado;
procedure cargarEmpleado(var emp: formato_empleado; p: string;
c:smallint; n:string);
begin
emp.plan:=p;
emp.cate:=c;
emp.nomb:=n;
end;
BEGIN
Assign (arch, 'empleados.dat');
rewrite(arch);
cargarEmpleado(reg, 'a', 1, 'Aranda');
write(arch, reg);
cargarEmpleado(reg, 'a', 1, 'Vera');
write(arch, reg);
cargarEmpleado(reg, 'a', 2, 'Arce');
write(arch, reg);
cargarEmpleado(reg, 'a', 2, 'Basualdo');
write(arch, reg);
cargarEmpleado(reg, 'a', 4, 'Corrales');
write(arch, reg);
cargarEmpleado(reg, 'a', 4, 'Derqui');
write(arch, reg);
cargarEmpleado(reg, 'a', 4, 'Gutierrez');
write(arch, reg);
cargarEmpleado(reg, 'a', 4, 'Morales');
write(arch, reg);
cargarEmpleado(reg, 'b', 1, 'Biondi');
write(arch, reg);
cargarEmpleado(reg, 'b', 1, 'Casas');
write(arch, reg);
cargarEmpleado(reg, 'b', 1, 'Davis');
write(arch, reg);
cargarEmpleado(reg, 'b', 1, 'Nogueira');
write(arch, reg);
cargarEmpleado(reg, 'b', 2, 'Milán');
write(arch, reg);
cargarEmpleado(reg, 'b', 2, 'Molina');
write(arch, reg);
cargarEmpleado(reg, 'b', 2, 'Portales');
write(arch, reg);
cargarEmpleado(reg, 'b', 3, 'Cortezs');
write(arch, reg);
cargarEmpleado(reg, 'b', 3, 'Mendez');
write(arch, reg);
cargarEmpleado(reg, 'b', 3, 'Nogueira');
write(arch, reg);
cargarEmpleado(reg, 'b', 3, 'Salerno');
write(arch, reg);
cargarEmpleado(reg, 'b', 4, 'Casas');
write(arch, reg);
 close(arch);
END.