<script>
	let { data } = $props(); 
	let session = data.user.session;
	let notes = $state();

	let save = () => {
		fetch("/notes", {
			method: "POST",
			body: JSON.stringify({
				notes, session
			})
		}).then(res => res.json()).then(res => {
			console.log(res);
		}).catch(console.error);
	}
</script>

<div style="margin: 1em;">
	<div class="row">
		<div class="col-md-5">
			<textarea bind:value={notes} class="form-control">{data.user.notes}</textarea>
			<br>
			<button class="btn btn-primary" on:click={save}>Save</button>
			<br>
		</div>
		<div class="col-md-9">
			<h1>Notes: <span style="color: blue">{notes}</span></h1>
		</div>
	</div>
</div>