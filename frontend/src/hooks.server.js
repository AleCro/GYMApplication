export async function handle({ event, resolve }) {
    const session = event.cookies.get("session");
    event.locals.user = session ? { id: 1, name: "Alejandre" } : null;
    console.log(session);
    return resolve(event);
}
