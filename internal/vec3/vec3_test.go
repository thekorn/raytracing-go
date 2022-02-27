package vec3

import (
	"math"
	"testing"
)

func TestCreateNullVector(t *testing.T) {
	v := MakeVec3(0, 0, 0)
	if v.x != 0 || v.y != 0 || v.z != 0 {
		t.Errorf("failed creating a NullVector")
	}
}

func TestCreateNewVector(t *testing.T) {
	v := MakeVec3(1, 2, 3)
	if v.x != 1 || v.y != 2 || v.z != 3 {
		t.Errorf("failed creating a Vector, got %v", v)
	}
}

func TestCreateNewColor(t *testing.T) {
	v := Color{1, 2, 3}
	if v.x != 1 || v.y != 2 || v.z != 3 {
		t.Errorf("failed creating a Color, got %s", v)
	}
}

func TestCreateNewPoint3(t *testing.T) {
	v := Point3{1, 2, 3}
	if v.x != 1 || v.y != 2 || v.z != 3 {
		t.Errorf("failed creating a Point3, got %v", v)
	}
}

func TestCalculateScalarProduct(t *testing.T) {
	v := MakeVec3(1, 2, 3)

	s := v.ScalarProd(6)
	if s.x != 6 || s.y != 12 || s.z != 18 {
		t.Errorf("failed calculating scalar product of a Vector, got %v", s)
	}
}

func TestCalculateScalarProductStructEqual(t *testing.T) {
	v := MakeVec3(1, 2, 3)

	s := v.ScalarProd(6)
	o := Vec3{6, 12, 18}
	if s != o {
		t.Errorf("failed calculating scalar product of a Vector, got %v", s)
	}
}

func TestCalculateNegativeScalarProduct(t *testing.T) {
	v := MakeVec3(1, 2, 3)

	s := v.ScalarProd(-6)
	if s.x != -6 || s.y != -12 || s.z != -18 {
		t.Errorf("failed calculating scalar product of a Vector, got %v", s)
	}
}

func TestAddTwoVectors(t *testing.T) {
	a := MakeVec3(1, 2, 3)
	b := MakeVec3(4, 5, 6)

	v := a.Add(b)
	if v.x != 5 || v.y != 7 || v.z != 9 {
		t.Errorf("failed adding two vectors, got %v", v)
	}
}

func TestSubstractTwoVectors(t *testing.T) {
	a := MakeVec3(1, 2, 3)
	b := MakeVec3(4, 8, 12)

	v := a.Sub(b)
	if v.x != -3 || v.y != -6 || v.z != -9 {
		t.Errorf("failed substracting two vectors, got %s", v)
	}
}

func TestGetLengthOfVector(t *testing.T) {
	a := MakeVec3(0, 0, 0)
	if a.Length() != 0 {
		t.Errorf("failed getting length of NullVector, got %v", a.Length())
	}

	a = MakeVec3(3, 0, 0)
	if a.Length() != 3 {
		t.Errorf("failed getting length of Vector, got %v", a.Length())
	}

	a = MakeVec3(1, 2, 3)
	if a.Length() != math.Sqrt(14) {
		t.Errorf("failed getting length of Vector, got %v", a.Length())
	}
}

func TestGetUnitVector(t *testing.T) {
	v := MakeVec3(2, 0, 0)
	if !v.UnitVec().Equals(Vec3{1, 0, 0}) {
		t.Errorf("failed getting the UnitVector")
	}
}

func TestGetDotProductTwoVectors(t *testing.T) {
	a := MakeVec3(1, 2, 3)
	b := MakeVec3(4, 8, 12)

	v := a.Dot(b)
	if v != 56 {
		t.Errorf("failed getting dot product of two vectors, got %f", v)
	}
}

func TestGetCrossProductTwoVectors(t *testing.T) {
	a := MakeVec3(1, 2, 3)
	b := MakeVec3(4, 8, 12)

	v := a.Cross(b)
	if !v.Equals(Vec3{}) {
		t.Errorf("failed getting cross product of two vectors, got %v", v)
	}
}
