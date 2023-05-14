package converter

import "github.com/aldyN25/todolist/app/models"

func MapTodosToTodosRes(todos []models.Todos) (todosRes []models.TodosRes) {
	for _, v := range todos {
		todosRes = append(todosRes, *v.ToTodosRes())
	}
	return
}

func MapActivitiesToActivitiesRes(activities []models.Activities) (activitiesRes []models.ActivitiesRes) {
	for _, v := range activities {
		activitiesRes = append(activitiesRes, *v.ToActivitiesRes())
	}
	return
}
