<script>
  import { onMount } from 'svelte';

  export let data;

  // helpers
  const mkId = () =>
    typeof crypto !== 'undefined' && crypto.randomUUID
      ? crypto.randomUUID()
      : `${Date.now()}-${Math.floor(Math.random() * 1e6)}`;

  const safeJson = async (res) => {
    try {
      return await res.json();
    } catch {
      const txt = await res.text().catch(() => '');
      try { return txt ? JSON.parse(txt) : {}; } catch { return {}; }
    }
  };

  // session & notes
  let session = data?.user?.session;
  // notesLocal: array of { id, text }
  let notesLocal = [];

  function mapServerNotes(rawNotes = []) {
    // server may send strings or objects; normalize to { id, text }
    return rawNotes.map((n, idx) => {
      if (typeof n === 'string') return { id: mkId(), text: n };
      if (n && typeof n === 'object') {
        return { id: n.id ?? mkId(), text: n.note ?? n.text ?? String(n) };
      }
      return { id: mkId(), text: String(n) };
    });
  }

  const API = '/notes';

  // fetch notes from server and normalize
  async function fetchNotes() {
   
  }


  // if parent replaces data completely, reflect it (but keep UI stable otherwise)
  $: if (data) {
    const raw = data?.user?.notes || [];
    if (raw.length !== notesLocal.length) {
      notesLocal = mapServerNotes(raw);
    }
  }

  // per-note save state & debounce timers
  let saveState = {}; // id -> { state, message }
  const timers = new Map();

  function setSaveState(id, state, message = '') {
    saveState = { ...saveState, [id]: { state, message } };
  }

  // Add a new note (optimistic push, then POST; then re-fetch canonical list)
  async function addNote() {
    const tempId = mkId();
    notesLocal = [{ id: tempId, text: '' }, ...notesLocal];
    setSaveState(tempId, 'saving');

    try {
      const res = await fetch(API, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ note: '', session })
      });
      if (!res.ok) throw new Error(`Server error ${res.status}`);
      // After a successful mutation, re-fetch canonical notes (robust even if server returns nothing)
      setSaveState(tempId, 'saved');
      // clear saved marker after brief delay
      setTimeout(() => { if (saveState[tempId]?.state === 'saved') setSaveState(tempId, 'idle'); }, 1200);
	  window.location.reload();
    } catch (err) {
      console.error('addNote error:', err);
      setSaveState(tempId, 'error', err.message || 'Failed to add');
      // keep temp note in UI so user can retry/edit
    }
  }

  // Save note by index (PATCH) then re-fetch
  async function saveNoteToServer(idx) {
    const noteObj = notesLocal[idx];
    if (!noteObj) return;
    const id = noteObj.id;
    setSaveState(id, 'saving');

    try {
      const res = await fetch(API, {
        method: 'PATCH',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ note: noteObj.text, session, i: idx })
      });
      if (!res.ok) throw new Error(`Server error ${res.status}`);
      setSaveState(id, 'saved');
      setTimeout(() => { if (saveState[id]?.state === 'saved') setSaveState(id, 'idle'); }, 1200);
    } catch (err) {
      console.error('saveNoteToServer error:', err);
      setSaveState(id, 'error', err.message || 'Save failed');
    }
  }

  // Debounced save when user types
  function scheduleSave(idx, delay = 800) {
    const note = notesLocal[idx];
    if (!note) return;
    const id = note.id;
    if (timers.has(id)) clearTimeout(timers.get(id));
    const t = setTimeout(() => {
      const curIdx = notesLocal.findIndex(n => n.id === id);
      if (curIdx !== -1) saveNoteToServer(curIdx);
      timers.delete(id);
    }, delay);
    timers.set(id, t);
    setSaveState(id, 'saving');
  }

  async function flushSave(idx) {
    const note = notesLocal[idx];
    if (!note) return;
    const id = note.id;
    if (timers.has(id)) {
      clearTimeout(timers.get(id));
      timers.delete(id);
    }
    await saveNoteToServer(idx);
  }

  // Delete note then re-fetch
  async function deleteNote(idx) {
    const note = notesLocal[idx];
    if (!note) return;
    const ok = window.confirm('Delete this note? This cannot be undone.');
    if (!ok) return;

    setSaveState(note.id, 'saving');
    try {
      const res = await fetch(API, {
        method: 'DELETE',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ session, i: idx })
      });
      window.location.reload();
    } catch (err) {
      console.error('deleteNote error:', err);
      setSaveState(note.id, 'error', err.message || 'Delete failed');
    }
  }
</script>

<div class="container my-4">
  <div class="d-flex justify-content-between align-items-center mb-3">
    <h4 class="mb-0">My Notes</h4>
    <div>
      <button class="btn btn-success" on:click={addNote}>
        <i class="bi bi-plus-circle me-1"></i> Create Note
      </button>
    </div>
  </div>

  {#if notesLocal.length === 0}
    <div class="alert alert-info">No notes yet. Click <strong>Create Note</strong> to add one.</div>
  {/if}

  <div class="row g-3">
    {#each notesLocal as note, idx (note.id)}
      <div class="col-12 col-md-6 col-lg-4">
        <div class="card h-100 shadow-sm">
          <div class="card-body d-flex flex-column">
            <div class="d-flex justify-content-between align-items-start mb-2">
              <div class="small text-muted">Note {idx + 1}</div>
              <div class="d-flex align-items-center gap-2">
                {#if saveState[note.id]?.state === 'saving'}
                  <span class="spinner-border spinner-border-sm text-primary" role="status" aria-hidden="true"></span>
                {:else if saveState[note.id]?.state === 'saved'}
                  <span class="badge bg-success">Saved</span>
                {:else if saveState[note.id]?.state === 'error'}
                  <span class="badge bg-danger" title={saveState[note.id].message}>Error</span>
                {:else}
                  <span class="text-muted small">Idle</span>
                {/if}
                <button class="btn btn-sm btn-outline-danger" on:click={() => deleteNote(idx)} aria-label="Delete note">
                  🗑️
                </button>
              </div>
            </div>

            <textarea
              class="form-control mb-2 flex-grow-1"
              bind:value={note.text}
              rows="6"
              on:input={() => scheduleSave(idx)}
              on:blur={() => flushSave(idx)}
            ></textarea>

            <div class="d-flex justify-content-between align-items-center mt-2">
              <div class="small text-muted">---</div>
              <div>
                <button class="btn btn-sm btn-primary me-1" on:click={() => flushSave(idx)}>
                  Save
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
    {/each}
  </div>
</div>

<style>
  .card { border-radius: .6rem; }
  textarea.form-control { resize: vertical; min-height: 6rem; }
  .spinner-border-sm { width: 1rem; height: 1rem; }
  .bi { vertical-align: -.125em; }
  @media (max-width: 575px) {
    .card { font-size: .95rem; }
  }
</style>
