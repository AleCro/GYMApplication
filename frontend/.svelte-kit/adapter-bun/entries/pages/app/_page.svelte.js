import { $ as head, a4 as ensure_array_like, a0 as escape_html, a7 as attr_class } from "../../../chunks/index2.js";
import "../../../chunks/cookie.js";
function _page($$renderer, $$props) {
  $$renderer.component(($$renderer2) => {
    let stats = [
      { title: "Total Notes", value: "0", trend: "", color: "indigo" },
      {
        title: "Upcoming Events",
        value: "0",
        trend: "",
        color: "green"
      },
      {
        title: "Progress Entries",
        value: "0",
        trend: "",
        color: "primary"
      },
      {
        title: "Workouts Logged",
        value: "Coming Soon",
        trend: "",
        color: "yellow"
      }
    ];
    let recentActivity = [];
    head($$renderer2, ($$renderer3) => {
      $$renderer3.title(($$renderer4) => {
        $$renderer4.push(`<title>Dashboard - AleGYM</title>`);
      });
    });
    $$renderer2.push(`<div class="space-y-10"><h1 class="text-4xl font-bold text-transparent bg-clip-text bg-gradient-to-r from-primary to-secondary">Dashboard</h1> <p class="text-gray-400">Welcome back! Here's what's happening with your fitness journey.</p> <section><h2 class="text-2xl font-semibold text-white mb-6 flex items-center gap-2"><span class="text-primary">📊</span> Overview</h2> <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-6"><!--[-->`);
    const each_array = ensure_array_like(stats);
    for (let $$index = 0, $$length = each_array.length; $$index < $$length; $$index++) {
      let stat = each_array[$$index];
      $$renderer2.push(`<div class="p-6 rounded-2xl glass-card border border-white/5 shadow-lg hover:shadow-primary/10 transition"><h3 class="text-sm font-bold text-gray-400 uppercase tracking-wider mb-2">${escape_html(stat.title)}</h3> <div class="text-3xl font-bold text-white">${escape_html(stat.value)}</div></div>`);
    }
    $$renderer2.push(`<!--]--></div></section> <div class="grid grid-cols-1 lg:grid-cols-3 gap-8"><div class="lg:col-span-2 p-8 rounded-2xl glass-card border border-white/5 shadow-xl space-y-6"><h2 class="text-xl font-bold text-white border-b border-white/10 pb-4">Recent Activity</h2> `);
    if (recentActivity.length > 0) {
      $$renderer2.push("<!--[-->");
      $$renderer2.push(`<ul class="space-y-4"><!--[-->`);
      const each_array_1 = ensure_array_like(recentActivity);
      for (let $$index_1 = 0, $$length = each_array_1.length; $$index_1 < $$length; $$index_1++) {
        let activity = each_array_1[$$index_1];
        $$renderer2.push(`<li class="flex items-center gap-4 p-4 rounded-xl bg-white/5 hover:bg-white/10 transition border border-white/5"><div${attr_class("w-3 h-3 rounded-full shrink-0", void 0, {
          "bg-indigo-500": activity.color === "indigo",
          "bg-green-500": activity.color === "green",
          "bg-primary": activity.color === "primary"
        })}></div> <div class="flex-grow"><p class="text-sm text-white font-medium">${escape_html(activity.type)}: <span class="text-gray-300">${escape_html(activity.title)}</span></p> <span class="text-xs text-gray-500">${escape_html(activity.date.toLocaleString())}</span></div></li>`);
      }
      $$renderer2.push(`<!--]--></ul>`);
    } else {
      $$renderer2.push("<!--[!-->");
      $$renderer2.push(`<p class="text-gray-500 italic">No recent activity found.</p>`);
    }
    $$renderer2.push(`<!--]--></div> <div class="p-8 rounded-2xl glass-card border border-white/5 shadow-xl space-y-6"><h2 class="text-xl font-bold text-white border-b border-white/10 pb-4">Quick Actions</h2> <div class="space-y-3"><a href="/app/workouts" class="flex items-center gap-3 p-4 rounded-xl bg-white/5 hover:bg-primary/20 hover:border-primary/30 border border-white/5 transition group"><span class="text-2xl group-hover:scale-110 transition">💪</span> <div class="text-sm font-bold text-white">Start Workout</div></a> <a href="/app/progress" class="flex items-center gap-3 p-4 rounded-xl bg-white/5 hover:bg-primary/20 hover:border-primary/30 border border-white/5 transition group"><span class="text-2xl group-hover:scale-110 transition">📸</span> <div class="text-sm font-bold text-white">Log Progress</div></a> <a href="/app/calendar" class="flex items-center gap-3 p-4 rounded-xl bg-white/5 hover:bg-primary/20 hover:border-primary/30 border border-white/5 transition group"><span class="text-2xl group-hover:scale-110 transition">📅</span> <div class="text-sm font-bold text-white">Schedule Event</div></a></div></div></div></div>`);
  });
}
export {
  _page as default
};
