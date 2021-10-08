package mdTool

import (
	"testing"
)

func TestAddImg(test *testing.T){
	err := Add_image_url()
	if err != nil {
		test.Errorf("Errors: \n%v",err)
		return
	}
}