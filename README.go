package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

// User struct to map API response
type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Website  string `json:"website"`
	Address  struct {
		Street  string `json:"street"`
		City    string `json:"city"`
		Zipcode string `json:"zipcode"`
	} `json:"address"`
}

func main() {
	fmt.Println("🚀 Go CLI API Fetcher - Beginner's Toolkit")
	fmt.Println("==========================================")
	
	if len(os.Args) > 1 {
		userID, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Printf("Error: Please provide a valid user ID (1-10)\n")
			os.Exit(1)
		}
		
		if userID < 1 || userID > 10 {
			fmt.Printf("Error: User ID must be between 1 and 10\n")
			os.Exit(1)
		}
		
		user, err := fetchUser(userID)
		if err != nil {
			fmt.Printf("Error fetching user: %v\n", err)
			os.Exit(1)
		}
		displayUser(user)
	} else {
		// Default behavior - fetch multiple users
		fmt.Println("Fetching multiple users...")
		fetchMultipleUsers()
	}
}

// fetchUser gets user data from JSONPlaceholder API
func fetchUser(userID int) (*User, error) {
	url := fmt.Sprintf("https://jsonplaceholder.typicode.com/users/%d", userID)
	
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("network error: %v", err)
	}
	defer resp.Body.Close()
	
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("API returned status: %s", resp.Status)
	}
	
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %v", err)
	}
	
	var user User
	err = json.Unmarshal(body, &user)
	if err != nil {
		return nil, fmt.Errorf("error parsing JSON: %v", err)
	}
	
	return &user, nil
}

// fetchMultipleUsers demonstrates fetching multiple users
func fetchMultipleUsers() {
	userIDs := []int{1, 2, 3}
	
	for _, id := range userIDs {
		user, err := fetchUser(id)
		if err != nil {
			fmt.Printf("Failed to fetch user %d: %v\n", id, err)
			continue
		}
		fmt.Printf("\n--- User %d ---\n", id)
		fmt.Printf("Name: %s\n", user.Name)
		fmt.Printf("Email: %s\n", user.Email)
		fmt.Printf("City: %s\n", user.Address.City)
	}
}

// displayUser prints user information in a formatted way
func displayUser(user *User) {
	fmt.Println("\n" + "="*50)
	fmt.Println("👤 USER INFORMATION")
	fmt.Println("="*50)
	fmt.Printf("ID: %d\n", user.ID)
	fmt.Printf("Name: %s\n", user.Name)
	fmt.Printf("Username: %s\n", user.Username)
	fmt.Printf("Email: %s\n", user.Email)
	fmt.Printf("Phone: %s\n", user.Phone)
	fmt.Printf("Website: %s\n", user.Website)
	fmt.Printf("Address: %s, %s, %s\n", user.Address.Street, user.Address.City, user.Address.Zipcode)
	fmt.Println("="*50)
}
