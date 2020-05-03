package main

import (
	//"context"
	//	"crypto/tls"
	//"encoding/json"
	//"strconv"
	//	"strings"

	//"errors"
	//"fmt"
	//"os"
	//"os/exec"
	//"time"
	//"io/ioutil"
	//"math/rand"
	//"net/http"
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	//"sync"
	//"time"
	orderslib "./orders"
	quotes "./quotes"
	//	"github.com/chrisftw/ezconf"
	//	log "github.com/sirupsen/logrus"
)

func main() {

	runtime.GOMAXPROCS(2)
	reader := bufio.NewReader(os.Stdin)
	for {
		printMenu()
		fmt.Print("$ ")
		cmdString, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		err = runCommand(cmdString)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
	//asyncGo()
	//asyncGo2()
	//asyncUsingChannelsToAvoidRaceConditions()
	//WaitingOnAllRoutinesToFinish()
	//AsyncSequentialCallToRoutines(wgGlobal)
	// AsyncAbortRoutine()
	//AsyncSequenciallyUsingChannelsAndSlices()
	// WaitGroupTest()
	//WGTest2()
	//TestParrallProc()
	// token := orderslib.GetToken()
	// fmt.Println(token)
	// orderslib.GetOrdersParallel(token)
	//GetQuotesParallel()

	//lineListSource(context.Background(), "Hello \n World \n Goodbye \n Foo")
}

func runCommand(commandStr string) error {
	commandStr = strings.TrimSuffix(commandStr, "\n")
	arrCommandStr := strings.Fields(commandStr)
	switch arrCommandStr[0] {
	case "a":
		token := orderslib.GetToken()
		fmt.Println(token)
		orderslib.GetOrdersParallel(token)
	case "b":
		quotes.GetQuotesParallel()
	case "exit":
		os.Exit(0)
		// add another case here for custom commands.
	}
	cmd := exec.Command(arrCommandStr[0], arrCommandStr[1:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

func printMenu() {
	kvs := map[string]string{
		"a": "GetQuotesParallel ",
		"b": "GetOrdersParallel ",
	}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}
}
