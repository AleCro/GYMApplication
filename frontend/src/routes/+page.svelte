<script>
	let name = $state('Alejandro');
</script>

<nav class="navbar navbar-expand-lg bg-body-tertiary">
	<div class="container-fluid">
		<a class="navbar-brand" href="#">AleGYM</a>
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
					<a class="nav-link active" aria-current="page" href="/">Home</a>
				</li>
				<li class="nav-item">
					<a class="nav-link" href="#">Calendar</a>
				</li>
			</ul>
		</div>
	</div>
</nav>

<div style="margin: 1em;">
	<div class="row">
		<div class="col-md-3">
			<input bind:value={name} class="form-control" />
		</div>
		<div class="col-md-9">
			<h1>Welcome back, <span style="color: blue">{name}</span>!</h1>
		</div>
	</div>

	
</div>
