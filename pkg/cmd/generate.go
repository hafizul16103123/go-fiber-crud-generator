package cmd

import (
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/spf13/cobra"
)

var (
	resourceName string
)
var rootPath string = "src/modules/"

func init() {
	generateCmd.Flags().StringVarP(&resourceName, "name", "n", "", "Name of the resource (required)")
	generateCmd.MarkFlagRequired("name")
	rootCmd.AddCommand(generateCmd)
}

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a CRUD resource",
	Run: func(cmd *cobra.Command, args []string) {
		if resourceName == "" {
			fmt.Println("Error: Resource name is required")
			return
		}

		// Create directories
		os.MkdirAll(rootPath+"models", os.ModePerm)
		os.MkdirAll(rootPath+"repository", os.ModePerm)
		os.MkdirAll(rootPath+"services", os.ModePerm)
		os.MkdirAll(rootPath+"controllers", os.ModePerm)
		os.MkdirAll(rootPath+"routes", os.ModePerm)

		// Generate files
		generateModel()
		generateRepository()
		generateService()
		generateController()
		generateRoutes()

		// Update main.go to include the new routes
		if err := updateMainGo(); err != nil {
			fmt.Println("Error updating main.go:", err)
		}

		fmt.Printf("CRUD resource '%s' generated successfully!\n", resourceName)
	},
}

func getModuleName() string {
	// if info, ok := debug.ReadBuildInfo(); ok {
	//     return info.Main.Path // Returns the module name from go.mod
	// }
	// return ""

	return "go-crud"
}

func generateModel() {
	tmpl := template.Must(template.ParseFiles("pkg/templates/model.tmpl"))
	file, _ := os.Create(fmt.Sprintf(rootPath+"models/%s.go", strings.ToLower(resourceName)))
	defer file.Close()

	tmpl.Execute(file, map[string]string{
		"ResourceName": resourceName,
	})
}

func generateRepository() {
	tmpl := template.Must(template.ParseFiles("pkg/templates/repository.tmpl"))
	file, _ := os.Create(rootPath + "repository/repository.go")
	defer file.Close()

	tmpl.Execute(file, nil)
}

func generateService() {
	tmpl := template.Must(template.ParseFiles("pkg/templates/service.tmpl"))
	file, _ := os.Create(fmt.Sprintf(rootPath+"services/%s_service.go", strings.ToLower(resourceName)))
	defer file.Close()

	tmpl.Execute(file, map[string]string{
		"BaseModuleName":    getModuleName(),
		"ResourceName":      resourceName,
		"ResourceNameLower": strings.ToLower(resourceName),
	})
}

func generateController() {
	tmpl := template.Must(template.ParseFiles("pkg/templates/controller.tmpl"))
	file, _ := os.Create(fmt.Sprintf(rootPath+"controllers/%s_controller.go", strings.ToLower(resourceName)))
	defer file.Close()

	tmpl.Execute(file, map[string]string{
		"BaseModuleName":    getModuleName(),
		"ResourceName":      resourceName,
		"ResourceNameLower": strings.ToLower(resourceName),
	})
}

func generateRoutes() {
	tmpl := template.Must(template.ParseFiles("pkg/templates/routes.tmpl"))
	file, _ := os.Create(fmt.Sprintf(rootPath+"routes/%s_routes.go", strings.ToLower(resourceName)))
	defer file.Close()

	tmpl.Execute(file, map[string]string{
		"BaseModuleName":    getModuleName(),
		"ResourceName":      resourceName,
		"ResourceNameLower": strings.ToLower(resourceName),
	})
}

// updateMainGo updates the main.go file to include the new routes
func updateMainGo() error {
	mainGoPath := "src/main.go" // Path to your main.go file
	importPath := fmt.Sprintf("\"go-crud/%sroutes\"", "modules/")
	routeSetup := fmt.Sprintf("routes.Setup%sRoutes(app, repo)", resourceName)

	// Read the main.go file
	content, err := os.ReadFile(mainGoPath)
	if err != nil {
		return fmt.Errorf("failed to read main.go: %v", err)
	}

	// Convert content to string
	fileContent := string(content)

	// Add imports if they don't already exist
	requiredImports := []string{
		`"context"`,
		`"log"`,
		`"go.mongodb.org/mongo-driver/mongo"`,
		`"go.mongodb.org/mongo-driver/mongo/options"`,
		`"go-crud/modules/repository"`,
		importPath,
	}

	for _, imp := range requiredImports {
		if !strings.Contains(fileContent, imp) {
			importIndex := strings.Index(fileContent, "import (")
			if importIndex == -1 {
				return fmt.Errorf("could not find import block in main.go")
			}

			// Insert the new import
			fileContent = fileContent[:importIndex+8] + "\n\t" + imp + fileContent[importIndex+8:]
		}
	}

	// Add MongoDB connection and repository initialization after app := fiber.New()
	appIndex := strings.Index(fileContent, "app := fiber.New()")
	if appIndex == -1 {
		return fmt.Errorf("could not find 'app := fiber.New()' in main.go")
	}

	// Find the end of the line where app is initialized
	appLineEnd := strings.Index(fileContent[appIndex:], "\n")
	if appLineEnd == -1 {
		return fmt.Errorf("could not find end of line after 'app := fiber.New()'")
	}

	// MongoDB connection and repository initialization code
	mongoIndex := strings.Index(fileContent, "// Connect to MongoDB")
	fmt.Println("mongoIndex:", mongoIndex)
	mongoCode := ""
	if mongoIndex == -1 {
		mongoCode = `
		// Connect to MongoDB
		client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
		if err != nil {
			log.Fatal("Failed to connect to MongoDB:", err)
		}
		db := client.Database("go_db")
	
		// Initialize repository
		repo := repository.NewRepository(db)
		`
	}

	// Insert the MongoDB connection and repository initialization code
	fileContent = fileContent[:appIndex+appLineEnd+1] + mongoCode + fileContent[appIndex+appLineEnd+1:]

	// Find the app.Listen(":3000") line
	listenIndex := strings.Index(fileContent, "app.Listen(\":3000\")")
	if listenIndex == -1 {
		return fmt.Errorf("could not find 'app.Listen(\":3000\")' in main.go")
	}

	// Insert the route setup before app.Listen(":3000")
	fileContent = fileContent[:listenIndex] + "\n\t" + routeSetup + "\n" + fileContent[listenIndex:]

	// Write the updated content back to main.go
	if err := os.WriteFile(mainGoPath, []byte(fileContent), 0644); err != nil {
		return fmt.Errorf("failed to write main.go: %v", err)
	}

	return nil
}
