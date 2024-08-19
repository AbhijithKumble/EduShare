-- +goose Up 
CREATE TABLE userinfo( 
    UserID  UUID NOT NULL, 
    FirstName VARCHAR(20),
    MiddleName VARCHAR(10),
    LastName VARCHAR(10),
    Dept  VARCHAR(8),                
    IsDeptAdmin BOOLEAN DEFAULT FALSE,      
    FOREIGN KEY(UserID) REFERENCES users(UserID) 
);
-- +goose Down 
DROP TABLE userinfo;
