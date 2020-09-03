package commands

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	domainCommands "github.com/markonesgava/take-care/care-taker/domain/commands"
	entities "github.com/markonesgava/take-care/care-taker/domain/entities"
	mock_repository "github.com/markonesgava/take-care/care-taker/domain/repository/mocks"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Test_Expect_error_when_dont_receive_a_name(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	repositoryMock := mock_repository.NewMockTakerRepository(controller)

	handler := newCreateTakeHandler(repositoryMock)

	_, err := handler.Handle(domainCommands.NewCreateTaker(""))

	if err == nil {
		t.Errorf("when doesn't receive a name expect an error")
	}
}

func Test_Expect_no_error_when_receive_a_name(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	repositoryMock := mock_repository.NewMockTakerRepository(controller)

	handler := newCreateTakeHandler(repositoryMock)

	repositoryMock.EXPECT().Add(gomock.Any()).Return(nil)

	_, err := handler.Handle(domainCommands.NewCreateTaker("Houstera"))

	if err != nil {
		t.Errorf("received error when sent taker with name")
	}
}

func Test_Expect_same_name_on_returned_entity(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	repositoryMock := mock_repository.NewMockTakerRepository(controller)

	handler := newCreateTakeHandler(repositoryMock)

	repositoryMock.EXPECT().Add(gomock.Any()).Return(nil)

	taker := domainCommands.NewCreateTaker("Houstera")

	taker1, _ := handler.Handle(taker)

	if taker.Name() != taker1.(*entities.Taker).Name {
		t.Errorf("received entities with diferent names")
	}
}

func Test_Expect_entity_with_id_after_handle(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	repositoryMock := mock_repository.NewMockTakerRepository(controller)

	handler := newCreateTakeHandler(repositoryMock)

	repositoryMock.EXPECT().Add(gomock.Any()).DoAndReturn(func(entity *entities.Taker) error {
		entity.ID = primitive.NewObjectID()
		return nil
	})

	taker := domainCommands.NewCreateTaker("Houstera")

	taker1, _ := handler.Handle(taker)

	if taker1.(*entities.Taker).ID.IsZero() {
		t.Errorf("expected ID after handle and didn't received it")
	}
}

func Test_Expect_same_error_from_repository(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	repositoryMock := mock_repository.NewMockTakerRepository(controller)

	handler := newCreateTakeHandler(repositoryMock)

	err := fmt.Errorf("Erro no repositorio")

	repositoryMock.EXPECT().Add(gomock.Any()).Return(err)

	taker := domainCommands.NewCreateTaker("Houstera")

	_, err1 := handler.Handle(taker)

	if err != err1 {
		t.Errorf("didn't received same error when called the handler")
	}
}
