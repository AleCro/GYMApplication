<script>
	export let data = {};
	let userSession = data?.user?.session
	import { onMount } from 'svelte';

	let goals = [];
	let newGoal = '';
	let newStep = {};
	let loading = false;
	let error = '';

	onMount(async () => {
		await fetchGoals();
	});

	async function fetchGoals() {
		try {
			loading = true;
			const res = await fetch('/goals', {
				method: 'POST',
				headers: { 'Content-Type': 'application/json' },
				credentials: 'include',
				body: JSON.stringify({ endpoint: 'getgoals', data: { session: userSession } })
			});
			if (!res.ok) throw new Error('Failed to fetch goals');
			const data = await res.json();
			goals = data || [];
		} catch (err) {
			error = err.message;
		} finally {
			loading = false;
		}
	}

	async function addGoal() {
		if (!newGoal.trim()) return;
		try {
			const res = await fetch('/goals', {
				method: 'POST',
				headers: { 'Content-Type': 'application/json' },
				credentials: 'include',
				body: JSON.stringify({ 
					endpoint: 'addgoal', 
					"data": {
						session: userSession,
						title: newGoal.trim() 
					}
				})
			});
			if (!res.ok) throw new Error('Failed to add goal');
			newGoal = '';
			await fetchGoals();
		} catch (err) {
			error = err.message;
		}
	}

	async function addStep(goalId, stepText) {
		if (!stepText.trim()) return;
		try {
			const res = await fetch('/goals', {
				method: 'POST',
				headers: { 'Content-Type': 'application/json' },
				credentials: 'include',
				body: JSON.stringify({
					endpoint: 'updategoal',
					data: {
						session: userSession,
						id: goalId,
						action: 'addStep',
						step: stepText.trim()
					}
				})
			});
			if (!res.ok) throw new Error('Failed to add step');
			newStep[goalId] = '';
			await fetchGoals();
		} catch (err) {
			error = err.message;
		}
	}

	async function toggleGoal(goalId) {
		try {
			const res = await fetch('/goals', {
				method: 'POST',
				headers: { 'Content-Type': 'application/json' },
				credentials: 'include',
				body: JSON.stringify({
					endpoint: 'updategoal',
					data: {
						session: userSession,
						id: goalId,
						action: 'toggleGoal'
					}
				})
			});
			if (!res.ok) throw new Error('Failed to toggle goal');
			await fetchGoals();
		} catch (err) {
			error = err.message;
		}
	}

	async function toggleStep(goalId, stepIndex) {
		try {
			const res = await fetch('/goals', {
				method: 'POST',
				headers: { 'Content-Type': 'application/json' },
				credentials: 'include',
				body: JSON.stringify({
					endpoint: 'updategoal',
					data: {
						session: userSession,
						id: goalId,
						action: 'toggleStep',
						stepIndex
					}
				})
			});
			if (!res.ok) throw new Error('Failed to toggle step');
			await fetchGoals();
		} catch (err) {
			error = err.message;
		}
	}

	async function deleteGoal(goalIndex) { // Renamed goalId to goalIndex for clarity
		try {
			const res = await fetch('/goals', {
				method: 'POST',
				headers: { 'Content-Type': 'application/json' },
				credentials: 'include',
				body: JSON.stringify({ 
					endpoint: 'deletegoal', 
					data: { 
						session: userSession, 
						i: goalIndex // **UPDATED FIELD NAME FROM 'id' TO 'i'**
					} 
				})
			});
			if (!res.ok) throw new Error('Failed to delete goal');
			await fetchGoals();
		} catch (err) {
			error = err.message;
		}
	}
</script>

<div class="goals-page">
	<header>
		<h1>🎯 My Goals</h1>
		<p>Break down your ambitions into small, achievable steps.</p>
	</header>

	<section class="goal-create">
		<input
			type="text"
			placeholder="Set a new goal..."
			bind:value={newGoal}
			on:keydown={(e) => e.key === 'Enter' && addGoal()}
		/>
		<button on:click={addGoal} disabled={loading}>Add Goal</button>
	</section>

	<section class="goal-list">
		{#if loading}
			<p class="empty">Loading...</p>
		{:else if error}
			<p class="empty">{error}</p>
		{:else if goals.length === 0}
			<p class="empty">No goals yet — start by adding one ✨</p>
		{:else}
			<div class="goal-grid">
				{#each goals as goal, i} <div class="goal-card">
						<div class="goal-header">
							<h2 class:done={goal.done} on:click={() => toggleGoal(goal.id)}>
								{goal.title}
							</h2>
							<button class="delete" on:click={() => deleteGoal(i)}>×</button> </div>

						<div class="steps">
							{#each goal.steps as step, i}
								<div class="step" on:click={() => toggleStep(goal.id, i)}>
									<input type="checkbox" checked={step.done} />
									<span class:done={step.done}>{step.text}</span>
								</div>
							{/each}

							<div class="add-step">
								<input
									type="text"
									placeholder="Add step..."
									bind:value={newStep[goal.id]}
									on:keydown={(e) => e.key === 'Enter' && addStep(goal.id, newStep[goal.id])}
								/>
								<button on:click={() => addStep(goal.id, newStep[goal.id])}>+</button>
							</div>
						</div>
					</div>
				{/each}
			</div>
		{/if}
	</section>
</div>

<style>
	:root {
		--blue: #0d6efd;
		--blue-light: #4da3ff;
		--text: #e6eef6;
		--muted: #adb5bd;
		--bg-dark: #0b0f14;
		--card-bg: #1a1f25;
		--border: #343a40;
		--radius: 14px;
	}

	body {
		background: radial-gradient(circle at top, #1a1f25 0%, #0b0f14 100%);
		color: var(--text);
		font-family: 'Inter', system-ui, sans-serif;
	}

	.goals-page {
		max-width: 1000px;
		margin: 2rem auto;
		padding: 1rem;
	}

	header {
		text-align: center;
		margin-bottom: 2rem;
	}
	header h1 {
		color: var(--blue);
		font-size: 2rem;
		margin-bottom: 0.4rem;
	}
	header p {
		color: var(--muted);
	}

	.goal-create {
		display: flex;
		gap: 0.5rem;
		margin-bottom: 1.5rem;
	}
	.goal-create input {
		flex: 1;
		padding: 0.7rem;
		border-radius: var(--radius);
		border: 1px solid var(--border);
		background: var(--card-bg);
		color: var(--text);
	}
	.goal-create button {
		background-color: var(--blue);
		border: none;
		padding: 0.7rem 1.2rem;
		border-radius: var(--radius);
		color: #fff;
		font-weight: 600;
		cursor: pointer;
		transition: transform 0.2s, box-shadow 0.2s, background-color 0.2s;
	}
	.goal-create button:hover {
		background-color: var(--blue-light);
		transform: translateY(-1px);
		box-shadow: 0 4px 15px rgba(13, 110, 253, 0.3);
	}

	.goal-list {
		margin-top: 1rem;
	}
	.goal-grid {
		display: grid;
		grid-template-columns: repeat(auto-fit, minmax(260px, 1fr));
		gap: 1rem;
	}

	.goal-card {
		background: var(--card-bg);
		border: 1px solid var(--border);
		border-radius: var(--radius);
		padding: 1rem;
		box-shadow: 0 4px 12px rgba(0, 0, 0, 0.3);
		transition: transform 0.1s ease, box-shadow 0.15s ease;
		min-height: 200px;
	}
	.goal-card:hover {
		transform: translateY(-2px);
		box-shadow: 0 6px 16px rgba(0, 0, 0, 0.4);
	}

	.goal-header {
		display: flex;
		justify-content: space-between;
		align-items: center;
		margin-bottom: 0.8rem;
	}
	.goal-header h2 {
		margin: 0;
		cursor: pointer;
		font-size: 1.1rem;
		color: var(--blue);
		word-break: break-word;
	}
	.goal-header h2.done {
		text-decoration: line-through;
		color: var(--muted);
	}
	.goal-header .delete {
		background: none;
		border: none;
		color: #dc3545;
		font-size: 1.2rem;
		cursor: pointer;
	}

	.steps {
		display: flex;
		flex-direction: column;
		gap: 0.4rem;
		font-size: 0.9rem;
	}

	.step {
		display: flex;
		align-items: center;
		gap: 0.4rem;
		cursor: pointer;
	}
	.step input {
		cursor: pointer;
	}
	.step span.done {
		text-decoration: line-through;
		color: var(--muted);
	}

	.add-step {
		display: flex;
		gap: 0.4rem;
		margin-top: 0.6rem;
	}
	.add-step input {
		flex: 1;
		padding: 0.45rem;
		border-radius: 8px;
		border: 1px solid var(--border);
		background: #212529;
		color: var(--text);
		font-size: 0.9rem;
	}
	.add-step button {
		padding: 0.4rem 0.8rem;
		border-radius: 8px;
		border: none;
		background-color: var(--blue);
		color: white;
		cursor: pointer;
		font-weight: 700;
		font-size: 1rem;
		transition: transform 0.2s ease, box-shadow 0.2s ease, background-color 0.2s;
	}
	.add-step button:hover {
		background-color: var(--blue-light);
		transform: translateY(-1px);
		box-shadow: 0 4px 10px rgba(13, 110, 253, 0.4);
	}

	.empty {
		text-align: center;
		color: var(--muted);
		font-style: italic;
		margin-top: 2rem;
	}
</style>