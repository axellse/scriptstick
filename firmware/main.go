package main

import (
	_ "embed"
	"fmt"
	"time"

	"axell.me/cbufirmware/drivers"
	"axell.me/cbufirmware/drivers/dip"
	"github.com/axellse/gopher-lua"
)

//go:embed target.txt
var BuildTarget drivers.Board

func main() {
	//note that flash is not implemented since i dont have the device.
	fmt.Println("booted scriptstick firmware v0.0.1")
	fmt.Println("built for " + drivers.BoardName(BuildTarget))
	fmt.Println("NOTE: This firmware is incomplete, very basic and has not been tested.")
	fmt.Println("")

	dip.Initalize(BuildTarget)
	if dip.Read() == 15 {
		fmt.Println("Programming mode enabled!")
		//requires flash implementation
	}

	state := lua.NewState()
	defer state.Close()

	//example
	if err := state.DoString(`print("hello")`); err != nil {
   		panic(err)
	}
	for {
		time.Sleep(2 * time.Second)
	}
}