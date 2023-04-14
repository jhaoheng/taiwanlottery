## query timeout

- `SET @@local.net_read_timeout=1200;`

## flow

1. 取得 flow_a 的落點預測
2. 排除掉 flow_a 的數據, 放入 flow_b 過濾資料, 寫入 db


- 目前需要驗證
- 產生出報表後，計算 index 的總數



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
- 完成 flow_b, 進行驗證
  - 很重要，手動驗證三期，都在合理範圍
  - 驗證流程
    - 如果要預測 112000041
    1. 需要抓取 112000040 的樂透結果
    2. 計算出趨勢圖(plan_g), 112000040, num_index_sum
    3. 產生出 flow_a(Test_flowa_GetRankAndCSV) 的圖表, 112000040
    4. 判斷出不可能中獎的號碼 
    5. 取得 112000041 的中獎號碼, 進行比對
- 經過 flow_a 的計算後，排除掉連續五個靠在一起的數字, 補滿六個進行過濾
- 當日的天數與中獎號碼的關係，看了 13 組，似乎都沒碰到
- 計算歷史中獎號碼，每個號碼的相依性
  - 目的: 
    - 計算所有數字的相依性, ex: 1 最常跟哪些號碼碰撞
    - 用來延伸選號參考, ex: 選擇 1 後，下一個選擇哪個號碼
  - 所以一次性計算全部的資料
  - 欄位: 起始日期、結束日期、下期號碼、本期號碼、num(1~49)、1-49 的相依號碼
  - 每個 num 與 相依號碼，若在該期碰撞到，則內容計算為 1
  - 最後計算總和


## muli-flow-a 缺點
- 計算範圍: 103000001~112000039, 其中使用 plan_a (7 get 5) 過濾
- 下一次的中獎號碼為 112000040，會直接打點在 圖表 上
- 112000040 並沒有過濾掉 7 get 5，所以這個是小缺點


## 預測 112000042
