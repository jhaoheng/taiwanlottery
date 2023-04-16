## query timeout

- `SET @@local.net_read_timeout=1200;`

## flow

- [ ] : 如果是 num_idx 是否有自己的趨勢？
	- 因為這不是數字，而是此 index 被排列到的可能性
	- 取出幾個數字，換算成折線圖看一下好了


- [ ] : 問題是排列過後的 num_idex，在 plan_g 後的圖表，最小的數字，被選中的頻率如何？
	- 所以 plan_g 輸出的 sum 寫入 txt 記錄下來即可
	- 也要記錄中獎號碼捏
	- 中獎號碼要紀錄的話，就要查表，plan_f


```
	// /*
	//     5. 取得 plan_e 的數字（消耗時間: 53.774µs），進行過濾

	//    - 只取得最後一次中獎數字
	// */
	// all_sets_map = func() map[string]struct{} {
	// 	plan_e := plan.NewPlanE()
	// 	filter_combinations := plan_e.GetSpecificNums(flow.Hits, 1)
	// 	fmt.Println("準備過濾的數字: ", filter_combinations)
	// 	all_sets_map = plan_e.RunFilter(all_sets_map, filter_combinations)
	// 	fmt.Println("剩下 =>", len(all_sets_map))
	// 	return all_sets_map
	// }()
```

## 需要做的事情
   
- 經過 flow_a 的計算後，排除掉連續五個靠在一起的數字, 補滿六個進行過濾
- 計算歷史中獎號碼，每個號碼的相依性
  - 目的: 
    - 計算所有數字的相依性, ex: 1 最常跟哪些號碼碰撞
    - 用來延伸選號參考, ex: 選擇 1 後，下一個選擇哪個號碼
  - 所以一次性計算全部的資料
  - 欄位: 起始日期、結束日期、下期號碼、本期號碼、num(1~49)、1-49 的相依號碼
  - 每個 num 與 相依號碼，若在該期碰撞到，則內容計算為 1
  - 最後計算總和

