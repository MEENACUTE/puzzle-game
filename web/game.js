let currentState = null;


async function startGame() {

    const nameInput =
        document.getElementById("nameInput");

    const name = nameInput.value.trim();

    const response = await fetch("/api/start", {

        method: "POST",

        headers: {
            "Content-Type": "application/json"
        },

        body: JSON.stringify({
            name: name
        })
    });

    currentState = await response.json();

    document
        .getElementById("startScreen")
        .classList.add("hidden");

    document
        .getElementById("gameScreen")
        .classList.remove("hidden");

    await updateScreen();
}


async function updateScreen() {

    const response =
        await fetch("/api/state");

    currentState =
        await response.json();

    document
        .getElementById("playerName")
        .textContent =
        currentState.playerName;

    document
        .getElementById("score")
        .textContent =
        currentState.score;

    document
        .getElementById("lives")
        .textContent =
        currentState.lives;

    document
        .getElementById("hints")
        .textContent =
        currentState.hints;

    await loadCurrentLevel();
}


async function loadCurrentLevel() {

    if (
        currentState.gameOver ||
        currentState.finished
    ) {
        return;
    }

    const response =
        await fetch("/api/state");

    const state =
        await response.json();

    currentState = state;

    document
        .getElementById("levelNumber")
        .textContent =
        "LEVEL " + state.level;

    /*
        The actual level data is returned
        from a new API endpoint.
    */

    const levelResponse =
        await fetch("/api/level");

    const level =
        await levelResponse.json();

    document
        .getElementById("levelTitle")
        .textContent =
        level.title;

    document
        .getElementById("levelDescription")
        .textContent =
        level.description;

    document
        .getElementById("question")
        .textContent =
        level.question;
}


async function submitAnswer() {

    const input =
        document.getElementById("answerInput");

    const answer =
        input.value.trim();

    if (answer === "") {

        showMessage(
            "Please enter an answer."
        );

        return;
    }

    const response =
        await fetch("/api/answer", {

            method: "POST",

            headers: {
                "Content-Type":
                    "application/json"
            },

            body: JSON.stringify({
                answer: answer
            })
        });

    const result =
        await response.json();

    currentState =
        result.game;

    input.value = "";

    if (result.correct) {

        if (currentState.finished) {

            showEndScreen(
                "CONGRATULATIONS!",
                "You solved all three puzzles!"
            );

            return;
        }

        showMessage(
            "Correct! Moving to the next puzzle..."
        );

        await updateScreen();

        return;
    }

    if (currentState.gameOver) {

        showEndScreen(
            "GAME OVER",
            "You ran out of lives."
        );

        return;
    }

    showMessage(
        "Wrong answer! Try again."
    );

    await updateScreen();
}


async function getHint() {

    const response =
        await fetch("/api/hint", {

            method: "POST"
        });

    const result =
        await response.json();

    currentState =
        result.game;

    if (result.hint === "") {

        showMessage(
            "You have no hints left."
        );

        return;
    }

    showMessage(
        "HINT: " + result.hint
    );

    await updateScreen();
}


function handleEnter(event) {

    if (event.key === "Enter") {

        submitAnswer();
    }
}


function showMessage(message) {

    document
        .getElementById("message")
        .textContent =
        message;
}


function showEndScreen(title, message) {

    document
        .getElementById("gameScreen")
        .classList.add("hidden");

    document
        .getElementById("endScreen")
        .classList.remove("hidden");

    document
        .getElementById("endTitle")
        .textContent =
        title;

    document
        .getElementById("endMessage")
        .textContent =
        message;

    document
        .getElementById("finalScore")
        .textContent =
        currentState.score;
}