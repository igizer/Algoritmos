program probandoalgo;

type
    empleados = record
        Planta: string[50];
        Categoria: 1..4;
        Nombre: string[50];  // Asegúrate de que el campo Nombre tenga la longitud correcta
    end;

var
    arch: file of empleados;
    reg: empleados;
    error: integer;
    cont_cat, cont_planta, cont_tot: integer;
    resg_cat: 1..4;
    resg_planta: string[50];

procedure corte_categoria();
begin
    Writeln('La cantidad de empleados que hay en la categoría ', resg_cat, ' es de ', cont_cat);
    cont_planta := cont_planta + cont_cat;
    cont_cat := 0;
end;

procedure corte_planta();
begin
    corte_categoria;
    Writeln('La cantidad de empleados que hay en la planta ', resg_planta, ' es de ', cont_planta);
    cont_tot := cont_tot + cont_planta;
    cont_planta := 0;
end;

begin
    Assign(arch, 'empleados.dat');  // Usa la ruta correcta del archivo .dat

    {$I-}
    Reset(arch);  // Abrimos el archivo en modo de lectura
    error := IOResult;
    {$I+}
    if error <> 0 then
    begin
        Writeln('Error: el archivo empleados.dat existe?');
        halt(2);
    end;

    cont_cat := 0;
    cont_planta := 0;
    cont_tot := 0;

    // Leer el primer registro
    read(arch, reg);
    resg_cat := reg.Categoria;
    resg_planta := reg.Planta;

    while not EOF(arch) do
    begin
        // Cambiar de planta o de categoría
        if resg_planta <> reg.Planta then
            corte_planta
        else if resg_cat <> reg.Categoria then
            corte_categoria;

        // Mostrar el nombre del empleado
        Writeln('Empleado: ', reg.Nombre);

        // Contar el empleado en la categoría actual
        cont_cat := cont_cat + 1;

        // Leer siguiente registro
        read(arch, reg);
    end;

    corte_planta;  // Asegurarse de realizar el corte final

    Writeln('El total general de empleados es de ', cont_tot);

    Close(arch);
    ReadLn;
end.