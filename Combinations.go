package main

import "fmt"

func main()  {
	fmt.Println(combine(4,2))
}

//C(n,k)=C(n-1,k-1)+C(n-1,k)
func combine(n int, k int) [][]int {
	if n <= 0 || k<= 0 {
		return nil
	}
	var result [][]int
	if k == 1 {
		for i:=1;i<=n;i++{
			result = append(result, []int{i})
		}
		return result
	}
	if k == n {
		var all []int
		for i:=1;i<=n;i++{
			all = append(all, i)
		}
		return [][]int{all}
	}
	in := combine(n-1,k-1)//n在组合中的情况
	for _, i := range in {
		nIn := make([]int, len(i)+1)
		nIn[0] = n
		copy(nIn[1:], i)
		result = append(result, nIn)
	}
	out := combine(n-1,k)//n不在组合中的情况
	result = append(result, out...)
	return result
}