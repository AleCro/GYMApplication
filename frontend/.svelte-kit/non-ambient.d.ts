
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
		RouteId(): "/" | "/app" | "/app/calendar" | "/app/documentation" | "/app/me" | "/app/notes" | "/app/progress" | "/app/user-management" | "/app/workouts" | "/login" | "/logout" | "/privacy";
		RouteParams(): {
			
		};
		LayoutParams(): {
			"/": Record<string, never>;
			"/app": Record<string, never>;
			"/app/calendar": Record<string, never>;
			"/app/documentation": Record<string, never>;
			"/app/me": Record<string, never>;
			"/app/notes": Record<string, never>;
			"/app/progress": Record<string, never>;
			"/app/user-management": Record<string, never>;
			"/app/workouts": Record<string, never>;
			"/login": Record<string, never>;
			"/logout": Record<string, never>;
			"/privacy": Record<string, never>
		};
		Pathname(): "/" | "/app" | "/app/" | "/app/calendar" | "/app/calendar/" | "/app/documentation" | "/app/documentation/" | "/app/me" | "/app/me/" | "/app/notes" | "/app/notes/" | "/app/progress" | "/app/progress/" | "/app/user-management" | "/app/user-management/" | "/app/workouts" | "/app/workouts/" | "/login" | "/login/" | "/logout" | "/logout/" | "/privacy" | "/privacy/";
		ResolvedPathname(): `${"" | `/${string}`}${ReturnType<AppTypes['Pathname']>}`;
		Asset(): "/robots.txt" | string & {};
	}
}