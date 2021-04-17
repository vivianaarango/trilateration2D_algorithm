package main

import "fmt"

func main()  {
	kenobi := Point{
		X: 4.0,
		Y: 4.0,
	}

	skywalker := Point{
		X: 9.0,
		Y: 7.0,
	}

	sato := Point{
		X: 9.0,
		Y: 1.0,
	}

	trilateration := NewSatellites(kenobi, skywalker, sato)
	result := trilateration.GetLocation(4, 3, 3.25)
	fmt.Println(result)
}