#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <conio.h>
#include <ctype.h>

// Acá efinimos el tamaño del tablero y otros parámetros importantes
#define TAMANIO_TABLERO 10 // El tablero es de 10x10
#define NUM_BARCOS 5 // Cada jugador tiene 5 barcos

// Función para limpiar la pantalla (esto simula un "refresco" en consola)
void limpiarPantalla() {
    system("cls");
}

// Función para limpiar el buffer de entrada (útil después de usar scanf)
void limpiarBufferEntrada() {
    int c;
    while ((c = getchar()) != '\n' && c != EOF); // Consume todo hasta el siguiente salto de línea
}

typedef struct {
    int x, y; // Las coordenadas iniciales del barco
    int longitud; // La longitud del barco (número de casillas que ocupa)
    int direccion; // Dirección: 0=ARRIBA, 1=DERECHA, 2=ABAJO, 3=IZQUIERDA
} Barco;

// Ahora definimos los tableros de ambos jugadores como matrices 2D
char tableroJugador1[TAMANIO_TABLERO][TAMANIO_TABLERO];
char tableroJugador2[TAMANIO_TABLERO][TAMANIO_TABLERO];

// Estructura para los jugadores en el ranking
typedef struct NodoJugador {
    char nombre[50]; // Nombre del jugador
    int tirosMinimos; // Récord de tiros mínimos necesarios para ganar
    int victorias; // Cantidad de partidas ganadas
    struct NodoJugador *siguiente; // Puntero al siguiente jugador en la lista
} NodoJugador;

NodoJugador *ranking = NULL; // La lista enlazada empieza vacía

// Crea un nuevo jugador y lo devuelve como un nodo
NodoJugador *crearJugador(char *nombre, int tiros) {
    NodoJugador *nuevo = (NodoJugador *)malloc(sizeof(NodoJugador));
    strcpy(nuevo->nombre, nombre); // Copiamos el nombre del jugador
    nuevo->tirosMinimos = tiros; // Guardamos el récord inicial
    nuevo->victorias = 1; // El jugador empieza con 1 victoria (sino claramente no estaría en el ranking)
    nuevo->siguiente = NULL; // No tiene siguiente por ahora
    return nuevo;
}

// Declaraciones de funciones que se usan en otras partes del código para el ranking
void guardarRanking();
void ordenarRanking();
void actualizarRanking(char *nombre, int tiros);

// Limpia la memoria usada por la lista del ranking
void liberarRanking() {
    NodoJugador *actual = ranking;
    while (actual != NULL) {
        NodoJugador *temp = actual;
        actual = actual->siguiente;
        free(temp); // Libera cada nodo individualmente
    }
    ranking = NULL; // Libera cada nodo individualmente
}


// Ordena el ranking según el criterio dado (que serían victorias o tiros mínimos)
typedef enum { ORDEN_VICTORIAS, ORDEN_TIROS } CriterioOrden;

void ordenarRanking(CriterioOrden criterio) {
    if (ranking == NULL || ranking->siguiente == NULL) return;

    NodoJugador *actual, *siguiente;
    int cambiado;

    // Este algoritmo implementa una burbuja optimizada para listas enlazadas, o sea esto reduce el número de comparaciones innecesarias
    do {
        cambiado = 0;
        actual = ranking;

        while (actual->siguiente != NULL) {
            siguiente = actual->siguiente;

            int debeIntercambiar = 0;
            if (criterio == ORDEN_VICTORIAS && 
                (actual->victorias < siguiente->victorias ||
                 (actual->victorias == siguiente->victorias &&
                  actual->tirosMinimos > siguiente->tirosMinimos))) {
                debeIntercambiar = 1;
            } else if (criterio == ORDEN_TIROS && actual->tirosMinimos > siguiente->tirosMinimos) {
                debeIntercambiar = 1;
            }

            if (debeIntercambiar) {
                // Intercambiamos los datos de los nodos en lugar de moverlos, de esta forma simplificamos un poco el ordenamiento y más sencillo
                NodoJugador temp = *actual;
                actual->victorias = siguiente->victorias;
                actual->tirosMinimos = siguiente->tirosMinimos;
                strcpy(actual->nombre, siguiente->nombre);

                siguiente->victorias = temp.victorias;
                siguiente->tirosMinimos = temp.tirosMinimos;
                strcpy(siguiente->nombre, temp.nombre);

                cambiado = 1; // Así indicamos que hubo un cambio para luego-
            }
            actual = actual->siguiente; // -para luego pasar al siguiente nodo
        }
    } while (cambiado); // y repetimos mientras haya más cambios
}

// Actualiza el ranking: añade un nuevo jugador o mejora los datos de uno existente
void actualizarRanking(char *nombre, int tiros) {
    NodoJugador *actual = ranking, *anterior = NULL;

    // Buscamos si el jugador ya está en el ranking, claramente si ganó alguna vez
    while (actual != NULL) {
        if (strcmp(actual->nombre, nombre) == 0) {
            // Si ya está en el ranking, actualizamos sus estadísticas
            if (tiros < actual->tirosMinimos || actual->tirosMinimos == 0)
                actual->tirosMinimos = tiros; // Actualiza si logra un récord mejor
            actual->victorias++; // Suma una victoria más
            return;
        }
        anterior = actual;
        actual = actual->siguiente;
    }

    // Si no encontramos al jugador, lo añadimos como nuevo
    NodoJugador *nuevo = crearJugador(nombre, tiros);
    if (anterior == NULL) {
        ranking = nuevo; // Es el primer elemento de la lista
    } else {
        anterior->siguiente = nuevo; // Lo añadimos al final de la lista
    }
}

// Función que guarda el ranking en un archivo de texto
void guardarRanking() {
    FILE *archivo = fopen("ranking.txt", "w");
    if (archivo) {
        NodoJugador *actual = ranking; // Empezamos desde el inicio de la lista
        while (actual) {
            // Escribimos cada jugador en el archivo con formato CSV
            fprintf(archivo, "%s,%d,%d\n", actual->nombre, actual->tirosMinimos, actual->victorias);
            actual = actual->siguiente; // Ahora pasamos al siguiente nodo
        }
        fclose(archivo);
    } else {
        printf("No se pudo abrir el archivo para guardar el ranking.\n");
    }
}
 

// Función para cargar el ranking desde un archivo
void cargarRanking() {
    FILE *archivo = fopen("ranking.txt", "r"); // De esta forma intentamos abrir el archivo de ranking
    if (archivo) {
        char nombre[50];
        int tirosMinimos, victorias;

        liberarRanking(); // Limpiamos cualquier ranking anterior cargado

        // Leemos el archivo línea por línea y cargamos los datos
        while (fscanf(archivo, "%[^,],%d,%d\n", nombre, &tirosMinimos, &victorias) == 3) {
            // Creamos un nuevo jugador con los datos leídos
            NodoJugador *nuevo = (NodoJugador *)malloc(sizeof(NodoJugador));
            strcpy(nuevo->nombre, nombre);
            nuevo->tirosMinimos = tirosMinimos;
            nuevo->victorias = victorias;
            nuevo->siguiente = NULL;

            // Insertamos el jugador al final de la lista
            if (ranking == NULL) {
                ranking = nuevo; // Si la lista está vacía, el nuevo jugador será el primero
            } else {
                NodoJugador *actual = ranking;
                // Buscamos el final de la lista
                while (actual->siguiente != NULL) {
                    actual = actual->siguiente; // Avanza hasta el último nodo
                }
                actual->siguiente = nuevo; // Insertamos al final de la lista
            }
        }
        fclose(archivo); // Ahora cerramos el archivo una vez terminado
    } else {
        printf("No se pudo abrir el archivo para cargar el ranking.\n");
    }
}

// Función para contar la cantidad de jugadores en el ranking (util para ordenar en ranking)
int contarJugadores(NodoJugador *ranking) {
    int contador = 0;
    NodoJugador *actual = ranking;
    while (actual) {
        contador++;
        actual = actual->siguiente;
    }
    return contador;
}

// Función para mostrar el ranking en consola
void mostrarRanking(NodoJugador *ranking) {
    limpiarPantalla(); // Limpiamos la pantalla para mostrar el ranking y que no se mezcle con el manú
    printf("===================================\n");
    printf("         Ranking de Jugadores      \n");
    printf("===================================\n");

    int totalJugadores = contarJugadores(ranking); // Calculamos el total de jugadores en el ranking

    if (totalJugadores == 0) {
        printf("Aún no hay jugadores en el ranking.\n");
    } else {
        NodoJugador *actual = ranking;
        int posicion = 1;
        // Recorremos la lista e imprimimos la posición de cada jugador
        while (actual) {
            printf("%d. %s - Tiros mínimos: %d - Victorias: %d\n",
                   posicion++, actual->nombre, actual->tirosMinimos, actual->victorias);
            actual = actual->siguiente;
        }
    }

    printf("\nPresiona cualquier tecla para continuar...\n");
    getch(); // Esperamos que el usuario presione una tecla para continuar, más que nada con el fin de crear una pausa que permita leer el ranking
    limpiarPantalla();
}

// Función para limpiar el ranking
void limpiarRanking() {
    limpiarPantalla();
    liberarRanking(); // Liberamos la memoria ocupada por la lista enlazada de jugadores
    
    if (remove("ranking.txt") == 0) {
        printf("El ranking ha sido limpiado exitosamente.\n");
    } else {
        printf("No se encontró un archivo de ranking para eliminar.\n");
    }

    printf("\nPresiona cualquier tecla para continuar...\n");
    getch(); //Justamente para que presione una tecla otra vez
}

// Esta función muestra el menú principal ni bien iniciamos el juego.
// Se encarga de limpiar la pantalla y luego mostrar las opciones disponibles para el usuario,
// como comenzar el juego, ver el ranking, limpiar el ranking, ver instrucciones, ver los créditos o salir.
void mostrarMenu() {
    limpiarPantalla();
    printf("===================================\n");
    printf("     Bienvenido a Batalla Naval    \n");
    printf("===================================\n");
    printf("1- Comenzar el juego\n");
    printf("2- Ranking de Jugadores\n");
    printf("3- Limpiar Ranking\n"); 
    printf("4- Instrucciones\n");
    printf("5- Créditos\n");
    printf("6- Salir\n");
    printf("Elija una opción: "); // Acá nos referimos a algún número del 1 al 6
}

// Muestra las instrucciones
void mostrarInstrucciones() {
    limpiarPantalla();
    printf("===================================\n");
    printf("           Instrucciones           \n");
    printf("===================================\n");
    printf("1. Cada jugador tiene un tablero 10x10.\n");
    printf("2. Al inicio, cada jugador debe colocar\n");
    printf("   sus barcos en el tablero.\n");
    printf("   - Usa las teclas W/A/S/D para mover\n");
    printf("     los barcos.\n");
    printf("   - Usa 'R' para rotar los barcos.\n");
    printf("   - Usa 'C' para confirmar su posición.\n");
    printf("3. Durante la partida, los jugadores se\n");
    printf("   alternan turnos para atacar.\n");
    printf("   - Introduce coordenadas (Ej: B5).\n");
    printf("   - 'X' indica un impacto y 'O' un fallo.\n");
    printf("4. Gana el jugador que hunda todos los\n");
    printf("   barcos del oponente primero.\n");
    printf("\nPresiona cualquier tecla para volver al menú.\n");
    getch();
}

// Muestra los créditos
void mostrarCreditos() {
    limpiarPantalla();
    printf("===================================\n");
    printf("              Créditos             \n");
    printf("===================================\n");
    printf("Creado por:\n");
    printf("  - Scherf, Sergio Agustín de Jesús\n");
    printf("  - Navarro, Fabio Ignacio\n");
    printf("\nEstudiantes de la Universidad Tecnológica\n");
    printf("Nacional, Cátedra de Algoritmos y\n");
    printf("Estructuras de Datos, Comisión K1.6.\n");
    printf("\nPresiona cualquier tecla para volver al menú.\n");
    getch();
}

// Función principal del menú (justamente según la tecla que presiona en el menú inicial, del 1 al 6)
int menuPrincipal() {
    char opcion;
    do {
        mostrarMenu();
        opcion = getch(); // Leer opción del usuario
        switch (opcion) {
            case '1':
                return 1; // Comenzar juego
            case '2':
                mostrarRanking(ranking); // Mostrar el ranking
                break;
            case '3':
                limpiarRanking(); // Limpiar el ranking
                break;
            case '4':
                mostrarInstrucciones();
                break;
            case '5':
                mostrarCreditos();
                break;
            case '6':
                return 6; // Salir
                break;
            default:
                printf("Opción inválida. Inténtalo nuevamente.\n");
                getch();
                break;
        }
    } while (1);
}

// A continuación tenemos funciones dedicadas al tablero/matriz de la partida

void inicializarTablero(char tablero[TAMANIO_TABLERO][TAMANIO_TABLERO]) {
    for (int i = 0; i < TAMANIO_TABLERO; i++) {
        for (int j = 0; j < TAMANIO_TABLERO; j++) {
            tablero[i][j] = '~'; // Elegimos este signo porque creemos que representa como las olas del agua
        }
    }
}

void dibujarTablero(char tablero[TAMANIO_TABLERO][TAMANIO_TABLERO], int ocultarBarcos) {
    // Imprimir cabecera (letras A-J)
    printf("   "); // Las columnas quedaban fuera de posición, con estos espacios se acomodan a las casillas
    for (int i = 0; i < TAMANIO_TABLERO; i++) {
        printf(" %c ", 'A' + i); // Para organizar e imprimir las columnas (de la A hasta la J)
    }
    printf("\n");

    // Imprimir las filas con números alineados
    for (int i = 0; i < TAMANIO_TABLERO; i++) {
        printf("%2d ", i + 1); // Para organizar e imprimir las filas (del 1 hasta al 10)
        for (int j = 0; j < TAMANIO_TABLERO; j++) {
            if (ocultarBarcos && tablero[i][j] == 'B') {
                printf(" ~ ");
            } else {
                printf(" %c ", tablero[i][j]); // De lo contrario, mostramos el símbolo real en la casilla.
            }
        }
        printf("\n");
    }
}

// Necesitamos distinguir los barcos colocados con los que vamos a colocar/acomodar, por esto los P son los temporales
void mostrarBarcoTemporal(Barco barco, char tablero[TAMANIO_TABLERO][TAMANIO_TABLERO]) {
    for (int i = 0; i < barco.longitud; i++) {
        int x = barco.x, y = barco.y;
        // Calculamos la posición de cada parte del barco dependiendo de su dirección.
        if (barco.direccion == 0) y -= i; // Dirección 0: vertical hacia arriba.
        if (barco.direccion == 1) x += i; // Dirección 1: horizontal hacia la derecha.
        if (barco.direccion == 2) y += i; // Dirección 2: vertical hacia abajo.
        if (barco.direccion == 3) x -= i; // Dirección 3: horizontal hacia la izquierda.

        // Si las coordenadas son válidas, marcamos la casilla temporalmente con P copmo dijimos arriba
        if (x >= 0 && x < TAMANIO_TABLERO && y >= 0 && y < TAMANIO_TABLERO && tablero[y][x] == '~')
            tablero[y][x] = 'P';
    }
}

// En general, esto es para una vez que movemos el barco, la P ya no debería estar marcando que está acá, por ende vuelve a ser con ~
void limpiarBarcoTemporal(Barco barco, char tablero[TAMANIO_TABLERO][TAMANIO_TABLERO]) {
    for (int i = 0; i < barco.longitud; i++) {
        int x = barco.x, y = barco.y;
        // Calculamos la posición
        if (barco.direccion == 0) y -= i;
        if (barco.direccion == 1) x += i;
        if (barco.direccion == 2) y += i;
        if (barco.direccion == 3) x -= i;

        // Si encontramos una marca temporal 'P', la limpiamos, volviendo la casilla a '~'.
        if (x >= 0 && x < TAMANIO_TABLERO && y >= 0 && y < TAMANIO_TABLERO && tablero[y][x] == 'P')
            tablero[y][x] = '~';
    }
}

// Prohibe que los barcos se ubiquen por fuera de la matriz 10x10
int validarColocacion(Barco barco, char tablero[TAMANIO_TABLERO][TAMANIO_TABLERO]) {
    for (int i = 0; i < barco.longitud; i++) {
        int x = barco.x, y = barco.y;
        if (barco.direccion == 0) y -= i;
        if (barco.direccion == 1) x += i;
        if (barco.direccion == 2) y += i;
        if (barco.direccion == 3) x -= i;

        // Si alguna parte del barco se sale del tablero o hay un barco en esa posición, la colocación no es válida.
        if (x < 0 || x >= TAMANIO_TABLERO || y < 0 || y >= TAMANIO_TABLERO || tablero[y][x] != '~')
            return 0; // Retorna 0 (falso) si la colocación no es válida.
    }
    return 1; // Retorna 1 (verdadero) si la colocación es válida.
}

// Esta función coloca definitivamente un barco en el tablero una vez que su colocación ha sido validada, Cambiando las letras a B
void colocarBarco(Barco barco, char tablero[TAMANIO_TABLERO][TAMANIO_TABLERO]) {
    for (int i = 0; i < barco.longitud; i++) {
        int x = barco.x, y = barco.y;
        if (barco.direccion == 0) y -= i;
        if (barco.direccion == 1) x += i;
        if (barco.direccion == 2) y += i;
        if (barco.direccion == 3) x -= i;

        if (x >= 0 && x < TAMANIO_TABLERO && y >= 0 && y < TAMANIO_TABLERO)
            tablero[y][x] = 'B';
    }
}

// Esta función es utilizada para que el jugador coloque sus barcos en el tablero.
void colocarBarcosJugador(char tablero[TAMANIO_TABLERO][TAMANIO_TABLERO], char *nombreJugador) {
    Barco barcos[NUM_BARCOS] = {
        {0, 0, 5, 1}, {0, 0, 4, 1}, {0, 0, 3, 1}, {0, 0, 3, 1}, {0, 0, 2, 1}
    };

    printf("%s, es tu turno de colocar los barcos.\n", nombreJugador);
    for (int i = 0; i < NUM_BARCOS; i++) {
        int colocado = 0;
        while (!colocado) {
            limpiarPantalla();
            printf("Colocando barco #%d (longitud: %d)\n", i + 1, barcos[i].longitud);

            mostrarBarcoTemporal(barcos[i], tablero);
            dibujarTablero(tablero, 0);
            limpiarBarcoTemporal(barcos[i], tablero);

            printf("Usa W/A/S/D para mover, R para rotar, C para confirmar.\n");

            char tecla = getch();
            if (tecla == 'w' && barcos[i].y > 0) barcos[i].y--;
            if (tecla == 's' && barcos[i].y < TAMANIO_TABLERO - 1) barcos[i].y++;
            if (tecla == 'a' && barcos[i].x > 0) barcos[i].x--;
            if (tecla == 'd' && barcos[i].x < TAMANIO_TABLERO - 1) barcos[i].x++;
            if (tecla == 'r') barcos[i].direccion = (barcos[i].direccion + 1) % 4;
            if (tecla == 'c' && validarColocacion(barcos[i], tablero)) {
                colocarBarco(barcos[i], tablero);
                colocado = 1;
            } else if (tecla == 'c') {
                printf("No puedes colocar el barco aquí.\n");
                getch();
            }
        }
    }
}

int validarAtaque(int x, int y, char tablero[TAMANIO_TABLERO][TAMANIO_TABLERO]) {
    return x >= 0 && x < TAMANIO_TABLERO && y >= 0 && y < TAMANIO_TABLERO && tablero[y][x] != 'O' && tablero[y][x] != 'X';
}

// Función que verifica si un jugador ha ganado el juego, cuando un jugador ya no tenga casillas de barco
int verificarVictoria(char tablero[TAMANIO_TABLERO][TAMANIO_TABLERO]) {
    for (int i = 0; i < TAMANIO_TABLERO; i++) {
        for (int j = 0; j < TAMANIO_TABLERO; j++) {
            if (tablero[i][j] == 'B') return 0; // a través de una busqueda corriendo el tablero encuentra una B, el juego sigue
        }
    }
    return 1;
}

// Función para obtener las coordenadas de ataque del jugador. 
// El jugador introduce la columna (A-J) y la fila (1-10) para realizar un ataque.
void obtenerCoordenadasAtaque(int *columna, int *fila) {
    char c;
    while (1) {
        printf("Introduce la columna (A-J) y fila (1-10) para atacar: ");
        scanf(" %c%d", &c, fila);
        *columna = toupper(c) - 'A';
        (*fila)--;

        if (*columna >= 0 && *columna < TAMANIO_TABLERO && *fila >= 0 && *fila < TAMANIO_TABLERO) {
            break;
        } else {
            printf("Coordenadas inválidas, inténtalo nuevamente.\n");
        }
    }
}

// Función jugar, es la principal del juego
// Controla el turno de cada jugador, muestra los tableros, obtiene los ataques y determina la victoria.
void jugar(char tablero1[TAMANIO_TABLERO][TAMANIO_TABLERO], char tablero2[TAMANIO_TABLERO][TAMANIO_TABLERO], char *nombreJugador1, char *nombreJugador2) {
    int turno = 1; // Comienza con el turno del jugador 1
    int tirosJugador1 = 0, tirosJugador2 = 0; // Contadores de tiros para cada jugador

    while (1) {
        int turnoExtra = 1; // Variable que controla si el jugador sigue teniendo turnos extra

        while (turnoExtra) {
            limpiarPantalla(); // Limpia la pantalla para mostrar la partida y turno actualizada
            printf("Turno de %s\n\n", turno == 1 ? nombreJugador1 : nombreJugador2);

            // Muestra el tablero del jugador contrario y el propio
            printf("Tablero enemigo:\n");
            dibujarTablero(turno == 1 ? tablero2 : tablero1, 1); // Muestra el tablero enemigo ocultando los barcos
            printf("\nTu tablero:\n");
            dibujarTablero(turno == 1 ? tablero1 : tablero2, 0); // Muestra el tablero propio

            // Solicita las coordenadas de ataque al jugador
            int columna, fila;
            obtenerCoordenadasAtaque(&columna, &fila);

            // Acá verifica si el ataque es válido
            if (validarAtaque(columna, fila, turno == 1 ? tablero2 : tablero1)) {
                // Incrementa el contador de tiros del jugador
                if (turno == 1) {
                    tirosJugador1++;
                } else {
                    tirosJugador2++;
                }

                // Procesa el ataque según el turno del jugador
                if (turno == 1) {
                    if (tablero2[fila][columna] == 'B') {
                        printf("¡Impacto!\n");
                        getch();
                        tablero2[fila][columna] = 'X';
                    } else {
                        printf("Agua.\n");
                        getch();
                        tablero2[fila][columna] = 'O';
                        turnoExtra = 0;
                    }
                } else {
                    if (tablero1[fila][columna] == 'B') {
                        printf("¡Impacto!\n");
                        getch();
                        tablero1[fila][columna] = 'X';
                    } else {
                        printf("Agua.\n");
                        getch();
                        tablero1[fila][columna] = 'O';
                        turnoExtra = 0;
                    }
                }

                // Verifica si el jugador ha ganado
                if (verificarVictoria(turno == 1 ? tablero2 : tablero1)) {
                    limpiarPantalla();
                    printf("¡%s ha ganado el juego!\n", turno == 1 ? nombreJugador1 : nombreJugador2);
                    getch();
                    // Actualiza el ranking del jugador ganador
                    if (turno == 1) {
                        actualizarRanking(nombreJugador1, tirosJugador1);
                    } else {
                        actualizarRanking(nombreJugador2, tirosJugador2);
                    }

                    // Notifica que el ranking se actualizó y lo muestra
                    printf("\nRanking actualizado:\n");
                    ordenarRanking(ORDEN_VICTORIAS);
                    mostrarRanking(ranking);
                    guardarRanking();

                    // Pregunta qué hacer después de la partida
                    printf("¿Qué deseas hacer?\n");
                    printf("1. Volver al menú principal\n");
                    printf("2. Cerrar el juego\n");
                    printf("Selecciona una opción: ");
                    char opcion = getch();
                    if (opcion == '1') {
                        return; // Vuelve al menú principal
                    } else {
                        guardarRanking(); // Guarda el ranking antes de salir
                        liberarRanking(); // Libera memoria del ranking
                        exit(0); // Sale del juego
                    }
                }
            } else {
                printf("Coordenadas inválidas o ya atacadas. Intenta nuevamente.\n");
                getch(); // Pausa para permitir que el jugador vea el mensaje
            }
        }

        // Cambia el turno al siguiente jugador
        turno = (turno == 1) ? 2 : 1;
    }
}

// Función más importante del juego
int main() {

    cargarRanking(); // Cargamos el ranking de jugadores desde el archivo al inicio del juego

    system("chcp 65001");  // Cambiamos a codificación UTF-8 para asegurarnos de que los caracteres especiales se muestren correctamente

    while (1) {  // Bucle principal del menú, el juego continuará ejecutándose hasta que el jugador elija salir

        int opcion = menuPrincipal(); // Mostramos el menú principal y guardamos la opción elegida por el jugador

        if (opcion == 6) { // Si el jugador elige la opción 5, salimos del juego
            break;
        } else if (opcion == 1) { // Si el jugador elige la opción 1, comenzamos un nuevo juego
            char nombreJugador1[50], nombreJugador2[50]; // Definimos los nombres de los jugadores

            // Solicitar nombres de los jugadores
            limpiarBufferEntrada(); // Limpiamos el buffer de entrada para evitar problemas con fgets
            limpiarPantalla();
            
            // Acá ambos jugadores escriben en consola sus nombres
            printf("Introduce el nombre del Jugador 1: ");
            fgets(nombreJugador1, sizeof(nombreJugador1), stdin);
            nombreJugador1[strcspn(nombreJugador1, "\n")] = '\0';  // Eliminar salto de línea al final

            printf("Introduce el nombre del Jugador 2: ");
            fgets(nombreJugador2, sizeof(nombreJugador2), stdin);
            nombreJugador2[strcspn(nombreJugador2, "\n")] = '\0';  // Eliminar salto de línea al final

            // Inicializamos los tableros de ambos jugadores
            char tableroJugador1[TAMANIO_TABLERO][TAMANIO_TABLERO];
            char tableroJugador2[TAMANIO_TABLERO][TAMANIO_TABLERO];
            inicializarTablero(tableroJugador1); // Llenamos el tablero del Jugador 1 con ~
            inicializarTablero(tableroJugador2); // y ahora el del jugador 2

            // Pedimos a cada jugador que coloque sus barcos en el tablero
            colocarBarcosJugador(tableroJugador1, nombreJugador1); // Colocamos los barcos del-
            colocarBarcosJugador(tableroJugador2, nombreJugador2); // - jugador 1 y 2

            // Jugar
            jugar(tableroJugador1, tableroJugador2, nombreJugador1, nombreJugador2); // Ejecutamos el ciclo del juego entre los dos jugadores
        }
    }
    return 0; // Finaliza el programa (el juego se cerrará aquí si se elige salir)
}