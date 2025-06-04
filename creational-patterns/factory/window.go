package main

type WindowGui struct {
	button string
}

func (w *WindowGui) setButton(button string) {
	w.button = button
}

func (w *WindowGui) getButton() string {
	return w.button
}
