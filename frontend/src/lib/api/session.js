import { env } from '$env/dynamic/public';
import { resolve } from "$app/paths";
import { request } from "$lib/api/util";

export let create = (username="", password="") => {
    return new Promise((resolve, reject) => {
        request("/session", "POST", JSON.stringify({username, password}), null).then(res => {
            if (res?.session) {
                resolve(res?.session)
            } else if (res?.message) {
                reject(res?.message);
            } else {
                reject(null);
            }
        }).catch(reject);
    })
};

export let refresh = (session) => {
    return new Promise((resolve, reject) => {
        request("/session", "PATCH", null, {"Authorization": `Bearer ${session}`}).then(res => {
            if (res?.session) {
                resolve(res?.session)
            } else if (res?.message) {
                reject(res?.message);
            } else {
                reject(null);
            }
        }).catch(reject);
    })
};

export let remove = (session) => {
    return new Promise((resolve, reject) => {
        request("/session", "DELETE", null, {"Authorization": session}).then(res => {
            if (res?.message) {
                reject(res?.message)
            } else {
                resolve(null);
            }
        }).catch(reject);
    })
};