<script>
  import { onMount } from 'svelte';
  import { fade } from 'svelte/transition';

  // --- API Base URL ---
  const API_URL = 'http://localhost:7284';
  // const API_URL = 'https://api-gym.alecro.click'; // for deployment

  // --- Muscles ---
  let muscles = [
    'chest', 'back', 'shoulders', 'arms', 'biceps', 'triceps',
    'forearms', 'abs', 'obliques', 'legs', 'quads', 'hamstrings',
    'glutes', 'calves', 'traps', 'neck',
    'Lower_back', 'adductors', 'abductors', 'lats', 'delts', 'core' // ✅ newly added
  ];

  // --- Icons ---
  const muscleIcons = {
    chest: '💪', back: '🦾', shoulders: '🏋️‍♂️', arms: '🤜', biceps: '💪',
    triceps: '🤛', forearms: '✊', abs: '🧘‍♂️', obliques: '🧘‍♀️', legs: '🦵',
    quads: '🚴‍♂️', hamstrings: '🏃‍♂️', glutes: '🍑', calves: '🦿',
    traps: '🐍', neck: '🧍‍♂️',
    Lower_back: '🪵', adductors: '🦿', abductors: '🦵', lats: '🪶', delts: '💪', core: '🔥' // ✅ new icons
  };

  // --- State ---
  let searchTerm = '';
  let selectedMuscle = '';
  let exercises = [];
  let allExercises = [];
  let loading = false;
  let error = '';
  let expandedExercise = null;

  // --- Computed search ---
  $: filteredResults = (() => {
    const term = searchTerm.toLowerCase().trim();
    if (!term) return { muscles, exercises: [] };

    const matchedMuscles = muscles.filter(m => m.toLowerCase().includes(term));
    const matchedExercises = allExercises.filter(ex =>
      ex.name.toLowerCase().includes(term)
    );

    return { muscles: matchedMuscles, exercises: matchedExercises };
  })();

  // --- Toggle expand/collapse ---
  function toggleExercise(name) {
    expandedExercise = expandedExercise === name ? null : name;
  }

  // --- Clear search ---
  function clearSearch() {
    searchTerm = '';
    selectedMuscle = '';
    exercises = [];
    expandedExercise = null;
    const el = document.getElementById('muscle-search');
    el && el.focus();
  }

  // --- Fetch exercises for a muscle ---
  async function getExercises(muscle) {
    if (!muscle) return;

    if (selectedMuscle === muscle) {
      selectedMuscle = '';
      exercises = [];
      expandedExercise = null;
      return;
    }

    selectedMuscle = muscle;
    exercises = [];
    error = '';
    loading = true;

    try {
      const res = await fetch(`${API_URL}/exercise?muscle=${encodeURIComponent(muscle.toLowerCase())}`);
      if (!res.ok) throw new Error('Muscle not found or API error');
      const data = await res.json();

      // Ensure unique IDs to prevent duplicate key errors
      exercises = (data.exercises || []).map((e, i) => ({
        id: `${muscle}-${i}`,
        name: e.name,
        description: e.description || '',
        steps: e.steps || []
      }));
    } catch (err) {
      error = err.message || 'Failed to fetch exercises';
      selectedMuscle = '';
    } finally {
      loading = false;
    }
  }

  // --- Prefetch all exercises for global search ---
  onMount(async () => {
    loading = true;
    try {
      const responses = await Promise.all(
        muscles.map(m =>
          fetch(`${API_URL}/exercise?muscle=${encodeURIComponent(m)}`)
            .then(r => (r.ok ? r.json() : null))
            .catch(() => null)
        )
      );

      allExercises = responses.flatMap((data, i) => {
        if (data && data.exercises) {
          return data.exercises.map((e, idx) => ({
            id: `${muscles[i]}-${idx}`,
            name: e.name,
            description: e.description || '',
            steps: e.steps || [],
            muscle: muscles[i]
          }));
        }
        return [];
      });

      console.log(`✅ Loaded ${allExercises.length} exercises`);
    } catch (err) {
      console.warn('Prefetch failed:', err);
    } finally {
      loading = false;
      searchTerm = searchTerm; // trigger refresh
    }
  });
</script>

<!-- =============================== -->
<!--           FRONTEND UI           -->
<!-- =============================== -->

<div class="container my-4">
  <h1 class="h4 mb-3">Find Exercises</h1>

  <!-- Search -->
  <div class="mb-3">
    <div class="input-group">
      <input
        id="muscle-search"
        class="form-control"
        type="search"
        placeholder="Search for a muscle or exercise..."
        bind:value={searchTerm}
        aria-label="Search muscles or exercises"
      />
      <button class="btn btn-outline-secondary" on:click={clearSearch}>Clear</button>
    </div>
    <div class="form-text mt-1">
      Type an exercise (like “push-up”) or a muscle (like “chest”).
    </div>
  </div>

  {#if error}
    <div class="alert alert-danger" transition:fade>{error}</div>
  {/if}

  <div class="row">
    <!-- MUSCLES -->
    <div class="col-lg-4 mb-3">
      <div class="card h-100">
        <div class="card-body">
          <h5 class="card-title">Muscles</h5>
          <p class="text-muted small">Click a muscle to view exercises.</p>

          {#if filteredResults.muscles.length > 0}
            <p class="fw-semibold mb-2">Muscles</p>
            <div class="d-flex flex-wrap gap-2 mb-3">
              {#each filteredResults.muscles as muscle (muscle)}
                <button
                  class="btn btn-outline-primary muscle-chip d-flex align-items-center"
                  class:selected={selectedMuscle === muscle}
                  on:click={() => getExercises(muscle)}>
                  {muscleIcons[muscle] || '💪'}
                  <span class="ms-1 text-capitalize">{muscle}</span>
                </button>
              {/each}
            </div>
          {/if}

          {#if filteredResults.exercises.length > 0}
            <p class="fw-semibold mb-2">Exercises</p>
            <div class="list-group small">
              {#each filteredResults.exercises as ex (ex.id)}
                <button
                  class="list-group-item list-group-item-action text-start"
                  on:click={async () => {
                    await getExercises(ex.muscle);
                    expandedExercise = ex.name;
                  }}>
                  {ex.name} <span class="text-muted">({ex.muscle})</span>
                </button>
              {/each}
            </div>
          {/if}

          {#if filteredResults.muscles.length === 0 && filteredResults.exercises.length === 0 && searchTerm}
            <div class="text-muted small">No results found for “{searchTerm}”</div>
          {/if}
        </div>

        <div class="card-footer d-flex justify-content-between align-items-center">
          <small class="text-muted">{muscles.length} muscles</small>
          <div>
            <button
              class="btn btn-sm btn-outline-secondary me-1"
              on:click={clearSearch}>
              Reset
            </button>
            <button
              class="btn btn-sm btn-primary"
              disabled={loading || !filteredResults.muscles.length}
              on:click={() => getExercises(filteredResults.muscles[0])}>
              Quick Load
            </button>
          </div>
        </div>
      </div>
    </div>

    <!-- EXERCISES -->
    <div class="col-lg-8">
      <div class="d-flex justify-content-between align-items-center mb-2">
        <div>
          {#if selectedMuscle}
            <h5 class="mb-0">Exercises for {selectedMuscle}</h5>
            <small class="text-muted">
              {exercises.length} exercise{exercises.length === 1 ? '' : 's'}
            </small>
          {:else}
            <h5 class="text-muted mb-0">Select a muscle or search an exercise</h5>
          {/if}
        </div>

        {#if loading}
          <div class="spinner-border text-primary" role="status"></div>
        {/if}
      </div>

      {#if selectedMuscle && exercises.length > 0}
        <div class="row g-3">
          {#each exercises as ex (ex.id)}
            <div>
              <article class="card" on:click={() => toggleExercise(ex.name)}>
                <div class="card-body">
                  <h6 class="fw-bold text-capitalize mb-1">
                    {ex.name}
                    <span class="float-end small text-muted">
                      {expandedExercise === ex.name ? '▲' : '▼'}
                    </span>
                  </h6>

                  {#if expandedExercise === ex.name}
                    <div class="expanded-content enter">
                      {#if ex.description}
                        <p class="card-text small text-muted mt-2">{ex.description}</p>
                      {/if}
                      {#if ex.steps?.length > 0}
                        <ol class="small ps-3 mb-0">
                          {#each ex.steps as step}
                            <li>{step}</li>
                          {/each}
                        </ol>
                      {/if}
                    </div>
                  {/if}
                </div>
              </article>
            </div>
          {/each}
        </div>
      {/if}
    </div>
  </div>
</div>

<style>
  h1 { font-weight: 600; }

  .muscle-chip {
    padding: .35rem .6rem;
    border-radius: .5rem;
    text-transform: capitalize;
    transition: transform .08s ease, box-shadow .12s ease;
  }
  .muscle-chip:hover { transform: translateY(-2px); }
  .muscle-chip.selected {
    background-color: var(--bs-primary);
    color: #fff;
  }

  @media (max-width: 767px) {
    .muscle-chip { flex: 1 0 48%; }
  }

  .row.g-3 {
    display: grid;
    grid-template-columns: repeat(auto-fill, minmax(260px, 1fr));
    gap: 1rem;
    align-items: start;
  }

  .card {
    border-radius: .6rem;
    cursor: pointer;
    transition: transform .1s ease, box-shadow .15s ease;
    overflow: hidden;
  }
  .card:hover {
    transform: translateY(-3px);
    box-shadow: 0 6px 15px rgba(0,0,0,0.08);
  }

  .expanded-content {
    overflow: hidden;
    transition: max-height 0.35s ease, opacity 0.25s ease;
    max-height: 0;
    opacity: 0;
  }
  .expanded-content.enter {
    max-height: 1000px;
    opacity: 1;
    margin-top: .5rem;
  }
</style>
