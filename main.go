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

type statusMsg struct {
	status int
}

type errMsg struct {
	err error
}

func checkServer() tea.Msg {
	c := &http.Client{Timeout: time.Second * 10}
	res, err := c.Get(url)
	if err != nil {
		return errMsg{err}
	}
	return statusMsg{res.StatusCode}
}

func (e errMsg) Error() string { return e.err.Error() }

// We don't call the function bubbletea runtime will do that when the
// time is right
func (m model) Init() tea.Cmd {
	return checkServer
}
