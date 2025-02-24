package logger

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

// Кастомный Handler, который пишет логи в файл и Kafka
type KafkaLogHandler struct {
	fileHandler slog.Handler
	producer    *kafka.Producer
	topic       string
}

// Создаем новый Kafka-логгер
func NewKafkaLogHandler(fileHandler slog.Handler, brokers, topic string) *KafkaLogHandler {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": brokers})
	if err != nil {
		log.Fatal("Ошибка подключения к Kafka:", err)
	}

	return &KafkaLogHandler{
		fileHandler: fileHandler,
		producer:    producer,
		topic:       topic,
	}
}

// Метод Handle отправляет логи в Kafka + файл
func (h *KafkaLogHandler) Handle(ctx context.Context, r slog.Record) error {
	// Сначала записываем лог в файл
	h.fileHandler.Handle(ctx, r)

	// Форматируем лог в текст
	msg := fmt.Sprintf("[%s] %s", r.Level, r.Message)

	// Отправляем в Kafka

	if r.Level == -4 {
		err := h.producer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &h.topic, Partition: kafka.PartitionAny},
			Value:          []byte(msg),
		}, nil)
		if err != nil {
			log.Println("Ошибка отправки лога в Kafka:", err)
		}
	}

	return nil
}

// Метод Enabled нужен для проверки, активен ли этот Handler
func (h *KafkaLogHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return true // Всегда включен
}

// Метод WithAttrs нужен для поддержки атрибутов (доп. данных в логах)
func (h *KafkaLogHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &KafkaLogHandler{
		fileHandler: h.fileHandler.WithAttrs(attrs),
		producer:    h.producer,
		topic:       h.topic,
	}
}

// Метод WithGroup нужен для поддержки группировки логов
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

func ensureTopicExists(brokers, topic string) {
	adminClient, err := kafka.NewAdminClient(&kafka.ConfigMap{"bootstrap.servers": brokers})
	if err != nil {
		log.Fatal("Ошибка создания AdminClient:", err.Error())
	}
	defer adminClient.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	results, err := adminClient.CreateTopics(
		ctx,
		[]kafka.TopicSpecification{{
			Topic:             topic,
			NumPartitions:     3, // Количество партиций
			ReplicationFactor: 1, // Репликация (поставь 2-3 для продакшена)
		}},
		kafka.SetAdminOperationTimeout(5*time.Second),
	)
	if err != nil {
		log.Fatal("Ошибка создания топика:", err.Error())
	}

	for _, result := range results {
		if result.Error.Code() == kafka.ErrTopicAlreadyExists {
			fmt.Println("Топик уже существует:", topic)
		} else if result.Error.Code() != kafka.ErrNoError {
			fmt.Println("Ошибка создания топика:", result.Error)
		} else {
			fmt.Println("Топик создан:", topic)
		}
	}
}

// Создаем новый логгер
func NewLogger() *slog.Logger {
	opts := slog.HandlerOptions{Level: slog.LevelDebug}

	// Открываем файл для логов
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal("Ошибка открытия файла лога:", err)
	}

	// Создаем обычный логгер для файла
	fileHandler := slog.NewTextHandler(file, &opts)

	// Добавляем поддержку Kafka
	ensureTopicExists("38.242.212.205:9092", "logs")
	kafkaHandler := NewKafkaLogHandler(fileHandler, "38.242.212.205:9092", "logs")

	log.Println("connected to kafka")

	// Создаем общий логгер
	return slog.New(kafkaHandler)
}
