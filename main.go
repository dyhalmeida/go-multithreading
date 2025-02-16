/*
# O que é o problema de concorrência nesse caso?
No código original, várias goroutines podem acessar e modificar count ao mesmo tempo,
causando uma condição de corrida. Isso resulta em valores inconsistentes.

# Como o código resolve isso?
A solução está no uso do mutex (sync.Mutex), que funciona como um cadeado para garantir
que apenas uma goroutine por vez possa modificar count.

- Criamos um mutex (m), que será usado para controlar o acesso à variável count.
- m.Lock() trava o acesso à variável count. Enquanto uma goroutine estiver executando esse trecho, outras terão que esperar.
- Incrementa count com segurança.
- m.Unlock() libera o acesso para que outra goroutine possa modificar count.

# Como o Mutex resolve o problema?
- Antes de modificar count, uma goroutine trava o mutex (m.Lock()).
- Enquanto o mutex está travado, nenhuma outra goroutine pode modificar count.
- Depois de atualizar count, o mutex é destravado (m.Unlock()), permitindo que a próxima goroutine continue.
- Assim, cada requisição manipula count de forma segura e ordenada.

# Para o teste foi utilizado o comando do apache benchmark (ab):

ab -n 10000 -c 100 http://localhost:8080/

-n 10000 -> Número total de requisições que serão enviadas (10.000 requisições).
-c 100 -> Número de conexões concorrentes, ou seja, quantas requisições serão feitas ao mesmo tempo (100 requisições simultâneas).
http://localhost:8080/ -> URL do servidor que será testado (localhost na porta 8080)

# Para detectar race condition em modo de desenvolvimento, foi utilizando o parâmetro -race no comando go run

go run -race main.go
*/
package main

import (
	"fmt"
	"net/http"
	"sync"
)

var count int64 = 0

func main() {
	m := sync.Mutex{}
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		m.Lock()
		count = count + 1
		m.Unlock()
		res.WriteHeader(http.StatusOK)
		res.Write([]byte(fmt.Sprintf("Você é o visitante número: %d", count)))

	})

	http.ListenAndServe(":8080", nil)
}
