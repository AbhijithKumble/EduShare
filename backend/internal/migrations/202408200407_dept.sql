-- +goose Up 
CREATE TABLE dept( 
  DeptCode VARCHAR(10) NOT NULL,    
  DeptName VARCHAR(40) NOT NULL,
);
-- +goose Down 
DROP TABLE dept;
