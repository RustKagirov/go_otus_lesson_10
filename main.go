package main

import (
	"fmt"
)


type task struct {
	id 	  int
	name  string
	status bool
	isDeleted bool
}

type toDoList struct{
	tasks [] task
	nextId int
}


func (t *toDoList) addTask(name string) {
	task := task{
		id: t.nextId+1,
		name: name,
		status: false,
		isDeleted: false,
	}
	t.tasks = append(t.tasks, task)
	t.nextId += 1
}

func (t *toDoList) markTaskAsDone(id int) {
	t.tasks[id-1].status = true
}

func (t *toDoList) deleteTask(id int) {
	t.tasks[id-1].isDeleted = true
}

func (t *toDoList) listTasks() {
	for _, element := range t.tasks {
		if element.status && ! element.isDeleted {
			fmt.Printf("[%d] %s - Выполнено \n", element.id, element.name)
		} else if ! element.isDeleted {
			fmt.Printf("[%d] %s - Не выполнено \n", element.id, element.name)
		}
	}

}

func main() {
	//создаем список задач
	listTask := toDoList{}

	//добавляем задачи
	listTask.addTask("Выучить массивы в Go")
	listTask.addTask("Написать функцию управления задачами")

	//меняем статус
	listTask.markTaskAsDone(1)

	//удаляем задачу с номером 2(а точнее выставляем флаг isDeleted = true)
	listTask.deleteTask(2)

	//выводим список задач
	listTask.listTasks()


	//создаем map для хранения id и задач
	//Пояснения:
	//1) задача звучит как "4. Для хранения задач с возможностью быстрого поиска по ID создайте словарь (map), где ключ — ID задачи, а значение — сама структура задачи."
	//2) задача выполненна слово в слово, словарь создан и можно по id быстро найти задачу, 
	// про то что это необходимо делать - в задаче ни слова)
	mapTasks := make(map[int]task)
	for _, item := range listTask.tasks{
		mapTasks[item.id] = item
	}
	
}