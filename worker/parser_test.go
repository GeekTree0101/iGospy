package worker_test

import (
	"testing"

	"github.com/GeekTree0101/iSpygo/model"

	"github.com/GeekTree0101/iSpygo/worker"
)

func Test_Processing_Valid_And_Simple_Usecase(t *testing.T) {
	// given
	var rawstr string = `
	enum Feed {

		enum Reload {

			enum Req {

			}

			enum Res {

				var cards: [Card]
				var error: Error?
			}

			enum ViewModel {

				var error: Error?
			}
		}

		enum Next {

			enum Req {

			}

			struct Res {

				var cards: [Card]
				var error: Error?
			}

			struct ViewModel {

				var changeSet: [Change<CardViewModel>]
				var error: Error?
			}
		}
	}
	`

	parser := worker.NewParser()

	// when
	result, err := parser.Processing(rawstr)

	// then
	if err != nil {
		t.Error(err)
		return
	}

	// title

	if result.Type != model.Title {
		t.Errorf("root is not Title")
		return
	}

	if len(result.Children) != 2 {
		t.Errorf("children count doesn't 2, %d", len(result.Children))
		return
	}

	// context

	if result.Children[0].Type != model.Context {
		t.Errorf("%s type doesn't Context", result.Children[0].Name)
		return
	}

	if result.Children[1].Type != model.Context {
		t.Errorf("%s type doesn't Context", result.Children[1].Name)
		return
	}

	if len(result.Children[0].Children) != 3 {
		t.Errorf("children count doesn't 3, %d", len(result.Children[0].Children))
		return
	}

	// first context

	if result.Children[0].Children[0].Type != model.Behavior {
		t.Errorf("%s type doesn't Behavior", result.Children[0].Children[0].Name)
		return
	}

	if result.Children[0].Children[1].Type != model.Behavior {
		t.Errorf("%s type doesn't Behavior", result.Children[0].Children[1].Name)
		return
	}

	if result.Children[0].Children[2].Type != model.Behavior {
		t.Errorf("%s type doesn't Behavior", result.Children[0].Children[2].Name)
		return
	}

	if len(result.Children[0].Children[0].Children) != 0 {
		t.Errorf("%s children should be zero", result.Children[0].Children[0].Name)
		return
	}

	if len(result.Children[0].Children[1].Children) != 0 {
		t.Errorf("%s children should be zero", result.Children[0].Children[1].Name)
		return
	}

	if len(result.Children[0].Children[2].Children) != 0 {
		t.Errorf("%s children should be zero", result.Children[0].Children[2].Name)
		return
	}

	// second context

	if result.Children[1].Children[0].Type != model.Behavior {
		t.Errorf("%s type doesn't Behavior", result.Children[1].Children[0].Name)
		return
	}

	if result.Children[1].Children[1].Type != model.Behavior {
		t.Errorf("%s type doesn't Behavior", result.Children[1].Children[1].Name)
		return
	}

	if result.Children[1].Children[2].Type != model.Behavior {
		t.Errorf("%s type doesn't Behavior", result.Children[1].Children[2].Name)
		return
	}

	if len(result.Children[1].Children[0].Children) != 0 {
		t.Errorf("%s children should be zero", result.Children[1].Children[0].Name)
		return
	}

	if len(result.Children[1].Children[1].Children) != 0 {
		t.Errorf("%s children should be zero", result.Children[1].Children[1].Name)
		return
	}

	if len(result.Children[1].Children[2].Children) != 0 {
		t.Errorf("%s children should be zero", result.Children[1].Children[2].Name)
		return
	}
}

func Test_Processing_Valid_And_Complicated_Usecase(t *testing.T) {
	// given
	var rawstr string = `
	enum Feed {

		enum Reload {

			enum Req {

			}

			enum Res {

				var cards: [Card]
				var error: Error?
			}

			enum ViewModel {

				var error: Error?
			}
		}

		enum Next {

			enum Req {

			}

			struct Res {

				var cards: [Card]
				var error: Error?
			}

			struct ViewModel {

				var changeSet: [Change<CardViewModel>]
				var error: Error?
			}
		}
	}
	`

	parser := worker.NewParser()

	// when
	result, err := parser.Processing(rawstr)

	// then
	if err != nil {
		t.Error(err)
		return
	}

	// title

	if result.Type != model.Title {
		t.Errorf("root is not Title")
		return
	}

	if len(result.Children) != 2 {
		t.Errorf("children count doesn't 2, %d", len(result.Children))
		return
	}

	// context

	if result.Children[0].Type != model.Context {
		t.Errorf("%s type doesn't Context", result.Children[0].Name)
		return
	}

	if result.Children[1].Type != model.Context {
		t.Errorf("%s type doesn't Context", result.Children[1].Name)
		return
	}

	if len(result.Children[0].Children) != 3 {
		t.Errorf("children count doesn't 3, %d", len(result.Children[0].Children))
		return
	}

	// first context

	if result.Children[0].Children[0].Type != model.Behavior {
		t.Errorf("%s type doesn't Behavior", result.Children[0].Children[0].Name)
		return
	}

	if result.Children[0].Children[1].Type != model.Behavior {
		t.Errorf("%s type doesn't Behavior", result.Children[0].Children[1].Name)
		return
	}

	if result.Children[0].Children[2].Type != model.Behavior {
		t.Errorf("%s type doesn't Behavior", result.Children[0].Children[2].Name)
		return
	}

	if len(result.Children[0].Children[0].Children) != 0 {
		t.Errorf("%s children should be zero", result.Children[0].Children[0].Name)
		return
	}

	if len(result.Children[0].Children[1].Children) != 0 {
		t.Errorf("%s children should be zero", result.Children[0].Children[1].Name)
		return
	}

	if len(result.Children[0].Children[2].Children) != 0 {
		t.Errorf("%s children should be zero", result.Children[0].Children[2].Name)
		return
	}

	// second context

	if result.Children[1].Children[0].Type != model.Behavior {
		t.Errorf("%s type doesn't Behavior", result.Children[1].Children[0].Name)
		return
	}

	if result.Children[1].Children[1].Type != model.Behavior {
		t.Errorf("%s type doesn't Behavior", result.Children[1].Children[1].Name)
		return
	}

	if result.Children[1].Children[2].Type != model.Behavior {
		t.Errorf("%s type doesn't Behavior", result.Children[1].Children[2].Name)
		return
	}

	if len(result.Children[1].Children[0].Children) != 0 {
		t.Errorf("%s children should be zero", result.Children[1].Children[0].Name)
		return
	}

	if len(result.Children[1].Children[1].Children) != 0 {
		t.Errorf("%s children should be zero", result.Children[1].Children[1].Name)
		return
	}

	if len(result.Children[1].Children[2].Children) != 0 {
		t.Errorf("%s children should be zero", result.Children[1].Children[2].Name)
		return
	}
}
