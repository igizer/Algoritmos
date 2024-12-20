//Un banco necesita saber operaciones válidas luego de realizar varias transacciones.
//El archivo de datos problema4.txt contiene registros de transacciones con número de
//transacción (termina con '#'), tipo (D para depósito, R para retiro), y un caracter
//que indica si la operación ha sido válida (S,N).
//Pregunta: ¿Qué porcentaje representan las operaciones válidas tipo Retiro sobre el
//total de Retiros?

program ejercicio1;
var
    sec: file of Char;
    v, resg_v1, resg_v2: Char;
    retiro_valido, retiro_total, promedio: real;
BEGIN


    Assign(sec, 'C:\Users\xpoke\Downloads\Actividades en clase\0 LABORATORIO\1 PASCAL\EXAMEN\PROBLEMA4.txt');

    {$I-}
    Reset(sec);
    {$I+}

    if IOResult <> 0 then
    begin
        writeln('Error: el archivo existe?');
        halt(2);
    end;

    read(sec, v);

    retiro_total:= 0;
    retiro_valido:= 0;

    While Not EoF(sec) do
    begin
        While (v <> '#') and Not EoF(sec) do 
        read(sec,v);
        
        if Not EoF(sec) then
        begin

            read(sec,v);

            

            resg_v1:= v;
            read(sec, v);
            resg_v2:= v;
            read(sec, v);

            if resg_v1 = 'R' then
            begin
                retiro_total:= retiro_total + 1;
                if resg_v2 = 'S' then
                    retiro_valido:= retiro_valido + 1;
            end;
        
        end;
    end;

        promedio:= retiro_valido * 100 / retiro_total;

        writeln('El porcentaje del retiro valido sobre el total de retiros es de: %', promedio:0:2);

END.
















program ValidTransactions;

uses
  SysUtils;

type
  TTransaction = record
    Number: string;
    TypeOp: char;
    Valid: char;
  end;

var
  InputFile: TextFile;
  Line: string;
  Transaction: TTransaction;
  ValidCount: Integer;

function ParseTransaction(Line: string): TTransaction;
var
  HashPos: Integer;
begin
  HashPos := Pos('#', Line);
  if HashPos > 0 then
  begin
    Result.Number := Copy(Line, 1, HashPos - 1);
    Result.TypeOp := Line[HashPos + 1];
    Result.Valid := Line[HashPos + 2];
  end
  else
  begin
    Result.Number := '';
    Result.TypeOp := ' ';
    Result.Valid := ' ';
  end;
end;

begin
  AssignFile(InputFile, 'C:\Users\xpoke\Downloads\Actividades en clase\0 LABORATORIO\1 PASCAL\EXAMEN\PROBLEMA4.txt');
  Reset(InputFile);
  ValidCount := 0;

  while not Eof(InputFile) do
  begin
    ReadLn(InputFile, Line);
    Transaction := ParseTransaction(Line);
    if Transaction.Valid = 'S' then
      Inc(ValidCount);
  end;

  CloseFile(InputFile);

  WriteLn('Número de transacciones válidas: ', ValidCount);
end.