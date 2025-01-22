//1) Usando el tipo de dato Record, cargar 3 libros en memoria y después realizar un listado
//de todos los libros. De los libros se deben almacenar: El título, el nombre del autor, un breve
//resumen y el id del libro.
//Teniendo el siguiente código, realice el codigo de una subacción donde el usuario pueda
//cargar a mano los 3 libros.
program repasoej1reg;
type
    formato_libro= record
        titulo: string[50];
        autor: string[20];
        resumen: string[250];
        id: longint;
end;

var
    libro1, libro2, libro3: formato_libro;

procedure MostrarLibro(unlibro:formato_libro);
begin
    Writeln('Titulo', unlibro.titulo);
    Writeln('Autor', unlibro.autor);
    Writeln('Resumen', unlibro.resumen);
    Writeln('Id Libro', unlibro.id);
end;

BEGIN
    libro1.titulo:= 'El increible libro';
    libro1.autor:= 'Don increible';
    libro1.resumen:= 'Un libro increible.';
    libro1.id:= 423232;
    libro2.titulo:= 'Otra titulo';
    libro2.autor:= 'El otro autor';
    libro2.resumen:= 'un libro cuando no esta el primero.';
    libro2.id:= 342345;
    libro3.titulo:= 'Este es un titulo';
    libro3.autor:= 'Este es el autor';
    libro3.resumen:= 'Esto, es un pequeño resumen del libro.';
    libro3.id:= 234534;
    mostrarlibro(libro1);
    writeln();
    mostrarlibro(libro2);
    writeln();
    mostrarlibro(libro3);
END.    