package htmltopdf

import (
	"bytes"
	"fmt"
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"html/template"
	"io/ioutil"
	"os"
)

func HtmlToPdf(templateDir string, templateFileName string, data interface{}, pdfFileName string) (pdfGenerator *wkhtmltopdf.PDFGenerator, err error){
	// create new generator
	pdfGenerator = new(wkhtmltopdf.PDFGenerator)

	// parsing template into buffer
	buffer, err := ParseTemplateToBuffer(templateDir+templateFileName, data)
	if err != nil{
		return pdfGenerator, err
	}

	// write buffer to temp html file
	htmlFile, err := CreateHtmlTempFileAndWriteBuffer(templateDir, buffer)
	if err != nil{
		return pdfGenerator, err
	}
	defer os.Remove(htmlFile.Name())

	// generate pdf from html
	tempPdfFilePath := templateDir + pdfFileName
	pdfGenerator, err = CreatePDFGeneratorAndParseHTMLtoPDF(htmlFile.Name(), wkhtmltopdf.OrientationLandscape, tempPdfFilePath)
	if err != nil{
		return pdfGenerator, err
	}
	defer os.Remove(tempPdfFilePath)

	return pdfGenerator, err
}

func ParseTemplateToBuffer(filePathToParse string, data interface{}) (tpl *bytes.Buffer, err error){

	// new buffer
	tpl = new(bytes.Buffer)

	// parse file to template
	tmpl, err := template.ParseFiles(filePathToParse)
	if err != nil {
		fmt.Println(err.Error())
		return tpl, err
	}

	// execute template into buffer
	err = tmpl.Execute(tpl, data)
	if err != nil {
		fmt.Println(err.Error())
		return tpl, err
	}

	return tpl, err
}

func CreateHtmlTempFileAndWriteBuffer(fileDir string, tpl *bytes.Buffer) (file *os.File, err error){

	// create new temp html template file
	file, err = ioutil.TempFile(fileDir,"temp.*.html") //Create temp because NewPageReader(pdf generator) doesn't render local image
	if err != nil {
		return file, err
	}

	// write result of execute template
	_, err = file.Write(tpl.Bytes()) //Same as file.writeString(tpl.String())
	if err != nil {
		return file, err
	}

	return file, err // dont forget to remove file os.Remove(file.Name())
}


func CreatePDFGeneratorAndParseHTMLtoPDF(htmlTempFilePath string, orientation string, pdfTempFilePath string) (generator *wkhtmltopdf.PDFGenerator, err error){

	// create pdf generator
	pdfGenerator, err := wkhtmltopdf.NewPDFGenerator() // first install wkhtmltopdf ubuntu with apt install wkhtmltopdf
	if err != nil {
		return pdfGenerator, err
	}

	// add pdf page and set configuration
	pdfGenerator.AddPage(wkhtmltopdf.NewPage(htmlTempFilePath)) // https://github.com/SebastiaanKlippert/go-wkhtmltopdf/issues/28#issuecomment-585308645 -> cannot render images
	pdfGenerator.Orientation.Set(orientation)
	pdfGenerator.Dpi.Set(300)

	// create new pdf file
	err = pdfGenerator.Create()
	if err != nil {
		return pdfGenerator, err
	}

	// write buffer into new pdf file
	err = pdfGenerator.WriteFile(pdfTempFilePath)
	if err != nil {
		return pdfGenerator, err
	}

	return pdfGenerator, err
}
