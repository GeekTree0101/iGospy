package worker_test

import (
	"testing"

	"github.com/GeekTree0101/iGospy/worker"

	"github.com/GeekTree0101/iGospy/model"
)

func Test_Build_Interactor_Spy_Object(t *testing.T) {
	// given
	expectedOutput := `var reloadCalled: Int = 0
var reloadReq: Feed.Reload.Request?
func reload(req: Feed.Reload.Request) {
  self.reloadCalled += 1
  self.reloadReq = req
}

var nextCalled: Int = 0
var nextReq: Feed.Next.Request?
func next(req: Feed.Next.Request) {
  self.nextCalled += 1
  self.nextReq = req
}

`

	usecase := model.Usecase{
		Title:    "Feed",
		Contexts: []string{"Reload", "Next"},
	}

	b := worker.NewBuilder(usecase)

	// when
	out, err := b.GetInteractor()

	// then
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	if out != expectedOutput {
		t.Errorf("unexpected output\n [output]\n%s\n [expected]\n%s\n", out, expectedOutput)
	}
}

func Test_Build_Presenter_Spy_Object(t *testing.T) {
	// given
	expectedOutput := `var presentReloadCalled: Int = 0
var presentReloadResponse: Feed.Reload.Response?
func presentReload(res: Feed.Reload.Response) {
  self.presentReloadCalled += 1
  self.presentReloadResponse = res
}

var presentNextCalled: Int = 0
var presentNextResponse: Feed.Next.Response?
func presentNext(res: Feed.Next.Response) {
  self.presentNextCalled += 1
  self.presentNextResponse = res
}

`

	usecase := model.Usecase{
		Title:    "Feed",
		Contexts: []string{"Reload", "Next"},
	}

	b := worker.NewBuilder(usecase)

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

	usecase := model.Usecase{
		Title:    "Feed",
		Contexts: []string{"Reload", "Next"},
	}

	b := worker.NewBuilder(usecase)

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
