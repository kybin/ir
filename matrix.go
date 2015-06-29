package main

type matrix4 struct {
	aa, ab, ac, ad float64
	ba, bb, bc, bd float64
	ca, cb, cc, cd float64
	da, db, dc, dd float64
}

func IdentityMatrix4() matrix4 {
	return matrix4{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}
}

func (m matrix4) Transpose() matrix4 {
	return matrix4{
		m.aa, m.ba, m.ca, m.da,
		m.ab, m.bb, m.cb, m.db,
		m.ac, m.bc, m.cc, m.dc,
		m.ad, m.bd, m.cd, m.dd,
	}
}

func (m matrix4) Multiply(m2 matrix4) matrix4 {
	return matrix4{
		// row a
		m.aa*m2.aa + m.ab*m2.ba + m.ac*m2.ca + m.ad*m2.da,
		m.aa*m2.ab + m.ab*m2.bb + m.ac*m2.cb + m.ad*m2.db,
		m.aa*m2.ac + m.ab*m2.bc + m.ac*m2.cc + m.ad*m2.dc,
		m.aa*m2.ad + m.ab*m2.bd + m.ac*m2.cd + m.ad*m2.dd,
		// row b
		m.ba*m2.aa + m.bb*m2.ba + m.bc*m2.ca + m.bd*m2.da,
		m.ba*m2.ab + m.bb*m2.bb + m.bc*m2.cb + m.bd*m2.db,
		m.ba*m2.ac + m.bb*m2.bc + m.bc*m2.cc + m.bd*m2.dc,
		m.ba*m2.ad + m.bb*m2.bd + m.bc*m2.cd + m.bd*m2.dd,
		// row c
		m.ca*m2.aa + m.cb*m2.ba + m.cc*m2.ca + m.cd*m2.da,
		m.ca*m2.ab + m.cb*m2.bb + m.cc*m2.cb + m.cd*m2.db,
		m.ca*m2.ac + m.cb*m2.bc + m.cc*m2.cc + m.cd*m2.dc,
		m.ca*m2.ad + m.cb*m2.bd + m.cc*m2.cd + m.cd*m2.dd,
		// row d
		m.da*m2.aa + m.db*m2.ba + m.dc*m2.ca + m.dd*m2.da,
		m.da*m2.ab + m.db*m2.bb + m.dc*m2.cb + m.dd*m2.db,
		m.da*m2.ac + m.db*m2.bc + m.dc*m2.cc + m.dd*m2.dc,
		m.da*m2.ad + m.db*m2.bd + m.dc*m2.cd + m.dd*m2.dd,
	}
}
