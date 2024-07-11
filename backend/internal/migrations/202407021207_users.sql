-- +goose Up 
CREATE TABLE users( 
    UserID  UUID NOT NULL PRIMARY KEY, 
    FirstName VARCHAR(20),
    MiddleName VARCHAR(10),
    LastName VARCHAR(10),
    CreatedAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UpdatedAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    Email VARCHAR(255) UNIQUE NOT NULL,
    Password VARCHAR(255) NOT NULL,
    Dept  VARCHAR(8),                
    IsDeptAdmin BOOLEAN DEFAULT FALSE,      
	IsVerified  BOOLEAN DEFAULT FALSE,                   
	VerificationToken VARCHAR(64),         
	VerificationTokenExpiry  TIMESTAMP,  
	ForgotPasswordToken      VARCHAR(64),    
	ForgotPasswordTokenExpiry TIMESTAMP
);

-- +goose Down 
DROP TABLE users;
