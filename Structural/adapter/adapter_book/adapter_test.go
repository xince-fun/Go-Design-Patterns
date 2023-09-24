package adapter

import "testing"

func TestAdapter(t *testing.T) {
	msg := "Hello World!"
	adapter := PrinterAdapter{OldPrinter: &MyLegayPrinter{}, Msg: msg}
	returnedMsg := adapter.PrintStored()

	if returnedMsg != "Legacy Printer: Adapter: Hello World!\n" {
		t.Errorf("Message didn't match: %s\n", returnedMsg)
	}
}
