package main

import (
	"bufio"

	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"text/template"
)

type Resume struct {
	Name         string
	Email        string
	Phone        string
	Location     string
	Github       string
	Introduction string
	Skills       Skills
	Education    []Education
	Experience   []Experience
	Projects     []Project
	Achievements []Achievement
}

type Skills struct {
	Programming []string
	Frameworks  []string
	Other       []string
	Soft        []string
}

type Education struct {
	Institution string
	Degree      string
	Major       string
	Year        string
	Location    string
}

type Experience struct {
	Title   string
	Company string
	Start   string
	End     string
	Bullets []string
}

type Project struct {
	Name    string
	Tech    string
	Summary string
}

type Achievement struct {
	Name  string
	Event string
	Date  string
}

func readInput(reader *bufio.Reader, prompt string) string {
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func readMultiple(reader *bufio.Reader, prompt string) []string {
	fmt.Println(prompt, "(Enter empty line to stop):")
	var items []string
	for {
		input := readInput(reader, "- ")
		if input == "" {
			break
		}
		items = append(items, input)
	}
	return items
}

func readEducation(reader *bufio.Reader) []Education {
	var educations []Education
	var count int
	fmt.Print("How many Education entries? ")
	fmt.Scanf("%d\n", &count)
	for i := 0; i < count; i++ {
		fmt.Printf("Education #%d:\n", i+1)
		educations = append(educations, Education{
			Institution: readInput(reader, "Institution: "),
			Degree:      readInput(reader, "Degree: "),
			Major:       readInput(reader, "Major: "),
			Year:        readInput(reader, "Year: "),
			Location:    readInput(reader, "Location: "),
		})
	}
	return educations
}

func readExperience(reader *bufio.Reader) []Experience {
	var experiences []Experience
	var count int
	fmt.Print("How many Experience entries? ")
	fmt.Scanf("%d\n", &count)
	for i := 0; i < count; i++ {
		fmt.Printf("Experience #%d:\n", i+1)
		var bullets []string
		var bulletCount int
		fmt.Print("How many bullet points? ")
		fmt.Scanf("%d\n", &bulletCount)
		for j := 0; j < bulletCount; j++ {
			bullets = append(bullets, readInput(reader, fmt.Sprintf("Bullet #%d: ", j+1)))
		}
		experiences = append(experiences, Experience{
			Title:   readInput(reader, "Title: "),
			Company: readInput(reader, "Company: "),
			Start:   readInput(reader, "Start Date: "),
			End:     readInput(reader, "End Date: "),
			Bullets: bullets,
		})
	}
	return experiences
}

func readProject(reader *bufio.Reader) []Project {
	var projects []Project
	var count int
	fmt.Print("How many Project entries? ")
	fmt.Scanf("%d\n", &count)
	for i := 0; i < count; i++ {
		fmt.Printf("Project #%d:\n", i+1)
		projects = append(projects, Project{
			Name:    readInput(reader, "Name: "),
			Tech:    readInput(reader, "Technologies Used: "),
			Summary: readInput(reader, "Summary: "),
		})
	}
	return projects
}

func readAchievement(reader *bufio.Reader) []Achievement {
	var achievements []Achievement
	var count int
	fmt.Print("How many Achievement entries? ")
	fmt.Scanf("%d\n", &count)
	for i := 0; i < count; i++ {
		fmt.Printf("Achievement #%d:\n", i+1)
		achievements = append(achievements, Achievement{
			Name:  readInput(reader, "Name: "),
			Event: readInput(reader, "Event: "),
			Date:  readInput(reader, "Date: "),
		})
	}
	return achievements
}

var temp *template.Template

func init() {
	temp = template.Must(template.ParseFiles("template/resume.tex"))
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	//reader := bufio.NewReader(os.Stdin)

	resume := Resume{
		Name:         readInput(reader, "Name: "),
		Email:        readInput(reader, "Email: "),
		Phone:        readInput(reader, "Phone: "),
		Location:     readInput(reader, "Location: "),
		Github:       readInput(reader, "Github: "),
		Introduction: readInput(reader, "Introduction: "),
		Skills: Skills{
			Programming: readMultiple(reader, "Enter Programming Skills"),
			Frameworks:  readMultiple(reader, "Enter Frameworks"),
			Other:       readMultiple(reader, "Enter Other Skills"),
			Soft:        readMultiple(reader, "Enter Soft Skills"),
		},
		Education:    readEducation(reader),
		Experience:   readExperience(reader),
		Projects:     readProject(reader),
		Achievements: readAchievement(reader),
	}

	os.MkdirAll("output", os.ModePerm)

	var outFileName string
	fmt.Printf("Enter the name of file:")
	fmt.Scan(&outFileName)

	fullPath := fmt.Sprintf("output/%s.tex", outFileName)

	outFile, err := os.Create(fullPath)
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()

	err = temp.Execute(outFile, resume)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("\n%v.tex generated successfully in output folder!", outFileName)

	//compile pdf now;
	//

	fmt.Printf("Do you want to compile to pdf now? (y/n)")
	char, _, err := reader.ReadRune()
	if err != nil {
		fmt.Println("Error reading character:", err)
		return
	}

	if char == 121 || char == 89 {
		cmd := exec.Command("pdflatex", "-output-directory=output", fullPath)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("\n%v.pdf generated successfully in output folder!", outFileName)
	} else {
		fmt.Println("pdf compilation skipped. you can compile pdf later !!")
	}

}
