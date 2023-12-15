# GoInventoryTracker

GoInventoryTracker is a command-line tool for managing inventory in retail stores or warehouses. It utilizes a stack, a binary tree, and a hashmap to provide a comprehensive solution for inventory management. Written for funsies and DSA implementations.

## Features

1. **Stack for Items on Sale:**
   - `sale push <items>`: Add an item to the list of items on sale.
   - `sale pop`: Remove the last item added to the sale list.
   - `sale peek`: View the current item on sale.
   - `sale display`: Show the list of items currently on sale.

2. **Binary Tree for Product Hierarchy:**
   - `product add <category> <product>`: Add a product to a specific category in the product hierarchy.
   - `product remove <product>`: Remove a product from the product hierarchy.
   - `product search <product>`: Search for a product in the hierarchy.
   - `product display`: Display the product hierarchy.

3. **HashMap for Stock Levels:**
   - `stock add <product> <quantity>`: Add stock for a specific product.
   - `stock remove <product> <quantity>`: Remove stock for a specific product.
   - `stock check <product>`: Check the current stock level for a product.
   - `stock display`: Display the overall stock levels.

4. **General Operations:**
   - `report sales`: Generate a report of recent sales.
   - `report low-stock`: Identify products with low stock levels.
   - `help`: Display a list of available commands and their descriptions.
   - `exit`: Close the program.

## Usage

```bash
$ ./go-inventory_tracker.go
> sale push Laptop
> product add Electronics Smartphone
> stock add Smartphone 50
> report low-stock
Low Stock Report:
- Smartphone: 50 units
```

## Getting Started

### Prerequisites

Golang installed

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/cs50-romain/GoInventoryTracker.git
   cd GoInventoryTracker
