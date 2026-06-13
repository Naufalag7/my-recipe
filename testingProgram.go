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
		fmt.Println("4. Search by ingredient")
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
			search(&recipes, amount)
		case 5:
			displayRecipesName(&recipes, amount)
			displayRecipesNameSubMenu(&recipes, amount)
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
	var choice, ingredientNumber, i, index int

	if n == 0 {
		fmt.Println("\n[No recipes to edit. Please add some recipes first.]")
		return
	}

	displayRecipesName(recipes, n)
	fmt.Println("=======================================")
	fmt.Print("\nEnter the name of the recipe to edit: ")
	fmt.Scan(&title)
	SortbyNameAscending(recipes, n)
	index = findIndexRecipe(recipes, n, title)
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
			for i = 0; i < recipes[index].countSteps; i++ {
				fmt.Printf("Step %d: ", i+1)
				fmt.Scan(&recipes[index].steps[i])
			}
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
	var index, i int

	if *n == 0 {
		fmt.Println("\n[No recipes to delete. Please add some recipes first.]")
		return
	}

	displayRecipesName(recipes, *n)
	fmt.Println("=======================================")
	fmt.Print("\nEnter title to delete: ")
	fmt.Scan(&title)
	SortbyNameAscending(recipes, *n)
	index = findIndexRecipe(recipes, *n, title)

	if index != -1 {
		for i = index; i < *n-1; i++ {
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

func displayRecipesNameSubMenu(recipes *recipeList, n int) {
	var choice string
	var status bool
	status = true
	for status {
		fmt.Println("=======================================")
		fmt.Print(" Ascending [A], Descending [D], Main Menu [M] : ")
		fmt.Scan(&choice)
		switch choice {
		case "A", "a":
			SortbyNameAscending(recipes, n)
			displayRecipesName(recipes, n)
		case "D", "d":
			SortbyNameDescending(recipes, n)
			displayRecipesName(recipes, n)
		case "M", "m":
			status = false
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
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
	var temp Recipe
	for i = 0; i < n-1; i++ {
		minIdx = i
		for j = i + 1; j < n; j++ {
			if recipes[j].name < recipes[minIdx].name {
				minIdx = j
			}
		}
		temp = recipes[i]
		recipes[i] = recipes[minIdx]
		recipes[minIdx] = temp
	}
}

func SortbyNameDescending(recipes *recipeList, n int) {
	//Selection Sort
	//Please make the Descending one aswell
	var i, j, minIdx int
	var temp Recipe
	for i = 0; i < n-1; i++ {
		minIdx = i
		for j = i + 1; j < n; j++ {
			if recipes[j].name > recipes[minIdx].name {
				minIdx = j
			}
		}
		temp = recipes[i]
		recipes[i] = recipes[minIdx]
		recipes[minIdx] = temp
	}
}

func SortbyTimeAscending(recipes *recipeList, n int) {
	//Insertion Sort
	//Please make the Descending one aswell
	var i, j int
	var temp Recipe
	for i = 1; i < n; i++ {
		temp = recipes[i]
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
	var temp Recipe
	for i = 1; i < n; i++ {
		temp = recipes[i]
		j = i - 1
		for j >= 0 && recipes[j].cookingTime < temp.cookingTime {
			recipes[j+1] = recipes[j]
			j--
		}
		recipes[j+1] = temp
	}
}

func search(recipes *recipeList, n int) {
	var choice string
	var status bool
	status = true
	for status {
		fmt.Printf("Search by: Sequential [S], Binary [B], Done [D]\n")
		fmt.Scan(&choice)
		switch choice {
		case "S", "s":
			searchByIngredientSequential(recipes, n)
		case "B", "b":
			searchByIngredientBinary(recipes, n)
		case "D", "d":
			status = false
		default:
			fmt.Println("Invalid search method. Please try again.")
		}
	}
}

/* Procedure to find where the
index is located on the array , use binary search*/
func findIndexRecipe(recipes *recipeList, n int, title string) int {
	var left, right, mid, idx int
	left = 0
	right = n - 1
	idx = -1
	for left <= right && idx == -1 {
		mid = left + (right-left)/2
		if recipes[mid].name == title {
			idx = mid
		} else if recipes[mid].name < title {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return idx
}

/* Procedure mainly to display the statistics of
recipes per ingredient category and a list of
the most frequently searched menus */
func viewStatistics(recipes recipeList, n int) {
	var i, j, categoryTotal, maxIdx int
	var categories [NMAX]string
	var categoryCounts [NMAX]int
	var sorted [NMAX]Recipe
	var found, hasSearch bool
	var temp Recipe

	if n == 0 {
		fmt.Println("\n[No recipe data available yet]")
		return
	}

	fmt.Println("\n=== Statistics ===")
	fmt.Printf("Total Recipes: %d\n", n)

	//Count recipes per category
	categoryTotal = 0
	for i = 0; i < n; i++ {
		found = false
		for j = 0; j < categoryTotal; j++ {
			if categories[j] == recipes[i].category {
				categoryCounts[j]++
				found = true
			}
		}
		if !found {
			categories[categoryTotal] = recipes[i].category
			categoryCounts[categoryTotal] = 1
			categoryTotal++
		}
	}

	fmt.Println("\n=== Recipes per Category ===")
	for i = 0; i < categoryTotal; i++ {
		fmt.Printf("  %s: %d recipe(s)\n", categories[i], categoryCounts[i])
	}

	//List most frequently searched recipes
	hasSearch = false
	for i = 0; i < n; i++ {
		if recipes[i].searchCount > 0 {
			hasSearch = true
		}
	}

	fmt.Println("\n=== Most Searched Recipes ===")
	if !hasSearch {
		fmt.Println("[No search history yet]")
	} else {
		for i = 0; i < n; i++ {
			sorted[i] = recipes[i]
		}
		for i = 0; i < n-1; i++ {
			maxIdx = i
			for j = i + 1; j < n; j++ {
				if sorted[j].searchCount > sorted[maxIdx].searchCount {
					maxIdx = j
				}
			}
			temp = sorted[i]
			sorted[i] = sorted[maxIdx]
			sorted[maxIdx] = temp
		}
		for i = 0; i < n; i++ {
			if sorted[i].searchCount > 0 {
				fmt.Printf("  %d. %s - searched %d time(s)\n", i+1, sorted[i].name, sorted[i].searchCount)
			}
		}
	}
}

/* Procedure to print all the details
of the recipe that are listed/added */
func printRecipeDetails(recipe Recipe) {
	var i int
	fmt.Println()
	fmt.Printf("Recipe: %s\n", recipe.name)
	fmt.Printf("Category: %s\n", recipe.category)
	fmt.Printf("Cooking time: %d minutes\n", recipe.cookingTime)
	fmt.Printf("Ingredients:\n")
	for i = 0; i < recipe.countIngredients; i++ {
		fmt.Printf("  - %s\n", recipe.ingredients[i])
	}
	fmt.Println("-------------------------------")
}

//helper function
func checkIngredient(recipe Recipe, ingredient string) bool {
	var i int
	var found bool
	found = false
	for i = 0; i < recipe.countIngredients; i++ {
		if recipe.ingredients[i] == ingredient {
			found = true
		}
	}
	return found
}

func sortByIngredientsAscending(recipes *recipeList, n int) {
	var i, j int
	var temp Recipe
	for i = 0; i < n-1; i++ {
		for j = 0; j < n-i-1; j++ {
			if recipes[j].countIngredients > recipes[j+1].countIngredients {
				temp = recipes[j]
				recipes[j] = recipes[j+1]
				recipes[j+1] = temp
			}
		}
	}
}

func searchByIngredientSequential(recipes *recipeList, n int) { //sequential search
	var ingredient string
	var found, i int
	found = 0
	fmt.Print("Enter ingredient to search [Use _ for spacing]: ")
	fmt.Scan(&ingredient)
	for i = 0; i < n; i++ {
		if checkIngredient(recipes[i], ingredient) {
			recipes[i].searchCount++
			printRecipeDetails(recipes[i])
			found++
		}
	}
	if found == 0 {
		fmt.Println("Ingredient not found in any recipe.")
	}
}

func searchByIngredientBinary(recipes *recipeList, n int) { //binary search
	var ingredient string
	var found, mid, left, right int
	found = 0
	fmt.Print("Enter ingredient to search [Use _ for spacing]: ")
	fmt.Scan(&ingredient)
	sortByIngredientsAscending(recipes, n)
	left = 0
	right = n - 1
	for left <= right {
		mid = left + (right-left)/2
		if checkIngredient(recipes[mid], ingredient) {
			recipes[mid].searchCount++
			printRecipeDetails(recipes[mid])
			found++
		} else if recipes[mid].ingredients[0] < ingredient {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if found == 0 {
		fmt.Println("Ingredient not found in any recipe.")
	}
}
