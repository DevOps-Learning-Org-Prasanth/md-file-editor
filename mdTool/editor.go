package mdTool

import (
	"fmt"
	// "errors"
	"path/filepath"
	"io/ioutil"
	"strings"
	"regexp"
)

func Add_image_url()(error){
	// get repo link
	fmt.Println("Getting repo link")
	// get the number of files change **
	fmt.Println("Getting the list of files modified")
	// add the image link
	files_path := "../TestFiles"
	git_repo := "github.com/DevOps/notes/cloud"
	errMsg := ""
	add_image := func()(error){
		files,err := filepath.Glob(files_path + "/*")
		if err != nil {
			return err
		}
		errs := []error{}
		regex,_ := regexp.Compile("!\\[[A-Za-z\\-_0-9]+\\]")

		for _,file := range files{
			if !(isMd(file)){
				continue
			}
			fmt.Println(file)
			data , err := ioutil.ReadFile(file)
			if err != nil {
				errs = append(errs,fmt.Errorf("Filename: %v Error: %v",file,err))
			}

			lines := strings.Split(string(data),"\n")
			for i,line := range lines {
				if !regex.MatchString(line){
					continue
				}
				fmt.Printf("lineno: %d Value: %s",i,line)
				lines[i] = line + fmt.Sprintf("(%v/images/%v.png)",git_repo,line)
			}
			output := strings.Join(lines,"\n")
			err = ioutil.WriteFile(file,[]byte(output),0644)
			if err != nil {
				errs = append(errs,fmt.Errorf("filename: %v, error: %v",file,err))
			}
		}
		for _,err = range errs {
			if err == nil{
				continue
			}
			errMsg += fmt.Sprintf("%v\n",err)
		}
		
		return nil
	}
	err := add_image()
	if errMsg != "" || err != nil {
		return fmt.Errorf("Error within function: %v \n Error: %v",errMsg,err)
	}
	return nil
}

func isMd(file string) bool {
	fileExt := filepath.Ext(file)
	return fileExt == ".md"
}