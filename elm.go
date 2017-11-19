package elm

import (
	"math/rand"

	"github.com/Yamashou/mpinverse"
	"gonum.org/v1/gonum/mat"
)

type ELM struct {
	W     mat.Dense
	Beta  mat.Dense
	Train DataSet
	Test  DataSet
}

func (e *ELM) getWeightMatrix(X mat.Dense) mat.Dense {
	var b mat.Dense
	b.Mul(&e.W, X.T())
	return mpinverse.NewMPInverse(setSigmoid(b))
}

func (e *ELM) Fit(d *DataSet, hidNum int) {
	var data mat.Dense

	xArray := e.getAddBiasArray(d)
	rundArray := getRundomArray(hidNum, d.XSize+1)
	yArray := mat.NewDense(len(d.Y), d.YSize, d.Y)
	e.W = *rundArray

	H := e.getWeightMatrix(*xArray)
	data.Mul(H.T(), yArray)
	e.Beta = data
}

func (e *ELM) Score(d *DataSet) {
	var data mat.Dense

	testArray := e.getAddBiasArray(d)
	data.Mul(&e.W, testArray.T())

	gData := setSigmoid(data)
	var data2 mat.Dense
	data2.Mul(gData.T(), &e.Beta)
	evaluationCheck(data2, d.Y)
}

func (e *ELM) getAddBiasArray(d *DataSet) *mat.Dense {
	dataSize := d.XSize
	t := addBias(d.X, len(d.X)/dataSize, dataSize)
	return mat.NewDense(len(t)/(dataSize+1), dataSize+1, t)
}

func getRundomArray(n, m int) *mat.Dense {
	data := make([]float64, n*m)
	for i := range data {
		data[i] = rand.NormFloat64() / 10
	}
	return mat.NewDense(n, m, data)
}

func addBias(X []float64, n, m int) []float64 {
	result := make([]float64, n*(m+1))
	k := 0
	count := 0
	for _, v := range X {
		result[k] = v
		if count == (m - 1) {
			result[k+1] = 1.0
			k += 2
			count = 0
			continue
		}
		k++
		count++
	}
	return result
}