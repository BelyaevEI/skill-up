package main

type Printer interface {
	Print(document string)
}

type Scanner interface {
	Scan(document string)
}

// MultiFunctionDevice наследует от обоих интерфейсов
type MultiFunctionDevice interface {
	Printer
	Scanner
}

// класс, реализующий только функцию печати
// Не заставляем SimplePrinter реализовывать функции сканирования, которые он не использует, соблюдая ISP.
type SimplePrinter struct{}

func (p *SimplePrinter) Print(document string) {
	// реализация печати
}

// класс, реализующий обе функции
type AdvancedPrinter struct{}

func (p *AdvancedPrinter) Print(document string) {
}

func (p *AdvancedPrinter) Scan(document string) {
}
