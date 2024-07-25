package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "html/template"
    "strings"
)

// Структура для обработки ответа от Telegram API
type TelegramResponse struct {
    Ok          bool   `json:"ok"`
    Description string `json:"description"`
}

// Обработчик для главной страницы
func HomePageHandler(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.ParseFiles("templates/index.html"))
    err := tmpl.Execute(w, nil)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

// Обработчик для отправки сообщения
func SendMessageHandler(w http.ResponseWriter, r *http.Request) {
    chatID := r.URL.Query().Get("chat_id")
    message := r.URL.Query().Get("message")
    token := r.URL.Query().Get("token")

    if chatID == "" || message == "" || token == "" {
        http.Error(w, "Missing parameters", http.StatusBadRequest)
        return
    }

    // Формируем URL для отправки сообщения через Telegram API
    url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", token)
    reqBody := fmt.Sprintf(`{"chat_id":%s,"text":"%s"}`, chatID, message)

    req, err := http.NewRequest("POST", url, strings.NewReader(reqBody))
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    req.Header.Set("Content-Type", "application/json")
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer resp.Body.Close()

    var telegramResponse TelegramResponse
    json.NewDecoder(resp.Body).Decode(&telegramResponse)

    if !telegramResponse.Ok {
        http.Error(w, telegramResponse.Description, http.StatusInternalServerError)
        return
    }

    w.Write([]byte(`{"ok": true}`))
}

func main() {
    http.HandleFunc("/", HomePageHandler)
    http.HandleFunc("/sendMessage", SendMessageHandler)
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

    fmt.Println("Server is running on http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
