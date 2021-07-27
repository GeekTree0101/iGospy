package worker_test

import (
	"testing"

	"github.com/GeekTree0101/iGospy/worker"

	"github.com/GeekTree0101/iGospy/model"
)

func Test_Build_Interactor_Spy_Object(t *testing.T) {
	// given
	expectedOutput := `var reloadCalled: Int = 0
var reloadRequest: Feed.Reload.Request?

func reload(request: Feed.Reload.Request) {
  self.reloadCalled += 1
  self.reloadRequest = request
}

var nextCalled: Int = 0
var nextRequest: Feed.Next.Request?

func next(request: Feed.Next.Request) {
  self.nextCalled += 1
  self.nextRequest = request
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

	if out.SpyGroup != expectedOutput {
		t.Errorf("unexpected output\n [output]\n%s\n [expected]\n%s\n", out.SpyGroup, expectedOutput)
	}
}

func Test_Build_Presenter_Spy_Object(t *testing.T) {
	// given
	expectedOutput := `var presentReloadCalled: Int = 0
var presentReloadResponse: Feed.Reload.Response?

func presentReload(response: Feed.Reload.Response) {
  self.presentReloadCalled += 1
  self.presentReloadResponse = response
}

var presentNextCalled: Int = 0
var presentNextResponse: Feed.Next.Response?

func presentNext(response: Feed.Next.Response) {
  self.presentNextCalled += 1
  self.presentNextResponse = response
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

	if out.SpyGroup != expectedOutput {
		t.Errorf("unexpected output\n [output]\n%s\n [expected]\n%s\n", out.SpyGroup, expectedOutput)
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

	if out.SpyGroup != expectedOutput {
		t.Errorf("unexpected output\n [output]\n%s\n [expected]\n%s\n", out.SpyGroup, expectedOutput)
	}
}
