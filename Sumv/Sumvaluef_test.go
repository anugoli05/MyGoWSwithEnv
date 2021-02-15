package sumvalue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//TestSumvalue function validates Sumvalue function
func TestSumvalue(t *testing.T) {
	t.Logf("\n-----------------------------------------------------------------------")
	t.Logf("\n-----------------------------------------------------------------------")
	t.Logf("\nThis is message before calling Sum function the test go file")
	/*Calling max function this is a comment */
	var Sumresultvalue int
	Sumresultvalue = Sumvalue(15, 5)
	//To log any info while test is running the test has to be run in command prompt with go test -v command to show log
	t.Logf("\nThis is message from t.Logf and Maxvalue returned is  %d", Sumresultvalue)
	// assert.Equal(t,Maxvalue,15)

	// if Sumresultvalue != 20 {

	// 	t.Errorf("Expected Max value is  %d and Actual Value is %d", 20, Sumresultvalue)
	// 	// t.FailNow()
	// }
	assert.Equal(t, 20, Sumresultvalue)
	t.Log("Testpassed after the assert logic of max function in test go file.")

}
