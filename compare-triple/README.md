Alice and Bob each created one problem for HackerRank. A reviewer rates the two challenges, awarding points on a scale from 1 to 100 for three categories: problem clarity, originality, and difficulty.

The rating for Alice's challenge is the triplet a = (a[0], a[1], a[2]), and the rating for Bob's challenge is the triplet b = (b[0], b[1], b[2]).

The task is to find their comparison points by comparing a[0] with b[0], a[1] with b[1], and a[2] with b[2].

If a[i] > b[i], then Alice is awarded 1 point.
If a[i] < b[i], then Bob is awarded 1 point.
If a[i] = b[i], then neither person receives a point.
Comparison points is the total points a person earned.

Given a and b, determine their respective comparison points.

## Análisis del código

Para analizar la complejidad de tiempo y espacio del código proporcionado, vamos a evaluar ambas métricas por separado.

Complejidad de tiempo:
El código consta de un bucle for que itera a través de las listas 'a' y 'b', ambas de longitud n (len(a) == len(b) == n). En cada iteración del bucle, se realiza una comparación entre los elementos de 'a' y 'b' y, dependiendo de la comparación, se incrementa el valor de 'scoreA' o 'scoreB'.

El tiempo de ejecución del bucle for es directamente proporcional al tamaño de entrada 'n', ya que cada iteración implica operaciones constantes (comparaciones y asignaciones). En consecuencia, la complejidad de tiempo es O(n), lo que significa que el tiempo de ejecución aumentará linealmente con el tamaño de las listas 'a' y 'b'.

Complejidad de espacio:
El código utiliza algunas variables adicionales para realizar el seguimiento de las puntuaciones 'scoreA' y 'scoreB'. Sin embargo, estas variables son de tipo int32 y ocupan una cantidad constante de espacio en memoria, independientemente del tamaño de las listas de entrada 'a' y 'b'.

El único espacio adicional utilizado es el arreglo de salida que contiene las dos puntuaciones (scoreA y scoreB). Dado que siempre tiene una longitud fija de 2 elementos, su tamaño no depende del tamaño de entrada 'n' y, por lo tanto, su complejidad de espacio es constante, es decir, O(1).

En resumen, la complejidad de tiempo del código es O(n) y la complejidad de espacio es O(1). Esto significa que el tiempo de ejecución aumentará linealmente con el tamaño de las listas 'a' y 'b', mientras que el consumo de memoria se mantendrá constante, independientemente del tamaño de entrada.