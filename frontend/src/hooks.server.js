import { env } from '$env/dynamic/private';
import { jwtVerify } from 'jose';

const SECRET = new TextEncoder().encode(env.JWT_SECRET.trim());

export const handle = async ({ event, resolve }) => {
	const token = event.cookies.get("token");
	
	if (token) {
		try {
			const { payload } = await jwtVerify(token, SECRET);
			event.locals.user = payload;
		} catch (err) {
			event.locals.user = null;
			console.warn('JWT validation failed:', err.message);
			event.cookies.delete('session', { path: '/' });
		}
	} else {
		event.locals.user = null;
	}

	return await resolve(event);
};
