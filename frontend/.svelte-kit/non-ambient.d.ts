
// this file is generated — do not edit it


declare module "svelte/elements" {
	export interface HTMLAttributes<T> {
		'data-sveltekit-keepfocus'?: true | '' | 'off' | undefined | null;
		'data-sveltekit-noscroll'?: true | '' | 'off' | undefined | null;
		'data-sveltekit-preload-code'?:
			| true
			| ''
			| 'eager'
			| 'viewport'
			| 'hover'
			| 'tap'
			| 'off'
			| undefined
			| null;
		'data-sveltekit-preload-data'?: true | '' | 'hover' | 'tap' | 'off' | undefined | null;
		'data-sveltekit-reload'?: true | '' | 'off' | undefined | null;
		'data-sveltekit-replacestate'?: true | '' | 'off' | undefined | null;
	}
}

export {};


declare module "$app/types" {
	export interface AppTypes {
		RouteId(): "/" | "/addgoal" | "/addprogress" | "/calendar" | "/deleteevent" | "/exercise" | "/getgoals" | "/getprogress" | "/goals" | "/login" | "/notes" | "/progress" | "/updategoal";
		RouteParams(): {
			
		};
		LayoutParams(): {
			"/": Record<string, never>;
			"/addgoal": Record<string, never>;
			"/addprogress": Record<string, never>;
			"/calendar": Record<string, never>;
			"/deleteevent": Record<string, never>;
			"/exercise": Record<string, never>;
			"/getgoals": Record<string, never>;
			"/getprogress": Record<string, never>;
			"/goals": Record<string, never>;
			"/login": Record<string, never>;
			"/notes": Record<string, never>;
			"/progress": Record<string, never>;
			"/updategoal": Record<string, never>
		};
		Pathname(): "/" | "/addgoal" | "/addgoal/" | "/addprogress" | "/addprogress/" | "/calendar" | "/calendar/" | "/deleteevent" | "/deleteevent/" | "/exercise" | "/exercise/" | "/getgoals" | "/getgoals/" | "/getprogress" | "/getprogress/" | "/goals" | "/goals/" | "/login" | "/login/" | "/notes" | "/notes/" | "/progress" | "/progress/" | "/updategoal" | "/updategoal/";
		ResolvedPathname(): `${"" | `/${string}`}${ReturnType<AppTypes['Pathname']>}`;
		Asset(): "/app.css" | "/robots.txt" | string & {};
	}
}