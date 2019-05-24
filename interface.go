package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)
type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {

	food,walk,voice string
}
func (anime Cow) Eat()  {

	fmt.Println(anime.food)
}

func (anime Cow) Speak()  {

	fmt.Println(anime.voice)
}
func (anime Cow) Move()  {

	fmt.Println(anime.walk)
}

type Bird struct {

	food,walk,voice string
}
func (anime Bird) Eat()  {

	fmt.Println(anime.food)
}

func (anime Bird) Speak()  {

	fmt.Println(anime.voice)
}
func (anime Bird) Move()  {

	fmt.Println(anime.walk)
}

type Snake struct {

	food,walk,voice string
}
func (anime Snake) Eat()  {

	fmt.Println(anime.food)
}

func (anime Snake) Speak()  {

	fmt.Println(anime.voice)
}
func (anime Snake) Move()  {

	fmt.Println(anime.walk)
}

func main() {

	data := make(map[string]Animal)

	var s string

	for {

		fmt.Println("Enter 1 of the two commands or Enter X to exit: ")

		reader := bufio.NewReader(os.Stdin)

		s, _ = reader.ReadString('\n')

		s = strings.TrimSpace(s)

		if s == "X" {
			break
		}

		req := strings.Split(s, " ")

		if req[0] == "newanimal"{

			var m Animal

			switch req[2] {
			case "cow":

				m = Cow{"Grass","Walk","moo"}

			case "snake":

				m = Snake{"mice","slither","hsss"}

			case "bird":

				m = Bird{"worm","fly","peep"}

			default:
				fmt.Println("Invalid Input of animal type")
				continue

			}

			data[req[1]] = m
			fmt.Println("Created it!")

		} else if req[0] == "query" {

			switch req[2] {

			case "eat":

				data[req[1]].Eat()

			case "move":

				data[req[1]].Move()

			case "speak":

				data[req[1]].Speak()

			default:
				fmt.Println("Invalid data for animal")

			}

		}

	}
}