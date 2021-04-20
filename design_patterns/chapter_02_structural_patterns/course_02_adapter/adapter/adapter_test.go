package adapter

import "testing"

func TestAdapter(t *testing.T) {
	msg := "Hello World"

	adapter := PrinterAdapter{
		OldPrinter: &MyLegacyPrinter{},
		Msg:        msg,
	}

	retMsg := adapter.PrintStored()

	if retMsg != "Legacy Printer: Adapter: Hello World" {
		t.Errorf("Message didn't match: %s", retMsg)
	}
}
