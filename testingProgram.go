/* Application: Culinary Recipe Management & Search App (MyRecipe)
Description: Final project program for managing, sorting, and searching recipes. */

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

/* Main function acts as the starting point of the app.
Uses a simple boolean variable to keep the menu running in a loop until the user wants to quit.
Takes the user's choice and calls the right function, passing the recipes array so everything stays updated. */
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

/* Function to let users add a new recipe into the system.
Checks if our array has reached the NMAX limit so the program doesn't crash.
If there is still space, it asks for the basic info like name and time,
and then loops to get all the ingredients and cooking steps one by one. */
func addRecipe(recipes *recipeList, n *int) {
	var i int

	if *n >= NMAX {
		fmt.Println("Recipe list is full. Cannot add more recipes.")
		return
	}

	fmt.Print("Recipe Name (Use '_' for Spaces): ")
	fmt.Scan(&recipes[*n].name)
	fmt.Print("Category (Use '_' for Spaces): ")
	fmt.Scan(&recipes[*n].category)
	fmt.Print("Cooking Time (in minutes): ")
	fmt.Scan(&recipes[*n].cookingTime)

	fmt.Print("Number of Ingredients (Use '_' for Spaces): ")
	fmt.Scan(&recipes[*n].countIngredients)
	for i = 0; i < recipes[*n].countIngredients; i++ {
		if i == 0 {
			fmt.Printf("Ingredient %d: ", i+1)
		} else {
			fmt.Printf("Ingredient %d: ", i+1)
		}
		fmt.Scan(&recipes[*n].ingredients[i])
	}

	fmt.Print("Number of Cooking Steps (Use '_' for Spaces): ")
	fmt.Scan(&recipes[*n].countSteps)
	for i = 0; i < recipes[*n].countSteps; i++ {
		fmt.Printf("Step %d: ", i+1)
		fmt.Scan(&recipes[*n].steps[i])
	}

	recipes[*n].searchCount = 0
	*n++
	fmt.Println("\n[Recipe added successfully.]")
}

/* Function that handles the editing part.
Asks for the recipe's name, searches for it, and if exists, it opens up a smaller sub menu.
Inside the sub menu, user can pick exactly what they want to fix, e.g like replacing a wrong ingredient
or changing the cooking steps without having to rewrite the whole recipe from scratch. */
func editRecipe(recipes *recipeList, n int) {
	var title string
	var choice, ingredientNumber, i, index int

	if n == 0 {
		fmt.Println("\n[No recipes to edit. Please add some recipes first.]")
		return
	}

	displayRecipesName(recipes, n)
	fmt.Println("=======================================")
	fmt.Print("\nEnter the exact name of the Recipe to edit: ")
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
		fmt.Println("5. Main Menu")
		fmt.Println("=======================================")
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Print("New Title (Use '_' for Spaces]: ")
			fmt.Scan(&recipes[index].name)
			fmt.Println("Title changed.")
		case 2:
			if recipes[index].countIngredients == 0 {
				fmt.Println("This recipe has no ingredients.")
			} else {
				fmt.Println("Current Ingredients:")
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
				fmt.Print("Enter the new ingredient (Use '_' for Spaces): ")
				fmt.Scan(&recipes[index].ingredients[ingredientNumber-1])
				fmt.Println("Ingredient updated.")
			}
		case 3:
			fmt.Print("Enter new cooking time (in minutes): ")
			fmt.Scan(&recipes[index].cookingTime)
			fmt.Println("Cooking time updated.")
		case 4:
			fmt.Print("Enter new cooking steps (Use '_' for Spaces): ")
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

/* Function to remove/delete the recipe, if user wants to throw away a recipe.
Finds the exact location(index) of the recipe in the array.
To delete, basically shift all the recipes that come after it one step to the left,
covering the old data, and then reduce the total recipe count. */
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

/* A quick display function.
Only prints the names of the recipes without the heavy details e.g like ingredients or steps.
Usually called before edit or delete so the user can easily see what's currently available in the array. */
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

/* Shows the simple name list.
Gives the user a chance to quickly sort the names alphabetically (A-Z or Z-A)
to make it easier to read, or they can just go back to the main menu. */
func displayRecipesSubMenu(recipes *recipeList, n int) {
	var choice string
	var status bool

	if n == 0 {
		return
	}

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

/* A display function for the users to see everything about the recipes.
Gives options to sort the data first, either by name or by how fast it takes to cook.
After sorting, it loops through everything and prints the full details. */
func displayRecipes(recipes *recipeList, n int) {
	var i, choice int

	if n == 0 {
		fmt.Println("\n[No recipes to display. Please add some recipes.]")
		return
	}

	fmt.Println("\n=======================================")
	fmt.Println("           ALL RECIPES DETAILS         ")
	fmt.Println("=======================================")
	for i = 0; i < n; i++ {
		printRecipeDetails(recipes[i])
	}

	status = true
	for status {
		fmt.Println("=======================================")
		fmt.Print("Sort by Name [N], Time [T], Menu [M] : ")
		fmt.Scan(&choice)

		switch choice {
		case "N", "n":
			fmt.Print("Ascending [A] or Descending [D] : ")
			fmt.Scan(&pick)
			switch pick {
			case "A", "a":
				for i = 0; i < n; i++ {
					SortbyNameAscending(recipes, n)
					printRecipeDetails(recipes[i])
				}
			case "D", "d":
				for i = 0; i < n; i++ {
					SortbyNameDescending(recipes, n)
					printRecipeDetails(recipes[i])
				}
			}
		case "T", "t":
			fmt.Print("Ascending [A] or Descending [D] : ")
			fmt.Scan(&pick)
			switch pick {
			case "A", "a":
				SortbyTimeAscending(recipes, n)
				for i = 0; i < n; i++ {

					printRecipeDetails(recipes[i])
				}
			case "D", "d":
				SortbyTimeDescending(recipes, n)
				for i = 0; i < n; i++ {
					printRecipeDetails(recipes[i])
				}
			}
		case "M", "m":
			status = false
		default:
			fmt.Println("\n[Invalid choice. Please try again.]")
		}
	}
}

/* standard selection sort algorithm. arranging the recipe list
alphabetically from A-Z based on the recipe name. */
func SortbyNameAscending(recipes *recipeList, n int) {
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

/* Selection sort, sorts the recipes in reverse alphabetical order, Z-A. */
func SortbyNameDescending(recipes *recipeList, n int) {
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

/* For cooking times, uses the insertion sort algorithm. Organizing the list
starting from the fastest recipe to cook up to the longest one. */
func SortbyTimeAscending(recipes *recipeList, n int) {
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

/* Reverse version of the cooking time sort using Insertion Sort.
Puts the recipes that take the longest time to cook at the very top. */
func SortbyTimeDescending(recipes *recipeList, n int) {
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

/* Acts as a gateway for the search feature. We check if the array
is empty first so we don't waste the user's time. Then it asks whether they want
to use the basic Sequential search or the faster Binary search. */
func search(recipes *recipeList, n int) {
	var choice string
	var status bool

	if n == 0 {
		fmt.Println("\n[No recipes available to search. Please add recipes first.]")
		return
	}

	status = true
	for status {
		fmt.Printf("\nSearch by: Sequential [S], Binary [B], Main Menu [M]\n")
		fmt.Print("Enter choice: ")
		fmt.Scan(&choice)

		switch choice {
		case "S", "s":
			searchByIngredientSequential(recipes, n)
		case "B", "b":
			searchByIngredientBinary(recipes, n)
		case "M", "m":
			status = false
		default:
			fmt.Println("Invalid search method. Please try again.")
		}
	}
}

/* A helper function. Uses binary search to quickly find
the exact array index of a recipe by matching its name. */
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

/* Function to view the statistics. Counts how many recipes we
have in each category. Creates a copy of the array, sorts it based
on how many times a recipe has been searched and prints out the most searched ones. */
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

/* Helper function to print the full recipe details */
func printRecipeDetails(recipe Recipe) {
	var i int

	fmt.Println()
	fmt.Printf("Recipe: %s\n", recipe.name)
	fmt.Printf("Category: %s\n", recipe.category)
	fmt.Printf("Cooking time: %d minutes\n", recipe.cookingTime)

	fmt.Printf("Ingredients:\n")
	for i = 0; i < recipe.countIngredients; i++ {
		if i == 0 {
			fmt.Printf("  - %s (main ingredient)\n", recipe.ingredients[i])
		} else {
			fmt.Printf("  - %s\n", recipe.ingredients[i])
		}
	}

	fmt.Printf("Steps:\n")
	for i = 0; i < recipe.countSteps; i++ {
		fmt.Printf("  %d. %s\n", i+1, recipe.steps[i])
	}
	fmt.Print("=")
}

/* Binary search has a strict rule: the data must be sorted first before searching.
Because we are searching by the main ingredient, this function specifically sorts the recipes
looking at the very first string in their ingredients array. */
func sortByMainIngredientAscending(recipes *recipeList, n int) {
	var i, j, minIdx int
	var temp Recipe

	for i = 0; i < n-1; i++ {
		minIdx = i
		for j = i + 1; j < n; j++ {
			if recipes[j].ingredients[0] < recipes[minIdx].ingredients[0] {
				minIdx = j
			}
		}
		temp = recipes[i]
		recipes[i] = recipes[minIdx]
		recipes[minIdx] = temp
	}
}

/* Sequential search method.
Checks every single recipe one by one from the first index to the last.
check if the recipe's main ingredient matches what the user typed. */
func searchByIngredientSequential(recipes *recipeList, n int) {
	var ingredient string
	var found, i int

	found = 0
	fmt.Print("\nEnter main ingredient to search [Use _ for spacing]: ")
	fmt.Scan(&ingredient)

	for i = 0; i < n; i++ {
		if recipes[i].countIngredients > 0 && recipes[i].ingredients[0] == ingredient {
			recipes[i].searchCount++
			printRecipeDetails(recipes[i])
			found++
		}
	}

	if found == 0 {
		fmt.Println("Ingredient not found in any recipe.")
	} else {
		fmt.Printf("\n[%d matching recipe(s).]\n", found)
	}
}

/* Binary Search method for finding ingredients, faster for a larger data.
Splits the search area in half repeatedly.
Once it finds a matching main ingredient, sweep slightly to the right just in case if
there are multiple different recipes that use the exact same main ingredient. */
func searchByIngredientBinary(recipes *recipeList, n int) {
	var ingredient string
	var found, mid, left, right int

	found = 0
	fmt.Print("\nEnter main ingredient to search [Use _ for spacing]: ")
	fmt.Scan(&ingredient)

	sortByMainIngredientAscending(recipes, n)

	left = 0
	right = n - 1

	for left <= right {
		mid = left + (right-left)/2
		if recipes[mid].ingredients[0] >= ingredient {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	for left < n && recipes[left].ingredients[0] == ingredient {
		recipes[left].searchCount++
		printRecipeDetails(recipes[left])
		found++
		left++
	}

	if found == 0 {
		fmt.Println("\n[Not found.]")
	} else {
		fmt.Printf("\n[%d matching recipe(s).]\n", found)
	}
}
