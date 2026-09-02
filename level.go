package main

type PuzzleType string

const (
	PuzzleNumber    PuzzleType = "number"
	PuzzleSequence  PuzzleType = "sequence"
	PuzzleChoice    PuzzleType = "choice"
	PuzzleTrueFalse PuzzleType = "true_false"
)

type Level struct {
	ID          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Type        PuzzleType `json:"type"`
	Question    string     `json:"question"`
	Options     []string   `json:"options,omitempty"`
	Answer      string     `json:"-"`
	Hint        string     `json:"hint"`
	Points      int        `json:"points"`
	Difficulty  string     `json:"difficulty"`
}

var levels = []Level{

	{
		ID:          1,
		Title:       "The Secret Number",
		Description: "Find the hidden number.",
		Type:        PuzzleNumber,
		Question: `
I am a number between 1 and 20.

Clues:
- I am even.
- I am greater than 10.
- I am less than 16.
- My digits add up to 3.

What number am I?`,
		Answer:     "12",
		Hint:       "There are only two even numbers between 10 and 16.",
		Points:     100,
		Difficulty: "Easy",
	},

	{
		ID:          2,
		Title:       "The Sequence",
		Description: "Find the missing number.",
		Type:        PuzzleSequence,
		Question: `
What number comes next?

2, 6, 12, 20, 30, ?`,
		Answer:     "42",
		Hint:       "Look at the difference between each number.",
		Points:     120,
		Difficulty: "Easy",
	},

	{
		ID:          3,
		Title:       "The Right Door",
		Description: "Choose the door that leads to safety.",
		Type:        PuzzleChoice,
		Question: `
Three doors are in front of you.

Door A says:
"The treasure is behind Door B."

Door B says:
"The treasure is not behind Door C."

Door C says:
"The treasure is behind Door C."

Only ONE statement is true.

Which door contains the treasure?`,
		Options: []string{
			"Door A",
			"Door B",
			"Door C",
		},
		Answer:     "door c",
		Hint:       "Test each possible door and count the true statements.",
		Points:     150,
		Difficulty: "Medium",
	},

	{
		ID:          4,
		Title:       "True or False",
		Description: "Think carefully before answering.",
		Type:        PuzzleTrueFalse,
		Question: `
Every square is a rectangle.

True or False?`,
		Options: []string{
			"True",
			"False",
		},
		Answer:     "true",
		Hint:       "A square has four right angles.",
		Points:     100,
		Difficulty: "Easy",
	},

	{
		ID:          5,
		Title:       "The Odd One",
		Description: "Find the number that does not belong.",
		Type:        PuzzleChoice,
		Question: `
Which number does NOT belong?

2, 4, 8, 16, 31, 64`,
		Options: []string{
			"16",
			"31",
			"64",
		},
		Answer:     "31",
		Hint:       "Look at what happens when you multiply by 2.",
		Points:     150,
		Difficulty: "Medium",
	},

	{
		ID:          6,
		Title:       "The Clock",
		Description: "Solve the time puzzle.",
		Type:        PuzzleNumber,
		Question: `
A clock shows 3:00.

What is the angle between the hour hand
and the minute hand?`,
		Answer:     "90",
		Hint:       "At 3:00, the hands form a right angle.",
		Points:     150,
		Difficulty: "Medium",
	},

	{
		ID:          7,
		Title:       "The Sequence II",
		Description: "This sequence is getting harder.",
		Type:        PuzzleSequence,
		Question: `
What number comes next?

1, 4, 9, 16, 25, ?`,
		Answer:     "36",
		Hint:       "Think about square numbers.",
		Points:     180,
		Difficulty: "Medium",
	},

	{
		ID:          8,
		Title:       "The Truth Teller",
		Description: "One statement is correct.",
		Type:        PuzzleChoice,
		Question: `
Alice says:
"The treasure is in Box A."

Bob says:
"The treasure is not in Box A."

Exactly ONE person is telling the truth.

Where is the treasure?`,
		Options: []string{
			"Box A",
			"Box B",
			"Impossible to know",
		},
		Answer:     "impossible to know",
		Hint:       "If Alice is telling the truth, what about Bob?",
		Points:     200,
		Difficulty: "Hard",
	},

	{
		ID:          9,
		Title:       "The Final Sequence",
		Description: "Find the hidden pattern.",
		Type:        PuzzleNumber,
		Question: `
What number is missing?

3, 9, 27, 81, ?`,
		Answer:     "243",
		Hint:       "Each number is multiplied by the same value.",
		Points:     200,
		Difficulty: "Hard",
	},

	{
		ID:          10,
		Title:       "The Final Door",
		Description: "Solve the final puzzle and unlock the treasure.",
		Type:        PuzzleTrueFalse,
		Question: `
A farmer has 17 sheep.

All but 9 run away.

How many sheep remain?

Answer True if the answer is 9.
Answer False otherwise.`,
		Options: []string{
			"True",
			"False",
		},
		Answer:     "true",
		Hint:       "Read the sentence carefully: 'all but 9'.",
		Points:     300,
		Difficulty: "Hard",
	},
}
