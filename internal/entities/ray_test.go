package entities

import (
	"testing"

	"github.com/thekorn/raytracing-go/internal/vec3"
)

func TestCreateDefaultRay(t *testing.T) {
	var r Ray
	var v vec3.Vec3
	if !r.Origin.Equals(v) || !r.Direction.Equals(v) {
		t.Errorf("failed creating a default ray, got %s", r)
	}
}
