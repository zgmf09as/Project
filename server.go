package main

import (
	"flag"
	"log"
	"os"
	"runtime/pprof"
	"time"
	"fmt"
	"math/rand"
	"runtime"
	_ "os/signal"
)
func run_main() {
	var rundur = flag.Int("rundur", 60*60*24*365, "run time sec")
	var profilefilename = flag.String("pfilename", "", "profile filename")
	var config = flag.String("config", "", "config filename")
	flag.Parse()
	log.Printf("rundur:%vs profile:%v config:%v",
		*rundur, *profilefilename, *config)

	//if *config != "" {
	//	ok := shootbase.GameConst.Load(*config)
	//	if !ok {
	//		log.Fatal("config load fail")
	//	}
	//}

	if *profilefilename != "" {
		f, err := os.Create(*profilefilename)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	service := NewService()

	time.Sleep(time.Duration(*rundur) * time.Second)
	service.SendGoCmd("quit", nil, nil)
	time.Sleep(1 * time.Second)
}

func f() {

	defer func() {
		s:=recover()
		fmt.Println(s)
	}()

	a := [...]int{1,2}
	for i :=0; i < 5; i++ {
		fmt.Println(a[i])
	}
}

func hello( i int ) {
	rand.Seed(time.Now().Unix())
	r := rand.Intn(100)
	time.Sleep(time.Duration(r))
	fmt.Println(i, r)
}

func sum(a int, b int, c chan int) {
	c <- a+b
}

func main() {
	//run_main()
	////test defer
	//f()
	//
	//fmt.Println("Hello, world")
	////end test

	////test Ducktyping
	//var basicSoldier Soldier
	//var basicGeneral General
	//
	//Command(basicGeneral)
	//Command(basicSoldier)
	////end test

	////test goroutine
	//runtime.GOMAXPROCS(runtime.NumCPU())
	//runtime.GOMAXPROCS(1)
	//
	//fmt.Println("Use MAX CPU : ",runtime.GOMAXPROCS(0))
	//
	//for i :=0 ; i < 100; i++ {
	//	//go hello(i)
	//	go func(n int) {
	//		fmt.Println("goroutine : ", n)
	//	}(i)
	//}
	////end test

	////test channel
	//done := make(chan bool) //create channel and synchronous chanel(동기채널)
	//count := 3
	//
	//go func() {
	//	for i :=0; i < count; i++ {
	//		done <- true
	//		fmt.Println("고루틴 : ", i)
	//		time.Sleep(1 * time.Second)
	//	}
	//}()
	//
	//for i := 0; i < count; i++ {
	//	<-done
	//	fmt.Println("메인 함수 : ", i)
	//}
	////end test

	////test channel buffering
	runtime.GOMAXPROCS(1)

	done := make(chan bool, 2) //2 buffer chanel
	count := 4

	go func(){
		for i:= 0; i < count; i++ {
			done <- true
			fmt.Println("고루틴 : ", i)
			//r := rand.Intn(100)
			//fmt.Println("고루틴 : ", i, " r:", r)
			//time.Sleep(time.Duration(r))
		}
	}()

	for i:=0; i < count; i++ {
		<-done
		r := rand.Intn(100)
		fmt.Println("메인함수 : ", i, " r:", r)
		time.Sleep(time.Duration(r))
	}

	fmt.Scanln()

	//sigChan := make(chan os.Signal, 1)
	//signal.Notify(sigChan, os.Interrupt, os.Kill)
	//<-sigChan
}
