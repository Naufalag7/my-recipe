//By the program name, it is mainly used to test the program workability

package main

import "fmt"

const NMAX int = 1000

type Recipe struct {
	name             string
	category         string
	ingredients      [100]string
	steps            [100]string
	countIngredients int
	countSteps       int
	cookingTime      int
	searchCount      int
}

type recipeList [NMAX]Recipe

func main() {
	var recipes recipeList
	var choice, amount int

	for {
		fmt.Println("\n=======================================================")
		fmt.Println("   CULINARY RECIPE MANAGEMENT & SEARCH APP (MyRecipe)") 
		fmt.Println("=======================================================")
		fmt.Println("1. Add recipe")
		fmt.Println("2. Edit recipe")
		fmt.Println("3. Delete recipe")
		fmt.Println("4. Search recipe")
		fmt.Println("5. Display all recipes")
		fmt.Println("6. View Statistics")
		fmt.Println("7. Exit")
		fmt.Println("========================================================")
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			addRecipe(&recipes, &amount)
		case 2:
			editRecipe(&recipes, amount)
		case 3:
			deleteRecipe(&recipes, &amount)
		case 4:
			searchByRecipe(&recipes, amount)
		case 5:
			displayRecipes(&recipes, amount)
		case 6:
			viewStatistics(recipes, amount)
		case 7:
			fmt.Println("\nExiting the application. Goodbye! 😘😘😘")
			return
		default:
			fmt.Println("\nInvalid choice. Please try again.\n")
		}
	}
}

func addRecipe(recipes *recipeList, n *int) {
	var i int
	if *n >= NMAX {
		fmt.Println("Recipe list is full. Cannot add more recipes.")
		return
	}
	fmt.Print("Enter recipe name: ")
	fmt.Scan(&recipes[*n].name)
	fmt.Print("Enter category: ")
	fmt.Scan(&recipes[*n].category)
	fmt.Print("Enter cooking time (in minutes): ")
	fmt.Scan(&recipes[*n].cookingTime)
	fmt.Print("Enter number of ingredients: ")
	fmt.Scan(&recipes[*n].countIngredients)
	for i = 0; i < recipes[*n].countIngredients; i++ {
		fmt.Printf("Ingredient %d: ", i+1)
		fmt.Scan(&recipes[*n].ingredients[i])
	}

	fmt.Print("Enter number of cooking steps: ")
	fmt.Scan(&recipes[*n].countSteps)
	for i = 0; i < recipes[*n].countSteps; i++ {
		fmt.Printf("Step %d: ", i+1)
		fmt.Scan(&recipes[*n].steps[i])
	}
	recipes[*n].searchCount = 0
	*n++
	fmt.Println("\n[Recipe added successfully.]")
}

func editRecipe(recipes *recipeList, n int) {
	var title string
	var choice, ingredientNumber, i int

	if n == 0 {
		fmt.Println("\n[No recipes to edit. Please add some recipes first.]")
		return
	}

	displayRecipesName(recipes, n)
	fmt.Println("=======================================")
	fmt.Print("\nEnter the name of the recipe to edit: ")
	fmt.Scan(&title)
	index := findIndexRecipe(recipes, n, title)
	if index == -1 {
		fmt.Println("\n[Recipe not found.]")
		return
	}

	for {
		fmt.Println("\n=======================================")
		fmt.Println("Edit recipe:", recipes[index].name)
		fmt.Println("=======================================")
		fmt.Println("1. Change title")
		fmt.Println("2. Change ingredient")
		fmt.Println("3. Change cooking time")
		fmt.Println("4. Change cooking steps")
		fmt.Println("5. Done")
		fmt.Println("=======================================")
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Print("Enter new title: ")
			fmt.Scan(&recipes[index].name)
			fmt.Println("Title changed.")
		case 2:
			if recipes[index].countIngredients == 0 {
				fmt.Println("This recipe has no ingredients.")
			} else {
				fmt.Println("Current ingredients:")
				for i = 0; i < recipes[index].countIngredients; i++ {
					fmt.Printf("%d. %s\n", i+1, recipes[index].ingredients[i])
				}
				fmt.Print("Enter ingredient number to change: ")
				fmt.Scan(&ingredientNumber)
				for ingredientNumber < 1 || ingredientNumber > recipes[index].countIngredients {
					fmt.Println("Invalid ingredient number.")
					fmt.Print("Enter ingredient number to change: ")
					fmt.Scan(&ingredientNumber)
				}
				fmt.Print("Enter the new ingredient: ")
				fmt.Scan(&recipes[index].ingredients[ingredientNumber-1])
				fmt.Println("Ingredient updated.")
			}
		case 3:
			fmt.Print("Enter new cooking time (in minutes): ")
			fmt.Scan(&recipes[index].cookingTime)
			fmt.Println("Cooking time updated.")
		case 4:
			fmt.Print("Enter new cooking steps: ")
			fmt.Scan(&recipes[index].steps)
			fmt.Println("Cooking steps updated.")
		case 5:
			fmt.Println("Finished editing recipe.")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func deleteRecipe(recipes *recipeList, n *int) {
	var title string

	if *n == 0 {
		fmt.Println("\n[No recipes to delete. Please add some recipes first.]")
		return
	}
	
	displayRecipesName(recipes, *n)
	fmt.Println("=======================================")
	fmt.Print("\nEnter title to delete: ")
	fmt.Scan(&title)
	index := findIndexRecipe(recipes, *n, title)

	if index != -1 {
		for i := index; i < *n-1; i++ {
			recipes[i] = recipes[i+1]
		}
		*n--
		fmt.Println("\n[Recipe is deleted]")
	} else {
		fmt.Println("\n[Recipe not found]")
	}
}

/*Procedure to display the recipes that are listed in
the array aswell as sorting (Ascending/Descending) */

func displayRecipesName(recipes *recipeList, n int) {
	var i int
	if n == 0 {
		fmt.Println("\n[No recipes to display. Please add some recipes.]")
		return
	}

	fmt.Println("\n=======================================")
	fmt.Println("           AVAILABLE RECIPES           ")
	fmt.Println("=======================================")
	for i = 0; i < n; i++ {
		fmt.Printf("%d. %s\n", i+1, recipes[i].name)
	}
}

func displayRecipes(recipes *recipeList, n int) {
	var i, choice int
	if n == 0 {
		fmt.Println("\n[No recipes to display. Please add some recipes.]")
		return
	}

	fmt.Println("\n=======================================")
	fmt.Println("         RECIPE DISPLAY OPTIONS        ")
	fmt.Println("=======================================")
	fmt.Println("1. Sort by name (Ascending)")
	fmt.Println("2. Sort by name (Descending)")
	fmt.Println("3. Sort by cooking time (Ascending)")
	fmt.Println("4. Sort by cooking time (Descending)")
	fmt.Println("=======================================")

	fmt.Print("Enter your choice: ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		SortbyNameAscending(recipes, n)
	case 2:
		SortbyNameDescending(recipes, n)
	case 3:
		SortbyTimeAscending(recipes, n)
	case 4:
		SortbyTimeDescending(recipes, n)
	default:
		fmt.Println("Invalid choice. Please try again.")
		return
	}

	for i = 0; i < n; i++ {
		printRecipeDetails(recipes[i])
	}
	fmt.Println()
}

func SortbyNameAscending(recipes *recipeList, n int) {
	//Selection Sort
	//Please make the Descending one aswell
	var i, j, minIdx int
	for i = 0; i < n-1; i++ {
		minIdx = i
		for j = i + 1; j < n; j++ {
			if recipes[j].name < recipes[minIdx].name {
				minIdx = j
			}
		}
		temp := recipes[i]
		recipes[i] = recipes[minIdx]
		recipes[minIdx] = temp
	}
}

func SortbyNameDescending(recipes *recipeList, n int) {
	//Selection Sort
	//Please make the Descending one aswell
	var i, j, minIdx int
	for i = 0; i < n-1; i++ {
		minIdx = i
		for j = i + 1; j < n; j++ {
			if recipes[j].name > recipes[minIdx].name {
				minIdx = j
			}
		}
		temp := recipes[i]
		recipes[i] = recipes[minIdx]
		recipes[minIdx] = temp
	}
}

func SortbyTimeAscending(recipes *recipeList, n int) {
	//Insertion Sort
	//Please make the Descending one aswell
	var i, j int
	for i = 1; i < n; i++ {
		temp := recipes[i]
		j = i - 1
		for j >= 0 && recipes[j].cookingTime > temp.cookingTime {
			recipes[j+1] = recipes[j]
			j--
		}
		recipes[j+1] = temp
	}
}

func SortbyTimeDescending(recipes *recipeList, n int) {
	//Insertion Sort
	var i, j int
	for i = 1; i < n; i++ {
		temp := recipes[i]
		j = i - 1
		for j >= 0 && recipes[j].cookingTime < temp.cookingTime {
			recipes[j+1] = recipes[j]
			j--
		}
		recipes[j+1] = temp
	}
}

func searchByRecipe(recipes *recipeList, n int) {
	var ingredient string
	var i, j, found int

	fmt.Print("Enter ingredient to search for: ")
	fmt.Scan(&ingredient)

	found = 0
	for i = 0; i < n; i++ {
		for j = 0; j < recipes[i].countIngredients; j++ {
			if recipes[i].ingredients[j] == ingredient {
				recipes[i].searchCount++
				printRecipeDetails(recipes[i])
				found++
				j = recipes[i].countIngredients
			}
		}
	}
	if found == 0 {
		fmt.Println("Ingredient not found in any recipe.")
	}
}

/* Procedure to find where the
index is located on the array */
func findIndexRecipe(recipes *recipeList, n int, title string) int {
	var find int = -1
	for i := 0; i < n; i++ {
		if recipes[i].name == title {
			find = i
		}
	}
	return find
}

/* Procedure mainly to display the statistics of
recipes per ingredient category and a list of
the most frequently searched menus */
func viewStatistics(recipes recipeList, n int) {
	var i, maxIdx int

	if n == 0 {
		fmt.Println("\n[No recipe data available yet]")
		return
	}

	maxIdx = 0
	for i = 1; i < n; i++ {
		if recipes[i].searchCount > recipes[maxIdx].searchCount {
			maxIdx = i
		}
	}

	fmt.Println("\n=== Statistics ===")
	fmt.Printf("Total Recipes: %d\n", n)

	if recipes[maxIdx].searchCount > 0 {
		fmt.Printf("Most Searched Recipe: %v (Searched %v times)\n", recipes[maxIdx].name, recipes[maxIdx].searchCount)
	} else {
		fmt.Println("\n[No search history yet]")
	}
}

/* Procedure to print all the details
of the recipe that are listed/added */
func printRecipeDetails(recipe Recipe) {

	fmt.Println()
	fmt.Printf("Recipe: %s\n", recipe.name)
	fmt.Printf("Category: %s\n", recipe.category)
	fmt.Printf("Cooking time: %d minutes\n", recipe.cookingTime)
	fmt.Printf("Ingredients:\n")
	for i := 0; i < recipe.countIngredients; i++ {
		fmt.Printf("  - %s\n", recipe.ingredients[i])
	}
	fmt.Println("-------------------------------")
}
