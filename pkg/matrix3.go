package algebra

import "math"

type Matrix3 [3]Vector3

func cos(angle float64) float64 {
	return math.Cos(angle)
}

func sin(angle float64) float64 {
	return math.Sin(angle)
}

func Matrix3Identidy() Matrix3 {
	return Matrix3{
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
	}
}

func Matrix3Roll(angle float64) Matrix3 {
	return Matrix3{
		{1, 0, 0},
		{0, cos(angle), -sin(angle)},
		{0, sin(angle), cos(angle)},
	}
}

func Matrix3Pitch(angle float64) Matrix3 {
	return Matrix3{
		{cos(angle), 0, sin(angle)},
		{0, 1, 0},
		{-sin(angle), 0, cos(angle)},
	}
}

func Matrix3Yaw(angle float64) Matrix3 {
	return Matrix3{
		{cos(angle), -sin(angle), 0},
		{sin(angle), cos(angle), 0},
		{0, 0, 1},
	}
}
