package ray

import (
	"testing"

	"github.com/thekorn/raytracing-go/internal/vec3"
)

func TestCreateDefaultRay(t *testing.T) {
	r := Ray{}
	if !r.origin.Equals(vec3.Vec3{}) || !r.direction.Equals(vec3.Vec3{}) {
		t.Errorf("failed creating a default ray, got %s", r)
	}
}
