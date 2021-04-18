package main

import solution_lsp "github.com/fabriciovz/golang/src/LSP/solution"

//type Product struct {
//	discount float64
//}
//
func main() {

	solution_lsp.Run()

}

///////////////////////////
//type A struct {
//}
//func (a A) Test() {
//	fmt.Println("Printing A")
//}
//
/////////////////////////////
//type B struct {
//	A
//}
///////////////////////
//
//func ImpossibleLiskovSubstitution(a A) {
//	a.Test()
//}
//func main() {
//
//	a := B{}
//	ImpossibleLiskovSubstitution(a) // PANIC : cannot use a (type B) as
//	//type A in argument to
//	//ImpossibleLiskovSubstitution
