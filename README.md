# ResepKu - Culinary Recipe Management App

A console-based recipe management application built with Go that demonstrates fundamental data structures, searching algorithms, and sorting algorithms through a real-world culinary use case.

## Team Members

### Naufal Abdurrahman Gumelar

**Role:** Lead Developer, Logic Architecture, Code Review

**Contributions:**

* Data structure design and array management
* Sequential Search and Binary Search implementation
* Edit and Delete operation constraints
* Statistics and analytics features

### Achmad Ramadhanu Putra Abdi

**Role:** Co-Developer, UI/UX Formatting, Algorithm Testing

**Contributions:**

* Selection Sort and Insertion Sort implementation
* Menu navigation and application flow
* Console display formatting
* Documentation and testing

> Final Project for Algoritma Pemrograman 2 - Telkom University

---

## Overview

ResepKu helps users manage and search cooking recipes based on available ingredients. The application stores recipe data in memory and applies classic searching and sorting algorithms to support efficient data retrieval and organization.

### Key Features

* **Recipe Management**

  * Add, edit, delete, and view recipes
  * Manage titles, categories, ingredients, cooking steps, and cooking duration

* **Sorting Algorithms**

  * Sort recipes alphabetically (A-Z / Z-A)
  * Sort recipes by cooking time (fastest to longest)

* **Ingredient Search**

  * Search recipes by main ingredient
  * Supports Sequential Search and Binary Search

* **Statistics Dashboard**

  * Recipe count by category
  * Most searched recipe leaderboard

---

## Getting Started

### Requirements

* Go 1.20 or later

### Run the Application

```bash
git clone https://github.com/yourusername/resepku-golang-app.git

cd resepku-golang-app

go run main.go
```

### Main Menu

```text
1. Add Recipe
2. Edit Recipe
3. Delete Recipe
4. Search by Ingredient
5. Display All Recipes
6. View Statistics
7. Exit
```

---

## Project Structure

```text
resepku-golang-app/
├── main.go
│   ├── Data Structures
│   ├── CRUD Operations
│   ├── Search Algorithms
│   ├── Sorting Algorithms
│   └── Display & Statistics
└── README.md
```

---

## Data Structures

### Recipe

```go
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
```

### Recipe Database

```go
const NMAX int = 1000

type recipeList [NMAX]Recipe
```

---

## Algorithms Implemented

### Searching

| Algorithm         | Purpose                                                |
| ----------------- | ------------------------------------------------------ |
| Sequential Search | Search recipes by main ingredient                      |
| Binary Search     | Search recipes by main ingredient after sorting        |
| Binary Search     | Locate recipes by title for edit and delete operations |

### Sorting

| Algorithm      | Purpose                          |
| -------------- | -------------------------------- |
| Selection Sort | Sort recipes alphabetically      |
| Selection Sort | Sort recipes by main ingredient  |
| Insertion Sort | Sort recipes by cooking duration |

---

## Application Workflow

### Add Recipe

* Validates storage capacity
* Prevents duplicate recipe names
* Stores recipe details and ingredients

### Edit & Delete Recipe

* Sorts recipes alphabetically
* Uses Binary Search to locate recipes by title
* Updates or removes selected recipes

### Search Recipe

* Search by main ingredient
* Supports Sequential Search and Binary Search
* Tracks recipe search frequency

### View Statistics

* Counts recipes by category
* Displays most searched recipes
* Uses a temporary sorted copy to preserve original data

---

## Core Functions

| Function               | Description                             |
| ---------------------- | --------------------------------------- |
| `addRecipe()`          | Add a new recipe                        |
| `editRecipe()`         | Modify an existing recipe               |
| `deleteRecipe()`       | Remove a recipe                         |
| `viewStatistics()`     | Display analytics and rankings          |
| `printRecipeDetails()` | Display recipe information consistently |

---

## Learning Outcomes

This project demonstrates:

* Array-based data management
* Struct implementation in Go
* CRUD operations
* Sequential Search
* Binary Search
* Selection Sort
* Insertion Sort
* Console application development
* Basic data analytics and reporting

---

## License

Created for academic purposes.
