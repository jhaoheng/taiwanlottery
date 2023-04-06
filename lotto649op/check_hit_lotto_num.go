package lotto649op

/*
- 檢查歷史中, 已經中過獎的號碼, 中獎機率多高
*/

type CheckHitLottoResult struct {
	CheckHit Lotto649OPData
	MatchHit []Lotto649OPData
	HitTime  int
}

func (op *Lotto649OP) CheckHitLotto(hit_num_count int) (results []CheckHitLottoResult) {

	results = []CheckHitLottoResult{}

	for _, xdata := range op.Datas {
		result := CheckHitLottoResult{
			CheckHit: xdata,
			MatchHit: []Lotto649OPData{},
			HitTime:  0,
		}

		choice := map[string]bool{
			xdata.Num_1:      true,
			xdata.Num_2:      true,
			xdata.Num_3:      true,
			xdata.Num_4:      true,
			xdata.Num_5:      true,
			xdata.Num_6:      true,
			xdata.NumSpecial: true,
		}

		//
		hit_lottery := []Lotto649OPData{}
		for _, data := range op.Datas {
			if xdata.SerialID == data.SerialID {
				continue
			}
			if data.Date.Unix() < xdata.Date.Unix() {
				continue
			}

			hits := 0
			if _, ok := choice[data.Num_1]; ok {
				hits++
			}
			if _, ok := choice[data.Num_2]; ok {
				hits++
			}
			if _, ok := choice[data.Num_3]; ok {
				hits++
			}
			if _, ok := choice[data.Num_4]; ok {
				hits++
			}
			if _, ok := choice[data.Num_5]; ok {
				hits++
			}
			if _, ok := choice[data.Num_6]; ok {
				hits++
			}
			if _, ok := choice[data.NumSpecial]; ok {
				hits++
			}
			//
			if hits >= hit_num_count {
				hit_lottery = append(hit_lottery, data)
			}
		}
		if len(hit_lottery) >= 1 {
			result.MatchHit = hit_lottery
			result.HitTime = len(hit_lottery)
			results = append(results, result)
		}
	}
	return results
}
