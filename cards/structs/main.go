package main

import (
	"fmt"
)

type pessoa struct{
	nome string
	sobrenome string
	contato contactInfo // ou somente contactInfo, neste caso o campo seria referenciado 
						// como contactInfo
}

type contactInfo struct {
	email string
	zip int
}

func (p pessoa) print(){
	fmt.Println(p) //{Alex Anderson {test@gmail.com 123456789}}
}

func (ponteiroPessoa *pessoa) updateName(name string){
	ponteiroPessoa.nome = name

}

func main() {

	
	name := "bill"
 
	namePointer := &name
	
	fmt.Println(&namePointer)
	printPointer(namePointer)
   }
	
   func printPointer(namePointer *string) {
	fmt.Println(&namePointer)
   }

