package logger

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"

	"github.com/IBM/sarama"
)

// KafkaLogHandler с реализацией slog.Handler
type KafkaLogHandler struct {
	fileHandler slog.Handler
	producer    sarama.SyncProducer
	topic       string
}

// Создание нового Kafka-логгера
func NewKafkaLogHandler(fileHandler slog.Handler, brokers, topic string) *KafkaLogHandler {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer([]string{brokers}, config)
	if err != nil {
		log.Fatal("Ошибка подключения к Kafka:", err)
	}

	return &KafkaLogHandler{
		fileHandler: fileHandler,
		producer:    producer,
		topic:       topic,
	}
}

// Handle отправляет логи в Kafka + файл
func (h *KafkaLogHandler) Handle(ctx context.Context, r slog.Record) error {
	h.fileHandler.Handle(ctx, r)

	msg := fmt.Sprintf("[%s] %s", r.Level, r.Message)

	if r.Level == -4 {

		_, _, err := h.producer.SendMessage(&sarama.ProducerMessage{
			Topic: h.topic,
			Value: sarama.StringEncoder(msg),
		})
		if err != nil {
			log.Println("Ошибка отправки лога в Kafka:", err)
		}
	}

	return nil
}

// Enabled проверяет, активен ли этот Handler
func (h *KafkaLogHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return true
}

// WithAttrs добавляет атрибуты
func (h *KafkaLogHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &KafkaLogHandler{
		fileHandler: h.fileHandler.WithAttrs(attrs),
		producer:    h.producer,
		topic:       h.topic,
	}
}

// WithGroup добавляет поддержку группировки
func (h *KafkaLogHandler) WithGroup(name string) slog.Handler {
	return &KafkaLogHandler{
		fileHandler: h.fileHandler.WithGroup(name),
		producer:    h.producer,
		topic:       h.topic,
	}
}

// Закрываем Kafka Producer
func (h *KafkaLogHandler) Close() {
	h.producer.Close()
}

// Создание нового логгера
func NewLogger() *slog.Logger {
	opts := slog.HandlerOptions{Level: slog.LevelDebug}

	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal("Ошибка открытия файла лога:", err)
	}

	fileHandler := slog.NewTextHandler(file, &opts)

	kafkaHandler := NewKafkaLogHandler(fileHandler, "38.242.212.205:9092", "logs")

	log.Println("Connected to Kafka with Sarama")

	return slog.New(kafkaHandler)
}
