package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type Question struct {
	ID      int
	Text    string
	Options []string
	Answer  int
}

type Answer struct {
	PlayerID       int
	QuestionID     int
	SelectedOption int
}

func generateQuestions(ctx context.Context, questions chan<- Question) {
	questionID := 1
	for {
		select {
		case <-ctx.Done():
			close(questions)
			return
		case <-time.After(10 * time.Second):
			question := Question{
				ID:      questionID,
				Text:    fmt.Sprintf("Question %d: What is %d + %d?", questionID, questionID, questionID),
				Options: []string{"1", "2", "3", "4"},
				Answer:  (questionID + questionID) % 4,
			}
			questions <- question
			questionID++
		}
	}
}

func player(ctx context.Context, playerID int, questions <-chan Question, answers chan<- Answer) {
	for {
		select {
		case <-ctx.Done():
			return
		case question, ok := <-questions:
			if !ok {
				return
			}
			selectedOption := rand.Intn(len(question.Options)) // Randomly selecting an answer
			answer := Answer{
				PlayerID:       playerID,
				QuestionID:     question.ID,
				SelectedOption: selectedOption,
			}
			answers <- answer
		}
	}
}

func tallyAnswers(ctx context.Context, answers <-chan Answer, results chan<- map[int]int) {
	for {
		select {
		case <-ctx.Done():
			close(results)
			return
		case answer, ok := <-answers:
			if !ok {
				return
			}
			result := make(map[int]int)
			result[answer.SelectedOption]++
			results <- result
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	questions := make(chan Question)
	answers := make(chan Answer)
	results := make(chan map[int]int)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		generateQuestions(ctx, questions)
	}()

	for i := 1; i <= 5; i++ { // 5 players
		wg.Add(1)
		go func(playerID int) {
			defer wg.Done()
			player(ctx, playerID, questions, answers)
		}(i)
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		tallyAnswers(ctx, answers, results)
	}()

	// Handling OS interrupts for graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-stop
		cancel()
	}()

	for {
		select {
		case <-ctx.Done():
			wg.Wait()
			fmt.Println("Game ended")
			return
		case result, ok := <-results:
			if !ok {
				return
			}
			fmt.Printf("Results: %v\n", result)
		}
	}
}
