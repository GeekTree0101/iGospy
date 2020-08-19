package worker_test

import (
	"testing"

	"github.com/GeekTree0101/iGospy/worker"
)

func Test_Make_Usecase_With_Valid_And_Simple_Raw_Usecase(t *testing.T) {
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
	result, err := parser.MakeUsecase(rawstr)

	// then
	if err != nil {
		t.Error(err)
		return
	}

	if result.Title != "Feed" {
		t.Errorf("invalid title")
		return
	}

	if len(result.Contexts) != 2 {
		t.Errorf("invalid contexts length")
		return
	}

	if result.Contexts[0] != "Reload" {
		t.Errorf("invalid first context")
		return
	}

	if result.Contexts[1] != "Next" {
		t.Errorf("invalid second context")
		return
	}
}

func Test_Make_Usecase_With_Valid_And_Complicated_Raw_Usecase(t *testing.T) {
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
	result, err := parser.MakeUsecase(rawstr)

	// then
	if err != nil {
		t.Error(err)
		return
	}

	if result.Title != "Feed" {
		t.Errorf("invalid title")
		return
	}

	if len(result.Contexts) != 2 {
		t.Errorf("invalid contexts length")
		return
	}

	if result.Contexts[0] != "Reload" {
		t.Errorf("invalid first context")
		return
	}

	if result.Contexts[1] != "Next" {
		t.Errorf("invalid second context")
		return
	}
}

func Test_Make_Usecase_With_RealData(t *testing.T) {
	rawstr := "	enum Feed {↵↵		enum Reload {↵↵			enum Req {↵↵			}↵↵			enum Res {↵↵				var cards: [Card]↵				var error: Error?↵			}↵↵			enum ViewModel {↵↵				var error: Error?↵			}↵		}↵↵		enum Next {↵↵			enum Req {↵↵			}↵↵			struct Res {↵↵				var cards: [Card]↵				var error: Error?↵			}↵↵			struct ViewModel {↵↵				var changeSet: [Change<CardViewModel>]↵				var error: Error?↵			}↵		}↵	}"

	parser := worker.NewParser()

	// when
	result, err := parser.MakeUsecase(rawstr)

	// then
	if err != nil {
		t.Error(err)
		return
	}

	if result.Title != "Feed" {
		t.Errorf("invalid title")
		return
	}

	if len(result.Contexts) != 2 {
		t.Errorf("invalid contexts length")
		return
	}

	if result.Contexts[0] != "Reload" {
		t.Errorf("invalid first context")
		return
	}

	if result.Contexts[1] != "Next" {
		t.Errorf("invalid second context")
		return
	}
}
