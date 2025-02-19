/*
# O que é o problema de concorrência nesse caso?
No código original, várias goroutines podem acessar e modificar count ao mesmo tempo,
causando uma condição de corrida. Isso resulta em valores inconsistentes.

# Como o código resolve isso?
O código resolve o problema de concorrência utilizando um channel para sincronizar o acesso ao contador count.

- A goroutine anônima escuta o canal countChannel em loop infinito, ou seja fica aguardando o canal countChannel
receber um valor e ser esvaziado, atribuindo o valor para increment.

- Quando uma requisição chega, o handler envia uma mensagem (incremento) para o canal countChannel.
(countChannel <- 1  // Envia um incremento para o canal), Isso faz com que a goroutine que escuta o canal
receba a mensagem e execute a operação de incrementar count.

# Como o channel resolve o problema de concorrência?
O canal age como um sinalizador de sincronização entre a goroutine do servidor HTTP e a goroutine que atualiza o contador count.

## Comunicação entre goroutines:
O canal countChannel garante que apenas uma goroutine possa atualizar o contador count por vez.
Sempre que uma requisição chega, ela envia um valor para o canal. O canal armazena e organiza essas mensagens.

## Garantia de atualização sequencial:
Como as mensagens enviadas para o canal são processadas uma de cada vez pela goroutine que escuta o canal,
o contador count nunca será atualizado simultaneamente por múltiplas goroutines.
Isso evita condições de corrida e garante que o valor de count seja atualizado de forma sequencial e consistente.

## Sincronização sem bloqueio explícito:
Diferente de usar um mutex (onde você precisaria bloquear e desbloquear o acesso à variável), o canal resolve a concorrência
de maneira implícita e sem a necessidade de locks. Ele faz isso ordenando as mensagens e permitindo que apenas uma goroutine
acesse o valor de count de cada vez.
*/
package main

import (
	"fmt"
	"net/http"
)

var count int64 = 0
var countChannel = make(chan int64)

func main() {

	go func() {
		for {
			increment := <-countChannel // Esvazia o canal, atribui o valor para increment
			count += increment
		}
	}()

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		countChannel <- 1 // Enche o canal, publica 1 no canal
		res.WriteHeader(http.StatusOK)
		res.Write([]byte(fmt.Sprintf("Você é o visitante número: %d", count)))

	})

	http.ListenAndServe(":8080", nil)
}
