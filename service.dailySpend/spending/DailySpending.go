package spending

type dailySpending struct {
	TotalSpend, TotalTransactions *int
}

func (s *dailySpending) AddTo(spend int) {

	var spent = *s.TotalSpend + spend
	var trans = *s.TotalTransactions + 1

	s.TotalSpend = &spent
	s.TotalTransactions = &trans
}
