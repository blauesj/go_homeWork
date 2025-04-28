package main

func main() {

}

type MyCalendar struct {
	bookList [][2]int
}

func Constructor() MyCalendar {
	return MyCalendar{}
}

func (this *MyCalendar) Book(startTime int, endTime int) bool {
	for _, booked := range this.bookList {
		if booked[0] < endTime && booked[1] > startTime {
			return false
		}
	}
	this.bookList = append(this.bookList, [2]int{startTime, endTime})
	return true
}
