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

program IMC;
var
    edad, imc, peso, altura: integer
procedure CalcularIMC();
begin
    imc:= peso / (altura*altura);
    if edad>=20 then //Mayor de 20
    begin
        case imc of
            0..18: Writeln('Su imc (', imc, ') es peso bajo');
            19..24: Writeln('Su imc (', imc, ') es peso normal');
            25..29: Writeln('Su imc (', imc, ') es sobrepeso');
        else
            Writeln('Su imc (', imc, ') es peso obeso')
        end;
    end
    else //Menor de 20 ERROR LINEA 29
    begin
        case imc of
            0..14: Writeln('Su imc (', imc, ') es peso bajo');
            14..19: Writeln('Su imc (', imc, ') es peso normal');
            20..24: Writeln('Su imc (', imc, ') es sobrepeso');
        else
            Writeln('Su imc (', imc, ') es peso obeso')
        end;
    end;
end;

begin
    Writeln('Escriba su edad');
    Readln(edad);
    Writeln('Escriba su peso');
    Readln(peso);
    Writeln('Escriba su altura');
    Readln(altura);
    CalcularIMC();
end.



























program ej01guia2;
var
 edad, peso, altura: integer;
 alturam: real; //Altura en metros
 imc : integer;
 resultado: string; 
begin
 writeln('Ingrese la edad que este entre 4 y 140 años');
 readln(edad);
 writeln('Ingrese el peso que este entre 10 y 300 kg');
 readln(peso);
 writeln('Ingrese el altura que este entre 10 y 240 cm');
 readln(altura);
 // Paso de cm a metros
 alturam := altura / 100;
 imc := Round(peso / (alturam * alturam));
 if edad >= 20 then
 begin
 case imc of 
 1..18 : resultado := 'Bajo Peso';
 19..24: resultado := 'Normal';
 25..29: resultado := 'Sobrepeso';
 else
 resultado := 'Obeso';
 end;
 end
 else
 begin
 case imc of
 1..13 : resultado := 'Bajo Peso';
 14..19: resultado := 'Normal';
 20..24: resultado := 'Sobrepeso';
 else
 resultado := 'Obeso';
 end;
 end;
 writeln('Su IMC es: ', IMC, ' usted tiene: ', resultado);
end.

//Siguiendo la forma propuesta de resolver el ejercicio, codificar una subacción que valide
//que los datos a la entrada sean los correctos (que estén dentro de los rangos propuestos),
//si el usuario ingresa un valor no válido, volver a pedirlo.