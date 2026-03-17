package main

import (
	"bufio"
	"fmt"
	"os"
)

type Client struct {
	balance int
	name    string
}

func NewClient(name string) *Client {
	return &Client{
		name:    name,
		balance: 0,
	}
}

func (c *Client) income(p int) {
	if c.balance > 0 {
		c.balance += (c.balance * p) / 100
	}
}

func (c *Client) deposit(d int) {
	c.balance += d
}

func (c *Client) withdraw(d int) {
	c.balance -= d
}

func (c Client) getBalance() int {
	return c.balance
}

type Bank struct {
	base map[string]*Client
}

func NewBank() *Bank {
	return &Bank{
		base: make(map[string]*Client),
	}
}

func (b *Bank) getOrCreateClient(name string) *Client {
	if _, ok := b.base[name]; !ok {
		b.base[name] = NewClient(name)
	}
	return b.base[name]
}

func (b *Bank) Income(p int) {
	for _, c := range b.base {
		c.income(p)
	}
}

func (b *Bank) Deposit(name string, p int) {
	client := b.getOrCreateClient(name)
	client.deposit(p)
}

func (b *Bank) Withdraw(name string, p int) {
	client := b.getOrCreateClient(name)
	client.withdraw(p)
}

func (b *Bank) Transfer(name1, name2 string, sum int) {
	client1 := b.getOrCreateClient(name1)
	client2 := b.getOrCreateClient(name2)

	client1.withdraw(sum)
	client2.deposit(sum)
}

func (b *Bank) Balance(name string) (int, bool) {
	client, ok := b.base[name]
	if !ok {
		return 0, false
	}
	return client.getBalance(), true
}

func main() {
	in := bufio.NewReader(os.Stdin)
	bank := NewBank()

	for {
		var command string
		_, err := fmt.Fscan(in, &command)
		if err != nil {
			break
		}

		switch command {
		case "DEPOSIT":
			var name string
			var sum int
			fmt.Fscan(in, &name, &sum)
			bank.Deposit(name, sum)

		case "WITHDRAW":
			var name string
			var sum int
			fmt.Fscan(in, &name, &sum)
			bank.Withdraw(name, sum)

		case "BALANCE":
			var name string
			fmt.Fscan(in, &name)
			balance, ok := bank.Balance(name)
			if !ok {
				fmt.Println("ERROR")
			} else {
				fmt.Println(balance)
			}

		case "TRANSFER":
			var name1, name2 string
			var sum int
			fmt.Fscan(in, &name1, &name2, &sum)
			bank.Transfer(name1, name2, sum)

		case "INCOME":
			var p int
			fmt.Fscan(in, &p)
			bank.Income(p)
		}
	}
}
