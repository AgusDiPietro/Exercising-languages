package tasks

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	ID          int
	Name        string
	Description string
	Completed   bool
}

//Añadir tareas
func Add(db *gorm.DB, task Task) *Task {
	task.Completed = false
	db.Create(&task)
	return &task
}

//Obtener todas las tareas
func GetAll(db *gorm.DB) []Task {
	var tasks []Task
	db.Find(&tasks)
	return tasks
}

//Obtener una tarea a travez del ID
func GetByID(db *gorm.DB, id int) *Task {
	var task Task
	db.Find(&task, id)
	return &task
}

//Borrar filtrado por id
func DeleteByID(db *gorm.DB, id int) {
	result := db.Delete(&Task{}, id)
	if result.Error != nil {
		panic(result.Error) //se para el programa y sale msj en la terminal
	}
}

//Edita la tabla de la DB con los registros
func UpdateByID(db *gorm.DB, id int, task Task) *Task {
	var taskToUpdate Task
	resultTask := db.First(&taskToUpdate, id)
	if resultTask.Error != nil {
		panic(resultTask.Error)
	}

	taskToUpdate = task

	resultSave := db.Save(&taskToUpdate)

	if resultSave.Error != nil {
		panic(resultSave.Error)
	}

	return &taskToUpdate
}
