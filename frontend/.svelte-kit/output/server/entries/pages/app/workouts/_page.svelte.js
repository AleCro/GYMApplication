import { a4 as ensure_array_like, a7 as attr_class, a0 as escape_html, a3 as stringify } from "../../../../chunks/index2.js";
function _page($$renderer) {
  let selectedGroup = "Chest";
  const workouts = {
    Chest: [
      { name: "Bench Press", sets: "4", reps: "8-12" },
      { name: "Incline Dumbbell Press", sets: "3", reps: "10-12" },
      { name: "Cable Flyes", sets: "3", reps: "12-15" },
      { name: "Pushups", sets: "3", reps: "Failure" }
    ],
    Back: [
      { name: "Pullups", sets: "4", reps: "8-12" },
      { name: "Barbell Rows", sets: "4", reps: "8-10" },
      { name: "Lat Pulldowns", sets: "3", reps: "10-12" },
      { name: "Face Pulls", sets: "3", reps: "15-20" }
    ],
    Legs: [
      { name: "Squats", sets: "4", reps: "6-10" },
      { name: "Romanian Deadlifts", sets: "3", reps: "8-12" },
      { name: "Leg Press", sets: "3", reps: "10-15" },
      { name: "Calf Raises", sets: "4", reps: "15-20" }
    ],
    Shoulders: [
      { name: "Overhead Press", sets: "4", reps: "8-10" },
      { name: "Lateral Raises", sets: "4", reps: "12-15" },
      { name: "Front Raises", sets: "3", reps: "12-15" },
      { name: "Rear Delt Flyes", sets: "3", reps: "15-20" }
    ],
    Arms: [
      { name: "Barbell Curls", sets: "3", reps: "10-12" },
      { name: "Tricep Pushdowns", sets: "3", reps: "12-15" },
      { name: "Hammer Curls", sets: "3", reps: "10-12" },
      { name: "Skullcrushers", sets: "3", reps: "10-12" }
    ],
    Abs: [
      { name: "Crunches", sets: "3", reps: "15-20" },
      { name: "Leg Raises", sets: "3", reps: "12-15" },
      { name: "Plank", sets: "3", reps: "60s" }
    ]
  };
  const groups = Object.keys(workouts);
  $$renderer.push(`<div class="space-y-8"><h1 class="text-4xl font-bold text-transparent bg-clip-text bg-gradient-to-r from-primary to-secondary">Workouts</h1> <div class="flex flex-wrap gap-3"><!--[-->`);
  const each_array = ensure_array_like(groups);
  for (let $$index = 0, $$length = each_array.length; $$index < $$length; $$index++) {
    let group = each_array[$$index];
    $$renderer.push(`<button${attr_class(`px-6 py-3 rounded-lg font-semibold transition shadow-md ${stringify(selectedGroup === group ? "bg-primary text-black shadow-lg shadow-primary/20 transform -translate-y-0.5" : "glass text-gray-300 hover:bg-white/10 hover:text-white")}`)}>${escape_html(group)}</button>`);
  }
  $$renderer.push(`<!--]--></div> <div class="glass-card p-8 rounded-2xl"><h2 class="text-3xl font-bold text-white mb-8 border-b border-white/10 pb-4 flex items-center gap-3"><span class="text-primary">#</span> ${escape_html(selectedGroup)} Routine</h2> <div class="space-y-4"><!--[-->`);
  const each_array_1 = ensure_array_like(workouts[selectedGroup]);
  for (let $$index_1 = 0, $$length = each_array_1.length; $$index_1 < $$length; $$index_1++) {
    let exercise = each_array_1[$$index_1];
    $$renderer.push(`<div class="flex items-center justify-between bg-white/5 p-5 rounded-xl hover:bg-white/10 transition border border-white/5 hover:border-primary/30 group"><div><div class="font-bold text-xl text-white group-hover:text-primary transition">${escape_html(exercise.name)}</div></div> <div class="text-right flex gap-6"><div class="flex flex-col items-end"><span class="text-xs text-gray-400 uppercase tracking-wider">Sets</span> <span class="text-secondary font-mono font-bold text-lg">${escape_html(exercise.sets)}</span></div> <div class="flex flex-col items-end w-16"><span class="text-xs text-gray-400 uppercase tracking-wider">Reps</span> <span class="text-gray-200 font-mono text-lg">${escape_html(exercise.reps)}</span></div></div></div>`);
  }
  $$renderer.push(`<!--]--></div></div></div>`);
}
export {
  _page as default
};
