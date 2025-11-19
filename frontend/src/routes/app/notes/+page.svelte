<script>
    import { onMount } from "svelte";
    import { cookies } from "$lib/api/cookie";
    import { request } from "$lib/api/util";

    let notes = [];
    let title = "";
    let content = "";
    let editingId = null;

    async function fetchNotes() {
        try {
            notes = await request("/notes", "GET", null, { Authorization: `Bearer ${cookies.token}` });
        } catch (e) {
            console.error(e);
        }
    }

    async function saveNote() {
        const payload = JSON.stringify({ title, content });
        const headers = {
            Authorization: `Bearer ${cookies.token}`,
            "Content-Type": "application/json",
        };

        try {
            if (editingId) {
                await request(`/notes/${editingId}`, "PUT", payload, headers);
            } else {
                await request("/notes", "POST", payload, headers);
            }
            title = "";
            content = "";
            editingId = null;
            fetchNotes();
        } catch (e) {
            console.error(e);
        }
    }

    async function deleteNote(id) {
        if (!confirm("Are you sure?")) return;
        try {
            await request(`/notes/${id}`, "DELETE", null, { Authorization: `Bearer ${cookies.token}` });
            fetchNotes();
        } catch (e) {
            console.error(e);
        }
    }

    function editNote(note) {
        title = note.title;
        content = note.content;
        editingId = note.id;
    }

    function cancelEdit() {
        title = "";
        content = "";
        editingId = null;
    }

    onMount(fetchNotes);
</script>

<div class="space-y-8">
    <div class="flex items-center justify-between">
        <h1 class="text-4xl font-bold text-transparent bg-clip-text bg-gradient-to-r from-primary to-secondary">Notes</h1>
    </div>

    <div class="glass-card p-6">
        <h2 class="text-xl font-semibold mb-4 text-primary">{editingId ? "Edit Note" : "New Note"}</h2>
        <div class="space-y-4">
            <input
                bind:value={title}
                placeholder="Title"
                class="w-full glass-input px-4 py-3 rounded-lg focus:ring-2 focus:ring-primary transition"
            />
            <textarea
                bind:value={content}
                placeholder="Content"
                rows="4"
                class="w-full glass-input px-4 py-3 rounded-lg focus:ring-2 focus:ring-primary transition resize-none"
            ></textarea>
            <div class="flex gap-3">
                <button
                    on:click={saveNote}
                    class="glass-button px-6 py-2 rounded-lg hover:shadow-lg hover:shadow-primary/20 transition transform hover:-translate-y-0.5"
                >
                    {editingId ? "Update" : "Add"} Note
                </button>
                {#if editingId}
                    <button
                        on:click={cancelEdit}
                        class="px-6 py-2 rounded-lg border border-white/10 hover:bg-white/10 transition text-gray-300"
                    >
                        Cancel
                    </button>
                {/if}
            </div>
        </div>
    </div>

    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        {#each notes as note}
            <div class="glass-card p-6 flex flex-col relative group hover:border-primary/50 transition duration-300">
                <h3 class="text-xl font-bold text-white mb-2">{note.title}</h3>
                <p class="text-gray-300 whitespace-pre-wrap flex-grow leading-relaxed">{note.content}</p>
                <div class="mt-6 flex justify-end gap-3 pt-4 border-t border-white/5">
                    <button
                        on:click={() => editNote(note)}
                        class="text-primary hover:text-green-300 font-medium text-sm"
                    >
                        Edit
                    </button>
                    <button
                        on:click={() => deleteNote(note.id)}
                        class="text-accent hover:text-red-300 font-medium text-sm"
                    >
                        Delete
                    </button>
                </div>
            </div>
        {/each}
    </div>
</div>
