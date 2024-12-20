El sector Farmacia de un importante Hospital cuenta con un archivo secuencial con datos de todos los medicamentos que posee:
STOCK_MEDICAMENTOS (ordenado por Categoría, NroLote,  IdMedicamento)
|Categoría|NroLote|IdMedicamento|/Stock/FechaVencimiento/

En donde el campo categoría indica:
	“I”: Presumiblemente seguro
	“II”: Produce efectos adversos, leves o moderados.
	“III”: Produce efectos adversos, graves o potencialmente peligrosos.

Se desea mostrar un reporte con las estadísticas de cantidad de medicamentos próximos a vencer por categorías de la siguiente manera:
TOTAL CATEGORÍA I PROXIMOS A VENCER: xxxxxxx
TOTAL CATEGORÍA I NO PRÓXIMOS A VENCER: xxxxxxx


TOTAL CATEGORÍA II PROXIMOS A VENCER: xxxxxxx
TOTAL CATEGORÍA II NO PRÓXIMOS A VENCER: xxxxxxx


TOTAL CATEGORÍA III PRÓXIMOS A VENCER: xxxxxxx
TOTAL CATEGORÍA III NO PROXIMOS A VENCER: xxxxxxx


TOTAL MEDICAMENTOS PRÓXIMOS A VENCER: xxxxxxx
TOTAL MEDICAMENTOS NO PRÓXIMOS A VENCER: xxxxxxx


TOTAL DE MEDICAMENTOS: xxxxxxx
Para determinar los medicamentos próximos a vencer , existe una función proximoAVencer() que recibe como parámetro un número de Lote, e indica Verdadero o Falso si es que el mismo se encuentra próximo a vencer.

ACCION Corte3 ES
	AMBIENTE
		Stock = REGISTRO
			Categoría : ("I", "II", "III")
			NroLote : AN(20)
			IdMedicamento : N(15)
			Stock : N(3)
			FechaVencimiento = REGISTRO
				AAAA: N(4)
				DD: (1..31)
				MM: (1..12)
		FINREGISTRO
		STOCK_MEDICAMENTOS:= Archivo de Stock Ordenado por Categoría, NroLote, IdMedicamento
		Reg_Stock:= Stock
		
		Resg_Categoría: ("I", "II", "III")
		TotCatProx: entero
		TotCatNoProx: entero
		

		Procedimiento Corte_Categoría ES
			ESC("TOTAL CATEGORÍA", Resg_Categoría, "PROXIMOS A VENCER:" TotCatProx)
			ESC("TOTAL CATEGORÍA", Resg_Categoría, "NO PRÓXIMOS A VENCER:" TotCatNoProx)
			Resg_Categoría:= Reg_Stock.Categoría
			TotMediProx:= TotMediProx + TotCatProx
			TotMediNoProx:= TotMediNoProx + TotCatNoProx
			TotCatProx:= 0
			TotCatNoProx:= 0





		
		ABRIR E/(STOCK_MEDICAMENTOS)
		Leer(STOCK_MEDICAMENTOS, Reg_Stock)
		Resg_Categoría:= Reg_Stock.Categoría
		TotMediProx:= 0
		TotMediNoProx:= 0
		MIENTRAS NFDA(STOCK_MEDICAMENTOS) HACER
			SI Reg_Stock.Categoría <> Resg_Categoría
				Corte_Categoría
			FINSI
			SI proximoAVencer(NroLote) = v
				TotCatProx:= TotCatProx + 1
			SINO
				TotCatNoProx:= TotCatNoProx + 1
			FINSI
			Leer(STOCK_MEDICAMENTOS, Reg_Stock)
		FINMIENTRAS

		ESC("TOTAL MEDICAMENTOS PRÓXIMOS A VENCER: ", TotMediProx)
		ESC("TOTAL MEDICAMENTOS NO PRÓXIMOS A VENCER: ", TotMediNoProx)