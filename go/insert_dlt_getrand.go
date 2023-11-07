package main

import (
	"fmt"
	"math/rand"
	"slices"
)

type RandomizedSet struct {
	values map[int]int
	vList  []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		values: map[int]int{},
		vList:  make([]int, 0, 500),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.values[val]; ok {
		return false
	}
	this.vList = append(this.vList, val)
	this.values[val] = len(this.vList) - 1
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.values[val]; !ok {
		return false
	}
	i := this.values[val]
	this.vList = slices.Delete[[]int](this.vList, i, 1)
	fmt.Println(this.vList)
	delete(this.values, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.vList[rand.Intn(len(this.vList))]
}

func RunInsertDltGetRand() {
	obj := Constructor()
	obj.Insert(0)
	obj.Insert(1)
	obj.Remove(0)
	obj.Insert(2)
	obj.Remove(1)
	fmt.Println(obj.GetRandom())
}
