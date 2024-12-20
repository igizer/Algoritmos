//1) Calcular el IMC (índice de masa corporal) para esto se debe tener en cuenta los
//siguientes datos: Altura, Edad y Peso.
//El cálculo del IMC es: IMC = Peso / (altura ^2)
//Con el Peso en Kg y la altura en metros.
//La altura puede ser de 10 cm a 240 cm, la edad puede ser de 4 a 140 años y el Peso puede ser
//de 10 kg a 300 kg.
//Según el valor que da el calculo podemos tener:
//- Para los mayores o iguales a 20 años, el IMC menor a 19 es bajo peso, entre 19 y
//24 es normal, entre 25 y 29 es sobrepeso, y más de 30 es obeso.
//- Para menores de 20 años el IMC es: menos de 14 es bajo peso, entre entre 14 y 19 es
//normal, entre 20 y 24 es sobrepeso y más de 25 es obeso.
program CALCULADORAIMC;
var
i: integer;
imc, peso, edad, altura: real;

begin

    for i:= 0 to 30 do
        begin
            writeln();
        end;

    writeln('CALCULADORA DE INDICE CORPORAL (IMC)');
    writeln('A continuacion escriba su altura en cm (Ej: 170)');
    readln(altura);
    writeln('Ahora escriba su peso en Kg (Ej: 60)');
    readln(peso);

    altura := altura;
    imc:= (peso / (altura*altura))*10000;
    

    writeln('Su Indice de Masa Muscular (IMC) es ', imc);
end.