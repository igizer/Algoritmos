NIVEL 1

accion pincho_vago es
Ambiente
    Paises: secuencia de caracter
    Edades: secuencia de entero
    P: caracter
    E: entero
    Cant_jugadores: entero
    Acumul_edades: entero
    Promedio_edades: real
    May, Men: entero
Proceso
    Arr(Paises); Avz(Paises, P)
    Arr(Edades); Avz(Edades, E)
    May:=0; Men:=0
    Cant_jugadores:=0
    Acumul_edades:=0
    Promedio_edades:=0
    Mientras NFDS(Paises) hacer
        Mientras P<> "#" hacer
            Avz(Paises, P)
        FinMientras
        Avz(Paises, P); Cant_jugadores:= Cant_jugadores+26
        Para i:1 a 26 hacer
            Acumul_edades:= Acumul_edades+E
            Si E>May entonces
                May:=E
            FinSi
            Si E<Men entonces
                Men:=E
            FinSi
            Avz(Edades, E)
        FinPara
    FinMientras
    Promedio_edades:= Acumul_edades/Cant_jugadores
    Escribir("El promedio de edades de jugadores es de", Promedio_edades)
    Escribir("El jugador mas viejo tiene", May, "y el mas joven", Men)
    Cerrar(Paises)
    Cerrar(Edades)
Finaccion


Accion pincho_vago2 es
Ambiente
    Jugadores=registro
        Seleccion: AN(20)
        Posicion: ("ARQ", "DEF", "MED", "DEL", "CT")
        Dorsal: N(2)
        ApyNom: AN(30)
        Equipo: AN(20)
    FinRegistro

    arch: archivo de Jugadores ordenado por Seleccion, Posicion, Dorsal
    reg: Jugadores
    resg_seleccion: AN(20)
    resg_posicion: ("ARQ", "DEF", "MED", "DEL", "CT")
    Tot_general, Tot_seleccion, Tot_posicion: entero

    Procedimiento Corte_posicion() es
        Escribir("En la seleccion", resg_seleccion, "hay", Tot_posicion, "jugadores en la posicion", resg_posicion)
        Tot_seleccion:= Tot_seleccion + Tot_posicion
        Tot_posicion:=0
        resg_posicion:= reg.Posicion
    FinProcedimiento

    Procedimiento Corte_seleccion() es
        Corte_posicion()
        Escribir("En la seleccion", resg_seleccion, "hay", Tot_seleccion, "jugadores")
        Tot_general:= Tot_general + Tot_seleccion
        Tot_seleccion:= 0
        resg_seleccion:= reg.Seleccion
    FinProcedimiento

Proceso
    Abrir E/(arch)
    Leer(arch, Reg)
    resg_seleccion:= reg.Seleccion
    resg_posicion:= reg.Posicion
    Tot_general:=0; Tot_seleccion:=0; Tot_posicion:=0
    Mientras NFDA(arch) hacer
        Si resg_seleccion <> reg.Seleccion entonces
            Corte_seleccion()
        SiNo
            Si resg_posicion <> reg.Posicion entonces
                Corte_posicion()
            FinSi
        FinSi
        Tot_posicion:= Tot_posicion + 1
        Leer(arch, reg)
    FinMientras
    Corte_seleccion()
    Escribir("En total hay", Tot_general, "jugadores")
    Cerrar(arch)
Finaccion

NIVEL 2

Accion pincho_vago3 es
Ambiente
    Paises: secuencia de caracter
    Edades: secuencia de entero
    P: caracter
    E: entero
    Cant_jugadores_dorsal: entero
    Acumul_edades_dorsal: entero
    Promedio_edad_dorsal: real
    Dorsal_usuario: entero
    May_equipo, Men_equipo: entero
Proceso
    Arr(Paises); Avz(Paises, P)
    Arr(Edades); Avz(Edades, E)
    May_equipo:=0; Men_equipo:=0
    Cant_jugadores_dorsal:=0
    Acumul_edades_dorsal:=0
    Promedio_edad_dorsal:=0
    Escribir("Ingrese una dorsal")
    Leer(Dorsal_usuario)
    Mientras NFDS(Paises) hacer
        Mientras P<> "#" hacer
            Avz(Paises, P)
        FinMientras
        Avz(Paises, P);
        Para i:1 a 26 hacer
            Si E = Dorsal_usuario entonces
                Cant_jugadores_dorsal:= Cant_jugadores_dorsal + 1
                Avz(Edades, E)
                Acumul_edades_dorsal:= Acumul_edades_dorsal + E
            SiNo
                Avz(Edades, E)
            FinSi                
            Si E>May_equipo entonces
                May_equipo:=E
            FinSi
            Si E<Men_equipo entonces
                Men_equipo:=E
            FinSi
            Avz(Edades, E)
        FinPara
        Escribir("El jugador mas viejo tiene", May_equipo)
        Escribir("El jugador mas joven tiene", Men_equipo)
        May_equipo:=0
        Men_equipo:=0
    FinMientras
    Promedio_edad_dorsal:= Acumul_edades_dorsal/Cant_jugadores_dorsal
    Escribir("El promedio de edades de jugadores de la dorsal",
            Dorsal_usuario, "es de", Promedio_edad_dorsal)
    Cerrar(Paises)
    Cerrar(Edades)
Finaccion

Accion pincho_vago4 es
Ambiente
    Jugadores=registro
        Seleccion: AN(20)
        Posicion: ("ARQ", "DEF", "MED", "DEL", "CT")
        Dorsal: N(2)
        ApyNom: AN(30)
        Edad: N(3)
        Equipo: AN(20)
    FinRegistro

    arch: archivo de Jugadores ordenado por Seleccion, Posicion, Dorsal
    reg: Jugadores
    resg_seleccion: AN(20)
    resg_posicion: ("ARQ", "DEF", "MED", "DEL", "CT")
    Tot_general, Tot_seleccion, Tot_posicion: entero
    Acumul_edades_seleccion: N(5)
    Cant_jugadores_seleccion: N(2)

    Procedimiento Corte_posicion() es
        Escribir("En la seleccion", resg_seleccion, "hay", Tot_posicion, "jugadores en la posicion", resg_posicion)
        Tot_seleccion:= Tot_seleccion + Tot_posicion
        Tot_posicion:=0
        resg_posicion:= reg.Posicion
    FinProcedimiento

    Procedimiento Corte_seleccion() es
        Corte_posicion()
        Escribir("En la seleccion", resg_seleccion, "hay", Tot_seleccion, "jugadores")
        Escribir("El promedio de edades de la seleccion", resg_seleccion,
                "es de", Acumul_edades_seleccion/Cant_jugadores_seleccion)
        Tot_general:= Tot_general + Tot_seleccion
        Tot_seleccion:= 0
        resg_seleccion:= reg.Seleccion
    FinProcedimiento

Proceso
    Abrir E/(arch)
    Leer(arch, Reg)
    resg_seleccion:= reg.Seleccion
    resg_posicion:= reg.Posicion
    Tot_general:=0; Tot_seleccion:=0; Tot_posicion:=0
    Cant_jugadores_seleccion:=0
    Acumul_edades_seleccion:=0
    Mientras NFDA(arch) hacer
        Si resg_seleccion <> reg.Seleccion entonces
            Corte_seleccion()
        SiNo
            Si resg_posicion <> reg.Posicion entonces
                Corte_posicion()
            FinSi
        FinSi
        
        Si reg.Posicion <> "CT" entonces
            Tot_posicion:= Tot_posicion + 1
        FinSi
        Cant_jugadores_seleccion:= Cant_jugadores_seleccion + 1
        Acumul_edades_seleccion:= Acumul_edades_seleccion + reg.Edad
        Leer(arch, reg)
    FinMientras
    Corte_seleccion()
    Escribir("En total hay", Tot_general, "jugadores")
    Cerrar(arch)
Finaccion

NIVEL 3

