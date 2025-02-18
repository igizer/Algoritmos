//Primer repaso 00:39:01 corrigiendo total 00:53:26
//EJERCICIO 8
//8) Se tiene un archivo con los empleados de una empresa que posee los siguientes
//campos: Planta, Categoría y Nombre del empleado.
//La categoría es un número entero que va del 1 al 4.
//Se pide mostrar el nombre de cada trabajador, y además la cantidad de trabajadores que
//hay por categoría, planta y total general.
//Para poder trabajar con este ejercicio se tiene que correr el siguiente programa que genera
//un archivo trabajadores.dat, con los datos necesarios para ejecutar el programa.

program repasogpt.pas;

type
    empleados = record
    Planta: string[50];
    Categoria: 1..4;
    Nombre: string [30];
    end;

var
    arch: file of empleados;
    reg: empleados;
    error: integer;
    cont_cat, cont_planta, cont_tot: integer;
    resg_cat: 1..4;
    resg_planta: string [50];

procedure corte_categoria();
    begin
        Writeln('La cantidad de empleados que hay en la categoría ', ord(resg_cat), ' es de ', cont_cat); // esto me olvidaba
        //el ord(resguardo) sive porque 1..4 es un subrango, tengo que hacerlo enrealidad un entero de esta forma
        //porque SINO NO PUEDO IMPRIMIRLO EN PANTALLA
        cont_planta:= cont_planta + cont_cat;
        cont_cat:= 0;
        resg_cat:= reg.Categoria
    end;

procedure corte_planta();
    begin
        corte_categoria;//corrigiendo agregué esto, me suelo olvidar, OJO A NO OLVIDARME ESTO
        Writeln('La cantidad de empleados que hay en la planta ', resg_planta, ' es de ', cont_planta);
        cont_tot:= cont_tot + cont_planta;
        cont_planta:= 0;
        resg_planta:= reg.Planta
    end;

begin

Assign(arch, 'C:\Codigo\Algoritmos\0 LABORATORIO\1 PASCAL\CORTES DE CONTROL\1\empleados.dat');

{$I-}
Reset(arch);
error:= IOResult;
{$I+}
if error <> 0 then
begin
    Writeln('Error: el archivo empleados.dat exite?');
    halt(2);
end;

read(arch, reg);
cont_cat:= 0;
cont_planta:= 0;
cont_tot:= 0;
resg_cat:= reg.Categoria;
resg_planta:= reg.Planta;

While not EOF(arch) do
begin

    If resg_planta <> reg.Planta then //no hace falta en este if ni en el else los begin end, porque solo manejan 1 linea
        corte_planta
    else If resg_cat <> reg.Categoria then
        corte_categoria;

    Writeln(reg.Nombre);
    cont_cat:= cont_cat + 1;

    read(arch, reg);

end;

corte_planta;

Writeln('El total general de empleados es de ', cont_tot);
Close(arch);
ReadLn;//se usa para esperar una entrada y poder leer el resultado

end.

