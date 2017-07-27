package card

func GetPaixing(hand_cards []*Card) int {
	if len(hand_cards) != 5 {
		return DouniuType_Meiniu
	}

	//检查是否有特殊牌型
	huapai_num := 0
	total_score := 0
	same_card_num := 0
	for i := 0; i < 5; i++ {
		if hand_cards[i].CardNo > 10 {
			huapai_num ++
		}
		total_score += hand_cards[i].GetScore()
	}
	for i := 0; i < 4; i++ {
		for j := i + 1; j < 5; j++ {
			if hand_cards[i].SameCardNoAs(hand_cards[j]) {
				same_card_num++
			}
		}
	}
	if total_score <= 10 {
		return DouniuType_Wuxiao
	}
	if same_card_num == 6 {
		return DouniuType_Zhadan
	}
	if huapai_num == 5 {
		return DouniuType_Wuhua
	}

	//检查常规牌型
	left_score := 0
	for i := 0; i < 3; i++ {
		cardi_score := hand_cards[i].GetScore()
		for j := i + 1; j < 4; j++ {
			cardj_score := hand_cards[j].GetScore()
			for k := j + 1; k < 5; k++ {
				cardk_score := hand_cards[k].GetScore()
				three_cards_score := cardi_score + cardj_score + cardk_score
				if three_cards_score % 10 == 0 {
					left_score = total_score - three_cards_score
					return GetLeftScorePaixing(left_score)
				}
			}
		}
	}

	return DouniuType_Meiniu
}

func GetLeftScorePaixing(score int) int {
	if 0 == score {
		return DouniuType_Meiniu
	}

	left_score := score % 10
	switch left_score {
	case 0:
		return DouniuType_Niuniu
	case 1:
		return DouniuType_Niu1
	case 2:
		return DouniuType_Niu2
	case 3:
		return DouniuType_Niu3
	case 4:
		return DouniuType_Niu4
	case 5:
		return DouniuType_Niu5
	case 6:
		return DouniuType_Niu6
	case 7:
		return DouniuType_Niu7
	case 8:
		return DouniuType_Niu8
	case 9:
		return DouniuType_Niu9
	}

	return DouniuType_Meiniu
}

func GetPaixingMultiple(paixing int) int {
	switch paixing {
	case DouniuType_Meiniu:
		return 1
	case DouniuType_Niu7, DouniuType_Niu8:
		return 2
	case DouniuType_Niu9:
		return 3
	case DouniuType_Niuniu:
		return 4
	case DouniuType_Wuhua:
		return 5
	case DouniuType_Zhadan:
		return 6
	case DouniuType_Wuxiao:
		return 7
	}
	return 1
}

func GetCardsMaxid(hand_cards []*Card) int {
	if len(hand_cards) != 5 {
		return 0
	}

	maxid := hand_cards[0].CardId
	for i := 1; i < 5; i++ {
		if hand_cards[i].CardId > maxid {
			maxid = hand_cards[i].CardId
		}
	}

	return maxid
}