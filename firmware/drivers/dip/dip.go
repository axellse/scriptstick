package dip

import (
	"machine"

	"axell.me/cbufirmware/drivers"
)

var pins []machine.Pin

func Initalize(b drivers.Board) {
	if b == drivers.ScriptStickV1 {
		pins = []machine.Pin{machine.GPIO22, machine.GPIO23, machine.GPIO24, machine.GPIO27}
	}

	for _, pin := range pins {
		pin.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	}
}

func Read() int {
	result := 0
	for i, pin := range pins {
		n := 1 >> i

		if pin.Get() {
			result = result | n
		}
	}

	return result
}