package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// 58. 区间和
// https://kamacoder.com/problempage.php?pid=1070
// 第一行输入为整数数组 Array 的长度 n，接下来 n 行，每行一个整数，表示数组的元素。随后的输入为需要计算总和的区间下标：a，b （b > = a），直至文件结束。
func main() {
	var n int
	fmt.Scanf("%d", &n)
	// 获取数组元素的同时计算前缀和，一般建议切片开大一点防止各种越界问题
	arr := make([]int, n+1)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &arr[i])
		if i > 0 {
			arr[i] += arr[i-1]
		}
	}

	/*
	   区间[l, r]的和可以使用区间[0, r]和[0, l - 1]相减得到，
	   在代码中即为arr[r]-arr[l-1]。这里需要注意l-1是否越界
	*/
	for {
		var l, r int
		_, err := fmt.Scanf("%d %d", &l, &r)
		if err != nil {
			return
		}

		if l > 0 {
			fmt.Println(arr[r] - arr[l-1])
		} else {
			fmt.Println(arr[r])
		}
	}
}

func main2() {
	// bufio中读取数据的接口，因为数据卡的比较严，导致使用fmt.Scan会超时
	scanner := bufio.NewScanner(os.Stdin)

	// 获取数组大小
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	// 获取数组元素的同时计算前缀和，一般建议切片开大一点防止各种越界问题
	arr := make([]int, n+1)
	for i := 0; i < n; i++ {
		scanner.Scan()
		arr[i], _ = strconv.Atoi(scanner.Text())
		if i != 0 {
			arr[i] += arr[i-1]
		}
	}

	/*
	   区间[l, r]的和可以使用区间[0, r]和[0, l - 1]相减得到，
	   在代码中即为arr[r]-arr[l-1]。这里需要注意l-1是否越界
	*/
	for {
		var l, r int
		scanner.Scan()
		_, err := fmt.Sscanf(scanner.Text(), "%d %d", &l, &r)
		if err != nil {
			return
		}

		if l > 0 {
			fmt.Println(arr[r] - arr[l-1])
		} else {
			fmt.Println(arr[r])
		}
	}
}

// 44. 开发商购买土地
// 在一个城市区域内，被划分成了n * m个连续的区块，每个区块都拥有不同的权值，代表着其土地价值。目前，有两家开发公司，A 公司和 B 公司，希望购买这个城市区域的土地。
// 现在，需要将这个城市区域的所有区块分配给 A 公司和 B 公司。
// 然而，由于城市规划的限制，只允许将区域按横向或纵向划分成两个子区域，而且每个子区域都必须包含一个或多个区块。
// 为了确保公平竞争，你需要找到一种分配方式，使得 A 公司和 B 公司各自的子区域内的土地总价值之差最小。
// 注意：区块不可再分。
