package Geometric

const pi float64 = 3.14

func AreaCircle(rad int)float64{
	
	return (pi*float64(rad*rad))
}

func CircumCircle(rad int)float64{
	return (2*pi*float64(rad))
}
