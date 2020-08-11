package worker_test

import (
	"testing"

	"github.com/GeekTree0101/iGospy/worker"

	"github.com/GeekTree0101/iGospy/model"
)

func Test_Build_Presenter_Spy_Object(t *testing.T) {
	// given
	expectedOutput := `var presentReloadCalled: Int = 0
var presentReloadRes: Feed.Reload.Res?
func presentReload(res: Feed.Reload.Res) {
  self.presentReloadCalled += 1
  self.presentReloadRes = res
}

var presentNextCalled: Int = 0
var presentNextRes: Feed.Next.Res?
func presentNext(res: Feed.Next.Res) {
  self.presentNextCalled += 1
  self.presentNextRes = res
}

`

	node := model.Node{
		Type: model.Title,
		Name: "Feed",
		Children: []model.Node{
			model.Node{
				Type: model.Context,
				Name: "Reload",
				Children: []model.Node{
					model.Node{
						Type:     model.Behavior,
						Name:     "Req",
						Children: nil,
					},
					model.Node{
						Type:     model.Behavior,
						Name:     "Res",
						Children: nil,
					},
					model.Node{
						Type:     model.Behavior,
						Name:     "ViewModel",
						Children: nil,
					},
				},
			},
			model.Node{
				Type: model.Context,
				Name: "Next",
				Children: []model.Node{
					model.Node{
						Type:     model.Behavior,
						Name:     "Req",
						Children: nil,
					},
					model.Node{
						Type:     model.Behavior,
						Name:     "Res",
						Children: nil,
					},
					model.Node{
						Type:     model.Behavior,
						Name:     "ViewModel",
						Children: nil,
					},
				},
			},
		},
	}

	b := worker.NewBuilder(node)

	// when
	out, err := b.GetPresenter()

	// then
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	if out != expectedOutput {
		t.Errorf("unexpected output\n [output]\n%s\n [expected]\n%s\n", out, expectedOutput)
	}
}

func Test_Build_Displayer_Spy_Object(t *testing.T) {
	// given
	expectedOutput := `var displayReloadViewModel: Feed.Reload.ViewModel?
func displayReload(viewModel: Feed.Reload.ViewModel) {
  self.displayReloadViewModel = viewModel
}

var displayNextViewModel: Feed.Next.ViewModel?
func displayNext(viewModel: Feed.Next.ViewModel) {
  self.displayNextViewModel = viewModel
}

`

	node := model.Node{
		Type: model.Title,
		Name: "Feed",
		Children: []model.Node{
			model.Node{
				Type: model.Context,
				Name: "Reload",
				Children: []model.Node{
					model.Node{
						Type:     model.Behavior,
						Name:     "Req",
						Children: nil,
					},
					model.Node{
						Type:     model.Behavior,
						Name:     "Res",
						Children: nil,
					},
					model.Node{
						Type:     model.Behavior,
						Name:     "ViewModel",
						Children: nil,
					},
				},
			},
			model.Node{
				Type: model.Context,
				Name: "Next",
				Children: []model.Node{
					model.Node{
						Type:     model.Behavior,
						Name:     "Req",
						Children: nil,
					},
					model.Node{
						Type:     model.Behavior,
						Name:     "Res",
						Children: nil,
					},
					model.Node{
						Type:     model.Behavior,
						Name:     "ViewModel",
						Children: nil,
					},
				},
			},
		},
	}

	b := worker.NewBuilder(node)

	// when
	out, err := b.GetDisplayer()

	// then
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	if out != expectedOutput {
		t.Errorf("unexpected output\n [output]\n%s\n [expected]\n%s\n", out, expectedOutput)
	}
}
