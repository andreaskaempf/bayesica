# Bayesian inference in Go

This project aims to create a basic facility for Bayesian inference in Go. It
is currently in scripts, but is being migrated to a library that will allow
models to be specified using common distributions, and underlying parameters to
be discovered from data using Metropolis and Gibbs sampling.

Inspiration is from *Bayesian Methods for Hackers* by Cameron Davidson-Pilon
and *Doing Bayesian Data Analysis*, second edition, by John Kruschke. The
project will aim to recreate the examples from these books as test cases for
the library.

The only dependency is [Gonum](http://www.gonum.org).

To install and run the test code:
* Install go (currently using version 1.22.1)
* Clone this respository, e.g.,`git clone git@github.com:andreaskaempf/bayesica.git`
* Change to the directory, i.e., `cd bayesica`
* Run `go mod init bayesica`
* Run `go get` to install dependencies
* `go build` to compile and build
* `./bayesica` to run the script (will be replaced with test cases)

-- Andreas Kaempf
