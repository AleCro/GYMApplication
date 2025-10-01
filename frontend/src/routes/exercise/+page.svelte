<script>
	import { onMount } from 'svelte';

	let muscles = ['chest', 'back', 'legs', 'arms', 'shoulders'];
	let searchTerm = '';
	let selectedMuscle = '';
	let exercises = [];

	const API_URL = 'https://api-gym.alecro.click';

	// Filter muscles based on search input
	$: filteredMuscles = muscles.filter(m =>
		m.toLowerCase().includes(searchTerm.toLowerCase())
	);

	async function getExercises(muscle) {
		selectedMuscle = muscle;
		exercises = [];

		try {
			const res = await fetch(`${API_URL}/exercise?muscle=${muscle}`);
			if (!res.ok) throw new Error('Muscle not found');
			const data = await res.json();
			exercises = data.exercises || [];
		} catch (err) {
			alert(err.message);
			selectedMuscle = '';
		}
	}
</script>

<style>
	input { padding: 5px; margin-bottom: 10px; width: 200px; }
	ul { list-style: none; padding-left: 0; }
	li { margin: 5px 0; cursor: pointer; color: blue; }
</style>

<h1>Exercise List</h1>

<input placeholder="Search for a muscle..." bind:value={searchTerm} />

<ul>
	{#each filteredMuscles as muscle}
		<li role="button" on:click={() => getExercises(muscle)}>{muscle}</li>
	{/each}
</ul>

{#if selectedMuscle}
	<h2>Exercises for {selectedMuscle}</h2>
	<ul>
		{#each exercises as ex}
			<li>{ex}</li>
		{/each}
	</ul>
{/if}
