generate-mock:
	mockery --name=IProducer --structname=ProducerMocked --inpackage=true --filename=producer_mock.go --outpkg=queue --dir=./infra/queue