package ui

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
)

var (
	ErrViewNotFound = fmt.Errorf("view not found")
	ErrNoViews      = fmt.Errorf("no views registered")
)

// View is a UI surface that can both draw and receive events.
type View interface {
	Component
	EventHandler
}

// Manager controls which View is active and provides common UI behavior
// like global key callbacks and modal overlay messages.
type Manager struct {
	views            map[string]View
	activeName       string
	active           View
	keyEventCallback func(*tcell.EventKey)
	modal            struct {
		isActive bool
		text     string
	}
}

func NewManager() *Manager {
	return &Manager{
		views: make(map[string]View),
	}
}

// AddView registers a view under a name. If this is the first view added,
// it becomes the active view automatically.
func (m *Manager) AddView(name string, v View) {
	if m.views == nil {
		m.views = make(map[string]View)
	}
	m.views[name] = v
	if m.active == nil {
		m.activeName = name
		m.active = v
	}
}

// SwitchView makes the named view active.
func (m *Manager) SwitchView(name string) error {
	if m.views == nil {
		return ErrNoViews
	}
	v, ok := m.views[name]
	if !ok {
		return fmt.Errorf("%w: %q", ErrViewNotFound, name)
	}
	m.activeName = name
	m.active = v
	return nil
}

func (m *Manager) ActiveViewName() string {
	return m.activeName
}

func (m *Manager) view() View {
	return m.active
}

// SetKeyEventCallback registers an optional global key handler that runs before the active view receives the key.
func (m *Manager) SetKeyEventCallback(callback func(*tcell.EventKey)) {
	m.keyEventCallback = callback
}

func (m *Manager) Handle(ev tcell.Event) {
	if m.active == nil {
		return
	}

	if keyEv, ok := ev.(*tcell.EventKey); ok {
		if m.keyEventCallback != nil {
			m.keyEventCallback(keyEv)
		}
	}

	m.active.Handle(ev)
}

func (m *Manager) Draw(scrn tcell.Screen) {
	if m.active == nil {
		return
	}

	m.active.Draw(scrn)

	if m.modal.isActive {
		ShowMessage(m.active, m.modal.text, scrn)
	}
}

func (m *Manager) ShowModal(text string) {
	m.modal.text = text
	m.modal.isActive = true
}

func (m *Manager) HideModal() {
	m.modal.text = ""
	m.modal.isActive = false
}

func (m *Manager) ModalVisible() bool {
	return m.modal.isActive
}

func (m *Manager) Width() int {
	if m.active == nil {
		return 0
	}
	return m.active.Width()
}

func (m *Manager) Height() int {
	if m.active == nil {
		return 0
	}
	return m.active.Height()
}
