package main

import "fmt"

func main() {
	var input int

	PrintTitle("Selamat datang di Tugas Besar Alpro Aplikasi Analisis Sentimen Kelompok 2")

	for input != 4 {
		err := PrintMenu("Pilih Menu", [255]string{"Login", "Register", "Admin", "Exit"}, 4, &input)
		if err != nil {
			return
		}

		switch input {
		case 1:
			LoginView()
		case 2:
			RegisterView()
		case 3:
			AdminMenuView()
		}
	}
}

// View

// LoginView displays the login screen interface.
// It renders a navigation breadcrumb showing the current location
// and prints a formatted LOGIN title header.
func LoginView() {
	PrintBreadcrumbs([255]string{"Login"}, 1)
	PrintTitle("LOGIN")
}

// RegisterView displays the registration screen interface.
// It renders a navigation breadcrumb showing the current location
// and prints a formatted REGISTER title header.
func RegisterView() {
	PrintBreadcrumbs([255]string{"Register"}, 1)
	PrintTitle("REGISTER")
}

// AdminMenuView displays the administrator menu interface.
// It renders a navigation breadcrumb showing the current location
// and prints a formatted ADMIN MENU title header.
func AdminMenuView() {
	PrintBreadcrumbs([255]string{"Admin Menu"}, 1)
	PrintTitle("ADMIN MENU")
}

// Helper

// PrintTitle formats and displays the given title text within a bordered box.
// If the title is longer than the predefined width (38 characters), it will be
// split into multiple lines, breaking at spaces when possible.
//
// The function centers each line of text and adds decorative borders
// around the entire title.
func PrintTitle(title string) {
	const width int = 38
	var start, currentPos, lastSpace int

	if len(title) <= width {
		printBorder()
		printCenteredText(title, width)
		printBorder()
		return
	}

	printBorder()

	start = 0
	currentPos = 0
	lastSpace = -1

	for currentPos < len(title) {
		if currentPos-start >= width {
			if lastSpace > start {
				printCenteredText(title[start:lastSpace], width)
				start = lastSpace + 1
				currentPos = start
				lastSpace = -1
			} else {
				printCenteredText(title[start:currentPos], width)
				start = currentPos
			}
		} else if title[currentPos] == ' ' {
			lastSpace = currentPos
			currentPos++
		} else {
			currentPos++
		}
	}

	if start < len(title) {
		printCenteredText(title[start:], width)
	}

	printBorder()
}

// printBorder prints a horizontal border consisting of 42 equal signs.
// Used to create the top and bottom borders of the title box.
func printBorder() {
	fmt.Println("==========================================")
}

// printCenteredText formats and prints a single line of text, centered within
// the specified width. The text is surrounded by "= " on the left and " =" on the right.
//
// For odd-length text, an extra space is added to the right padding to maintain
// proper centering and border alignment.
func printCenteredText(text string, width int) {
	var leftPadding, rightPadding int

	leftPadding = (width - len(text)) / 2
	rightPadding = (width - len(text)) / 2

	if len(text)%2 != 0 {
		rightPadding++
	}

	fmt.Print("= ")
	for i := 0; i < leftPadding; i++ {
		fmt.Print(" ")
	}
	fmt.Print(text)
	for i := 0; i < rightPadding; i++ {
		fmt.Print(" ")
	}
	fmt.Println(" =")
}

// PrintMenu displays a menu of options and captures the user's selection.
// It prints numbered menu items from the provided menu array up to n items,
// prompts the user with menuTitle, and stores their validated choice in answer.
// Returns an error if input scanning fails or the selection is out of range (1-n).
func PrintMenu(menuTitle string, menu [255]string, n int, answer *int) error {
	var input int

	for i := 0; i < n; i++ {
		fmt.Printf("%d. %s\n", i+1, menu[i])
	}

	fmt.Printf("%s (1-%d): ", menuTitle, n)
	_, err := fmt.Scan(&input)
	if err != nil {
		return err
	}

	if input < 1 || input > n {
		return fmt.Errorf("pilihan tidak valid, silakan pilih antara 1 dan %d", n)
	}

	*answer = input

	return nil
}

// PrintBreadcrumbs displays a hierarchical navigation path starting with "Main Menu".
// It prints the first n elements from the links array, separated by " > " characters.
// The last element is printed without a trailing separator.
//
// Parameters:
//   - links: an array of string elements representing navigation levels
//   - n: the number of elements from links to include in the breadcrumb trail
func PrintBreadcrumbs(links [255]string, n int) {
	fmt.Print("Main Menu > ")
	for i := 0; i < n; i++ {
		if i == n-1 {
			fmt.Print(links[i])
		} else {
			fmt.Print(links[i] + " > ")
		}
	}
	fmt.Println()
}
