package module

/*
[目的]
- 帶入指定 datas, 在資料表中過濾掉 datas 的資料
*/

/*
	- 過濾 plan_a=> 取得 中獎號碼, 7 取 5 排列組合
	- 過濾{指定號碼} => datas := [][]string{0:{"1","2","3","4","5","6"}}
	- 過濾{最後一次的中獎號碼, 七個數字} => datas := M_Get_N([]int{},6), 取得排列組合
*/

// type INewFilterTool interface {
// 	DirectlyDel(datas [][]int)
// 	SearchLikeAndDel(datas [][]int)
// 	SearchLikeAndDelWithGoroutine(datas [][]int)
// 	GetFilteredAndDoFilter(datas [][]int) []model.Lotto649AllSets
// }

// type FilterTool struct {
// }

// func NewFilterTool() INewFilterTool {
// 	return &FilterTool{}
// }

// // 直接刪除
// func (filter *FilterTool) DirectlyDel(datas [][]int) {
// 	fmt.Println("=== 開始過濾, 直接刪除 ===")
// 	start := time.Now()
// 	//
// 	for index, data := range datas {
// 		fmt.Printf("=> %v, 號碼: %02d\n", index, data)
// 		if len(data) != 6 {
// 			panic("資料錯誤, 必須 6 個號碼")
// 		}
// 		//
// 		if err := model.NewLotto649Filtered().SetNums(data).Delete(); err != nil {
// 			panic(err)
// 		}
// 	}
// 	fmt.Printf("執行時間: %v\n", -time.Until(start))
// }

// // -
// func (filter *FilterTool) SearchLikeAndDel(datas [][]int) {
// 	fmt.Println("=== 開始過濾 ===")
// 	start := time.Now()
// 	//
// 	del_count := 0
// 	for index, data := range datas {
// 		fmt.Printf("=> %v, 查詢號碼: %02d, ", index, data)
// 		if len(data) > 6 {
// 			panic("資料錯誤, 不得超過 6 個號碼")
// 		}
// 		//
// 		text := ""
// 		for _, n := range data {
// 			text = text + "%" + fmt.Sprintf("%02d", n) + "%"
// 		}
// 		finds, err := model.NewLotto649Filtered().FindNumsLike([]string{text})
// 		if err != nil {
// 			panic(err)
// 		}
// 		fmt.Printf("找到 %v\n", len(finds))
// 		if len(finds) != 0 {
// 			if err := model.NewLotto649Filtered().BatchDelete(finds); err != nil {
// 				panic(err)
// 			}
// 			del_count = del_count + len(finds)
// 		}
// 	}
// 	fmt.Printf("總共刪除: %v, 執行時間: %v\n", del_count, -time.Until(start))
// }

// // -
// func (filter *FilterTool) SearchLikeAndDelWithGoroutine(datas [][]int) {
// 	fmt.Println("=== 開始過濾 ===")
// 	start := time.Now()
// 	del_count := 0

// 	//
// 	done_chan := make(chan struct{}, len(datas))
// 	var waitgroup sync.WaitGroup
// 	doSomehting := func(text string) {
// 		finds, err := model.NewLotto649Filtered().FindNumsLike([]string{text})
// 		if err != nil {
// 			fmt.Printf("text: %v,err: %v\n", text, err.Error())
// 			return
// 		}
// 		fmt.Printf("=> %v, %v\n", text, len(finds))
// 		if len(finds) != 0 {
// 			err := model.NewLotto649Filtered().BatchDelete(finds)
// 			if err != nil {
// 				fmt.Printf("finds: %v,err: %v\n", finds, err.Error())
// 				return
// 			}
// 		}
// 		waitgroup.Done()
// 		done_chan <- struct{}{}
// 	}

// 	//
// 	for index, data := range datas {
// 		fmt.Printf("=> %v, 查詢號碼: %02d\n", index, data)
// 		if len(data) > 6 {
// 			panic("資料錯誤, 不得超過 6 個號碼")
// 		}

// 		//
// 		text := ""
// 		for _, n := range data {
// 			text = text + "%" + fmt.Sprintf("%02d", n) + "%"
// 		}
// 		waitgroup.Add(1)
// 		go doSomehting(text)
// 	}
// 	waitgroup.Wait()

// 	//
// 	close(done_chan)
// 	for range done_chan {
// 		del_count++
// 	}
// 	fmt.Printf("總共刪除: %v, 執行時間: %v\n", del_count, -time.Until(start))
// }

// func (filter *FilterTool) GetFilteredAndDoFilter(datas [][]int) []model.Lotto649AllSets {
// 	start := time.Now()
// 	//
// 	finds, err := model.NewLotto649AllSets().FindAll()
// 	if err != nil {
// 		panic(err)
// 	}
// 	tmp := []model.Lotto649AllSets{}

// 	for _, data := range datas {
// 		fmt.Printf("=== 開始: %v ===\n", data)
// 		sort.Ints(data)
// 		num_1 := fmt.Sprintf("%02d", data[0])
// 		num_2 := fmt.Sprintf("%02d", data[1])
// 		num_3 := fmt.Sprintf("%02d", data[2])
// 		num_4 := fmt.Sprintf("%02d", data[3])
// 		num_5 := fmt.Sprintf("%02d", data[4])
// 		num_6 := fmt.Sprintf("%02d", data[5])
// 		//
// 		text := fmt.Sprintf("%v,%v,%v,%v,%v,%v", num_1, num_2, num_3, num_4, num_5, num_6)
// 		//
// 		for _, find := range finds {
// 			if find.Nums == text {
// 				tmp = append(tmp, find)
// 			}
// 		}
// 		fmt.Println("=== 結束 ===")
// 	}
// 	//
// 	fmt.Printf("總共: %v, 執行時間: %v\n", len(tmp), -time.Until(start))
// 	return tmp
// }
