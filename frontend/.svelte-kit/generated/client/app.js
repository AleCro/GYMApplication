export { matchers } from './matchers.js';

export const nodes = [
	() => import('./nodes/0'),
	() => import('./nodes/1'),
	() => import('./nodes/2'),
	() => import('./nodes/3'),
	() => import('./nodes/4'),
	() => import('./nodes/5'),
	() => import('./nodes/6'),
	() => import('./nodes/7'),
	() => import('./nodes/8'),
	() => import('./nodes/9'),
	() => import('./nodes/10'),
	() => import('./nodes/11'),
	() => import('./nodes/12'),
	() => import('./nodes/13'),
	() => import('./nodes/14'),
	() => import('./nodes/15')
];

export const server_loads = [0,2,3];

export const dictionary = {
		"/": [4],
		"/app": [6,[2]],
		"/app/calendar": [9,[2]],
		"/app/documentation": [7,[2]],
		"/app/me": [8,[2]],
		"/app/notes": [11,[2]],
		"/app/progress": [12,[2]],
		"/app/user-management": [10,[2]],
		"/app/workouts": [13,[2]],
		"/login": [14,[3]],
		"/logout": [15],
		"/privacy": [5]
	};

export const hooks = {
	handleError: (({ error }) => { console.error(error) }),
	
	reroute: (() => {}),
	transport: {}
};

export const decoders = Object.fromEntries(Object.entries(hooks.transport).map(([k, v]) => [k, v.decode]));

export const hash = false;

export const decode = (type, value) => decoders[type](value);

export { default as root } from '../root.js';