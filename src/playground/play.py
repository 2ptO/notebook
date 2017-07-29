#playing with python constructs

#maps and lists (slices in Go)

class Student:
    def __init__(self, id, name):
        self.id = id
        self.name = name

def enrollStudents(course, students):
    for s in students:
        course[s.id] = s

def listStudents(course):
    for s in course:
        print course[s].id, course[s].name

def main():
    students = [Student(1,"alice"), Student(2, "bob"), Student(3,"Cathy")]
    coursePhysics = {}
    enrollStudents(coursePhysics, students)
    listStudents(coursePhysics)

if __name__ == "__main__":
    main()
