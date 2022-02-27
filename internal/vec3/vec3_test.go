package vec3

import (
	"math"
	"testing"
)

func shouldPanic(t *testing.T, f func()) {
	defer func() { recover() }()
	f()
	t.Errorf("should have panicked")
}

func TestCreateNullVector(t *testing.T) {
	v := MakeVec3(0, 0, 0)
	if v.X != 0 || v.Y != 0 || v.Z != 0 {
		t.Errorf("failed creating a NullVector")
	}
}

func TestCreateNewVector(t *testing.T) {
	v := MakeVec3(1, 2, 3)
	if v.X != 1 || v.Y != 2 || v.Z != 3 {
		t.Errorf("failed creating a Vector, got %v", v)
	}
}

func TestCreateNewColor(t *testing.T) {
	v := MakeColor(0, 0.4, 0.6)
	if v.X != 0 || v.Y != 0.4 || v.Z != 0.6 {
		t.Errorf("failed creating a Color, got %s", v)
	}
}

func TestCreateNewPoint3(t *testing.T) {
	v := MakePoint3(1, 2, 3)
	if v.X != 1 || v.Y != 2 || v.Z != 3 {
		t.Errorf("failed creating a Point3, got %v", v)
	}
}

func TestCalculateScalarProduct(t *testing.T) {
	v := MakeVec3(1, 2, 3)

	s := v.ScalarProd(6)
	if s.X != 6 || s.Y != 12 || s.Z != 18 {
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
	if s.X != -6 || s.Y != -12 || s.Z != -18 {
		t.Errorf("failed calculating scalar product of a Vector, got %v", s)
	}
}

func TestAddTwoVectors(t *testing.T) {
	a := MakeVec3(1, 2, 3)
	b := MakeVec3(4, 5, 6)

	v := a.Add(b)
	if v.X != 5 || v.Y != 7 || v.Z != 9 {
		t.Errorf("failed adding two vectors, got %v", v)
	}
}

func TestSubstractTwoVectors(t *testing.T) {
	a := MakeVec3(1, 2, 3)
	b := MakeVec3(4, 8, 12)

	v := a.Sub(b)
	if v.X != -3 || v.Y != -6 || v.Z != -9 {
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
	var o Vec3
	if !v.Equals(o) {
		t.Errorf("failed getting cross product of two vectors, got %v", v)
	}
}

func TestCreateBlackColor(t *testing.T) {
	var black Color
	if black.X != 0 || black.Y != 0 || black.Z != 0 {
		t.Errorf("failed creating black")
	}
}

func TestCreateInvalidColor(t *testing.T) {
	createInvalidColor := func() {
		MakeColor(10, 3, 5)
	}
	shouldPanic(t, createInvalidColor)
}
