-- +goose Up 
CREATE TABLE dept( 
  DeptCode VARCHAR(40) NOT NULL,    
  DeptName VARCHAR(40) NOT NULL,
  Primary Key(DeptCode)
);
-- +goose Down 
DROP TABLE dept;
