package main

import (
	"net/http"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

const url = "https://charm.sh/"

type model struct {
	status int
	err    error
}

type (
	statusMsg int
	errMsg    struct {
		err error
	}
)

func checkServer() tea.Msg {
	c := &http.Client{Timeout: time.Second * 10}
	res, err := c.Get(url)
	if err != nil {
		return errMsg{err}
	}
	return statusMsg(res.StatusCode)
}

func (e errMsg) Error() string { return e.err.Error() }

// We don't call the function bubbletea runtime will do that when the
// time is right
func (m model) Init() tea.Cmd {
	return checkServer
}

// Internally `Cmd`s run in a goroutine, and the `Msg` they return is sent
// to the Update function for handling
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case statusMsg:
		m.status = int(msg)
		return m, tea.Quit
	case errMsg:
		m.err = msg
		return m, tea.Quit
	case tea.KeyMsg:
		if msg.Type == tea.KeyCtrlC {
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() string { return "" }
