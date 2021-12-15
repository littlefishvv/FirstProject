package main

func main() {
	testPainc()
}
func testPainc(){
	defer func(){
		if err:=recover();err!=nil{
			println(err.(string))
		}
	}()
	panic("test panic")
}
