package main

type Level struct {
	Number      int    `json:"number"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Question    string `json:"question"`
	Answer      string `json:"-"`
	Hint        string `json:"hint"`
	Points      int    `json:"points"`
}

var levels = []Level{
	{
		Number:      1,
		Title:       "The Secret Number",
		Description: "Find the hidden number.",
		Question: `
I am a number between 1 and 20.

Clues:
- I am even.
- I am greater than 10.
- I am less than 16.
- My digits add up to 3.

What number am I?`,
		Answer: "12",
		Hint:   "There are only two even numbers between 10 and 16.",
		Points: 100,
	},

	{
		Number:      2,
		Title:       "The Password",
		Description: "Crack the mysterious password.",
		Question: `
A strange lock displays:

2, 6, 12, 20, 30, ?

The pattern is:

1 x 2 = 2
2 x 3 = 6
3 x 4 = 12
4 x 5 = 20
5 x 6 = 30

What comes next?`,
		Answer: "42",
		Hint:   "Try multiplying consecutive numbers.",
		Points: 200,
	},

	{
		Number:      3,
		Title:       "The Final Door",
		Description: "Solve the final logic puzzle.",
		Question: `
There are three boxes:

BOX A
BOX B
BOX C

Only ONE box contains treasure.

The labels say:

BOX A:
"The treasure is not in Box B."

BOX B:
"The treasure is in Box A."

BOX C:
"The treasure is not in this box."

Exactly ONE statement is TRUE.

Which box contains the treasure?`,
		Answer: "c",
		Hint:   "Test each box and count how many statements are true.",
		Points: 300,
	},
}
