package main

import "fmt"

type Pet struct{
   name string
   age int
}
type Dog struct{
    p *Pet
}

func (d *Dog) Speak(){
	d.p.Speak()
}
func (d *Dog) SpeakTo(host string){
	d.p.SpeakTo(host);
}
func (p *Pet) Speak(){
	fmt.Println(p.name,p.age,"...")
}
func (p *Pet) SpeakTo(host string){
	p.Speak();
	fmt.Println(" ",host)
}
func main() {
	dog:=Dog{&Pet{name:"狗",
		   			age:5}}
	dog.SpeakTo("cat")
}
