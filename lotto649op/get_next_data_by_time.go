package lotto649op

import "time"

func (op *Lotto649OP) GetNextDataByTime(the_time time.Time) (the_data Lotto649OPData) {
	var min int64 = 1000000000.0
	the_data = Lotto649OPData{}

	for _, data := range op.Datas {
		if data.Date.Unix()-the_time.Unix() > 0 {
			if (data.Date.Unix() - the_time.Unix()) < min {
				min = data.Date.Unix() - the_time.Unix()
				the_data = data
			}
		}
	}
	return
}
