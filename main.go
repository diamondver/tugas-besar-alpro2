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

// passwordAdmin is the authentication credential for the administrator account.
const passwordAdmin string = "admin123"

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
			LihatSemuaKomentarView(false)
		case 2:
			BuatKomentarView(user, false)
		case 3:
			EditKomentarView(user, false)
		case 4:
			HapusKomentarView(user, false)
		}
	}
}

// LihatSemuaKomentarView displays all comments and provides options for searching,
// sorting, and refreshing the comment list.
func LihatSemuaKomentarView(isAdmin bool) {
	var input int
	var commentsData [NMAX]Comment
	var isFirstRun bool = true

	for {
		if isAdmin {
			PrintBreadcrumbs([255]string{"Admin Menu", "Lihat Komentar", "Lihat Semua Komentar"}, 3)
		} else {
			PrintBreadcrumbs([255]string{"User Menu", "Lihat Semua Komentar"}, 2)
		}
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
			_, err = fmt.Scan(&search)
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
func BuatKomentarView(user User, isAdmin bool) {
	var komentar, kategori string

	if isAdmin {
		PrintBreadcrumbs([255]string{"Admin Menu", "Lihat Komentar", "Buat Komentar"}, 3)
	} else {
		PrintBreadcrumbs([255]string{"User Menu", "Buat Komentar"}, 2)
	}
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
func EditKomentarView(user User, isAdmin bool) {
	var commentsData [NMAX]Comment

	if isAdmin {
		PrintBreadcrumbs([255]string{"Admin Menu", "Lihat Komentar", "Edit Komentar"}, 3)
	} else {
		PrintBreadcrumbs([255]string{"User Menu", "Edit Komentar"}, 2)
	}
	PrintTitle("EDIT KOMENTAR")

	err := GetComments(&commentsData)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Scanln()
		return
	}

	var n int = 1
	for i := 0; i < nComment; i++ {
		if commentsData[i].userId == user.id && !isAdmin {
			fmt.Printf("%d. ID: %d, Komentar: %s, Kategori: %s\n", n, commentsData[i].id, commentsData[i].komentar, commentsData[i].kategori)
			n++
		} else if isAdmin {
			fmt.Printf("%d. ID: %d, User ID: %d, Komentar: %s, Kategori: %s\n", n, commentsData[i].id, commentsData[i].userId, commentsData[i].komentar, commentsData[i].kategori)
			n++
		}
	}

	var inputId int
	var commentToEdit Comment
	var komentar, kategori string

	for {
		fmt.Print("ID: ")
		_, err := fmt.Scan(&inputId)
		if err != nil {
			fmt.Println(err.Error())
		} else if err := FindCommentById(inputId, &commentToEdit); err != nil {
			fmt.Println(err.Error())
		} else if commentToEdit.userId != user.id && !isAdmin {
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
// and prints a formatted HAPUS KOMENTAR (Delete Comment) title header.
func HapusKomentarView(user User, isAdmin bool) {
	var commentsData [NMAX]Comment

	if isAdmin {
		PrintBreadcrumbs([255]string{"Admin Menu", "Lihat Komentar", "Hapus Komentar"}, 3)
	} else {
		PrintBreadcrumbs([255]string{"User Menu", "Hapus Komentar"}, 2)
	}
	PrintTitle("HAPUS KOMENTAR")

	err := GetComments(&commentsData)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Scanln()
		return
	}

	var n int = 1
	for i := 0; i < nComment; i++ {
		if commentsData[i].userId == user.id && !isAdmin {
			fmt.Printf("%d. ID: %d, Komentar: %s, Kategori: %s\n", n, commentsData[i].id, commentsData[i].komentar, commentsData[i].kategori)
			n++
		} else if isAdmin {
			fmt.Printf("%d. ID: %d, User ID: %d, Komentar: %s, Kategori: %s\n", n, commentsData[i].id, commentsData[i].userId, commentsData[i].komentar, commentsData[i].kategori)
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
		} else if commentToDelete.userId != user.id && !isAdmin {
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
func RegisterView() {
	var username, password string

	PrintBreadcrumbs([255]string{"Register"}, 1)
	PrintTitle("REGISTER")

	for {
		if err := RegisterForm(&username, &password, false); err != nil {
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

// AdminMenuView displays the administrator menu interface with authentication.
// It renders a navigation breadcrumb showing the current location in the application
// and prints a formatted ADMIN MENU title header.
func AdminMenuView() {
	var password string
	var isLoggedIn bool = false
	var input int

	for {
		PrintBreadcrumbs([255]string{"Admin Menu"}, 1)
		PrintTitle("ADMIN MENU")

		if passwordAdmin != "" && !isLoggedIn {
			fmt.Print("Masukkan Password Admin: ")
			_, err := fmt.Scan(&password)
			if err != nil {
				fmt.Println("Terjadi kesalahan saat membaca input:", err.Error())
				return
			}

			if password != passwordAdmin {
				fmt.Println("Password salah!")
				if err := ConfirmForm("Apakah Anda ingin mencoba lagi?"); err != nil {
					return
				}
				continue
			}

			isLoggedIn = true
		}

		err := PrintMenu("Pilih Menu", [255]string{"Lihat Komentar", "Lihat User", "Lihat Grafik", "Keluar"}, 4, &input)
		if err != nil {
			return
		}

		if input == 4 {
			break
		}

		switch input {
		case 1:
			LihatKomentarAdminView()
		case 2:
			LihatUserView()
		case 3:
			LihatGrafikView()
		}
	}
}

// LihatKomentarAdminView displays the comment management interface for administrators.
// It renders a navigation breadcrumb showing the current location in the application
// and prints a formatted LIHAT KOMENTAR (View Comments) title header.
func LihatKomentarAdminView() {
	var input int

	for {
		PrintBreadcrumbs([255]string{"Admin Menu", "Lihat Komentar"}, 2)
		PrintTitle("LIHAT KOMENTAR")

		err := PrintMenu("Pilih Menu", [255]string{"Lihat Semua Komentar", "Buat Komentar", "Ubah Komentar", "Delete Komentar", "Kembali"}, 5, &input)
		if err != nil {
			return
		}

		if input == 5 {
			break
		}

		switch input {
		case 1:
			LihatSemuaKomentarView(true)
		case 2:
			BuatKomentarView(User{}, true)
		case 3:
			EditKomentarView(User{}, true)
		case 4:
			HapusKomentarView(User{}, true)
		}
	}
}

// LihatUserView displays the user management interface for administrators.
// It renders a navigation breadcrumb showing the current location in the application
// and prints a formatted LIHAT USER (View Users) title header.
func LihatUserView() {
	var input int
	for {
		PrintBreadcrumbs([255]string{"Admin Menu", "Lihat User"}, 2)
		PrintTitle("LIHAT USER")

		err := PrintMenu("Pilih Menu", [255]string{"Lihat Semua User", "Buat User", "Ubah User", "Hapus User", "Kembali"}, 5, &input)
		if err != nil {
			return
		}

		if input == 5 {
			break
		}

		switch input {
		case 1:
			LihatSemuaUserAdminView()
		case 2:
			BuatUserAdminView()
		case 3:
			EditUserAdminView()
		case 4:
			HapusUserAdminView()
		}
	}
}

// LihatSemuaUserAdminView displays all users in the system for administrative review.
// It renders a navigation breadcrumb showing the current location in the application
// and prints a formatted LIHAT SEMUA USER (View All Users) title header.
func LihatSemuaUserAdminView() {
	var input int
	var usersData [NMAX]User
	var isFirstRun bool = true

	for {
		PrintBreadcrumbs([255]string{"Admin Menu", "Lihat User", "Lihat Semua User"}, 3)
		PrintTitle("LIHAT SEMUA USER")

		if nUser == 0 {
			fmt.Println("Tidak ada user yang terdaftar.")
			if err := ConfirmForm("Apakah Anda ingin kembali?"); err != nil {
				return
			}
			break
		}

		if isFirstRun {
			err := GetUsers(&usersData)
			if err != nil {
				fmt.Println(err.Error())
				fmt.Scanln()
				return
			}
		}

		var n int = 1
		for i := 0; i < nUser; i++ {
			if usersData[i].id != 0 {
				fmt.Printf("%d. ID: %d, Username: %s\n", n, usersData[i].id, usersData[i].username)
				n++
			}
		}

		err := PrintMenu("Pilih Menu", [255]string{"Cari User", "Sortir User", "Refresh", "Kembali"}, 4, &input)
		if err != nil {
			return
		}

		isFirstRun = false

		if input == 4 {
			break
		}

		switch input {
		case 1:
			var search string
			fmt.Print("Masukkan kata kunci untuk mencari user: ")
			_, err = fmt.Scan(&search)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			err = GetUsersSearch(&usersData, search)
			if err != nil {
				fmt.Println(err.Error())
				fmt.Scanln()
				continue
			}
		case 2:
			err = GetUsersSort(&usersData)
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

// BuatUserAdminView displays the user creation interface for administrators.
// It renders a navigation breadcrumb showing the current location in the application
// and prints a formatted BUAT USER (Create User) title header.
func BuatUserAdminView() {
	PrintBreadcrumbs([255]string{"Admin Menu", "Lihat User", "Buat User"}, 3)
	PrintTitle("BUAT USER")

	var username, password string
	for {
		if err := RegisterForm(&username, &password, false); err != nil {
			fmt.Println(err.Error())
		} else if err := CreateUser(username, password); err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("User berhasil dibuat!")
			break
		}

		if err := ConfirmForm("Apakah Anda ingin mencoba lagi?"); err != nil {
			break
		}
	}
}

// EditUserAdminView displays the user editing interface for administrators.
// It renders a navigation breadcrumb showing the current location in the application
// and prints a formatted UBAH USER (Edit User) title header.
func EditUserAdminView() {
	var usersData [NMAX]User

	PrintBreadcrumbs([255]string{"Admin Menu", "Lihat User", "Ubah User"}, 3)
	PrintTitle("UBAH USER")

	err := GetUsers(&usersData)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Scanln()
		return
	}

	var n int = 1
	for i := 0; i < nUser; i++ {
		fmt.Printf("%d. ID: %d, Username: %s\n", n, usersData[i].id, usersData[i].username)
		n++
	}

	var inputId int
	var userToEdit User
	var username, password string

	for {
		fmt.Print("ID: ")
		_, err := fmt.Scan(&inputId)
		if err != nil {
			fmt.Println(err.Error())
		} else if err := FindUserById(inputId, &userToEdit); err != nil {
			fmt.Println(err.Error())
		} else if err := RegisterForm(&username, &password, true); err != nil {
			fmt.Println(err.Error())
		} else if err := EditUser(username, password, userToEdit.id); err != nil {
			fmt.Println(err.Error())
		} else {
			break
		}

		if err := ConfirmForm("Apakah Anda ingin mencoba lagi?"); err != nil {
			return
		}
	}
}

// HapusUserAdminView displays the user deletion interface for administrators.
// It renders a navigation breadcrumb showing the current location in the application
// and prints a formatted title header.
func HapusUserAdminView() {
	var usersData [NMAX]User

	PrintBreadcrumbs([255]string{"Admin Menu", "Lihat User", "Hapus User"}, 3)
	PrintTitle("HAPUS USER")

	err := GetUsers(&usersData)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Scanln()
		return
	}

	var n int = 1
	for i := 0; i < nUser; i++ {
		fmt.Printf("%d. ID: %d, Username: %s\n", n, usersData[i].id, usersData[i].username)
		n++
	}

	var inputId int
	var userToDelete User

	for {
		fmt.Print("ID: ")
		_, err := fmt.Scan(&inputId)
		if err != nil {
			fmt.Println(err.Error())
		} else if err := FindUserById(inputId, &userToDelete); err != nil {
			fmt.Println(err.Error())
		} else if err := DeleteUser(userToDelete.id); err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("User berhasil dihapus!")
			break
		}

		if err := ConfirmForm("Apakah Anda ingin mencoba lagi?"); err != nil {
			break
		}
	}
}

// LihatGrafikView displays statistics and analytics for the sentiment analysis system.
// It renders a navigation breadcrumb showing the current location in the application
// and prints a formatted LIHAT GRAFIK (View Graph/Statistics) title header.
func LihatGrafikView() {
	PrintBreadcrumbs([255]string{"Admin Menu", "Lihat Grafik"}, 2)
	PrintTitle("LIHAT GRAFIK")
	fmt.Println("Jumlah User:", nUser)
	fmt.Println("Jumlah Komentar:", nComment)
	fmt.Println("Jumlah Komentar Positif:", CountCommentsByCategory("positif"))
	fmt.Println("Jumlah Komentar Netral:", CountCommentsByCategory("netral"))
	fmt.Println("Jumlah Komentar Negatif:", CountCommentsByCategory("negatif"))
	fmt.Scan()
}

// Form

// LoginForm prompts the user to enter their username and password.
// It reads the inputs from standard input and validates that neither field is empty.
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
func RegisterForm(username, password *string, editMode bool) error {
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

	if !editMode && (*username == "" || *password == "" || confirmPassword == "") {
		return fmt.Errorf("username, password, dan konfirmasi password tidak boleh kosong")
	}

	if *password != confirmPassword {
		return fmt.Errorf("password dan konfirmasi password tidak cocok")
	}

	return nil
}

// KomentarForm prompts the user to enter comment text and a sentiment category.
// It reads the inputs from standard input and validates them according to application rules.
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

// GetUsers retrieves all registered users from the system and copies them to the provided array.
func GetUsers(usersInput *[NMAX]User) error {
	if nUser == 0 {
		return fmt.Errorf("tidak ada pengguna yang terdaftar")
	}

	*usersInput = users

	return nil
}

// GetUsersSearch searches for users whose usernames contain the specified substring.
// It performs a case-insensitive search by converting both the search term and
// usernames to lowercase before comparison.
func GetUsersSearch(usersInput *[NMAX]User, search string) error {
	var matchCount int
	var isMatch bool

	if nUser == 0 {
		return fmt.Errorf("tidak ada pengguna yang terdaftar")
	}

	var tempUsers [NMAX]User
	matchCount = 0

	search = toLower(search)

	for i := 0; i < nUser; i++ {
		userLower := toLower(users[i].username)
		isMatch = false

		for j := 0; j <= len(userLower)-len(search); j++ {
			isMatch = true

			for k := 0; k < len(search); k++ {
				if userLower[j+k] != search[k] {
					isMatch = false
					break
				}
			}

			if isMatch {
				tempUsers[matchCount] = users[i]
				matchCount++
				break
			}
		}
	}

	if matchCount == 0 {
		return fmt.Errorf("tidak ada username yang sesuai dengan pencarian")
	}

	for i := 0; i < NMAX; i++ {
		if i < matchCount {
			usersInput[i] = tempUsers[i]
		} else {
			usersInput[i] = User{}
		}
	}

	return nil
}

// GetUsersSort sorts the users array by ID and stores the result in the provided usersInput.
// It prompts the user to choose between ascending or descending sort order through a menu interface.
// Selection sort is used for ascending order, and insertion sort is used for descending order.
func GetUsersSort(usersInput *[NMAX]User) error {
	var input int
	var key User

	if nUser == 0 {
		return fmt.Errorf("tidak ada user yang tersedia")
	}

	err := PrintMenu("Pilih Urutan", [255]string{"Ascending (A-Z)", "Descending (Z-A)"}, 2, &input)
	if err != nil {
		return err
	}

	*usersInput = users

	if input == 1 {
		for i := 0; i < nUser-1; i++ {
			minIdx := i
			for j := i + 1; j < nUser; j++ {
				if usersInput[j].id < usersInput[minIdx].id {
					minIdx = j
				}
			}

			usersInput[i], usersInput[minIdx] = usersInput[minIdx], usersInput[i]
		}
	} else {
		for i := 1; i < nUser; i++ {
			key = usersInput[i]
			j := i - 1

			for j >= 0 && usersInput[j].id < key.id {
				usersInput[j+1] = usersInput[j]
				j--
			}

			usersInput[j+1] = key
		}
	}

	return nil
}

// FindUserByUsername searches for a user with the specified username in the users array.
// If found, it copies the user data to the provided user pointer.
func FindUserByUsername(username string, user *User) error {
	for i := 0; i < nUser; i++ {
		if users[i].username == username {
			*user = users[i]
			return nil
		}
	}
	return fmt.Errorf("pengguna dengan username '%s' tidak ditemukan", username)
}

// FindUserById searches for a user with the specified ID using binary search algorithm.
// It assumes that the global users array is sorted by ID in ascending order.
// If found, it copies the user data to the provided user pointer.
func FindUserById(userId int, user *User) error {
	var left, right, mid int

	left = 0
	right = nUser - 1

	for left <= right {
		mid = (left + right) / 2

		if users[mid].id == userId {
			*user = users[mid]
			return nil
		}

		if users[mid].id < userId {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return fmt.Errorf("pengguna dengan ID %d tidak ditemukan", userId)
}

// CreateUser creates a new user with the specified username and password.
// It adds the user to the users array and assigns a unique ID.
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

// EditUser updates a user's username and/or password using binary search to find the user.
// It assumes that the users array is sorted by ID in ascending order.
func EditUser(username, password string, userId int) error {
	var left, right, mid int

	left = 0
	right = nUser - 1

	for left <= right {
		mid = (left + right) / 2

		if users[mid].id == userId {
			if username != "" {
				users[mid].username = username
			}
			if password != "" {
				users[mid].password = password
			}
			return nil
		}

		if users[mid].id < userId {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return fmt.Errorf("pengguna dengan ID %d tidak ditemukan", userId)
}

// DeleteUser removes a user with the specified ID from the users array using binary search.
// It assumes that the users array is sorted by ID in ascending order.
// Once found, it deletes the user by shifting all subsequent elements one
// position to the left to fill the gap, then decrements the user counter.
func DeleteUser(userId int) error {
	var left, right, mid int
	left = 0
	right = nUser - 1

	for left <= right {
		mid = (left + right) / 2

		if users[mid].id == userId {
			for j := mid; j < nUser-1; j++ {
				users[j] = users[j+1]
			}
			users[nUser-1] = User{}
			nUser--
			return nil
		}

		if users[mid].id < userId {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return fmt.Errorf("pengguna dengan ID %d tidak ditemukan", userId)
}

// CreateComment adds a new comment to the system with the specified content and category.
// It assigns a unique ID to the comment and associates it with the given user.
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

// CountCommentsByCategory counts the number of comments that match the specified category.
// It iterates through all comments in the global comments array and increments a counter
// each time it finds a comment with a matching kategori field.
func CountCommentsByCategory(category string) int {
	var count int

	for i := 0; i < nComment; i++ {
		if comments[i].kategori == category {
			count++
		}
	}

	return count
}

// GetComments retrieves all available comments from the system and copies them to the provided array.
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
// Selection sort is used for ascending order, and insertion sort is used for descending order.
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

// FindCommentById searches for a comment with the specified ID using binary search.
// It assumes that the comments array is sorted by ID in ascending order.
// If found, it copies the comment data to the provided comment pointer.
func FindCommentById(id int, comment *Comment) error {
	var left, right, mid int

	left = 0
	right = nComment - 1

	for left <= right {
		mid = (left + right) / 2

		if comments[mid].id == id {
			*comment = comments[mid]
			return nil
		}

		if comments[mid].id < id {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return fmt.Errorf("komentar dengan ID %d tidak ditemukan", id)
}

// EditComment updates an existing comment's text and/or category with the provided values.
// It searches for a comment with the specified ID in the global comments array.
func EditComment(komen, kategori string, id int) error {
	var left, right, mid int

	left = 0
	right = nComment - 1

	for left <= right {
		mid = (left + right) / 2

		if comments[mid].id == id {
			if komen != "" {
				comments[mid].komentar = komen
			}
			if kategori != "" {
				comments[mid].kategori = kategori
			}
			return nil
		}

		if comments[mid].id < id {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return fmt.Errorf("komentar dengan ID %d tidak ditemukan", id)
}

// DeleteComment removes a comment with the specified ID from the comments array using binary search.
// It assumes that the comments array is sorted by ID in ascending order.
// Once found, it deletes the comment by shifting all subsequent elements one
// position to the left to fill the gap, then decrements the comment counter.
func DeleteComment(id int) error {
	var left, right, mid int

	left = 0
	right = nComment - 1

	for left <= right {
		mid = (left + right) / 2

		if comments[mid].id == id {
			for j := mid; j < nComment-1; j++ {
				comments[j] = comments[j+1]
			}
			comments[nComment-1] = Comment{}
			nComment--
			return nil
		}

		if comments[mid].id < id {
			left = mid + 1
		} else {
			right = mid - 1
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
// The function first displays all menu options as a numbered list. It then enters a loop
// that prompts the user for input until a valid selection is made. If the user enters
// an invalid option (outside the range 1-n), an error message is displayed and the user
// is prompted again. This continues until a valid selection is made or an input error occurs.
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
func toLower(s string) string {
	for i := 0; i < len(s); i++ {
		if s[i] >= 'A' && s[i] <= 'Z' {
			s = s[:i] + string(s[i]+32) + s[i+1:]
		}
	}

	return s
}
