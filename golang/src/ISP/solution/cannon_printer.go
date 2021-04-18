package solution

type CannonPrinter struct {
}

func NewCannonPrinter() IPrint {
	return &CannonPrinter{}
}

func(x *CannonPrinter) print() {
	// Real printing code starts here
}

func(x *CannonPrinter) getPrintSpoolDetails() {
	// Real print spool details code starts here
}