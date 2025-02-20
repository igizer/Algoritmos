//3) Escribir un algoritmo donde se ingresan 3 números distintos y se deben ordenar de
//menor a mayor, e imprimir por pantalla.

//10:00

program menormayor;
var
x1, x2, x3: integer;

procedure intercambio(var x, y: integer); //OJO CON LA PALABRA VAR, POR QUE SE AGREGA?

    var
        aux: integer;
    begin
        aux:= x;
        x:= y;
        y:= aux;
    end;

//CON RESPECTO AL VER USADO EN (VAR x, y: integer)
//Esto significa que cualquier cambio que realices sobre esos parámetros
//dentro del procedimiento afectará directamente a las variables originales
//que se pasaron al llamar al procedimiento. Es decir, no solo se pasa el valor
//de las variables, sino también una referencia a las mismas.
//Por defecto, cuando pasas una variable a un procedimiento o función,
//se pasa por valor. Esto significa que el procedimiento recibe una copia del valor de la variable,
//y cualquier cambio que hagas dentro del procedimiento no afecta a la variable original.
//Si quieres que el procedimiento modifique directamente las variables que pasas, entonces debes usar var en los parámetros.
//Esto indica que esas variables se pasan por referencia y no por valor.

BEGIN

    Writeln('Te ayudamos a ordenar 3 número de menor a mayor!');
    Writeln('Ingrese el primer numero');
    readln(x1);
    Writeln('Ingrese el segundo numero');
    readln(x2);
    Writeln('Ingrese el tercer numero');
    readln(x3);

    if x1 > x2 then
        intercambio(x1, x2);
    if x1 > x3 then
        intercambio(x1, x3);
    if x2 > x3 then
        intercambio(x2, x3);

    writeln('El orden de los numeros de menor a mayor es: ');
    writeln(x1, ', ', x2, ', ', x3);

END.
        
