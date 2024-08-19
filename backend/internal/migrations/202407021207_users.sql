-- +goose Up 
CREATE TABLE users( 
    UserID  UUID NOT NULL PRIMARY KEY, 
    Email VARCHAR(255) UNIQUE NOT NULL,
    Password VARCHAR(255) NOT NULL,
    CreatedAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UpdatedAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  	IsVerified  BOOLEAN DEFAULT FALSE,                   
    IsAdmin BOOLEAN DEFAULT FALSE,
  	VerificationToken VARCHAR(64) DEFAULT NULL,         
  	VerificationTokenExpiry  TIMESTAMP DEFAULT NULL,  
  	ForgotPasswordToken      VARCHAR(64) DEFAULT NULL,    
  	ForgotPasswordTokenExpiry TIMESTAMP DEFAULT NULL
);
-- +goose Down 
DROP TABLE users;
