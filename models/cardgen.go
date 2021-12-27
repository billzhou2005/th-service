package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"time"

	"github.com/jaracil/ei"
)

type Player struct {
	Name       string  `json:"name"`
	Cards      [3]Card `gorm:"embedded" json:"cards"`
	Cardstype  string  `json:"cardstype"`
	CIfirst    int     `json:"cifirst"`
	CIsecond   int     `json:"cisecond"`
	CIthird    int     `json:"cithird"`
	Cardsscore int     `json:"cardsscore"`
}

type Card struct {
	Points int `json:"points"`
	Suits  int `json:"suits"`
}

func Cardgen(numofplayers int) [9]Player {
	var players [9]Player

	cardtypecount := make(map[string]int)

	for i := 0; i < numofplayers; i++ {
		players[i].Name = "player" + strconv.Itoa(i+1)
	}
	//t1 := time.Now().UnixNano() / 1e6 //1564552562
	//fmt.Println(t1)
	//for i := 0; i < 3; i++ {
	//黑桃 1（spade 0-12）、红桃 2（heart 13-25）、梅花 3（club 26-38）、方块 4（dianmond 39-51）
	nums := generateRandomNumber(0, 51, 27)
	//nums[0] = 45
	//nums[1] = 9
	//nums[2] = 35
	//fmt.Println(nums)
	for i := 0; i < numofplayers; i++ {
		players[i].Cards[0].Points = nums[3*i] % 13
		players[i].Cards[0].Suits = int(nums[3*i]/13) + 1
		players[i].Cards[1].Points = nums[3*i+1] % 13
		players[i].Cards[1].Suits = int(nums[3*i+1]/13) + 1
		players[i].Cards[2].Points = nums[3*i+2] % 13
		players[i].Cards[2].Suits = int(nums[3*i+2]/13) + 1

		players[i] = cardsTypeAndCI(players[i])
		cardtypecount[players[i].Cardstype] += 1
	}
	fmt.Println(cardtypecount, time.Now().String())
	//fmt.Printf("%+v\n", players)
	//t2 := time.Now().UnixNano() / 1e6 //1564552562
	//fmt.Println(t2 - t1)
	jsonfile, err := os.Open("./models/jhlevel.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonfile.Close()
	bytevalue, _ := ioutil.ReadAll(jsonfile)

	var jhlevel interface{}
	json.Unmarshal(bytevalue, &jhlevel)
	for j := 0; j < numofplayers-1; j++ {
		for k := j + 1; k < numofplayers; k++ {
			//fmt.Println(j, k, ei.N(jhlevel).M(players[j].Cardstype).IntZ(), ei.N(jhlevel).M(players[k].Cardstype).IntZ())
			if ei.N(jhlevel).M(players[j].Cardstype).IntZ() > ei.N(jhlevel).M(players[k].Cardstype).IntZ() {
				players[j].Cardsscore += 1
			} else if ei.N(jhlevel).M(players[j].Cardstype).IntZ() < ei.N(jhlevel).M(players[k].Cardstype).IntZ() {
				players[k].Cardsscore += 1
			} else {
				switch players[j].Cardstype {
				case "highcard":
					if players[j].CIfirst > players[k].CIfirst {
						players[j].Cardsscore += 1
					} else if players[j].CIfirst < players[k].CIfirst {
						players[k].Cardsscore += 1
					} else {
						if players[j].CIsecond > players[k].CIsecond {
							players[j].Cardsscore += 1
						} else if players[j].CIsecond < players[k].CIsecond {
							players[k].Cardsscore += 1
						} else {
							if players[j].CIthird > players[k].CIthird {
								players[j].Cardsscore += 1
							} else if players[j].CIthird < players[k].CIthird {
								players[k].Cardsscore += 1
							}
						}
					}
				case "pair":
					if players[j].CIfirst > players[k].CIfirst {
						players[j].Cardsscore += 1
					} else if players[j].CIfirst < players[k].CIfirst {
						players[k].Cardsscore += 1
					} else {
						if players[j].CIthird > players[k].CIthird {
							players[j].Cardsscore += 1
						} else if players[j].CIthird < players[k].CIthird {
							players[k].Cardsscore += 1
						}
					}
				case "straight":
					if players[j].CIfirst > players[k].CIfirst {
						players[j].Cardsscore += 1
					} else if players[j].CIfirst < players[k].CIfirst {
						players[k].Cardsscore += 1
					}
				case "flush":
					if players[j].CIfirst > players[k].CIfirst {
						players[j].Cardsscore += 1
					} else if players[j].CIfirst < players[k].CIfirst {
						players[k].Cardsscore += 1
					} else {
						if players[j].CIsecond > players[k].CIsecond {
							players[j].Cardsscore += 1
						} else if players[j].CIsecond < players[k].CIsecond {
							players[k].Cardsscore += 1
						} else {
							if players[j].CIthird > players[k].CIthird {
								players[j].Cardsscore += 1
							} else if players[j].CIthird < players[k].CIthird {
								players[k].Cardsscore += 1
							}
						}
					}
				case "straightflush":
					if players[j].CIfirst > players[k].CIfirst {
						players[j].Cardsscore += 1
					} else if players[j].CIfirst < players[k].CIfirst {
						players[k].Cardsscore += 1
					}
				case "bomb":
					if players[j].CIfirst > players[k].CIfirst {
						players[j].Cardsscore += 1
					} else if players[j].CIfirst < players[k].CIfirst {
						players[k].Cardsscore += 1
					}
				default:
					fmt.Println("players[j].Cardstype valuee error")
				}
			}
		}
	}

	return players
}

func cardsTypeAndCI(players Player) Player {
	cards := players.Cards
	cardspoints := make([]int, 0)
	cardstype := "highcard"

	cardspoints = append(cardspoints, cards[0].Points)
	cardspoints = append(cardspoints, cards[1].Points)
	cardspoints = append(cardspoints, cards[2].Points)

	if cards[0].Points == cards[1].Points {
		cardstype = "pair"
		if cards[1].Points == cards[2].Points {
			cardstype = "bomb"
		}
	} else if cards[0].Points == cards[2].Points {
		cardstype = "pair"
	} else if cards[1].Points == cards[2].Points {
		cardstype = "pair"
	}

	if cards[0].Suits == cards[1].Suits {
		if cards[1].Suits == cards[2].Suits {
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
				cardspoints[0] = 11
				cardspoints[1] = 12
				cardspoints[2] = 13
			}
		}
	}
	players.Cardstype = cardstype
	switch players.Cardstype {
	case "highcard":
		if cardspoints[0] == 0 {
			players.CIfirst = cardspoints[0] + 13
			players.CIsecond = cardspoints[2]
			players.CIthird = cardspoints[1]
		} else {
			players.CIfirst = cardspoints[2]
			players.CIsecond = cardspoints[1]
			players.CIthird = cardspoints[0]
		}
	case "pair":
		if cardspoints[0] == cardspoints[1] {
			players.CIfirst = cardspoints[0]
			if cardspoints[0] == 0 {
				players.CIfirst = cardspoints[0] + 13
			}
			players.CIthird = cardspoints[2]
		} else if cardspoints[1] == cardspoints[2] {
			players.CIfirst = cardspoints[1]
			players.CIthird = cardspoints[0]
			if cardspoints[0] == 0 {
				players.CIthird = cardspoints[0] + 13
			}
		}

	case "straight":
		players.CIfirst = cardspoints[2]
	case "flush":
		if cardspoints[0] == 0 {
			players.CIfirst = cardspoints[0] + 13
			players.CIsecond = cardspoints[2]
			players.CIthird = cardspoints[1]
		} else {
			players.CIfirst = cardspoints[2]
			players.CIsecond = cardspoints[1]
			players.CIthird = cardspoints[0]
		}
	case "straightflush":
		players.CIfirst = cardspoints[2]
	case "bomb":
		players.CIfirst = cardspoints[2]
	default:
		fmt.Println("players.Cardstype error")
	}

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
