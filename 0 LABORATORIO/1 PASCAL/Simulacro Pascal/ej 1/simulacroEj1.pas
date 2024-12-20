//Crea un programa en Pascal que lea una lista de calificaciones (números enteros)
//desde un archivo de texto llamado problema1.txt, calcule el promedio de estas
//calificaciones y determine cuántos estudiantes obtuvieron una calificación mayor
//o igual al promedio. El programa debe mostrar el promedio y el número de estudiantes
//que superaron o igualaron el promedio.
//C:\Users\xpoke\Downloads\Actividades en clase\0 LABORATORIO\1 PASCAL\Simulacro Pascal\ej 1\PROBLEMA.txt
program simulacroEj1;
var
    sec: text;
    v: double;
    promedio, acumulador, personas, supera_promedio : real;

BEGIN
    Assign(sec, 'C:\Users\xpoke\Downloads\Actividades en clase\0 LABORATORIO\1 PASCAL\Simulacro Pascal\ej 1\PROBLEMA1.txt');

    {$I-}
    reset(sec);
    {$I+}

    if IOResult <> 0 then
    begin
        writeln('No se encontró el archivo.txt');
        halt(2);
    end;

    read(sec, v);

    personas:= 0;
    acumulador:= 0;

    while Not EoF(sec) do
    begin
        personas:= personas + 1;
        acumulador:= acumulador + v;
        read(sec, v);
    end;
    close(sec);

    promedio:= acumulador/personas;

    Writeln('El promedio de las notas es de ', promedio:0:2);

    {$I-}
    reset(sec);
    {$I+}

    read(sec, v);

    supera_promedio:= 0;

    while Not EoF(sec) do
    begin
        if v >= promedio then
            supera_promedio:= supera_promedio + 1;
        read(sec, v);

    end;
    close(sec);

    writeln('La cantidad de personas que superan el promedio es de ', supera_promedio:0:2);

END.





























program simulacroEj1;
var
    sec: text;
    v: real;
    personas, acumulador, promedio, mayores_al_promedio: real;
BEGIN
    Assign(sec, 'C:\Users\xpoke\Downloads\Actividades en clase\0 LABORATORIO\1 PASCAL\Simulacro Pascal\ej 1\PROBLEMA1.txt');

    {$I-}
    Reset(sec);
    {$I+}

    if IOResult <> 0 then
        begin
            writeln('No se encontró ningún archivo');
            halt(2);
        end;
    
    read(sec, v);

    personas:= 0;
    acumulador:= 0;

    While Not EoF(sec) do
        begin
            acumulador:= acumulador + v;
            personas:= personas + 1;
            read(sec, v);
        end;
    
    promedio:= acumulador/personas;
    writeln('El promedio de estas calificaciones es de ', promedio:0:4);
    close(sec);

    {$I-}
    Reset(sec);
    {$I+}

    mayores_al_promedio:= 0;

    read(sec, v);

    While Not EoF(sec) do
        begin
            if v >= promedio then
                mayores_al_promedio:= mayores_al_promedio + 1;
            read(sec,v)
        end;
    
    Writeln('Y las personas que sacaron mayor nota al promedio son ', mayores_al_promedio:0:4);
    close(sec);
        
    

END.