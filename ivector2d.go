// int position in 2d space
package ivector2d

import (
	"fmt"

	"github.com/kasworld/rand"
)

type Vt [2]int

func (v Vt) String() string {
	return fmt.Sprintf("[%v,%v]", v[0], v[1])
}

var Zero = Vt{0, 0}
var UnitX = Vt{1, 0}
var UnitY = Vt{0, 1}

// func (p Vt) Copy() Vt {
//  return Vt{p[0], p[1]}
// }
func (p Vt) Eq(other Vt) bool {
	return p == other
	//return p[0] == other[0] && p[1] == other[1]
}
func (p Vt) Ne(other Vt) bool {
	return !p.Eq(other)
}
func (p Vt) IsZero() bool {
	return p.Eq(Zero)
}
func (p Vt) Add(other Vt) Vt {
	return Vt{p[0] + other[0], p[1] + other[1]}
}
func (p Vt) Neg() Vt {
	return Vt{-p[0], -p[1]}
}
func (p Vt) Sub(other Vt) Vt {
	return Vt{p[0] - other[0], p[1] - other[1]}
}
func (p Vt) Mul(other Vt) Vt {
	return Vt{p[0] * other[0], p[1] * other[1]}
}
func (p Vt) Imul(other int) Vt {
	return Vt{p[0] * other, p[1] * other}
}
func (p Vt) Idiv(other int) Vt {
	return Vt{p[0] / other, p[1] / other}
}
func (p Vt) Sqd(q Vt) int {
	var sum int
	for dim, pCoord := range p {
		d := pCoord - q[dim]
		sum += d * d
	}
	return sum
}

func (p Vt) Dot(other Vt) int {
	return p[0]*other[0] + p[1]*other[1]
}

// reflect plane( == normal vector )
func (p Vt) Reflect(normal Vt) Vt {
	d := 2 * (p[0]*normal[0] + p[1]*normal[1])
	return Vt{p[0] - d*normal[0], p[1] - d*normal[1]}
}

func RandVector2D(rnd *rand.Rand, st, end int) Vt {
	return Vt{
		rnd.IntRange(st, end),
		rnd.IntRange(st, end),
	}
}

func RandVector(rnd *rand.Rand, st, end Vt) Vt {
	return Vt{
		rnd.IntRange(st[0], end[0]),
		rnd.IntRange(st[1], end[1]),
	}
}
