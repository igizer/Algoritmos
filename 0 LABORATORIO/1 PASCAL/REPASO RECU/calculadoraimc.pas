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

program calculadoraimc;
var
edad, peso: integer;
altura, imc: Real;
condicion= ('bajo peso', 'normal', 'sobrepeso', 'obeso');

BEGIN
writeln('CALCULADORA DE INDICE DE MASA CORPORAL');

writeln('Escriba su edad');
readln(edad);
Writeln('Escriba a continuación su peso en kg');
readln(peso);
writeln('Escriba a continuación su altura en metros');
readln(altura);

imc:= peso / (altura*altura);

writeln('Su Indice de masa corporal de de ', imc:0:2, ' puntos');
if edad >= 20 then
    begin
        if imc < 19 then
            condicion := 'bajo peso'
        else if imc <= 24 then
            condicion := 'normal'
        else if imc <= 29 then
            condicion := 'sobrepeso'
        else
            condicion := 'obeso';
    end
    else
    begin
        if imc < 14 then
            condicion := 'bajo peso'
        else if imc <= 19 then
            condicion := 'normal'
        else if imc <= 24 then
            condicion := 'sobrepeso'
        else
            condicion := 'obeso';
    end;

//if edad>19 then
//    begin
//        case imc of
//            0..19: condicion:= 'bajo peso';
//            19..24: condicion:= 'normal';
//            24,0001..29: condicion:= 'sobrepeso';
//            29..1000: condicion:= 'obeso';
//        end;
//    end;
//else
//    begin
//        case imc of
//            0..14: condicion:= 'bajo peso';
//            14..19: condicion:= 'normal';
//            20,0001..24: condicion:= 'sobrepeso';
//            25..1000: condicion:= 'obeso';
//        end;
//end;

writeln('Su condicion es: ', condicion);


writeln('Según el valor que da el calculo podemos tener:');
writeln('- Para los mayores o iguales a 20 años, el IMC menor a 19 es bajo peso, entre 19 y');
writeln('24 es normal, entre 25 y 29 es sobrepeso, y más de 30 es obeso.');
writeln('- Para menores de 20 años el IMC es: menos de 14 es bajo peso, entre entre 14 y 19 es');
writeln('normal, entre 20 y 24 es sobrepeso y más de 25 es obeso.');

END.