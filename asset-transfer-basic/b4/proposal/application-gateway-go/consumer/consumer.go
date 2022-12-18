package consumer

import (
	"fmt"
	//"io/ioutil"
	//"log"
	//"path"
	"time"
	// "math"
)

/* 
複数の需要家
type1: 普通充電 昼に充電
type2: 普通充電 夜に充電
type3: 急速充電 昼
*/
/* ユーザごとに充電開始時間と希望充電時間を私が決める（ランダムではない）*/
var startTime time.Time
var auctionInterval time.Duration

// ゴールーチンで各ユーザ起動
// input: シミュレーション開始時間
func AllConsumers(start time.Time, interval int) {
	startTime = start
	fmt.Println(startTime)
	auctionInterval = time.Duration(interval)
	consumer("consumer1", 1, 100000, 1000, 8, 1)

}
// 充電開始時間(差分)、バッテリー容量(Wh)、チャージ済み(Wh)、充電時間(hour)、最終的なバッテリー残量(0から1)
func consumer(username string, add time.Duration, battery float64, chargedEnergy float64, chargeTime float64, finalLife float64) {
	fmt.Println(startTime.Add(add * time.Hour))
	batteryLife := chargedEnergy / battery
	fmt.Println(batteryLife)
	amountPerMin := battery * finalLife / chargeTime / 60
	fmt.Println(amountPerMin)

	ticker := time.NewTicker(time.Minute * auctionInterval)
	zeroCount := 0
	// 1回目: amountPerMin * 2入札
	var getEnergy float64 = 0
	// getEnergy := bid(math.Ceil(amountPerMin * 2), lat, lon, username, batteryLife)
	if getEnergy == 0 {
		zeroCount++
		fmt.Printf("zeroCount: %d\n", zeroCount)
	} else {
		chargedEnergy += getEnergy
		zeroCount = 0
	}

	for {
		if chargedEnergy >= battery || zeroCount == 3 {
			ticker.Stop()
			fmt.Printf("break\n")
			break
		}
		// getEnergy := bid(math.Ceil(amountPerMin * 2), lat, lon, username, batteryLife)
		// ログ
		<-ticker.C
		getEnergy = 0
		if getEnergy == 0 {
			zeroCount++
		} else {
			zeroCount = 0
			chargedEnergy += getEnergy
		}
		fmt.Printf("zeroCount: %d\n", zeroCount)
	}

}
