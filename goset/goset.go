package goset

import (
	"fmt"

	set "github.com/deckarep/golang-set"
)

func Run() {
	kide := set.NewSet()
	kide.Add("xiaorui.cc")
	kide.Add("blog.xiaorui.cc")
	kide.Add("vps.xiaorui.cc")
	kide.Add("linode.xiaorui.cc")

	special := []interface{}{"Biology", "Chemistry"}
	scienceClasses := set.NewSetFromSlice(special)

	address := set.NewSet()
	address.Add("beijing")
	address.Add("nanjing")
	address.Add("shanghai")

	bonusClasses := set.NewSet()
	bonusClasses.Add("Go Programming")
	bonusClasses.Add("Python Programming")

	//一个并集的运算
	allClasses := kide.Union(scienceClasses).Union(address).Union(bonusClasses)
	fmt.Println("并集：", allClasses)

	//是否包含"Cookiing"
	fmt.Println("是否包含：", scienceClasses.Contains("Cooking")) //false

	//两个集合的差集
	fmt.Println("差集：", allClasses.Difference(scienceClasses)) //Set{Music, Automotive, Go Programming, Python Programming, Cooking, English, Math, Welding}

	//两个集合的交集
	fmt.Println("交集：", scienceClasses.Intersect(kide)) //Set{Biology}

	//有多少基数
	fmt.Println(bonusClasses.Cardinality()) //2

}

func Run1() {
	newViewList := set.NewSet()
	newViewList.Add("dx_bj")
	newViewList.Add("yd_bj")
	newViewList.Add("lt_bj")
	newViewList.Add("yd_sz")
	newViewList.Add("lt_sz")

	oldViewList := set.NewSet()
	oldViewList.Add("yd_bj")
	oldViewList.Add("lt_bj")
	oldViewList.Add("dx_sz")
	oldViewList.Add("yd_sz")
	oldViewList.Add("lt_sz")

	diffList := oldViewList.Difference(newViewList)
	if diffList.Cardinality() > 0 {
		fmt.Printf("view changed:%v\n", diffList)
	}
}
