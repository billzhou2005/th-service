package models

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"time"
)

type Player struct {
	Username  string `json:"username"`
	Cards     Cards  `json:"cards"`
	Cardstype string `json:"cardstype"`
	Tablerank int    `json:"tablerank"`
}

type Cards struct {
	Cardone   Card `json:"cardone"`
	Cardtwo   Card `json:"cardtwo"`
	Cardthree Card `json:"cardthree"`
}
type Card struct {
	Points int `json:"points"`
	Suits  int `json:"suits"`
}

//黑桃 1（spade 0-12）、红桃 2（heart 13-25）、梅花 3（club 26-38）、方块 4（dianmond 39-51）

func Cardgen() [9]Player {

	var players [9]Player

	for i := 0; i < 9; i++ {
		players[i].Username = "player" + strconv.Itoa(i+1)
	}
	//t1 := time.Now().UnixNano() / 1e6 //1564552562
	//fmt.Println(t1)
	//for i := 0; i < 3; i++ {
	nums := generateRandomNumber(0, 51, 27)
	//nums[0] = 44
	//nums[1] = 45
	//nums[2] = 46
	fmt.Println(nums)
	for i := 0; i < 9; i++ {
		players[i].Cards.Cardone.Points = nums[3*i] % 13
		players[i].Cards.Cardone.Suits = int(nums[3*i]/13) + 1
		players[i].Cards.Cardtwo.Points = nums[3*i+1] % 13
		players[i].Cards.Cardtwo.Suits = int(nums[3*i+1]/13) + 1
		players[i].Cards.Cardthree.Points = nums[3*i+2] % 13
		players[i].Cards.Cardthree.Suits = int(nums[3*i+2]/13) + 1

		players[i].Cardstype = cardsType(players[i].Cards)
	}
	//fmt.Println(players)
	//}
	//fmt.Printf("%+v\n", players)
	//t2 := time.Now().UnixNano() / 1e6 //1564552562
	//fmt.Println(t2 - t1)

	return players
}

func cardsType(cards Cards) string {
	cardspoints := make([]int, 0)
	cardstype := "highcard"

	cardspoints = append(cardspoints, cards.Cardone.Points)
	cardspoints = append(cardspoints, cards.Cardtwo.Points)
	cardspoints = append(cardspoints, cards.Cardthree.Points)

	if cards.Cardone.Points == cards.Cardtwo.Points {
		cardstype = "pair"
		if cards.Cardtwo.Points == cards.Cardthree.Points {
			cardstype = "bomb"
		}
	} else if cards.Cardone.Points == cards.Cardthree.Points {
		cardstype = "pair"
	}

	if cards.Cardone.Suits == cards.Cardtwo.Suits {
		if cards.Cardtwo.Suits == cards.Cardthree.Suits {
			cardstype = "flush"
		}
	}

	sort.Ints(cardspoints)
	if cardspoints[0]+1 == cardspoints[1] {
		if cardspoints[1]+1 == cardspoints[2] {
			if cardstype == "flush" {
				cardstype = "straightflush"
			} else {
				cardstype = "straight"
			}
		}
	}
	if cardspoints[1] == 11 {
		if cardspoints[2] == 12 {
			if cardspoints[0] == 0 {
				if cardstype == "flush" {
					cardstype = "straightflush"
				} else {
					cardstype = "straight"
				}
			}
		}
	}
	return cardstype
}

func generateRandomNumber(start int, end int, count int) []int {
	if end < start || (end-start) < count {
		return nil
	}

	//slice saved as result
	nums := make([]int, 0)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for len(nums) < count {
		num := r.Intn((end - start)) + start

		exist := false
		for _, v := range nums {
			if v == num {
				exist = true
				break
			}
		}

		if !exist {
			nums = append(nums, num)
		}
	}

	return nums
}
