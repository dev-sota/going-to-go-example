package invitation

import (
	"github.com/ispec-inc/going-to-go-example/pkg/domain/model"
)

type FindCodeOutput struct {
	Invitation model.Invitation
}

type AddCodeOutput struct {
	Invitation model.Invitation
}
