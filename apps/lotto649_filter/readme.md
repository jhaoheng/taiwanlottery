## query timeout

- `SET @@local.net_read_timeout=1200;`

## copy
- `INSERT INTO lotto649_filtered SELECT * FROM lotto649_all_sets;`
- 直接 copy table 比較快

## issue

- like 的作法還是不行
- 可能應該是，每一次算出來的總數，透過運算，透過比對，在寫入到 db