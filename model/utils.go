package model

import "math"

// 计算Sigmoid函数
func sigmoid(x float64) float64 {
	return 1 / (1 + math.Exp(-x))
}
