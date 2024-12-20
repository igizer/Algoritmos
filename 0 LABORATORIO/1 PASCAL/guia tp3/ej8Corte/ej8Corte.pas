//8) Se tiene un archivo con los empleados de una empresa que posee los siguientes
//campos: Planta, Categoría y Nombre del empleado.
//La categoría es un número entero que va del 1 al 4. 
//Se pide mostrar el nombre de cada trabajador, y además la cantidad de trabajadores que
//hay por categoría, planta y total general.
//Para poder trabajar con este ejercicio se tiene que correr el siguiente programa que genera
//un archivo trabajadores.dat, con los datos necesarios para ejecutar el programa.
program ej8Corte;

type
formato = record
    planta: string[20];
    categoria: 1..4;
    nombre: string[50]
end;

var

    reg: formato;
    arch: File of formato;


    arch_sal: file of formato;

    resg_planta: string[20];
    resg_categoria: 1..4;

    cont_categoria, cont_planta, cont_totg: integer;

procedure Corte_cate();
begin
    Writeln('La cantidad de personas en la categoria ', resg_categoria, ' es: ', cont_categoria);
    cont_planta:= cont_planta + cont_categoria;
    cont_categoria:= 0;
    resg_categoria:= reg.categoria;
end;

procedure Corte_plan();
begin
    Corte_cate();
    Writeln('La cantidad de personas en la planta ', resg_planta, ' es: ', cont_planta);
    cont_totg:= cont_totg + cont_planta;
    cont_planta:= 0;
    resg_planta:= reg.planta;
end;

BEGIN

Assign(arch, 'C:\Users\xpoke\Downloads\Actividades en clase\0 LABORATORIO\1 PASCAL\guia tp3\ej8Corte\ejecutar.dat');

{$I-}
reset(arch);
{$I+}

Assign(arch_sal, 'C:\Users\xpoke\Downloads\Actividades en clase\0 LABORATORIO\1 PASCAL\guia tp3\ej8Corte\salida.dat');

rewrite(arch_sal);

if IOResult <> 0 then
begin
    write('ERROR: no se ha encontrado ningun archivo ".dat"');
    halt(2)
end;

read(arch, reg);

While Not EoF(arch) do
begin

    if resg_planta <> reg.planta then
    Corte_plan()
    else
    begin
        if resg_categoria <> reg.categoria then
        Corte_cate()
    end;

    Writeln('Nombre: ', reg.nombre);
    cont_categoria:= cont_categoria + 1;

    if reg.categoria = 4 then
            write(arch_sal, reg);

    read(arch, reg);

end;

Corte_plan();

Writeln('El total general de perdonas en cada planta es de', cont_totg);

close(arch);

END.




















program ej8Corte;

type
formato_empleado = record
 plan: string[20];
 cate: 1..4;
 nomb: string[50];
end;

var
 arch: file of formato_empleado;
 reg: formato_empleado;

 res_plan: string[20];
 res_cate: 1..4;

 contcate, contplan, conttot: integer;

procedure corte_cate();
begin
writeln('en la categoria: ', res_cate, ' la cantidad de personas son: ', contcate);
contplan:= contplan + contcate;
contcate:= 0;
res_cate:=reg.cate;
end;

procedure corte_plan();
begin
corte_cate();
writeln('El total de empleados en la planta: ', res_plan, ' son ', contplan);
conttot:= conttot + contplan;
contplan:= 0;
res_plan:=reg.plan;
end;

BEGIN

    Assign (arch, 'C:\Users\xpoke\Downloads\Actividades en clase\0 LABORATORIO\1 PASCAL\guia tp3\ej8Corte\ejecutar.dat');

    {$I-}
    Reset (arch);
    {$I+}

    if IOResult <> 0 then
    begin
        writeln('Error: el archivo empleados.dat exite?');
        halt(2); 
    end;

    contcate:=0;
    contplan:=0;
    conttot:=0;

    read(arch, reg);

    res_plan:=reg.plan;
    res_cate:=reg.cate;

    while not EOF(arch) do
    begin
        
        if res_plan <> reg.plan then
            corte_plan()
        else
        begin
            if res_cate <> reg.cate then
            corte_cate();
        end;

        writeln('Nombre: ',reg.nomb);
        contcate := contcate + 1;
        read(arch, reg);

        

    end;

    corte_plan();

    writeln('El total de empleados es de: ', conttot);
    close(arch);


END.
