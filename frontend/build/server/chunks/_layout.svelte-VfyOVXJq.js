import { attr, bind_props, ensure_array_like, escape_html, head, slot, store_get, stringify, unsubscribe_stores } from "./index2-BLSeo-Vt.js";
import "./server-CzoYAPU4.js";
import "./clsx-cC83_lR5.js";
import { writable } from "./chunks-DNg0X1yh.js";
import { onDestroy, userStore } from "./stores-C-8AFukU.js";
import "./cookie-toVUY8nI.js";

//#region .svelte-kit/adapter-bun/entries/pages/app/_layout.svelte.js
function _layout($$renderer, $$props) {
	$$renderer.component(($$renderer2) => {
		var $$store_subs;
		let user, userInitial;
		let data = $$props["data"];
		userStore.set(data?.user ?? null);
		let unsubscribe = userStore.subscribe((value) => {
			user = value;
		});
		let sessionRefreshInterval = null;
		let mobileOpen = writable(false);
		onDestroy(() => {
			clearInterval(sessionRefreshInterval);
			sessionRefreshInterval = null;
			unsubscribe();
		});
		let mainNav = [
			{
				href: "/app",
				label: "Dashboard",
				icon: "🏠"
			},
			{
				href: "/app/notes",
				label: "Notes",
				icon: "📝"
			},
			{
				href: "/app/calendar",
				label: "Calendar",
				icon: "📅"
			},
			{
				href: "/app/workouts",
				label: "Workouts",
				icon: "💪"
			},
			{
				href: "/app/progress",
				label: "Progress",
				icon: "📈"
			},
			{
				href: "/app/user-management",
				label: "User Management",
				icon: "👤",
				group: 255
			}
		];
		let secondaryNav = [{
			href: "/app/documentation",
			label: "Documentation",
			icon: "📚"
		}, {
			href: "/logout",
			label: "Sign out",
			icon: "🚪"
		}];
		user = data.user;
		userInitial = user ? user.username ? user.username.charAt(0).toUpperCase() : "G" : "U";
		head($$renderer2, ($$renderer3) => {
			$$renderer3.title(($$renderer4) => {
				$$renderer4.push(`<title>${escape_html(data?.title ?? "AleGYM")}</title>`);
			});
			$$renderer3.push(`<meta charset="utf-8"/> <meta name="viewport" content="width=device-width,initial-scale=1"/> <meta name="description" content="AleGYM — Your Ultimate Fitness Companion"/> <link rel="icon" href="/favicon.ico"/>`);
		});
		$$renderer2.push(`<div class="min-h-screen bg-background text-text font-sans flex flex-col bg-[url('https://images.unsplash.com/photo-1534438327276-14e5300c3a48?q=80&amp;w=2070&amp;auto=format&amp;fit=crop')] bg-cover bg-center bg-fixed bg-no-repeat"><div class="absolute inset-0 bg-black/60 backdrop-blur-[2px] z-0"></div> <header class="glass sticky top-0 z-50 border-b-0"><div class="max-w-screen-2xl mx-auto px-4 sm:px-6 lg:px-8"><div class="flex items-center justify-between h-16"><div class="flex items-center gap-4"><button class="p-2 rounded-md hover:bg-white/10 transition focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-primary lg:hidden" aria-label="Toggle sidebar">`);
		if (store_get($$store_subs ??= {}, "$mobileOpen", mobileOpen)) {
			$$renderer2.push("<!--[-->");
			$$renderer2.push(`<svg class="w-6 h-6" viewBox="0 0 24 24" fill="none" stroke="currentColor"><path stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round" d="M6 18L18 6M6 6l12 12"></path></svg>`);
		} else {
			$$renderer2.push("<!--[!-->");
			$$renderer2.push(`<svg class="w-6 h-6" viewBox="0 0 24 24" fill="none" stroke="currentColor"><path stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round" d="M4 7h16M4 12h16M4 17h16"></path></svg>`);
		}
		$$renderer2.push(`<!--]--></button> <a href="/app" class="flex items-center gap-3 focus:outline-none focus-visible:ring-2 focus-visible:ring-primary rounded-lg -m-1 p-1"><div class="w-9 h-9 rounded-lg bg-gradient-to-br from-primary to-secondary flex items-center justify-center text-white font-bold shadow-md shrink-0">AG</div> <div class="hidden sm:block"><div class="font-bold text-lg tracking-wide">AleGYM</div></div></a></div> <div class="flex items-center gap-4"><div class="hidden md:flex items-center glass-input rounded-lg px-3 py-1.5 focus-within:ring-2 focus-within:ring-primary transition"><svg class="w-5 h-5 mr-2 text-gray-400" viewBox="0 0 24 24" fill="none" aria-hidden="" stroke="currentColor"><path stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round" d="M21 21l-4.35-4.35M11 18a7 7 0 100-14 7 7 0 000 14z"></path></svg> <input placeholder="Search..." class="bg-transparent outline-none text-sm placeholder-gray-400 w-40 text-text"/></div> <a href="/app/me" class="flex items-center gap-2 p-1.5 rounded-full hover:bg-white/10 transition focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-primary"><div class="w-8 h-8 rounded-full bg-secondary text-white flex items-center justify-center text-sm font-medium shadow-md shrink-0">${escape_html(userInitial)}</div> <span class="hidden lg:block text-sm font-medium text-text mr-1"${attr("title", `Session expires: ${stringify(user.exp - (/* @__PURE__ */ new Date()).valueOf() / 1e3)}m`)}>${escape_html(data?.user?.username ?? "Guest")}</span></a> <button class="hidden sm:block px-3 py-2 rounded-lg bg-accent/20 text-accent text-sm font-medium hover:bg-accent/30 transition shadow-sm focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-accent">Sign out</button></div></div></div></header> <div class="flex-grow flex max-w-screen-2xl mx-auto w-full min-w-0 z-10 relative"><aside class="hidden lg:block w-64 flex-shrink-0 glass border-y-0 border-l-0 border-r-0 pt-6 px-4 pb-4 sticky top-16 h-[calc(100vh-4rem)] overflow-hidden ml-4 my-4 rounded-xl"><nav class="space-y-6"><div class="space-y-1"><div class="text-xs uppercase tracking-wider text-gray-400 font-semibold mb-2">Workspace</div> <ul class="space-y-1 text-sm"><!--[-->`);
		const each_array = ensure_array_like(mainNav);
		for (let $$index = 0, $$length = each_array.length; $$index < $$length; $$index++) {
			let item = each_array[$$index];
			if (!item.group || item.group == data.user?.group) {
				$$renderer2.push("<!--[-->");
				$$renderer2.push(`<li><a${attr("href", item.href)} class="flex items-center gap-3 px-3 py-2 rounded-lg hover:bg-white/10 transition font-medium text-gray-300 hover:text-primary focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-primary"><span class="text-lg w-5 flex justify-center">${escape_html(item.icon)}</span> ${escape_html(item.label)}</a></li>`);
			} else $$renderer2.push("<!--[!-->");
			$$renderer2.push(`<!--]-->`);
		}
		$$renderer2.push(`<!--]--></ul></div> <div class="space-y-1 pt-4 border-t border-white/10"><div class="text-xs uppercase tracking-wider text-gray-400 font-semibold mb-2">General</div> <ul class="space-y-1 text-sm"><!--[-->`);
		const each_array_1 = ensure_array_like(secondaryNav);
		for (let $$index_1 = 0, $$length = each_array_1.length; $$index_1 < $$length; $$index_1++) {
			let item = each_array_1[$$index_1];
			if (!item.group || item.group == data.user?.group) {
				$$renderer2.push("<!--[-->");
				$$renderer2.push(`<li><a${attr("href", item.href)} class="flex items-center gap-3 px-3 py-2 rounded-lg hover:bg-white/10 transition font-medium text-gray-300 hover:text-primary focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-primary"><span class="text-lg w-5 flex justify-center">${escape_html(item.icon)}</span> ${escape_html(item.label)}</a></li>`);
			} else $$renderer2.push("<!--[!-->");
			$$renderer2.push(`<!--]-->`);
		}
		$$renderer2.push(`<!--]--></ul></div></nav></aside> `);
		if (store_get($$store_subs ??= {}, "$mobileOpen", mobileOpen)) {
			$$renderer2.push("<!--[-->");
			$$renderer2.push(`<div class="lg:hidden fixed inset-0 z-40 bg-black/70 transition-opacity duration-300 ease-in-out"></div> <aside id="mobile-sidebar" class="lg:hidden fixed left-0 top-0 bottom-0 z-50 w-72 p-6 glass shadow-2xl overflow-y-auto"><div class="flex items-center justify-between mb-8"><a href="/" class="flex items-center gap-3"><div class="w-9 h-9 rounded-lg bg-gradient-to-br from-primary to-secondary flex items-center justify-center text-white font-bold">AG</div> <div class="text-base font-bold">AleGYM</div></a> <button aria-label="Close drawer" class="p-2 rounded-md hover:bg-white/10 transition focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-primary"><svg class="w-6 h-6" viewBox="0 0 24 24" fill="none" stroke="currentColor"><path stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round" d="M6 18L18 6M6 6l12 12"></path></svg></button></div> <nav class="space-y-6"><div class="space-y-1"><div class="text-xs uppercase tracking-wider text-gray-400 font-semibold mb-2">Workspace</div> <ul class="space-y-1 text-base"><!--[-->`);
			const each_array_2 = ensure_array_like([...mainNav, ...secondaryNav]);
			for (let $$index_2 = 0, $$length = each_array_2.length; $$index_2 < $$length; $$index_2++) {
				let item = each_array_2[$$index_2];
				if (!item.group || item.group == data.user?.group) {
					$$renderer2.push("<!--[-->");
					$$renderer2.push(`<li><a${attr("href", item.href)} class="flex items-center gap-3 px-3 py-2 rounded-lg hover:bg-white/10 transition font-medium text-gray-300 hover:text-primary focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-primary"><span class="text-xl w-5 flex justify-center">${escape_html(item.icon)}</span> ${escape_html(item.label)}</a></li>`);
				} else $$renderer2.push("<!--[!-->");
				$$renderer2.push(`<!--]-->`);
			}
			$$renderer2.push(`<!--]--></ul></div> <div class="mt-10 pt-4 border-t border-white/10"><a href="/me" class="flex items-center gap-3 px-2 py-2 rounded-lg hover:bg-white/10 transition focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-primary"><div class="w-8 h-8 rounded-full bg-secondary text-white flex items-center justify-center text-sm font-medium shrink-0">${escape_html(userInitial)}</div> <div><div class="font-medium">${escape_html(data?.user?.username ?? "Guest")}</div> <div class="text-xs text-gray-400">View profile</div></div></a> <button class="mt-4 inline-block w-full text-center px-3 py-2 rounded-lg bg-accent/20 text-accent text-sm font-medium hover:bg-accent/30 transition focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-accent">Sign out</button></div></nav></aside>`);
		} else $$renderer2.push("<!--[!-->");
		$$renderer2.push(`<!--]--> <main class="flex-grow p-4 sm:p-6 lg:p-8 min-w-0"><div class="p-6 rounded-xl glass transition shadow-xl w-full max-w-full overflow-hidden"><!--[-->`);
		slot($$renderer2, $$props, "default", {});
		$$renderer2.push(`<!--]--></div></main></div> <footer class="border-t border-white/10 glass mt-auto"><div class="max-w-screen-2xl mx-auto px-4 sm:px-6 lg:px-8 py-4 text-sm text-gray-400"><div class="flex items-center justify-between"><div>© ${escape_html((/* @__PURE__ */ new Date()).getFullYear())} AleGYM</div> <div class="flex items-center gap-4"><a href="/privacy" class="hover:text-primary transition focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-primary rounded-sm">Privacy</a></div></div></div></footer></div>`);
		if ($$store_subs) unsubscribe_stores($$store_subs);
		bind_props($$props, { data });
	});
}

//#endregion
export { _layout as default };
//# sourceMappingURL=_layout.svelte-VfyOVXJq.js.map