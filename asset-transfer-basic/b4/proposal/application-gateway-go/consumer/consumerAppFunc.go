package consumer

import (
	//"bytes"
	//"encoding/json"
	"fmt"
	"time"
	//"strconv"
	//"math"
	//"sort"
	//"sync"
	//"net/http"
	
	"github.com/hyperledger/fabric-gateway/pkg/client"
)

const (
	earthRadius = 6378137.0
	pricePerMater = 0.000001
	kmPerBattery = 0.05 // battery(%) * kmPerBattery = x km
	layout = "2006-01-02T15:04:05+09:00"
)

func Bid(contract *client.Contract, userName string, latitude float64, longitude float64, energyAmount int, batteryLife float64) (
	[]Energy, time.Time, error) {
	searchRange := (100 - float64(batteryLife)) * kmPerBattery * 1000 // 1000m->500mに変更
	fmt.Printf("searchRange:%g\n", searchRange)

	var successList []Energy
	var autcionStartMin time.Time
	var err error
	return successList, autcionStartMin, err
}

func BidResult(contract *client.Contract, userName string, auctionStartMin time.Time) {

}