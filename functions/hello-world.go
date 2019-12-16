package main

// Hellow prog

// func main() {
// 	fmt.Println("hello world")
// }

// Values

// func main() {
// 	fmt.Println("go" + "lang")
// 	fmt.Println("1+1=", 1+1)
// 	fmt.Println("7.0/3.0=", 7.0/3.0)
// 	fmt.Println(false && false)
// 	fmt.Println(true && false)
// 	fmt.Println(true || false)
// 	fmt.Println(!true)
// }

// channel

// func main() {
// 	message := make(chan string, 3)

// 	message <- "Hemant"
// 	message <- "Kumar"
// 	message <- "Joshi"

// 	fmt.Println(<-message)
// 	fmt.Println(<-message)
// 	fmt.Println(<-message)
// }

// channel synchronization

// func master(do chan bool) {

// 	fmt.Println("Make me happy")
// 	do <- true
// }

// func worker(done chan bool) {
// 	fmt.Println("workin....")
// 	time.Sleep(time.Second)
// 	fmt.Println("done")

// 	done <- true
// }

// func main() {

// 	done := make(chan bool, 1)
// 	go worker(done)

// 	do := make(chan bool, 1)
// 	go master(do)

// 	<-done
// 	<-do

// }

// select

// func main() {
// 	c1 := make(chan string)
// 	c2 := make(chan string)

// 	go func() {
// 		time.Sleep(1 * time.Second)
// 		c1 <- "one"
// 	}()

// 	go func() {
// 		time.Sleep(5 * time.Second)
// 		c1 <- "two"
// 	}()

// 	for i := 0; i < 2; i++ {
// 		select {
// 		case msg1 := <-c1:
// 			fmt.Println("received", msg1)
// 		case msg2 := <-c2:
// 			fmt.Println("received", msg2)

// 		}
// 	}
// }

// timeouts
