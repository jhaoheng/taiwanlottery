# 目的
- 建立 台灣彩券 本地端 sqlite 資料庫

# Docs
## 台灣彩券舊資料（csv）下載
- https://data.gov.tw/datasets/search?p=1&size=10&s=pubdate.date_desc&rtt=27699
- 已下載請參考 taiwan_lotto_csvs
- 台灣彩券每一季會更新 csv 檔案
- 匯入 app 請參考, `./apps/imports_csvs`

## 爬蟲
- 若想抓取最新資料, 使用爬蟲
- 爬蟲 app 請參考, `./apps/crawler_lotto649`