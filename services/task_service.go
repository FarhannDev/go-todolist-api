package services

import (
	"fmt"
	"sync"
	"time"
	"todolist-api/models"
)

var mu sync.Mutex
var wg sync.WaitGroup
var ch = make(chan string, 10)

func AddTask(task string, description string) string {
	wg.Add(1)
	go func() {
		defer wg.Done()

		mu.Lock()
		newTask := models.Task{
			ID:          len(models.Tasks) + 1,
			Task:        task,
			Description: description,
			Completed:   false,
		}
		models.Tasks = append(models.Tasks, newTask)
		mu.Unlock()

		time.Sleep(1 * time.Second)
		ch <- fmt.Sprintln("Tugas berhasil ditambahkan!", task)
	}()

	wg.Wait()
	return <-ch
}

func Gettask() []models.Task {
	mu.Lock()
	defer mu.Unlock()
	return models.Tasks
}
func CompleteTask(id int) string {
	wg.Add(1)
	go func() {
		defer wg.Done()

		mu.Lock()
		for i, task := range models.Tasks {
			if task.ID == id {
				models.Tasks[i].Completed = true
				mu.Unlock()
				time.Sleep(1 * time.Second)
				ch <- fmt.Sprintf("✅ Tugas '%s' telah selesai!", task.Task)
				return
			}
		}
		mu.Unlock()
		ch <- fmt.Sprintf("Tugas dengan ID %d tidak ditemukan!", id)
	}()

	wg.Wait()
	return <-ch
}

func DeleteTask(id int) string {
	wg.Add(1)
	go func() {
		defer wg.Done()

		mu.Lock()
		for i, task := range models.Tasks {
			if task.ID == id {
				models.Tasks = append(models.Tasks[:i], models.Tasks[i+1:]...)
				mu.Unlock()
				time.Sleep(1 * time.Second)
				ch <- fmt.Sprintf("Tugas '%s' telah dihapus!", task.Task)
				return
			}
		}
		mu.Unlock()
		ch <- fmt.Sprintf(" Tugas dengan ID %d tidak ditemukan!", id)
	}()
	wg.Wait()
	return <-ch
}
