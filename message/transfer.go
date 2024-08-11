package message

import (
	"encoding/json"

	"github.com/streadway/amqp"
)

type Transfer struct {
	Amount     float64 `json:"amount"`
	TargetUser string  `json:"target_user"`
	Remarks    string  `json:"remarks"`
	TrfSource  string  `json:"trf_source"`
}

type TransferRepo interface {
	Publish(t *Transfer) error
}

type transferRepo struct {
	rmq *amqp.Connection
}

func NewTransferRepo(rmq *amqp.Connection) TransferRepo {
	return &transferRepo{
		rmq,
	}
}

func (rpo *transferRepo) Publish(t *Transfer) error {
	ch, err := rpo.rmq.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	b, err := json.Marshal(t)
	if err != nil {
		return err
	}

	err = ch.Publish(
		"",
		"transfer",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        b,
		},
	)
	if err != nil {
		return err
	}

	return nil
}
