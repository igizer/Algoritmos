// COM RECURSANTES 1.6 Integrantes: Scherf Sergio, Torras Agustin, Aveiro Alejo, Paredes Fabricio
program Sudoku;

const
  N = 9;

type
  TBoard = array[1..N, 1..N] of Integer;

var
  board: TBoard;

procedure DisplayBoard(var b: TBoard);
var
  i, j: Integer;
begin
  writeln('Tablero de Sudoku:');
  for i := 1 to N do
  begin
    for j := 1 to N do
    begin
      if b[i, j] = 0 then
        write('. ')
      else
        write(b[i, j], ' ');
    end;
    writeln;
  end;
end;

function IsValidMove(var b: TBoard; row, col, num: Integer): Boolean;
var
  i, j, startRow, startCol: Integer;
begin

  for i := 1 to N do
  begin
    if (b[row, i] = num) or (b[i, col] = num) then
      Exit(False);
  end;

  startRow := ((row - 1) div 3) * 3 + 1;
  startCol := ((col - 1) div 3) * 3 + 1;
  for i := startRow to startRow + 2 do
    for j := startCol to startCol + 2 do
      if b[i, j] = num then
        Exit(False);

  IsValidMove := True;
end;

procedure FillBoard(var b: TBoard);
begin
  b[1, 1] := 5; b[1, 2] := 3; b[1, 5] := 7;
  b[2, 1] := 6; b[2, 4] := 1; b[2, 5] := 9; b[2, 6] := 5;
  b[3, 2] := 9; b[3, 3] := 8; b[3, 8] := 6;
end;

procedure PlaySudoku(var b: TBoard);
var
  row, col, num: Integer;
begin
  repeat
    writeln('Ingresa la fila (1-9), columna (1-9) y número (1-9) separados por espacios (0 0 0 para salir):');
    readln(row, col, num);

    if (row = 0) and (col = 0) and (num = 0) then
      break;

    if (row < 1) or (row > 9) or (col < 1) or (col > 9) or (num < 1) or (num > 9) then
    begin
      writeln('Entrada no válida. Inténtalo de nuevo.');
      continue;
    end;

    if IsValidMove(b, row, col, num) then
    begin
      b[row, col] := num;
      DisplayBoard(b);
    end
    else
      writeln('Movimiento no válido. Inténtalo de nuevo.');
  until False;
end;

begin
  FillBoard(board);
  DisplayBoard(board);
  PlaySudoku(board);
  writeln('Gracias por jugar al Sudoku!');
end.