// Distributions for Bayesian models

package main

import (
	"fmt"

	"gonum.org/v1/gonum/stat/distuv"
)

// Distribution types
const (
	Normal = iota
	Beta
	Poisson
	Uniform
)

// Test stub for testing distributions
func testDistributions() {

	// Create a normal distribution
	n := NewNormal(5, .5)
	fmt.Println(n.SamplePrior(), n.Prob(5)) //, n.Prob(5), n.LogProb(5))

	// Create another normal distibution, with the first one as mean
	n2 := NewNormal(0, .5)
	n2.Deps = append(n2.Deps, &n)
	fmt.Println(n2.SamplePrior(), n2.Prob(5))
}

// Distribution with parameters, which may in turn be other
// distribution (set in the Deps array)
type Distribution struct {
	Type    int             // type of distribution
	Priors  []float64       // prior parameters
	Current []float64       // current parameters
	Deps    []*Distribution // parameters that come from other distributions
}

// Create a new normal parameter
func NewNormal(mu, sd float64) Distribution {
	return Distribution{Type: Normal, Priors: []float64{mu, sd}}
}

// Sample from prior of a normal distribution
func (d Distribution) SamplePrior() float64 {
	priors := d.getPriors()
	switch d.Type {
	case Normal:
		norm := distuv.Normal{priors[0], priors[1], src}
		return distuv.Normal.Rand(norm)
	default:
		panic("Unknown distribution type")
	}
}

// Generate probability for a value from the prior of a normal distribution
func (d Distribution) Prob(x float64) float64 {
	priors := d.getPriors()
	switch d.Type {
	case Normal:
		norm := distuv.Normal{priors[0], priors[1], src}
		return norm.Prob(x)
	default:
		panic("Unknown distribution type")
	}
}

// Generate log probability for a value from the prior of a distribution
func (d Distribution) LogProb(x float64) float64 {
	priors := d.getPriors()
	switch d.Type {
	case Normal:
		norm := distuv.Normal{priors[0], priors[1], src}
		return norm.LogProb(x)
	default:
		panic("Unknown distribution type")
	}
}

// Get priors, recursively from dependencies if present, otherwise values provided
// as data when creating the distribution
func (d Distribution) getPriors() []float64 {
	res := make([]float64, len(d.Priors))
	for i, p := range d.Priors {
		if len(d.Deps) > i && d.Deps[i] != nil {
			res[i] = d.Deps[i].SamplePrior()
		} else {
			res[i] = p
		}
	}
	return res
}
