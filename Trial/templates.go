package Trial

import (
	"fmt"
	"html/template"
	"k8s.io/apimachinery/pkg/api/resource"
	"os"
)

const isNotReadReplica = `
	(and
		(eq (5==5) "true")
	)`
const (
	ConvergeImagesDefaultTemplate = "converge-images"
)

var isNotReadReplica_var = fmt.Sprintf("{{and %s}}", isNotReadReplica)

type OperationTemplate struct {
	Guard          string `json:"guard,omitempty"`
	RetryOnFailure bool   `json:"retry_on_failure"`
}

func TestTemplate() {
	name := "Aman CHhabra"
	fmt.Println(name)
	fmt.Println(isNotReadReplica_var)
	operation := OperationTemplate{
		Guard:          isNotReadReplica_var,
		RetryOnFailure: false,
	}
	fmt.Println(operation)
	labels := map[string]string{
		"type":     "postgresql",
		"template": ConvergeImagesDefaultTemplate,
	}
	tmp, _ := template.New("Hello").Parse(isNotReadReplica_var)
	tmp.ExecuteTemplate(os.Stdout, "", labels)
}

func TestQuantity() {

	// Creating a quantity with a value of 1 GB:
	//quantity := resource.MustParse("1Gi")

	// Converting a quantity to a string:
	//str := quantity.String()
	//fmt.Println(str)

	// Adding two quantities together:
	quantity1, err := resource.ParseQuantity("4Gi")
	if err != nil {
		fmt.Println(err)
	}
	//quantity2 := resource.MustParse("200Mi")
	quantity3, err := resource.ParseQuantity("1073741824")
	fmt.Println(quantity1.Value())
	fmt.Println(quantity3.Value())
	//if err != nil {
	//	fmt.Println(err)
	//}
	//quantity4 := resource.NewQuantity(4, "DecimalSI")
	//quantity1.Add(quantity3)
	//fmt.Println(resource.Format("2147483648"))
	//if quantity3.Value() < quantity1.Value() {
	//	fmt.Println(true)
	//} else {
	//	fmt.Println(false)
	//}
	fmt.Println(quantity1.String())
	//fmt.Println(quantity4)
}
