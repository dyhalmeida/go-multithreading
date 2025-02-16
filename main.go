/*
Esse código apresenta um problema de concorrência quando múltiplas requisições HTTP são feitas simultaneamente.

O Problema de Concorrência
A variável count é uma variável global e está sendo modificada diretamente dentro do handler HTTP,
sem nenhuma forma de sincronização.

O que acontece quando múltiplas requisições chegam ao mesmo tempo?

Como o servidor HTTP do Go lida com cada requisição em uma goroutine separada,
várias goroutines podem tentar atualizar count ao mesmo tempo. Isso leva a uma condição de corrida (race condition),
onde duas ou mais goroutines podem ler e escrever a variável simultaneamente, resultando em um valor inconsistente de count.

Exemplo do problema:
Duas requisições chegam quase ao mesmo tempo.
Ambas as goroutines leem o valor atual de count (suponha que seja 10).
Ambas calculam count + 1 e atribuem o novo valor (11).
O valor de count deveria ser 12, mas como ambas escreveram 11, perdemos um incremento.

Para o teste foi utilizado o comando do apache benchmark (ab):

ab -n 10000 -c 100 http://localhost:8080/

-n 10000 -> Número total de requisições que serão enviadas (10.000 requisições).
-c 100 -> Número de conexões concorrentes, ou seja, quantas requisições serão feitas ao mesmo tempo (100 requisições simultâneas).
http://localhost:8080/ -> URL do servidor que será testado (localhost na porta 8080)
*/
package main

import (
	"fmt"
	"net/http"
)

var count int64 = 0

func main() {

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		count = count + 1
		res.WriteHeader(http.StatusOK)
		res.Write([]byte(fmt.Sprintf("Você é o visitante número: %d", count)))

	})

	http.ListenAndServe(":8080", nil)
}
