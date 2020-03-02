package main

import (
	"errors"
	"fmt"
	"math/rand"
	"sort"
	"time"

	gerr "github.com/pkg/errors"
)

var pool = []*PoolItem{
	{
		PoolSeq: 1,
		AwardItem: AwardItem{
			Pid:        100,
			Name:       "一百元话费",
			Img:        "pic",
			Num:        3,
			ExpireVal:  2,
			ExpireType: 73000,
		},
		GrantLimit: 1,
		GrantSum:   0,
		AwardRate:  30000000,
		AwardPrice: 100,
	},
	{
		PoolSeq: 2,
		AwardItem: AwardItem{
			Pid:        200,
			Name:       "er百元话费",
			Img:        "pic",
			Num:        3,
			ExpireVal:  2,
			ExpireType: 73000,
		},
		GrantLimit: 1,
		GrantSum:   0,
		AwardRate:  20000000,
		AwardPrice: 100,
	},
	{
		PoolSeq: 3,
		AwardItem: AwardItem{
			Pid:        300,
			Name:       "三百元话费",
			Img:        "pic",
			Num:        3,
			ExpireVal:  2,
			ExpireType: 73000,
		},
		GrantLimit: 1,
		GrantSum:   0,
		AwardRate:  10000000,
		AwardPrice: 100,
	},
	{
		PoolSeq: 4,
		AwardItem: AwardItem{
			Pid:        400,
			Name:       "四百元话费",
			Img:        "pic",
			Num:        3,
			ExpireVal:  2,
			ExpireType: 73000,
		},
		GrantLimit: 1,
		GrantSum:   0,
		AwardRate:  10000000,
		AwardPrice: 100,
	},
	{
		PoolSeq: 5,
		AwardItem: AwardItem{
			Pid:        500,
			Name:       "五百元话费",
			Img:        "pic",
			Num:        3,
			ExpireVal:  2,
			ExpireType: 73000,
		},
		GrantLimit: 1,
		GrantSum:   0,
		AwardRate:  10000000,
		AwardPrice: 100,
	},
	{
		PoolSeq: 6,
		AwardItem: AwardItem{
			Pid:        600,
			Name:       "六百元话费",
			Img:        "pic",
			Num:        3,
			ExpireVal:  2,
			ExpireType: 73000,
		},
		GrantLimit: 1,
		GrantSum:   0,
		AwardRate:  20000000,
		AwardPrice: 100,
	},
}

func main() {
	award, seq, err := GetRandomAward(pool)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(seq, award)
	incrAwardGrantSum(seq, pool)
}

func incrAwardGrantSum(seq int32, pool []*PoolItem) {
	for i := 0; i < len(pool); i++ {
		if seq == pool[i].PoolSeq {
			pool[i].GrantSum++
			break
		}
	}
}

var ErrStockNotEnough = errors.New("Stock Not Enough")
var ErrArgumentsInvalid = errors.New("Arguments Invalid")

const awardRateRadix int32 = 100000000

// TaskAward 任务奖励
type AwardItem struct {
	Pid        int64  `bson:"pid"`         // 奖励id
	Name       string `bson:"name"`        // 奖励名称
	Img        string `bson:"img"`         // 奖励图片
	Num        int32  `bson:"num"`         // 奖励数据
	ExpireVal  int64  `bson:"expire_val"`  // 过期值
	ExpireType int32  `bson:"expire_type"` // 过期类型
}

type PoolItem struct {
	PoolSeq    int32     //奖励序号
	AwardItem  AwardItem //奖励项
	GrantLimit int32     //发放限制
	GrantSum   int32     //已发放数量
	AwardRate  int32     //获奖率
	AwardPrice int32     //奖项价值
}

type Range struct {
	liftInterval  int32
	rightInterval int32
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func GetRandomAward(aPool []*PoolItem) (award AwardItem, poolSeq int32, err error) {
	if len(aPool) == 0 {
		return AwardItem{}, 0, ErrArgumentsInvalid
	}

	var pool []*PoolItem
	for i := 0; i < len(aPool); i++ {
		pool = append(pool, &PoolItem{
			PoolSeq:    aPool[i].PoolSeq,
			AwardItem:  aPool[i].AwardItem,
			GrantLimit: aPool[i].GrantLimit,
			GrantSum:   aPool[i].GrantSum,
			AwardRate:  aPool[i].AwardRate,
		})
	}
	poolSeq = -1
	bucket := initBucket(pool)
Label:
	randNum := rand.Int31n(awardRateRadix)
	for key := range bucket {
		if randNum >= key.liftInterval && randNum < key.rightInterval {
			element := bucket[key]

			if element.GrantLimit != 0 && element.GrantSum >= element.GrantLimit {
				fmt.Println("到达发放上限 ==》》", element.PoolSeq)
				bucket = initBucket(kickElement(pool))
				goto Label
			} else {
				award = element.AwardItem
				poolSeq = element.PoolSeq
				break
			}
		}
	}
	if poolSeq == -1 {
		err = ErrStockNotEnough
	}
	return award, poolSeq, gerr.WithStack(err)
}

func kickElement(pool []*PoolItem) []*PoolItem {
	for i := 0; i < len(pool); i++ {
		if pool[i].GrantSum >= pool[i].GrantLimit {
			for j := len(pool) - 1; j >= 0; j-- {
				if pool[j].GrantSum < pool[j].GrantLimit {
					pool[j].AwardRate += pool[i].AwardRate
					break
				}
			}
		}
	}
	temp := []*PoolItem{}
	for k := 0; k < len(pool); k++ {
		if pool[k].GrantSum < pool[k].GrantLimit {
			temp = append(temp, pool[k])
		}
	}
	return temp
}

func initBucket(items []*PoolItem) map[Range]*PoolItem {
	sort.Sort(AwardPoolSort(items))

	bucket := make(map[Range]*PoolItem, len(items))
	var loopCount int32
	for i := 0; i < len(items); i++ {
		if i == 0 {
			sec := Range{
				liftInterval:  0,
				rightInterval: items[i].AwardRate,
			}
			bucket[sec] = items[i]
		} else {
			sec := Range{
				liftInterval:  loopCount,
				rightInterval: items[i].AwardRate + loopCount,
			}
			bucket[sec] = items[i]
		}
		loopCount += items[i].AwardRate
	}
	return bucket
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
type AwardPoolSort []*PoolItem

//Swap  交换
func (a AwardPoolSort) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

// Len 长度
func (a AwardPoolSort) Len() int {
	return len(a)
}

// Less 大小比较
func (a AwardPoolSort) Less(i, j int) bool {
	if a[i].AwardRate < a[j].AwardRate {
		return true
	} else {
		return false
	}
}
