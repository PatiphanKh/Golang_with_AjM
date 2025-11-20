# MAUI Application for Student Enrollment
- .Net Maui
- ทำ Application ตัวอย่างของ Application สำหรับลงทะเบียนเรียนทุกหน้า ที่เปลี่ยนไปมาได้ พร้อมข้อมูลจำลอง จากตัวอย่างข้อมูลที่มีให้
- ไม่จำเป็นต้องเชื่อมต่อกับ Back-End แค่ทำ Demo Application

# Create Back-end using Go
- Gorm + Gin
- สร้างเส้น API CRUD ทุกเส้น
- สร้างเส้น API ที่เกี่ยวกับการลงทะเบียนเรียนทั้งหมด
- มีการตรวจสอบและป้องกันข้อผิดพลาด ของการประมวลผลข้อมูลที่เหมาะสม

# Presentation day : In the final exam period

## Create tables and data (MySQL)

```sql
	CREATE TABLE students (
	    student_id INT PRIMARY KEY AUTO_INCREMENT,
	    first_name VARCHAR(50) NOT NULL,
	    last_name VARCHAR(50) NOT NULL,
	    password VARCHAR(255) NULL,
	    image VARCHAR(1000) NOT NULL,
	    date_of_birth DATE,
	    email VARCHAR(100) UNIQUE,
	    phone_number VARCHAR(20)
	);
```

```sql
CREATE TABLE courses (
    course_id INT PRIMARY KEY AUTO_INCREMENT,
    course_name VARCHAR(100) NOT NULL,
    course_code VARCHAR(20) UNIQUE,
    instructor VARCHAR(50)
);
```

```sql
CREATE TABLE enrollment (
    enrollment_id INT PRIMARY KEY AUTO_INCREMENT,
    student_id INT,
    course_id INT,
    enrollment_date DATE,
    FOREIGN KEY (student_id) REFERENCES students(student_id),
    FOREIGN KEY (course_id) REFERENCES courses(course_id)
);
```

## Insert data

```sql
-- Insert 10 Courses
INSERT INTO courses (course_name, course_code, instructor)
VALUES
('Introduction to Programming', 'CS101', 'Prof. Smith'),
('Database Systems', 'DB201', 'Prof. Johnson'),
('Web Development', 'WEB301', 'Prof. Lee'),
('Data Structures and Algorithms', 'DSA401', 'Prof. Patel'),
('Artificial Intelligence', 'AI501', 'Prof. Kim'),
('Machine Learning', 'ML601', 'Prof. Chen'),
('Cybersecurity', 'CYB701', 'Prof. Harris'),
('Network Security', 'NS801', 'Prof. Wilson'),
('Software Engineering', 'SE901', 'Prof. Taylor'),
('Cloud Computing', 'CC1001', 'Prof. Garcia');

-- Insert 10 Students (without passwords)
INSERT INTO students (first_name, last_name, image, date_of_birth, email, phone_number)
VALUES
('Alice', 'Johnson', 'student1.jpg', '2000-01-01', 'alice.johnson@example.com', '123-456-7890'),
('Bob', 'Smith', 'student2.jpg', '2001-02-02', 'bob.smith@example.com', '987-654-3210'),
('Charlie', 'Brown', 'student3.jpg', '2002-03-03', 'charlie.brown@example.com', '555-555-5555'),
('David', 'Lee', 'student4.jpg', '2003-04-04', 'david.lee@example.com', '111-222-3333'),
('Emily', 'Davis', 'student5.jpg', '2004-05-05', 'emily.davis@example.com', '999-888-7777'),
('Frank', 'Miller', 'student6.jpg', '2005-06-06', 'frank.miller@example.com', '444-333-2222'),
('Grace', 'Wilson', 'student7.jpg', '2006-07-07', 'grace.wilson@example.com', '777-666-5555'),
('Henry', 'Moore', 'student8.jpg', '2007-08-08', 'henry.moore@example.com', '222-111-0000'),
('Isabella', 'Taylor', 'student9.jpg', '2008-09-09', 'isabella.taylor@example.com', '333-222-1111'),
('Jack', 'Anderson', 'student10.jpg', '2009-10-10', 'jack.anderson@example.com', '666-555-4444');

-- Insert Enrollment Data
INSERT INTO enrollment (student_id, course_id, enrollment_date)
VALUES
(1, 1, '2023-09-01'),
(1, 2, '2023-09-01'),
(2, 1, '2023-09-01'),
(2, 3, '2023-09-01'),
(3, 2, '2023-09-01'),
(3, 4, '2023-09-01'),
(4, 3, '2023-09-01'),
(4, 5, '2023-09-01'),
(5, 4, '2023-09-01'),
(5, 6, '2023-09-01'),
(6, 5, '2023-09-01'),
(6, 7, '2023-09-01');
```

