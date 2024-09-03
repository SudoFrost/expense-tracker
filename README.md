# Expense Tracker
=================

Expense Tracker is a command line application that allows you to track your expenses.

## Installation
------------

To install the expense tracker, follow these steps:

1. Make sure you have Go installed on your machine.
2. Clone the repository: `git clone https://github.com/sudofrost/expense-tracker.git`
3. Navigate to the project directory: `cd expense-tracker`
4. Build the project: `go build`
5. Run the application: `./expense-tracker`

## Usage
-----

### Listing Expenses

To list all expenses, run the following command:

```bash
./expense-tracker list
```

### Adding an Expense

To add an expense, run the following command:

```bash
./expense-tracker add --description "Grocery shopping" --amount 50.00
```

### Updating an Expense

To update an expense, run the following command:

```bash
./expense-tracker update 1 --amount 20.00
```

### Deleting an Expense

To delete an expense, run the following command:

```bash
./expense-tracker delete --id 1
```

### Summary of Expenses

To get a summary of expenses for a specific month, run the following command:

```bash
./expense-tracker summary --month 1
```

### Project Page

Visit the project page for more information: [Project URL](https://roadmap.sh/projects/expense-tracker)