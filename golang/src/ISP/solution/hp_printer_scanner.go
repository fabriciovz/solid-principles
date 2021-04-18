package solution

type HPPrinterNScanner struct {
}
type IHPPrinterNScanner interface {
	IPrint
	IScan
}
func NewHPPrinterNScanner() IHPPrinterNScanner {
	return &HPPrinterNScanner{}
}

func(x *HPPrinterNScanner) print() {
	// Real printing code starts here
}

func(x *HPPrinterNScanner) getPrintSpoolDetails() {
	// Real print spool details code starts here
}

func(x *HPPrinterNScanner) scan() {
	// Real scan code starts here
}

func(x *HPPrinterNScanner) scanPhoto() {
	// Real photo scan code starts here
}
