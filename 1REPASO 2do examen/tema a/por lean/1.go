ACCION temab ES 
    AMBIENTE
        CLAVEL = REGISTRO 
            id_sucursal: N(5)
            id_cliente: N(6)
        FINREGISTRO 
        FECHA = REGISTRO 
            aa: N(4)
            mm: 1..12
            dd: 1..31
        FINREGISTRO
        CLIENTES = REGISTRO 
            clave: CLAVEL
            nom_y_ape: AN(30)
            saldo: N(6)
            fecha_alta: FECHA
            fecha_baja: FECHA
        FINREGISTRO
        MOVIMIENTOS = REGISTRO 
            clave: CLAVEL
            cod_movimiento: 0..99 
            nom_y_ape: AN(30)
            fecha_mov: FECHA
            monto: N(6)
            detalle: AN(30)
            categoria: 1..6
            tipo = ("I", "E")
        FINREGISTRO
        contador: entero 
        reg_mae, reg_aux, reg_sal: CLIENTES 
        reg_mov: MOVIMIENTOS
        MAE, SAL: ARCHIVO DE CLIENTES ORDENADO POR clave
        MOV: ARCHIVO DE MOVIMIENTOS ORDENADO POR clave y cod_movimiento
        PROCEDIMIENTO leer_mae() ES 
            LEER(MAE, reg_mae)
            SI FDA(MAE) ENTONCES 
                reg_mae.clave := HV
            FINSI
        FINPROCEDIMIENTO 
        PROCEDIMIENTO leer_mov() ES 
            LEER(MOV, reg_mov)
            SI FDA(MOV) ENTONCES 
                reg_mov.clave := HV
            FINSI
        FINPROCEDIMIENTO 
        PROCESO 
            contador:= 0
            ABRIR E/(MOV); ABRIR(MAE); ABRIR S/(SAL)
            leer_mae(); leer_mov()
            MIENTRAS (reg_mae.clave <> HV ) O (reg_mov.clave <> HV) HACER
                SI reg_mae.clave < reg_mov.clave ENTONCES
                    GRABAR(SAL, reg_mae)
                    leer_mae()
                SINO 
                    SI reg_mae.clave = reg_mov.clave ENTONCES
                        SI reg_mov.cod_movimiento = 0 ENTONCES
                            ESCRIBIR("ERROR")
                        FINSI
                        reg_aux:= reg_mae
                        MIENTRAS reg_mov.clave = reg_aux.clave HACER 
                            SI reg_mov.cod_movimiento > 0 y reg_mov.cod_movimiento < 99 ENTONCES 
                                SEGUN reg_mov.tipo HACER 
                                    "I": reg_aux.saldo:= reg_aux.saldo + reg_mov.monto
                                    "E": reg_aux.saldo:= reg_aux.saldo - reg_mov.monto 
                                FINSEGUN
                                reg_sal:= reg_aux
                                GRABAR(SAL, reg_sal)
                                leer_mov()
                            FINSI 
                            SI reg_mov.cod_movimiento = 99 ENTONCES
                                reg_aux.fecha_baja:= reg_mov.fecha_mov
                                reg_sal:= reg_aux
                                GRABAR(SAL, reg_sal)
                                leer_mov()
                                contador:= contador + 1
                            FINSI 
                        FINMIENTRAS 
                        leer_mae()
                    SINO 
                        reg_aux:= reg_mae
                        SI reg_aux.clave > reg_mov.clave ENTONCES 
                            SI reg_mov.cod_movimiento > 0 ENTONCES
                                ESCRIBIR("ERROR")
                            FINSI
                            reg_aux.clave:= reg_mov.clave 
                            reg_aux.nom_y_ape:= reg_mov.nom_y_ape
                            reg_aux.saldo:= reg_mov.monto
                            reg_aux.fecha_alta:= reg_mov.fecha_mov
                            reg_aux.fecha_baja.aa:= ""
                            reg_aux.fecha_baja.mm:= ""
                            reg_aux.fecha_baja.dd:= ""
                            reg_sal:= reg_aux
                            GRABAR(SAL, reg_sal)
                            leer_mov()
                            MIENTRAS reg_mov.clave = reg_aux.clave HACER 
                                SI reg_mov.cod_movimiento > 0 y reg_mov.cod_movimiento < 99 ENTONCES 
                                    SEGUN reg_mov.tipo HACER 
                                        "I": reg_aux.saldo:= reg_aux.saldo + reg_mov.monto
                                        "E": reg_aux.saldo:= reg_aux.saldo - reg_mov.monto 
                                    FINSEGUN
                                    reg_sal:= reg_aux
                                    GRABAR(SAL, reg_sal)
                                    leer_mov()
                                FINSI 
                                SI reg_mov.cod_movimiento = 99 ENTONCES
                                    reg_aux.fecha_baja:= reg_mov.fecha_mov
                                    reg_sal:= reg_aux
                                    GRABAR(SAL, reg_sal)
                                    leer_mov()
                                    contador:= contador + 1
                                FINSI 
                            FINMIENTRAS 
                            leer_mae()
                        FINSI
                    FINSI
                FINSI
            FINMIENTRAS
            ESCRIBIR("se dieron ", contador, "bajas")
            CERRAR(MAE)
            CERRAR(MOV)
            CERRAR(SAL)
FINACCION