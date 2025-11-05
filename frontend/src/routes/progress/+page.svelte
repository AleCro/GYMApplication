<script>
	import { onMount } from 'svelte';

	let weight = '';
	let message = '';
	let photoFile = null;
	let photoURL = '';
	let progress = [];
	let loading = false;
	let uploading = false;
	let uploadProgress = 0; // NEW: For tracking file upload percentage (0-100)
	let successMsg = '';
	let errorMsg = '';
	let dragActive = false;
	let removing = {};
	
	export let data;
	const userSession = data?.user?.session;
	
	// **NEW:** Unified fetch function using the central proxy
async function callApi(endpoint, payload) {
    const res = await fetch('/api/progress', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({
            endpoint: endpoint,
            // IMPORTANT: The Go endpoint expects 'session' inside the payload for progress/goals
            data: { session: userSession, ...payload } 
        })
    });
    
    if (!res.ok) {
        let errorMessage = `[${endpoint}] Server error: ${res.status} - ${res.statusText}`;
        
        // Attempt to read error body as JSON, but use a fallback if it fails (which caused your original error)
        try {
            const errorBody = await res.clone().json();
            if (errorBody && errorBody.error) {
                errorMessage = `[${endpoint}] Server error: ${res.status} - ${errorBody.error}`;
            } else if (errorBody && errorBody.message) {
                 // Used for some generic handler responses
                errorMessage = `[${endpoint}] Server error: ${res.status} - ${errorBody.message}`;
            }
        } catch (e) {
            // If .json() fails (i.e., Go returned plain text), use the status text and log the issue.
            console.error("Failed to parse error response as JSON:", e);
        }
        
        throw new Error(errorMessage);
    }
    
    // If res.ok is true, we expect successful JSON
    return res.json(); 
}

	// --------------------------------
	// FILE UPLOAD LOGIC (Unchanged - uses `/progress` directly for multipart upload)
	// --------------------------------
	async function uploadPhoto(file) {
		try {
			uploading = true;
			uploadProgress = 0; // Reset progress
			const formData = new FormData();
			formData.append('image', file);

			// **Multipart forms MUST be sent directly to /progress, not the JSON proxy**
			return new Promise((resolve, reject) => {
				const xhr = new XMLHttpRequest();
				xhr.open('POST', `/progress`, true); // <--- Still targets dedicated upload proxy

				xhr.upload.onprogress = (event) => {
					if (event.lengthComputable) {
						uploadProgress = Math.round((event.loaded / event.total) * 100);
					}
				};

				xhr.onload = () => {
					uploading = false;
					if (xhr.status === 200) {
						try {
							const data = JSON.parse(xhr.responseText);
							photoURL = `${data.url}`;
							successMsg = '📸 Photo uploaded successfully!';
							errorMsg = '';
							resolve();
						} catch (e) {
							reject(new Error('Invalid response from server.'));
						}
					} else {
						reject(new Error(`Upload failed with status: ${xhr.status}`));
					}
				};

				xhr.onerror = () => {
					uploading = false;
					reject(new Error('Network error during upload.'));
				};

				xhr.send(formData);
			});
		} catch (err) {
			console.error(err);
			errorMsg = err.message || 'Failed to upload image.';
		} finally {
			uploading = false;
		}
	}

	async function handleFileChange(e) {
		const file = e.target.files[0];
		if (!file) return;
		photoFile = file;
		await uploadPhoto(file);
	}

	async function handleDrop(e) {
		e.preventDefault();
		e.stopPropagation();
		dragActive = false;
		const file = e.dataTransfer.files[0];
		if (file) {
			photoFile = file;
			await uploadPhoto(file);
		}
	}

	function handleDragOver(e) {
		e.preventDefault();
		e.stopPropagation();
		dragActive = true;
	}

	function handleDragLeave(e) {
		e.preventDefault();
		e.stopPropagation();
		dragActive = false;
	}

	// --------------------------------
	// LOAD PROGRESS (MODIFIED)
	// --------------------------------
	async function loadProgress() {
		try {
			// **MODIFIED:** Using unified callApi function
			const data = await callApi('/getprogress', {});
			progress = data;
		} catch (err) {
			console.error('Error loading progress:', err);
		}
	}

	// --------------------------------
	// SUBMIT PROGRESS (MODIFIED)
	// --------------------------------
	async function submitProgress(e) {
		e.preventDefault();
		successMsg = '';
		errorMsg = '';
		loading = true;

		try {
			const payload = {
				weight: parseFloat(weight),
				message,
				photo: photoURL
			};

			// **MODIFIED:** Using unified callApi function
			const data = await callApi('/addprogress', payload);
			
			// Ensure we are handling the backend response structure correctly
			const newEntry = data.progress || data; 
			progress = [newEntry, ...progress];
			successMsg = 'Progress saved 🎉 Keep going!';
			weight = '';
			message = '';
			photoFile = null;
			photoURL = '';
		} catch (err) {
			errorMsg = err.message || 'Failed to save progress';
		} finally {
			loading = false;
		}
	}
    
    // --------------------------------
	// REMOVE PROGRESS (MODIFIED)
	// --------------------------------
	async function removeProgress(id) {
		if (!confirm('Are you sure you want to delete this entry?')) return;
		
		removing = { ...removing, [id]: true };
		successMsg = '';
		errorMsg = '';
		
		try {
			// **MODIFIED:** Using unified callApi function
			// Note: The Go backend DeleteProgressForm expects 'ID' (capitalized in Go struct) 
			// and 'session' (handled in callApi).
			const payload = { id: id };
			await callApi('/removeprogress', payload);

			// Optimistically update the UI by filtering the item out
			progress = progress.filter(entry => entry.id !== id);
			successMsg = 'Entry removed successfully 👋';
		} catch (err) {
			errorMsg = err.message || 'Failed to remove progress entry.';
		} finally {
			removing = { ...removing, [id]: false };
		}
	}


	function scrollToSection(id) {
		const el = document.getElementById(id);
		if (el) el.scrollIntoView({ behavior: 'smooth', block: 'start' });
	}

	onMount(loadProgress);
</script>

<div class="container my-4">
	<header class="text-center mb-4">
		<h1 class="h3 fw-semibold text-primary">Welcome back, Champion 🏆</h1>
		<p class="text-muted">Your consistency builds strength — let’s see your latest progress.</p>
	</header>

	<nav class="d-flex justify-content-center gap-2 mb-4">
		<button class="btn btn-outline-primary" on:click={() => scrollToSection('tracker')}>Track</button>
		<button class="btn btn-outline-primary" on:click={() => scrollToSection('history')}>Journey</button>
		<a href="/goals" class="btn btn-primary">Goals →</a>
	</nav>

	<section id="tracker" class="card shadow-sm mb-4">
		<div class="card-body">
			<h2 class="h5 text-primary mb-3">📈 Track Today’s Progress</h2>

			<form on:submit|preventDefault={submitProgress}>
				<div class="mb-3">
					<label for="weight" class="form-label fw-semibold">Weight (lbs)</label>
					<input
						id="weight"
						type="number"
						class="form-control"
						placeholder="e.g., 165"
						bind:value={weight}
						required
					/>
				</div>

				<div class="mb-3">
					<label for="message" class="form-label fw-semibold">Note</label>
					<textarea
						id="message"
						class="form-control"
						rows="3"
						placeholder="How are you feeling today?"
						bind:value={message}
					></textarea>
				</div>

				<div
					class="upload-box mb-3"
					on:drop={handleDrop}
					on:dragover={handleDragOver}
					on:dragleave={handleDragLeave}
					class:active={dragActive}
					on:click={() => { if (!uploading) document.getElementById('photo').click() }}
				>
					{#if photoURL}
						<img src={photoURL} alt="Preview" class="preview" />
					{:else if uploading}
						<div class="w-100 p-3 text-center">
							<p class="upload-text mb-2">Uploading photo...</p>
							<div class="progress-container">
								<div class="progress-bar" style="width: {uploadProgress}%;"></div>
							</div>
							<p class="text-sm mt-1">{uploadProgress}%</p>
						</div>
					{:else}
						<p class="upload-text">📸 Add or drag image</p>
					{/if}
					<input
						id="photo"
						type="file"
						accept="image/*"
						on:change={handleFileChange}
						class="hidden-input"
					/>
				</div>

				<div class="d-flex gap-2">
					<button type="submit" class="btn btn-primary" disabled={loading || uploading}>
						{#if loading}
							<span class="spinner-border spinner-border-sm"></span> Saving...
						{:else}
							💾 Save Progress
						{/if}
					</button>

					<label class="btn btn-outline-light mb-0" for="photo">Choose File</label>
				</div>

				{#if successMsg}<div class="alert alert-success mt-3">{successMsg}</div>{/if}
				{#if errorMsg}<div class="alert alert-danger mt-3">{errorMsg}</div>{/if}
			</form>
		</div>
	</section>

	<section id="history" class="card shadow-sm">
		<div class="card-body">
			<h2 class="h5 text-primary mb-3">🕒 Your Journey</h2>
			{#if progress.length === 0}
				<p class="text-muted fst-italic text-center">
					No entries yet — your journey starts here 💪
				</p>
			{:else}
				<div class="row g-3">
					{#each progress as entry (entry.id)}
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
										on:click={() => removeProgress(entry.id)}
										disabled={removing[entry.id]}
									>
										{#if removing[entry.id]}
											Removing...
										{:else}
											🗑️ Remove
										{/if}
									</button>
									</div>
							</div>
						</div>
					{/each}
				</div>
			{/if}
		</div>
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
		position: relative;
		border: 2px dashed rgba(255, 255, 255, 0.2);
		border-radius: 12px;
		text-align: center;
		padding: 2rem;
		cursor: pointer;
		background: rgba(255, 255, 255, 0.03);
		transition: all 0.2s ease;
		min-height: 220px;
		display: flex;
		align-items: center;
		justify-content: center;
	}
	.upload-box:hover,
	.upload-box.active {
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
		max-height: 300px;
		object-fit: cover;
		border-radius: 10px;
		box-shadow: 0 2px 10px rgba(0, 0, 0, 0.4);
	}
    
    /* **NEW PROGRESS BAR STYLES** */
	.progress-container {
		width: 100%;
		background-color: rgba(255, 255, 255, 0.1);
		border-radius: 6px;
		overflow: hidden;
		height: 10px;
		max-width: 250px;
		margin: 0 auto;
	}
	.progress-bar {
		height: 100%;
		background-color: #0d6efd; /* Primary blue */
		transition: width 0.3s ease;
	}
    
    /* **NEW BUTTON STYLES (DANGER/REMOVE)** */
    .btn-danger {
        --bs-btn-bg: #dc3545;
        --bs-btn-border-color: #dc3545;
        --bs-btn-hover-bg: #bb2d3b;
        --bs-btn-hover-border-color: #bb2d3b;
        --bs-btn-active-bg: #bb2d3b;
        --bs-btn-active-border-color: #bb2d3b;
        --bs-btn-disabled-bg: #dc3545;
    }
    .card-img-top {
        object-fit: cover;
        height: 150px; /* fixed height for consistent card size */
    }
</style>