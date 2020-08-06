<script>
	let usecase = `
enum WriteForm {

  enum Post {

    struct Req {
		var isOn: Bool
	}

	struct Res {
		var error: Error?
	}

	struct ViewModel {
		var result: Int
	}
  }

  enum Edit {

	struct Request {
		var isOn: Bool
	}

	struct Response {
		var error: Error?
	}

	struct ViewModel {
		var result: Int
	}
  }

}
	`;
	let displayPresenter = "Ready to make presenter";
	let displayDisplayer = "Ready to make displayer";

	async function didTapMake() {
		const res = await fetch('http://localhost:7777/make', {
			method: 'POST',
			body: JSON.stringify({
				usecase
			})
		})
		
		const result = await res.json();
		displayPresenter = result.presenter
		displayDisplayer  = result.displayer
	}

</script>

<main>
	<h1>iSpygo</h1>
	<div class="field-container">
		<div class="field-container">
			<h2>Input: Usecase Content</h2>
			<textarea class="before-field" type="text" name="usecase" bind:value={usecase}></textarea>
		</div>
		<button class="make-button" on:click={didTapMake}>Generate</button>
		<div class="field-container">
			<h2>Output: Presenter</h2>
			<textarea class="after-field" type="text" name="spy" value={displayPresenter} readonly></textarea>
		</div>
		<div class="field-container">
			<h2>Output: Display</h2>
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