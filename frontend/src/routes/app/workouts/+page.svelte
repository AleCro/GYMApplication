<script>
    let selectedGroup = "Chest";

    const workouts = {
        Chest: [
            { name: "Bench Press", sets: "4", reps: "8-12" },
            { name: "Incline Dumbbell Press", sets: "3", reps: "10-12" },
            { name: "Cable Flyes", sets: "3", reps: "12-15" },
            { name: "Pushups", sets: "3", reps: "Failure" },
        ],
        Back: [
            { name: "Pullups", sets: "4", reps: "8-12" },
            { name: "Barbell Rows", sets: "4", reps: "8-10" },
            { name: "Lat Pulldowns", sets: "3", reps: "10-12" },
            { name: "Face Pulls", sets: "3", reps: "15-20" },
        ],
        Legs: [
            { name: "Squats", sets: "4", reps: "6-10" },
            { name: "Romanian Deadlifts", sets: "3", reps: "8-12" },
            { name: "Leg Press", sets: "3", reps: "10-15" },
            { name: "Calf Raises", sets: "4", reps: "15-20" },
        ],
        Shoulders: [
            { name: "Overhead Press", sets: "4", reps: "8-10" },
            { name: "Lateral Raises", sets: "4", reps: "12-15" },
            { name: "Front Raises", sets: "3", reps: "12-15" },
            { name: "Rear Delt Flyes", sets: "3", reps: "15-20" },
        ],
        Arms: [
            { name: "Barbell Curls", sets: "3", reps: "10-12" },
            { name: "Tricep Pushdowns", sets: "3", reps: "12-15" },
            { name: "Hammer Curls", sets: "3", reps: "10-12" },
            { name: "Skullcrushers", sets: "3", reps: "10-12" },
        ],
        Abs: [
            { name: "Crunches", sets: "3", reps: "15-20" },
            { name: "Leg Raises", sets: "3", reps: "12-15" },
            { name: "Plank", sets: "3", reps: "60s" },
        ]
    };

    const groups = Object.keys(workouts);
</script>


<div class="space-y-8">
    <h1 class="text-4xl font-bold text-transparent bg-clip-text bg-gradient-to-r from-primary to-secondary">Workouts</h1>

    <div class="flex flex-wrap gap-3">
        {#each groups as group}
            <button
                class="px-6 py-3 rounded-lg font-semibold transition shadow-md {selectedGroup === group ? 'bg-primary text-black shadow-lg shadow-primary/20 transform -translate-y-0.5' : 'glass text-gray-300 hover:bg-white/10 hover:text-white'}"
                on:click={() => selectedGroup = group}
            >
                {group}
            </button>
        {/each}
    </div>

    <div class="glass-card p-8 rounded-2xl">
        <h2 class="text-3xl font-bold text-white mb-8 border-b border-white/10 pb-4 flex items-center gap-3">
            <span class="text-primary">#</span> {selectedGroup} Routine
        </h2>
        <div class="space-y-4">
            {#each workouts[selectedGroup] as exercise}
                <div class="flex items-center justify-between bg-white/5 p-5 rounded-xl hover:bg-white/10 transition border border-white/5 hover:border-primary/30 group">
                    <div>
                        <div class="font-bold text-xl text-white group-hover:text-primary transition">{exercise.name}</div>
                    </div>
                    <div class="text-right flex gap-6">
                        <div class="flex flex-col items-end">
                            <span class="text-xs text-gray-400 uppercase tracking-wider">Sets</span>
                            <span class="text-secondary font-mono font-bold text-lg">{exercise.sets}</span>
                        </div>
                        <div class="flex flex-col items-end w-16">
                            <span class="text-xs text-gray-400 uppercase tracking-wider">Reps</span>
                            <span class="text-gray-200 font-mono text-lg">{exercise.reps}</span>
                        </div>
                    </div>
                </div>
            {/each}
        </div>
    </div>
</div>
