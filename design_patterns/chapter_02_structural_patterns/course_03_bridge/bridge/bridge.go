package bridge

import (
	"errors"
	"fmt"
	"io"
)

type PrinterAPI interface {
	PrintMessage(string) error
}

type PrinterImpl1 struct {
}

func (p *PrinterImpl1) PrintMessage(msg string) error {
	fmt.Printf("%s\n", msg)
	return nil
}

type PrinterImpl2 struct {
	Writer io.Writer
}

func (p2 *PrinterImpl2) PrintMessage(msg string) error {
	if p2.Writer == nil {
		return errors.New("you need to pass an io.Writer to PrinterImpl2")
	}
	_, err := fmt.Fprintf(p2.Writer, "%s", msg)
	if err != nil {
		return err
	}
	return nil
}

type TestWriter struct {
	Msg string
}

func (t *TestWriter) Write(p []byte) (n int, err error) {
	n = len(p)
	if n > 0 {
		t.Msg = string(p)
		return
	}
	err = errors.New("content received on Writer was empty")
	return
}

type PrinterAbstraction interface {
	Print() error
}

type NormalPrinter struct {
	Msg     string
	Printer PrinterAPI
}

func (np *NormalPrinter) Print() error {
	err := np.Printer.PrintMessage(np.Msg)
	if err != nil {
		return err
	}
	return nil
}
