# Staircase
Staircase detail

This is a staircase of size :

   *
  **
 ***
****
Its base and height are both equal to . It is drawn using * symbols and spaces. The last line is not preceded by any spaces.

Write a program that prints a staircase of size .

Function Description

Complete the staircase function in the editor below.

staircase has the following parameter(s):

int n: an integer
Print

Print a staircase as described above.

## Análisis de código

Complejidad de tiempo:
El código consta de un bucle for que itera desde 1 hasta 'n' (el valor de entrada). En cada iteración, se imprimen dos strings repetidos 'n - i' y 'i' veces, respectivamente, utilizando la función strings.Repeat(). Como ambas operaciones de repetición de strings tienen una complejidad de tiempo lineal con respecto a su longitud, el tiempo de ejecución total del bucle es la suma de las longitudes de los dos strings repetidos en cada iteración.

En el peor caso, 'n' es el máximo valor posible para el tipo int32, que es 2,147,483,647. Sin embargo, dado que 'n' en el mundo real rara vez será tan grande y que solo nos interesa la tendencia asintótica, podemos decir que la complejidad de tiempo es O(n).

Complejidad de espacio:
El código utiliza solo variables int32 y strings para formatear y almacenar la salida. La cantidad de espacio requerido para estas variables no depende del valor de 'n' y, por lo tanto, es constante en relación con el tamaño de entrada. La función strings.Repeat() no crea duplicados reales de los strings, sino que usa una referencia interna al string original, lo que hace que la complejidad de espacio para la repetición sea O(1).

En resumen, la complejidad de tiempo es O(n), donde 'n' es el valor de entrada, y la complejidad de espacio es O(1), es decir, constante e independiente del tamaño de entrada. El código es eficiente en términos de tiempo y espacio, y su rendimiento es predecible incluso para valores de 'n' grandes.