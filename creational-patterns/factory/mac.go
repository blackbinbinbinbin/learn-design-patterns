package main

type MacGui struct {
	button string
}

func (m *MacGui) setButton(button string) {
	m.button = button
}

func (m *MacGui) getButton() string {
	return m.button
}
