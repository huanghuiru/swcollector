package cron

import (
	"github.com/huanghuiru/swcollector/funcs"
	"github.com/gaochao1/swcollector/g"
	"github.com/open-falcon/common/model"
	"log"
	"math"
	"time"
)

func Collect() {
	if !g.Config().Transfer.Enabled {
		return
	}

	if g.Config().Transfer.Addr == "" {
		return
	}

	for _, v := range funcs.Mappers {
		go collect(int64(v.Interval), v.Fs)
	}
}

func collect(sec int64, fns []func() []*model.MetricValue) {
	//for {
	go MetricToTransfer(sec, fns)
		//time.Sleep(time.Duration(sec) * time.Second)
	//}
}

func MetricToTransfer(sec int64, fns []func() []*model.MetricValue) {
	mvs := []*model.MetricValue{}
	log.Println("*fns",fns)
	
	for _, fn := range fns {
		items := fn()
		log.Println("firstitems",items)
		if items == nil {
			log.Println("items is nil")
			continue
		}

		if len(items) == 0 {
			log.Println("items is 0")
			continue
		}

		log.Println("test for start,fn:",fn)
		for _, mv := range items {
			mvs = append(mvs, mv)
		}
		log.Println("test for end")
	}

	startTime := time.Now()

	//分批次传给transfer
	n := 100
	lenMvs := len(mvs)
	log.Println("mvslen",lenMvs)			   

	div := lenMvs / n
	mod := math.Mod(float64(lenMvs), float64(n))

	mvsSend := []*model.MetricValue{}
	for i := 1; i <= div+1; i++ {

		if i < div+1 {
			mvsSend = mvs[n*(i-1) : n*i]
		} else {
			mvsSend = mvs[n*(i-1) : (n*(i-1))+int(mod)]
			log.Println("*mvsSend:",mvsSend)
		}
		time.Sleep(100 * time.Millisecond)

		go g.SendToTransfer(mvsSend)
	}

	endTime := time.Now()
	log.Println("INFO : Send metrics to transfer running in the background. Process time :", endTime.Sub(startTime), "Send metrics :", len(mvs))
}
