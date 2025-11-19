import { env } from '$env/dynamic/public';

export let request = (path = "/", method = "GET", body = null, headers = null) => {
    return new Promise((resolve, reject) => {
        let req = {method};
        if (headers) {
            req.headers = headers;
        }
        if (body) {
            req.body = body;
        }
        fetch(`${env.PUBLIC_API_URL}${path}`, req).then(res => res.json()).then(resolve).catch(reject);
    })
}