package main

import (
	//	"diy_package/string" //under GOPATH
	"errors"
	"fmt"
	"runtime"
	"sync/atomic"
	"time"

	"math/rand"
	"os"
	"sort"
	"strings"
	"sync"
)

func jsonMarshal() {

	type Response1 struct {
		Page   int
		Fruits []string
	}
	type Response2 struct {
		Page   int      `json:"page,d"`
		Fruits []string `json:"fruits"`
	}
	println(buffer)
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))
	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))
	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))
	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	res1D := &Response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	res2D := &Response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	// println(res2D)
	fmt.Printf("\n%+v", res2D)
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))
	// var tmp interface{}
	var tmp map[string]interface{}
	if err := json.Unmarshal([]byte(res2B), &tmp); err != nil {
		panic(err)
	}
	fmt.Printf("\n%+v", tmp)

	type ColorGroup struct {
		ID     int `json:"id"`
		Name   string
		Colors []string
	}
	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}

	b, err := json.Marshal(group)
	if err != nil {
		fmt.Println("error:", err)

	}
	os.Stdout.Write(b)
	var tmp interface{}
	if err := json.Unmarshal([]byte(b), &tmp); err != nil {
		panic(err)
	}
	fmt.Printf("\n%+v", tmp)
	// println("aaa: ", tmp)
	// tmp := make(map[int]string)
	// tmp[1] = "a"
	// tmp[2] = "b"
	// delete(tmp, 1)
	// fmt.Println(tmp)

}
func helloWorld() {
	fmt.Println("like print with newline ending")
	fmt.Println("-------------------------------")
	fmt.Print(3, "ouput directly", 333)
	fmt.Println("\n-------------------------------\n")
	fmt.Printf("format ouput %d", 100)
	// num := 'a'

	// for i := 0; i < 10; i++ {
	// 	fmt.Printf("number of %c:%d", num, num)
	// 	num++
	//
	// }
	fmt.Println(`what you see is what you get
	fjeifjewifwe
	fjewjfiew
	new line 
	fefj
	`)
}
func values() {
	fmt.Println("go" + "lang")
	fmt.Println("3+3=", 3+3)
	fmt.Println("3.0/43=", 3.0/43)
	fmt.Println("true&&false", true && false)
	fmt.Println("true||false", true || false)
	fmt.Println(!true)
}
func variables() {
	var a, b int = 1, 2
	fmt.Println(a, b)
	var e int
	fmt.Println(e)
}
func forloop() {
	i := 1
	for i <= 10 {
		fmt.Println("i is ", i)
		i++
	}
	fmt.Println()
	for i := 10; i <= 20; i++ {
		if i == 15 {
			fmt.Println("stop point ", i)
			break
		} else if i == 12 { //else must be follow }
			fmt.Println("jump point ")
			continue
		} else {
			fmt.Println("i is ", i)
		}
	}
}

func ifElse() {
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

}
func Arrays() {
	var a [6]int
	fmt.Println("empty array", a)
	b := [6]int{1, 2, 3, 4, 6, 65}
	fmt.Println("decleare array", b)
	var twoDimetion [9][9]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			fmt.Print(i, "*", j, "=", i*j)
			twoDimetion[i][j] = i * j
		}
	}
}

func slice() {
	_rune := make([]rune, 3)
	character := 'a'
	for name := range _rune {
		_rune[name] = character
		// _rune[name] = character++  not support
		character++
		fmt.Println(name, "byte:", _rune)
	}
	fmt.Printf("get slice: %c\n", _rune)

	_rune = append(_rune, 'd', 'e')
	fmt.Printf("append slice: %c", _rune)

	c := make([]rune, len(_rune))
	copy(c, _rune)

	fmt.Printf("copy slice: %c", _rune)

	l := _rune[2:4]
	fmt.Println("slice 2:4", l)
	l = _rune[:3] //can't use  := because l is new by := already
	fmt.Println("slice :4", l)
	l = _rune[2:]
	fmt.Println("slice 2:", l)

	twoD := make([][]int, 9)
	for i := 0; i < 9; i++ {
		innerLenth := i + 1
		twoD[i] = make([]int, innerLenth)
		for j := 0; j < innerLenth; j++ {
			twoD[i][j] = (1 + i) * (j + 1)
		}
	}
	for i := 0; i < len(twoD); i++ {
		fmt.Println((i + 1), ":", twoD[i], " length of this dimetion:", len(twoD[i]))
	}

}
func maps() {
	maps := make(map[string]int) //type maps
	maps["jialin"] = 27
	maps["nonglei"] = 23
	fmt.Println("maps:", maps, " length:", len(maps))

	map_1 := maps["jialin"]
	fmt.Println("element:", map_1)
	map_1 = 100
	fmt.Println("element:", map_1)

	delete(maps, "nonglei")
	fmt.Println("delete map:", maps)
	values, pairOrNot := maps["jialin"]
	fmt.Println("value:", values, " pair?:", pairOrNot)

	mapinital := map[string]int{"a": 1, "b": 2}
	fmt.Println("map:", mapinital)

	maps["demo"] = 10
	map_c := make(map[string]int)
	for key, value := range maps {
		//map_c[key] = maps[key] //same to map_c[key] = value
		map_c[key] = value
	}
	sliceMaps := make([]map[string]int, 3)
	sliceMaps[0] = maps
	sliceMaps[1] = map_c
	//sliceMaps[2] = map_c
	for name := range sliceMaps {
		fmt.Println(sliceMaps[name])
	}

}
func ranges() {
	nums := []string{"1", "fjjewif", "ijefel"}
	for index, value := range nums {
		fmt.Println("index:", index, "value:", value)
	}
	for i, c := range "ab cd " {
		fmt.Printf("%d->%c\n", i, c)
		// fmt.Println(i, rune(c))
	}
	for key, value := range map[string]int{"a": 1, "b": 2, "c": 3} {
		fmt.Println("key:", key, " value:", value)
	}
}
func variadicParameter(nums ...int) {
	fmt.Println(nums, " ")
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println(sum)
}
func closures() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}
func pointer(i *int) {
	(*i)++
	fmt.Println("address of i in function: ", i)
}

type person struct {
	name   string
	age    int
	height float64
	weight float64
}

type consumption interface {
	decoration() float64
	tools() float64
}
type boy struct {
	name       string
	computer   float64
	cloth      float64
	priceWeith float64
	total      float64
}
type girl struct {
	name       string
	computer   float64
	cloth      float64
	priceWeith float64
	total      float64
}

/*
if use interface can't use pointer type*/
func (sp boy) decoration() float64 {
	return sp.cloth * sp.priceWeith
}
func (sp girl) decoration() float64 {
	return sp.cloth * sp.priceWeith
}
func (sp girl) tools() float64 {
	return sp.computer * sp.priceWeith
}
func (sp boy) tools() float64 {
	return sp.computer * sp.priceWeith
}
func measure(ip consumption) { //can't use pointer type

	fmt.Println("values of member: ", ip)
	fmt.Println("cost of decoration:", ip.decoration())
	fmt.Println("cost of tools:", ip.tools())
}

// BMI Categories:
// Underweight = <18.5
// Normal weight = 18.5–24.9
// Overweight = 25–29.9
// Obesity = BMI of 30 or greater
func (sp *person) BMI() (string, float64) {
	BMI := sp.weight / (sp.height * sp.height)
	switch {
	case BMI < 18.5:
		return "under weight", BMI
	case (18.5 <= BMI && BMI <= 24.9):
		return "normal weight", BMI
	case 24.9 < BMI && BMI <= 29.9:
		return "over weight", BMI
	case 30 < BMI:
		return "over weight", BMI
	}
	return " ", 0.0
}

func Errors(arg int) (int, error) {
	if arg == 2 {
		return -1, errors.New("error occurs with number 2")
	}
	return arg * 10, nil
}
func Errors2(arg int) (int, error) {
	if arg == 2 {
		//  &argError{arg, "error occurs with number 2"}  it's type is error? calling argError's Error()

		return -1, &argError{arg, "detail(error occurs with number 2)"}
	}
	return arg * 10, nil
}

type argError struct {
	arg     int
	details string
}

func (sp *argError) Error() string {
	return fmt.Sprintf("%d -demo- %s ", sp.arg, sp.details)
}

/*goroutines just like multi thread*/
func goroutines(str string) {
	byt := []rune(str)
	for name := range byt {
		fmt.Printf("%c ", byt[name])
	}
}
func channel(messageChannel chan string) {
	go func() {
		messageChannel <- "string from channel function"
	}()
	msg := <-messageChannel
	// msg2 := <-messageChannel   //one comes only one outs
	fmt.Println(msg)
	// fmt.Println(msg2)
}

/* if the channel's buffer is 2,channel's write port is lock when two message finish writting,which means only reading can work and no writting more for that time*/
func channelBuffer(messageChannel chan string) {
	go func() {

		messageChannel <- "writer: message one"
		fmt.Println("write firtst msg done!")
		messageChannel <- "writer: message two"
		fmt.Println("write second msg done!")
		messageChannel <- "writer: message three"
		fmt.Println("write third msg done!")
	}()
	// msg := <-messageChannel
	// fmt.Println("reader:", msg)
	// //	time.Sleep(time.Second * 3)
	// msg = <-messageChannel
	// fmt.Println("reader:", msg)
	// msg = <-messageChannel
	// fmt.Println("reader:", msg)

}
func channelSynchronization(cha chan bool) {
	fmt.Println("writting!")
	cha <- true //try put this line under next two line
	fmt.Println("write done!")
	//cha <- true //in this position can also work,why, because output's time is shortter and the synchronization process?   first output then sleep
	time.Sleep(time.Second * 3)
}

/* if use channel direction, the channel parameter must be declear with 1 or more buffer directly*/
func channelDirection(sender <-chan string, receiver chan<- string) {
	// receiver <- "msg for receiver"
	msg := <-sender
	receiver <- msg
}
func chnnelRceiver(cha chan<- string, str string) {
	cha <- str
}
func selectF() {
	c1 := make(chan string, 1)
	c2 := make(chan string, 1)
	go func() {
		// time.Sleep(time.Second * 1)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(time.Second * 3)
		c2 <- "two"
	}()
	for i := 0; i < 2; i++ {

		time.Sleep(time.Second * 1)
		select {
		case msg1 := <-c2:
			fmt.Println("channel :", msg1)
		case msg2 := <-c1:
			fmt.Println("channel :", msg2)
		default:
			fmt.Println("no channel comming")
		}

	}
}
func timeout() {
	cha1 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		cha1 <- "channel one"
	}()
	select {
	case msgOne := <-cha1:
		fmt.Println(msgOne)
	case <-time.After(time.Second * 2):
		fmt.Println("one timeout")
	}
	cha2 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		cha2 <- "msgTwo"
	}()
	select {
	case msgtwo := <-cha2:
		fmt.Println(msgtwo)
	case <-time.After(time.Second * 1):
		fmt.Println("two timeout")
	}
}

//Closing a channel indicates that no more values will be sent on it.
func closeChannel() {
	jobs := make(chan int, 6)
	done := make(chan bool)
	go func() {
		for i := 0; i < 10; i++ {
			tmpInt, open := <-jobs
			if open {
				fmt.Println("received job", tmpInt)
			} else {
				fmt.Println("all jobs received")
				done <- true // as all channel done signal
				return
			}
		}
	}()
	for i := 0; i < 11; i++ {
		jobs <- i
		fmt.Println("sent job", i)
		if i == 7 {
			close(jobs)
			break
		}
	}
	<-done //waiting for all channel exit
}
func rangeChannel() {
	jobs := make(chan int, 6)
	done := make(chan bool)
	go func() {
		for i := 0; i < 11; i++ {
			jobs <- i
			fmt.Println("sent job", i)
			if i == 7 {
				close(jobs)
				done <- true
				break
			}
		}
	}()
	for element := range jobs {
		fmt.Println(element)
	}
	<-done
}

func ticker() {
	//when try microsecond it seems  obscure
	tick := time.NewTicker(time.Millisecond * 100)
	// tick := time.NewTicker(time.Microsecond * 100)
	tickF := func() {
		for time_ := range tick.C {
			fmt.Println("Tick at ", time_, " so you can make a schedule in this for loop")
		}
	}
	go tickF()
	time.Sleep(time.Millisecond * 600)
	tick.Stop()
	fmt.Println("ticker stopped")

}
func timer() {
	time1 := time.NewTimer(time.Second)
	<-time1.C
	fmt.Println("time one expired,times up")
	time2 := time.NewTimer(time.Second)
	go func() {
		<-time2.C
		//it won't be time expired,because stop() involed alreader before that happen
		fmt.Println("time two expired")
	}()
	time.Sleep(time.Second * 2)
	stopped := time2.Stop()
	if stopped {
		fmt.Println("timer two is stopped")
	}

}
func factory(id int, creator <-chan int, customer chan<- int) {
	for singleJob := range creator {
		fmt.Println("creator", id, "th processing job with data ", singleJob)
		time.Sleep(time.Millisecond * 500)
		customer <- singleJob * 100
	}
}

//actually,just make a time channel's cycle change
func rateLimit() {
	requests := make(chan int, 6)
	for i := 0; i < 6; i++ {
		requests <- i
	}
	close(requests) //no more requests can be received
	responseCycle := (time.Tick(time.Second))
	for item := range requests {
		<-responseCycle
		fmt.Println("respond to ", item, "at ", time.Now())
	}

	burstyResponseCycle := make(chan time.Time, 3) //burstyLimit
	for i := 0; i < 3; i++ {
		burstyResponseCycle <- time.Now() //make the first three cycle run immediately
	}
	//fill the burstyResponseCycle immediately when it is not full
	go func() {
		for t := range time.Tick(time.Second) {
			burstyResponseCycle <- t
		}
	}()

	burstyRequests := make(chan int, 6)
	for i := 0; i < 6; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for burstyRes := range burstyRequests {
		<-burstyResponseCycle //once get out of one Cycle from burstyResponseCycle  ,the goroutine above 12th line will fill the burstyResponseCycle
		fmt.Println("respond burstly to ", burstyRes, "at time ", time.Now())
	}
}

func atomicF() {
	var counter uint64 = 0
	for i := 0; i < 50; i++ {
		go func() {
			for {
				atomic.AddUint64(&counter, 1)
				runtime.Gosched() //Allow other goroutines to proceed.
			}
		}()
	}
	time.Sleep(time.Millisecond * 100) //counter adding in one second
	//In order to safely use the counter while it’s still being updated by other goroutines, we extract a copy of the current value into counterFinal via LoadUint64.
	counterFinal := atomic.LoadUint64(&counter)
	fmt.Println("final counter is ", counterFinal)
}
func mutexF() {
	state := make(map[int]int)
	mutex := &sync.Mutex{}
	var counter uint64 = 0
	total := 0
	//100 goroutines to keep reading and calulate the 5 state randly
	for i := 0; i < 100; i++ {
		go func() {
			tmpTotal := 0
			for {
				key := rand.Intn(5)
				mutex.Lock()
				total += state[key]
				tmpTotal += state[key]
				fmt.Println("goroutine ", i, " tmptotal number:", tmpTotal)
				mutex.Unlock()
				atomic.AddUint64(&counter, 1)
				runtime.Gosched()
			}
		}()
	}
	//10 goroutines to keep writing the 5 state randly
	for i := 0; i < 10; i++ {
		go func() {
			for {
				key := rand.Intn(5)
				val := rand.Intn(100)
				mutex.Lock()
				state[key] = val
				mutex.Unlock()
				atomic.AddUint64(&counter, 1)
				runtime.Gosched()
			}
		}()
	}

	//do all above in one second
	time.Sleep(time.Second)
	counterFinal := atomic.LoadUint64(&counter)
	fmt.Println("final counter is ", counterFinal)
	mutex.Lock()
	fmt.Println("final states are", state, " total number: ", total)
	mutex.Unlock()

}
func sortF() {
	strs := []string{"c", "fe", "e", "a"}
	sort.Strings(strs)
	fmt.Println("sorted string ", strs)

	isSorted := sort.StringsAreSorted(strs)
	fmt.Println("sorted string? ", isSorted)
}

type Str []string

//We implement sort.Interface - Len, Less, and Swap - on our type so we can use the sort package’s generic Sort function.
//can also add function for []string
func (s Str) Len() int {
	return len(s)
}
func (s Str) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s Str) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func stateful() {
	type reader struct {
		key  int
		resp chan int
	}
	type writer struct {
		key     int
		val     int
		respVal int
		resp    chan bool
	}
	var opsCounter uint64 = 0
	readP := make(chan *reader)
	writeP := make(chan *writer)
	go func() {
		//This goroutine repeatedly selects on the reads and writes channels, responding to requests as they arrive. A response is executed by first performing the requested operation and then sending a value on the response channel resp to indicate success
		var state = make(map[int]int)
		for {
			//channel's buffer block guarantee that responder work before reveiver work
			select {
			case read := <-readP:
				// fmt.Println("response to reader with value ", state[read.key])
				read.resp <- state[read.key] //this will block until state[key] if full(buffer is 1) which means the writer is finish write the state[key], the state[key].value comes from writer bellow
			case write := <-writeP:
				state[write.key] = write.val
				write.respVal = write.val * 10
				fmt.Println("respond to writer for  value ", write.val)
				write.resp <- true
			}
		}
	}()
	for i := 0; i < 100; i++ {
		go func() {
			read := &reader{
				key:  -1,
				resp: make(chan int)} //channel buffer is 1
			readP <- read //select read case will catch readP
			// value := <-read.resp //this will block until read.resp have been responded
			<-read.resp //this will block until read.resp have been responded
			// fmt.Println("receive from reader with value ", value)
			atomic.AddUint64(&opsCounter, 1)
		}()
	}

	for i := 0; i < 100; i++ {
		go func() {
			for {
				write := &writer{
					key:     rand.Intn(5),
					val:     rand.Intn(100),
					respVal: -1,
					resp:    make(chan bool)}
				writeP <- write //invoke the select write case:
				<-write.resp    //wait for the write.resp to be true in select
				fmt.Println("receive from select resond with value ", write.respVal)
				atomic.AddUint64(&opsCounter, 1)
			}
		}()
	}
	time.Sleep(time.Second)
	opsCounterFinal := atomic.LoadUint64(&opsCounter)
	fmt.Println("final operation counter is :", opsCounterFinal)
}

//Note that unlike some languages which use exceptions for handling of many errors, in Go it is idiomatic to use error-indicating return values wherever possible.
//A common use of panic is to abort if a function returns an error value that we don’t know how to (or want to) handle.
func panicF() {
	// panic("a problem occurs")
	fileDescription, err := os.Create("./tmp.file")
	if err != nil {
		panic(err)
	}
	fmt.Println(fileDescription)
	// var str string
	// str = "good demo"
	// str := "goodfjewjfiewmeo"
	// retLen, err := whatisit.WriteString(str)
	// fmt.Println(whatisit, retLen, err)
}

//usually for purposes of cleanup. defer is often used where e.g. ensure and finally would be used in other languages.
func deferF() {
	fd := createFile("./c.file")
	defer closeFile(fd)
	writeToFile(fd)
}
func createFile(p string) *os.File {
	fd, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return fd
}
func writeToFile(fd *os.File) {
	fmt.Println("writing to file", fd)
	fmt.Fprintf(fd, "write data to file %x", fd)
}
func closeFile(fd *os.File) {
	fmt.Println("closing file ", fd)
	fd.Close()
}
func Index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}
func Include(vs []string, t string) bool {
	return Index(vs, t) >= 0
}

// f func(string) bool : like a function pointer
func findByAnyPrefixChacter(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}
func Delete(vs []string, fs string) []string {
	fmt.Println("original string: ", vs)
	vsNew := make([]string, 0) //string slice
	for _, v := range vs {
		if fs == v {
			continue
		}
		vsNew = append(vsNew, v)
	}
	return vsNew
}
func MapIndex(vs []string, f func(string) bool) map[string]int {
	vsmap := make(map[string]int)
	for i, v := range vs {
		if f(v) {
			vsmap[v] = i
		}
	}
	return vsmap
}

const s = "jialin wu"
const tmpNumber = 2774534854

func main() {

	// helloWorld()
	//
	// str := "just a dmeo"
	// str_byte := []byte(str)
	// for i := 0; i < len(str_byte); i++ {
	// 	fmt.Printf("%c\n", str_byte[i])
	// }
	// //fmt.Println(string.Reverse("Hello, golanger"))
	// values()
	// variables()
	// fmt.Println(s, tmpNumber)
	// forloop()
	// ifElse()
	// slice()
	//
	// maps()
	// ranges()

	// 	num := []int{1, 23, 4, 6, 5}
	// 	variadicParameter(num...)
	// 	variadicParameter(4, 56546, 456, 45, 654)

	/*:just like all elements in the function closures is keep update until recall it which will restore all the element
	 */
	//
	// nextInt := closures() //addres of anonymous function
	// fmt.Println(nextInt)
	// fmt.Println(nextInt()) //call the anonymous function
	// fmt.Println(nextInt())
	// fmt.Println(nextInt())
	// nextInt_b := closures()
	// a := nextInt_b()
	// fmt.Println(nextInt_b)
	// fmt.Printf("%d", a)

	// i := 0
	// pointer(&i)
	// pointer(&i)
	// fmt.Println(i)
	// pointer(&i)
	// fmt.Println(i)
	// fmt.Println("address of i: out of function", &i)
	//
	//
	// s := person{name: "jialin", height: 1.65, age: 23, weight: 50}
	// fmt.Println(s)
	// fmt.Println(s.name)
	// as := &s
	// fmt.Println(as.name)
	// as.name = "jialin wu"
	// fmt.Println(s.name)
	// fmt.Println(s.BMI())

	// b := boy{name: "boy", priceWeith: 1.0, cloth: 80, computer: 10000}
	// g := girl{name: "girl", priceWeith: 2.0, cloth: 180, computer: 2000}
	// measure(b)
	// measure(g)
	// for _, name := range []int{1, 100} {
	// 	if retValue, e := Errors(name); e != nil {
	// 		fmt.Println("failed:", e)
	// 	} else {
	// 		fmt.Println("worked:", retValue)
	// 	}
	// }
	//
	// for name, _ := range [10]int{0} {
	// 	if retValue, e := Errors2(name); e != nil {
	// 		fmt.Println("failed:", e) //this will call Error()
	// 		if argErrorPointer, returnNormal := e.(*argError); returnNormal {
	// 			fmt.Println("failed:", argErrorPointer.arg, argErrorPointer.details)
	// 		}
	// 	} else {
	// 		fmt.Println("worked:", retValue)
	// 	})
	// }

	// go goroutines("jialinwu")
	// go func(str string) {
	// 	fmt.Println("just a demo fo goroutine:", str)
	// }("hahahhaha")
	// go goroutines("the world is but a little quiet")
	// var input string
	// fmt.Scanln(&input)
	// fmt.Println("done")

	//chanel
	// messageChannel := make(chan string)
	// channel(messageChannel)
	//	chanel bufer
	// messageChannelBuffer := make(chan string, 1)
	// channelBuffer(messageChannelBuffer)
	// msg := <-messageChannelBuffer
	// fmt.Println("reader:", msg)
	//	time.Sleep(time.Second * 3)
	// msg = <-messageChannelBuffer
	// fmt.Println("reader:", msg)
	// msg = <-messageChannelBuffer
	// fmt.Println("reader:", msg)
	//channel syncronization
	// ch := make(chan bool)
	// go channelSynchronization(ch)
	// <-ch

	//channel direction
	// s := make(chan string, 1) // it seems no sense with s:= make (chan<- string) or make(<-chan string ) without function
	// r := make(chan string, 1)
	// chnnelRceiver(s, "msg for receiver")
	// channelDirection(s, r)
	// fmt.Println(<-r)
	//fmt.Println("out of function:", <-s)

	//select for channel
	// selectF()
	// timeout()
	// closeChannel()
	// rangeChannel()
	// ticker()
	// timer()

	//worker:channel goroutines
	// creator := make(chan int, 100)
	// customer := make(chan int, 100)
	// flag := make(chan bool)
	//implement workers
	// for i := 0; i < 3; i++ {
	// 	go factory(i, creator, customer) // start goroutines first ,because channel creator can be block when nothing comes,and the goroutines will keep running
	// }
	// for i := 0; i < 20; i++ {
	// 	creator <- i
	// }
	// close(creator)
	// for i := 0; i < 20; i++ {
	// 	fmt.Println(<-customer)
	// }
	//why? when using range error occurs, learn range for channel
	// for item := range customer {
	// 	fmt.Println("deal result: ", item)
	// }
	// close(customer)

	// rateLimit()

	// atomicF()

	// mutexF()

	// stateful()

	// tmpstr := []string{"fjwfwef", "yf", "jjjel", "wohlhfe"}
	// //sort by element's length
	// sort.Sort(Str(tmpstr))
	// //sort by the first character'ascii of each element
	// // sort.Strings(tmpstr)
	// fmt.Println(tmpstr)
	// // fmt.Println("diy string len ", Str(tmpstr).Len())

	// panicF()

	// deferF()

	//collection functions
	// var strs = []string{"fjwfwef", "yf", "fdddddd", "jjjel", "wohlhfe", "eee", "b"}
	// fmt.Println("index b ", Index(strs, "b"))
	// fmt.Println("include b? ", Include(strs, "b"))
	// //actuall v represent single string in strs
	// fmt.Println(findByAnyPrefixChacter(strs, func(v string) bool {
	// 	return strings.HasPrefix(v, "y")
	// }))
	// fmt.Println(Delete(strs, "eee"))
	// fmt.Println(MapIndex(strs, func(v string) bool {
	// 	return strings.HasPrefix(v, "f")
	// }))

}
