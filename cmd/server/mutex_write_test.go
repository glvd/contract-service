package main

import (
	"log"
	"math/rand"
	"sync"
	"testing"
	"time"
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
	m := &sync.Mutex{}
	m.Lock()
	defer m.Unlock()
	for i := 0; i < WriteMax; i++ {
		go func(i int) {
			log.Println("sleep", i)
			time.Sleep(time.Duration(i) * time.Second)
		}(i)
	}
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
