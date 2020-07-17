package main

// Cake は
type Cake struct{ state string }

func baker(cooked chan<- *Cake) {
	for {
		cake := new(Cake)
		cake.state = "cooked"
		cooked <- cake //  bakerはこのcakeには決して再び触れない
	}
}

func icer(iced chan<- *Cake, cooked <-chan *Cake) {
	for cake := range cooked {
		cake.state = "iced"
		iced <- cake // icerはこのcakeには決して再び触れない
	}
}

func main() {

}
