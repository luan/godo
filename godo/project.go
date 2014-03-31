package godo

type Project struct {
	ID     int
	Name   string
}

type ProjectDomain struct {
	Tasks []Task
	Name   string
}

func NewProject(name string) Project {
	return Project{Name: name}
}

func NewProjectDomain(name string) ProjectDomain {
	return ProjectDomain{Name: name}
}

