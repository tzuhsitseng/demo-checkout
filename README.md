# demo-checkout

## Concepts
- 將題目所提到的三種收費模式組成一個 strategy，針對第二題提出的需求變更再新增另一個 strategy
  - 透過 env 來設定類似 feature flag 的效果去做新舊版本的切換
- 目前是將折扣以及折抵比例放進 env 進行動態調整
  - 若有後台的話將這些資料放進 DB 或是 Config System 是比較好的選擇

## Contents
- 套件依賴管理使用 `Go Module`
- HTTP framework 使用 `Iris`
  - 可以使用其他 HTTP framework 像是 [Gin](https://github.com/gin-gonic/gin) 或是 [echo](https://github.com/labstack/echo) 來取代 
- Environment variables
  - 可依所需情境使用像是 [viper](https://github.com/spf13/viper) 等第三方套件來管理
- Dependency Injection
  - 可依所需情境使用像是 [dig](https://github.com/uber-go/dig) 或是 [wire](https://github.com/google/wire) 來作元件依賴管理

## Demo
```bash
## Start server which hosts on 8080 port
./bin/server 

## Call API
curl http://localhost:8080/calculation?user_level=2&user_points=100&price=1000

## 回傳將會使用多少平台幣與平台點數
## -> {"coin":710,"points":100}
```