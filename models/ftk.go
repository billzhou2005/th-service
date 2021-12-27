package models

import "gorm.io/gorm"

type FtkDb struct {
	gorm.Model
	RID     int    `json:"rID"`
	Player1 string `json:"player1"`
	Player2 string `json:"player2"`
	Player3 string `json:"player3"`
	Player4 string `json:"player4"`
}

type FtkCards struct {
	RID          int            `json:"rID"`
	PlayersCards [4]PlayerCards `json:"playersCards"`
}
type PlayerCards struct {
	Cards          [26]Card `gorm:"embedded" json:"cards"`
	CardsScore     int      `json:"cardsScore"`
	KindsGather    []Kinds  `json:"kindsGather"`
	BombFtkCounter int      `json:"bombFtkCounter"`
	BombFtkPure    []int    `json:"bombFtkPure"`
}
type Kinds struct {
	Points  int `json:"points"`
	Counter int `json:"counter"`
}

// create Table
func CreateFtkDb(db *gorm.DB, FtkDb *FtkDb) (err error) {
	err = db.Create(FtkDb).Error
	if err != nil {
		return err
	}
	return nil
}

//get tables
func GetFtkDbs(db *gorm.DB, FtkDb *[]FtkDb) (err error) {
	err = db.Find(FtkDb).Error
	if err != nil {
		return err
	}
	return nil
}

//get Table by id
func GetFtkDb(db *gorm.DB, FtkDb *FtkDb, id string) (err error) {
	err = db.Where("id = ?", id).First(FtkDb).Error
	if err != nil {
		return err
	}
	return nil
}

//update Table
func UpdateFtkDb(db *gorm.DB, FtkDb *FtkDb) (err error) {
	db.Save(FtkDb)
	return nil
}

//delete Table
func DeleteFtkDb(db *gorm.DB, FtkDb *FtkDb, id string) (err error) {
	db.Where("id = ?", id).Delete(FtkDb)
	return nil
}

/*
func testPrintJson(j interface{}) {
	b, _ := json.Marshal(j)
	log.Println(string(b))
}

func main() {

	for i := 0; i < 100; i++ {
		ftkCards := FtkGen(i)
		testPrintJson(ftkCards)
		// log.Println(ftkCards)
		for j := 0; j < 4; j++ {
			log.Println(ftkCards.PlayersCards[j].CardsScore)
		}
		time.Sleep(100 * time.Microsecond)
	}
}
*/
func FtkGen(rID int) FtkCards {
	var ftkCards FtkCards

	ftkCards.RID = rID
	//黑桃 1（spade 0-12）、红桃 2（heart 13-25）、梅花 3（club 26-38）、方块 4（dianmond 39-51）
	cards := getCards()
	for i := 0; i < 4; i++ {
		// 3 offest, points 3-15, 15 displayed as 2
		for j := 0; j < 26; j++ {
			ftkCards.PlayersCards[i].Cards[j].Points = cards[i*26+j]%13 + 3
			suits := int(cards[i*26+j] / 13)
			suits %= 4
			suits += 1
			ftkCards.PlayersCards[i].Cards[j].Suits = suits
		}
	}

	for i := 0; i < 4; i++ {
		ftkCards.PlayersCards[i].Cards = cardsBigToSmall(ftkCards.PlayersCards[i].Cards)
		ftkCards.PlayersCards[i] = cardsKindsGather(ftkCards.PlayersCards[i])
		ftkCards.PlayersCards[i] = cardsBombFtk(ftkCards.PlayersCards[i])
		ftkCards.PlayersCards[i] = cardsScore(ftkCards.PlayersCards[i])
	}

	return ftkCards
}

func cardsScore(playerCards PlayerCards) PlayerCards {
	// pair 20, 3kinds 50, 4kinds 300, 5kinds 400, 6kinds 500, 7kinds 700, 8kinds 1000
	// ftk 100, ftkpure 200
	score := 0

	for i := range playerCards.KindsGather {
		switch playerCards.KindsGather[i].Counter {
		case 8:
			score += 1000
		case 7:
			score += 700
		case 6:
			score += 500
		case 5:
			score += 400
		case 4:
			score += 300
		case 3:
			score += 50
		case 2:
			score += 20
		default:
			score += playerCards.KindsGather[i].Points * 10
		}
	}

	score += playerCards.BombFtkCounter * 100
	score += len(playerCards.BombFtkPure) * 100

	playerCards.CardsScore = score
	return playerCards
}

func cardsBombFtk(playerCards PlayerCards) PlayerCards {
	var cardsFive []Card
	var cardsTen []Card
	var cardsThirteen []Card

	for i := range playerCards.Cards {
		if playerCards.Cards[i].Points == 5 {
			cardsFive = append(cardsFive, playerCards.Cards[i])
		}
		if playerCards.Cards[i].Points == 10 {
			cardsTen = append(cardsTen, playerCards.Cards[i])
		}
		if playerCards.Cards[i].Points == 13 {
			cardsThirteen = append(cardsThirteen, playerCards.Cards[i])
		}
	}

	min := len(cardsFive)
	minFlag := 5
	if min > len(cardsTen) {
		min = len(cardsTen)
		minFlag = 10
	}
	if min > len(cardsThirteen) {
		min = len(cardsThirteen)
		minFlag = 13
	}
	playerCards.BombFtkCounter = min

	if min == 0 {
		return playerCards
	}

	var suits []int

	if minFlag == 5 {
		for i := range cardsFive {
			suits = append(suits, cardsFive[i].Suits)
		}
	} else if minFlag == 10 {
		for i := range cardsTen {
			suits = append(suits, cardsTen[i].Suits)
		}
	} else if minFlag == 13 {
		for i := range cardsThirteen {
			suits = append(suits, cardsThirteen[i].Suits)
		}
	}

	suits = RemoveRepByLoop(suits)
	for i := range suits {
		fivePureFlag := false
		tenPureFlag := false
		thirteenPureFlag := false
		for j := range cardsFive {
			if suits[i] == cardsFive[j].Suits {
				fivePureFlag = true
			}
		}
		for j := range cardsTen {
			if suits[i] == cardsTen[j].Suits {
				tenPureFlag = true
			}
		}
		for j := range cardsThirteen {
			if suits[i] == cardsThirteen[j].Suits {
				thirteenPureFlag = true
			}
		}

		if fivePureFlag && tenPureFlag && thirteenPureFlag {
			playerCards.BombFtkPure = append(playerCards.BombFtkPure, suits[i])
		}
	}
	return playerCards
}

func RemoveRepByLoop(slc []int) []int {
	result := []int{}
	for i := range slc {
		flag := true
		for j := range result {
			if slc[i] == result[j] {
				flag = false
				break
			}
		}
		if flag {
			result = append(result, slc[i])
		}
	}
	return result
}

func cardsKindsGather(playerCards PlayerCards) PlayerCards {
	var kinds Kinds

	nums := len(playerCards.Cards)
	kinds.Points = playerCards.Cards[0].Points
	kinds.Counter = 1

	for i := 1; i < nums; i++ {
		if playerCards.Cards[i].Points == kinds.Points {
			kinds.Counter++
		} else {
			playerCards.KindsGather = append(playerCards.KindsGather, kinds)
			kinds.Points = playerCards.Cards[i].Points
			kinds.Counter = 1
		}

		if i == nums-1 {
			playerCards.KindsGather = append(playerCards.KindsGather, kinds)
		}
	}
	return playerCards
}

func cardsBigToSmall(cards [26]Card) [26]Card {
	nums := len(cards)
	for i := 0; i < nums; i++ {
		for j := i + 1; j < nums; j++ {
			if cards[i].Points < cards[j].Points {
				temp := cards[i]
				cards[i] = cards[j]
				cards[j] = temp
			}
		}
	}
	return cards
}

func getCards() []int {
	var cards [104]int

	nums := generateRandomNumber(0, 104, 78)
	for i := 0; i < 104; i++ {
		cards[i] = i
	}

	for i := 0; i < 78; i++ {
		cards[nums[i]] = 255
	}
	for i := 0; i < 104; i++ {
		if cards[i] != 255 {
			nums = append(nums, cards[i])
		}
	}
	return nums
}
