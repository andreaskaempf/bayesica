// Simple Metropolis sampler, estimates the posterior distribution of a normal
// distribution from synthetically generated data.
//
// AK, 9.04.2024

package main

import (
	"fmt"
	"math"

	"time"

	"golang.org/x/exp/rand"

	"gonum.org/v1/gonum/stat/distuv"
)

// To be initialized to random source
var src rand.Source

func main() {

	// Seed random source with current time
	t := time.Now()
	src = rand.NewSource(uint64(t.Unix()))

	// Generate random samples from a normal distribution
	data := generateData(2, 1, 2000)
	fmt.Println("Mean of data =", mean(data))

	// Extract posterior samples for mean, that fit data
	// Prior assumes mean 5, sd 2
	traces := sampleNormalMu(5, 2, 0.2, data, 2000)
	fmt.Println("Mean of traces =", mean(traces[500:])) // ignore burn-in

	// Create graphs of traces, should match the mean used to create data
	histogram(traces[500:], 20)
	lineChart(traces)
}

// Use Metropolis algorithm to sample traces of mean from posterior distribution.
// Assumes the sigma (std dev) is fixed to its prior value.
//
// Arguments:
//
//	priorMu: mean assumed for prior
//	priorSd: std dev assumed for prior
//	propSd: std dev used for generating proposals
//	data: list of observed data points
//	n: number of samples to generate
//
// Returns:
//
//	list of samples (traces) of means generated
func sampleNormalMu(priorMu, priorSd, propSd float64, data []float64, n int) []float64 {

	// Start by assuming that mu and std dev are equal to priors
	curMu := priorMu
	curSd := priorSd // this never changes

	// Iterate through each sample generated
	trace := []float64{} // list of generated posterior values
	accepted := 0        // number of proposals accepted
	for i := 0; i < n; i++ {

		// Suggest new value for the parameter, just a random value from the
		// normal distribution, using the current value for mean
		normGuess := distuv.Normal{curMu, propSd, src}
		propMu := distuv.Normal.Rand(normGuess)

		// Compute likelihoods for both the current and proposed means.
		// Most direct approach would be to multiply together the probabilities
		// of each data point, but this converges to zero for lots of data points,
		// so add up the log probabilities instead.
		normCurrent := distuv.Normal{curMu, curSd, src}
		normProposal := distuv.Normal{propMu, curSd, src}
		var likelihoodProposal, likelihoodCurrent float64
		for _, x := range data {
			likelihoodCurrent += normCurrent.LogProb(x)
			likelihoodProposal += normProposal.LogProb(x)
		}

		// Compute probability of acceptance, as ratio of prior probability of proposed and
		// current mu. Since we are using log probabilities, add the likelihood and priors
		// instead of multiplying them, and subtract the probabilities instead of dividing.
		normPrior := distuv.Normal{priorMu, priorSd, src}
		priorCurrent := normPrior.LogProb(curMu)
		priorProposal := normPrior.LogProb(propMu)
		pCurrent := likelihoodCurrent + priorCurrent
		pProposal := likelihoodProposal + priorProposal
		pAccept := pProposal - pCurrent // subtract instead of divide

		// Update if proposal is accepted
		p := rand.Float64() // probability 0-1
		p = math.Log(p)     // take log since comparing Log probabilities
		if p < pAccept {
			curMu = propMu
			accepted++
		}

		// Add current value to posterior trace, even if not changed
		trace = append(trace, curMu)
	}

	fmt.Println("Accepted", accepted, "of", n, "=",
		float64(accepted)/float64(n)*100, "%")
	return trace
}

// Generate a normally distributed list of numbers
func generateData(mu, sd float64, n int) []float64 {
	data := make([]float64, n)
	dist := distuv.Normal{mu, sd, src}
	for i := 0; i < n; i++ {
		data[i] = distuv.Normal.Rand(dist)
	}
	return data
}

// Mean of a list of numbers
func mean(nums []float64) float64 {
	var res float64
	for _, x := range nums {
		res += x
	}
	return res / float64(len(nums))
}
