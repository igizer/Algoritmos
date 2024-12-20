program Pruebafranco;
var x, y, suma, resta, multiplicacion, divE, modE: Integer;
division : real;
begin

    x := 55;
    Y := 30;

    suma := x + y;
    writeln('el resultasdo de la suma entre ', x, ' y ', y, ' es ', suma);

    resta := x - y;
    writeln('el resultasdo de la resta entre ', x, ' y ', y, ' es ', resta);

    division := x / y;
    writeln('el resultasdo de la division entre ', x, ' y ', y, ' es ', division:0:3);

    multiplicacion := x * y;
    writeln('el resultasdo de la multiplicacion entre ', x, ' y ', y, ' es ', multiplicacion);

    divE := x div y;
    writeln('el resultasdo de la Division Entera entre ', x, ' y ', y, ' es ', divE);

    modE := x mod y;
    writeln('el resultasdo del resto entre ', x, ' y ', y, ' es ', modE);
end.

