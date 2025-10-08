<script>
	import { onMount } from 'svelte';

	export let data;

	let session = data.user.session;
	let notes = [...(data.user.notes || [])];

	async function addNote() {
		try {
			const res = await fetch('/notes', {
				method: 'POST',
				headers: { 'Content-Type': 'application/json' },
				body: JSON.stringify({
					note: '',
					session
				})
			});

			const json = await res.text();
			notes = json.notes;
			window.location.reload();
		} catch (err) {
			console.error(err);
		}
	}

	async function updateNote(i, note) {
		try {
			const res = await fetch('/notes', {
				method: 'PATCH',
				headers: { 'Content-Type': 'application/json' },
				body: JSON.stringify({
					note,
					session,
					i
				})
			});

			const json = await res.text();
			notes = json.notes;
			window.location.reload();
		} catch (err) {
			console.error(err);
		}
	}

	async function deleteNote(i) {
		try {
			const res = await fetch('/notes', {
				method: 'DELETE',
				headers: { 'Content-Type': 'application/json' },
				body: JSON.stringify({
					session,
					i
				})
			});
			
			const json = await res.text();
			notes = json.notes;
			window.location.reload();
		} catch (err) {
			console.error(err);
		}
	}
</script>

<div style="margin: 1em;">
	<div>
		<button class="btn btn-success" on:click={addNote}>+ Create Note</button>
	</div>

	<div class="row">
		{#each notes as note, i}
			<div class="col-md-5" style="margin: 1em;">
				<p>Note {i + 1}</p>
				<textarea
					class="form-control"
					bind:value={notes[i]}
					on:blur={() => updateNote(i, notes[i])}
				></textarea>

				<div style="margin-top: 0.5em;">
					<button class="btn btn-primary" on:click={() => updateNote(i, notes[i])}>Save</button>
					<button class="btn btn-danger" on:click={() => deleteNote(i)}>Delete</button>
				</div>
			</div>
		{/each}
	</div>
</div>
