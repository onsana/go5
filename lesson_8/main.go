package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"time"
)

type Question struct {
	Text    string
	Options []string
	Correct int
}

type Answer struct {
	PlayerID int
	Option   int
}

type Result struct {
	Correct   int
	Incorrect int
}

func main() {
	questions := make(chan Question)
	answers := make(chan Answer)
	results := make(chan Result)
	ctx, cancel := context.WithCancel(context.Background())

	var wg sync.WaitGroup

	wg.Add(1)
	go questionGenerator(ctx, questions, &wg)

	playerCount := 3
	for i := 0; i < playerCount; i++ {
		wg.Add(1)
		go player(ctx, i, questions, answers, &wg)
	}

	wg.Add(1)
	go answerCounter(ctx, questions, answers, results, &wg)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for result := range results {
			fmt.Printf("Correct: %d, Incorrect: %d\n", result.Correct, result.Incorrect)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	<-sigChan

	cancel()
	close(questions)
	close(answers)
	close(results)
	wg.Wait()
	fmt.Println("Program exited gracefully.")
}

func questionGenerator(ctx context.Context, questions chan<- Question, wg *sync.WaitGroup) {
	defer wg.Done()
	round := 0
	for {
		select {
		case <-ctx.Done():
			return
		case <-time.After(10 * time.Second):
			round++
			question := Question{
				Text:    fmt.Sprintf("Round %d: What is the capital of France?", round),
				Options: []string{"Paris", "London", "Berlin", "Rome"},
				Correct: 0,
			}
			fmt.Printf("Generated question: %s\n", question.Text)
			questions <- question
		}
	}
}

func player(ctx context.Context, id int, questions <-chan Question, answers chan<- Answer, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case question := <-questions:
			answer := Answer{
				PlayerID: id,
				Option:   rand.Intn(len(question.Options)),
			}
			fmt.Printf("Player %d answered: %d\n", id, answer.Option)
			answers <- answer
		}
	}
}

func answerCounter(ctx context.Context, questions <-chan Question, answers <-chan Answer, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case question := <-questions:
			var correct, incorrect int
			for i := 0; i < 3; i++ {
				answer := <-answers
				if answer.Option == question.Correct {
					correct++
				} else {
					incorrect++
				}
			}
			results <- Result{Correct: correct, Incorrect: incorrect}
		}
	}
}
