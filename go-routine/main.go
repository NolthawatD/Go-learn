package main

import (
	"fmt"
	"sync"
	"time"

	"math/rand"

	"github.com/gofiber/fiber/v2"
	"github.com/robfig/cron/v3"
)

func main() {
	// funcGoRoutine()
	// funcGoChannel()
	// funcRoutineSelect()
	// funcWaitGroup()
	// funcMutex()
	// funcOnce()
	// funcNewCond()
	// funcPubSub()
	// funcFiberPubsub()
	funcCronJob()
}

func funcCronJob() {
	c := cron.New(cron.WithSeconds())
	c.AddFunc("*/10 * * * * *", func() {
		fmt.Println("Hello wolrd every 10 seconds")
	})

	c.Start()

	select {}
}

var pubsub = &PubSub{
	Subscribers: make(map[uint]*Subscriber),
}

func funcFiberPubsub() {
	app := fiber.New()

	// Start a subscription routine
	go func() {
		subscriber := pubsub.Subscribe()
		defer pubsub.Unsubscribe(subscriber.ID)

		for msg := range subscriber.Channel {
			// Handle the message, e.g., log, process, etc.
			fmt.Printf("Received message: %s\n", msg.Content)
		}
	}()

	// Publish endpoint
	app.Post("/publish", func(c *fiber.Ctx) error {
		var msg Message
		if err := c.BodyParser(&msg); err != nil {
			return err
		}

		go pubsub.Publish(msg)
		return c.SendString("Message published")
	})

	app.Listen(":8888")
}

func funcPubSub() {
	// สร้าง channel เพื่อส่งข้อความ
	ch := make(chan string)

	// สร้าง goroutine เพื่อส่งข้อความไปยัง channel
	go func() {
		for i := 0; i < 10; i++ {
			ch <- fmt.Sprintf("Hello, world! %d", i)
			time.Sleep(1 * time.Second)
		}
	}()

	// สร้าง goroutine เพื่อรับข้อความจาก channel
	go func() {
		for {
			msg := <-ch
			fmt.Println(msg)
		}
	}()

	// รอให้ goroutines ทำงานเสร็จสิ้น
	time.Sleep(5 * time.Second)
}

// ## NewCond go-routine ล็อกทรัพยากรเมื่อกำลังรอเงื่อนไข และปลดล็อกเมื่อเข้าเงื่อนไข
func funcNewCond() {
	// Create a new condition variable
	var mutex sync.Mutex
	cond := sync.NewCond(&mutex)

	// A shared resource
	ready := false

	// A goroutine that waits for a condition
	go func() {
		fmt.Println("Goroutine: Waiting for the condition...")

		mutex.Lock()
		for !ready {
			cond.Wait() // Wait for the condition
		}
		fmt.Println("Goroutine: Condition met, proceeding...")
		mutex.Unlock()
	}()

	// Simulate some work (e.g., loading resources)
	time.Sleep(2 * time.Second)

	// Signal the condition
	mutex.Lock()
	ready = true
	cond.Signal() // Signal one waiting goroutine
	mutex.Unlock()
	fmt.Println("Push signal !")

	// Give some time for the goroutine to complete
	time.Sleep(1 * time.Second)
	fmt.Println("Main: Work is done.")
}

// ## Once เรียกใช้งานเพียงครั้งเดียว
func funcOnce() {
	var once sync.Once
	var wg sync.WaitGroup

	initialize := func() {
		fmt.Println("Initializing only once")
	}

	doWork := func(workerId int) {
		defer wg.Done()
		fmt.Printf("Worker %d started\n", workerId)
		once.Do(initialize) // This will only be executed once
		fmt.Printf("Worker %d done\n", workerId)
	}

	numWorkers := 5
	wg.Add(numWorkers)

	// Launch several goroutines
	for i := 0; i < numWorkers; i++ {
		go doWork(i)
	}

	// Wait for all goroutines to complete
	wg.Wait()
	fmt.Println("All workers completed")
}

// ## Mutex lock ล็อคให้ go routine หนึ่งใช้ทรัพยากร
// Counter struct holds a value and a mutex
type Counter struct {
	value int
	mu    sync.Mutex
}

// Increment method increments the counter's value safely using the mutex
func (c *Counter) Increment() {
	c.mu.Lock()   // Lock the mutex before accessing the value
	c.value++     // Increment the value
	c.mu.Unlock() // Unlock the mutex after accessing the value
}

// Value method returns the current value of the counter
func (c *Counter) Value() int {
	return c.value
}

func funcMutex() {
	var wg sync.WaitGroup
	counter := Counter{}

	// Start 10 goroutines
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				counter.Increment()
			}
		}()
	}

	wg.Wait() // Wait for all goroutines to finish
	fmt.Println("Final counter value:", counter.Value())
}

// ## WaitGroup จะรอ gorountine เสร็จเป็นจำนวนเท่าไหร่
func funcWaitGroup() {
	var wg sync.WaitGroup

	// Launch several goroutines and increment the WaitGroup counter for each
	wg.Add(5)
	for i := 1; i <= 5; i++ {
		// wg.Add(1) // เพิ่มทีละ 1
		go worker(i, &wg)
	}

	wg.Wait() // Block until the WaitGroup counter goes back to 0; all workers are done

	fmt.Println("All workers completed")
}

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement the counter when the goroutine completes

	fmt.Printf("Worker %d starting\n", id)

	// Simulate some work by sleeping
	sleepDuration := time.Duration(rand.Intn(1000)) * time.Millisecond
	time.Sleep(sleepDuration)

	fmt.Printf("Worker %d done\n", id)
}

// # Race condition

// ## Go-channel and select
func funcRoutineSelect() {
	channel1 := make(chan int)
	channel2 := make(chan int)

	go func() {
		channel1 <- 10
		close(channel1)
	}()
	go func() {
		channel2 <- 20
		close(channel2)
	}()

	closeChannel1, closeChannel2 := false, false

	for {
		if closeChannel1 && closeChannel2 {
			break
		}
		select {
		case v, ok := <-channel1:
			if !ok {
				closeChannel1 = true
				continue
			}
			fmt.Println("Channel 1", v)
		case v, ok := <-channel2:
			if !ok {
				closeChannel2 = true
				continue
			}
			fmt.Println("Channel 2", v)
		}
	}
}

// ## Go-channel
func funcGoChannel() {
	ch := make(chan int)

	go func() {
		ch <- 10
		ch <- 20
		ch <- 30
		ch <- 40
		close(ch) // close channel
	}()

	for v := range ch {
		fmt.Println(v)
	}
}

// ## Go-rountine
func funcGoRoutine() {
	// สร้าง Goroutines ใหม่ 2 ตัว
	go doSomething1()
	go doSomething2()

	// รอให้ Goroutines ทั้งสองตัวทำงานเสร็จสิ้น
	time.Sleep(time.Second)

	// พิมพ์ข้อความ
	fmt.Println("Done")
}

func doSomething1() {
	fmt.Println("Doing something 1")
}

func doSomething2() {
	fmt.Println("Doing something 2")
}
