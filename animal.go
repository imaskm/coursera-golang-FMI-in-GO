package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type animals struct {
	food,locomotion,noise string
}

func main() {


	cow := animals{"grass","walk","moo"}
	bird := animals{"worms","fly","peep"}
	snake := animals{"mice","slither","hsss"}

	var s string

	for{


		fmt.Println("Enter the name of the Animal(cow,snake,bird) and Info(move,eat,speak) you want or Enter X to exit: ")
		reader := bufio.NewReader(os.Stdin)

		s, _ = reader.ReadString('\n')

		s = strings.TrimSpace(s)

		if s=="X" {
			break
		}

		req := strings.Split(s," ")


		if req[0] == "cow" {

			if req[1] == "eat"{
				cow.Eat()
			}else if req[1] == "move"{
				cow.Move()
			}else if req[1] == "speak"{
				cow.Speak()
			}else{
				fmt.Println("Invalid Request")
			}


		} else if req[0] == "bird" {

			if req[1] == "eat"{
				bird.Eat()
			}else if req[1] == "move"{
				bird.Move()
			}else if req[1] == "speak"{
				bird.Speak()
			}else{
				fmt.Println("Invalid Request")
			}

		}else if req[0] == "snake" {

			if req[1] == "eat"{
				snake.Eat()
			}else if req[1] == "move"{
				snake.Move()
			}else if req[1] == "speak"{
				snake.Speak()
			}else{
				fmt.Println("Invalid Request")
			}

		}
	}
}

func (anim animals) Eat()  {

	fmt.Println( anim.food )
	
}

func (anim animals)  Move()  {
	
	fmt.Println(anim.locomotion)
	
}

func (anim animals)  Speak()  {

	fmt.Println(anim.noise)
}
