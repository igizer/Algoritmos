//12) Dado el ejercicio anterior, realizar el mismo algoritmo, solo que ahora no se muestre el 
//valor de la contraseña cuandos se escribe por el teclado.
program Ej12;
 uses Crt;
 const
 user = 'Usuario';
 pass = 'password';
 var
 usuario, contrasena: string;
 letra: char;
 begin
 write('Ingrese el usuario: ');
 readln(usuario);
 write('Ingrese el password: ');
 contraseña :='';
 Repeat 
 letra:=readkey; 
 If letra<>#13 Then 
 Begin 
 contraseña :=contraseña + letra; 
 Write(chr(42)) 
 End; 
 until letra=#13; 
 writeln();
 if user = usuario then
 begin
 if pass = contraseña then
 begin
 writeln('Usuario y contraseña válido.');
 end;
 end
 else 
 writeln('Usuario o contraseña inválida.'); 
end.
