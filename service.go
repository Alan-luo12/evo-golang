package main

import (
	"log"
	"time"
)

func Processtask(id int64, delay_time int) {

	_, err := DB.Exec("UPDATE tasks SET status='running', updated_at = CURRENT_TIMESTAMP WHERE id = ?", id)
	if err != nil {
		log.Println("failed to running task")
		return
	}

	time.Sleep(time.Duration(delay_time) * time.Millisecond)

	_, err = DB.Exec("UPDATE tasks SET status = 'done', updated_at = CURRENT_TIMESTAMP WHERE id = ?", id)
	if err != nil {
		log.Printf("failed to done task %v Errror %v", id, err)
		return
	}
}
