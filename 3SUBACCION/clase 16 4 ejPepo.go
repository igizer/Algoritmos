//Un importante supermercado del país, le solicita desarrollar un algoritmo que asista en el
//manejo de sus líneas de cajas. El programa registra cada compra realizada listando al final el
//ticket de compra. Será utilizado por un cajero y el ingreso de datos se hará de manera 
//interactiva, pudiendo el cajero atender varios clientes y cerrar la caja en cualquier
//momento (una vez terminado de atender un cliente).
Por primera vez, se solicitan los siguientes datos:
Nro de Caja
Apellido y nombre del cajero.
Luego por cada arribo de un cliente a la caja, se le solicitará (también por unica vez):
Dni del cliente
Cantidad de productos (distintos) a comprar
Por cada producto que compra el usuario se deberá ingresar el Código de barras (alfanumérico) correspondiente, y las unidades que lleva para dicho producto.

Sub Acciones ya implementadas:
obtener_precio_producto(codigo_barras) Brinda el precio unitario del producto.
es_socio(dni) Permite determinar si el cliente pertenece al programa de fidelización del supermercado de "clientes frecuentes". Los resultados posibles son: ORO | PLATA | DIAMANTE | NO SOCIO
obtener_fecha_hora() Brinda la fecha y hora del sistema

