<script>
	import { isPasswordSecure } from "$lib/passwordCheck";
	import { slide } from "svelte/transition";
	import { userStore } from "$lib/api/stores.js";
    import { onDestroy } from "svelte";
    import { request } from "$lib/api/util.js";
    import { cookies } from "$lib/api/cookie.js";
    import { createRemoteJWKSet } from "jose";

	export let data;
	userStore.set(data?.user ?? null);
	$: user = data.user;
	let unsubscribe = userStore.subscribe((value) => {
		user = value;
	});

	const userInitial = (user.username ?? "G").charAt(0).toUpperCase();

	let oldPassword = "";
	let newPassword = "";
	let confirmNewPassword = "";

	let loading = false;
	let successMessage = "";
	let errorMessage = "";

	$: passwordCheck = isPasswordSecure(newPassword);
	$: passwordsMatch =
		newPassword.length > 0 && newPassword === confirmNewPassword;
	$: isFormValid =
		passwordCheck.secure &&
		passwordsMatch &&
		oldPassword.length > 0 &&
		!loading;

	async function handleChangePassword() {
		if (!isFormValid) {
			errorMessage = "Please correct the errors in the form.";
			return;
		}

		loading = true;
		errorMessage = "";
		successMessage = "";

		request("/user/change-password", "POST", JSON.stringify({
			"password": oldPassword,
			"new-password": newPassword
		}), {
			Authorization: `Bearer ${cookies.token}`
		}).then(res => {
			loading = false;
			if (res?.session) {
				cookies.token = res?.session;
				successMessage = `Changed password`;
				oldPassword = "";
				newPassword = "";
				confirmNewPassword = "";
			} else if (res?.message) {
				errorMessage = `Error resetting password: ${res?.message}`;
			} else {
				errorMessage = "Error resetting password";
			}

			console.log(res);
		}).catch((err) => {
			loading = false;
			errorMessage = err.toString();
		})
	}

	onDestroy(() => {
		unsubscribe();
	});
</script>

<svelte:head>
	<title>My Profile</title>
</svelte:head>

<div class="space-y-8 max-w-2xl mx-auto">
	<section>
		<h1 class="text-2xl font-semibold text-white mb-6">My Profile</h1>
		<div class="flex items-center gap-6 p-6 bg-gray-800 rounded-lg">
			<div
				class="w-24 h-24 rounded-full bg-indigo-600 text-white flex items-center justify-center text-5xl font-medium shrink-0"
			>
				{userInitial}
			</div>
			<div>
				<div class="text-xs text-gray-400 uppercase tracking-wider">
					Username
				</div>
				<div class="text-3xl font-bold text-gray-100">
					{user.username}
				</div>
				{#if user.email}
					<div class="text-sm text-gray-300 mt-1">{user.email}</div>
				{/if}
			</div>
		</div>
	</section>

	<section>
		<h2 class="text-xl font-semibold text-white mb-5">Change Password</h2>
		<form
			on:submit|preventDefault={handleChangePassword}
			class="p-6 bg-gray-800 rounded-lg space-y-6"
		>
			{#if successMessage}
				<div
					class="p-4 text-sm text-green-200 bg-green-900/40 border border-green-700 rounded-md"
					transition:slide
				>
					{successMessage}
				</div>
			{/if}
			{#if errorMessage}
				<div
					class="p-4 text-sm text-red-200 bg-red-900/40 border border-red-700 rounded-md"
					transition:slide
				>
					{errorMessage}
				</div>
			{/if}

			<div>
				<label
					for="old-password"
					class="block text-sm font-medium text-gray-300 mb-2"
					>Old Password</label
				>
				<input
					id="old-password"
					type="password"
					bind:value={oldPassword}
					required
					class="w-full px-3 py-2 bg-gray-900 border border-gray-700 rounded-md placeholder-gray-500 text-white focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500"
				/>
			</div>

			<div>
				<label
					for="new-password"
					class="block text-sm font-medium text-gray-300 mb-2"
					>New Password</label
				>
				<input
					id="new-password"
					type="password"
					bind:value={newPassword}
					required
					aria-invalid={newPassword.length > 0 &&
						!passwordCheck.secure}
					class="w-full px-3 py-2 bg-gray-900 border border-gray-700 rounded-md placeholder-gray-500 text-white focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500"
				/>
			</div>

			{#if newPassword.length > 0 && !passwordCheck.secure}
				<div
					class="p-4 bg-gray-900 border border-gray-700 rounded-md"
					transition:slide
				>
					<ul class="space-y-1 text-sm text-gray-400">
						{#each passwordCheck.issues as issue}
							<li class="flex items-center">
								<svg
									class="w-4 h-4 mr-2 text-red-400 shrink-0"
									fill="none"
									stroke-linecap="round"
									stroke-linejoin="round"
									stroke-width="2"
									viewBox="0 0 24 24"
									stroke="currentColor"
									><path d="M6 18L18 6M6 6l12 12"></path></svg
								>
								<span>{issue}</span>
							</li>
						{/each}
					</ul>
				</div>
			{/if}

			<div>
				<label
					for="confirm-password"
					class="block text-sm font-medium text-gray-300 mb-2"
					>Confirm New Password</label
				>
				<input
					id="confirm-password"
					type="password"
					bind:value={confirmNewPassword}
					required
					aria-invalid={confirmNewPassword.length > 0 &&
						!passwordsMatch}
					class="w-full px-3 py-2 bg-gray-900 border border-gray-700 rounded-md placeholder-gray-500 text-white focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500"
				/>
				{#if confirmNewPassword.length > 0 && !passwordsMatch}
					<p class="mt-2 text-sm text-red-400" transition:slide>
						Passwords do not match.
					</p>
				{/if}
			</div>

			<div class="pt-2">
				<button
					type="submit"
					disabled={!isFormValid || loading}
					class="w-full sm:w-auto px-6 py-2.5 bg-indigo-600 text-white font-semibold rounded-md shadow-md hover:bg-indigo-500 transition duration-200 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-offset-gray-800 focus:ring-indigo-500 disabled:opacity-50 disabled:cursor-not-allowed"
				>
					{#if loading}
						<span>Saving...</span>
					{:else}
						Change Password
					{/if}
				</button>
			</div>
		</form>
	</section>
</div>
