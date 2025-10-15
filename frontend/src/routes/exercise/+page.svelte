<script>
  import { onMount } from 'svelte';
  import { fade, slide } from 'svelte/transition';

  let muscles = ['chest', 'back', 'legs', 'arms', 'shoulders'];
  let searchTerm = '';
  let selectedMuscle = '';
  let exercises = [];
  let loading = false;
  let error = '';

  const API_URL = 'https://api-gym.alecro.click';

  // Filter muscles based on search input
  $: filteredMuscles = muscles.filter(m =>
    m.toLowerCase().includes(searchTerm.toLowerCase())
  );

  // Allow Enter to pick the first filtered muscle
  function handleSearchKeydown(e) {
    if (e.key === 'Enter' && filteredMuscles.length > 0) {
      getExercises(filteredMuscles[0]);
    }
  }

  function clearSearch() {
    searchTerm = '';
    // optional focus -- will work if input has bind:this
    const el = document.getElementById('muscle-search');
    el && el.focus();
  }

  async function getExercises(muscle) {
    if (!muscle) return;
    selectedMuscle = muscle;
    exercises = [];
    error = '';
    loading = true;

    try {
      const res = await fetch(`${API_URL}/exercise?muscle=${encodeURIComponent(muscle)}`);
      if (!res.ok) throw new Error('Muscle not found or API error');
      const data = await res.json();
      // handle a few possible shapes (strings or objects)
      exercises = (data && data.exercises) ? data.exercises.map(x => {
        if (typeof x === 'string') return { name: x };
        if (x && x.name) return x;
        // fall back stringify
        return { name: String(x) };
      }) : [];
    } catch (err) {
      console.error(err);
      error = err.message || 'Failed to fetch exercises';
      selectedMuscle = '';
    } finally {
      loading = false;
    }
  }

  // optional: prefetch the first muscle on mount
  onMount(() => {
    // getExercises(muscles[0]); // uncomment to auto-load first muscle
  });
</script>

<div class="container my-4">
  <h1 class="h4 mb-3">Find Exercises</h1>

  <!-- Search -->
  <div class="mb-3">
    <div class="input-group">
      <input
        id="muscle-search"
        class="form-control"
        type="search"
        placeholder="Search for a muscle (e.g., chest)"
        bind:value={searchTerm}
        on:keydown={handleSearchKeydown}
        aria-label="Search muscles"
      />
      <button class="btn btn-outline-secondary" type="button" on:click={clearSearch} aria-label="Clear search">
        Clear
      </button>
    </div>
    <div class="form-text mt-1">Press Enter to select the top match.</div>
  </div>

  <!-- Error -->
  {#if error}
    <div class="alert alert-danger" role="alert" transition:fade>
      {error}
    </div>
  {/if}

  <div class="row">
    <!-- Muscles column -->
    <div class="col-lg-4 mb-3">
      <div class="card h-100">
        <div class="card-body">
          <h5 class="card-title">Muscles</h5>
          <p class="text-muted small">Click a muscle to view exercises.</p>

          <!-- chips / list -->
          <div class="d-flex flex-wrap gap-2" role="list">
            {#each filteredMuscles as muscle (muscle)}
              <button
                class="btn btn-outline-primary muscle-chip d-flex align-items-center"
                class:selected={selectedMuscle === muscle}
                on:click={() => getExercises(muscle)}
                role="listitem"
                aria-pressed={selectedMuscle === muscle}
              >
                💪
                <span class="text-capitalize">{muscle}</span>
              </button>
            {/each}

            {#if filteredMuscles.length === 0}
              <div class="text-muted small">No muscles match “{searchTerm}”</div>
            {/if}
          </div>
        </div>

        <div class="card-footer d-flex justify-content-between align-items-center">
          <small class="text-muted">{muscles.length} muscles</small>
          <div>
            <button class="btn btn-sm btn-outline-secondary me-1" on:click={() => { searchTerm=''; }}>
              Reset
            </button>
            <button class="btn btn-sm btn-primary" disabled={loading || !filteredMuscles.length} on:click={() => getExercises(filteredMuscles[0])}>
              Quick load
            </button>
          </div>
        </div>
      </div>
    </div>

    <!-- Exercises column -->
    <div class="col-lg-8">
      <div class="d-flex justify-content-between align-items-center mb-2">
        <div>
          <h5 class="mb-0">{selectedMuscle ? `Exercises for ${selectedMuscle}` : 'Select a muscle'}</h5>
          <small class="text-muted">{selectedMuscle ? `${exercises.length} exercise${exercises.length === 1 ? '' : 's'}` : 'No muscle selected'}</small>
        </div>

        {#if loading}
          <div class="spinner-border text-primary" role="status" aria-hidden="true"></div>
        {/if}
      </div>

      <div>
        {#if !selectedMuscle && !loading}
          <div class="card p-3 text-center text-muted">
            Select a muscle to view exercises.
          </div>
        {/if}

        {#if selectedMuscle && !loading}
          {#if exercises.length === 0}
            <div class="alert alert-warning" role="alert" transition:fade>
              No exercises found for <strong class="text-capitalize">{selectedMuscle}</strong>.
            </div>
          {/if}

          <div class="row g-3">
            {#each exercises as ex (ex.name + '-' + Math.random())}
              <div class="col-md-6">
                <article class="card h-100" transition:slide>
                  <div class="card-body">
                    <h6 class="card-title mb-1 text-capitalize">{ex.name}</h6>
                    {#if ex.type}
                      <div class="mb-1"><small class="text-muted">Type: {ex.type}</small></div>
                    {/if}
                    {#if ex.primary}
                      <div><span class="badge bg-secondary text-capitalize">{ex.primary}</span></div>
                    {/if}
                    {#if ex.description}
                      <p class="card-text mt-2 small text-muted">{ex.description}</p>
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
  .muscle-chip.selected,
  .muscle-chip[aria-pressed="true"] {
    background-color: var(--bs-primary);
    color: #fff;
    box-shadow: 0 6px 18px rgba(13,110,253,.12);
  }

  /* small responsive tweak */
  @media (max-width: 767px) {
    .muscle-chip { flex: 1 0 48%; }
  }

  /* card visuals */
  .card { border-radius: .6rem; }
  .card .card-body h6 { font-size: .95rem; }
</style>
