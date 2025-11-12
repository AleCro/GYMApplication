<script>
  import { onMount } from 'svelte';
  export let data = {};
  let userSession = data?.user?.session;

  let weight = '';
  let message = '';
  let photoBase64 = '';
  let photoPreview = '';
  let progress = [];
  let loading = false;
  let successMsg = '';
  let errorMsg = '';
  let removing = {};

  onMount(loadProgress);

  // ==========================
  // 📸 FILE HANDLER
  // ==========================
  function handleFileChange(e) {
    const file = e.target.files?.[0];
    if (!file) return;

    const reader = new FileReader();
    reader.onload = (ev) => {
      photoBase64 = ev.target.result;
      photoPreview = ev.target.result;
    };
    reader.readAsDataURL(file);
  }

  // ==========================
  // 💾 ADD PROGRESS
  // ==========================
  async function submitProgress() {
    successMsg = '';
    errorMsg = '';

    if (!weight || Number(weight) <= 0) {
      errorMsg = 'Please enter a valid weight.';
      return;
    }

    try {
      loading = true;
      const res = await fetch('/progress', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        credentials: 'include',
        body: JSON.stringify({
          endpoint: 'addprogress',
          data: { session: userSession, weight, message, photo: photoBase64 }
        })
      });

      const data = await res.json();
      if (data.success) {
        successMsg = data.message || 'Progress saved!';
        const entry = data.data || {};

        // ✅ Normalize returned ID and add to local state
        const newEntry = {
          id: entry.id || entry._id || crypto.randomUUID(),
          date: entry.date,
          weight: entry.weight,
          message: entry.message,
          photo: entry.photoBase64 || entry.photo || ''
        };

        progress = [newEntry, ...progress];

        // Clear form fields
        weight = '';
        message = '';
        photoBase64 = '';
        photoPreview = '';

        // Background refresh for consistency
        setTimeout(loadProgress, 800);
      } else {
        errorMsg = data.message || 'Failed to save progress.';
      }
    } catch (err) {
      console.error('[SubmitProgress]', err);
      errorMsg = err.message || 'Network error saving progress.';
    } finally {
      loading = false;
      setTimeout(() => {
        successMsg = '';
        errorMsg = '';
      }, 4000);
    }
  }

  // ==========================
  // 📥 LOAD PROGRESS
  // ==========================
  async function loadProgress() {
    try {
      const res = await fetch('/progress', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        credentials: 'include',
        body: JSON.stringify({
          endpoint: 'getprogress',
          data: { session: userSession }
        })
      });

      const data = await res.json();
      if (Array.isArray(data)) {
        progress = data.map((p) => {
          const id = p.id || p._id || null;
          if (!id) console.warn('[Missing ID]', p);
          return {
            id,
            date: p.date,
            weight: p.weight,
            message: p.message,
            photo: p.photoBase64 || p.photo || ''
          };
        });
        console.log('[Loaded Progress Entries]', progress);
      } else {
        progress = [];
      }
    } catch (err) {
      console.error('[LoadProgress]', err);
      errorMsg = 'Failed to load progress.';
      progress = [];
    }
  }

  // ==========================
  // 🗑️ DELETE PROGRESS
  // ==========================
  async function removeProgress(entryId, index) {
    if (!entryId) {
      alert('Cannot delete: missing ID');
      console.error('[RemoveProgress] Missing ID for entry:', progress[index]);
      return;
    }

    if (!confirm('Delete this progress entry?')) return;

    try {
      removing[index] = true;

      const res = await fetch('/progress', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        credentials: 'include',
        body: JSON.stringify({
          endpoint: 'deleteprogress',
          data: { session: userSession, id: entryId }
        })
      });

      const data = await res.json();
      if (data.success) {
        // ✅ Instant front-end removal
        progress.splice(index, 1);
        progress = [...progress];
        successMsg = data.message || 'Entry deleted successfully.';
      } else {
        errorMsg = data.message || 'Failed to delete progress.';
      }
    } catch (err) {
      console.error('[RemoveProgress]', err);
      errorMsg = err.message || 'Error deleting progress.';
    } finally {
      removing[index] = false;
      setTimeout(() => {
        successMsg = '';
        errorMsg = '';
      }, 4000);
    }
  }
</script>

<!-- ========================== -->
<!-- MARKUP / UI -->
<!-- ========================== -->
<div class="container my-4">
  <header class="text-center mb-4">
    <h1 class="h3 fw-semibold text-primary">📈 Progress Tracker</h1>
    <p class="text-muted">Track your fitness journey — one entry at a time.</p>
  </header>

  <section id="tracker" class="card shadow-sm mb-4 p-3">
    <form on:submit|preventDefault={submitProgress}>
      <div class="mb-3">
        <label class="form-label fw-semibold">Weight (lbs)</label>
        <input type="number" class="form-control" placeholder="e.g., 165" bind:value={weight} required />
      </div>

      <div class="mb-3">
        <label class="form-label fw-semibold">Note</label>
        <textarea class="form-control" rows="3" placeholder="How are you feeling today?" bind:value={message}></textarea>
      </div>

      <div class="upload-box mb-3" on:click={() => document.getElementById('photo').click()}>
        {#if photoPreview}
          <div class="preview-wrapper">
            <img src={photoPreview} alt="Preview" class="preview" />
            <button
              type="button"
              class="btn btn-sm btn-outline-danger mt-2"
              on:click={(e) => {
                e.stopPropagation();
                photoBase64 = '';
                photoPreview = '';
              }}
            >
              Remove
            </button>
          </div>
        {:else}
          <p class="upload-text">📸 Add or drag image</p>
        {/if}
        <input id="photo" type="file" accept="image/*" on:change={handleFileChange} class="hidden-input" />
      </div>

      <button type="submit" class="btn btn-primary w-100" disabled={loading}>
        {#if loading}
          <span class="spinner-border spinner-border-sm"></span> Saving...
        {:else}
          💾 Save Progress
        {/if}
      </button>

      {#if successMsg}<div class="alert alert-success mt-3">{successMsg}</div>{/if}
      {#if errorMsg}<div class="alert alert-danger mt-3">{errorMsg}</div>{/if}
    </form>
  </section>

  <section id="history" class="card shadow-sm p-3">
    <h2 class="h5 text-primary mb-3">🕒 Your Journey</h2>

    {#if !progress.length}
      <p class="text-muted fst-italic text-center">No entries yet — your journey starts here 💪</p>
    {:else}
      <div class="row g-3">
        {#each progress as entry, i}
          <div class="col-md-4">
            <div class="card h-100 border-primary-subtle shadow-sm">
              {#if entry.photo}
                <img src={entry.photo} alt="Progress photo" class="card-img-top" />
              {/if}
              <div class="card-body d-flex flex-column">
                <h5 class="fw-bold text-primary">{entry.weight} lbs</h5>
                <p class="text-muted small">{entry.message}</p>
                <small class="text-secondary mt-auto">{entry.date}</small>
                <button
                  class="btn btn-sm btn-danger mt-3"
                  on:click={() => removeProgress(entry.id, i)}
                  disabled={removing[i]}
                >
                  {#if removing[i]}Removing...{:else}🗑️ Remove{/if}
                </button>
              </div>
            </div>
          </div>
        {/each}
      </div>
    {/if}
  </section>
</div>

<style>
  body {
    background: radial-gradient(circle at top, #0d1117 0%, #0b0f14 100%);
    color: #e6eef6;
    font-family: system-ui, -apple-system, 'Segoe UI', Roboto, sans-serif;
  }

  .card {
    background: rgba(255, 255, 255, 0.06);
    backdrop-filter: blur(10px);
    border-radius: 0.6rem;
    border: 1px solid rgba(255, 255, 255, 0.08);
  }

  .upload-box {
    border: 2px dashed rgba(255, 255, 255, 0.2);
    border-radius: 12px;
    text-align: center;
    padding: 2rem;
    cursor: pointer;
    background: rgba(255, 255, 255, 0.03);
    transition: all 0.2s ease;
    min-height: 200px;
    display: flex;
    align-items: center;
    justify-content: center;
  }

  .upload-box:hover {
    border-color: #0d6efd;
    background: rgba(13, 110, 253, 0.1);
  }

  .upload-text {
    font-size: 1.1rem;
    color: rgba(230, 238, 246, 0.7);
  }

  .hidden-input {
    display: none;
  }

  .preview {
    max-width: 100%;
    max-height: 280px;
    object-fit: cover;
    border-radius: 10px;
    box-shadow: 0 2px 10px rgba(0, 0, 0, 0.4);
  }

  .preview-wrapper {
    display: flex;
    flex-direction: column;
    align-items: center;
  }
</style>
