program guia1ej1;
var
x, y, suma, resta, multi, resto, division: integer;
BEGIN
    x:= 55;
    y:= 30;
    suma:= x + y;
    resta:= x - y;
    multi:= x * y;
    division:= x div y;
    resto:= x mod y;

    writeln('suma: 55+30=', suma);
    writeln('resta: 55-30=', resta);
    writeln('multiplicacion: 55*30=', multi);
    writeln('division entera: 55 div 30=', division);
    writeln('resto division entera: 55 mod 30=', resto);
END.
//1) Escribir un programa en pascal que dado dos números enteros realice la suma, resta,
//multiplicación, división entera y resto de los mismos. Los numero son: X = 55 e Y = 30.