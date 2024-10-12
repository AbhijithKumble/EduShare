package types

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type UserAcc struct {
	UserID                    uuid.UUID  `json:"id"`
	Email                     string     `json:"email"`
	Password                  string     `json:"-"`
	CreatedAt                 time.Time  `json:"createdAt"`
	UpdatedAt                 time.Time  `json:"updatedAt"`
	IsVerified                bool       `json:"isVerified"`
	IsAdmin                   bool       `json:"isAdmin"`
	VerificationToken         *string    `json:"verificationToken"`
	VerificationTokenExpiry   *time.Time `json:"verificationTokenExpiry"`
	ForgotPasswordToken       *string    `json:"forgotPasswordToken"`
	ForgotPasswordTokenExpiry *time.Time `json:"forgotPasswordTokenExpiry"`
}

type UserInfo struct {
	UserID      uuid.UUID `json:"id"`
	FirstName   string    `json:"firstName"`
	MiddleName  string    `json:"middleName"`
	LastName    string    `json:"lastName"`
	Dept        string    `json:"dept"`
	IsDeptAdmin bool      `json:"isDeptAdmin"`
}

type UserInfoPayLoad struct {
	Email      string `json:"email"`
	FirstName  string `json:"firstName"`
	MiddleName string `json:"middleName"`
	LastName   string `json:"lastName"`
	Dept       string `json:"dept"`
}

type LoginUserPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type ForgotPasswordPayLoad struct {
	Email string `json:"email" validate:"required,email"`
}

type EmailPayload struct {
	From struct {
		Email string `json:"email"`
		Name  string `json:"name"`
	} `json:"from"`
	To []struct {
		Email string `json:"email"`
	} `json:"to"`
	Subject  string `json:"subject"`
	Text     string `json:"text"`
	Category string `json:"category"`
}

type UserStore interface {
	GetUserByEmail(c context.Context, email string) (UserAcc, error, int)
	CreateUser(c context.Context, user UserAcc) (*uuid.UUID, error)
	VerifyOtp(c context.Context, email string, password string, otp string) (int, error)
}

type OtpPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Otp      string `json:"otp"`
}

type Dept struct {
	DeptCode string `json:"deptCode"`
	DeptName string `json:"deptName"`
}

type DeptPayload struct {
	DeptCode string `json:"deptCode" validate:"required, deptCode"`
	DeptName string `json:"deptName" validate:"required, deptName"`
}

type DeptStore interface {
	GetDepts(c context.Context) ([]Dept, error)
	CreateDepts(c context.Context, dept DeptPayload) error
}

type Course struct {
	Dept
	CourseCode string `json:"courseCode"` //course num
	CourseName string `json:"courseName"`
}

type CoursePayload struct {
	DeptPayload
	CourseCode string `json:"courseCode" validate:"required, courseCode"`
	CourseName string `json:"courseName" validate:"required, courseName"`
}

type CourseStore interface {
	GetCourses(c context.Context) ([]Course, error)
	CreateCourses(c context.Context, courses CoursePayload) error
}

type Favourites struct {
	UserId string `json:"id"`
	Course
}

type Resources struct {
	Course
	FileType     string `json:"fileType"`
	ExamType     string `json:"examType"`
	Year         int    `json:"year"`
	FileLocation string `json:"fileLocation"`
	IsApproved   bool   `json:"isApproved"`
}

type ResourcePayload struct {
	CoursePayload
	FileType string `json:"fileType" validate:"required"`
	ExamType string `json:"examType" validate:"required"`
	Year     int    `json:"year" validate:"required"`
}
