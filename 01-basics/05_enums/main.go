package main

import "fmt"

// Example 1: Basic Enum
type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func (w Weekday) String() string {
	return [...]string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}[w]
}

// Example 2: Enum with Custom Starting Value
type Month int

const (
	January Month = iota + 1
	February
	March
	April
	May
	June
	July
	August
	September
	October
	November
	December
)

func (m Month) String() string {
	months := [...]string{"", "January", "February", "March", "April", "May", "June",
		"July", "August", "September", "October", "November", "December"}
	if m < January || m > December {
		return "Invalid"
	}
	return months[m]
}

// Example 3: Bit Flags
type Permission uint

const (
	Read Permission = 1 << iota
	Write
	Execute
	Admin
)

func (p Permission) String() string {
	var perms []string
	if p&Read != 0 {
		perms = append(perms, "Read")
	}
	if p&Write != 0 {
		perms = append(perms, "Write")
	}
	if p&Execute != 0 {
		perms = append(perms, "Execute")
	}
	if p&Admin != 0 {
		perms = append(perms, "Admin")
	}
	if len(perms) == 0 {
		return "None"
	}
	return fmt.Sprintf("%v", perms)
}

func (p Permission) Has(perm Permission) bool {
	return p&perm != 0
}

// Example 4: File Sizes
type ByteSize int64

const (
	_           = iota // ignore first value by assigning to blank identifier
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
)

func (b ByteSize) String() string {
	switch {
	case b >= PB:
		return fmt.Sprintf("%.2f PB", float64(b)/float64(PB))
	case b >= TB:
		return fmt.Sprintf("%.2f TB", float64(b)/float64(TB))
	case b >= GB:
		return fmt.Sprintf("%.2f GB", float64(b)/float64(GB))
	case b >= MB:
		return fmt.Sprintf("%.2f MB", float64(b)/float64(MB))
	case b >= KB:
		return fmt.Sprintf("%.2f KB", float64(b)/float64(KB))
	default:
		return fmt.Sprintf("%d B", b)
	}
}

// Example 5: Status with Validation
type Status int

const (
	Unknown Status = iota
	Pending
	Active
	Completed
	Failed
	maxStatus // sentinel value for validation
)

func (s Status) String() string {
	return [...]string{"Unknown", "Pending", "Active", "Completed", "Failed"}[s]
}

func (s Status) IsValid() bool {
	return s >= Unknown && s < maxStatus
}

func (s Status) IsTerminal() bool {
	return s == Completed || s == Failed
}

// Example 6: HTTP Methods
type HTTPMethod string

const (
	GET    HTTPMethod = "GET"
	POST   HTTPMethod = "POST"
	PUT    HTTPMethod = "PUT"
	DELETE HTTPMethod = "DELETE"
	PATCH  HTTPMethod = "PATCH"
)

func (h HTTPMethod) IsSafe() bool {
	return h == GET
}

func (h HTTPMethod) IsIdempotent() bool {
	return h == GET || h == PUT || h == DELETE
}

func main() {
	fmt.Println("=== Example 1: Weekdays ===")
	today := Wednesday
	fmt.Printf("Today is %s (value: %d)\n", today, today)
	fmt.Printf("Weekend starts on %s\n", Saturday)

	fmt.Println("\n=== Example 2: Months ===")
	currentMonth := March
	fmt.Printf("Current month: %s (value: %d)\n", currentMonth, currentMonth)

	fmt.Println("\n=== Example 3: Bit Flags (Permissions) ===")
	userPerms := Read | Write
	adminPerms := Read | Write | Execute | Admin

	fmt.Printf("User permissions: %s\n", userPerms)
	fmt.Printf("Admin permissions: %s\n", adminPerms)
	fmt.Printf("User has Read? %v\n", userPerms.Has(Read))
	fmt.Printf("User has Execute? %v\n", userPerms.Has(Execute))
	fmt.Printf("Admin has Execute? %v\n", adminPerms.Has(Execute))

	// Adding permissions
	userPerms |= Execute
	fmt.Printf("User permissions after adding Execute: %s\n", userPerms)

	// Removing permissions
	userPerms &^= Write
	fmt.Printf("User permissions after removing Write: %s\n", userPerms)

	fmt.Println("\n=== Example 4: File Sizes ===")
	fileSize := ByteSize(1024 * 1024 * 500) // 500 MB
	fmt.Printf("File size: %s\n", fileSize)
	fmt.Printf("1 GB = %s\n", GB)
	fmt.Printf("5 TB = %s\n", 5*TB)

	fmt.Println("\n=== Example 5: Status with Validation ===")
	var status Status = Pending
	fmt.Printf("Status: %s\n", status)
	fmt.Printf("Is valid? %v\n", status.IsValid())
	fmt.Printf("Is terminal? %v\n", status.IsTerminal())

	status = Completed
	fmt.Printf("New status: %s\n", status)
	fmt.Printf("Is terminal? %v\n", status.IsTerminal())

	invalidStatus := Status(999)
	fmt.Printf("Invalid status is valid? %v\n", invalidStatus.IsValid())

	fmt.Println("\n=== Example 6: HTTP Methods ===")
	method := GET
	fmt.Printf("Method: %s\n", method)
	fmt.Printf("Is safe? %v\n", method.IsSafe())
	fmt.Printf("Is idempotent? %v\n", method.IsIdempotent())

	method = POST
	fmt.Printf("Method: %s\n", method)
	fmt.Printf("Is safe? %v\n", method.IsSafe())
	fmt.Printf("Is idempotent? %v\n", method.IsIdempotent())

	fmt.Println("\n=== Type Safety Demo ===")
	// This demonstrates type safety
	var day Weekday = Monday
	var month Month = January

	// The following would cause a compile error:
	// day = month  // ❌ Cannot assign Month to Weekday

	fmt.Printf("Day: %s, Month: %s\n", day, month)
	fmt.Println("Note: You cannot assign a Month to a Weekday variable due to type safety!")
}
