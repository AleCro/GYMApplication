import { API_URL } from '$lib/config.js';

/** @type {import('@sveltejs/kit').Handle} */
export async function handle({ event, resolve }) {
    const sessionId = event.cookies.get('session');
    console.log(sessionId);

    try {
        let res = await fetch(API_URL + "/session", {
            method: "POST",
            body: JSON.stringify({ s: sessionId })
        });
        // let x = (res.text());
        
        let resp = await res.json();
        event.locals.user = resp;
        event.locals.user.session = sessionId;
    } catch (e) {}


    return resolve(event);
}