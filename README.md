# UNIVERSIDAD PERUANA DE CIENCIAS APLICADAS

## Programación Concurrente y Distribuida - CC65
### Tarea académica 3

---

### INTEGRANTES

- Charapaqui Reluz, Alcides Rommel: u202021294
- Benites Diaz, Juan Alejandro: u202010449
- Humbser Meza, Diego Fernando: u202012711

---

### SECCIÓN
SW72

---

### DOCENTE
Jara Garcia, Carlos Alberto

---

### CICLO
2023-02

---

## Descripción del Juego: Ludo Modificado

El Ludo modificado es una versión ampliada y adaptada del popular juego de mesa Ludo. En esta versión, los jugadores compiten para llevar a sus personajes a través de un peligroso laberinto lleno de obstáculos y desafíos. Cada jugador tiene la tarea de guiar a sus personajes a través del laberinto y llegar a la meta antes que los demás.

### Reglas del Juego:

- **Tablero del Laberinto**: El tablero está dividido en casillas con caminos y obstáculos. Cada jugador tiene cuatro personajes que comienzan en puntos de partida específicos.
  
- **Turnos y Movimientos**: Los jugadores se turnan para lanzar un dado y mover a sus personajes. Los jugadores lanzan tres dados, dos dados normales (del 1 al 6) y uno con la operación (sumar o restar) para determinar cuántos pasos pueden avanzar o retroceder en su turno. Los jugadores pueden mover un solo personaje por turno.

- **Obstáculos**: El laberinto está lleno de obstáculos como paredes, trampas y criaturas que bloquean el paso de los personajes en varias casillas. Si al personaje le toca avanzar hacia una casilla con obstáculo, entonces el jugador pierde el turno y continúa el siguiente jugador.

- **Objetivo**: El objetivo es llevar a los cuatro personajes desde los puntos de partida hasta la meta en el menor número de turnos posible. El primer jugador en llevar a todos sus personajes a la meta gana el juego.

### Implementación en Go:

El juego ha sido implementado en el lenguaje de programación Go, utilizando características de concurrencia y canales para simular el juego en tiempo real con múltiples jugadores.

### Cómo jugar:

1. Ejecute el programa en su máquina.
2. Observe cómo los jugadores lanzan los dados y mueven sus peones de manera concurrente.
3. El juego terminará cuando un jugador haya movido todos sus peones al final del tablero.

---

Titulo: "Ludo Modificado"
Descripcion: "Una versión ampliada y adaptada del popular juego de mesa Ludo implementada en Go."

Estructuras_y_Constantes:
  - Constantes:
    - NumPlayers: "Número de jugadores en el juego."
    - NumPawns: "Número de peones por jugador."
    - MaxSquare: "Número máximo de casillas en el tablero."
    - Start: "Valor inicial para los peones que aún no han comenzado a moverse."
    - Finish: "Valor que indica que un peón ha llegado al final."
    - Obstacle: "Valor que indica un obstáculo en el tablero."
  - Board:
    - Descripcion: "Representa el tablero del juego."
    - Atributos:
      - positions: "Matriz que almacena las posiciones actuales de todos los peones de todos los jugadores."
      - squares: "Array que representa el tablero. Puede contener obstáculos."
      - mutex: "Mutex para garantizar la sincronización cuando se accede al tablero."
      - gameOver: "Indica si el juego ha terminado."
  - Player:
    - Descripcion: "Representa a un jugador en el juego."
    - Atributos:
      - id: "Identificador único del jugador."
      - board: "Referencia al tablero del juego."
      - channel: "Canal de comunicación para el jugador."

Funciones:
  - hasWon:
    - Descripcion: "Verifica si todos los peones de un jugador han llegado al final."
    - Concurrency: "No"
  - play:
    - Descripcion: "Función principal que simula el juego para un jugador."
    - Concurrency: "Sí, cada jugador juega en su propia goroutine."
    - Canales: "Cada jugador tiene un canal para comunicarse con el tablero."
    - Detalles:
      - "Los jugadores lanzan un dado y mueven uno de sus peones en su turno."
      - "El juego continúa hasta que un jugador mueve todos sus peones al final del tablero."
  - main:
    - Descripcion: "Función principal que inicializa el juego y comienza la simulación."
    - Concurrency: "Sí, se crean múltiples goroutines para los jugadores."
    - Canales: "Se crean canales para cada jugador para comunicarse con el tablero."

Enfasis_en_Concurrencia_y_Canales:
  - Concurrency:
    - Descripcion: "El juego utiliza goroutines para simular múltiples jugadores jugando de manera concurrente."
    - Uso: "Cada jugador juega en su propia goroutine, lo que permite que todos jueguen al mismo tiempo."
  - Canales:
    - Descripcion: "Los canales se utilizan para la comunicación entre jugadores y el tablero."
    - Uso: "Cada jugador tiene un canal para enviar y recibir información del tablero. Esto garantiza que los movimientos de los jugadores se manejen de manera sincronizada y que un jugador no pueda moverse fuera de su turno."

Conclusión:
  - "La implementación del juego Ludo modificado en Go demuestra la potencia y flexibilidad de la programación concurrente. A través de este proyecto, hemos aprendido cómo manejar múltiples goroutines, sincronizar el acceso a recursos compartidos y comunicarnos entre diferentes partes del programa usando canales."


### Conclusión:

La implementación del juego Ludo modificado en Go demuestra la potencia y flexibilidad de la programación concurrente. A través de este proyecto, hemos aprendido cómo manejar múltiples goroutines, sincronizar el acceso a recursos compartidos y comunicarnos entre diferentes partes del programa usando canales.

---

