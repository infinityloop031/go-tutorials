package main


func main(){
	// var bookings [50]string
	// bookings[0]="bill"
	// bookings[1]="shabbir"
	// bookings[2]="peter"

	// println(bookings[0])
	// println(len(bookings))

	// slices

	var booking []string
	booking=append(booking,"nil")
	booking=append(booking,"bye")
	booking=append(booking, "shabbir")

	for i := 0; i < len(booking); i++ {
		println(booking[i])
	}


}