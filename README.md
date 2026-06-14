# ResepKu - Culinary Recipe Management App

A console-based recipe management application built with Go that demonstrates fundamental data structures, searching algorithms, and sorting algorithms through a real-world culinary use case.

## Team Members

This project was developed collaboratively by both team members, with responsibilities distributed across different components of the application.

### Naufal Abdurrahman Gumelar (103012540018)

**Responsibilities:**

* Developed and refined sorting algorithms (Selection Sort and Insertion Sort)
* Implemented search algorithms (Sequential Search and Binary Search)
* Improved and optimized edit functionality
* Updated and enhanced overall program logic
* Contributed to feature improvements and algorithm integration

### Achmad Ramadhanu Putra Abdi (103012540021)

**Responsibilities:**

* Developed the main program structure and menu navigation flow
* Implemented CRUD operations (Add, Edit, and Delete)
* Conducted program testing and debugging
* Fixed syntax and runtime issues throughout development
* Improved search functionality and user interaction flow
* Assisted with application documentation and validation

### Shared Contributions

Both members collaborated on:

* Application design and workflow planning
* Feature discussion and implementation decisions
* Testing and verification of program functionality
* Code review and refinement
* Final documentation and project presentation

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

  * Number of recipes per ingredient
  * Most searched recipe leaderboard

---

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
├── myrecipe.go
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

* Array based data management
* Struct implementation in Go
* CRUD operations
* Sequential Search
* Binary Search
* Selection Sort
* Insertion Sort
* Console application development
* Basic data analytics 

---
