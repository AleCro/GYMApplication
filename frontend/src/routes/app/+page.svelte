<script>
    import { onMount } from "svelte";
    import { cookies } from "$lib/api/cookie";
    import { request } from "$lib/api/util";

    let stats = [
        { title: 'Total Notes', value: '0', trend: '', color: 'indigo' },
        { title: 'Upcoming Events', value: '0', trend: '', color: 'green' },
        { title: 'Progress Entries', value: '0', trend: '', color: 'primary' },
        { title: 'Workouts Logged', value: 'Coming Soon', trend: '', color: 'yellow' },
    ];

    let recentActivity = [];

    async function fetchData() {
        try {
            const [notes, events, progress] = await Promise.all([
                request("/notes", "GET", null, { Authorization: `Bearer ${cookies.token}` }),
                request("/events", "GET", null, { Authorization: `Bearer ${cookies.token}` }),
                request("/progress", "GET", null, { Authorization: `Bearer ${cookies.token}` })
            ]);

            // Update Stats
            stats[0].value = notes.length.toString();
            stats[1].value = events.filter(e => new Date(e.date) >= new Date()).length.toString();
            stats[2].value = progress.length.toString();

            // Combine for recent activity
            const noteActivity = notes.map(n => ({ type: 'Note', title: n.title, date: new Date(n.updatedAt), color: 'indigo' }));
            const eventActivity = events.map(e => ({ type: 'Event', title: e.title, date: new Date(e.createdAt), color: 'green' }));
            const progressActivity = progress.map(p => ({ type: 'Progress', title: p.title, date: new Date(p.createdAt), color: 'primary' }));

            recentActivity = [...noteActivity, ...eventActivity, ...progressActivity]
                .sort((a, b) => b.date - a.date)
                .slice(0, 5);

        } catch (e) {
            console.error(e);
        }
    }

    onMount(fetchData);
</script>

<svelte:head>
    <title>Dashboard - AleGYM</title>
</svelte:head>

<div class="space-y-10">

    <h1 class="text-4xl font-bold text-transparent bg-clip-text bg-gradient-to-r from-primary to-secondary">Dashboard</h1>
    <p class="text-gray-400">Welcome back! Here's what's happening with your fitness journey.</p>

    <section>
        <h2 class="text-2xl font-semibold text-white mb-6 flex items-center gap-2">
            <span class="text-primary">📊</span> Overview
        </h2>
        <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-6">
            {#each stats as stat}
                <div class="p-6 rounded-2xl glass-card border border-white/5 shadow-lg hover:shadow-primary/10 transition">
                    <h3 class="text-sm font-bold text-gray-400 uppercase tracking-wider mb-2">{stat.title}</h3>
                    <div class="text-3xl font-bold text-white">{stat.value}</div>
                </div>
            {/each}
        </div>
    </section>

    <div class="grid grid-cols-1 lg:grid-cols-3 gap-8">
        
        <div class="lg:col-span-2 p-8 rounded-2xl glass-card border border-white/5 shadow-xl space-y-6">
            <h2 class="text-xl font-bold text-white border-b border-white/10 pb-4">Recent Activity</h2>
            {#if recentActivity.length > 0}
                <ul class="space-y-4">
                    {#each recentActivity as activity}
                        <li class="flex items-center gap-4 p-4 rounded-xl bg-white/5 hover:bg-white/10 transition border border-white/5">
                            <div class="w-3 h-3 rounded-full shrink-0" 
                                class:bg-indigo-500={activity.color === 'indigo'}
                                class:bg-green-500={activity.color === 'green'}
                                class:bg-primary={activity.color === 'primary'}
                            ></div>
                            <div class="flex-grow">
                                <p class="text-sm text-white font-medium">
                                    {activity.type}: <span class="text-gray-300">{activity.title}</span>
                                </p>
                                <span class="text-xs text-gray-500">{activity.date.toLocaleString()}</span>
                            </div>
                        </li>
                    {/each}
                </ul>
            {:else}
                <p class="text-gray-500 italic">No recent activity found.</p>
            {/if}
        </div>

        <div class="p-8 rounded-2xl glass-card border border-white/5 shadow-xl space-y-6">
            <h2 class="text-xl font-bold text-white border-b border-white/10 pb-4">Quick Actions</h2>
            <div class="space-y-3">
                <a href="/app/workouts" class="flex items-center gap-3 p-4 rounded-xl bg-white/5 hover:bg-primary/20 hover:border-primary/30 border border-white/5 transition group">
                    <span class="text-2xl group-hover:scale-110 transition">💪</span>
                    <div class="text-sm font-bold text-white">Start Workout</div>
                </a>
                <a href="/app/progress" class="flex items-center gap-3 p-4 rounded-xl bg-white/5 hover:bg-primary/20 hover:border-primary/30 border border-white/5 transition group">
                    <span class="text-2xl group-hover:scale-110 transition">📸</span>
                    <div class="text-sm font-bold text-white">Log Progress</div>
                </a>
                <a href="/app/calendar" class="flex items-center gap-3 p-4 rounded-xl bg-white/5 hover:bg-primary/20 hover:border-primary/30 border border-white/5 transition group">
                    <span class="text-2xl group-hover:scale-110 transition">📅</span>
                    <div class="text-sm font-bold text-white">Schedule Event</div>
                </a>
            </div>
        </div>
    </div>
</div>