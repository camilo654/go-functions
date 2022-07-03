package app

import (
	"os"
	"time"

	"github.com/camilo654/go-functions/src/services"
	"github.com/camilo654/go-functions/src/util"
	"github.com/camilo654/go-functions/src/view"
)

func Run() {
	for {
		view.ShowMenu()
		input := util.ReadInt()
		if util.Valid(input) {
			runOption(input)
		}
		time.Sleep(1 * time.Second)
	}
}

func runOption(input int) {
	switch input {
	case 1:
		view.GetNames()
		names := util.ReadString()
		view.ShowResponse(services.GetNames(names))
	case 2:
		view.GetPokemonId()
		input := util.ReadInt()
		view.ShowResponse(services.GetPokemonName(input))
	case 3:
		view.GetWords()
		strA, strB := util.ReadStrings()
		view.ShowResponse(services.AreFriends(strA, strB))
	case 4:
		os.Exit(0)
	}
	view.ShowSeparator()
}
