package entities

import (
	"testing"

	"github.com/thekorn/raytracing-go/internal/vec3"
)

func TestCreateDefaultRay(t *testing.T) {
	r := Ray{}
	if !r.Origin.Equals(vec3.Vec3{}) || !r.Direction.Equals(vec3.Vec3{}) {
		t.Errorf("failed creating a default ray, got %s", r)
	}
}
