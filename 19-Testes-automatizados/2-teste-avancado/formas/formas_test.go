package formas

import (
	"testing"
	"math"
)

func TestArea(t *testing.T) {

	t.Run("Retangulo", func(t *testing.T) {
		ret := Retangulo{10, 12}
		esperado := 120.0
		entregue := ret.Area()

		if esperado != entregue {
			t.Errorf("Esperado %v, entregue %v.", esperado, entregue)
		}
	})

	t.Run("Círculo", func(t *testing.T)  {
		circ := Circulo{10}
		esperado := float64(math.Pi * 100)
		entregue := circ.Area()

		if esperado != entregue {
			t.Errorf("Esperado %v, entregue %v.", esperado, entregue)
		}
	})
}