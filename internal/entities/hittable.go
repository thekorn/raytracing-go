package entities

import (
	"github.com/thekorn/raytracing-go/internal/vec3"
)

type HitRecord struct {
	P          vec3.Point3
	Normal     vec3.Vec3
	T          float64
	Front_face bool
	Material   Material
}

func (h *HitRecord) Set_face_normal(r Ray, outward_normal vec3.Vec3) {
	h.Front_face = r.Direction.Dot(outward_normal) < 0
	if h.Front_face {
		h.Normal = outward_normal
	} else {
		h.Normal = outward_normal.ScalarProd(-1)
	}
}

func (h *HitRecord) Update(n HitRecord) {
	h.P = n.P
	h.Normal = n.Normal
	h.T = n.T
	h.Front_face = n.Front_face
	h.Material = n.Material
}

type Hittable interface {
	Hit(r Ray, t_min float64, t_max float64, rec *HitRecord) bool
}

type HittableList struct {
	List []Hittable
}

func (hl *HittableList) Clear() {
	hl.List = []Hittable{}
}

func (hl *HittableList) Add(object Hittable) {
	hl.List = append(hl.List, object)
}

func (hl HittableList) Hit(r Ray, t_min float64, t_max float64, rec *HitRecord) bool {
	var temp_rec HitRecord
	hit_anything := false
	closest_so_far := t_max

	for _, object := range hl.List {
		if object.Hit(r, t_min, closest_so_far, &temp_rec) {
			hit_anything = true
			closest_so_far = temp_rec.T
			rec.Update(temp_rec)
		}
	}
	return hit_anything
}

func MakeHittableList() HittableList {
	var h HittableList
	return h
}

func MakeHittableListWithObject(object Hittable) HittableList {
	h := MakeHittableList()
	h.Add(object)
	return h
}
