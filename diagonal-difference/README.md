# Diagonal Difference

Given a square matrix, calculate the absolute difference between the sums of its diagonals.

## Análisis de código

El código proporcionado calcula la diferencia absoluta entre las sumas de las dos diagonales de una matriz cuadrada y tiene una complejidad de tiempo lineal y una complejidad de espacio constante.

Complejidad de tiempo:
El código contiene un bucle for que itera a través de la diagonal de la matriz, donde la matriz tiene una dimensión de n x n (en este caso, n=3). El bucle recorre una vez por cada elemento de la diagonal, por lo que la cantidad de iteraciones del bucle es igual al número de elementos en la diagonal, que es n. Por lo tanto, la complejidad de tiempo es O(n), donde 'n' es el número de filas (o columnas) de la matriz cuadrada.

Complejidad de espacio:
El código utiliza solo un conjunto limitado de variables, independientemente del tamaño de la matriz de entrada. Las variables son la matriz de entrada 'arr', las variables 'diagonalA' y 'diagonalB' para almacenar las sumas de las dos diagonales, y el índice 'i' del bucle for. Todas estas variables ocupan una cantidad constante de espacio en memoria, y no hay crecimiento proporcional al tamaño de la matriz de entrada. Por lo tanto, la complejidad de espacio es O(1), lo que significa que el consumo de memoria es constante, independientemente del tamaño de entrada.

En resumen, tanto la complejidad de tiempo como la complejidad de espacio son lineales con respecto al número de filas (o columnas) de la matriz cuadrada de entrada 'arr'. Por lo tanto, el código es eficiente y su rendimiento es predecible incluso para matrices cuadradas grandes.