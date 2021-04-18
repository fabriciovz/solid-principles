package problem

type CannonPrinter struct {
}

func NewCannonPrinter() IMultiFunction {
	return &CannonPrinter{}
}

func(x *CannonPrinter) print() {
	// Real printing code starts here
}

func(x *CannonPrinter) getPrintSpoolDetails() {
	// Real print spool details code starts here
}

func(x *CannonPrinter) scan() {
	// Real scan code starts here
}

func(x *CannonPrinter) scanPhoto() {
	// Real photo scan code starts here
}

func(x *CannonPrinter) fax() {
	// Real fax code starts here
}

func(x *CannonPrinter) internetFax() {
	// Real internet fax code starts here
}
