program imc1;
var
  peso, altura, imc: real;

begin
    writeln('cual es su estatura en cm: ');
    readln(altura);
    writeln('cual es su peso en kg: ');
    readln(peso);
    imc := peso / (altura/100 * altura/100);
    writeln('su imc es: ', imc:0:1);

    if imc < 18.5 then
        writeln('bajo peso')
    else if (imc < 24.9) And (imc >= 18.5) then
        writeln('peso normal')
    else if (imc < 29.9) And (imc >= 25.0) then
        writeln('sobrepeso')
    else
        writeln('terrible gorda');
    readln;
end.