# Mini-Max Sum
Given five positive integers, find the minimum and maximum values that can be calculated by summing exactly four of the five integers. Then print the respective minimum and maximum values as a single line of two space-separated long integers.

## Análisis de código

La complejidad de tiempo del código es O(n^2), y la complejidad de espacio es O(1).

Complejidad de tiempo:
El código contiene dos bucles for anidados. El bucle exterior itera a través del arreglo 'arr' de longitud 'n' (len(arr)). El bucle interior también itera a través del mismo arreglo con una longitud 'n'. Dentro del bucle interior, se realiza una operación de suma para calcular la suma de los elementos del arreglo, excluyendo el elemento en la posición 'i'.

El bucle interior tiene una complejidad de tiempo O(n), y dado que el bucle interior se ejecuta 'n' veces para cada iteración del bucle exterior, la complejidad de tiempo total del código es O(n * n), que se simplifica a O(n^2), donde 'n' es la longitud del arreglo de entrada 'arr'.

Complejidad de espacio:
El código utiliza un conjunto limitado de variables, independientes del tamaño del arreglo de entrada 'arr'. Las variables son 'min', 'max', 'sum', 'founded', 'i' y 'j'. Todas estas variables ocupan una cantidad constante de espacio en memoria, independientemente del tamaño del arreglo de entrada 'arr'.

Además de estas variables, el arreglo de entrada 'arr' ocupa espacio en memoria proporcional a su longitud 'n'. Sin embargo, en términos de complejidad de espacio, no consideramos esto porque se trata del espacio ocupado por los datos de entrada, que no depende del algoritmo en sí.

Por lo tanto, la complejidad de espacio es O(1), lo que significa que el consumo de memoria es constante, independientemente del tamaño de entrada.

En resumen, la complejidad de tiempo del código es O(n^2), y la complejidad de espacio es O(1). La complejidad de tiempo cuadrática es una señal de que el algoritmo podría ser optimizado para listas de entrada grandes, ya que podría volverse ineficiente.