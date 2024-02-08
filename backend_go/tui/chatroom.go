package tui

import (
	"fmt"
	"github.com/charmbracelet/log"
	// "net/http"
	"strings"
	// "time"

	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

/* This is the frontend for the chat websocket
It's a TUI, an alternative to the web frontend */

var messages []string

func StartChatroomInterface(wsUrl string) {
	log.Debug("Starting chatroom interface")

	// Connect to the websocket
	err := connectToWebsocket(wsUrl)
	if err != nil {
		log.Error("Failed to connect to websocket", "err", err)
	}

	// Listen for websocket messages
	go func() {
		for {
			_, msg, err := wsConn.ReadMessage()
			if err != nil {
				log.Error("Failed to read message from websocket", "err", err)
				return
			}
			log.Debug("Received message", "msg", string(msg))
			messages = append(messages, string(msg))
		}
	}()

	// Create the TUI
	p := tea.NewProgram(initialModel())
	// Run the TUI
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

type model struct {
	viewport    viewport.Model
	textarea    textarea.Model
	senderStyle lipgloss.Style
	err         error
}

type errMsg error

func initialModel() model {
	ta := textarea.New()
	ta.Placeholder = "Send a message..."
	ta.Focus()

	ta.Prompt = "┃ "
	ta.CharLimit = 280

	ta.SetWidth(30)
	ta.SetHeight(3)

	// Remove cursor line styling
	ta.FocusedStyle.CursorLine = lipgloss.NewStyle()

	ta.ShowLineNumbers = false

	vp := viewport.New(30, 5)
	vp.SetContent(`Welcome to the chat room!
Type a message and press Enter to send.`)

	ta.KeyMap.InsertNewline.SetEnabled(false)

	return model{
		textarea:    ta,
		viewport:    vp,
		senderStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("5")),
		err:         nil,
	}
}

func (m model) Init() tea.Cmd {
	return textarea.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		tiCmd tea.Cmd
		vpCmd tea.Cmd
	)

	m.viewport.SetContent(strings.Join(messages, "\n"))

	m.textarea, tiCmd = m.textarea.Update(msg)
	m.viewport, vpCmd = m.viewport.Update(msg)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			fmt.Println(m.textarea.Value())
			return m, tea.Quit
		case tea.KeyEnter:
			m.viewport.SetContent(strings.Join(messages, "\n"))
			m.viewport.GotoBottom()
			sendMessage(m.textarea.Value())
			m.textarea.Reset()
		}

	// We handle errors just like any other message
	case errMsg:
		m.err = msg
		return m, nil
	}

	return m, tea.Batch(tiCmd, vpCmd)
}

func (m model) View() string {
	return fmt.Sprintf(
		"%s\n\n%s",
		m.viewport.View(),
		m.textarea.View(),
	) + "\n\n"
}
