/*
# Problema de Concorrência no Código
O código está simulando a recepção de mensagens de dois sistemas diferentes,
como se fosse uma integração com RabbitMQ e Kafka, através de duas goroutines
que enviam mensagens para canais diferentes (chn1 e chn2).
O problema de concorrência surge quando tentamos ler e processar as mensagens
de ambos os canais ao mesmo tempo, pois precisamos de uma maneira eficiente e
segura de gerenciar essa comunicação paralela sem causar bloqueios ou sobrecargas.

# Como o código resolve o problema de concorrência?
O código resolve esse problema de concorrência utilizando o mecanismo de goroutines e channels juntamente com o select.

## Goroutines
Duas goroutines estão sendo criadas:
- Uma simula a recepção de mensagens do RabbitMQ.
- A outra simula a recepção de mensagens do Kafka.
Cada goroutine está enviando uma mensagem para seu respectivo canal (chn1 ou chn2) em intervalos de tempo diferentes.

## Select
O select é a chave para resolver o problema de concorrência. Ele é usado para esperar por mensagens de ambos os canais
simultaneamente de forma eficiente. O select vai funcionar assim:

- Quando uma mensagem é recebida de chn1 (RabbitMQ), ela será processada.
- Quando uma mensagem é recebida de chn2 (Kafka), ela será processada.
- Se nenhuma mensagem chegar dentro de 1 segundo, o select dispara a opção do timeout, imprimindo a mensagem "timeout".

O select resolve o problema de concorrência de forma não bloqueante porque ele permite que o programa escute os dois canais
ao mesmo tempo e faça o que for necessário quando uma mensagem chega, sem precisar bloquear uma goroutine enquanto espera pela outra.

# Como o Select Resolve a Concorrência?
O select resolve a concorrência da seguinte forma:

- Ele escuta ambos os canais (chn1 e chn2) simultaneamente, sem que o código precise se preocupar em controlar qual canal está sendo lido primeiro.
- Como as goroutines podem estar operando em paralelo, o select ajuda a garantir que uma mensagem de qualquer
canal será processada assim que chegar, permitindo interação concorrente sem bloqueios.
- Se nenhum canal enviar uma mensagem dentro de 1 segundo, o select aciona o timeout, permitindo que o código não fique bloqueado esperando indefinidamente.
*/
package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type Message struct {
	ID  int64
	Msg string
}

func main() {

	chn1 := make(chan Message)
	chn2 := make(chan Message)

	var i int64 = 0

	// RabbitMQ
	go func() {
		for {
			atomic.AddInt64(&i, 1)
			time.Sleep(time.Second * 2)
			msg := Message{i, "Message from RabbitMQ"}
			chn1 <- msg
		}
	}()

	// Kafka
	go func() {
		for {
			atomic.AddInt64(&i, 1)
			time.Sleep(time.Second * 3)
			msg := Message{i, "Message from Kafka"}
			chn2 <- msg
		}
	}()

	for {
		select {
		case msg := <-chn1:
			fmt.Printf("received from rabbitmq id %d: %s \n", msg.ID, msg.Msg)
		case msg := <-chn2:
			fmt.Printf("received from kafka id %d: %s \n", msg.ID, msg.Msg)
		case <-time.After(time.Second):
			println("timeout")
		}
	}

}
