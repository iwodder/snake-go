package ui

import (
	"testing"

	"github.com/gdamore/tcell/v2"
	"github.com/stretchr/testify/require"
)

func Test_Manager(t *testing.T) {
	newKeyHandler := func() (func(*tcell.EventKey), *bool) {
		wasRun := new(bool)
		return func(key *tcell.EventKey) {
			*wasRun = true
		}, wasRun
	}

	t.Run("manager has no views when created", func(t *testing.T) {
		require.Empty(t, NewManager().views)
	})

	t.Run("first view added is the active view", func(t *testing.T) {
		mgr := NewManager()
		view := MockView{}
		viewName := "FirstView"

		mgr.AddView(viewName, &view)

		require.Equal(t, viewName, mgr.ActiveViewName())
		require.Equal(t, &view, mgr.view())
	})

	t.Run("active view is updated when view is switched", func(t *testing.T) {
		mgr := NewManager()
		mockViewOne := MockView{}
		mockViewTwo := MockView{}
		viewOne := "FirstView"
		viewTwo := "SecondView"

		mgr.AddView(viewOne, &mockViewOne)
		mgr.AddView(viewTwo, &mockViewTwo)

		err := mgr.SwitchView(viewTwo)

		require.NoError(t, err)
		require.Equal(t, viewTwo, mgr.ActiveViewName())
		require.Equal(t, &mockViewTwo, mgr.view())
	})

	t.Run("error is returned when switching to non-existent view", func(t *testing.T) {
		mgr := NewManager()
		viewName := "NonExistentView"

		err := mgr.SwitchView(viewName)

		require.ErrorIs(t, err, ErrViewNotFound)
	})

	t.Run("adding view to improperly created manager doesn't panic", func(t *testing.T) {
		mgr := Manager{}
		viewName := "TestView"
		view := MockView{}

		require.NotPanics(t, func() {
			mgr.AddView(viewName, &view)
		})
	})

	t.Run("switching views on an improperly created manager errors", func(t *testing.T) {
		mgr := Manager{}
		viewName := "TestView"

		require.ErrorIs(t, mgr.SwitchView(viewName), ErrNoViews)
	})

	t.Run("events are not propagated when view isn't active", func(t *testing.T) {
		mgr := NewManager()
		handler, wasRun := newKeyHandler()

		mgr.SetKeyEventCallback(handler)

		mgr.Handle(tcell.NewEventKey(tcell.KeyRune, 'a', tcell.ModNone))

		require.False(t, *wasRun)
	})

	t.Run("events are propagated to key handler when view is active", func(t *testing.T) {
		mgr := NewManager()
		handler, wasRun := newKeyHandler()

		mgr.AddView("TestView", &MockView{})
		mgr.SetKeyEventCallback(handler)

		mgr.Handle(tcell.NewEventKey(tcell.KeyRune, 'a', tcell.ModNone))

		require.True(t, *wasRun)
	})

	t.Run("events are propagated to active view", func(t *testing.T) {
		view := MockView{}
		mgr := NewManager()
		mgr.AddView("TestView", &view)

		mgr.Handle(tcell.NewEventKey(tcell.KeyRune, 'a', tcell.ModNone))

		view.assertWasNotified(t)
	})

	t.Run("calling draw on a manager with no views does nothing", func(t *testing.T) {
		mgr := NewManager()

		mgr.Draw(tcell.NewSimulationScreen("UTF-8"))
	})

	t.Run("calling draw on a manager with views draws them", func(t *testing.T) {
		mgr := NewManager()
		view := MockView{}
		mgr.AddView("TestView", &view)

		mgr.Draw(tcell.NewSimulationScreen("UTF-8"))
		view.assertWasDrawn(t)
	})

	t.Run("showing modal reports modal is visible", func(t *testing.T) {
		mgr := NewManager()

		mgr.ShowModal("My Modal")

		require.True(t, mgr.ModalVisible())
	})

	t.Run("showing and then hiding modal reports modal is not visible", func(t *testing.T) {
		mgr := NewManager()

		mgr.ShowModal("My Modal")
		mgr.HideModal()

		require.False(t, mgr.ModalVisible())
	})

	t.Run("height and width report 0 when no views are active", func(t *testing.T) {
		mgr := NewManager()

		require.Equal(t, 0, mgr.Height())
		require.Equal(t, 0, mgr.Width())
	})
}

type MockView struct {
	s tcell.Screen
	e tcell.Event
	composite
}

func (m *MockView) Width() int {
	return 0
}

func (m *MockView) Height() int {
	return 0
}

func (m *MockView) Draw(s tcell.Screen) {
	m.s = s
}

func (m *MockView) Handle(e tcell.Event) {
	m.e = e
}

func (m *MockView) assertWasNotified(t *testing.T) {
	require.NotNil(t, m.e)
}

func (m *MockView) assertWasDrawn(t *testing.T) {
	require.NotNil(t, m.s)
}
