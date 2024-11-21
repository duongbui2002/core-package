package cqrs

type TxRequest interface {
	Request

	isTxRequest()
}
