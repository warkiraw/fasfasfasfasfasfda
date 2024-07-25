document.addEventListener("DOMContentLoaded", function () {
    const tg = window.Telegram.WebApp;
    tg.expand(); // Расширяет веб-приложение на весь экран

    const startButton = document.getElementById("startButton");
    const timerDisplay = document.getElementById("timerDisplay");

    startButton.addEventListener("click", function () {
        startButton.disabled = true;
        let timeLeft = 60;
        timerDisplay.textContent = `Time left: ${timeLeft} seconds`;

        const countdown = setInterval(() => {
            timeLeft -= 1;
            timerDisplay.textContent = `Time left: ${timeLeft} seconds`;

            if (timeLeft <= 0) {
                clearInterval(countdown);
                sendTelegramMessage();
                startButton.disabled = false;
            }
        }, 1000);
    });

    function sendTelegramMessage() {
        const chatId = tg.initDataUnsafe.user.id; // Получение ID пользователя
        const botToken = "5766262321:AAFv5HPDSJU6amPJe38K7-Ho1KAhe9nS7uY"; // Вставьте сюда ваш токен бота
        const message = "One minute has passed!";

        fetch(`/sendMessage?chat_id=${chatId}&message=${message}&token=${botToken}`)
            .then(response => response.json())
            .then(data => {
                if (data.ok) {
                    alert("Message sent successfully!");
                } else {
                    alert("Failed to send message.");
                }
            })
            .catch(error => console.error("Error sending message:", error));
    }
});
