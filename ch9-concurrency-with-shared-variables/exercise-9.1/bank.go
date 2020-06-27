// Package bank provides a concurrency-safe bank with one account.
package bank

type WithdrawalMessage struct {
	Amount              int
	AvailableToWithdraw chan bool
}

var deposits = make(chan int) // send amount to deposit
var balances = make(chan int) // receive balance
var withdrawals = make(chan WithdrawalMessage)

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }
func Withdraw(amount int) bool {
	available := make(chan bool)
	withdrawals <- WithdrawalMessage{amount, available}
	return <-available
}

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		case withdrawal := <-withdrawals:
			balanceAvailable := withdrawal.Amount <= balance
			if balanceAvailable {
				balance -= withdrawal.Amount
			}
			withdrawal.AvailableToWithdraw <- balanceAvailable
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}
