package SpaceCookies

import (
	. "github.com/vova616/GarageEngine/Engine"
	//"Engine/Components"
	//"github.com/jteeuwen/glfw"
	"github.com/vova616/GarageEngine/Engine/Input"
	//"log"
	c "github.com/vova616/chipmunk"
	. "github.com/vova616/chipmunk/vect"
	//"fmt"
	//"time"
)

type ShipController struct {
	BaseComponent
	Speed           float32
	RotationSpeed   float32
	Missle          *Missle
	MisslesPosition []Vector
}

func NewShipController() *ShipController {
	return &ShipController{NewComponent(), 200000, 100, nil, []Vector{{10, -28, 0},
		{10, 28, 0}}}
}

func (sp *ShipController) OnComponentBind(binded *GameObject) {
	sp.GameObject().AddComponent(NewPhysics2(false, c.NewCircle(Vect{0, 0}, Float(15))))
}

func (sp *ShipController) Start() {
	ph := sp.GameObject().Physics
	ph.Body.SetMass(50)
	ph.Shape.Group = 1
	//sp.Physics.Shape.Friction = 0.5
}

func (sp *ShipController) Shoot() {
	if sp.Missle != nil {
		s := sp.Transform().Rotation2D()
		a := sp.Transform().Rotation()
		//scale := sp.Transform().Scale()
		for _, pos := range sp.MisslesPosition {
			p := sp.Transform().WorldPosition()
			_ = a
			m := Identity()
			//m.Scale(scale.X, scale.Y, scale.Z)
			m.Translate(pos.X, pos.Y, pos.Z)
			m.Rotate(a.Z, 0, 0, 1)
			m.Translate(p.X, p.Y, p.Z)
			p = m.Translation()

			nfire := sp.Missle.GameObject().Clone()
			nfire.Transform().SetParent2(GameSceneGeneral.Layer3)
			nfire.Transform().SetWorldPosition(p)
			nfire.Physics.Body.IgnoreGravity = true
			nfire.Physics.Body.SetMass(1)

			nfire.Physics.Body.AddForce(-s.X*30000, -s.Y*30000)

			nfire.Physics.Shape.Group = 1
			nfire.Physics.Body.SetMoment(Inf)
			nfire.Transform().SetRotation(sp.Transform().Rotation())
		}
	}
}

func (sp *ShipController) Update() {
	r := sp.Transform().Rotation()
	r2 := sp.Transform().Rotation2D()
	ph := sp.GameObject().Physics
	rx, ry := r2.X, r2.Y
	rx, ry = rx*DeltaTime(), ry*DeltaTime()

	if Input.KeyDown('S') {
		ph.Body.AddForce(sp.Speed*rx, sp.Speed*ry)
	}

	if Input.KeyDown('W') {
		ph.Body.AddForce(-sp.Speed*rx, -sp.Speed*ry)
	}

	if Input.KeyDown('A') {
		sp.Transform().SetRotationf(0, 0, r.Z-sp.RotationSpeed*DeltaTime())
	}
	if Input.KeyDown('D') {
		sp.Transform().SetRotationf(0, 0, r.Z+sp.RotationSpeed*DeltaTime())
	}

	if Input.KeyPress('F') {
		sp.Shoot()
	}

	if Input.KeyPress('P') {

		EnablePhysics = !EnablePhysics
	}
}

func (sp *ShipController) LateUpdate() {
	GameSceneGeneral.SceneData.Camera.Transform().SetPosition(NewVector3(sp.Transform().Position().X-float32(Width/2), sp.Transform().Position().Y-float32(Height/2), 0))
}