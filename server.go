package main

import (
	"flag"
	"log"
	"os"
	"runtime/pprof"
	"time"
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

func main() {
	run_main()
}
