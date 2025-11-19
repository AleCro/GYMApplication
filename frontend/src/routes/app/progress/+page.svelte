<script>
    import { onMount } from "svelte";
    import { cookies } from "$lib/api/cookie";
    import { request } from "$lib/api/util";

    let progressEntries = [];
    let title = "";
    let description = "";
    let files;
    let editingId = null;
    let uploading = false;

    async function fetchProgress() {
        try {
            progressEntries = await request("/progress", "GET", null, { Authorization: `Bearer ${cookies.token}` });
        } catch (e) {
            console.error(e);
        }
    }

    async function uploadProgress() {
        if (!files || files.length === 0) {
            alert("Please select an image");
            return;
        }
        if (!title) {
            alert("Please enter a title");
            return;
        }

        uploading = true;
        const formData = new FormData();
        formData.append("title", title);
        formData.append("description", description);
        formData.append("image", files[0]);

        try {
            // For FormData, we don't set Content-Type header manually, let browser do it
            await request("/progress", "POST", formData, { Authorization: `Bearer ${cookies.token}` });
            
            title = "";
            description = "";
            files = null;
            // Reset file input
            document.getElementById("fileInput").value = "";
            fetchProgress();
        } catch (e) {
            console.error(e);
            alert("Failed to upload");
        }
        uploading = false;
    }

    async function editProgress(entry) {
        editingId = entry.id;
        title = entry.title;
        description = entry.description;
        // We don't set files because we can't pre-fill file input, and we don't support updating image yet
    }

    function resetForm() {
        editingId = null;
        title = "";
        description = "";
        files = null;
        if (document.getElementById("fileInput")) {
            document.getElementById("fileInput").value = "";
        }
    }

    async function deleteProgress(id) {
        if (!confirm("Are you sure you want to delete this entry?")) return;
        try {
            await request(`/progress/${id}`, "DELETE", null, { Authorization: `Bearer ${cookies.token}` });
            fetchProgress();
        } catch (e) {
            console.error(e);
            alert(`Failed to delete: ${e.message || e}`);
        }
    }

    onMount(() => {
        fetchProgress();
    });
</script>
<div class="space-y-8">
    <h1 class="text-4xl font-bold text-transparent bg-clip-text bg-gradient-to-r from-primary to-secondary">Progress Tracker</h1>

    <div class="glass-card p-8 rounded-2xl">
        <h2 class="text-2xl font-bold text-white mb-6 flex items-center gap-2">
            <span class="text-primary">{editingId ? "✎" : "+"}</span> {editingId ? "Edit Entry" : "Add New Entry"}
        </h2>
        <div class="space-y-6">
            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
                <div class="space-y-2">
                    <label class="text-xs font-bold text-gray-400 uppercase tracking-wider ml-1">Title</label>
                    <input
                        bind:value={title}
                        placeholder="e.g., Week 1 Transformation"
                        class="w-full glass-input px-4 py-3 rounded-lg focus:ring-2 focus:ring-primary transition"
                    />
                </div>
                <div class="space-y-2">
                    <label class="text-xs font-bold text-gray-400 uppercase tracking-wider ml-1">Photo</label>
                    <input
                        type="file"
                        accept="image/*"
                        bind:files
                        id="fileInput"
                        disabled={!!editingId}
                        class="w-full glass-input px-4 py-2.5 rounded-lg focus:ring-2 focus:ring-primary transition file:mr-4 file:py-1 file:px-3 file:rounded-md file:border-0 file:text-xs file:font-semibold file:bg-primary file:text-black hover:file:bg-green-400 cursor-pointer disabled:opacity-50 disabled:cursor-not-allowed"
                    />
                    {#if editingId}
                        <p class="text-xs text-gray-500 ml-1">Photo cannot be changed during edit</p>
                    {/if}
                </div>
            </div>
            <div class="space-y-2">
                <label class="text-xs font-bold text-gray-400 uppercase tracking-wider ml-1">Description</label>
                <textarea
                    bind:value={description}
                    placeholder="How do you feel? Current weight? Personal records?"
                    rows="3"
                    class="w-full glass-input px-4 py-3 rounded-lg focus:ring-2 focus:ring-primary transition resize-none"
                ></textarea>
            </div>
            <div class="flex justify-end gap-3">
                {#if editingId}
                    <button
                        on:click={resetForm}
                        class="px-6 py-3 rounded-lg border border-white/10 hover:bg-white/10 transition text-gray-300"
                    >
                        Cancel
                    </button>
                {/if}
                <button
                    on:click={uploadProgress}
                    disabled={uploading}
                    class="glass-button px-8 py-3 rounded-lg hover:shadow-lg hover:shadow-primary/20 transition transform hover:-translate-y-0.5 disabled:opacity-50 disabled:cursor-not-allowed disabled:transform-none font-bold text-black"
                >
                    {uploading ? "Saving..." : (editingId ? "Update Entry" : "Save Progress")}
                </button>
            </div>
        </div>
    </div>

    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
        {#each progressEntries as entry}
            <div class="glass-card rounded-2xl overflow-hidden flex flex-col hover:shadow-2xl hover:shadow-primary/10 transition duration-300 group border border-white/5 relative">
                <div class="aspect-[4/3] bg-black/40 relative overflow-hidden">
                    <img
                        src={entry.imageData}
                        alt={entry.title}
                        class="w-full h-full object-cover transition duration-700 group-hover:scale-110"
                    />
                    <div class="absolute inset-0 bg-gradient-to-t from-black/80 via-transparent to-transparent opacity-0 group-hover:opacity-100 transition duration-300 flex items-end p-6">
                        <span class="text-white font-medium bg-primary/20 backdrop-blur-md px-3 py-1 rounded-full border border-primary/30 text-sm">
                            {new Date(entry.createdAt).toLocaleDateString(undefined, { year: 'numeric', month: 'long', day: 'numeric' })}
                        </span>
                    </div>
                </div>
                <div class="p-6 flex-grow flex flex-col gap-2">
                    <div class="flex justify-between items-start">
                        <h3 class="text-2xl font-bold text-white group-hover:text-primary transition">{entry.title}</h3>
                    </div>
                    <p class="text-gray-400 text-sm leading-relaxed">{entry.description}</p>
                    
                    <div class="mt-4 flex justify-end gap-3 pt-4 border-t border-white/5">
                        <button 
                            on:click={() => editProgress(entry)}
                            class="text-sm font-medium text-primary hover:text-green-300 transition"
                        >
                            Edit
                        </button>
                        <button 
                            on:click={() => deleteProgress(entry.id)}
                            class="text-sm font-medium text-accent hover:text-red-300 transition"
                        >
                            Delete
                        </button>
                    </div>
                </div>
            </div>
        {/each}
    </div>
</div>
