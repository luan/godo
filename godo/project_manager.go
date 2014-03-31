package godo

import ("database/sql")

type ProjectManager struct {
	manager
}

type taskProjectRow struct {
	Pid int
	Pname sql.NullString
	Tname sql.NullString
	Tstatus sql.NullString
	Tid sql.NullInt64
}

func NewProjectManager() *ProjectManager {
	return &ProjectManager{}
}

func (pm *ProjectManager) FindAllWithTasks() (projects []ProjectDomain, err error) {
	array := []taskProjectRow{}
	err = Dbmap.Select(&array, "select p.id as pid, p.name as pname, t.id as tid, t.name as tname, t.status as tstatus from projects p left join tasks t on p.id = t.projectid order by pid;")

	projects = []ProjectDomain{}
	previousPid := -1

	for _, row := range array {

		if (previousPid != row.Pid) {
			projects = append(projects, NewProjectDomain(row.Pname.String))
		}
		previousPid = row.Pid

		if (row.Tid.Valid) {
			projects[len(projects)-1].Tasks = append(projects[len(projects)-1].Tasks, NewTaskDomain(row.Tname.String, row.Tstatus.String))
		}
	}

	return
}

func (tm *ProjectManager) FindAll() (projects []Project, err error) {
	err = tm.manager.FindAll(&projects, "projects")
	return
}
