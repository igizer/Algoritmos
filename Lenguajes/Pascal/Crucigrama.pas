// PONGAN ACA ARRIBA LOS INTEGRANTES y COM 1.6
program Crucigrama;

const
  N = 5; // Tamaño del crucigrama

type
  TBoard = array[1..N, 1..N] of Char;
  TWord = record
    word: string;
    row, col: Integer;
    direction: Char; // 'H' para horizontal, 'V' para vertical
  end;

var
  board: TBoard;
  words: array[1..3] of TWord;
  i, j: Integer;

procedure InitializeBoard(var b: TBoard);
var
  i, j: Integer;
begin
  for i := 1 to N do
    for j := 1 to N do
      b[i, j] := '_';
end;

procedure PlaceWord(var b: TBoard; w: TWord);
var
  i: Integer;
begin
  if w.direction = 'H' then
    for i := 1 to Length(w.word) do
      b[w.row, w.col + i - 1] := w.word[i]
  else
    for i := 1 to Length(w.word) do
      b[w.row + i - 1, w.col] := w.word[i];
end;

procedure DisplayBoard(var b: TBoard);
var
  i, j: Integer;
begin
  for i := 1 to N do
  begin
    for j := 1 to N do
      write(b[i, j], ' ');
    writeln;
  end;
end;

procedure InitializeWords;
begin
  words[1].word := 'ARRAY';
  words[1].row := 1;
  words[1].col := 1;
  words[1].direction := 'H';

  words[2].word := 'UPDATE';
  words[2].row := 2;
  words[2].col := 1;
  words[2].direction := 'V';

  words[3].word := 'ALGO';
  words[3].row := 3;
  words[3].col := 3;
  words[3].direction := 'H';
end;

begin
  InitializeBoard(board);
  InitializeWords;
  
  for i := 1 to 3 do
    PlaceWord(board, words[i]);
  
  DisplayBoard(board);
end.