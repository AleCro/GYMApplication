<script>
	let { data } = $props(); 
	let session = data.user.session;
	let notes = $state();

	let save = () => {
		fetch("/", {
			method: "POST",
			body: JSON.stringify({
				notes, session
			})
		}).then(res => res.json()).then(res => {
			console.log(res);
		}).catch(console.error);
	}
</script>

<nav class="navbar navbar-expand-lg bg-body-tertiary">
	<div class="container-fluid">
		<a class="navbar-brand" href="/">AleGYM</a>
		<button
			class="navbar-toggler"
			type="button"
			data-bs-toggle="collapse"
			data-bs-target="#navbarNav"
			aria-controls="navbarNav"
			aria-expanded="false"
			aria-label="Toggle navigation"
		>
			<span class="navbar-toggler-icon"></span>
		</button>
		<div class="collapse navbar-collapse" id="navbarNav">
			<ul class="navbar-nav">
				<li class="nav-item">
					<a class="nav-link active" aria-current="page" href="/">Notes</a>
				</li>
				<li class="nav-item">
					<a class="nav-link" href="/calendar">Calendar</a>
				</li>
			</ul>
		</div>
	</div>
</nav>

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
