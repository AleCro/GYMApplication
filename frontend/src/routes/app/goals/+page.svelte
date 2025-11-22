<script>
    import { onMount } from "svelte";
    import { cookies } from "$lib/api/cookie";
    import { request } from "$lib/api/util";

    let goals = [];
    let title = "";
    let description = "";
    let subGoals = []; // Array of { id: string, title: string, completed: boolean }
    let newSubGoalTitle = "";
    let editingId = null;

    async function fetchGoals() {
        try {
            goals = await request("/goals", "GET", null, { Authorization: `Bearer ${cookies.token}` });
        } catch (e) {
            console.error(e);
        }
    }

    function addSubGoal() {
        if (!newSubGoalTitle.trim()) return;
        subGoals = [...subGoals, { id: crypto.randomUUID(), title: newSubGoalTitle, completed: false }];
        newSubGoalTitle = "";
    }

    function removeSubGoal(id) {
        subGoals = subGoals.filter(sg => sg.id !== id);
    }

    function toggleSubGoal(id) {
        subGoals = subGoals.map(sg => sg.id === id ? { ...sg, completed: !sg.completed } : sg);
    }

    async function saveGoal() {
        const payload = JSON.stringify({ title, description, subGoals });
        const headers = {
            Authorization: `Bearer ${cookies.token}`,
            "Content-Type": "application/json",
        };

        try {
            if (editingId) {
                await request(`/goals/${editingId}`, "PUT", payload, headers);
            } else {
                await request("/goals", "POST", payload, headers);
            }
            resetForm();
            fetchGoals();
        } catch (e) {
            console.error(e);
        }
    }

    async function deleteGoal(id) {
        if (!confirm("Are you sure you want to delete this goal?")) return;
        try {
            await request(`/goals/${id}`, "DELETE", null, { Authorization: `Bearer ${cookies.token}` });
            fetchGoals();
        } catch (e) {
            console.error(e);
        }
    }

    function editGoal(goal) {
        title = goal.title;
        description = goal.description;
        subGoals = goal.subGoals || [];
        editingId = goal.id;
    }

    async function toggleGoalSubtask(goal, subGoalId) {
        // Optimistic update
        const updatedSubGoals = goal.subGoals.map(sg =>
            sg.id === subGoalId ? { ...sg, completed: !sg.completed } : sg
        );
        const updatedGoal = { ...goal, subGoals: updatedSubGoals };

        // Update local state immediately
        goals = goals.map(g => g.id === goal.id ? updatedGoal : g);

        const payload = JSON.stringify({
            title: goal.title,
            description: goal.description,
            subGoals: updatedSubGoals
        });
        const headers = {
            Authorization: `Bearer ${cookies.token}`,
            "Content-Type": "application/json",
        };

        try {
            await request(`/goals/${goal.id}`, "PUT", payload, headers);
        } catch (e) {
            console.error(e);
            fetchGoals(); // Revert on error
        }
    }

    function resetForm() {
        title = "";
        description = "";
        subGoals = [];
        newSubGoalTitle = "";
        editingId = null;
    }

    onMount(fetchGoals);
</script>

<div class="space-y-8">
    <div class="flex items-center justify-between">
        <h1 class="text-4xl font-bold text-transparent bg-clip-text bg-gradient-to-r from-primary to-secondary">Goals</h1>
    </div>

    <div class="glass-card p-6">
        <h2 class="text-xl font-semibold mb-4 text-primary">{editingId ? "Edit Goal" : "New Goal"}</h2>
        <div class="space-y-4">
            <input
                bind:value={title}
                placeholder="Goal Title"
                class="w-full glass-input px-4 py-3 rounded-lg focus:ring-2 focus:ring-primary transition"
            />
            <textarea
                bind:value={description}
                placeholder="Description"
                rows="3"
                class="w-full glass-input px-4 py-3 rounded-lg focus:ring-2 focus:ring-primary transition resize-none"
            ></textarea>
            
            <div class="space-y-2">
                <h3 class="text-sm font-medium text-gray-300">Steps</h3>
                <div class="flex gap-2">
                    <input
                        bind:value={newSubGoalTitle}
                        placeholder="Add a subtask..."
                        class="flex-grow glass-input px-4 py-2 rounded-lg focus:ring-2 focus:ring-primary transition"
                        on:keydown={(e) => e.key === 'Enter' && addSubGoal()}
                    />
                    <button
                        on:click={addSubGoal}
                        class="glass-button px-4 py-2 rounded-lg hover:bg-white/10 transition"
                    >
                        Add
                    </button>
                </div>
                {#if subGoals.length > 0}
                    <ul class="space-y-2 mt-2">
                        {#each subGoals as subGoal (subGoal.id)}
                            <li class="flex items-center justify-between bg-white/5 p-2 rounded-lg">
                                <span class={subGoal.completed ? "line-through text-gray-500" : "text-gray-200"}>{subGoal.title}</span>
                                <button
                                    on:click={() => removeSubGoal(subGoal.id)}
                                    class="text-red-400 hover:text-red-300 text-sm"
                                >
                                    Remove
                                </button>
                            </li>
                        {/each}
                    </ul>
                {/if}
            </div>

            <div class="flex gap-3 pt-4">
                <button
                    on:click={saveGoal}
                    class="glass-button px-6 py-2 rounded-lg hover:shadow-lg hover:shadow-primary/20 transition transform hover:-translate-y-0.5"
                >
                    {editingId ? "Update" : "Add"} Goal
                </button>
                {#if editingId}
                    <button
                        on:click={resetForm}
                        class="px-6 py-2 rounded-lg border border-white/10 hover:bg-white/10 transition text-gray-300"
                    >
                        Cancel
                    </button>
                {/if}
            </div>
        </div>
    </div>

    <div class="grid grid-cols-1 gap-6">
        {#each goals as goal}
            <div class="glass-card p-6 flex flex-col relative group hover:border-primary/50 transition duration-300">
                <div class="flex justify-between items-start mb-4">
                    <div>
                        <h3 class="text-2xl font-bold text-white">{goal.title}</h3>
                        <p class="text-gray-400 mt-1">{goal.description}</p>
                    </div>
                    <div class="flex gap-2">
                        <button
                            on:click={() => editGoal(goal)}
                            class="text-primary hover:text-green-300 font-medium text-sm"
                        >
                            Edit
                        </button>
                        <button
                            on:click={() => deleteGoal(goal.id)}
                            class="text-accent hover:text-red-300 font-medium text-sm"
                        >
                            Delete
                        </button>
                    </div>
                </div>

                {#if goal.subGoals && goal.subGoals.length > 0}
                    <div class="mt-4 space-y-2">
                        <h4 class="text-sm font-semibold text-gray-300 uppercase tracking-wider">Steps</h4>
                        <div class="space-y-1">
                            {#each goal.subGoals as subGoal}
                                <button
                                    class="flex items-center gap-2 text-gray-300 hover:text-white transition group/item w-full text-left"
                                    on:click={() => toggleGoalSubtask(goal, subGoal.id)}
                                >
                                    <div class={`w-4 h-4 rounded border flex items-center justify-center transition ${subGoal.completed ? 'bg-green-500 border-green-500' : 'border-gray-500 group-hover/item:border-primary'}`}>
                                        {#if subGoal.completed}
                                            <svg class="w-3 h-3 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="3" d="M5 13l4 4L19 7"></path></svg>
                                        {/if}
                                    </div>
                                    <span class={subGoal.completed ? "line-through text-gray-500" : ""}>{subGoal.title}</span>
                                </button>
                            {/each}
                        </div>
                        
                        <!-- Progress Bar -->
                        <div class="mt-4">
                            <div class="flex justify-between text-xs text-gray-400 mb-1">
                                <span>Progress</span>
                                <span>{Math.round((goal.subGoals.filter(g => g.completed).length / goal.subGoals.length) * 100)}%</span>
                            </div>
                            <div class="w-full bg-gray-700 rounded-full h-2.5">
                                <div class="bg-primary h-2.5 rounded-full" style="width: {Math.round((goal.subGoals.filter(g => g.completed).length / goal.subGoals.length) * 100)}%"></div>
                            </div>
                        </div>
                    </div>
                {/if}
            </div>
        {/each}
    </div>
</div>
