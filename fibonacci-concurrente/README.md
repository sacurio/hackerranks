# Fibonacci concurrente
Implementa una función para calcular el n-ésimo número de Fibonacci de forma concurrente utilizando goroutines en Golang.

## Measuring the performance of three Fibonacci algorithms
<<inside of the container that contains the benchmark test file>>
``` bash
go install golang.org/x/perf/cmd/benchstat@latest
go test -bench=. -count 5 | tee old.txt
benchstat old.txt
```