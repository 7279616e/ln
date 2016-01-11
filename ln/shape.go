package ln

type Shape interface {
	Compile()
	BoundingBox() Box
	Intersect(Ray) Hit
	Paths() Paths
}

type TransformedShape struct {
	Shape
	Matrix  Matrix
	Inverse Matrix
}

func NewTransformedShape(s Shape, m Matrix) Shape {
	return &TransformedShape{s, m, m.Inverse()}
}

func (s *TransformedShape) BoundingBox() Box {
	return s.Matrix.MulBox(s.Shape.BoundingBox())
}

func (s *TransformedShape) Intersect(r Ray) Hit {
	return s.Shape.Intersect(s.Inverse.MulRay(r))
}

func (s *TransformedShape) Paths() Paths {
	return s.Shape.Paths().Transform(s.Matrix)
}
