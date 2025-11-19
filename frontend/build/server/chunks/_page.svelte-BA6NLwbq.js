import { attr, attr_class, ensure_array_like, escape_html } from "./index2-BLSeo-Vt.js";
import "./clsx-cC83_lR5.js";
import "./cookie-toVUY8nI.js";

//#region .svelte-kit/adapter-bun/entries/pages/app/calendar/_page.svelte.js
function _page($$renderer, $$props) {
	$$renderer.component(($$renderer2) => {
		let year, month, daysInMonth, firstDayOfMonth, monthName;
		let events = [];
		let currentDate = /* @__PURE__ */ new Date();
		function getEventsForDay(day) {
			return events.filter((e) => {
				return e.date.getDate() === day && e.date.getMonth() === month && e.date.getFullYear() === year;
			});
		}
		year = currentDate.getFullYear();
		month = currentDate.getMonth();
		daysInMonth = new Date(year, month + 1, 0).getDate();
		firstDayOfMonth = new Date(year, month, 1).getDay();
		monthName = currentDate.toLocaleString("default", { month: "long" });
		$$renderer2.push(`<div class="space-y-8"><div class="flex items-center justify-between"><h1 class="text-4xl font-bold text-transparent bg-clip-text bg-gradient-to-r from-primary to-secondary">Calendar</h1> <div class="flex items-center gap-4 glass p-2 rounded-lg"><button class="p-2 hover:bg-white/10 rounded-lg text-white transition">&lt;</button> <h2 class="text-xl font-semibold text-white w-40 text-center">${escape_html(monthName)} ${escape_html(year)}</h2> <button class="p-2 hover:bg-white/10 rounded-lg text-white transition">></button></div></div> <div class="grid grid-cols-7 gap-4 text-center text-gray-400 font-semibold mb-2"><div>Sun</div><div>Mon</div><div>Tue</div><div>Wed</div><div>Thu</div><div>Fri</div><div>Sat</div></div> <div class="grid grid-cols-7 gap-4"><!--[-->`);
		const each_array = ensure_array_like(Array(firstDayOfMonth));
		for (let $$index = 0, $$length = each_array.length; $$index < $$length; $$index++) {
			each_array[$$index];
			$$renderer2.push(`<div class="h-32 rounded-xl"></div>`);
		}
		$$renderer2.push(`<!--]--> <!--[-->`);
		const each_array_1 = ensure_array_like(Array(daysInMonth));
		for (let i = 0, $$length = each_array_1.length; i < $$length; i++) {
			each_array_1[i];
			const day = i + 1;
			const dayEvents = getEventsForDay(day);
			const isToday = day === (/* @__PURE__ */ new Date()).getDate() && month === (/* @__PURE__ */ new Date()).getMonth() && year === (/* @__PURE__ */ new Date()).getFullYear();
			$$renderer2.push(`<div${attr_class("h-32 glass-card p-2 cursor-pointer transition relative overflow-hidden hover:border-primary/50 group flex flex-col", void 0, {
				"border-primary": isToday,
				"bg-primary-10": isToday
			})}><span${attr_class("font-bold text-gray-300 group-hover:text-white transition", void 0, { "text-primary": isToday })}>${escape_html(day)}</span> <div class="mt-1 space-y-1 overflow-y-auto scrollbar-hide w-full flex-grow"><!--[-->`);
			const each_array_2 = ensure_array_like(dayEvents);
			for (let $$index_1 = 0, $$length2 = each_array_2.length; $$index_1 < $$length2; $$index_1++) {
				let event = each_array_2[$$index_1];
				$$renderer2.push(`<div class="text-xs bg-primary/80 text-black font-semibold px-2 py-1 rounded shadow-sm truncate w-full block"${attr("title", event.title)}>${escape_html(event.title)}</div>`);
			}
			$$renderer2.push(`<!--]--></div></div>`);
		}
		$$renderer2.push(`<!--]--></div></div> `);
		$$renderer2.push("<!--[!-->");
		$$renderer2.push(`<!--]-->`);
	});
}

//#endregion
export { _page as default };
//# sourceMappingURL=_page.svelte-BA6NLwbq.js.map