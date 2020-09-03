package commands

import (
	"testing"

	"github.com/golang/mock/gomock"
	domainCommands "github.com/markonesgava/take-care/care-taker/domain/commands"
	mock_repository "github.com/markonesgava/take-care/care-taker/domain/repository/mocks"
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
