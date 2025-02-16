/*
# O que é o problema de concorrência nesse caso?
No código original sem sincronização, múltiplas goroutines podiam acessar e modificar count ao mesmo tempo,
causando uma condição de corrida (race condition) e valores inconsistentes.

# Como o código resolve isso?
A solução aqui é utilizar atomic.AddInt64(&count, 1), que incrementa count de forma atômica.

- atomic.AddInt64() faz a soma de 1 ao valor de count, garantindo que a operação seja segura entre múltiplas goroutines.
- Diferente do sync.Mutex, não há bloqueios — as goroutines podem executar incrementos sem esperar.

# Observação
- sync/atomic é uma solução mais leve e eficiente para contadores e números.
- sync.Mutex é útil quando há múltiplas variáveis ou blocos mais complexos.

# Para o teste foi utilizado o comando do apache benchmark (ab):

ab -n 10000 -c 100 http://localhost:8080/

-n 10000 -> Número total de requisições que serão enviadas (10.000 requisições).
-c 100 -> Número de conexões concorrentes, ou seja, quantas requisições serão feitas ao mesmo tempo (100 requisições simultâneas).
http://localhost:8080/ -> URL do servidor que será testado (localhost na porta 8080)
*/
package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
)

var count int64 = 0

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		atomic.AddInt64(&count, 1)
		res.WriteHeader(http.StatusOK)
		res.Write([]byte(fmt.Sprintf("Você é o visitante número: %d", count)))

	})

	http.ListenAndServe(":8080", nil)
}
