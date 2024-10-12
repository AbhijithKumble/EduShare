-- +goose Up 
CREATE TABLE courses( 
  DeptCode    VARCHAR(10) NOT NULL, 
  CourseCode   VARCHAR(10) NOT NULL,    
  CourseName  VARCHAR(100) NOT NULL,
  PRIMARY KEY (DeptCode , CourseCode),
  FOREIGN KEY (DeptCode) REFERENCES dept(DeptCode)
);

-- +goose Down 
DROP TABLE courses;
