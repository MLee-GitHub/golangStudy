package leetcode

import (
	"fmt"
	"math"
)

// 五次及以上多项式方程没有根式解(没有求根公式)被伽罗瓦用群论做出的最著名结论，but, 没有王屠夫难道非得吃带毛猪???hehe...
// 工作生活中还是有诸多求解高次方程的真实需求，那这日子可怎么过下去啊🥺
// 1. 牛顿法可近似求解任意方程的解f(x)=0, 此处为f(x)=x^2-N=0  (切线(离切点越近误差越小)是曲线的线性逼近，那么此时
// 曲线的根即可近似用切线的根表示[切点的横坐标])，是否总可以保证收敛
// 2. 对于求解N的平方根问题，可转化为在1～N内采用二分法进行查找解的问题

func NewtonSqrt(x float64, precision int) float64 {
	if precision < 0 {
		return 0
	}
	ans := x
	err := math.Pow(10, float64(-precision)) // 1e-2
	for ;math.Abs(x - ans * ans) > err;{
		ans = (ans + x/ans)/2
		fmt.Println(ans)
	}
	return ans
}
