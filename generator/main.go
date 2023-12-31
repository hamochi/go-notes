package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
	"text/template"

	"github.com/Depado/bfchroma"
	"github.com/alecthomas/chroma/formatters/html"
	"github.com/alecthomas/chroma/styles"
	"github.com/russross/blackfriday/v2"
)

type Post struct {
	Name      string
	Content   string
	Level     int
	CreatedAt string
	UpdatedAt string
	Title     string
}

type TemplatePlaceHolders struct {
	Title     string
	Menu      string
	Content   string
	Level     string
	CreatedAt string
	UpdatedAt string
}

func main() {
	notesRoot := "notes"
	var posts []Post
	menuHtml := strings.Builder{}

	err := Walk(notesRoot, &menuHtml, "", 0, &posts)
	if err != nil {
		panic(err)
	}

	menu := strings.ReplaceAll(menuHtml.String(), "<ul></ul>", "")

	rawMainTemplate, err := os.ReadFile("template/main.template.html")
	if err != nil {
		panic(err)
	}

	stitchedRawMainTemplate, err := StitchTemplates(rawMainTemplate, "template/", "%s")
	if err != nil {
		panic(err)
	}

	//Create index.html
	index, err := GeneratePostHTML(notesRoot)
	if err != nil {
		panic(err)
	}
	index.Level = 0
	index.Name = "docs" + "/index.html"
	mainMenu := strings.ReplaceAll(menu, "{{.Level}}", "")
	err = CreateHTMLFile(index, mainMenu, "", stitchedRawMainTemplate)
	if err != nil {
		panic(err)
	}

	for _, post := range posts {
		levelDots := LevelToDots(post.Level)
		menu := strings.ReplaceAll(menu, "{{.Level}}", levelDots)

		err := CreateHTMLFile(post, menu, levelDots, stitchedRawMainTemplate)
		if err != nil {
			panic(err)
		}
	}

}

func CreateHTMLFile(post Post, menu string, levelDots string, stitchedRawMainTemplate []byte) error {
	d := TemplatePlaceHolders{
		Title:     post.Title,
		Content:   post.Content,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
		Menu:      menu,
		Level:     levelDots,
	}

	tmpl, err := template.New("main.template.html").Parse(string(stitchedRawMainTemplate))
	if err != nil {
		return err
	}

	f, err := os.OpenFile(post.Name, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	defer f.Close()

	err = tmpl.Execute(f, d)
	if err != nil {
		return err
	}

	return nil
}

func Walk(root string, str *strings.Builder, parentName string, level int, posts *[]Post) error {
	level++
	str.WriteString("<ul>")
	files, err := ioutil.ReadDir(root)
	if err != nil {
		return err
	}

	for _, f := range files {
		if f.IsDir() {
			cleanN := CleanName(f.Name())
			str.WriteString("<li>")

			fileName := fmt.Sprintf("%s/%s.html", parentName, cleanN)
			has, err := HasTemplateFile(root + "/" + f.Name())
			if err != nil {
				return err
			}

			if has {
				// Create directories in case they don't exist
				err := os.MkdirAll("docs"+parentName, 0755)
				if err != nil {
					return err
				}

				post, err := GeneratePostHTML(root + "/" + f.Name())
				if err != nil {
					return err
				}

				post.Level = level
				post.Name = "docs" + fileName
				*posts = append(*posts, post)
				str.WriteString(fmt.Sprintf(`<a href="{{.Level}}%s">%s</a>`, strings.TrimLeft(fileName, "/"), strings.ReplaceAll(cleanN, "_", " ")))
			} else {
				str.WriteString(cleanN)
			}

			Walk(root+"/"+f.Name(), str, parentName+"/"+cleanN, level, posts)
			str.WriteString("</li>")

		}
	}
	str.WriteString("</ul>")
	return nil
}

func HasTemplateFile(path string) (bool, error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return false, err
	}
	for _, f := range files {
		if f.Name() == "README.template.md" {
			return true, nil

		}
	}
	return false, nil
}

func GeneratePostHTML(path string) (Post, error) {
	p := Post{}

	rawMDBuffer, err := os.ReadFile(path + "/README.template.md")
	if err != nil {
		return Post{}, err
	}

	rawMDBuffer, err = StitchTemplates(rawMDBuffer, path, "```go\n%s```")
	if err != nil {
		return Post{}, err
	}

	// Extract meta data
	parts := bytes.Split(rawMDBuffer, []byte("---"))
	if len(parts) > 1 {
		var meta struct {
			CreatedAt string
			UpdatedAt string
			Title     string
		}
		err := json.Unmarshal(parts[0], &meta)
		if err != nil {
			return Post{}, err
		}
		p.CreatedAt = meta.CreatedAt
		p.UpdatedAt = meta.UpdatedAt
		p.Title = meta.Title
		rawMDBuffer = parts[1]
	}

	// Generate README.md from each template
	err = os.WriteFile(path+"/README.md", rawMDBuffer, os.ModePerm)
	if err != nil {
		return Post{}, err
	}

	r := bfchroma.NewRenderer(
		bfchroma.ChromaStyle(styles.Xcode),
		bfchroma.ChromaOptions(html.WithClasses(true)),
	)

	p.Content = string(blackfriday.Run(rawMDBuffer, blackfriday.WithRenderer(r)))

	return p, nil
}

func StitchTemplates(content []byte, path string, wrapIn string) ([]byte, error) {
	re := regexp.MustCompile(`(?m)\<<([^>]+)>>`)
	result := re.FindAll(content, -1)

	for _, filename := range result {
		filename = bytes.Trim(filename, "<<")
		filename = bytes.Trim(filename, ">>")

		file, err := os.ReadFile(path + "/" + string(filename))
		if err != nil {
			return nil, err
		}

		fileString := string(file)
		placeholder := fmt.Sprintf("<<%s>>", filename)
		codeBlock := fmt.Sprintf(wrapIn, fileString)
		content = bytes.Replace(content, []byte(placeholder), []byte(codeBlock), -1)
	}
	return content, nil
}

func LevelToDots(level int) string {
	var dots string
	for i := 0; i < level-1; i++ {
		dots = dots + "../"
	}

	return dots
}

func CleanName(s string) string {
	nameParts := strings.Split(s, ".")
	if len(nameParts) > 1 {
		return nameParts[1]
	}
	return nameParts[0]
}
