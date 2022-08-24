package tui

import "github.com/gdamore/tcell/v2"

type IScreen interface {
	tcell.Screen
}

type ISimulationScreen interface {
	tcell.SimulationScreen
}
