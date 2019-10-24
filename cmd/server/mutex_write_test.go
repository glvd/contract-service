package main

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"

	"github.com/glvd/seed/model"
)

// WriteOne ...
type WriteOne struct {
	Name  string
	Value string
}

// WriteTwo ...
type WriteTwo struct {
	Name  string
	Value string
}

// WriteThree ...
type WriteThree struct {
	Name  string
	Value string
}

// WriteMax ...
const WriteMax = 1000

// TestMutexWrite ...
func TestMutexWrite(t *testing.T) {
	eng, e := model.InitSQLite3("mutex1.db")
	if e != nil {
		return
	}
	e = eng.Sync2(WriteThree{}, WriteTwo{}, WriteOne{})
	if e != nil {
		return
	}
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		for i := 0; i < WriteMax; i++ {
			_, e := eng.InsertOne(&WriteOne{
				Name:  fmt.Sprintf("Name(%d)", i),
				Value: GenerateRandomString(64),
			})
			if e != nil {
				log.Error(e)
				return
			}
		}
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		for i := 0; i < WriteMax; i++ {
			_, e := eng.InsertOne(&WriteTwo{
				Name:  fmt.Sprintf("Name(%d)", i),
				Value: GenerateRandomString(64),
			})
			if e != nil {
				log.Error(e)
				return
			}
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < WriteMax; i++ {
			_, e := eng.InsertOne(&WriteThree{
				Name:  fmt.Sprintf("Name(%d)", i),
				Value: GenerateRandomString(64),
			})
			if e != nil {
				log.Error(e)
				return
			}
		}
		wg.Done()
	}()

	wg.Wait()
}

/*RandomKind RandomKind */
type RandomKind int

/*random kinds */
const (
	RandomNum      RandomKind = iota // 纯数字
	RandomLower                      // 小写字母
	RandomUpper                      // 大写字母
	RandomLowerNum                   // 数字、小写字母
	RandomUpperNum                   // 数字、大写字母
	RandomAll                        // 数字、大小写字母
)

/*RandomString defines */
var (
	RandomString = map[RandomKind]string{
		RandomNum:      "0123456789",
		RandomLower:    "abcdefghijklmnopqrstuvwxyz",
		RandomUpper:    "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		RandomLowerNum: "0123456789abcdefghijklmnopqrstuvwxyz",
		RandomUpperNum: "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		RandomAll:      "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ",
	}
)

//GenerateRandomString 随机字符串
func GenerateRandomString(size int, kind ...RandomKind) string {
	bytes := RandomString[RandomAll]
	if kind != nil {
		if k, b := RandomString[kind[0]]; b == true {
			bytes = k
		}
	}
	var result []byte
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		result = append(result, bytes[rand.Intn(len(bytes))])
	}
	return string(result)
}
