package main

import "fmt"

/*
剑指 Offer 57 - II. 和为s的连续正数序列
输入一个正整数 target ，输出所有和为 target 的连续正整数序列（至少含有两个数）。

序列内的数字由小到大排列，不同序列按照首个数字从小到大排列。
 */

func findContinuousSequence(target int) [][]int {
	res:=[][]int{}
//	for i:=1;i<target/2;i++{
//		sum:=0
//		i1:=i
//		tmp:=[]int{}
//		for sum<target{
//			tmp=append(tmp,i1)
//			i1++
//			sum+=i1
//
//		}
//		if sum==target{
//			tmp1:=make([]int,len(tmp))
//			copy(tmp1,tmp)
//			res=append(res,tmp1)
//		}
//	}
//	return res
	i,j:=1,2

	for i<=j{
		s:=(j+i)*(j-i+1)/2
		if s>target{
			i++
		}else if s<target{
			j++
		}else{
			if j-i>0{
				tmp:=[]int{}
				for k:=i;k<=j;k++{
					tmp=append(tmp,k)
				}
				res=append(res,tmp)

			}
			i++
		}
	}
	return res


}
func main() {
	res:=findContinuousSequence(9)
	fmt.Println(res)
}
