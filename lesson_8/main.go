// package main

// import (
// 	"context"
// 	"fmt"
// 	"math/rand"
// 	"os"
// 	"os/signal"
// 	"sync"
// 	"time"
// )

// type Question struct {
// 	Text    string
// 	Options []string
// 	Correct int
// }

// type Answer struct {
// 	PlayerID int
// 	Option   int
// }

// type Result struct {
// 	Correct   int
// 	Incorrect int
// }

// func main() {
// 	questions := make(chan Question)
// 	answers := make(chan Answer)
// 	results := make(chan Result)
// 	ctx, cancel := context.WithCancel(context.Background())

// 	var wg sync.WaitGroup

// 	wg.Add(1)
// 	go questionGenerator(ctx, questions, &wg)

// 	playerCount := 3
// 	for i := 0; i < playerCount; i++ {
// 		wg.Add(1)
// 		go player(ctx, i, questions, answers, &wg)
// 	}

// 	wg.Add(1)
// 	go answerCounter(ctx, questions, answers, results, &wg)

// 	wg.Add(1)
// 	go func() {
// 		defer wg.Done()
// 		for result := range results {
// 			fmt.Printf("Correct: %d, Incorrect: %d\n", result.Correct, result.Incorrect)
// 		}
// 	}()

// 	sigChan := make(chan os.Signal, 1)
// 	signal.Notify(sigChan, os.Interrupt)
// 	<-sigChan

// 	cancel()
// 	close(questions)
// 	close(answers)
// 	close(results)
// 	wg.Wait()
// 	fmt.Println("Program exited gracefully.")
// }

// func questionGenerator(ctx context.Context, questions chan<- Question, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	round := 0
// 	for {
// 		select {
// 		case <-ctx.Done():
// 			return
// 		case <-time.After(10 * time.Second):
// 			round++
// 			question := Question{
// 				Text:    fmt.Sprintf("Round %d: What is the capital of France?", round),
// 				Options: []string{"Paris", "London", "Berlin", "Rome"},
// 				Correct: 0,
// 			}
// 			fmt.Printf("Generated question: %s\n", question.Text)
// 			questions <- question
// 		}
// 	}
// }

// func player(ctx context.Context, id int, questions <-chan Question, answers chan<- Answer, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for {
// 		select {
// 		case <-ctx.Done():
// 			return
// 		case question := <-questions:
// 			answer := Answer{
// 				PlayerID: id,
// 				Option:   rand.Intn(len(question.Options)),
// 			}
// 			fmt.Printf("Player %d answered: %d\n", id, answer.Option)
// 			answers <- answer
// 		}
// 	}
// }

// func answerCounter(ctx context.Context, questions <-chan Question, answers <-chan Answer, results chan<- Result, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for {
// 		select {
// 		case <-ctx.Done():
// 			return
// 		case question := <-questions:
// 			var correct, incorrect int
// 			for i := 0; i < 3; i++ {
// 				answer := <-answers
// 				if answer.Option == question.Correct {
// 					correct++
// 				} else {
// 					incorrect++
// 				}
// 			}
// 			results <- Result{Correct: correct, Incorrect: incorrect}
// 		}
// 	}
// }

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
