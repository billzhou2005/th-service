package models

import (
	"math/rand"
	"strconv"
	"time"
)

type Player struct {
	Username string `json:"username"`
	Cards    Cards  `json:"cards"`
}

type Cards struct {
	Cardone   int `json:"cardone"`
	Cardtwo   int `json:"cardtwo"`
	Cardthree int `json:"cardthree"`
}

func Cardgen() [9]Player {

	var players [9]Player

	for i := 0; i < 9; i++ {
		players[i].Username = "player" + strconv.Itoa(i)
	}
	//t1 := time.Now().UnixNano() / 1e6 //1564552562
	//fmt.Println(t1)
	//for i := 0; i < 3; i++ {
	nums := generateRandomNumber(1, 52, 27)
	//fmt.Println(nums)
	for i := 0; i < 9; i++ {
		players[i].Cards.Cardone = nums[3*i]
		players[i].Cards.Cardtwo = nums[3*i+1]
		players[i].Cards.Cardthree = nums[3*i+2]
	}
	//fmt.Println(players)
	//}
	//fmt.Printf("%+v\n", players)
	//t2 := time.Now().UnixNano() / 1e6 //1564552562
	//fmt.Println(t2 - t1)

	return players
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
