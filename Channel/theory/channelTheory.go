package theory

import "fmt"

/*
channel - очередь значений, операции могут блокироваться при заполнении или пустоте
4 свойства каналов:
- передача по нулевому каналу всегда блокируется
- получение из нулевого канала всегда блокируется
- передача по закрытому каналу вызывает панику
- получение из закрытого канала возвращает zero vulae
*/

func main() {

	// send to a nil channel blocks forever
	var c chan string
	c <- "Hello" // deadlock, because the zero value for uninitalised channel is nil

	//receive from a nil channel blocks forever
	fmt.Println(<-c) //deadlock
	/*
		Explanation:
		- размер буфера канала не частью декларации его типа, поэтому должен быть частью
		значения канала
		- если канал неинициализирован, то размер его буфера будет равен 0
		- если размер буфера равен 0, то канал небуферизованный
		- если канала небуферизованный, то отправка будет блокироваться до тех пор, пока другая
		горутина не будет готова прочитать из него данные
		- если канал nil, то получать и отправитель не связаны друг с другом, они оба заблокированы
		навсегда
	*/

	// send to a closed channel panics

	var c_1 = make(chan int, 10)
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 10; j++ {
				c_1 <- j
			}
			close(c_1) //panic
		}()
	}

	/*
		Паника, потому что первая программа, достигшая 10, закроет канал до того, как братья
		завершат отправку своих сообщений в канал

	*/

	// receive from a closed channel returns the zero value immediately
	c_2 := make(chan int, 3)
	c_2 <- 1
	c_2 <- 2
	c_2 <- 3
	close(c_2)
	for i := 0; i < 4; i++ {
		fmt.Printf("%d ", <-c_2) // 1 2 3 0
	}

	/*
		Solution: use a for range loop
		for v, ok := <- c_2; ok; v, ok <- c_2 {
			...
		}
	*/
}
