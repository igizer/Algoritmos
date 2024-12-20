#Ejecutar consola con Cntrl+Alt+N




#----Variables----
texto = """Hola soy Sergio y sé sumar :P
a+b="""

a = 3

b = 2

c = a + b

print(texto, c) # c va a valer 5


print("-----------------------------------------------")

operacion = 3 + 2

print("sino puedo hacerlo de esta forma también:", operacion)

print("-----------------------------------------------")

#Otra cosa es que siempre se pone el último valor que asume la variable
var= 2
var= "ya no es un 2, ahora es esta propia aclaración"
print(var)

print("Pero si en vez de reemplazar, quiero sumar, puedo hacer esto con un +=")
print("suma=a, suma+=b, debe dar 5, verificamos abajo")
print()
suma= a
suma+= b
print(suma)