package main

import "fmt"

// NMAX defines the maximum number of users and comments that can be stored in the application.
// This constant is used to initialize the fixed-size arrays for both users and comments.
const NMAX int = 255

// User represents a user account in the system.
// Each user has a unique identifier, username, and password for authentication.
type User struct {
	id       int    // Unique identifier for the user
	username string // Username for login and display purposes
	password string // Password for authentication
}

// Comment represents a sentiment comment in the system.
// Each comment has a unique identifier, the user ID of the author,
// the comment text, and a category classification.
type Comment struct {
	id       int    // Unique identifier for the comment
	userId   string // Identifier of the user who created the comment
	komentar string // The actual comment text content
	kategori string // The sentiment category or classification of the comment
}

// users is an array storing all registered user accounts.
// The array has a fixed size determined by the NMAX constant.
var users [NMAX]User

// nUser tracks the current number of users stored in the users array.
var nUser int = 0

// idUser is a counter for generating unique user IDs, starting from 1.
var idUser int = 1

// comments is an array storing all sentiment comments.
// The array has a fixed size determined by the NMAX constant.
var comments [NMAX]Comment

// nComment tracks the current number of comments stored in the comments array.
var nComment int = 0

// idComment is a counter for generating unique comment IDs, starting from 1.
var idComment int = 1

func main() {
	var input int
	var userLogin User

	for input != 4 {
		PrintTitle("Selamat datang di Tugas Besar Alpro Aplikasi Analisis Sentimen Kelompok 2")
		err := PrintMenu("Pilih Menu", [255]string{"Login", "Register", "Admin", "Exit"}, 4, &input)
		if err != nil {
			return
		}

		switch input {
		case 1:
			LoginView(&userLogin)
		case 2:
			RegisterView()
		case 3:
			AdminMenuView()
		}
	}
}

// View

// LoginView displays the login screen interface and handles the user authentication process.
// It renders a navigation breadcrumb showing the current location in the application
// and prints a formatted LOGIN title header.
//
// The function implements a login loop that:
//   - Prompts for username and password using LoginForm
//   - Attempts to find the user by username
//   - Verifies the password matches
//   - Provides appropriate error messages on failure
//   - Offers to retry or exit on failed attempts
//
// Parameters:
//   - user: pointer to a User struct that will be populated with the authenticated user's data
//     when login is successful
//
// The function doesn't return any values but modifies the user parameter when authentication succeeds.
func LoginView(user *User) {
	var username, password string

	PrintBreadcrumbs([255]string{"Login"}, 1)
	PrintTitle("LOGIN")

	for {
		if err := LoginForm(&username, &password); err != nil {
			fmt.Println(err.Error())
		} else if err := FindUserByUsername(username, user); err != nil {
			fmt.Println(err.Error())
		} else if user.password != password {
			fmt.Println("Password salah!")
		} else {
			fmt.Println("Login berhasil!")
			break
		}

		if err := ConfirmForm("Apakah Anda ingin mencoba lagi?"); err != nil {
			break
		}
	}
}

// RegisterView displays the registration screen interface and handles the user registration process.
// It renders a navigation breadcrumb showing the current location in the application
// and prints a formatted REGISTER title header.
//
// The function implements a registration loop that:
//   - Prompts for username and password using RegisterForm
//   - Attempts to create a new user account with the provided credentials
//   - Provides appropriate error messages on failure
//   - Offers to retry or exit on failed attempts
//
// The function doesn't take any parameters and doesn't return any values.
// Upon successful registration, a confirmation message is displayed and the function returns.
func RegisterView() {
	var username, password string

	PrintBreadcrumbs([255]string{"Register"}, 1)
	PrintTitle("REGISTER")

	for {
		if err := RegisterForm(&username, &password); err != nil {
			fmt.Println(err.Error())
		} else {
			if err := CreateUser(username, password); err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("Registrasi berhasil!")
				break
			}
		}

		if err := ConfirmForm("Apakah Anda ingin mencoba lagi?"); err != nil {
			break
		}
	}
}

// AdminMenuView displays the administrator menu interface.
// It renders a navigation breadcrumb showing the current location
// and prints a formatted ADMIN MENU title header.
func AdminMenuView() {
	PrintBreadcrumbs([255]string{"Admin Menu"}, 1)
	PrintTitle("ADMIN MENU")
}

// Form

// LoginForm prompts the user to enter their username and password.
// It reads the inputs from standard input and validates that neither field is empty.
//
// Parameters:
//   - username: pointer to a string where the entered username will be stored
//   - password: pointer to a string where the entered password will be stored
//
// Returns:
//   - error: nil if the input was successfully read and validated, or an error if:
//   - scanning the input failed
//   - the username or password is empty
func LoginForm(username, password *string) error {
	fmt.Print("Masukkan Username: ")
	_, err := fmt.Scan(username)
	if err != nil {
		return err
	}

	fmt.Print("Masukkan Password: ")
	_, err = fmt.Scan(password)
	if err != nil {
		return err
	}

	if *username == "" || *password == "" {
		return fmt.Errorf("username dan password tidak boleh kosong")
	}

	return nil
}

// RegisterForm prompts the user to enter a username, password, and password confirmation.
// It reads the inputs from standard input and validates that no field is empty
// and that the password matches the confirmation password.
//
// Parameters:
//   - username: pointer to a string where the entered username will be stored
//   - password: pointer to a string where the entered password will be stored
//
// Returns:
//   - error: nil if the input was successfully read and validated, or an error if:
//   - scanning the input failed
//   - any of the fields is empty
//   - the password and confirmation password don't match
func RegisterForm(username, password *string) error {
	var confirmPassword string

	fmt.Print("Masukkan Username: ")
	_, err := fmt.Scan(username)
	if err != nil {
		return err
	}

	fmt.Print("Masukkan Password: ")
	_, err = fmt.Scan(password)
	if err != nil {
		return err
	}

	fmt.Print("Masukkan Konfirmasi Password: ")
	_, err = fmt.Scan(&confirmPassword)
	if err != nil {
		return err
	}

	if *username == "" || *password == "" || confirmPassword == "" {
		return fmt.Errorf("username, password, dan konfirmasi password tidak boleh kosong")
	}

	if *password != confirmPassword {
		return fmt.Errorf("password dan konfirmasi password tidak cocok")
	}

	return nil
}

// ConfirmForm prompts the user with a yes/no question and returns the result.
// It displays the provided title followed by options for Yes (1) or No (2),
// then reads the user's selection from standard input.
//
// Parameters:
//   - title: the prompt text to display to the user
//
// Returns:
//   - error: nil if the user selects Yes (1), an error with message "cancel"
//     if the user selects No (2), or any error encountered while reading input
//
// The function loops until the user provides valid input (1 or 2) or until
// an input reading error occurs.
func ConfirmForm(title string) error {
	var input int

	for {
		fmt.Printf("%s (1. Ya, 2. Tidak): ", title)
		_, err := fmt.Scan(&input)
		if err != nil {
			return err
		}

		if input == 1 {
			return nil
		} else if input == 2 {
			return fmt.Errorf("cancel")
		} else {
			fmt.Println("Pilihan tidak valid, silakan pilih 1 atau 2.")
		}
	}
}

// Data

// FindUserByUsername searches for a user with the specified username in the users array.
// If found, it copies the user data to the provided user pointer.
//
// Parameters:
//   - username: the username to search for
//   - user: pointer to a User struct where the found user data will be stored
//
// Returns:
//   - error: nil if a user with the matching username is found, otherwise an error
//     with a message indicating the username was not found
func FindUserByUsername(username string, user *User) error {
	for i := 0; i < nUser; i++ {
		if users[i].username == username {
			*user = users[i]
			return nil
		}
	}
	return fmt.Errorf("pengguna dengan username '%s' tidak ditemukan", username)
}

// CreateUser creates a new user with the specified username and password.
// It adds the user to the users array and assigns a unique ID.
//
// Parameters:
//   - username: the username for the new user account
//   - password: the password for the new user account
//
// Returns:
//   - error: nil if the user was successfully created, or an error if:
//   - the maximum number of users has been reached (nUser >= NMAX)
//   - the username already exists in the system
func CreateUser(username, password string) error {
	if nUser >= NMAX {
		return fmt.Errorf("jumlah pengguna sudah mencapai batas maksimum")
	}

	for i := 0; i < nUser; i++ {
		if users[i].username == username {
			return fmt.Errorf("username '%s' sudah terdaftar", username)
		}
	}

	users[nUser] = User{
		id:       idUser,
		username: username,
		password: password,
	}
	nUser++
	idUser++
	return nil
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
