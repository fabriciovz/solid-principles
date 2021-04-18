package problem

type XeroxWorkCentre struct {
}

func NewXeroxWorkCentre() IMultiFunction {
	return &XeroxWorkCentre{}
}

func(x *XeroxWorkCentre) print() {
	// Real printing code starts here
}

func(x *XeroxWorkCentre) getPrintSpoolDetails() {
	// Real print spool details code starts here
}

func(x *XeroxWorkCentre) scan() {
	// Real scan code starts here
}

func(x *XeroxWorkCentre) scanPhoto() {
	// Real photo scan code starts here
}

func(x *XeroxWorkCentre) fax() {
	// Real fax code starts here
}

func(x *XeroxWorkCentre) internetFax() {
	// Real internet fax code starts here
}