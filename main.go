package main

import (
	"bufio"
	"fmt"
	"strconv"

	//"log"
	"os"
	"strings"

	"github.com/cs50-romain/GoInventoryTracker/sale"
	"github.com/cs50-romain/GoInventoryTracker/stock"
)

var sale_stack = sale.Init()
var stocks = stock.Init()

func saleFunctions(args []string) {
	if len(args) == 1 {
		fmt.Println("Need more arguments. Type help to view a list of commands.")
	} else if len(args) == 2 {
		sale_cmd := args[1]
		if sale_cmd == "push" {
			fmt.Println("Not enough arguments.")
		} else if sale_cmd == "peek" {
			fmt.Println(sale_stack.Peek())
		} else if sale_cmd == "display" {
			sale_stack.Display()
		} else if sale_cmd == "pop" {
			sale_stack.Pop()
		}
	} else if len(args) == 3 {
		sale_cmd := args[1]
		item := args[2]

		if sale_cmd == "push" {
			sale_stack.Push(item, 10)
		} else {
			fmt.Println("Invalid command. Type help to view a list of commands.")
		}
	}
}

func productFunctions(args []string) {
	fmt.Println(args)
}

func stockFunctions(args []string) {
	if len(args) == 1 {
		fmt.Println("Need more arguments. Type help to view a list of commands.")
	} else if len(args) == 2 {
		stock_cmd := args[1]
		if stock_cmd != "display" {
			fmt.Println("Not enough arguments.")
		} else {
			stocks.Display()
		}
	} else if len(args) >= 3 {
		stock_cmd := args[1]
		item := args[2]

		if stock_cmd == "add" {
			if len(args) < 3 {
				fmt.Println("Need more arguments. Type help to view a list of commands.")
			} else if len(args) == 3 {
				stocks.Add(item, 1)
			} else if len(args) == 4 {
				quantity, _ := strconv.Atoi(args[3])
				stocks.Add(item, quantity)
			}
		} else if stock_cmd == "remove" {
			if len(args) < 3 {
				fmt.Println("Need more arguments. Type help to view a list of commands.")
			} else if len(args) == 3 {
				stocks.Remove(item, 1)
			} else if len(args) == 4 {
				quantity, _ := strconv.Atoi(args[3])
				ok := stocks.Remove(item, quantity)
				if !ok {
					fmt.Println("[INFO] Item does not exist OR not enough inventory.")
				}
			}
		} else if stock_cmd == "check" {
			fmt.Println(stocks.Check(item))
		} else {
			fmt.Println("Invalid command. Type help to view a list of commands.")
		}
	}
}

func reportFunctions(args []string) {
	fmt.Println(args)
}

func parseCommands() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	cmd := scanner.Text()

	args := strings.Split(cmd, " ")

	if cmd == "help" {
		fmt.Println("Listing possible commands")
		fmt.Println(`
		- sale push <items>: Add an item to the list of items on sale.
	   	- sale pop: Remove the last item added to the sale list.
		- sale peek: View the current item on sale.
		- sale display: Show the list of items currently on sale.

		- product add <category> <product>: Add a product to a specific category in the product hierarchy.
		- product remove <product>: Remove a product from the product hierarchy.
		- product search <product>: Search for a product in the hierarchy.
		- product display: Display the product hierarchy.

		- stock add <product> <quantity>: Add stock for a specific product.
		- stock remove <product> <quantity>: Remove stock for a specific product.
		- stock check <product>: Check the current stock level for a product.
		- stock display: Display the overall stock levels.

		- report sales: Generate a report of recent sales.
		- report low-stock: Identify products with low stock levels.
		- help: Display a list of available commands and their descriptions.
		- exit: Close the program.
		`)
	} else if cmd == "exit" {
		os.Exit(1)
	} else if args[0] == "sale" {
		// Implements stack
		saleFunctions(args)
	} else if args[0] == "product" {
		// Implement binary tree
		productFunctions(args)
	} else if args[0] == "stock" {

		stockFunctions(args)
	} else if args[0] == "report" {
		reportFunctions(args)
	} else {
		fmt.Println("Invalid command. Type help to view a list of commands.")
	}
}

func main() {
	fmt.Println("Type help to view a list of possible commands")
	for {
		fmt.Print("> ")
		parseCommands()
	}
}
