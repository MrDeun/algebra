package algebra

import "math"

type Vector3 [3]float64

func Vector3Zero() Vector3 {
	return Vector3{0, 0, 0}
}

func (vec *Vector3) X() float64 {
	return vec[0]
}

func (vec *Vector3) Y() float64 {
	return vec[1]
}

func (vec *Vector3) Z() float64 {
	return vec[2]
}

func Vector3Up() Vector3 {
	return Vector3{0, 1, 0}
}

func NewVector3(x, y, z float64) Vector3 {
	return Vector3{x, y, z}
}

func (vec *Vector3) lengthSquared() float64 {
	return vec[0]*vec[0] + vec[1]*vec[1] + vec[2]*vec[2]
}

func (vec *Vector3) length() float64 {
	return math.Sqrt(vec.lengthSquared())
}

func (vec *Vector3) Normalize() {
	length := vec.length()
	vec[0] /= length
	vec[1] /= length
	vec[2] /= length

}

func Vector3Cross(vec1, vec2 Vector3) Vector3 {
	return NewVector3(
		vec1[1]*vec2[2]-vec1[2]*vec2[1],
		vec1[2]*vec2[0]-vec1[0]*vec2[2],
		vec1[0]*vec2[1]-vec1[1]*vec2[0],
	)
}

func Vector3Dot(vec1, vec2 Vector3) float64 {
	return vec1[0]*vec2[0] + vec1[1]*vec2[1] + vec1[2]*vec2[2]
}
