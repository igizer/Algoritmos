program BatallaNaval;

const
  N = 5;

type
  TBoard = array[1..N, 1..N] of Char;

var
  playerBoard, enemyBoard: TBoard;
  shipsRemaining: Integer;

procedure InitializeBoard(var b: TBoard);
var
  i, j: Integer;
begin
  for i := 1 to N do
    for j := 1 to N do
      b[i, j] := '~';
end;

procedure PlaceShip(var b: TBoard; row, col: Integer);
begin
  b[row, col] := 'B';
end;

procedure DisplayBoard(var b: TBoard);
var
  i, j: Integer;
begin
  writeln('Tablero:');
  for i := 1 to N do
  begin
    for j := 1 to N do
    begin
      if b[i, j] = 'B' then
        write('~ ', ' ')
      else
        write(b[i, j], ' ');
    end;
    writeln;
  end;
end;

procedure Attack(var b: TBoard; row, col: Integer; var shipsRemaining: Integer);
begin
  if b[row, col] = 'B' then
  begin
    writeln('¡Impacto!');
    b[row, col] := 'X';
    shipsRemaining := shipsRemaining - 1;
  end
  else if b[row, col] = '~' then
  begin
    writeln('Agua.');
    b[row, col] := 'O';
  end
  else
    writeln('Ya atacaste esta posición.');
end;

procedure PlaceEnemyShips(var b: TBoard);
begin
  PlaceShip(b, 1, 1);
  PlaceShip(b, 2, 3);
  PlaceShip(b, 4, 4);
  PlaceShip(b, 5, 2);
  PlaceShip(b, 3, 5);
end;

procedure PlayGame;
var
  row, col: Integer;
begin
  shipsRemaining := 5;
  InitializeBoard(playerBoard);
  InitializeBoard(enemyBoard);
  PlaceEnemyShips(enemyBoard);

  writeln('Bienvenido a Batalla Naval!');
  DisplayBoard(playerBoard);

  while shipsRemaining > 0 do
  begin
    writeln('Introduce las coordenadas de tu disparo (fila columna):');
    readln(row, col);

    if (row >= 1) and (row <= N) and (col >= 1) and (col <= N) then
    begin
      Attack(enemyBoard, row, col, shipsRemaining);
      DisplayBoard(playerBoard);
    end
    else
      writeln('Coordenadas fuera de rango. Inténtalo de nuevo.');

    writeln('Barcos restantes: ', shipsRemaining);
  end;

  writeln('¡Felicidades! Hundiste todos los barcos enemigos.');
end;

begin
  PlayGame;
end.