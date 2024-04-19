package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_NotifiesKeyListenersOfEvents(t *testing.T) {
	scn := setupDefaultScreen(t)

	g := game{
		scn: scn,
	}
	mock := mockKeyListener{}

	g.registerKeyListener(&mock)

	require.NoError(t, scn.PostEvent(&tcell.EventKey{}))
	require.NoError(t, scn.PostEvent(nil))

	g.pollEvents()

	mock.assertWasNotified(t)
}

type mockKeyListener struct {
	keyEvent *tcell.EventKey
}

func (m *mockKeyListener) notify(event *tcell.EventKey) {
	m.keyEvent = event
}

func (m *mockKeyListener) assertWasNotified(t *testing.T) {
	require.NotNil(t, m.keyEvent, "mockKeyListener should've been notified")
}
