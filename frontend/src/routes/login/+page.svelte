<script>
	import { onMount } from 'svelte';
	import { API_URL } from '$lib/config.js';

	function setCookie(name, value, days) {
		let expires = '';
		if (days) {
			let date = new Date();
			date.setTime(date.getTime() + days * 24 * 60 * 60 * 1000);
			expires = '; expires=' + date.toUTCString();
		}
		document.cookie = name + '=' + (value || '') + expires + '; path=/';
	}

	let username = '';
	let password = '';
	function aleLogin(username, password) {
		fetch(API_URL + '/login', {
			method: 'POST',
			body: JSON.stringify({ username, password })
		})
			.then((res) => res.json())
			.then((res) => {
				if (Object.keys(res) == 0) {
					return;
				}
				setCookie('session', res.session, 7);
				window.location = "/";
			})
			.catch(console.error);
	}
</script>

<div class="d-flex justify-content-center align-items-center" style="margin-top: 10em">
	<div class="login-card">
		<h4 class="text-center mb-4">Login</h4>
		<div>
			<div class="mb-3">
				<label for="username" class="form-label">Username</label>
				<input
					type="text"
					class="form-control"
					id="username"
					placeholder="Enter your username"
					bind:value={username}
				/>
			</div>
			<div class="mb-3">
				<label for="password" class="form-label">Password</label>
				<input
					type="password"
					class="form-control"
					id="password"
					placeholder="Enter your password"
					bind:value={password}
				/>
			</div>
			<button class="btn btn-primary w-100" on:click={aleLogin(username, password)}>Login</button>
		</div>
	</div>
</div>

<style>
	.login-card {
		width: 100%;
		max-width: 360px;
		padding: 2rem;
		border-radius: 1rem;
		box-shadow: 0 4px 12px rgba(188, 11, 215, 0.1);
	}

	/* Gradient primary button */
	.btn-primary {
		background: linear-gradient(135deg, #0d6efd, #1a73e8, #4285f4);
		border: none;
	}

	.btn-primary:hover {
		background: linear-gradient(135deg, #0b5ed7, #1669d6, #357ae8);
	}
</style>
