package main

import "fmt"

//func declaration
//err handling
//return a func
//...
//defer
//panic
//recover

//func foo(a, b string, x, y int) (string, int) {
//	c := a + b
//	z := x + y
//	return c,z
//}
//
//func zero(x, _ int) int {
//	return 0
//}

//func foo(zeroOrOne int) (int,error){
//	switch zeroOrOne {
//	case 0:
//		fmt.Println("I'm zero")
//		return 0, nil
//	case 1:
//		fmt.Println("I'm one")
//		return 1, nil
//	default:
//		return -1, fmt.Errorf("Input must be 0 or 1")
//	}
//}

//func squares() func()  {
//	var count int
//	return func() {
//		count++
//		fmt.Printf("This function has executed %d times \n",count)
//	}
//}

//var m = make(map[string] int)
//var mu = sync.Mutex{}
//
//func lookup(key string) int{
//	mu.Lock()
//	defer mu.Unlock()
//	value := m[key]
//	return value
//}

//func sum (vals...int) int{
//	total := 0
//	for _,val := range vals{
//		total += val
//	}
//	return total
//}

//func boo(){
//	resp,err := http.Get("https://www.baidu.com")
//	if err != nil{
//		return
//	}
//	defer resp.Body.Close()
//	bytes, _ := ioutil.ReadAll(resp.Body)
//	data := string(bytes)
//	fmt.Println(data)
//}

func foo(zeroOrOne int){
	defer func(){
		if p := recover(); p != nil{
			fmt.Println(p)
		}
	}()
	switch zeroOrOne {
	case 0:
		fmt.Println("I'm zero")
	case 1:
		fmt.Println("I'm one")
	default:
		panic("Input muse be 0 or 1")
	}
	fmt.Println("OK")
}

func main(){
	//a := LearningGo.Point{}
	//p := &a
	//fmt.Println(a.X)
	//fmt.Println(p.X)

	//alan := LearningGo.Student{ID:0,Name:"Alan"}
	//fmt.Println(alan)

	//c := LearningGo.Circle{Central: LearningGo.Point{X:0,Y:0},Radius: 1.5}
	//fmt.Println(c)

	//fmt.Println(foo("a","b",0,1))

	//resp, err := http.Get("https://www.baidu.com")
	//if err != nil{
	//	log.Fatalln("Error to fetch this url")
	//	return
	//}
	//fmt.Println(resp)

	//r, err := foo(2)
	//if err != nil{
	//	fmt.Println(err)
	//}
	//fmt.Println(r)

	//fmt.Println(sum(1,2,3,4))

	foo(2)

}

