package main

import (
	"fmt"
)

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
	userId   int    // Identifier of the user who created the comment
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
			UserMenuView(*user)
			break
		}

		if err := ConfirmForm("Apakah Anda ingin mencoba lagi?"); err != nil {
			break
		}
	}
}

// UserMenuView displays and handles the main user menu interface.
// It presents a navigation breadcrumb and menu options for the authenticated user.
//
// The function implements a menu loop that:
//   - Displays the user menu with 5 options (view comments, create/edit/delete comments, exit)
//   - Captures user input and validates it
//   - Routes to the appropriate view function based on selection
//   - Continues to display the menu until the user selects "Keluar" (exit)
//
// The function handles any input errors by returning immediately when they occur.
//
// Parameters:
//   - user: User struct containing the authenticated user's information
//     which is passed to child views that require user context
//
// No return value - function exits when user selects option 5 or when an input error occurs.
func UserMenuView(user User) {
	var input int

	for {
		PrintBreadcrumbs([255]string{"User Menu"}, 1)
		PrintTitle("USER MENU")

		err := PrintMenu("Pilih Menu", [255]string{"Lihat Semua Komentar", "Buat Komentar", "Edit Komentar", "Hapus Komentar", "Keluar"}, 5, &input)
		if err != nil {
			return
		}

		if input == 5 {
			break
		}

		switch input {
		case 1:
			LihatSemuaKomentarView()
		case 2:
			BuatKomentarView(user)
		case 3:
			EditKomentarView(user)
		case 4:
			HapusKomentarView(user)
		}
	}
}

// LihatSemuaKomentarView displays all comments and provides options for searching,
// sorting, and refreshing the comment list.
//
// This function implements a view that:
//   - Displays a navigation breadcrumb showing the current location
//   - Shows a formatted title header "LIHAT SEMUA KOMENTAR"
//   - Lists all available comments with their details (ID, User ID, content, category)
//   - Provides a menu with 4 options:
//     1. Search comments by keyword
//     2. Sort comments by ID (ascending or descending)
//     3. Refresh the comment list
//     4. Return to the previous menu
//
// The function handles all error conditions, displaying appropriate messages to the user.
// When errors occur in comment operations, the function pauses with fmt.Scanln() to allow
// the user to read the error message before continuing.
//
// The function maintains state between menu selections through the isFirstRun boolean,
// which determines whether to reload comments from the data store or continue using
// the current filtered/sorted view.
//
// No parameters or return values - the function runs until the user selects "Kembali" (back)
// or an unrecoverable error occurs during menu display.
func LihatSemuaKomentarView() {
	var input int
	var commentsData [NMAX]Comment
	var isFirstRun bool = true

	for {
		PrintBreadcrumbs([255]string{"User Menu", "Lihat Semua Komentar"}, 2)
		PrintTitle("LIHAT SEMUA KOMENTAR")

		if isFirstRun {
			err := GetComments(&commentsData)
			if err != nil {
				fmt.Println(err.Error())
				fmt.Scanln()
				return
			}
		}

		var n int = 1
		for i := 0; i < nComment; i++ {
			if commentsData[i].id != 0 {
				fmt.Printf("%d. ID: %d, User ID: %d, Komentar: %s, Kategori: %s\n", n, commentsData[i].id, commentsData[i].userId, commentsData[i].komentar, commentsData[i].kategori)
				n++
			}
		}

		err := PrintMenu("Pilih Menu", [255]string{"Cari Komentar", "Sortir Komentar", "Refresh", "Kembali"}, 4, &input)
		if err != nil {
			return
		}

		if input == 4 {
			break
		}

		isFirstRun = false

		switch input {
		case 1:
			var search string
			fmt.Print("Masukkan kata kunci untuk mencari komentar: ")
			_, err = fmt.Scanln(&search)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			err = GetCommentsSearch(&commentsData, search)
			if err != nil {
				fmt.Println(err.Error())
				fmt.Scanln()
				continue
			}
		case 2:
			err = GetCommentsSort(&commentsData)
			if err != nil {
				fmt.Println(err.Error())
				fmt.Scanln()
				continue
			}
		case 3:
			isFirstRun = true
		}
	}
}

// BuatKomentarView displays the comment creation interface and handles the process of creating a new comment.
// It renders a navigation breadcrumb showing the current location in the application
// and prints a formatted BUAT KOMENTAR (Create Comment) title header.
//
// The function implements a comment creation loop that:
//   - Prompts for comment text and category using KomentarForm
//   - Attempts to create a new comment with the provided data
//   - Provides appropriate error messages on failure
//   - Offers to retry or exit on failed attempts
//
// Parameters:
//   - user: User struct containing the authenticated user's information,
//     which is used to associate the created comment with the user
//
// The function doesn't return any values. Upon successful comment creation,
// a confirmation message is displayed and the function returns.
func BuatKomentarView(user User) {
	var komentar, kategori string

	PrintBreadcrumbs([255]string{"User Menu", "Buat Komentar"}, 2)
	PrintTitle("BUAT KOMENTAR")

	for {
		if err := KomentarForm(&komentar, &kategori, false); err != nil {
			fmt.Println(err.Error())
		} else if err := CreateComment(user, komentar, kategori); err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("Komentar berhasil dibuat!")
			break
		}

		if err := ConfirmForm("Apakah Anda ingin mencoba lagi?"); err != nil {
			break
		}
	}
}

// EditKomentarView displays the comment editing interface and handles the process of modifying existing comments.
// It renders a navigation breadcrumb showing the current location in the application
// and prints a formatted EDIT KOMENTAR (Edit Comment) title header.
//
// The function implements a workflow that:
//   - Retrieves and displays only the comments belonging to the current user
//   - Prompts the user to select a comment by ID
//   - Verifies the user has permission to edit the selected comment
//   - Uses KomentarForm to collect the modified comment text and category
//   - Attempts to update the comment with the new information
//   - Provides appropriate error messages on failure
//   - Offers to retry or exit on failed attempts
//
// The function handles multiple error conditions including:
//   - Failure to retrieve comments
//   - Invalid input ID
//   - Comment not found
//   - Permission denied (attempting to edit another user's comment)
//   - Form validation errors
//   - Comment update failures
//
// Parameters:
//   - user: User struct containing the authenticated user's information,
//     which is used to filter comments and verify edit permissions
//
// The function doesn't return any values. Upon successful comment editing,
// a confirmation message is displayed and the function returns.
func EditKomentarView(user User) {
	var commentsData [NMAX]Comment

	PrintBreadcrumbs([255]string{"User Menu", "Edit Komentar"}, 2)
	PrintTitle("EDIT KOMENTAR")

	err := GetComments(&commentsData)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Scanln()
		return
	}

	var n int = 1
	for i := 0; i < nComment; i++ {
		if commentsData[i].userId == user.id {
			fmt.Printf("%d. ID: %d, Komentar: %s, Kategori: %s\n", n, commentsData[i].id, commentsData[i].komentar, commentsData[i].kategori)
			n++
		}
	}

	var inputId int
	var commentToEdit Comment
	var komentar, kategori string

	for {
		fmt.Print("ID: ")
		_, err := fmt.Scanln(&inputId)
		if err != nil {
			fmt.Println(err.Error())
		} else if err := FindCommentById(inputId, &commentToEdit); err != nil {
			fmt.Println(err.Error())
		} else if commentToEdit.userId != user.id {
			fmt.Println("Anda tidak memiliki izin untuk mengedit komentar ini.")
		} else if err := KomentarForm(&komentar, &kategori, true); err != nil {
			fmt.Println(err.Error())
		} else if err := EditComment(komentar, kategori, commentToEdit.id); err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("Komentar berhasil diubah!")
			break
		}

		if err := ConfirmForm("Apakah Anda ingin mencoba lagi?"); err != nil {
			break
		}
	}
}

// HapusKomentarView displays the comment deletion interface and handles the process of removing existing comments.
// It renders a navigation breadcrumb showing the current location in the application
// and prints a formatted title header.
//
// The function implements a workflow that:
//   - Retrieves and displays only the comments belonging to the current user
//   - Prompts the user to select a comment by ID
//   - Verifies the user has permission to delete the selected comment
//   - Attempts to delete the comment from the system
//   - Provides appropriate error messages on failure
//   - Offers to retry or exit on failed attempts
//
// The function handles multiple error conditions including:
//   - Failure to retrieve comments
//   - Invalid input ID
//   - Comment not found
//   - Permission denied (attempting to delete another user's comment)
//   - Comment deletion failures
//
// Parameters:
//   - user: User struct containing the authenticated user's information,
//     which is used to filter comments and verify deletion permissions
//
// The function doesn't return any values. Upon successful comment deletion,
// a confirmation message is displayed and the function returns.
func HapusKomentarView(user User) {
	var commentsData [NMAX]Comment

	PrintBreadcrumbs([255]string{"User Menu", "Hapus Komentar"}, 2)
	PrintTitle("EDIT KOMENTAR")

	err := GetComments(&commentsData)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Scanln()
		return
	}

	var n int = 1
	for i := 0; i < nComment; i++ {
		if commentsData[i].userId == user.id {
			fmt.Printf("%d. ID: %d, Komentar: %s, Kategori: %s\n", n, commentsData[i].id, commentsData[i].komentar, commentsData[i].kategori)
			n++
		}
	}

	var inputId int
	var commentToDelete Comment

	for {
		fmt.Print("ID: ")
		_, err := fmt.Scan(&inputId)
		if err != nil {
			fmt.Println(err.Error())
		} else if err := FindCommentById(inputId, &commentToDelete); err != nil {
			fmt.Println(err.Error())
		} else if commentToDelete.userId != user.id {
			fmt.Println("Anda tidak memiliki izin untuk menghapus komentar ini.")
		} else if err := DeleteComment(commentToDelete.id); err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("Komentar berhasil dihapus!")
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

// KomentarForm prompts the user to enter comment text and a sentiment category.
// It reads the inputs from standard input and validates them according to application rules.
//
// Parameters:
//   - komentar: pointer to a string where the entered comment text will be stored
//   - kategori: pointer to a string where the entered category will be stored
//   - editMode: boolean flag that modifies validation behavior; when true, empty inputs
//     are allowed (for partial updates); when false, both fields are required
//
// The function enforces that categories, when provided, must be one of three valid values:
// "positif", "negatif", or "netral".
//
// Returns:
//   - error: nil if the input was successfully read and validated, or an error if:
//     1. scanning the input failed
//     2. required fields are empty (in non-edit mode)
//     3. an invalid category value was provided
func KomentarForm(komentar, kategori *string, editMode bool) error {
	fmt.Print("Masukkan Komentar: ")
	_, err := fmt.Scan(komentar)
	if err != nil {
		return err
	}

	fmt.Print("Masukkan Kategori: ")
	_, err = fmt.Scan(kategori)
	if err != nil {
		return err
	}

	if !editMode && (*komentar == "" || *kategori == "") {
		return fmt.Errorf("komentar dan kategori tidak boleh kosong")
	}

	if *kategori != "" && *kategori != "positif" && *kategori != "negatif" && *kategori != "netral" {
		return fmt.Errorf("kategori harus 'positif', 'negatif', atau 'netral'")
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

// CreateComment adds a new comment to the system with the specified content and category.
// It assigns a unique ID to the comment and associates it with the given user.
//
// Parameters:
//   - user: the User struct of the comment author, providing the user ID for association
//   - komentar: string containing the comment text content
//   - kategori: string specifying the sentiment category of the comment
//     (must be one of: "positif", "negatif", or "netral")
//
// Returns:
//   - error: nil if the comment was successfully created, or an error if
//     the maximum number of comments (NMAX) has been reached
//
// The function automatically increments both the comment counter (nComment)
// and the unique comment ID generator (idComment) after successful creation.
func CreateComment(user User, komentar, kategori string) error {
	if nComment >= NMAX {
		return fmt.Errorf("jumlah komentar sudah mencapai batas maksimum")
	}

	comments[nComment] = Comment{
		id:       idComment,
		userId:   user.id,
		komentar: komentar,
		kategori: kategori,
	}
	nComment++
	idComment++
	return nil
}

// GetComments retrieves all available comments from the system and copies them to the provided array.
//
// Parameters:
//   - commentsInput: pointer to an array of Comment structs where the comment data will be copied.
//     This array must have at least NMAX capacity.
//
// Returns:
//   - error: nil if comments were successfully retrieved, or an error if
//     no comments are available in the system (nComment == 0)
//
// The function performs a direct copy of the global comments array to the provided
// parameter without any filtering or sorting.
func GetComments(commentsInput *[NMAX]Comment) error {
	if nComment == 0 {
		return fmt.Errorf("tidak ada komentar yang tersedia")
	}

	*commentsInput = comments

	return nil
}

// GetCommentsSearch searches through all comments for those containing the specified search string.
// It performs a case-insensitive substring search by converting both the search term and
// comment text to lowercase before comparison.
//
// The function implements a manual substring matching algorithm that:
//   - Checks each possible starting position in the comment text
//   - Compares character by character with the search string
//   - Collects all matching comments in a temporary array
//
// Parameters:
//   - commentsInput: pointer to an array of Comment structs where the matching comments
//     will be copied. Non-matching positions will be set to empty Comment structs.
//   - search: the string to search for within comment text
//
// Returns:
//   - error: nil if at least one matching comment was found, or an error if:
//     1. no comments are available in the system (nComment == 0)
//     2. no comments match the search criteria
//
// The function performs a complete replacement of the contents of commentsInput,
// filling it with matching comments followed by empty Comment structs.
func GetCommentsSearch(commentsInput *[NMAX]Comment, search string) error {
	var matchCount int
	var isMatch bool

	if nComment == 0 {
		return fmt.Errorf("tidak ada komentar yang tersedia")
	}

	var tempComments [NMAX]Comment
	matchCount = 0

	search = toLower(search)

	for i := 0; i < nComment; i++ {
		commentLower := toLower(comments[i].komentar)
		isMatch = false

		for j := 0; j <= len(commentLower)-len(search); j++ {
			isMatch = true

			for k := 0; k < len(search); k++ {
				if commentLower[j+k] != search[k] {
					isMatch = false
					break
				}
			}

			if isMatch {
				tempComments[matchCount] = comments[i]
				matchCount++
				break
			}
		}
	}

	if matchCount == 0 {
		return fmt.Errorf("tidak ada komentar yang sesuai dengan pencarian")
	}

	for i := 0; i < NMAX; i++ {
		if i < matchCount {
			commentsInput[i] = tempComments[i]
		} else {
			commentsInput[i] = Comment{}
		}
	}

	return nil
}

// GetCommentsSort sorts the comments array by ID and stores the result in the provided commentsInput.
// It prompts the user to choose between ascending or descending sort order through a menu interface.
//
// The function implements two different sorting algorithms:
//   - Selection sort for ascending order (input == 1)
//   - Insertion sort for descending order (input == 2)
//
// Parameters:
//   - commentsInput: pointer to an array of Comment structs where the sorted comments will be stored.
//     The array must have at least NMAX capacity.
//
// Returns:
//   - error: nil if comments were successfully sorted, or an error if:
//     1. no comments are available in the system (nComment == 0)
//     2. the menu interface returns an error during user input
//
// The function performs a complete replacement of the contents of commentsInput with
// a sorted copy of the global comments array.
func GetCommentsSort(commentsInput *[NMAX]Comment) error {
	var input int
	var key Comment

	if nComment == 0 {
		return fmt.Errorf("tidak ada komentar yang tersedia")
	}

	err := PrintMenu("Pilih Urutan", [255]string{"Ascending (A-Z)", "Descending (Z-A)"}, 2, &input)
	if err != nil {
		return err
	}

	*commentsInput = comments

	if input == 1 {
		for i := 0; i < nComment-1; i++ {
			minIdx := i
			for j := i + 1; j < nComment; j++ {
				if commentsInput[j].id < commentsInput[minIdx].id {
					minIdx = j
				}
			}

			commentsInput[i], commentsInput[minIdx] = commentsInput[minIdx], commentsInput[i]
		}
	} else {
		for i := 1; i < nComment; i++ {
			key = commentsInput[i]
			j := i - 1

			for j >= 0 && commentsInput[j].id < key.id {
				commentsInput[j+1] = commentsInput[j]
				j--
			}

			commentsInput[j+1] = key
		}
	}

	return nil
}

// FindCommentById searches for a comment with the specified ID in the comments array.
// If found, it copies the comment data to the provided comment pointer.
//
// Parameters:
//   - id: the unique identifier of the comment to search for
//   - comment: pointer to a Comment struct where the found comment data will be stored
//
// Returns:
//   - error: nil if a comment with the matching ID is found, otherwise an error
//     with a message indicating the comment with the given ID was not found
func FindCommentById(id int, comment *Comment) error {
	for i := 0; i < nComment; i++ {
		if comments[i].id == id {
			*comment = comments[i]
			return nil
		}
	}

	return fmt.Errorf("komentar dengan ID %d tidak ditemukan", id)
}

// EditComment updates an existing comment's text and/or category with the provided values.
// It searches for a comment with the specified ID in the global comments array.
//
// Parameters:
//   - komen: the new comment text. If empty, the original text is preserved.
//   - kategori: the new comment category. If empty, the original category is preserved.
//   - id: the unique identifier of the comment to edit
//
// Returns:
//   - error: nil if the comment was successfully updated, or an error if
//     no comment with the matching ID was found
//
// The function uses a conditional update approach, where empty string values
// for komen or kategori will not overwrite the existing values.
func EditComment(komen, kategori string, id int) error {
	for i := 0; i < nComment; i++ {
		if comments[i].id == id {
			if komen != "" {
				comments[i].komentar = komen
			}
			if kategori != "" {
				comments[i].kategori = kategori
			}
			return nil
		}
	}
	return fmt.Errorf("komentar dengan ID %d tidak ditemukan", id)
}

// DeleteComment removes a comment with the specified ID from the comments array.
// It uses an in-place deletion approach by shifting all subsequent elements one
// position to the left to fill the gap, then decrements the comment counter.
//
// Parameters:
//   - id: the unique identifier of the comment to delete
//
// Returns:
//   - error: nil if the comment was successfully deleted, or an error if
//     no comment with the matching ID was found
//
// The function modifies the global comments array and nComment counter when successful.
func DeleteComment(id int) error {
	for i := 0; i < nComment; i++ {
		if comments[i].id == id {
			for j := i; j < nComment-1; j++ {
				comments[j] = comments[j+1]
			}
			nComment--
			return nil
		}
	}
	return fmt.Errorf("komentar dengan ID %d tidak ditemukan", id)
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
//
// The function first displays all menu options as a numbered list. It then enters a loop
// that prompts the user for input until a valid selection is made. If the user enters
// an invalid option (outside the range 1-n), an error message is displayed and the user
// is prompted again. This continues until a valid selection is made or an input error occurs.
//
// Parameters:
//   - menuTitle: string that appears in the prompt asking for user selection
//   - menu: array of strings containing the text for each menu option
//   - n: number of menu items to display from the array (must be between 1 and 255)
//   - answer: pointer to an integer where the validated user selection will be stored
//
// Returns:
//   - error: nil if a valid selection was made, or an error if input scanning fails
//
// The function handles input validation internally and will only return when either:
//   - A valid selection is made (stored in answer, returns nil)
//   - A scanning error occurs (returns the error)
func PrintMenu(menuTitle string, menu [255]string, n int, answer *int) error {
	for i := 0; i < n; i++ {
		fmt.Printf("%d. %s\n", i+1, menu[i])
	}

	for {
		var input int
		fmt.Printf("%s (1-%d): ", menuTitle, n)
		_, err := fmt.Scan(&input)

		if err != nil {
			fmt.Print(err.Error())
			continue
		}

		if input >= 1 && input <= n {
			*answer = input
			return nil
		}

		fmt.Printf("Pilihan tidak valid, silakan pilih antara 1 dan %d\n", n)
	}
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

// toLower converts a string to lowercase by changing any uppercase ASCII characters
// (A-Z) to their lowercase equivalents.
//
// This function manually implements case conversion by adding 32 to the ASCII value
// of uppercase letters, which converts them to lowercase. It works only for standard
// ASCII characters and doesn't support Unicode case conversion.
//
// Parameters:
//   - s: the input string to convert to lowercase
//
// Returns:
//   - string: a new string with all uppercase ASCII letters converted to lowercase
func toLower(s string) string {
	for i := 0; i < len(s); i++ {
		if s[i] >= 'A' && s[i] <= 'Z' {
			s = s[:i] + string(s[i]+32) + s[i+1:]
		}
	}

	return s
}
