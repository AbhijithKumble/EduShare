-- +goose Up 
CREATE TABLE courses( 
  DeptCode    VARCHAR(10) NOT NULL, 
  CourseNum   VARCHAR(10) NOT NULL,    
  CourseName  VARCHAR(25) NOT NULL,
  PRIMARY KEY (DeptCode , CourseNum),
  FOREIGN KEY (DeptCode) REFERENCES dept(DeptCode)
);
-- +goose Down 
DROP TABLE courses;
