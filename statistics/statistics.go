/*
package imlements mathematition statistics functions
*/
package statistics

import (
	"math"
)

/*
calculates  the arithmetic average.  
*/
func Average(nmbs []float64) float64 {
	if len(nmbs) == 0 {
		return 0.0
	}
	s := 0.0
	for _, n := range nmbs {
		s += n
	}
	return s / float64(len(nmbs))
}

/*
calculates the variance of a sets of integer numbers.  
*/
func Variance(nmbs []float64) float64 {
	if len(nmbs) == 0 {
		return 0.0
	}
	s := 0.0
	av := Average(nmbs)
	for _, n := range nmbs {
		s += (n - av) * (n - av)
	}
	return s / float64(len(nmbs))
}

/*
calculates the standard deviation of a sets of integer numbers.  
*/
func StandardDeviation(nmbs []float64) float64 {
	return math.Sqrt(Variance(nmbs))
}

/*
calculates  the covariance .  
*/
func Covariance(X,Y []float64) float64 {
	n:=len(X)
	if len(Y)<n {
		n=len(Y)
	}
	Ex,Ey:=Average(X),Average(Y)

	s := 0.0
	for i := 0; i < n; i++ {
		s += (X[i]-Ex)*(Y[i]-Ey)
	}
	return s / float64(n)
}


