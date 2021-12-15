package main

import "fmt"

type Pet struct{

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
	fmt.Println("...")
}
func (p *Pet) SpeakTo(host string){
	p.Speak();
	fmt.Println(" ",host)
}
func main() {
	dog:=new(Dog)
	dog.SpeakTo("cat")
}
