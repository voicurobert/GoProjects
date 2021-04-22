package bridge

import (
	"strings"
	"testing"
)

func TestPrintAPI1(t *testing.T) {
	api1 := PrinterImpl1{}

	err := api1.PrintMessage("Hello")

	if err != nil {
		t.Errorf("Error trying to use the API1 implementation: Message: %s \n", err.Error())
	}
}

func TestPrintAPI2(t *testing.T) {
	api2 := PrinterImpl2{}

	err := api2.PrintMessage("Hello")
	if err != nil {
		expectedErr := "you need to pass an io.Writer to PrinterImpl2"
		if !strings.Contains(err.Error(), expectedErr) {
			t.Errorf("Error message was not correct.\n Actual :%s\n Expected: %s\n", err.Error(), expectedErr)
		}
	}

	testWriter := TestWriter{}
	api2 = PrinterImpl2{Writer: &testWriter}

	expectedMsg := "Hello"
	err = api2.PrintMessage(expectedMsg)
	if err != nil {
		t.Errorf("Error trying to use the API2 implementation: %s\n", err.Error())
	}

	if testWriter.Msg != expectedMsg {
		t.Fatalf("API2 did not write correctly  on the io.writer.\n Actual: %s\nExpected: %s\n", testWriter.Msg, expectedMsg)
	}

}

func TestNormalPrinter(t *testing.T) {
	expectedMsg := "Hello io.Writer"

	normalPrinter := NormalPrinter{
		Msg:     expectedMsg,
		Printer: &PrinterImpl1{},
	}

	err := normalPrinter.Print()
	if err != nil {
		t.Errorf(err.Error())
	}

	testWriter := TestWriter{}

	normalPrinter = NormalPrinter{
		Msg:     expectedMsg,
		Printer: &PrinterImpl2{Writer: &testWriter},
	}

	err = normalPrinter.Print()
	if err != nil {
		t.Errorf(err.Error())
	}

	if testWriter.Msg != expectedMsg {
		t.Fatalf("The expected message on the io.writer does not match.\n Actual: %s\nExpected: %s\n",
			testWriter.Msg, expectedMsg)
	}
}
