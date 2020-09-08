<script>
	let usecase = `
enum Feed {

  enum Reload {

	enum Request {

	}

	enum Response {

		var cards: [Card]
		var error: Error?
	}

	enum ViewModel {

		var error: Error?
	}
  }

  enum Next {

	enum Request {

	}

	struct Response {

		var cards: [Card]
		var error: Error?
	}

	struct ViewModel {

		var changeSet: [Change<CardViewModel>]
		var error: Error?
	}
  }
}
`;

	let displayInteractorInterface = "Ready to make interactor"
	let displayPresenterInterface = "Ready to make presenter";
	let displayDisplayerInterface = "Ready to make displayer";

	let displayInteractor = "Ready to make interactor"
	let displayPresenter = "Ready to make presenter";
	let displayDisplayer = "Ready to make displayer";

	async function didTapMake() {
		const res = await fetch('http://localhost:7777/make', {
			method: 'POST',
			headers: {
    			'Accept': 'application/json',
    			'Content-Type': 'application/json'
  			},
			body: JSON.stringify({
				usecase
			})
		})
		
		const result = await res.json();
		displayInteractor = result.interactor
		displayPresenter = result.presenter
		displayDisplayer  = result.displayer
		displayInteractorInterface = result.interactorInterface
		displayPresenterInterface = result.presenterInterface
		displayDisplayerInterface = result.displayerInterface
	}

</script>

<main>
	<img width=100%, src="https://github.com/GeekTree0101/iGospy/raw/master/res/logo.png" alt="iGospy"/>
	<center>
		<a href="https://www.github.com/Geektree0101">Created by Geektree0101</a>
	</center>
	<div class="field-container">
		<div class="field-container">
			<h1>Usecase</h1>
			<textarea class="before-field" type="text" name="usecase" bind:value={usecase}></textarea>
		</div>
		<button class="make-button" on:click={didTapMake}>Generate</button>
		<div class="field-container">
			<h1>Interactor</h1>
			<h5>Interface</h5>
			<textarea class="after-field" type="text" name="spy" value={displayInteractorInterface} readonly></textarea>
			<h5>Spy</h5>
			<textarea class="after-field" type="text" name="spy" value={displayInteractor} readonly></textarea>
		</div>
		<div class="field-container">
			<h1>Presenter</h1>
			<h5>Interface</h5>
			<textarea class="after-field" type="text" name="spy" value={displayPresenterInterface} readonly></textarea>
			<h5>Spy</h5>
			<textarea class="after-field" type="text" name="spy" value={displayPresenter} readonly></textarea>
		</div>
		<div class="field-container">
			<h1>Display</h1>
			<h5>Interface</h5>
			<textarea class="after-field" type="text" name="spy" value={displayDisplayerInterface} readonly></textarea>
			<h5>Spy</h5>
			<textarea class="after-field" type="text" name="spy" value={displayDisplayer} readonly></textarea>
		</div>
	</div>
</main>

<style>
 .field-container {
	display: flex;
	flex-direction: column;
	justify-content: stretch;
 }
 .before-field {
	 flex-grow: 1.0;
	 min-height: 20em;
 }
 .after-field {
	 flex-grow: 1.0;
	 min-height: 20em;
 }
 .make-button {
	 color: white;
	 background-color: orange;
	 min-height: 40pt;
	 min-width: 120pt;
	 border-radius: 5pt;
	 border-color: white;
	 align-self: center;
 }
</style>