package main

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/jung-kurt/gofpdf/v2"
)

func main() {
	names := []string{
		"Ahmed Al-maidama",
		"Ahmed Elawadi",
		"Amal alsaloom",
		"Anood khaled",
		"Bader Alobed",
		"Fatima raed Alabri",
		"Ghalia alidrisi",
		"Hamad alshalawi",
		"Hind AlZamil",
		"Khaled Faisal Alanbar",
		"khalid alquhtani",
		"Khalid Mudathir Mohamed",
		"maram alraddadi",
		"Mohammed alhammad",
		"Mohammed AlMuhawis",
		"Nawaf Aleid",
		"Nawaf khalid alsoliman",
		"Nouf alsaadoun",
		"Raghad alabri",
		"Raghad alsaleh",
		"Rian saeed abullah bawazir",
		"Salaheddin Thabet",
		"Sema Abdulrahman Bairakdar",
		"Sharif Ahmad Alabedyahia",
		"Ibrahim Mahmoud",
		"Khaled Mohamed Husaini",
		"Sarah Rashed AlOthman",
		"Areeb alshuhail",
		"Sara Abdulaziz Bin Aksrsh",
		"Thamer alshakwan",
		"Nawaf Alharbi",
		"Basma almudaifeaa",
	} // Replace this with your list of names
	template := "template-participation.pdf" // Replace this with your template file name

	for i, name := range names {
		name = strings.ToUpper(name)
		pdf := gofpdf.New("L", "mm", "A4", "")

		pdf.AddPage()

		pdf.SetFont("Arial", "B", 48)
		wd := pdf.GetStringWidth(name)

		// Set X position to middle of page minus half the string width
		pdf.SetX((297 - wd) / 2) // 297 mm is the width of an A4 page in landscape mode
		pdf.Text(pdf.GetX(), 120, name)

		tmpFileName := fmt.Sprintf("out/temp_%d_%s.pdf", i, name)
		err := pdf.OutputFileAndClose(tmpFileName)
		if err != nil {
			panic(err)
		}

		// Overlay the generated PDF with the name onto the template
		outFileName := fmt.Sprintf("out/%d_%s.pdf", i, name)
		cmd := exec.Command("pdfcpu", "stamp", "add", "-mode", "pdf", "--", tmpFileName+":1", "rot:0, scale:1.0", template, outFileName)
		output, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Printf("Command output: %s\n", output)
			panic(err)
		}

		// Remove the temporary PDF
		err = exec.Command("rm", tmpFileName).Run()
		if err != nil {
			panic(err)
		}
	}
}
